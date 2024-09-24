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
	cidrMarkdownDescription        = "Parses an IP address and prefix length in CIDR notation as defined in [RFC 4632](https://rfc-editor.org/rfc/rfc4632.html) and [RFC 4291](https://rfc-editor.org/rfc/rfc4291.html)."
	cidrAttrMarkdownDescription    = "IP address and prefix length in CIDR notation."
	ipAttrMarkdownDescription    = "IP address."
	networkAttrMarkdownDescription    = "IP network."
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &cidrDataSource{}

// cidrDataSource defines the data source implementation.
type cidrDataSource struct{}

// cidrDataSource defines the data source model.
type cidrDataSourceModel struct {
	CIDR    types.String `tfsdk:"cidr"`
	IP      types.String `tfsdk:"ip"`
	Network   types.String `tfsdk:"network"`
}

func NewCIDRDataSource() datasource.DataSource {
	return &cidrDataSource{}
}

func (d *cidrDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cidr"
}

func (d *cidrDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: cidrMarkdownDescription,

		Attributes: map[string]schema.Attribute{
			"cidr": schema.StringAttribute{
				MarkdownDescription: cidrAttrMarkdownDescription,
				Required:            true,
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: ipAttrMarkdownDescription,
				Computed:            true,
			},
			"network": schema.StringAttribute{
				MarkdownDescription: networkAttrMarkdownDescription,
				Computed:            true,
			},
		},
	}
}

func (d *cidrDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data cidrDataSourceModel

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

func (d *cidrDataSourceModel) update(_ context.Context) error {
	cidr, err := netparse.ParseCIDR(d.CIDR.ValueString())
	if err != nil {
		return fmt.Errorf("failed to parse cidr: %w", err)
	}

	d.IP = types.StringValue(cidr.IP)
	d.Network = types.StringValue(cidr.Network)

	return nil
}
