// Package vpc contains the configuration for the FlexibleEngine VPC service.
package vpc

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_vpc_route_table
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_route_table
	p.AddResourceConfigurator("flexibleengine_vpc_route_table", func(r *config.Resource) {

		// Subnets is a list of Subnet IDs
		r.References["subnets"] = config.Reference{
			Type: "VPCSubnet",
		}
	})

	// flexibleengine_vpc_route
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_route
	p.AddResourceConfigurator("flexibleengine_vpc_route", func(r *config.Resource) {

		// route_table_id is the ID of the Route Table
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}

		// TODO nexthop
		/*
			nexthop (Required, String) - Specifies the next hop.
				If the route type is ecs, the value is an ECS instance ID in the VPC.
				If the route type is eni, the value is the extension NIC of an ECS in the VPC.
				If the route type is vip, the value is a virtual IP address.
				If the route type is nat, the value is a VPN gateway ID.
				If the route type is peering, the value is a VPC peering connection ID.
				If the route type is vpn, the value is a VPN gateway ID.
				If the route type is dc, the value is a Direct Connect gateway ID.
		*/

	})

	// flexibleengine_vpc_peering_connection_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_peering_v2
	p.AddResourceConfigurator("flexibleengine_vpc_peering_connection_v2", func(r *config.Resource) {

		// peer_vpc_id is the ID of the peer VPC
		r.References["peer_vpc_id"] = config.Reference{
			Type: "VPC",
		}

	})

	// flexibleengine_vpc_peering_connection_accepter_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_peering_accepter_v2
	p.AddResourceConfigurator("flexibleengine_vpc_peering_connection_accepter_v2", func(r *config.Resource) {

		// peering_id is the ID of the VPC Peering Connection
		r.References["vpc_peering_connection_id"] = config.Reference{
			Type: "PeeringConnection",
		}

	})

	// TODO This require LTS (Log Tank Service)
	// flexibleengine_vpc_flow_log_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_flow_log_v1
	// p.AddResourceConfigurator("flexibleengine_vpc_flow_log_v1", func(r *config.Resource) {

	// 	// TODO Check if this is correct
	// 	r.References["ressource_id"] = config.Reference{
	// 		Type: "NetworkPort",
	// 	}

	// 	r.References["log_group_id"] = config.Reference{
	// 		Type: "LogGroup",
	// 	}

	// })

	// flexibleengine_networking_secgroup_rule_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_secgroup_rule_v2
	p.AddResourceConfigurator("flexibleengine_networking_secgroup_rule_v2", func(r *config.Resource) {

		// remote_group_id is the ID of the security group to which this rule applies.
		r.References["remote_group_id"] = config.Reference{
			Type: "SecGroup",
		}
	})

}
