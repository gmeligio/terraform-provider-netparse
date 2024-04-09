terraform {
  required_providers {
    netparse = {
      source = "registry.terraform.io/gmeligio/netparse"
    }
  }
}

provider "netparse" {}

data "netparse_url" "example" {
  url = "https://abc:def@example.com:45/path/to/somewhere?foo=bar&baz=qux#231"
}

output "url" {
  value = data.netparse_url.example
}