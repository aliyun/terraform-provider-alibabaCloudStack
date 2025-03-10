package alibabacloudstack

import (
	"regexp"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	r_kvstore "github.com/aliyun/alibaba-cloud-sdk-go/services/r-kvstore"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/connectivity"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/errmsgs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceAlibabacloudStackKVStoreInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlibabacloudStackKVStoreInstancesRead,

		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{string(Normal), string(Creating), string(Changing), string(Inactive)}, false),
			},
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Memcache", "Redis"}, false),
			},
			"tags": tagsSchema(),
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"charge_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"expire_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_class": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"availability_zone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpc_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vswitch_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"capacity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bandwidth": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"connections": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"connection_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAlibabacloudStackKVStoreInstancesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AlibabacloudStackClient)

	request := r_kvstore.CreateDescribeInstancesRequest()
	client.InitRpcRequest(*request.RpcRequest)
	request.InstanceType = d.Get("instance_type").(string)
	request.InstanceStatus = d.Get("status").(string)
	request.PageSize = requests.NewInteger(PageSizeLarge)
	request.PageNumber = requests.NewInteger(1)

	var dbi []r_kvstore.KVStoreInstanceInDescribeInstances

	var nameRegex *regexp.Regexp
	if v, ok := d.GetOk("name_regex"); ok {
		if r, err := regexp.Compile(v.(string)); err == nil {
			nameRegex = r
		}
	}
	idsMap := make(map[string]string)
	if v, ok := d.GetOk("ids"); ok {
		for _, vv := range v.([]interface{}) {
			idsMap[vv.(string)] = vv.(string)
		}
	}
	if v, ok := d.GetOk("tags"); ok {
		var reqTags []r_kvstore.DescribeInstancesTag
		for key, value := range v.(map[string]interface{}) {
			reqTags = append(reqTags, r_kvstore.DescribeInstancesTag{
				Key:   key,
				Value: value.(string),
			})
		}
		request.Tag = &reqTags
	}
	for {
		raw, err := client.WithRkvClient(func(rkvClient *r_kvstore.Client) (interface{}, error) {
			return rkvClient.DescribeInstances(request)
		})
		response, ok := raw.(*r_kvstore.DescribeInstancesResponse)
		if err != nil {
			errmsg := ""
			if ok {
				errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
			}
			return errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, "alibabacloudstack_kvstore_instances", request.GetActionName(), errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		if len(response.Instances.KVStoreInstance) < 1 {
			break
		}

		for _, item := range response.Instances.KVStoreInstance {
			if nameRegex != nil {
				if !nameRegex.MatchString(item.InstanceName) {
					continue
				}
			}
			if len(idsMap) > 0 {
				if _, ok := idsMap[item.InstanceId]; !ok {
					continue
				}
			}
			dbi = append(dbi, item)
		}

		if len(response.Instances.KVStoreInstance) < PageSizeLarge {
			break
		}

		if page, err := getNextpageNumber(request.PageNumber); err != nil {
			return errmsgs.WrapError(err)
		} else {
			request.PageNumber = page
		}
	}

	return kvstoreInstancesDescription(d, dbi)
}

func kvstoreInstancesDescription(d *schema.ResourceData, dbi []r_kvstore.KVStoreInstanceInDescribeInstances) error {

	var ids []string
	var names []string
	var s []map[string]interface{}

	for _, item := range dbi {
		mapping := map[string]interface{}{
			"id":                 item.InstanceId,
			"name":               item.InstanceName,
			"charge_type":        item.ChargeType,
			"instance_type":      item.InstanceType,
			"instance_class":     item.InstanceClass,
			"region_id":          item.RegionId,
			"create_time":        item.CreateTime,
			"expire_time":        item.EndTime,
			"status":             item.InstanceStatus,
			"availability_zone":  item.ZoneId,
			"vpc_id":             item.VpcId,
			"vswitch_id":         item.VSwitchId,
			"private_ip":         item.PrivateIp,
			"port":               item.Port,
			"user_name":          item.UserName,
			"capacity":           item.Bandwidth,
			"bandwidth":          item.Bandwidth,
			"connections":        item.Connections,
			"connection_domain":  item.ConnectionDomain,
		}

		ids = append(ids, item.InstanceId)
		names = append(names, item.InstanceName)
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("ids", ids); err != nil {
		return errmsgs.WrapError(err)
	}
	if err := d.Set("names", names); err != nil {
		return errmsgs.WrapError(err)
	}
	if err := d.Set("instances", s); err != nil {
		return errmsgs.WrapError(err)
	}

	// create a json file in current directory and write data source to it
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		if err := writeToFile(output.(string), s); err != nil {
			return err
		}
	}
	return nil
}
