---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_adb_account"
sidebar_current: "docs-alibabacloudstack-resource-adb-account"
description: |-
  Provides a ADB account resource.
---

# alibabacloudstack\_adb\_account

Provides a [ADB](https://www.alibabacloud.com/help/product/92664.htm) account resource and used to manage databases.

## Example Usage

```
variable "creation" {
  default = "ADB"
}

variable "name" {
  default = "adbaccountmysql"
}

data "alibabacloudstack_zones" "default" {
  available_resource_creation = var.creation
}

resource "alibabacloudstack_vpc" "default" {
  name       = var.name
  cidr_block = "172.16.0.0/16"
}

resource "alibabacloudstack_vswitch" "default" {
  vpc_id            = alibabacloudstack_vpc.default.id
  cidr_block        = "172.16.0.0/24"
  zone_id           = data.alibabacloudstack_zones.default.zones[0].id
  name              = var.name
}

resource "alibabacloudstack_adb_db_cluster" "cluster" {
  db_cluster_version  = "3.0"
  db_cluster_category = "Cluster"
  db_node_class       = "C8"
  db_node_count       = 2
  db_node_storage     = 200
  pay_type            = "PostPaid"
  vswitch_id          = alibabacloudstack_vswitch.default.id
  description         = var.name
}

resource "alibabacloudstack_adb_account" "account" {
  db_cluster_id       = alibabacloudstack_adb_db_cluster.cluster.id
  account_name        = "tftestnormal"
  account_password    = "Test12345"
  account_description = var.name
}
```

## Argument Reference

The following arguments are supported:

* `account_type` - (Optional, ForceNew) The type of the database account. Default Value: Super. Valid values:
  * Normal: standard account. Up to 256 standard accounts can be created for a cluster.
  * Super: privileged account. Only a single privileged account can be created for a cluster.
* `db_cluster_id` - (Required, ForceNew) The Id of cluster in which account belongs.
* `account_name` - (Required, ForceNew) Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
* `account_password` - (Optional) Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
* `kms_encrypted_password` - (Optional) An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
* `kms_encryption_context` - (Optional) An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
* `account_description` - (Optional) Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.

## Attributes Reference

The following attributes are exported:

* `id` - The current account resource ID. Composed of instance ID and account name with format `<instance_id>:<name>`.

## Import

ADB account can be imported using the id, e.g.

```
$ terraform import alibabacloudstack_adb_account.example "am-12345:tf_account"
```
