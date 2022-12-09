package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/gaetanars/provider-flexibleengine/config/dedicatedelb"
	"github.com/gaetanars/provider-flexibleengine/config/ecs"
	"github.com/gaetanars/provider-flexibleengine/config/eip"
	"github.com/gaetanars/provider-flexibleengine/config/iam"
	"github.com/gaetanars/provider-flexibleengine/config/vpc"
	"github.com/gaetanars/provider-flexibleengine/config/vpcep"
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const ()

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), tools.ResourcePrefix, tools.ModulePath, providerMetadata,
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			defaultVersion(),
			GroupKindOverrides(),
			KnownReferencers(),
			KindOverrides(),
			// KindRemoveVersion END
			KindRemoveVersion(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		dedicatedelb.Configure,
		ecs.Configure,
		eip.Configure,
		iam.Configure,
		vpc.Configure,
		vpcep.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
