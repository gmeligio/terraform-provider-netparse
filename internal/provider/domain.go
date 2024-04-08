package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/net/publicsuffix"
)

// domainDataSourceModel describes the data source data model.
// References used
// https://github.com/lupomontero/psl
// https://github.com/jpillora/go-tld
// https://github.com/zomasec/tld
// https://github.com/bobesa/go-domain-util
// https://github.com/joeguo/tldextract
type domainDataSourceModel struct {
	Domain    types.String `tfsdk:"domain"`
	Host      types.String `tfsdk:"host"`
	Manager   types.String `tfsdk:"manager"`
	SLD       types.String `tfsdk:"sld"`
	Subdomain types.String `tfsdk:"subdomain"`
	TLD       types.String `tfsdk:"tld"`
}

// TODO: Use regexp from `psl.isValid` to validate and remove verification of manager
func (d domainDataSourceModel) validate(_ context.Context) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.Host.IsUnknown() || d.Host.IsNull() {
		return diags
	}

	host := d.Host.ValueString()

	eTLD, icann := publicsuffix.PublicSuffix(host)

	manager := findManager(icann, eTLD)

	if manager == "None" {
		diags.AddAttributeError(
			path.Root("host"),
			"Invalid Attribute Configuration",
			"Expected host to have as a manager either ICANN or Private.",
		)
	}

	return diags
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

func (d *domainDataSourceModel) update(_ context.Context) diag.Diagnostics {
	var diags diag.Diagnostics

	host := d.Host.ValueString()

	eTLD, icann := publicsuffix.PublicSuffix(host)
	d.TLD = types.StringValue(eTLD)

	sld, err := extractSld(host, eTLD)
	if err != nil {
		diags.AddAttributeError(
			path.Root("sld"),
			"Invalid Attribute Configuration",
			err.Error(),
		)
	}
	d.SLD = types.StringValue(sld)

	domain := sld + "." + eTLD
	d.Domain = types.StringValue(domain)

	manager := findManager(icann, eTLD)
	d.Manager = types.StringValue(manager)

	subdomain := extractSubdomain(host, domain)
	d.Subdomain = types.StringValue(subdomain)

	return diags
}

func extractSubdomain(host, domain string) string {
	// If the host is the same as the domain, there is no subdomain
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
