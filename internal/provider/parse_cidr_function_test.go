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

func TestParseCIDRFunction_Known(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccParseCIDRFunctionConfig_basic("192.0.2.1/24"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("ip"),
						knownvalue.StringExact("192.0.2.1"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("network"),
						knownvalue.StringExact("192.0.2.0/24"),
					),
				},
			},
			{
				Config: testAccParseCIDRFunctionConfig_basic("2001:db8:a0b:12f0::1/32"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("ip"),
						knownvalue.StringExact("2001:db8:a0b:12f0::1"),
					),
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("network"),
						knownvalue.StringExact("2001:db8::/32"),
					),
				},
			},
		},
	})
}

func TestParseCIDRFunction_Null(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::netparse::parse_cidr(null)
				}
				`,
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
		},
	})
}

func TestParseCIDRFunction_Unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = "192.0.2.1/24"
				}

				output "test" {
					value = provider::netparse::parse_cidr(terraform_data.test.output)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("ip"),
						knownvalue.StringExact("192.0.2.1"),
					),
				},
			},
		},
	})
}

func testAccParseCIDRFunctionConfig_basic(cidr string) string {
	return fmt.Sprintf(`
output "test" {
	value = provider::netparse::parse_cidr(%[1]q)
}
`, cidr)
}
