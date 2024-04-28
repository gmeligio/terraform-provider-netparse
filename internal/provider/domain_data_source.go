package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/net/publicsuffix"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &domainDataSource{}

// domainDataSource defines the data source implementation.
type domainDataSource struct{}

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

func (d *domainDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var data domainDataSourceModel

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = data.validate(ctx)
	resp.Diagnostics.Append(diags...)
}

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

// TODO: Use regexp from `psl.isValid` to validate and remove verification of manager.
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
