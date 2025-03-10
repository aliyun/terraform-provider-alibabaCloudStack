---
subcategory: "ECS"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_disks"
sidebar_current: "docs-alibabacloudstack-datasource-disks"
description: |-
Provides a list of disks to the user.
---

# alibabacloudstack\_disks

This data source provides the disks of the current Apsara Stack Cloud user.

## Example Usage

```
data "alibabacloudstack_disks" "disks_ds" {
  name_regex = "sample_disk"
}

output "first_disk_id" {
  value = "${data.alibabacloudstack_disks.disks_ds.disks.0.id}"
}
```

## Argument Reference

The following arguments are supported:

* `ids` - (Optional) A list of disks IDs.
* `name_regex` - (Optional) A regex string to filter results by disk name.
* `type` - (Optional) Disk type. Possible values: `system` and `data`.
* `category` - (Optional) Disk category. Possible values: `cloud` (basic cloud disk), `cloud_efficiency` (ultra cloud disk), `ephemeral_ssd` (local SSD cloud disk), `cloud_ssd` (SSD cloud disk), and `cloud_essd` (ESSD cloud disk).
* `instance_id` - (Optional) Filter the results by the specified ECS instance ID.
* `tags` - (Optional) A map of tags assigned to the disks. It must be in the format:
  ```
  data "alibabacloudstack_disks" "disks_ds" {
    tags = {
      tagKey1 = "tagValue1",
      tagKey2 = "tagValue2"
    }
  }
  ```
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `disks` - A list of disks. Each element contains the following attributes:
  * `id` - ID of the disk.
  * `name` - Disk name.
  * `description` - Disk description.
  * `region_id` - Region ID the disk belongs to.
  * `availability_zone` - Availability zone of the disk.
  * `status` - Current status. Possible values: `In_use`, `Available`, `Attaching`, `Detaching`, `Creating` and `ReIniting`.
  * `type` - Disk type. Possible values: `system` and `data`.
  * `category` - Disk category. Possible values: `cloud` (basic cloud disk), `cloud_efficiency` (ultra cloud disk), `ephemeral_ssd` (local SSD cloud disk), `cloud_ssd` (SSD cloud disk), and `cloud_essd` (ESSD cloud disk).
  * `size` - Disk size in GiB.
  * `image_id` - ID of the image from which the disk is created. It is null unless the disk is created using an image.
  * `snapshot_id` - Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
  * `instance_id` - ID of the related instance. It is `null` unless the `status` is `In_use`.
  * `kms_key_id` - The ID of the KMS key corresponding to the data disk.
  * `creation_time` - Disk creation time.
  * `attached_time` - Disk attachment time.
  * `detached_time` - Disk detachment time.
  * `storage_set_id` - The ID of the storage set to which the disk belongs.
  * `expiration_time` - Disk expiration time.
  * `tags` - A map of tags assigned to the disk.
 
