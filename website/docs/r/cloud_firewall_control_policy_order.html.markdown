---
subcategory: "Cloud Firewall"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_cloud_firewall_control_policy_order"
sidebar_current: "docs-alibabacloudstack-resource-cloud-firewall-control-policy-order"
description: |-
  Provides a Alibabacloudstack Cloud Firewall Control Policy Order resource.
---

# alibabacloudstack\_cloud\_firewall\_control\_policy\_order

Provides a Cloud Firewall Control Policy Order resource.

For information about Cloud Firewall Control Policy Order and how to use it, see [What is Control Policy Order](https://www.alibabacloud.com/help/doc-detail/138867.htm).

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alibabacloudstack_cloud_firewall_control_policy" "default" {
  direction        = "in"
  application_name = "ANY"
  description      = var.name
  acl_action       = "accept"
  source           = "127.0.0.1/32"
  source_type      = "net"
  destination      = "127.0.0.2/32"
  destination_type = "net"
  proto            = "ANY"
}

resource "alibabacloudstack_cloud_firewall_control_policy_order" "default" {
  acl_uuid  = alibabacloudstack_cloud_firewall_control_policy.default.acl_uuid
  direction = alibabacloudstack_cloud_firewall_control_policy.default.direction
  order     = 1
}
```

## Argument Reference

The following arguments are supported:

* `acl_uuid` - (Required, ForceNew) The unique ID of the access control policy.
* `direction` - (Required, ForceNew) The direction of the traffic to which the access control policy applies. Valid values: `in`, `out`.
* `order` - (Required, Int) The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority. **NOTE:** The value of `-1` indicates the lowest priority.
-> **NOTE:** From version 1.227.1, `order` must be set.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Control Policy Order. It formats as `<acl_uuid>:<direction>`.

## Import

Cloud Firewall Control Policy Order can be imported using the id, e.g.

```shell
$ terraform import alibabacloudstack_cloud_firewall_control_policy_order.example <acl_uuid>:<direction>
```
