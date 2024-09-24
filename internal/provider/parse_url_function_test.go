// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestParseURLFunction_Known(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccParseURLFunctionConfig_basic("https://abc:def@example.com:45/path/to/somewhere?foo=bar&baz=qux#231"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("authority"),
						knownvalue.StringExact("abc:def@example.com:45"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("credentials"),
						knownvalue.StringExact("abc:def"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("fragment"),
						knownvalue.StringExact("231"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("hash"),
						knownvalue.StringExact("#231"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("host"),
						knownvalue.StringExact("example.com"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("password"),
						knownvalue.StringExact("def"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("path"),
						knownvalue.StringExact("/path/to/somewhere"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("port"),
						knownvalue.StringExact("45"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("protocol"),
						knownvalue.StringExact("https:"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("query"),
						knownvalue.StringExact("foo=bar&baz=qux"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("scheme"),
						knownvalue.StringExact("https"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("search"),
						knownvalue.StringExact("?foo=bar&baz=qux"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("username"),
						knownvalue.StringExact("abc"),
					),
				},
			},
		},
	})
}

func TestParseURLFunction_Null(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::netparse::parse_url(null)
				}
				`,
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
		},
	})
}

func TestParseURLFunction_Unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = "https://example.com"
				}

				output "test" {
					value = provider::netparse::parse_url(terraform_data.test.output)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("authority"),
						knownvalue.StringExact("example.com"),
					),
				},
			},
		},
	})
}

func testAccParseURLFunctionConfig_basic(host string) string {
	return fmt.Sprintf(`
output "test" {
	value = provider::netparse::parse_url(%[1]q)
}
`, host)
}
