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
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestContainsIPFunction_Known(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContainsIPFunctionConfig_basic("192.0.2.0/24", "192.0.2.3"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.Bool(true),
					),
				},
			},
			{
				Config: testAccContainsIPFunctionConfig_basic("192.0.2.0/24", "192.1.0.0"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.Bool(false),
					),
				},
			},
			{
				Config: testAccContainsIPFunctionConfig_basic("2001:db8::/32", "2001:db8:a0b:12f0::5"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.Bool(true),
					),
				},
			},
		},
	})
}

func TestContainsIPFunction_Null(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::netparse::contains_ip(null, "192.1.0.0")
				}
				`,
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
			{
				Config: `
				output "test" {
					value = provider::netparse::contains_ip("192.0.2.0/24", null)
				}
				`,
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
		},
	})
}

func TestContainsIPFunction_Unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = {
						network = "192.0.2.0/24"
						ip      = "192.0.2.4"
					}
				}

				output "test" {
					value = provider::netparse::contains_ip(terraform_data.test.output.network, terraform_data.test.output.ip)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.Bool(true),
					),
				},
			},
		},
	})
}

func testAccContainsIPFunctionConfig_basic(network string, ip string) string {
	return fmt.Sprintf(`
output "test" {
	value = provider::netparse::contains_ip(%[1]q, %[2]q)
}
`, network, ip)
}
