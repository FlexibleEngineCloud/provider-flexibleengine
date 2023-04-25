package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/ag"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/agd"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/antiddos"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/as"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/bms"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/cce"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/ces"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/csbs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/css"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/dcs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/dds"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/dedicatedelb"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/dis"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/dli"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/dms"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/dns"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/drs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/dws"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/ecs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/eip"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/elb"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/eps"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/evs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/fgs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/iam"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/ims"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/kms"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/lts"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/modelarts"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/mrs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/nat"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/netacl"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/oss"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/rds"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/rts"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/sdrs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/sfs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/smn"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/sms"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/swr"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/tms"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/vbs"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/vpc"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/vpcep"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/config/waf"
	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

var skipList = []string{
	"flexibleengine_cts_tracker_v1$", // Only system tracker_name is available
	"flexibleengine_cce_pvc$",        // Ignored temporary due to issue https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/51
	"flexibleengine_fgs_dependency$", // Cannot upload a zip file to reference it to a dependency

	// Deprecated resources
	"flexibleengine_mrs_cluster_v1$",
	"flexibleengine_mrs_job_v1$",
	"flexibleengine_mrs_hybrid_cluster_v1$",
	"flexibleengine_elb_listener$",
	"flexibleengine_rds_instance_v1$",
	"flexibleengine_drs_replication_v2$",
	"flexibleengine_networking_router_route_v2$",
	"flexibleengine_elb_health$",
	"flexibleengine_drs_replicationconsistencygroup_v2$",
	"flexibleengine_elb_backend$",
	"flexibleengine_elb_loadbalancer$",

	// https://github.com/FlexibleEngineCloud/terraform-provider-flexibleengine/pull/869
	"flexibleengine_networking_floatingip_associate_v2$",
	"flexibleengine_networking_floatingip_v2$",
	"flexibleengine_networking_network_v2$",
	"flexibleengine_networking_subnet_v2$",
	"flexibleengine_networking_router_interface_v2$",
	"flexibleengine_networking_router_v2$",

	// Remove CSE that is in public test only
	"flexibleengine_cse_microservice$",
	"flexibleengine_cse_microservice_instance$",
	"flexibleengine_cse_microservice_engine$",
	// Duplicated resources
	"flexibleengine_fw_firewall_group_v2$", // Duplicated with flexibleengine_network_acl
	"flexibleengine_fw_policy_v2$",         // Duplicated with flexibleengine_network_acl_rule
	"flexibleengine_fw_rule_v2$",           // Duplicated with flexibleengine_network_acl_rule
}

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), tools.ResourcePrefix, tools.ModulePath, providerMetadata,
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithSkipList(skipList),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(), // 1
			defaultVersion(),             // 2
			GroupKindOverrides(),         // 3
			KnownReferencers(),           // 4
			KindOverrides(),              // 5
			KindRemoveVersion(),          // Always at the end
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		antiddos.Configure,
		dedicatedelb.Configure,
		as.Configure,
		ag.Configure,
		agd.Configure,
		dns.Configure,
		bms.Configure,
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
		dws.Configure,
		csbs.Configure,
		dms.Configure,
		sfs.Configure,
		cce.Configure,
		rts.Configure,
		sms.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
