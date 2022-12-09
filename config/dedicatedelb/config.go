package dedicatedelb

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
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
	p.AddResourceConfigurator("flexibleengine_lb_loadbalancer_v3", func(r *config.Resource) {
		r.References["loadbalancer_id"] = config.Reference{
			Type: "LoadBalancer",
		}
		r.References["ipv4_subnet_id"] = config.Reference{
			Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet",
		}
		r.References["ipv6_network_id"] = config.Reference{
			Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Network",
		}
		r.References["ipv4_eip_id"] = config.Reference{
			Type: "github.com/gaetanars/provider-flexibleengine/apis/eip/v1beta1.EIP",
		}
	})
	p.AddResourceConfigurator("flexibleengine_lb_member_v3", func(r *config.Resource) {
		r.References["pool_id"] = config.Reference{
			Type: "Pool",
		}
	})
	p.AddResourceConfigurator("flexibleengine_lb_monitor_v3", func(r *config.Resource) {
		r.References["pool_id"] = config.Reference{
			Type: "Pool",
		}
	})
	p.AddResourceConfigurator("flexibleengine_lb_pool_v3", func(r *config.Resource) {
		r.References["loadbalancer_id"] = config.Reference{
			Type: "LoadBalancer",
		}
		r.References["listener_id"] = config.Reference{
			Type: "Listener",
		}
	})
}
