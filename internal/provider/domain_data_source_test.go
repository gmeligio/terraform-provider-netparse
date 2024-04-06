package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDomainDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccDomainDataSourceConfigMultipleSubdomain,
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
				Config: testAccDomainDataSourceConfigSingleSubdomain,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.url_domain.test", "domain", "example.com"),
					resource.TestCheckResourceAttr("data.url_domain.test", "host", "foo.example.com"),
					resource.TestCheckResourceAttr("data.url_domain.test", "manager", "ICANN"),
					resource.TestCheckResourceAttr("data.url_domain.test", "sld", "example"),
					resource.TestCheckResourceAttr("data.url_domain.test", "subdomain", "foo"),
					resource.TestCheckResourceAttr("data.url_domain.test", "tld", "com"),
				),
			},
			// {
			// 	Config: testAccDomainDataSourceConfig,
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr("data.url_domain.test", "domain", "example.com"),
			// 		resource.TestCheckResourceAttr("data.url_domain.test", "host", "foo.example.com"),
			// 	),
			// },
		},
	})
}

const testAccDomainDataSourceConfigMultipleSubdomain = `
terraform {
  required_providers {
    url = {
      source = "registry.terraform.io/gmeligio/url"
    }
  }
}

data "url_domain" "test" {
  host = "foo.bar.example.com"
}
`

const testAccDomainDataSourceConfigSingleSubdomain = `
terraform {
  required_providers {
    url = {
      source = "registry.terraform.io/gmeligio/url"
    }
  }
}

data "url_domain" "test" {
  host = "foo.example.com"
}
`
