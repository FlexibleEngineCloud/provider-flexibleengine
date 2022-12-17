// Package as contains FlexibleEngine AutoScaling resources.
package as

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_as_configuration_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/as_configuration_v1
	p.AddResourceConfigurator("flexibleengine_as_configuration_v1", func(r *config.Resource) {

		r.References["instance_config.instance_id"] = config.Reference{
			Type: tools.GenerateType("ecs", "Instance"),
		}

		r.References["instance_config.key_name"] = config.Reference{
			Type: tools.GenerateType("ecs", "KeyPair"),
		}

		r.References["image"] = config.Reference{
			Type: tools.GenerateType("ims", "Image"),
		}

	})

	// flexibleengine_as_group_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/as_group_v1
	p.AddResourceConfigurator("flexibleengine_as_group_v1", func(r *config.Resource) {

		r.References["scaling_configuration_id"] = config.Reference{
			Type: "Configuration",
		}
		r.References["lbaas_listeners.pool_id"] = config.Reference{
			Type: tools.GenerateType("elb", "Pool"),
		}
		// TODO: Get protocol_port value from elb listener
		// r.References["lbaas_listeners.protocol_port"] = config.Reference{
		// 	Type:      tools.GenerateType("elb", "Listener"),
		// 	Extractor: common.PathProtocolPortExtractor,
		// }
		r.References["security_groups.id"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecurityGroup"),
		}
		r.References["networks.id"] = config.Reference{
			Type: tools.GenerateType("vpc", "Network"),
		}
	})

	// flexibleengine_as_lifecycle_hook_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/as_lifecycle_hook_v1
	p.AddResourceConfigurator("flexibleengine_as_lifecycle_hook_v1", func(r *config.Resource) {
		r.References["scaling_configuration_id"] = config.Reference{
			Type: "Configuration",
		}
	})

	// flexibleengine_as_policy_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/as_policy_v1
	p.AddResourceConfigurator("flexibleengine_as_policy_v1", func(r *config.Resource) {
		r.References["scaling_configuration_id"] = config.Reference{
			Type: "Configuration",
		}
		r.References["alarm_id"] = config.Reference{
			Type: tools.GenerateType("ces", "AlarmRule"),
		}
	})

}
