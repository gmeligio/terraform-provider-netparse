terraform {
  required_providers {
    url = {
      source = "registry.terraform.io/gmeligio/url"
    }
  }
}

provider "url" {}

data "url_components" "example" {
  url = "https://abc:def@example.com:45/path/to/somewhere?foo=bar&baz=qux#231"
}

output "components" {
  value = data.url_components.example
}