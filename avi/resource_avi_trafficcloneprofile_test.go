package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVITrafficCloneProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVITrafficCloneProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVITrafficCloneProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITrafficCloneProfileExists("avi_trafficcloneprofile.testtrafficcloneprofile"),
					resource.TestCheckResourceAttr(
						"avi_trafficcloneprofile.testtrafficcloneprofile", "name", "tp-test")),
			},
			{
				Config: testAccUpdatedAVITrafficCloneProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVITrafficCloneProfileExists("avi_trafficcloneprofile.testtrafficcloneprofile"),
					resource.TestCheckResourceAttr(
						"avi_trafficcloneprofile.testtrafficcloneprofile", "name", "tp-abc")),
			},
		},
	})

}

func testAccCheckAVITrafficCloneProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Traffic Clone Profile ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVITrafficCloneProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_trafficcloneprofile" {
			continue
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI Traffic Clone Profile still exists")
		}
	}
	return nil
}

const testAccAVITrafficCloneProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

resource "avi_trafficcloneprofile" "testtrafficcloneprofile" {
	name = "tp-test"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
}
`

const testAccUpdatedAVITrafficCloneProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

resource "avi_trafficcloneprofile" "testtrafficcloneprofile" {
	name = "tp-abc"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cloud_ref= "${data.avi_cloud.default_cloud.id}"
}
`
