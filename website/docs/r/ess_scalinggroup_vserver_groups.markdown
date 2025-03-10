---
subcategory: "Auto Scaling(ESS)"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_ess_scalinggroup_vserver_groups"
sidebar_current: "docs-alibabacloudstack-resource-ess_scalinggroup_vserver_groups"
description: |-
  Provides a ESS Attachment resource to attach or remove vserver groups.
---

# alibabacloudstack\_ess\_scalinggroup\_vserver\_groups

Attaches/Detaches vserver groups to a specified scaling group.

-> **NOTE:** The load balancer of which vserver groups belongs to must be in `active` status.

-> **NOTE:** If scaling group's network type is `VPC`, the vserver groups must be in the same `VPC`.
 
-> **NOTE:** A scaling group can have at most 5 vserver groups attached by default.

-> **NOTE:** Vserver groups and the default group of loadbalancer share the same backend server quota.

-> **NOTE:** When attach vserver groups to scaling group, existing ECS instances will be added to vserver groups; Instead, ECS instances will be removed from vserver group when detach.

-> **NOTE:** Detach action will be executed before attach action.

-> **NOTE:** Vserver group is defined uniquely by `loadbalancer_id`, `vserver_group_id`, `port`.

-> **NOTE:** Modifing `weight` attribute means detach vserver group first and then, attach with new weight parameter.


## Example Usage

```
variable "name" {
  default = "testAccEssVserverGroupsAttachment"
}

data "alibabacloudstack_zones" "default" {
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}

resource "alibabacloudstack_vpc" "default" {
  name       = "${var.name}"
  cidr_block = "172.16.0.0/16"
}

resource "alibabacloudstack_vswitch" "default" {
  vpc_id            = "${alibabacloudstack_vpc.default.id}"
  cidr_block        = "172.16.0.0/24"
  availability_zone = "${data.alibabacloudstack_zones.default.zones.0.id}"
  name              = "${var.name}"
}

resource "alibabacloudstack_slb" "default" {
  name = "${var.name}"
  vswitch_id = "${alibabacloudstack_vswitch.default.id}"
}

resource "alibabacloudstackalibabacloudstack_slb_server_group" "default" {
  load_balancer_id = "${alibabacloudstack_slb.default.id}"
  name = "test"
}
	
resource "alibabacloudstack_slb_listener" "default" {
  count = 2
  load_balancer_id = "${element(alibabacloudstack_slb.default.*.id, count.index)}"
  backend_port = "22"
  frontend_port = "22"
  protocol = "tcp"
  bandwidth = "10"
  health_check_type = "tcp"
}

resource "alibabacloudstack_ess_scaling_group" "default" {
  min_size = "2"
  max_size = "2"
  scaling_group_name = "${var.name}"
  vswitch_ids = ["${alibabacloudstack_vswitch.default.id}"]
  depends_on = ["alibabacloudstack_slb_listener.default"]
}

resource "alibabacloudstack_ess_scalinggroup_vserver_groups" "default" {
  scaling_group_id = "${alibabacloudstack_ess_scaling_group.default.id}"
  vserver_groups {
  loadbalancer_id = "${alibabacloudstack_slb.default.id}"
  vserver_attributes {
    vserver_group_id = "${alibabacloudstack_slb_server_group.default.id}"
    port = "100"
    weight = "60"
    }
  }
}

```

## Argument Reference

The following arguments are supported:

* `scaling_group_id` - (Required) ID of the scaling group.
* `vserver_groups` - (Optional) A list of vserver groups attached on scaling group. See [Block vserver_group](#block-vserver_group) below for details.
  * `loadbalancer_id` - (Required) Loadbalancer server ID of VServer Group.
  * `vserver_attributes` - (Required) A list of VServer Group attributes. See [Block vserver_attribute](#block-vserver_attribute) below for details.
    * `vserver_group_id` - (Required) ID of VServer Group.
    * `port` - (Required) - The port will be used for VServer Group backend server.
    * `weight` - (Required) The weight of an ECS instance attached to the VServer Group.
* `force` - (Optional) If instances of scaling group are attached/removed from slb backend server when attach/detach vserver group from scaling group. Default to true.

## Attributes Reference

The following attributes are exported:

* `id` - (Required, ForceNew) The ESS vserver groups attachment resource ID.
