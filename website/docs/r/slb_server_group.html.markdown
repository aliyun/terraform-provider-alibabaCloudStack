---
subcategory: "Server Load Balancer (SLB)"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_slb_server_group"
sidebar_current: "docs-alibabacloudstack-resource-slb-server-group"
description: |-
  Provides a Load Balancer Virtual Backend Server Group resource.
---

# alibabacloudstack\_slb\_server\_group

A virtual server group contains several ECS instances. The virtual server group can help you to define multiple listening dimension,
and to meet the personalized requirements of domain name and URL forwarding.

-> **NOTE:** One ECS instance can be added into multiple virtual server groups.

-> **NOTE:** One virtual server group can be attached with multiple listeners in one load balancer.

-> **NOTE:** One Classic and Internet load balancer, its virtual server group can add Classic and VPC ECS instances.

-> **NOTE:** One Classic and Intranet load balancer, its virtual server group can only add Classic ECS instances.

-> **NOTE:** One VPC load balancer, its virtual server group can only add the same VPC ECS instances.

## Example Usage

```
variable "name" {
  default = "slbservergroupvpc"
}
data "alibabacloudstack_zones" "default" {
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}
data "alibabacloudstack_instance_types" "default" {
  availability_zone = "${data.alibabacloudstack_zones.default.zones.0.id}"
  cpu_core_count    = 1
  memory_size       = 2
}
data "alibabacloudstack_images" "default" {
  name_regex  = "^ubuntu_18.*64"
  most_recent = true
  owners      = "system"
}
resource "alibabacloudstack_vpc" "default" {
  name       = "${var.name}"
  cidr_block = "172.16.0.0/16"
}
resource "alibabacloudstack_vswitch" "default" {
  vpc_id            = "${alibabacloudstack_vpc.default.id}"
  cidr_block        = "172.16.0.0/16"
  availability_zone = "${data.alibabacloudstack_zones.default.zones.0.id}"
  name              = "${var.name}"
}
resource "alibabacloudstack_security_group" "default" {
  name   = "${var.name}"
  vpc_id = "${alibabacloudstack_vpc.default.id}"
}
resource "alibabacloudstack_instance" "instance" {
  image_id                   = "${data.alibabacloudstack_images.default.images.0.id}"
  instance_type              = "${data.alibabacloudstack_instance_types.default.instance_types.0.id}"
  instance_name              = "${var.name}"
  count                      = "2"
  security_groups            = "${alibabacloudstack_security_group.default.*.id}"
  internet_max_bandwidth_out = "10"
  availability_zone          = "${data.alibabacloudstack_zones.default.zones.0.id}"
  system_disk_category       = "cloud_efficiency"
  vswitch_id                 = "${alibabacloudstack_vswitch.default.id}"
}
resource "alibabacloudstack_slb" "default" {
  name       = "${var.name}"
  vswitch_id = "${alibabacloudstack_vswitch.default.id}"
}
resource "alibabacloudstack_slb_server_group" "default" {
  load_balancer_id = "${alibabacloudstack_slb.default.id}"
  name             = "${var.name}"
  servers {
    server_ids = ["${alibabacloudstack_instance.instance.0.id}", "${alibabacloudstack_instance.instance.1.id}"]
    port       = 100
    weight     = 10
  }
  servers {
    server_ids = ["${alibabacloudstack_instance.instance.*.id}"]
    port       = 80
    weight     = 100
  }
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required, ForceNew) The Load Balancer ID which is used to launch a new virtual server group.
* `name` - (Optional) Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
* `delete_protection_validation` - (Optional) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
* `servers` - A list of ECS instances to be added. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows. The servers mapping supports the following:
  * `server_ids` - (Required) A list backend server ID (ECS instance ID).
  * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
  * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
  * `type` - (Optional) Type of the backend server. Valid value ecs, eni. Default to eni.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the virtual server group.
* `load_balancer_id` - The Load Balancer ID which is used to launch a new virtual server group.
* `name` - The name of the virtual server group.
* `servers` - A list of ECS instances that have be added.
