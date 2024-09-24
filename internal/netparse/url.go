package netparse

import (
	"net/url"
)

// UrlModel describes the data source data model.
// References used.
// https://registry.terraform.io/modules/matti/urlparse/external/latest
// https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/url_parse
type UrlModel struct {
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

func ParseUrl(u string) (*UrlModel, error) {
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

	return &UrlModel{
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

func (u *UrlModel) Validate() error {
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
