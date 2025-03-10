package alibabacloudstack

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/connectivity"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/errmsgs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"regexp"
)

func dataSourceAlibabacloudStackNetworkInterfaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlibabacloudstackNetworkInterfacesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				MinItems: 1,
				MaxItems: 100,
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Primary", "Secondary"}, false),
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": tagsSchema(),
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"names": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"interfaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vswitch_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpc_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_ips": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"mac": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": tagsSchema(),
					},
				},
			},
		},
	}
}

func dataSourceAlibabacloudstackNetworkInterfacesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AlibabacloudStackClient)
	request := ecs.CreateDescribeNetworkInterfacesRequest()
	client.InitRpcRequest(*request.RpcRequest)

	if networkInterfaceIds, ok := d.GetOk("ids"); ok {
		ids := expandStringList(networkInterfaceIds.(*schema.Set).List())
		request.NetworkInterfaceId = &ids
	}
	if vpcId, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = vpcId.(string)
	}

	if vswitchId, ok := d.GetOk("vswitch_id"); ok {
		request.VSwitchId = vswitchId.(string)
	}

	if privateIp, ok := d.GetOk("private_ip"); ok {
		request.PrimaryIpAddress = privateIp.(string)
	}

	if securityGroupId, ok := d.GetOk("security_group_id"); ok {
		request.SecurityGroupId = securityGroupId.(string)
	}

	if typ, ok := d.GetOk("type"); ok {
		request.Type = typ.(string)
	}

	if instanceId, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = instanceId.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		var tags []ecs.DescribeNetworkInterfacesTag
		for key, value := range v.(map[string]interface{}) {
			tags = append(tags, ecs.DescribeNetworkInterfacesTag{
				Key:   key,
				Value: value.(string),
			})
		}
		request.Tag = &tags
	}
	var allEnis []ecs.NetworkInterfaceSet
	request.PageSize = requests.NewInteger(PageSizeLarge)
	request.PageNumber = requests.NewInteger(1)
	for {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.DescribeNetworkInterfaces(request)
		})
		response, ok := raw.(*ecs.DescribeNetworkInterfacesResponse)
		if err != nil {
			errmsg := ""
			if ok {
				errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
			}
			return errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, "alibabacloudstack_network_interfaces", request.GetActionName(), errmsgs.AlibabacloudStackGoClientFailure, errmsg)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)

		allEnis = append(allEnis, response.NetworkInterfaceSets.NetworkInterfaceSet...)
		if len(response.NetworkInterfaceSets.NetworkInterfaceSet) < PageSizeLarge {
			break
		}
		page, err := getNextpageNumber(request.PageNumber)
		if err != nil {
			return errmsgs.WrapError(err)
		}
		request.PageNumber = page
	}

	var filterEnis []ecs.NetworkInterfaceSet
	nameRegex, ok := d.GetOk("name_regex")
	if ok && nameRegex.(string) != "" {
		var r *regexp.Regexp
		r = regexp.MustCompile(nameRegex.(string))

		for i := range allEnis {
			if r.MatchString(allEnis[i].NetworkInterfaceName) {
				filterEnis = append(filterEnis, allEnis[i])
			}
		}
	} else {
		filterEnis = allEnis
	}
	return networkInterfaceDescriptionAttributes(d, filterEnis, meta)
}

func networkInterfaceDescriptionAttributes(d *schema.ResourceData, enis []ecs.NetworkInterfaceSet, meta interface{}) error {
	var ids []string
	var names []string
	var s []map[string]interface{}
	for _, eni := range enis {
		var ips []string
		for _, ip := range eni.PrivateIpSets.PrivateIpSet {
			if ip.Primary {
				continue
			}
			ips = append(ips, ip.PrivateIpAddress)
		}
		mapping := map[string]interface{}{
			"id":              eni.NetworkInterfaceId,
			"name":            eni.NetworkInterfaceName,
			"status":          eni.Status,
			"vpc_id":          eni.VpcId,
			"vswitch_id":      eni.VSwitchId,
			"zone_id":         eni.ZoneId,
			"public_ip":       eni.AssociatedPublicIp.PublicIpAddress,
			"private_ip":      eni.PrivateIpAddress,
			"private_ips":     ips,
			"mac":             eni.MacAddress,
			"security_groups": eni.SecurityGroupIds.SecurityGroupId,
			"description":     eni.Description,
			"instance_id":     eni.InstanceId,
			"creation_time":   eni.CreationTime,
			"tags":            ecsTagsToMap(eni.Tags.Tag),
		}

		ids = append(ids, eni.NetworkInterfaceId)
		names = append(names, eni.NetworkInterfaceName)
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("names", names); err != nil {
		return err
	}
	if err := d.Set("ids", ids); err != nil {
		return err
	}
	if err := d.Set("interfaces", s); err != nil {
		return err
	}

	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		if err := writeToFile(output.(string), s); err != nil {
			return err
		}
	}

	return nil
}
