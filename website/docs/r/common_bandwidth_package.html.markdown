---
subcategory: "VPC"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_common_bandwidth_package"
sidebar_current: "docs-alibabacloudstack-resource-common-bandwidth-package"
description: |-
  Provides a Alibabacloudstack Common Bandwidth Package resource.
---

# alibabacloudstack\_common_bandwidth_package

Provides a common bandwidth package resource.

-> **NOTE:** Terraform will auto build common bandwidth package instance while it uses `alibabacloudstack_common_bandwidth_package` to build a common bandwidth package resource.

## Example Usage

Basic Usage

```
resource "alibabacloudstack_common_bandwidth_package" "foo" {
  bandwidth            = "200"
  name                 = "test-common-bandwidth-package"
  description          = "test-common-bandwidth-package"
}
```
## Argument Reference

The following arguments are supported:

* `bandwidth` - (Required) The bandwidth of the common bandwidth package, in Mbps.
* `name` - (Optional) The name of the common bandwidth package.
* `description` - (Optional) The description of the common bandwidth package instance.
* `internet_charge_type` - (Optional, ForceNew, Computed) The billing method of the Internet Shared Bandwidth instance. Set the value to 'PayByTraffic', which specifies the pay-by-data-transfer billing method.
* `ratio` - (Optional, ForceNew, Computed) The percentage of the minimum bandwidth commitment. Set the parameter to 20

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the common bandwidth package instance id.


