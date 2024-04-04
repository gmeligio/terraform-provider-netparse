package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/net/publicsuffix"
)

// domainDataSourceModel describes the data source data model.
type domainDataSourceModel struct {
	Domain types.String `tfsdk:"domain"`
}

func (d domainDataSourceModel) validate(_ context.Context) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.Domain.IsUnknown() || d.Domain.IsNull() {
		return diags
	}

	domain := d.Domain.ValueString()

	eTLD, icann := publicsuffix.PublicSuffix(domain)

	manager := "Unmanaged"
	if icann {
		manager = "ICANN Managed"
	} else if strings.IndexByte(eTLD, '.') >= 0 {
		manager = "Privately Managed"
	}

	if manager == "Unmanaged" {
		diags.AddAttributeError(
			path.Root("domain"),
			"Invalid Attribute Configuration",
			"Expected domain to be either ICANN managede or privately managed.",
		)
	}

	return diags
}

// func (d *domainDataSourceModel) update(ctx context.Context) diag.Diagnostics {
// 	var buffer bytes.Buffer
// 	var diags diag.Diagnostics
// 	var err error

// 	// cloudinit Provider 'v2.2.0' doesn't actually set default values in state properly, so we need to make sure
// 	// that we don't use any known empty values from previous versions of state
// 	diags.Append(d.setDefaults(ctx)...)
// 	if diags.HasError() {
// 		return diags
// 	}

// 	var configParts []configPartModel
// 	diags.Append(d.Parts.ElementsAs(ctx, &configParts, false)...)
// 	if diags.HasError() {
// 		return diags
// 	}

// 	if d.Gzip.ValueBool() {
// 		gzipWriter := gzip.NewWriter(&buffer)

// 		err = renderPartsToWriter(ctx, d.Boundary.ValueString(), configParts, gzipWriter)

// 		gzipWriter.Close()
// 	} else {
// 		err = renderPartsToWriter(ctx, d.Boundary.ValueString(), configParts, &buffer)
// 	}

// 	if err != nil {
// 		diags.AddError("Unable to render cloudinit config to MIME multi-part file", err.Error())
// 		return diags
// 	}

// 	output := ""
// 	if d.Base64Encode.ValueBool() {
// 		output = base64.StdEncoding.EncodeToString(buffer.Bytes())
// 	} else {
// 		output = buffer.String()
// 	}

// 	d.ID = types.StringValue(strconv.Itoa(hashcode.String(output)))
// 	d.Rendered = types.StringValue(output)

// 	return diags
// }
