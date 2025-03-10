---
subcategory: "Auto Scaling(ESS)"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_ess_scaling_configuration"
sidebar_current: "docs-alibabacloudstack-resource-ess-scaling-configuration"
description: |-
  Provides a ESS scaling configuration resource.
---

# alibabacloudstack\_ess\_scaling\_configuration

Provides a ESS scaling configuration resource.

-> **NOTE:** Several instance types have outdated in some regions and availability zones, such as `ecs.t1.*`, `ecs.s2.*`, `ecs.n1.*` and so on. If you want to keep them, you should set `is_outdated` to true. For more about the upgraded instance type, refer to `alibabacloudstack_instance_types` datasource.

## Example Usage

```
variable "name" {
  default = "essscalingconfiguration"
}

data "alibabacloudstack_zones" "default" {
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}

data "alibabacloudstack_instance_types" "default" {
  availability_zone = data.alibabacloudstack_zones.default.zones[0].id
  cpu_core_count    = 2
  memory_size       = 4
}

data "alibabacloudstack_images" "default" {
  name_regex  = "^ubuntu_18.*64"
  most_recent = true
  owners      = "system"
}

resource "alibabacloudstack_vpc" "default" {
  name       = var.name
  cidr_block = "172.16.0.0/16"
}

resource "alibabacloudstack_vswitch" "default" {
  vpc_id            = alibabacloudstack_vpc.default.id
  cidr_block        = "172.16.0.0/24"
  availability_zone = data.alibabacloudstack_zones.default.zones[0].id
  name              = var.name
}

resource "alibabacloudstack_security_group" "default" {
  name   = var.name
  vpc_id = alibabacloudstack_vpc.default.id
}

resource "alibabacloudstack_security_group_rule" "default" {
  type              = "ingress"
  ip_protocol       = "tcp"
  nic_type          = "intranet"
  policy            = "accept"
  port_range        = "22/22"
  priority          = 1
  security_group_id = alibabacloudstack_security_group.default.id
  cidr_ip           = "172.16.0.0/24"
}

resource "alibabacloudstack_ess_scaling_group" "default" {
  min_size           = 1
  max_size           = 1
  scaling_group_name = var.name
  removal_policies   = ["OldestInstance", "NewestInstance"]
  vswitch_ids        = [alibabacloudstack_vswitch.default.id]
}

resource "alibabacloudstack_ess_scaling_configuration" "default" {
  scaling_group_id  = alibabacloudstack_ess_scaling_group.default.id
  image_id          = data.alibabacloudstack_images.default.images[0].id
  instance_type     = data.alibabacloudstack_instance_types.default.instance_types[0].id
  security_group_ids = [alibabacloudstack_security_group.default.id]
  force_delete      = true
  active            = true
}
```
## Argument Reference

The following arguments are supported:

* `scaling_group_id` - (Required, ForceNew) ID of the scaling group of a scaling configuration.
* `image_id` - (Optional) ID of an image file, indicating the image resource selected when an instance is enabled.
* `instance_type` - (Optional) Resource type of an ECS instance.
* `instance_types` - (Optional) Resource types of an ECS instance.
* `instance_name` - (Optional) Name of an ECS instance. Default to "ESS-Instance".
* `is_outdated` - (Optional) Whether to use outdated instance type. Default to false.
* `security_group_ids` - (Required) List IDs of the security group used to create new instances.
* `deployment_set_id` - (Required) ID of Ecs Deployment Set.
* `scaling_configuration_name` - (Optional) Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
* `internet_max_bandwidth_in` - (Optional) Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
* `system_disk_category` - (Optional) Category of the system disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloud_efficiency`.
* `system_disk_size` - (Optional) Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
* `system_disk_auto_snapshot_policy_id` - (Optional) The id of auto snapshot policy for system disk.
* `enable` - (Optional) Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
* `active` - (Optional) Whether active current scaling configuration in the specified scaling group. Default to `false`.
* `substitute` - (Optional) The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
* `host_name` - (Optional) Hostname of an ECS instance.
* `user_data` - (Optional) User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
* `key_name` - (Optional) The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
* `role_name` - (Optional) Instance RAM role name. The name is provided and maintained by RAM. You can use `alibabacloudstack_ram_role` to create a new one.
* `force_delete` - (Optional) The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
* `data_disk` - (Optional) DataDisk mappings to attach to ecs instance. See [Block datadisk](#block-datadisk) below for details.
  * `size` - (Optional) Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk. 
  * `device` - (Optional) The mount point of data disk N. Valid values of N: 1 to 16. If this parameter is not specified, the system automatically allocates a mount point to created ECS instances. The name of the mount point ranges from /dev/xvdb to /dev/xvdz in alphabetical order.
  * `category` - (Optional) Category of data disk. The parameter value options are `ephemeral_ssd`, `cloud_efficiency`, `cloud_ssd` and `cloud`.
  * `snapshot_id` - (Optional) Snapshot used for creating the data disk. If this parameter is specified, the size parameter is neglected, and the size of the created disk is the size of the snapshot. 
  * `delete_with_instance` - (Optional) Whether to delete data disks attached on ecs when release ecs instance. Optional value: `true` or `false`, default to `true`.
  * `encrypted` - (Optional) Specifies whether data disk N is to be encrypted. Valid values of N: 1 to 16. Valid values: `true`: encrypted, `false`: not encrypted. Default value: `false`.
  * `kms_key_id` - (Optional) The CMK ID for data disk N. Valid values of N: 1 to 16.
  * `name` - (Optional) The name of data disk N. Valid values of N: 1 to 16. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
  * `description` - (Optional) The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
  * `auto_snapshot_policy_id` - (Optional) The id of auto snapshot policy for data disk.
* `tags` - (Optional) A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
    - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
    - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
* `zone_id` - (Required) The zone ID of the instance.
* `override` - (Optional) Indicates whether to overwrite the existing data. Default to false.

-> **NOTE:** Before enabling the scaling group, it must have a active scaling configuration.

-> **NOTE:** If the number of attached ECS instances by `instance_ids` is smaller than MinSize, the Auto Scaling Service will automatically create ECS Pay-As-You-Go instance to cater to MinSize. For example, MinSize=5 and 2 existing ECS instances has been attached to the scaling group. When the scaling group is enabled, it will create 3 instnaces automatically based on its current active scaling configuration.

-> **NOTE:** Restrictions on attaching ECS instances:

   - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
   - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
   - The attached ECS instances must in the running state.
   - The attached ECS instances has not been attached to other scaling groups.
   - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.

-> **NOTE:** The last scaling configuration can't be set to inactive and deleted alone.


## Block datadisk

The datadisk mapping supports the following:



## Attributes Reference

The following attributes are exported:

* `id` - The scaling configuration ID.
