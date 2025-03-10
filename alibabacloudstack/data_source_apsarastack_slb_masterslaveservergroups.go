package alibabacloudstack

import (
	"regexp"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/connectivity"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/errmsgs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceAlibabacloudStackSlbMasterSlaveServerGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlibabacloudStackSlbMasterSlaveServerGroupsRead,

		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
				ForceNew: true,
				MinItems: 1,
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				ForceNew:     true,
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed values
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"groups": {
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
						"servers": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"weight": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"server_type": {
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

func dataSourceAlibabacloudStackSlbMasterSlaveServerGroupsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AlibabacloudStackClient)

	request := slb.CreateDescribeMasterSlaveServerGroupsRequest()
	client.InitRpcRequest(*request.RpcRequest)
	request.LoadBalancerId = d.Get("load_balancer_id").(string)

	idsMap := make(map[string]string)
	if v, ok := d.GetOk("ids"); ok {
		for _, vv := range v.([]interface{}) {
			idsMap[Trim(vv.(string))] = Trim(vv.(string))
		}
	}

	raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
		return slbClient.DescribeMasterSlaveServerGroups(request)
	})
	response, ok := raw.(*slb.DescribeMasterSlaveServerGroupsResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		return errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, "alibabacloudstack_slb_master_slave_server_groups", request.GetActionName(), errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	var filteredServerGroupsTemp []slb.MasterSlaveServerGroup
	nameRegex, ok := d.GetOk("name_regex")
	if (ok && nameRegex.(string) != "") || (len(idsMap) > 0) {
		var r *regexp.Regexp
		if nameRegex != "" {
			r = regexp.MustCompile(nameRegex.(string))
		}
		for _, serverGroup := range response.MasterSlaveServerGroups.MasterSlaveServerGroup {
			if r != nil && !r.MatchString(serverGroup.MasterSlaveServerGroupName) {
				continue
			}
			if len(idsMap) > 0 {
				if _, ok := idsMap[serverGroup.MasterSlaveServerGroupId]; !ok {
					continue
				}
			}

			filteredServerGroupsTemp = append(filteredServerGroupsTemp, serverGroup)
		}
	} else {
		filteredServerGroupsTemp = response.MasterSlaveServerGroups.MasterSlaveServerGroup
	}

	return slbMasterSlaveServerGroupsDescriptionAttributes(d, filteredServerGroupsTemp, meta)
}

func slbMasterSlaveServerGroupsDescriptionAttributes(d *schema.ResourceData, serverGroups []slb.MasterSlaveServerGroup, meta interface{}) error {
	client := meta.(*connectivity.AlibabacloudStackClient)

	var ids []string
	var names []string
	var s []map[string]interface{}

	for _, serverGroup := range serverGroups {
		mapping := map[string]interface{}{
			"id":   serverGroup.MasterSlaveServerGroupId,
			"name": serverGroup.MasterSlaveServerGroupName,
		}

		request := slb.CreateDescribeMasterSlaveServerGroupAttributeRequest()
		client.InitRpcRequest(*request.RpcRequest)
		request.MasterSlaveServerGroupId = serverGroup.MasterSlaveServerGroupId
		raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
			return slbClient.DescribeMasterSlaveServerGroupAttribute(request)
		})
		if err != nil {
			errmsg := ""
			if response, ok := raw.(*slb.DescribeMasterSlaveServerGroupAttributeResponse); ok {
				errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
			}
			return errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, "alibabacloudstack_slb_master_slave_server_groups", request.GetActionName(), errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		response, _ := raw.(*slb.DescribeMasterSlaveServerGroupAttributeResponse)
		if response != nil && len(response.MasterSlaveBackendServers.MasterSlaveBackendServer) > 0 {
			var backendServerMappings []map[string]interface{}
			for _, backendServer := range response.MasterSlaveBackendServers.MasterSlaveBackendServer {
				backendServerMapping := map[string]interface{}{
					"instance_id": backendServer.ServerId,
					"weight":      backendServer.Weight,
					"server_type": backendServer.ServerType,
				}
				backendServerMappings = append(backendServerMappings, backendServerMapping)
			}
			mapping["servers"] = backendServerMappings
		}

		ids = append(ids, serverGroup.MasterSlaveServerGroupId)
		names = append(names, serverGroup.MasterSlaveServerGroupName)
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("groups", s); err != nil {
		return errmsgs.WrapError(err)
	}
	if err := d.Set("ids", ids); err != nil {
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
