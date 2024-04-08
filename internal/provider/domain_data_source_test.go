package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDomainDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDomainDataSource("foo.bar.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.url_domain.test", "domain", "example.com"),
					resource.TestCheckResourceAttr("data.url_domain.test", "host", "foo.bar.example.com"),
					resource.TestCheckResourceAttr("data.url_domain.test", "manager", "ICANN"),
					resource.TestCheckResourceAttr("data.url_domain.test", "sld", "example"),
					resource.TestCheckResourceAttr("data.url_domain.test", "subdomain", "foo.bar"),
					resource.TestCheckResourceAttr("data.url_domain.test", "tld", "com"),
				),
			},
			{
				Config: testAccDomainDataSource("foo.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.url_domain.test", "domain", "example.com"),
					resource.TestCheckResourceAttr("data.url_domain.test", "host", "foo.example.com"),
					resource.TestCheckResourceAttr("data.url_domain.test", "manager", "ICANN"),
					resource.TestCheckResourceAttr("data.url_domain.test", "sld", "example"),
					resource.TestCheckResourceAttr("data.url_domain.test", "subdomain", "foo"),
					resource.TestCheckResourceAttr("data.url_domain.test", "tld", "com"),
				),
			},
			{
				Config: testAccDomainDataSource("example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.url_domain.test", "domain", "example.com"),
					resource.TestCheckResourceAttr("data.url_domain.test", "host", "example.com"),
					resource.TestCheckResourceAttr("data.url_domain.test", "manager", "ICANN"),
					resource.TestCheckResourceAttr("data.url_domain.test", "sld", "example"),
					resource.TestCheckResourceAttr("data.url_domain.test", "subdomain", ""),
					resource.TestCheckResourceAttr("data.url_domain.test", "tld", "com"),
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
