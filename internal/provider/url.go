package provider

import (
	"net/url"
)

const (
	urlMarkdownDescription                  = "Parses URL components from a URL string."
	urlAttributeMarkdownDescription         = "The URL to parse."
	authorityAttributeMarkdownDescription   = "The concatenation of the username, password, host, and port. It's separated from the scheme by :// . For example: user1:123@example.com:3000 for http://user1:123@example.com:3000 ."
	schemeAttributeMarkdownDescription      = "The protocol used to access the domain. For example: http, https, ftp, sftp, file, etc."
	protocolAttributeMarkdownDescription    = "The concatenation of the scheme and the port. For example: http:, https:, ftp:, sftp:, file:, etc."
	credentialsAttributeMarkdownDescription = "The concatenation of the username and password. For example: user1:123 for https://user1:123@example.com ."
	usernameAttributeMarkdownDescription    = "The first component of the URL credentials. For example: user1 for https://user1:123@example.com ."
	passwordAttributeMarkdownDescription    = "The second component of the URL credentials. For example: 123 for https://user1:123@example.com ."
	hostAttributeMarkdownDescription        = "The domain part of the authority. For example: example.com for https://example.com ."
	portAttributeMarkdownDescription        = "The last component of the URL authority. For example: 443 for https://example.com:443 ."
	pathAttributeMarkdownDescription        = "The URL component after the authority. For example: /path/to/resource for https://example.com/path/to/resource ."
	searchAttributeMarkdownDescription      = "The URL component after the path. For example: ?key=value for https://example.com/path/to/resource?key=value ."
	queryAttributeMarkdownDescription       = "The URL component of the search starting at the ? and before the fragment. For example: key=value for https://example.com/path/to/resource?key=value#section ."
	fragmentAttributeMarkdownDescription    = "The URL component after the search. For example: section for https://example.com/path/to/resource?key=value#section ."
	hashAttributeMarkdownDescription        = "The concatenation of a # with the fragment. For example: #section for https://example.com/path/to/resource?key=value#section ."
)

// urlModel describes the data source data model.
// References used.
// https://registry.terraform.io/modules/matti/urlparse/external/latest
// https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/url_parse
type urlModel struct {
	Url         string
	Authority   string
	Protocol    string
	Scheme      string
	Credentials string
	Username    string
	Password    string
	Host        string
	Port        string
	Path        string
	Search      string
	Query       string
	Hash        string
	Fragment    string
}

func ParseUrl(u string) (*urlModel, error) {
	internalUrl, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	authority := renderAuthority(internalUrl)
	scheme := internalUrl.Scheme
	protocol := scheme + ":"
	credentials := internalUrl.User.String()
	username := internalUrl.User.Username()
	password, _ := internalUrl.User.Password()
	host := internalUrl.Hostname()
	port := internalUrl.Port()
	path := internalUrl.Path
	search := renderSearch(internalUrl)
	query := internalUrl.RawQuery
	fragment := internalUrl.Fragment
	hash := renderHash(internalUrl)

	return &urlModel{
		Url:         u,
		Authority:   authority,
		Protocol:    protocol,
		Scheme:      scheme,
		Credentials: credentials,
		Username:    username,
		Password:    password,
		Host:        host,
		Port:        port,
		Path:        path,
		Search:      search,
		Query:       query,
		Hash:        hash,
		Fragment:    fragment,
	}, nil
}

func (u *urlModel) validate() error {
	return nil
}

func renderHash(u *url.URL) string {
	fragment := u.Fragment
	if fragment == "" {
		return ""
	}

	return "#" + fragment
}

func renderSearch(u *url.URL) string {
	query := u.RawQuery
	if query == "" {
		return ""
	}

	return "?" + query
}

func renderAuthority(u *url.URL) string {
	credentials := u.User.String()
	port := u.Port()

	var credentialsComponent string
	if credentials != "" {
		credentialsComponent = credentials + "@"
	}

	var portComponent string
	if port != "" {
		portComponent = ":" + port
	}

	return credentialsComponent + u.Hostname() + portComponent
}
