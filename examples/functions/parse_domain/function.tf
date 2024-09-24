output "domain" {
  value = provider::netparse::parse_domain("foo.bar.example.com")

  # {
  #   host = "foo.bar.example.com"
  #   domain = "example.com"
  #   manager = "ICANN"
  #   sld = "example"
  #   subdomain = "foo.bar"
  #   tld = "com"
  # }
}
