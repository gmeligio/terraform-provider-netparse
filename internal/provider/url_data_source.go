package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &urlDataSource{}

func NewUrlDataSource() datasource.DataSource {
	return &urlDataSource{}
}

// urlDataSource defines the data source implementation.
type urlDataSource struct {
}

func (u *urlDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_url"
}

func (u *urlDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var data urlDataSourceModel

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = data.validate(ctx)
	resp.Diagnostics.Append(diags...)
}

func (u *urlDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Parses URL components from a URL string.",

		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				MarkdownDescription: "The URL to parse",
				Required:            true,
			},
			"authority": schema.StringAttribute{
				MarkdownDescription: "The concatenation of the username, password, host, and port. It's separated from the scheme by :// . For example: user1:123@example.com:3000 for http://user1:123@example.com:3000",
				Computed:            true,
			},
			"scheme": schema.StringAttribute{
				MarkdownDescription: "The protocol used to access the domain. For example: http, https, ftp, sftp, file, etc.",
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "The concatenation of the scheme and the port. For example: http:, https:, ftp:, sftp:, file:, etc.",
				Computed:            true,
			},
			"credentials": schema.StringAttribute{
				MarkdownDescription: "The concatenation of the username and password. For example: user1:123 for https://user1:123@example.com",
				Computed:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "The first component of the URL credentials. For example: user1 for https://user1:123@example.com",
				Computed:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The second component of the URL credentials. For example: 123 for https://user1:123@example.com",
				Computed:            true,
			},
			"host": schema.StringAttribute{
				MarkdownDescription: "The domain part of the authority. For example: example.com for https://example.com",
				Computed:            true,
			},
			"port": schema.StringAttribute{
				MarkdownDescription: "The last component of the URL authority. For example: 443 for https://example.com:443",
				Computed:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "The URL component after the authority. For example: /path/to/resource for https://example.com/path/to/resource",
				Computed:            true,
			},
			"search": schema.StringAttribute{
				MarkdownDescription: "The URL component after the path. For example: ?key=value for https://example.com/path/to/resource?key=value",
				Computed:            true,
			},
			"query": schema.StringAttribute{
				MarkdownDescription: "The URL component of the search starting at the ? and before the fragment. For example: key=value for https://example.com/path/to/resource?key=value#section",
				Computed:            true,
			},
			"fragment": schema.StringAttribute{
				MarkdownDescription: "The URL component after the search. For example: section for https://example.com/path/to/resource?key=value#section",
				Computed:            true,
			},
			"hash": schema.StringAttribute{
				MarkdownDescription: "The concatenation of a # with the fragment. For example: #section for https://example.com/path/to/resource?key=value#section",
				Computed:            true,
			},
		},
	}
}

func (u *urlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data urlDataSourceModel

	diags := req.Config.Get(ctx, &data)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = data.update(ctx)
	resp.Diagnostics.Append(diags...)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
