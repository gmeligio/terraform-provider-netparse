package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const resourceFqn = "data.netparse_url.test"

func TestAccURLDataSource_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceFqn,
				Config:       testAccURLDataSourceConfig_basic("://example.com"),
				ExpectError:  regexp.MustCompile("missing protocol scheme"),
			},
			{
				ResourceName: resourceFqn,
				Config:       testAccURLDataSourceConfig_basic("https://abc:def@example.com:45/path/to/somewhere?foo=bar&baz=qux#231"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "authority", "abc:def@example.com:45"),
					resource.TestCheckResourceAttr(resourceFqn, "credentials", "abc:def"),
					resource.TestCheckResourceAttr(resourceFqn, "fragment", "231"),
					resource.TestCheckResourceAttr(resourceFqn, "hash", "#231"),
					resource.TestCheckResourceAttr(resourceFqn, "host", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "password", "def"),
					resource.TestCheckResourceAttr(resourceFqn, "path", "/path/to/somewhere"),
					resource.TestCheckResourceAttr(resourceFqn, "port", "45"),
					resource.TestCheckResourceAttr(resourceFqn, "protocol", "https:"),
					resource.TestCheckResourceAttr(resourceFqn, "query", "foo=bar&baz=qux"),
					resource.TestCheckResourceAttr(resourceFqn, "scheme", "https"),
					resource.TestCheckResourceAttr(resourceFqn, "search", "?foo=bar&baz=qux"),
					resource.TestCheckResourceAttr(resourceFqn, "username", "abc"),
				),
			},
			{
				ResourceName: resourceFqn,

				Config: testAccURLDataSourceConfig_basic("https://example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "authority", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "credentials", ""),
					resource.TestCheckResourceAttr(resourceFqn, "fragment", ""),
					resource.TestCheckResourceAttr(resourceFqn, "hash", ""),
					resource.TestCheckResourceAttr(resourceFqn, "host", "example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "password", ""),
					resource.TestCheckResourceAttr(resourceFqn, "path", ""),
					resource.TestCheckResourceAttr(resourceFqn, "port", ""),
					resource.TestCheckResourceAttr(resourceFqn, "protocol", "https:"),
					resource.TestCheckResourceAttr(resourceFqn, "query", ""),
					resource.TestCheckResourceAttr(resourceFqn, "scheme", "https"),
					resource.TestCheckResourceAttr(resourceFqn, "search", ""),
					resource.TestCheckResourceAttr(resourceFqn, "username", ""),
				),
			},
			{
				ResourceName: resourceFqn,

				Config: testAccURLDataSourceConfig_basic("https://user:password@complex-subdomain.example.com:8080/path/to/resource?query1=value1&query2=value2#Section1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "authority", "user:password@complex-subdomain.example.com:8080"),
					resource.TestCheckResourceAttr(resourceFqn, "credentials", "user:password"),
					resource.TestCheckResourceAttr(resourceFqn, "fragment", "Section1"),
					resource.TestCheckResourceAttr(resourceFqn, "hash", "#Section1"),
					resource.TestCheckResourceAttr(resourceFqn, "host", "complex-subdomain.example.com"),
					resource.TestCheckResourceAttr(resourceFqn, "password", "password"),
					resource.TestCheckResourceAttr(resourceFqn, "path", "/path/to/resource"),
					resource.TestCheckResourceAttr(resourceFqn, "port", "8080"),
					resource.TestCheckResourceAttr(resourceFqn, "protocol", "https:"),
					resource.TestCheckResourceAttr(resourceFqn, "query", "query1=value1&query2=value2"),
					resource.TestCheckResourceAttr(resourceFqn, "scheme", "https"),
					resource.TestCheckResourceAttr(resourceFqn, "search", "?query1=value1&query2=value2"),
					resource.TestCheckResourceAttr(resourceFqn, "username", "user"),
				),
			},
			{
				ResourceName: resourceFqn,
				Config:       testAccURLDataSourceConfig_basic("https://example.org/api/v1/search/%E2%9C%93?query=%F0%9F%92%A9&lang=en#results"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceFqn, "authority", "example.org"),
					resource.TestCheckResourceAttr(resourceFqn, "credentials", ""),
					resource.TestCheckResourceAttr(resourceFqn, "fragment", "results"),
					resource.TestCheckResourceAttr(resourceFqn, "hash", "#results"),
					resource.TestCheckResourceAttr(resourceFqn, "host", "example.org"),
					resource.TestCheckResourceAttr(resourceFqn, "password", ""),
					resource.TestCheckResourceAttr(resourceFqn, "path", "/api/v1/search/✓"),
					resource.TestCheckResourceAttr(resourceFqn, "port", ""),
					resource.TestCheckResourceAttr(resourceFqn, "protocol", "https:"),
					resource.TestCheckResourceAttr(resourceFqn, "query", "query=%F0%9F%92%A9&lang=en"),
					resource.TestCheckResourceAttr(resourceFqn, "scheme", "https"),
					resource.TestCheckResourceAttr(resourceFqn, "search", "?query=%F0%9F%92%A9&lang=en"),
					resource.TestCheckResourceAttr(resourceFqn, "username", ""),
				),
			},
		},
	})
}

func testAccURLDataSourceConfig_basic(host string) string {
	return fmt.Sprintf(`
data "netparse_url" "test" {
  url = %[1]q
}
`, host)
}
