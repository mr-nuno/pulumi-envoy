package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Config defines provider-level configuration for authenticating with Loopia.
type Config struct {
	Path string `pulumi:"path" provider:"secret"` // Path for generated configuration file
}

// Annotate adds descriptions to provider config fields for documentation and codegen.
func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Path, "Path to Envoy configuration file.")
}

// NewProvider constructs a Pulumi provider using the given client factory.
func NewProvider() (p.Provider, error) {
	return infer.NewProviderBuilder().
		WithNamespace("envoy").
		WithConfig(infer.Config(&Config{})).
		WithResources(
			infer.Resource(&EnvoyRoute{}),
			infer.Resource(&EnvoyCluster{}),
			infer.Resource(&EnvoyListener{}),
		).
		WithModuleMap(map[tokens.ModuleName]tokens.ModuleName{
			"envoy": "index",
		}).
		Build()
}
