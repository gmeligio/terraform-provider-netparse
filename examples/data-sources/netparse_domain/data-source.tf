data "netparse_domain" "example" {
  host = "foo.bar.example.com"
}

output "domain" {
  value = data.netparse_domain.example

  # {
  #   host = "foo.bar.example.com"
  #   domain = "example.com"
  #   manager = "ICANN"
  #   sld = "example"
  #   subdomain = "foo.bar"
  #   tld = "com"
  # }
}
