package alibabacloudstack

import (
	"fmt"
	"regexp"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/connectivity"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/errmsgs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceAlibabacloudStackDmsEnterpriseUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlibabacloudStackDmsEnterpriseUsersRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				ForceNew:     true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"search_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DELETE", "DISABLE", "NORMAL"}, false),
			},
			"tid": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"names": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mobile": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_uid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"role_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeInt},
						},
						"role_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nick_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAlibabacloudStackDmsEnterpriseUsersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AlibabacloudStackClient)

	action := "ListUsers"
	request := make(map[string]interface{})
	if v, ok := d.GetOk("role"); ok {
		request["Role"] = v
	}
	if v, ok := d.GetOk("search_key"); ok {
		request["SearchKey"] = v
	}
	if v, ok := d.GetOk("status"); ok {
		request["UserState"] = v
	}
	if v, ok := d.GetOk("tid"); ok {
		request["Tid"] = v
	}
	request["PageSize"] = PageSizeLarge
	request["PageNumber"] = 1
	var objects []map[string]interface{}
	var userNameRegex *regexp.Regexp
	if v, ok := d.GetOk("name_regex"); ok {
		r, err := regexp.Compile(v.(string))
		if err != nil {
			return errmsgs.WrapError(err)
		}
		userNameRegex = r
	}

	idsMap := make(map[string]string)
	if v, ok := d.GetOk("ids"); ok {
		for _, vv := range v.([]interface{}) {
			if vv == nil {
				continue
			}
			idsMap[vv.(string)] = vv.(string)
		}
	}

	for {
		response, err := client.DoTeaRequest("POST", "dms-enterprise", "2018-11-01", action, "", nil, nil, request)
		if err != nil {
			return err
		}
		addDebug(action, response, request)

		resp, err := jsonpath.Get("$.UserList.User", response)
		if err != nil {
			return errmsgs.WrapErrorf(err, errmsgs.FailedGetAttributeMsg, action, "$.UserList.User", response)
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})
			if userNameRegex != nil {
				if !userNameRegex.MatchString(fmt.Sprint(item["NickName"])) {
					continue
				}
			}
			if len(idsMap) > 0 {
				if _, ok := idsMap[fmt.Sprint(item["Uid"])]; !ok {
					continue
				}
			}
			objects = append(objects, item)
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}
	ids := make([]string, 0)
	names := make([]interface{}, 0)
	s := make([]map[string]interface{}, 0)
	for _, object := range objects {
		mapping := map[string]interface{}{
			"mobile":     object["Mobile"],
			"parent_uid": formatInt(object["ParentUid"]),
			"role_ids":   object["RoleIdList"].(map[string]interface{})["RoleIds"],
			"role_names": object["RoleNameList"].(map[string]interface{})["RoleNames"],
			"status":     object["State"],
			"id":         fmt.Sprint(object["Uid"]),
			"uid":        fmt.Sprint(object["Uid"]),
			"user_id":    object["UserId"],
			"user_name":  object["NickName"],
			"nick_name":  object["NickName"],
		}
		ids = append(ids, fmt.Sprint(object["Uid"]))
		names = append(names, object["NickName"])
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("ids", ids); err != nil {
		return errmsgs.WrapError(err)
	}

	if err := d.Set("names", names); err != nil {
		return errmsgs.WrapError(err)
	}

	if err := d.Set("users", s); err != nil {
		return errmsgs.WrapError(err)
	}
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		if err := writeToFile(output.(string), s); err != nil {
			return err
		}
	}

	return nil
}
