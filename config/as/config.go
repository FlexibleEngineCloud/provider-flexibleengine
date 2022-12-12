// Package as contains FlexibleEngine AutoScaling resources.
package as

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("flexibleengine_as_configuration_v1", func(r *config.Resource) {

		r.References["instance_config.instance_id"] = config.Reference{
			Type: tools.GenerateType("ecs", "Instance"),
		}
		r.References["image"] = config.Reference{
			Type: tools.GenerateType("ims", "Image"),
		}

	})
	p.AddResourceConfigurator("flexibleengine_as_group_v1", func(r *config.Resource) {

		r.References["scaling_configuration_id"] = config.Reference{
			Type: "Configuration",
		}
		r.References["lbaas_listeners.pool_id"] = config.Reference{
			Type: tools.GenerateType("elb", "Pool"),
		}
		// TODO Get flexibleengine_lb_listener_v2.listener_1.protocol_port
		// r.References["lbaas_listeners.protocol_port"] = config.Reference{
		// 	Type: tools.GenerateType("elb", "Listener"),
		// }
		r.References["security_groups.id"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecGroup"),
		}
		r.References["networks.id"] = config.Reference{
			Type: tools.GenerateType("vpc", "Network"),
		}
	})
	p.AddResourceConfigurator("flexibleengine_as_lifecycle_hook_v1", func(r *config.Resource) {
		r.References["scaling_configuration_id"] = config.Reference{
			Type: "Configuration",
		}
	})
	p.AddResourceConfigurator("flexibleengine_as_policy_v1", func(r *config.Resource) {
		r.References["scaling_configuration_id"] = config.Reference{
			Type: "Configuration",
		}
		r.References["alarm_id"] = config.Reference{
			Type: tools.GenerateType("ces", "AlarmRule"),
		}
	})

}
