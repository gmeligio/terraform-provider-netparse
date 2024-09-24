data "netparse_url" "example" {
  url = "https://abc:def@example.com:45/path/to/somewhere?foo=bar&baz=qux#231"
}

output "url" {
  value = data.netparse_url.example

  # {
  #   url         = "https://abc:def@example.com:45/path/to/somewhere?foo=bar&baz=qux#231"
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
