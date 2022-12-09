package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/gaetanars/provider-flexibleengine/config/ecs"
	"github.com/gaetanars/provider-flexibleengine/config/eip"
	"github.com/gaetanars/provider-flexibleengine/config/iam"
	"github.com/gaetanars/provider-flexibleengine/config/vpc"
	"github.com/gaetanars/provider-flexibleengine/config/vpcep"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "flexibleengine"
	modulePath     = "github.com/gaetanars/provider-flexibleengine"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			KnownReferencers(),
			GroupKindOverrides(),
			KindOverrides(),
			// KindRemoveVersion END
			KindRemoveVersion(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		ecs.Configure,
		eip.Configure,
		iam.Configure,
		vpc.Configure,
		vpcep.Configure,
		eip.Configure,
		iam.Configure,
		ecs.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
