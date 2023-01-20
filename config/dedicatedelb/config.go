// Package dedicatedelb contains FlexibleEngine Dedicated ELB resources.
package dedicatedelb

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/references"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_lb_listener_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_listener_v3
	p.AddResourceConfigurator("flexibleengine_lb_listener_v3", func(r *config.Resource) {
		r.References["loadbalancer_id"] = config.Reference{
			Type: "LoadBalancer",
		}
		r.References["default_pool_id"] = config.Reference{
			Type: "Pool",
		}
		r.References["ip_group"] = config.Reference{
			Type: "IPGroup",
		}
		r.References["server_certificate"] = config.Reference{
			Type: "Certificate",
		}
		r.References["sni_certificate"] = config.Reference{
			Type: "Certificate",
		}
		r.References["ca_certificate"] = config.Reference{
			Type: "Certificate",
		}
	})

	// flexibleengine_lb_loadbalancer_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_loadbalancer_v3
	p.AddResourceConfigurator("flexibleengine_lb_loadbalancer_v3", func(r *config.Resource) {
		r.References["ipv4_subnet_id"] = references.TypeVPCSubnetIDIPV4().Get()

		r.References["ipv6_network_id"] = references.TypeVPCSubnetID().WithoutRefsSelectors().Get()
	})

	// flexibleengine_lb_member_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_member_v3
	p.AddResourceConfigurator("flexibleengine_lb_member_v3", func(r *config.Resource) {
		r.References["subnet_id"] = references.TypeVPCSubnetIDIPV4().Get()
	})

	// flexibleengine_lb_pool_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_pool_v3
	p.AddResourceConfigurator("flexibleengine_lb_pool_v3", func(r *config.Resource) {
		r.References["loadbalancer_id"] = config.Reference{
			Type: "LoadBalancer",
		}
		r.References["listener_id"] = config.Reference{
			Type: "Listener",
		}
	})

	// flexibleengine_lb_member_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_member_v3
	p.AddResourceConfigurator("flexibleengine_lb_member_v3", func(r *config.Resource) {
		// subnet_id is ipv4_subnet_id or ipv6_subnet_id
		delete(r.References, "subnet_id")
	})
}
