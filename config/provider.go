package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/gaetanars/provider-flexibleengine/config/ag"
	"github.com/gaetanars/provider-flexibleengine/config/agd"
	"github.com/gaetanars/provider-flexibleengine/config/bms"
	"github.com/gaetanars/provider-flexibleengine/config/dcs"
	"github.com/gaetanars/provider-flexibleengine/config/dds"
	"github.com/gaetanars/provider-flexibleengine/config/dedicatedelb"
	"github.com/gaetanars/provider-flexibleengine/config/drs"
	"github.com/gaetanars/provider-flexibleengine/config/ecs"
	"github.com/gaetanars/provider-flexibleengine/config/eip"
	"github.com/gaetanars/provider-flexibleengine/config/elb"
	"github.com/gaetanars/provider-flexibleengine/config/evs"
	"github.com/gaetanars/provider-flexibleengine/config/iam"
	"github.com/gaetanars/provider-flexibleengine/config/ims"
	"github.com/gaetanars/provider-flexibleengine/config/kms"
	"github.com/gaetanars/provider-flexibleengine/config/nat"
	"github.com/gaetanars/provider-flexibleengine/config/oss"
	"github.com/gaetanars/provider-flexibleengine/config/swr"
	"github.com/gaetanars/provider-flexibleengine/config/vbs"
	"github.com/gaetanars/provider-flexibleengine/config/vpc"
	"github.com/gaetanars/provider-flexibleengine/config/vpcep"
	"github.com/gaetanars/provider-flexibleengine/config/waf"
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

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
			// KindRemoveVersion Always at the end
			KindRemoveVersion(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		dedicatedelb.Configure,
		ag.Configure,
		agd.Configure,
		bms.Configure,
		dcs.Configure,
		drs.Configure,
		dds.Configure,
		swr.Configure,
		waf.Configure,
		evs.Configure,
		vbs.Configure,
		elb.Configure,
		ecs.Configure,
		eip.Configure,
		iam.Configure,
		ims.Configure,
		kms.Configure,
		vpc.Configure,
		vpcep.Configure,
		oss.Configure,
		nat.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
