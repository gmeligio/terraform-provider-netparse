locals {
  # Get the host from the URL
  url = provider::netparse::parse_url("https://abc:def@example.com:45/path/to/somewhere?foo=bar&baz=qux#231s")

  # {
  #   host = "foo.bar.example.com"
  #   domain = "example.com"
  #   manager = "ICANN"
  #   sld = "example"
  #   subdomain = "foo.bar"
  #   tld = "com"
  # }
}

# Then get the domain from the host
output "domain" {
  value = provider::netparse::parse_domain(url.host)

  # {
  #   host = "example.com"
  #   domain = "example.com"
  #   manager = "ICANN"
  #   sld = "example"
  #   subdomain = ""
  #   tld = "com"
  # }
}
