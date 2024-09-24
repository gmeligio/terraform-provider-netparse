package provider

import (
	"context"
	"fmt"

	"github.com/gmeligio/terraform-provider-netparse/internal/netparse"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	urlMarkdownDescription                  = "Parses URL components from a URL string. It uses the [net/url](https://pkg.go.dev/net/url) go package to parse the URL. For more details on the URL components, see [What is a URL?](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_is_a_URL) and [WHATWG URL Standard](https://url.spec.whatwg.org/#api)."
	urlAttributeMarkdownDescription         = "The URL to parse."
	authorityAttributeMarkdownDescription   = "The concatenation of the username, password, host, and port. It's separated from the scheme by `://`."
	credentialsAttributeMarkdownDescription = "The concatenation of the username and password."
	fragmentAttributeMarkdownDescription    = "The component after the search."
	hashAttributeMarkdownDescription        = "The concatenation of a `#` with the fragment."
	hostAttributeMarkdownDescription        = "The domain part of the authority."
	passwordAttributeMarkdownDescription    = "The second component of the credentials."
	protocolAttributeMarkdownDescription    = "The concatenation of the protocol scheme and `:`."
	schemeAttributeMarkdownDescription      = "The protocol scheme used to access the domain."
	usernameAttributeMarkdownDescription    = "The first component of the credentials."
	portAttributeMarkdownDescription        = "The last component of the authority."
	pathAttributeMarkdownDescription        = "The component after the authority."
	queryAttributeMarkdownDescription       = "A substring of the search component, after the `?` and before the fragment."
	searchAttributeMarkdownDescription      = "The component after the path."
)


// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &urlDataSource{}

func NewUrlDataSource() datasource.DataSource {
	return &urlDataSource{}
}

// urlDataSource defines the data source implementation.
type urlDataSource struct{}

type urlDataSourceModel struct {
	Url         types.String `tfsdk:"url"`
	Authority   types.String `tfsdk:"authority"`
	Protocol    types.String `tfsdk:"protocol"`
	Scheme      types.String `tfsdk:"scheme"`
	Credentials types.String `tfsdk:"credentials"`
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	Host        types.String `tfsdk:"host"`
	Port        types.String `tfsdk:"port"`
	Path        types.String `tfsdk:"path"`
	Search      types.String `tfsdk:"search"`
	Query       types.String `tfsdk:"query"`
	Hash        types.String `tfsdk:"hash"`
	Fragment    types.String `tfsdk:"fragment"`
}

func NewUrlDataSourceModel() *urlDataSourceModel {
	return &urlDataSourceModel{}
}

func (u *urlDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_url"
}

// func (u *urlDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
// 	var data urlDataSourceModel

// 	diags := req.Config.Get(ctx, &data)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	diags = data.validate(ctx)
// 	resp.Diagnostics.Append(diags...)
// }

func (u *urlDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: urlMarkdownDescription,

		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				MarkdownDescription: urlAttributeMarkdownDescription,
				Required:            true,
			},
			"authority": schema.StringAttribute{
				MarkdownDescription: authorityAttributeMarkdownDescription,
				Computed:            true,
			},
			"scheme": schema.StringAttribute{
				MarkdownDescription: schemeAttributeMarkdownDescription,
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: protocolAttributeMarkdownDescription,
				Computed:            true,
			},
			"credentials": schema.StringAttribute{
				MarkdownDescription: credentialsAttributeMarkdownDescription,
				Computed:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: usernameAttributeMarkdownDescription,
				Computed:            true,
				Sensitive:           true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: passwordAttributeMarkdownDescription,
				Computed:            true,
				Sensitive:           true,
			},
			"host": schema.StringAttribute{
				MarkdownDescription: hostAttributeMarkdownDescription,
				Computed:            true,
			},
			"port": schema.StringAttribute{
				MarkdownDescription: portAttributeMarkdownDescription,
				Computed:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: pathAttributeMarkdownDescription,
				Computed:            true,
			},
			"search": schema.StringAttribute{
				MarkdownDescription: searchAttributeMarkdownDescription,
				Computed:            true,
			},
			"query": schema.StringAttribute{
				MarkdownDescription: queryAttributeMarkdownDescription,
				Computed:            true,
			},
			"fragment": schema.StringAttribute{
				MarkdownDescription: fragmentAttributeMarkdownDescription,
				Computed:            true,
			},
			"hash": schema.StringAttribute{
				MarkdownDescription: hashAttributeMarkdownDescription,
				Computed:            true,
			},
		},
	}
}

func (u *urlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data urlDataSourceModel

	diags := req.Config.Get(ctx, &data)

	tflog.Trace(ctx, "Parsed URL data source model", map[string]interface{}{
		"data": data,
	})

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := data.update(ctx)
	if err != nil {
		resp.Diagnostics.AddError("failed to update data", err.Error())
	}

	tflog.Trace(ctx, "Updated URL data source model", map[string]interface{}{
		"data": data,
	})

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// func (d *urlDataSourceModel) validate(_ context.Context) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	if d.Url.IsUnknown() {
// 		return diags
// 	}

// 	if d.Url.IsNull() {
// 		diags.AddAttributeError(
// 			path.Root("url"),
// 			"Invalid Attribute Configuration",
// 			"Expected url to be non-null. Received a null value.",
// 		)
// 	}

// 	if netparse.UrlValidate(d.Url.ValueString()) != nil {
// 		diags.AddAttributeError(
// 			path.Root("url"),
// 			"Invalid Attribute Configuration",
// 			"Expected url to be valid. Received an invalid value.",
// 		)
// 	}

// 	return diags
// }

func (u *urlDataSourceModel) update(_ context.Context) error {
	url, err := netparse.ParseUrl(u.Url.ValueString())
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	u.Authority = types.StringValue(url.Authority)
	u.Protocol = types.StringValue(url.Protocol)
	u.Scheme = types.StringValue(url.Scheme)
	u.Credentials = types.StringValue(url.Credentials)
	u.Username = types.StringValue(url.Username)
	u.Password = types.StringValue(url.Password)
	u.Host = types.StringValue(url.Host)
	u.Port = types.StringValue(url.Port)
	u.Path = types.StringValue(url.Path)
	u.Search = types.StringValue(url.Search)
	u.Query = types.StringValue(url.Query)
	u.Hash = types.StringValue(url.Hash)
	u.Fragment = types.StringValue(url.Fragment)

	return nil
}
