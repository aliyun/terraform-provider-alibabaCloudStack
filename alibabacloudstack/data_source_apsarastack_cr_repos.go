package alibabacloudstack

import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/connectivity"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/errmsgs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceAlibabacloudStackCRRepos() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlibabacloudStackCRReposRead,

		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_details": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Computed values
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"repos": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"summary": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"repo_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"public": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"internal": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceAlibabacloudStackCRReposRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AlibabacloudStackClient)
	request := client.NewCommonRequest("GET", "cr", "2016-06-07", "GetRepoList", "/repos")
	request.Domain = client.Domain

	response, err := client.ProcessCommonRequest(request)
	addDebug(request.GetActionName(), response, request)
	if err != nil {
		if response == nil {
			return errmsgs.WrapErrorf(err, "Process Common Request Failed")
		}
		errmsg := errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		return errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, "alibabacloudstack_cr_namespace", request.GetActionName(), errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}

	repos := crResponseList{}
	log.Printf("response %v", response)
	err = json.Unmarshal(response.GetHttpContentBytes(), &repos)
	log.Printf("unmarshalled response %v", &repos)
	if err != nil {
		return errmsgs.WrapError(err)
	}

	var names []string
	var s []map[string]interface{}

	for _, repo := range repos.Data.Repos {

		if namespace, ok := d.GetOk("namespace"); ok {
			if repo.RepoNamespace != namespace {
				continue
			}
		}

		if nameRegex, ok := d.GetOk("name_regex"); ok {
			r := regexp.MustCompile(nameRegex.(string))
			if !r.MatchString(repo.RepoName) {
				continue
			}
		}
		domainList := make(map[string]string)
		domainList["public"] = repo.RepoDomainList.Public
		domainList["internal"] = repo.RepoDomainList.Internal
		domainList["vpc"] = repo.RepoDomainList.Vpc
		mapping := map[string]interface{}{
			"namespace":   repo.RepoNamespace,
			"name":        repo.RepoName,
			"summary":     repo.Summary,
			"repo_type":   repo.RepoType,
			"domain_list": domainList,
		}

		if detailedEnabled := d.Get("enable_details"); !detailedEnabled.(bool) {
			names = append(names, repo.RepoName)
			s = append(s, mapping)
			continue
		}

		names = append(names, repo.RepoName)
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(names))
	if err := d.Set("repos", s); err != nil {
		return errmsgs.WrapError(err)
	}
	if err := d.Set("ids", names); err != nil {
		return errmsgs.WrapError(err)
	}
	if err := d.Set("names", names); err != nil {
		return errmsgs.WrapError(err)
	}

	// create a json file in current directory and write data source to it.
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		if err := writeToFile(output.(string), s); err != nil {
			return err
		}
	}

	return nil
}
