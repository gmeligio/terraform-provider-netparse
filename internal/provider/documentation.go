package provider

import "fmt"

var parseDomainMarkdownDescription = describeFunction(domainMarkdownDescription, domainDataSourceTypeName)

const (
	domainMarkdownDescription        = "Parses Public Suffix List properties from a domain. It uses the [publicsuffix](https://pkg.go.dev/golang.org/x/net/publicsuffix) go package to parse the domain. For more details on the domain parts, see [What is a Domain Name?](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_domain_name)."
	domainAttrMarkdownDescription    = "The domain name. It's the tld plus one more label."
	hostAttrMarkdownDescription      = "The host that identifies the domain name."
	managerAttrMarkdownDescription   = "The manager is the entity that manages the domain. It can be one of: ICANN, Private, or None."
	sldAttrMarkdownDescription       = "The second-level domain (SLD) is the label to the left of the effective TLD."
	subdomainAttrMarkdownDescription = "The subdomain is the left part of the host that is not the domain."
	tldAttrMarkdownDescription       = "The effective top-level domain (eTLD) of the domain. This is the public suffix of the domain."
)

var parseURLMarkdownDescription = describeFunction(urlMarkdownDescription, urlDataSourceTypeName)

const (
	urlMarkdownDescription                  = "Parses URL components from a URL string. It uses the [net/url](https://pkg.go.dev/net/url) go package to parse the URL. For more details on the URL components, see [What is a URL?](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_is_a_URL) and [WHATWG URL Standard](https://url.spec.whatwg.org/#api)."
	urlAttributeMarkdownDescription         = "The URL to parse."
	authorityAttributeMarkdownDescription   = "The concatenation of the username, password, host, and port. It's separated from the scheme by `://`."
	credentialsAttributeMarkdownDescription = "The concatenation of the username and password."
	fragmentAttributeMarkdownDescription    = "The component after the search."
	hashAttributeMarkdownDescription        = "The concatenation of a `#` with the fragment."
	hostAttributeMarkdownDescription        = "The domain part of the authority."
	passwordAttributeMarkdownDescription    = "The second component of the credentials."
	protocolAttributeMarkdownDescription    = "The concatenation of the protocol scheme and `:`."
	schemeAttributeMarkdownDescription      = "The protocol scheme used to access the domain."
	usernameAttributeMarkdownDescription    = "The first component of the credentials."
	portAttributeMarkdownDescription        = "The last component of the authority."
	pathAttributeMarkdownDescription        = "The component after the authority."
	queryAttributeMarkdownDescription       = "A substring of the search component, after the `?` and before the fragment."
	searchAttributeMarkdownDescription      = "The component after the path."
)

var parseCIDRMarkdownDescription = describeFunction(cidrMarkdownDescription, cidrDataSourceTypeName)

const (
	cidrMarkdownDescription        = "Parses an IP address and prefix length in CIDR notation. It uses the [net/netip](https://pkg.go.dev/net/netip#Prefix.Masked) go package to parse the CDIR. For more details in CIDR notation, see [RFC 4632](https://rfc-editor.org/rfc/rfc4632.html) and [RFC 4291](https://rfc-editor.org/rfc/rfc4291.html)."
	cidrAttrMarkdownDescription    = "The IP address and prefix length in CIDR notation."
	ipAttrMarkdownDescription      = "The IP address."
	networkAttrMarkdownDescription = "The IP network."
)

const (
	containsIPMarkdownDescription = "Checks if an IP address is within a network."
)

func describeFunction(dataSourceDescription string, functionTypeName string) string {
	return fmt.Sprintf("%s The functionality is equivalent to the `%s` data source.", dataSourceDescription, functionTypeName)
}
