package provider

import (
	"context"
	"fmt"

	"github.com/gmeligio/terraform-provider-netparse/internal/netparse"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	domainMarkdownDescription        = "Parses Public Suffix List properties from a domain. It uses the [publicsuffix](https://pkg.go.dev/golang.org/x/net/publicsuffix) go package to parse the domain. For more details on the domain parts, see [What is a Domain Name?](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_domain_name)."
	domainAttrMarkdownDescription    = "The domain name. It's the tld plus one more label."
	hostAttrMarkdownDescription      = "The host that identifies the domain name."
	managerAttrMarkdownDescription   = "The manager is the entity that manages the domain. It can be one of: ICANN, Private, or None."
	sldAttrMarkdownDescription       = "The second-level domain (SLD) is the label to the left of the effective TLD."
	subdomainAttrMarkdownDescription = "The subdomain is the left part of the host that is not the domain."
	tldAttrMarkdownDescription       = "The effective top-level domain (eTLD) of the domain. This is the public suffix of the domain."
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &domainDataSource{}

// domainDataSource defines the data source implementation.
type domainDataSource struct{}

// domainDataSourceModel describes the data source model.
type domainDataSourceModel struct {
	Domain    types.String `tfsdk:"domain"`
	Host      types.String `tfsdk:"host"`
	Manager   types.String `tfsdk:"manager"`
	SLD       types.String `tfsdk:"sld"`
	Subdomain types.String `tfsdk:"subdomain"`
	TLD       types.String `tfsdk:"tld"`
}

func NewDomainDataSource() datasource.DataSource {
	return &domainDataSource{}
}

func (d *domainDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

// func (d *domainDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
// 	var data domainDataSourceModel

// 	diags := req.Config.Get(ctx, &data)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	diags = data.validate(ctx)
// 	resp.Diagnostics.Append(diags...)
// }

func (d *domainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: domainMarkdownDescription,

		Attributes: map[string]schema.Attribute{
			"domain": schema.StringAttribute{
				MarkdownDescription: domainAttrMarkdownDescription,
				Computed:            true,
			},
			"host": schema.StringAttribute{
				MarkdownDescription: hostAttrMarkdownDescription,
				Required:            true,
			},
			"manager": schema.StringAttribute{
				MarkdownDescription: managerAttrMarkdownDescription,
				Computed:            true,
			},
			"sld": schema.StringAttribute{
				MarkdownDescription: sldAttrMarkdownDescription,
				Computed:            true,
			},
			"subdomain": schema.StringAttribute{
				MarkdownDescription: subdomainAttrMarkdownDescription,
				Computed:            true,
			},
			"tld": schema.StringAttribute{
				MarkdownDescription: tldAttrMarkdownDescription,
				Computed:            true,
			},
		},
	}
}

func (d *domainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data domainDataSourceModel

	diags := req.Config.Get(ctx, &data)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := data.update(ctx)
	if err != nil {
		resp.Diagnostics.AddError("failed to update data", err.Error())
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// func (d domainDataSourceModel) validate(_ context.Context) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	if d.Host.IsUnknown(){
// 		return diags
// 	}

// 	if d.Host.IsNull() {
// 		diags.AddAttributeError(
// 			path.Root("host"),
// 			"Invalid Attribute Configuration",
// 			"Expected host to be non-null. Received a null value.",
// 		)
// 	}

// 	if netparse.DomainValidate(d.Host.ValueString()) != nil {
// 		diags.AddAttributeError(
// 			path.Root("host"),
// 			"Invalid Attribute Configuration",
// 			"Expected host to be valid. Received an invalid value.",
// 		)
// 	}

// 	return diags
// }

func (d *domainDataSourceModel) update(_ context.Context) error {
	domain, err := netparse.ParseDomain(d.Host.ValueString())
	if err != nil {
		return fmt.Errorf("failed to parse domain: %w", err)
	}

	if domain.Manager == "None" {
		return fmt.Errorf("unsupported manager: None")
	}

	d.Domain = types.StringValue(domain.Domain)
	d.Manager = types.StringValue(domain.Manager)
	d.SLD = types.StringValue(domain.SLD)
	d.Subdomain = types.StringValue(domain.Subdomain)
	d.TLD = types.StringValue(domain.TLD)

	return nil
}
