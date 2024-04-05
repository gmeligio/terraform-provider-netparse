package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure PublicsuffixProvider satisfies various provider interfaces.
var _ provider.Provider = &PublicsuffixProvider{}
var _ provider.ProviderWithFunctions = &PublicsuffixProvider{}

// PublicsuffixProvider defines the provider implementation.
type PublicsuffixProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// PublicsuffixProviderModel describes the provider data model.
type PublicsuffixProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *PublicsuffixProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "publicsuffix"
	resp.Version = p.version
}

func (p *PublicsuffixProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *PublicsuffixProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *PublicsuffixProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *PublicsuffixProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDomainDataSource,
	}
}

func (p *PublicsuffixProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PublicsuffixProvider{
			version: version,
		}
	}
}
