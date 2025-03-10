package alibabacloudstack

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAlibabacloudStackZonesDataSource_basic(t *testing.T) {
	ResourceTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlibabacloudStackZonesDataSourceBasicConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.foo"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.id"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.local_name"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_instance_types.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_resource_creation.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_disk_categories.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "ids.#"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.slb_slave_zone_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccAlibabacloudStackZonesDataSource_filter(t *testing.T) {

	ResourceTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlibabacloudStackZonesDataSourceFilter,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.foo"),
					testCheckZoneLength("data.alibabacloudstack_zones.foo"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.id"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.local_name"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_instance_types.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_resource_creation.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_disk_categories.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "ids.#"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.slb_slave_zone_ids.#", "0"),
				),
			},

			{
				Config: testAccCheckAlibabacloudStackZonesDataSourceFilterIoOptimized,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.foo"),
					testCheckZoneLength("data.alibabacloudstack_zones.foo"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.id"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.local_name"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_instance_types.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_resource_creation.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_disk_categories.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "ids.#"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.slb_slave_zone_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccAlibabacloudStackZonesDataSource_unitRegion(t *testing.T) {
	ResourceTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlibabacloudStackZonesDataSourceUnitRegion,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.foo"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.id"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.local_name"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_instance_types.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_resource_creation.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.available_disk_categories.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "ids.#"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.slb_slave_zone_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccAlibabacloudStackZonesDataSource_multiZone(t *testing.T) {
	ResourceTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlibabacloudStackZonesDataSourceMultiZone,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.default"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.id"),
					//resource.TestMatchResourceAttr("data.alibabacloudstack_zones.default", "zones.0.id", regexp.MustCompile(fmt.Sprintf(".%s.", MULTI_IZ_SYMBOL))),
					//resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.local_name", "a"),
					//resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.available_instance_types.#", "0"),
					//resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.available_resource_creation.#", "0"),
					//resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.available_disk_categories.#", "0"),

					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.local_name"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_instance_types.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_resource_creation.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_disk_categories.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.multi_zone_ids.#"),
					//resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.multi_zone_ids.0"),
					//resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.multi_zone_ids.1"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "ids.#"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.slb_slave_zone_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccAlibabacloudStackZonesDataSource_chargeType(t *testing.T) {
	ResourceTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlibabacloudStackZonesDataSourceChargeType,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.default"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.id"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.local_name"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_instance_types.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_resource_creation.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_disk_categories.#"),
					//resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.local_name", ""),
					//resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.available_instance_types.#", "0"),
					//resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.available_resource_creation.#", "0"),
					//resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.available_disk_categories.#", "0"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "ids.#"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.0.slb_slave_zone_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccAlibabacloudStackZonesDataSource_slb(t *testing.T) {
	ResourceTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlibabacloudStackZonesDataSource_slb,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.default"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.id"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.local_name"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_instance_types.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_resource_creation.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.available_disk_categories.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "ids.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "zones.0.slb_slave_zone_ids.#"),
				),
			},
		},
	})
}

func TestAccAlibabacloudStackZonesDataSource_enable_details(t *testing.T) {
	ResourceTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlibabacloudStackZonesDataSourceEnableDetails,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.foo"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.#"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "zones.0.id"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.local_name", ""),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.available_instance_types.#", "0"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.available_resource_creation.#", "0"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.available_disk_categories.#", "0"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.foo", "ids.#"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.foo", "zones.0.slb_slave_zone_ids.#", "0"),
				),
			},
		},
	})
}
func TestAccAlibabacloudStackZonesDataSource_empty(t *testing.T) {
	ResourceTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlibabacloudStackZonesDataSourceEmpty,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlibabacloudStackDataSourceID("data.alibabacloudstack_zones.default"),
					resource.TestCheckResourceAttr("data.alibabacloudstack_zones.default", "zones.#", "0"),
					resource.TestCheckNoResourceAttr("data.alibabacloudstack_zones.default", "zones.id"),
					resource.TestCheckNoResourceAttr("data.alibabacloudstack_zones.default", "zones.local_name"),
					resource.TestCheckNoResourceAttr("data.alibabacloudstack_zones.default", "zones.available_instance_types"),
					resource.TestCheckNoResourceAttr("data.alibabacloudstack_zones.default", "zones.available_resource_creation"),
					resource.TestCheckNoResourceAttr("data.alibabacloudstack_zones.default", "zones.available_disk_categories"),
					resource.TestCheckResourceAttrSet("data.alibabacloudstack_zones.default", "ids.#"),
				),
			},
		},
	})
}

// the zone length changed occasionally
// check by range to avoid test case failure
func testCheckZoneLength(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ms := s.RootModule()
		rs, ok := ms.Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s", name)
		}

		i, err := strconv.Atoi(is.Attributes["zones.#"])

		if err != nil {
			return fmt.Errorf("convert zone length err: %#v", err)
		}

		if i <= 0 {
			return fmt.Errorf("zone length expected greater than 0 got err: %d", i)
		}

		return nil
	}
}

const testAccCheckAlibabacloudStackZonesDataSourceBasicConfig = `
data "alibabacloudstack_zones" "foo" {
	enable_details = true
}
`

const testAccCheckAlibabacloudStackZonesDataSourceFilter = `
data "alibabacloudstack_zones" "foo" {
	available_resource_creation= "VSwitch"
	available_disk_category= "cloud_efficiency"
	enable_details = true
}
`

const testAccCheckAlibabacloudStackZonesDataSourceFilterIoOptimized = `
data "alibabacloudstack_zones" "foo" {
	available_resource_creation= "IoOptimized"
	available_disk_category= "cloud_efficiency"
	enable_details = true
}
`

const testAccCheckAlibabacloudStackZonesDataSourceUnitRegion = `
data "alibabacloudstack_zones" "foo" {
	available_resource_creation= "VSwitch"
	enable_details = true
}
`

const testAccCheckAlibabacloudStackZonesDataSourceMultiZone = `
data "alibabacloudstack_zones" "default" {
  available_resource_creation= "Rds"
  multi = true
  enable_details = true
}`

const testAccCheckAlibabacloudStackZonesDataSourceChargeType = `
data "alibabacloudstack_zones" "default" {
  instance_charge_type = "PrePaid"
  available_resource_creation= "Rds"
  multi = true
  enable_details = true
}`

const testAccCheckAlibabacloudStackZonesDataSource_slb = `
data "alibabacloudstack_zones" "default" {
  available_resource_creation= "Slb"
  enable_details = true
  available_slb_address_ip_version= "ipv4"
  available_slb_address_type="Vpc"
}`

const testAccCheckAlibabacloudStackZonesDataSourceEnableDetails = `
data "alibabacloudstack_zones" "foo" {}
`
const testAccCheckAlibabacloudStackZonesDataSourceEmpty = `
data "alibabacloudstack_zones" "default" {
  available_instance_type = "ecs.n1.fake"
}
`
