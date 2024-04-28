package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure NetparseProvider satisfies various provider interfaces.
var _ provider.Provider = &NetparseProvider{}
var _ provider.ProviderWithFunctions = &NetparseProvider{}

// NetparseProvider defines the provider implementation.
type NetparseProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *NetparseProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netparse"
	resp.Version = p.version
}

func (p *NetparseProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *NetparseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *NetparseProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *NetparseProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUrlDataSource,
		NewDomainDataSource,
	}
}

func (p *NetparseProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewParseUrlFunction,
		NewParseDomainFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NetparseProvider{
			version: version,
		}
	}
}
