---
subcategory: "VPC"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_vpcs"
sidebar_current: "docs-alibabacloudstack-datasource-vpcs"
description: |-
    Provides a list of VPCs owned by an AlibabacloudStack Cloud account.
---

# alibabacloudstack\_vpcs

This data source provides VPCs available to the user.

## Example Usage

```
data "alibabacloudstack_vpcs" "vpcs_ds" {
  cidr_block = "172.16.0.0/12"
  status     = "Available"
  name_regex = "^foo"
}

output "first_vpc_id" {
  value = "${data.alibabacloudstack_vpcs.vpcs_ds.vpcs.0.id}"
}
```

## Argument Reference

The following arguments are supported:

* `cidr_block` - (Optional) Filter results by a specific CIDR block. For example: "172.16.0.0/12".
* `status` - (Optional) Filter results by a specific status. Valid value are `Pending` and `Available`.
* `name_regex` - (Optional) A regex string to filter VPCs by name.
* `is_default` - (Optional, type: bool) Indicate whether the VPC is the default one in the specified region.
* `vswitch_id` - (Optional) Filter results by the specified VSwitch.
* `tags` - (Optional) A mapping of tags to assign to the resource.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).
* `ids` - (Optional) A list of VPC IDs.
* `resource_group_id` - (Optional) The Id of resource group which VPC belongs.
* `dhcp_options_set_id` - (Optional, ForceNew) The ID of dhcp options set.
* `dry_run` - (Optional, ForceNew) Indicates whether to check this request only. Valid values: `true` and `false`.
* `vpc_name` - (Optional, ForceNew) The name of the VPC.
* `vpc_owner_id` - (Optional, ForceNew) The owner ID of VPC.
* `enable_details` -(Optional) Default to `true`. Set it to true can output the `route_table_id`.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `ids` - A list of VPC IDs.
* `names` - A list of VPC names.
* `vpcs` - A list of VPCs. Each element contains the following attributes:
  * `id` - ID of the VPC.
  * `region_id` - ID of the region where the VPC is located.
  * `resource_group_id` - (Optional) The Id of resource group which VPC belongs.
  * `status` - Status of the VPC.
  * `vpc_name` - Name of the VPC.
  * `vswitch_ids` - List of VSwitch IDs in the specified VPC
  * `cidr_block` - CIDR block of the VPC.
  * `vrouter_id` - ID of the VRouter.
  * `route_table_id` - Route table ID of the VRouter.
  * `description` - Description of the VPC
  * `is_default` - Whether the VPC is the default VPC in the region.
  * `creation_time` - Time of creation.
  * `tags` - A map of tags assigned to the VPC.
  * `ipv6_cidr_block` - The IPv6 CIDR block of the VPC.
  * `router_id` - The ID of the VRouter.
  * `secondary_cidr_blocks` - A list of secondary IPv4 CIDR blocks of the VPC.
  * `user_cidrs` - A list of user CIDRs.
  * `vpc_id` - ID of the VPC.
  * `available_ip_address_count` - available ip count

 
