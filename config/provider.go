package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/FrangipaneTeam/provider-flexibleengine/config/ag"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/agd"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/antiddos"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/as"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/bms"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/cce"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/ces"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/csbs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/cse"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/css"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/cts"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dcs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dds"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dedicatedelb"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dis"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dli"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dms"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/drs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/dws"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/ecs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/eip"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/elb"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/eps"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/evs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/fgs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/iam"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/ims"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/kms"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/lts"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/modelarts"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/mrs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/nat"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/netacl"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/oss"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/rds"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/sdrs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/sfs"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/smn"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/swr"
	"github.com/FrangipaneTeam/provider-flexibleengine/config/tms"
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
			ExternalNameConfigurations(), // 1
			defaultVersion(),             // 2
			GroupKindOverrides(),         // 3
			KnownReferencers(),           // 4
			KindOverrides(),              // 5
			KindRemoveVersion(),          // Always at the end
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		antiddos.Configure,
		dedicatedelb.Configure,
		as.Configure,
		ag.Configure,
		agd.Configure,
		bms.Configure,
		cse.Configure,
		lts.Configure,
		dcs.Configure,
		fgs.Configure,
		lts.Configure,
		drs.Configure,
		smn.Configure,
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
		sdrs.Configure,
		ces.Configure,
		css.Configure,
		mrs.Configure,
		dli.Configure,
		tms.Configure,
		netacl.Configure,
		dis.Configure,
		cts.Configure,
		dws.Configure,
		csbs.Configure,
		dms.Configure,
		sfs.Configure,
		cce.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
