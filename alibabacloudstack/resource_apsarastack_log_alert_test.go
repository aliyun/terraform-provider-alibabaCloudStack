package alibabacloudstack

import (
	"fmt"
	"testing"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAlibabacloudStackLogAlert_basic(t *testing.T) {
	var alert *sls.Alert
	resourceId := "alibabacloudstack_log_alert.default"
	ra := resourceAttrInit(resourceId, logAlertMap)
	serviceFunc := func() interface{} {
		return &LogService{testAccProvider.Meta().(*connectivity.AlibabacloudStackClient)}
	}
	rc := resourceCheckInit(resourceId, &alert, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()

	rand := getAccTestRandInt(1000000, 9999999)
	name := fmt.Sprintf("tf-testacclogalert-%d", rand)
	displayname := fmt.Sprintf("alert_displayname-%d", rand)
	content := "aliyun sls alert test"
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceLogAlertDependence)

	ResourceTest(t, resource.TestCase{

		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"project_name":      "${alibabacloudstack_log_project.default.name}",
					"alert_name":        "alert_name",
					"alert_displayname": displayname,
					"condition":         "count >100",
					"dashboard":         "terraform-dashboard",
					"query_list": []map[string]interface{}{
						{
							"logstore":    "${alibabacloudstack_log_store.default.name}",
							"chart_title": "chart_title",
							"start":       "-60s",
							"end":         "20s",
							"query":       "* AND aliyun",
						},
					},
					"notification_list": []map[string]interface{}{
						{
							"type":        "SMS",
							"mobile_list": []string{"18865521787", "123456678"},
							"content":     content,
						},
						{
							"type":       "Email",
							"email_list": []string{"nihao@alibaba-inc.com", "test@123.com"},
							"content":    content,
						},
						{
							"type":        "DingTalk",
							"service_uri": "www.aliyun.com",
							"content":     content,
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alert_name":        "alert_name",
						"condition":         "count >100",
						"alert_displayname": displayname,
						"dashboard":         "terraform-dashboard",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"throttling": "1h",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"throttling": "1h",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"throttling": "60s",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"throttling": "60s",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"throttling": "10",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"throttling": "10",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"alert_displayname": "update_alert_name_new",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alert_displayname": "update_alert_name_new",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"condition": "count>999",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"condition": "count>999",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"dashboard": "dashboard_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dashboard": "dashboard_update",
					}),
				),
			},

			{
				Config: testAccConfig(map[string]interface{}{
					"alert_displayname": "update_alert_name",
					"condition":         "count<100",
					"dashboard":         "terraform-dashboard-update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alert_displayname": "update_alert_name",
						"condition":         "count<100",
						"dashboard":         "terraform-dashboard-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"query_list": []map[string]interface{}{
						{
							"logstore":    "${alibabacloudstack_log_store.default.name}",
							"chart_title": "chart_title",
							"start":       "-80s",
							"end":         "60s",
							"query":       "* AND aliyun_update",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"query_list.#":       "1",
						"query_list.0.start": "-80s",
						"query_list.0.end":   "60s",
						"query_list.0.query": "* AND aliyun_update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"notification_list": []map[string]interface{}{
						{
							"type":        "SMS",
							"mobile_list": []string{"456456", "456456456"},
							"content":     "updatecontent",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"notification_list.#":         "1",
						"notification_list.0.type":    "SMS",
						"notification_list.0.content": "updatecontent",
					}),
				),
			},
		},
	})
}

var logAlertMap = map[string]string{
	"project_name": CHECKSET,
	"alert_name":   CHECKSET,
}

func resourceLogAlertDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

resource "alibabacloudstack_log_project" "default"{
	name = "${var.name}"
	description = "create by terraform"
}
resource "alibabacloudstack_log_store" "default"{
  	project = "${alibabacloudstack_log_project.default.name}"
  	name = "${var.name}"
  	retention_period = 3650
  	shard_count = 3
  	auto_split = true
  	max_split_shard_count = 60
  	append_meta = true
}
`, name)
}