package netparse

import (
	"fmt"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// DomainModel describes the domain model.
// References used.
// https://github.com/lupomontero/psl
// https://github.com/jpillora/go-tld
// https://github.com/zomasec/tld
// https://github.com/bobesa/go-domain-util
// https://github.com/joeguo/tldextract
type DomainModel struct {
	Domain    string
	Host      string
	Manager   string
	SLD       string
	Subdomain string
	TLD       string
}

func ParseDomain(h string) (*DomainModel, error) {
	host := h
	eTLD, icann := publicsuffix.PublicSuffix(host)
	tld := eTLD

	sld, err := extractSld(host, eTLD)
	if err != nil {
		return nil, err
	}

	domain := sld + "." + eTLD
	manager := FindManager(icann, eTLD)
	subdomain := extractSubdomain(host, domain)

	return &DomainModel{
		Domain:    domain,
		Host:      host,
		Manager:   manager,
		SLD:       sld,
		Subdomain: subdomain,
		TLD:       tld,
	}, nil
}

func FindManager(icann bool, eTLD string) string {
	manager := "None"
	if icann {
		manager = "ICANN"
	} else if strings.IndexByte(eTLD, '.') >= 0 {
		manager = "Private"
	}

	return manager
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

func DomainValidate(u string) error {
	return nil
}
