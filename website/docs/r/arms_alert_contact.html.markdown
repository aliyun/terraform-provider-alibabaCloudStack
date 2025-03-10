---
subcategory: "Application Real-Time Monitoring Service (ARMS)"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_arms_alert_contact"
sidebar_current: "docs-alibabacloudstack-resource-arms-alert-contact"
description: |-
  Provides a Alibabacloudstack Application Real-Time Monitoring Service (ARMS) Alert Contact resource.
---

# alibabacloudstack_arms_alert_contact

Provides a Application Real-Time Monitoring Service (ARMS) Alert Contact resource.

For information about Application Real-Time Monitoring Service (ARMS) Alert Contact and how to use it, see [What is Alert Contact](https://www.alibabacloud.com/help/en/application-real-time-monitoring-service/latest/createalertcontact).


## Example Usage

Basic Usage

```terraform
resource "alibabacloudstack_arms_alert_contact" "example" {
  alert_contact_name     = "example_value"
  ding_robot_webhook_url = "https://oapi.dingtalk.com/robot/send?access_token=91f2f6****"
  email                  = "someone@example.com"
  phone_num              = "1381111****"
}

```

## Argument Reference

The following arguments are supported:

* `alert_contact_name` - (Optional) The name of the alert contact.
* `ding_robot_webhook_url` - (Optional) The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
* `email` - (Optional) The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
* `phone_num` - (Optional) The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
* `system_noc` - (Optional) Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Alert Contact.

## Import

Application Real-Time Monitoring Service (ARMS) Alert Contact can be imported using the id, e.g.

```shell
$ terraform import alibabacloudstack_arms_alert_contact.example <id>
```
