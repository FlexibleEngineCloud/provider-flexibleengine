// Package elb contains custom ResourceConfigurators for the elb package.
package elb

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_lb_l7policy_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_l7policy_v2
	p.AddResourceConfigurator("flexibleengine_lb_l7policy_v2", func(r *config.Resource) {
		r.References["listener_id"] = config.Reference{
			Type: "Listener",
		}
		r.References["redirect_pool_id"] = config.Reference{
			Type: "Pool",
		}
		r.References["redirect_listener_id"] = config.Reference{
			Type: "Listener",
		}
	})

	// flexibleengine_lb_l7rule_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_l7rule_v2
	p.AddResourceConfigurator("flexibleengine_lb_l7rule_v2", func(r *config.Resource) {
		r.References["l7policy_id"] = config.Reference{
			Type: "L7Policy",
		}
	})

	// flexibleengine_lb_listener_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_listener_v2
	p.AddResourceConfigurator("flexibleengine_lb_listener_v2", func(r *config.Resource) {
		r.References["loadbalancer_id"] = config.Reference{
			Type: "LoadBalancer",
		}
		r.References["default_pool_id"] = config.Reference{
			Type: "Pool",
		}
	})

	// flexibleengine_lb_loadbalancer_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_loadbalancer_v2
	p.AddResourceConfigurator("flexibleengine_lb_loadbalancer_v2", func(r *config.Resource) {
		r.References["vip_subnet_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPCSubnet"),
		}
		r.References["security_group_ids"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecurityGroup"),
		}
	})

	// flexibleengine_lb_member_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_member_v2
	p.AddResourceConfigurator("flexibleengine_lb_member_v2", func(r *config.Resource) {
		r.References["pool_id"] = config.Reference{
			Type: "Pool",
		}
	})

	// flexibleengine_lb_monitor_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_monitor_v2
	p.AddResourceConfigurator("flexibleengine_lb_monitor_v2", func(r *config.Resource) {
		r.References["pool_id"] = config.Reference{
			Type: "Pool",
		}
	})

	// flexibleengine_lb_pool_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_pool_v2
	p.AddResourceConfigurator("flexibleengine_lb_pool_v2", func(r *config.Resource) {
		r.References["loadbalancer_id"] = config.Reference{
			Type: "LoadBalancer",
		}
		r.References["listener_id"] = config.Reference{
			Type: "Listener",
		}
	})

	// flexibleengine_lb_whitelist_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_whitelist_v2
	p.AddResourceConfigurator("flexibleengine_lb_whitelist_v2", func(r *config.Resource) {
		r.References["listener_id"] = config.Reference{
			Type: "Listener",
		}
	})
}
