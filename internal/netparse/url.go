package netparse

import (
	"net/url"
)

// URLModel describes the data source data model.
// References used.
// https://registry.terraform.io/modules/matti/urlparse/external/latest
// https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/url_parse
type URLModel struct {
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

func ParseURL(u string) (*URLModel, error) {
	internalURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	authority := renderAuthority(internalURL)
	scheme := internalURL.Scheme
	protocol := scheme + ":"
	credentials := internalURL.User.String()
	username := internalURL.User.Username()
	password, _ := internalURL.User.Password()
	host := internalURL.Hostname()
	port := internalURL.Port()
	path := internalURL.Path
	search := renderSearch(internalURL)
	query := internalURL.RawQuery
	fragment := internalURL.Fragment
	hash := renderHash(internalURL)

	return &URLModel{
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

func (u *URLModel) Validate() error {
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

func UrlValidate(u string) error {
	return nil
}
