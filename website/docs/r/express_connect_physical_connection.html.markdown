---
subcategory: "Express Connect"
layout: "alicloud"
page_title: "Alibabacloudstack: alibabacloudstack_express_connect_physical_connection"
sidebar_current: "docs-alibabacloudstack-resource-express-connect-physical-connection"
description: |-
  Provides a Alibabacloudstack Express Connect Physical Connection resource.
---

# alibabacloudstack\_express\_connect\_physical\_connection

Provides a Express Connect Physical Connection resource.

For information about Express Connect Physical Connection and how to use it, see [What is Physical Connection](https://www.alibabacloud.com/help/doc-detail/44852.htm).

## Example Usage

Basic Usage

```terraform
resource "alibabacloudstack_express_connect_physical_connection" "domestic" {
  device_name              = "express_connect_foo"
  access_point_id          = "ap-cn-hangzhou-yh-B"
  line_operator            = "CT"
  peer_location            = "example_value"
  physical_connection_name = "example_value"
  type                     = "VPC"
  description              = "my domestic connection"
  port_type                = "1000Base-LX"
  bandwidth                = 100
}

```

## Argument Reference

The following arguments are supported:

* `device_name` - (Required) The Physical Device Name.
* `access_point_id` - (Required, ForceNew) The Physical Leased Line Access Point ID.
* `bandwidth` - (Optional, Computed) On the Bandwidth of the ECC Service and Physical Connection.
* `circuit_code` - (Optional) Operators for Physical Connection Circuit Provided Coding.
* `description` - (Optional) The Physical Connection to Which the Description.
* `line_operator` - (Required) Provides Access to the Physical Line Operator. Valid values:
  * CT: China Telecom
  * CU: China Unicom
  * CM: china Mobile
  * CO: Other Chinese
  * Equinix: Equinix
  * Other: Other Overseas.

* `peer_location` - (Required) and an on-Premises Data Center Location.
* `physical_connection_name` - (Optional) on Behalf of the Resource Name of the Resources-Attribute Field.
* `port_type` - (Optional) The Physical Leased Line Access Port Type. Valid value:
  * 100Base-T: Fast Electrical Ports
  * 1000Base-T: gigabit Electrical Ports
  * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
  * 10GBase-T: Gigabit Electrical Port
  * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
  * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
  * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.

**NOTE:** From in v1.185.0+, The `40GBase-LR` and `100GBase-LR` is valid. and Set these values based on the water levels of background ports. For details about the water levels, contact the business manager.

* `redundant_physical_connection_id` - (Optional) Redundant Physical Connection to Which the ID.
* `status` - (Optional, Computed) Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
* `type` - (Optional, Computed, ForceNew) Physical Private Line of Type. Default Value: VPC.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Physical Connection.

## Import

Express Connect Physical Connection can be imported using the id, e.g.

```
$ terraform import alibabacloudstack_express_connect_physical_connection.example <id>
```
