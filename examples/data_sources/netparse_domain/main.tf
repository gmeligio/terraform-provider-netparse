terraform {
  required_providers {
    netparse = {
      source = "registry.terraform.io/gmeligio/netparse"
    }
  }
}

provider "netparse" {}

data "netparse_domain" "example" {
  host = "example.com"
}

output "domain" {
  value = data.netparse_domain.example
}