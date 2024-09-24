package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCIDRDataSource(t *testing.T) {
	resourceFqn := "data.netparse_cidr.test"

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceFqn,
				Config:       testAccCIDRDataSource("192.0.2.1/24"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "ip", "192.0.2.1"),
					resource.TestCheckResourceAttr(resourceFqn, "network", "192.0.2.0/24"),
				),
			},
			{
				ResourceName: resourceFqn,
				Config:       testAccCIDRDataSource("2001:db8:a0b:12f0::1/32"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "ip", "2001:db8:a0b:12f0::1"),
					resource.TestCheckResourceAttr(resourceFqn, "network", "2001:db8::/32"),
				),
			},
		},
	})
}

func testAccCIDRDataSource(cidr string) string {
	return fmt.Sprintf(`
data "netparse_cidr" "test" {
  cidr = %[1]q
}
`, cidr)
}
