package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDomainDataSource(t *testing.T) {
	resourceFqn := "data.url_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDomainDataSource("foo.bar.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "domain", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "host", "foo.bar.example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "manager", "ICANN"),
					resource.TestCheckResourceAttr(resourceFqn, "sld", "example"),
					resource.TestCheckResourceAttr(resourceFqn, "subdomain", "foo.bar"),
					resource.TestCheckResourceAttr(resourceFqn, "tld", "com"),
				),
			},
			{
				Config: testAccDomainDataSource("foo.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "domain", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "host", "foo.example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "manager", "ICANN"),
					resource.TestCheckResourceAttr(resourceFqn, "sld", "example"),
					resource.TestCheckResourceAttr(resourceFqn, "subdomain", "foo"),
					resource.TestCheckResourceAttr(resourceFqn, "tld", "com"),
				),
			},
			{
				Config: testAccDomainDataSource("example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "domain", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "host", "example.com"),
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
data "url_domain" "test" {
  host = %[1]q
}
`, host)
}
