package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDomainDataSource(t *testing.T) {
	resourceFqn := "data.netparse_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceFqn,
				Config:       testAccDomainDataSource("foo.bar.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "domain", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "manager", "ICANN"),
					resource.TestCheckResourceAttr(resourceFqn, "sld", "example"),
					resource.TestCheckResourceAttr(resourceFqn, "subdomain", "foo.bar"),
					resource.TestCheckResourceAttr(resourceFqn, "tld", "com"),
				),
			},
			{
				ResourceName: resourceFqn,
				Config:       testAccDomainDataSource("foo.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "domain", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "manager", "ICANN"),
					resource.TestCheckResourceAttr(resourceFqn, "sld", "example"),
					resource.TestCheckResourceAttr(resourceFqn, "subdomain", "foo"),
					resource.TestCheckResourceAttr(resourceFqn, "tld", "com"),
				),
			},
			{
				ResourceName: resourceFqn,
				Config:       testAccDomainDataSource("example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "domain", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "manager", "ICANN"),
					resource.TestCheckResourceAttr(resourceFqn, "sld", "example"),
					resource.TestCheckResourceAttr(resourceFqn, "subdomain", ""),
					resource.TestCheckResourceAttr(resourceFqn, "tld", "com"),
				),
			},
		},
	})
}

func testAccDomainDataSource(host string) string {
	return fmt.Sprintf(`
data "netparse_domain" "test" {
  host = %[1]q
}
`, host)
}
