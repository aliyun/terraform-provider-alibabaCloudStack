---
subcategory: "Server Load Balancer (SLB)"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_slb_acl"
sidebar_current: "docs-alibabacloudstack-resource-slb-acl"
description: |-
  Provides a Load Banlancer Access Control List resource.
---

# alibabacloudstack\_slb\_acl

An access control list contains multiple IP addresses or CIDR blocks.
The access control list can help you to define multiple instance listening dimension,
and to meet the multiple usage for single access control list.

Server Load Balancer allows you to configure access control for listeners.
You can configure different whitelists or blacklists for different listeners.

You can configure access control
when you create a listener or change access control configuration after a listener is created.

-> **NOTE:** One access control list can be attached to many Listeners in different load balancer as whitelists or blacklists.

-> **NOTE:** The maximum number of access control lists per region  is 50.

-> **NOTE:** The maximum number of IP addresses added each time is 50.

-> **NOTE:** The maximum number of entries per access control list is 300.

-> **NOTE:** The maximum number of listeners that an access control list can be added to is 50.

For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).

For information about acl and how to use it, see [Configure an access control list](https://www.alibabacloud.com/help/doc-detail/85978.htm).


## Example Usage

```
variable "name" {
  default = "terraformslbaclconfig"
}
variable "ip_version" {
  default = "ipv4"
}

resource "alibabacloudstack_slb_acl" "default" {
  name       = "${var.name}"
  ip_version = "${var.ip_version}"
  entry_list {
    entry   = "10.10.10.0/24"
    comment = "first"
  }
  entry_list {
    entry   = "168.10.10.0/24"
    comment = "second"
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the access control list.
* `ip_version` - (Optional, ForceNew) The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
* `entry_list` - (Optional) A list of entry (IP addresses or CIDR blocks) to be added. At most 50 etnry can be supported in one resource. It contains two sub-fields as `Entry Block` follows.The entry mapping supports the following:
  * `entry` - (Required) An IP addresses or CIDR blocks.
  * `comment` - (Optional) the comment of the entry.
* `tags` - (Optional, Available in v1.66.0+) A mapping of tags to assign to the resource.
* `resource_group_id` - (Optional, ForceNew, Available in v1.67.0+) Resource group ID.


## Attributes Reference

The following attributes are exported:

* `id` - The Id of the access control list.
