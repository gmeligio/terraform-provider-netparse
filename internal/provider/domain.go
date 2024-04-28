package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/net/publicsuffix"
)

const (
	domainMarkdownDescription        = "Parses Public Suffix List properties from a domain"
	domainAttrMarkdownDescription    = "The domain name. It's the tld plus one more label. For example: example.com for host foo.example.com"
	hostAttrMarkdownDescription      = "The host that identifies the domain name"
	managerAttrMarkdownDescription   = "The manager is the entity that manages the domain. It can be one of the following: ICANN, Private, or None."
	sldAttrMarkdownDescription       = "The second-level domain (SLD) is the label to the left of the effective TLD. For example: example for example.com, or foo for foo.co.uk"
	subdomainAttrMarkdownDescription = "The subdomain is the left part of the host that is not the domain. For example: www for www.example.com, mail for mail.foo.org, blog for blog.bar.org"
	tldAttrMarkdownDescription       = "The effective top-level domain (eTLD) of the domain. This is the public suffix of the domain. For example: com for example.com, or co.uk for foo.co.uk"
)

// domainDataSourceModel describes the data source data model.
// References used.
// https://github.com/lupomontero/psl
// https://github.com/jpillora/go-tld
// https://github.com/zomasec/tld
// https://github.com/bobesa/go-domain-util
// https://github.com/joeguo/tldextract
type domainModel struct {
	Domain    string
	Host      string
	Manager   string
	SLD       string
	Subdomain string
	TLD       string
}

func ParseDomain(h string) (*domainModel, error) {
	host := h
	eTLD, icann := publicsuffix.PublicSuffix(host)
	tld := eTLD

	sld, err := extractSld(host, eTLD)
	if err != nil {
		return nil, err
	}

	domain := sld + "." + eTLD
	manager := findManager(icann, eTLD)
	subdomain := extractSubdomain(host, domain)

	return &domainModel{
		Domain:    domain,
		Host:      host,
		Manager:   manager,
		SLD:       sld,
		Subdomain: subdomain,
		TLD:       tld,
	}, nil
}

func findManager(icann bool, eTLD string) string {
	manager := "None"
	if icann {
		manager = "ICANN"
	} else if strings.IndexByte(eTLD, '.') >= 0 {
		manager = "Private"
	}

	return manager
}

func (d *domainDataSourceModel) update(_ context.Context) error {
	domain, err := ParseDomain(d.Host.ValueString())
	if err != nil {
		return fmt.Errorf("failed to parse domain: %w", err)
	}

	d.Domain = types.StringValue(domain.Domain)
	d.Manager = types.StringValue(domain.Manager)
	d.SLD = types.StringValue(domain.SLD)
	d.Subdomain = types.StringValue(domain.Subdomain)
	d.TLD = types.StringValue(domain.TLD)

	return nil
}

func extractSubdomain(host, domain string) string {
	// If the host is the same as the domain, there is no subdomain.
	if host == domain {
		return ""
	}

	return strings.TrimSuffix(host, "."+domain)
}

func extractSld(host, eTLD string) (string, error) {
	if strings.HasPrefix(host, ".") || strings.HasSuffix(host, ".") || strings.Contains(host, "..") {
		return "", fmt.Errorf("publicsuffix: empty label in domain %q", host)
	}

	if len(host) <= len(eTLD) {
		return "", fmt.Errorf("publicsuffix: cannot derive eTLD+1 for domain %q", host)
	}
	i := len(host) - len(eTLD) - 1
	if host[i] != '.' {
		return "", fmt.Errorf("publicsuffix: invalid public suffix %q for domain %q", eTLD, host)
	}

	leftTld := host[:i]
	lastDotInLeftTld := strings.LastIndex(leftTld, ".")

	return host[1+lastDotInLeftTld : i], nil
}
