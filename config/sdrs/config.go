package sdrs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("flexibleengine_sdrs_drill_v1", func(r *config.Resource) {

		// group_id
		r.References["group_id"] = config.Reference{
			Type: "ProtectionGroup",
		}
		r.References["drill_vpc_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPC"),
		}

	})
	p.AddResourceConfigurator("flexibleengine_sdrs_protectedinstance_v1", func(r *config.Resource) {

		// group_id
		r.References["group_id"] = config.Reference{
			Type: "ProtectionGroup",
		}
		r.References["server_id"] = config.Reference{
			Type: tools.GenerateType("ecs", "Instance"),
		}
		r.References["primary_subnet_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPCSubnet"),
		}

	})
	p.AddResourceConfigurator("flexibleengine_sdrs_protectiongroup_v1", func(r *config.Resource) {
		r.References["source_vpc_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPC"),
		}
	})
	p.AddResourceConfigurator("flexibleengine_sdrs_replication_attach_v1", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type: "ProtectedInstance",
		}
		r.References["replication_id"] = config.Reference{
			Type: "ReplicationPair",
		}
	})
	p.AddResourceConfigurator("flexibleengine_sdrs_replication_pair_v1", func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			Type: "ProtectionGroup",
		}
		r.References["volume_id"] = config.Reference{
			Type: tools.GenerateType("evs", "BlockStorageVolume"),
		}
	})

}
