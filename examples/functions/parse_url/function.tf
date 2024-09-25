locals {
  # Get the host from the URL
  url = provider::netparse::parse_url("https://abc:def@example.com:45/path/to/somewhere?foo=bar&baz=qux#231s")

  # {
  #   authority   = "abc:def@example.com:45"
  #   credentials = "abc:def"
  #   fragment    = "231"
  #   hash        = "#231"
  #   host        = "example.com"
  #   password    = "def"
  #   path        = "/path/to/somewhere"
  #   port        = 45
  #   protocol    = "https:"
  #   query       = "foo=bar&baz=qux"
  #   scheme      = "https"
  #   search      = "?foo=bar&baz=qux"
  #   username    = "abc"
  # }
}

# Then get the domain from the host
output "domain" {
  value = provider::netparse::parse_domain(url.host)

  # {
  #   host      = "example.com"
  #   domain    = "example.com"
  #   manager   = "ICANN"
  #   sld       = "example"
  #   subdomain = ""
  #   tld       = "com"
  # }
}
