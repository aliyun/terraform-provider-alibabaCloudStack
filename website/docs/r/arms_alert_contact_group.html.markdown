---
subcategory: "Application Real-Time Monitoring Service (ARMS)"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_arms_alert_contact_group"
sidebar_current: "docs-alibabacloudstack-resource-arms-alert-contact-group"
description: |-
  Provides a Alibabacloudstack Application Real-Time Monitoring Service (ARMS) Alert Contact Group resource.
---

# alibabacloudstack_arms_alert_contact_group

Provides a Application Real-Time Monitoring Service (ARMS) Alert Contact Group resource.

For information about Application Real-Time Monitoring Service (ARMS) Alert Contact Group and how to use it, see [What is Alert Contact Group](https://www.alibabacloud.com/help/en/doc-detail/130677.htm).


## Example Usage

Basic Usage

```terraform
resource "alibabacloudstack_arms_alert_contact" "example" {
  alert_contact_name     = "example_value"
  ding_robot_webhook_url = "https://oapi.dingtalk.com/robot/send?access_token=91f2f6****"
  email                  = "someone@example.com"
  phone_num              = "1381111****"
}
resource "alibabacloudstack_arms_alert_contact_group" "example" {
  alert_contact_group_name = "example_value"
  contact_ids              = [alibabacloudstack_arms_alert_contact.example.id]
}

```

## Argument Reference

The following arguments are supported:

* `alert_contact_group_name` - (Required) The name of the resource.
* `contact_ids` - (Optional) The list id of alert contact.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Alert Contact Group.

## Import

Application Real-Time Monitoring Service (ARMS) Alert Contact Group can be imported using the id, e.g.

```shell
$ terraform import alibabacloudstack_arms_alert_contact_group.example <id>
```
