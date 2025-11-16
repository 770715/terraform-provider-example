package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies expected interfaces
var (
	_ provider.Provider = &exampleProvider{}
)

type exampleProvider struct {
	version string
}

// New returns a provider factory function
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &exampleProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name
func (p *exampleProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "example"
	resp.Version = p.version
}

// Schema defines the provider-level configuration schema
func (p *exampleProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Example provider demonstrating Plugin Framework v6.",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				Description: "API endpoint for the example service.",
				Optional:    true,
			},
		},
	}
}

// Configure prepares provider data for resources and data sources
func (p *exampleProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider configuration
	var config struct {
		Endpoint string `tfsdk:"endpoint"`
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize API client here if needed
	// Store client in resp.ResourceData and resp.DataSourceData
}

// DataSources defines data sources implemented by the provider
func (p *exampleProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Resources defines resources implemented by the provider
func (p *exampleProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}
