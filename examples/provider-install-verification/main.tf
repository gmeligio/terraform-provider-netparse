terraform {
  required_providers {
    publicsuffix = {
      source = "registry.terraform.io/gmeligio/publicsuffix"
    }
  }
}

provider "publicsuffix" {}

data "publicsuffix_domain" "example" {
  domain = "www.example.com"
}
