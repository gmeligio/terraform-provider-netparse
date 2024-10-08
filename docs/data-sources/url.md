---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netparse_url Data Source - netparse"
subcategory: ""
description: |-
  Parses URL components from a URL string. It uses the net/url https://pkg.go.dev/net/url go package to parse the URL. For more details on the URL components, see What is a URL? https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_is_a_URL and WHATWG URL Standard https://url.spec.whatwg.org/#api.
---

# netparse_url (Data Source)

Parses URL components from a URL string. It uses the [net/url](https://pkg.go.dev/net/url) go package to parse the URL. For more details on the URL components, see [What is a URL?](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_is_a_URL) and [WHATWG URL Standard](https://url.spec.whatwg.org/#api).

## Example Usage

```terraform
# Get the host from the URL
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

# Then get the domain from the host
data "netparse_domain" "example" {
  host = data.netparse_url.example.host
}

output "domain" {
  value = data.netparse_domain.example

  # {
  #   host      = "example.com"
  #   domain    = "example.com"
  #   manager   = "ICANN"
  #   sld       = "example"
  #   subdomain = ""
  #   tld       = "com"
  # }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `url` (String) The URL to parse.

### Read-Only

- `authority` (String) The concatenation of the username, password, host, and port. It's separated from the scheme by `://`.
- `credentials` (String) The concatenation of the username and password.
- `fragment` (String) The component after the search.
- `hash` (String) The concatenation of a `#` with the fragment.
- `host` (String) The domain part of the authority.
- `password` (String, Sensitive) The second component of the credentials.
- `path` (String) The component after the authority.
- `port` (String) The last component of the authority.
- `protocol` (String) The concatenation of the protocol scheme and `:`.
- `query` (String) A substring of the search component, after the `?` and before the fragment.
- `scheme` (String) The protocol scheme used to access the domain.
- `search` (String) The component after the path.
- `username` (String, Sensitive) The first component of the credentials.
