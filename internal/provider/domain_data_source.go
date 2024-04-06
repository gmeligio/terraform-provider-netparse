package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &domainDataSource{}

func NewDomainDataSource() datasource.DataSource {
	return &domainDataSource{}
}

// domainDataSource defines the data source implementation.
type domainDataSource struct {
}

func (d *domainDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

func (d *domainDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var data domainDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(data.validate(ctx)...)
}

func (d *domainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Parses Public Suffix List properties from a domain",

		Attributes: map[string]schema.Attribute{
			"domain": schema.StringAttribute{
				MarkdownDescription: "The domain name. It's the tld plus one more label. For example: example.com for host foo.example.com",
				Computed:            true,
			},
			"host": schema.StringAttribute{
				MarkdownDescription: "The host that identifies the domain name",
				Required:            true,
			},
			"manager": schema.StringAttribute{
				MarkdownDescription: "The manager is the entity that manages the domain. It can be one of the following: ICANN, Private, or None.",
				Computed:            true,
			},
			"sld": schema.StringAttribute{
				MarkdownDescription: "The second-level domain (SLD) is the label to the left of the effective TLD. For example: example for example.com, or foo for foo.co.uk",
				Computed:            true,
			},
			"subdomain": schema.StringAttribute{
				MarkdownDescription: "The subdomain is the left part of the host that is not the domain. For example: www for www.example.com, mail for mail.foo.org, blog for blog.bar.org",
				Computed:            true,
			},
			"tld": schema.StringAttribute{
				MarkdownDescription: "The effective top-level domain (eTLD) of the domain. This is the public suffix of the domain. For example: com for example.com, or co.uk for foo.co.uk",
				Computed:            true,
			},
		},
	}
}

func (d *domainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data domainDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "read a data source")

	resp.Diagnostics.Append(data.update(ctx)...)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
