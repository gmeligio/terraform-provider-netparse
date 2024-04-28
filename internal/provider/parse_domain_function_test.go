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


func TestParseDomainFunction_Known(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccParseDomainFunctionConfig_basic("foo.bar.example.com"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("domain"),
						knownvalue.StringExact("example.com"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("manager"),
						knownvalue.StringExact("ICANN"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("sld"),
						knownvalue.StringExact("example"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("subdomain"),
						knownvalue.StringExact("foo.bar"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("tld"),
						knownvalue.StringExact("com"),
					),
				},
			},
		},
	})
}

func TestParseDomainFunction_Null(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::netparse::parse_domain(null)
				}
				`,
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
		},
	})
}

func TestParseDomainFunction_Unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = "example.com"
				}

				output "test" {
					value = provider::netparse::parse_domain(terraform_data.test.output)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("domain"),
						knownvalue.StringExact("example.com"),
					),
				},
			},
		},
	})
}

func testAccParseDomainFunctionConfig_basic(host string) string {
	return fmt.Sprintf(`
output "test" {
	value = provider::netparse::parse_domain(%[1]q)
}
`, host)
}
