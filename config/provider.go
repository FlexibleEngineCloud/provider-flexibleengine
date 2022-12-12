package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/FrangipaneTeam/provider-flexibleengine/config/ag"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/agd"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/bms"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/ces"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dcs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dds"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dedicatedelb"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/drs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/ecs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/eip"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/elb"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/eps"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/evs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/iam"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/ims"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/kms"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/modelarts"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/nat"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/oss"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/rds"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/swr"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/vbs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/vpc"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/vpcep"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/waf"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"

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
		eps.Configure,
		swr.Configure,
		waf.Configure,
		modelarts.Configure,
		evs.Configure,
		vbs.Configure,
		elb.Configure,
		ecs.Configure,
		rds.Configure,
		eip.Configure,
		iam.Configure,
		ims.Configure,
		kms.Configure,
		vpc.Configure,
		vpcep.Configure,
		oss.Configure,
		nat.Configure,
		ces.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
