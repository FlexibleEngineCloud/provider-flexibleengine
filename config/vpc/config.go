// Package vpc contains the configuration for the FlexibleEngine VPC service.
package vpc

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_vpc_route_table
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_route_table
	p.AddResourceConfigurator("flexibleengine_vpc_route_table", func(r *config.Resource) {

		// Subnets is a list of Subnet IDs
		r.References["subnets"] = config.Reference{
			Type:              "VPCSubnet",
			Extractor:         common.PathIDExtractor,
			SelectorFieldName: "SubnetSelector", // Selector is only one reference
		}

	})

	// flexibleengine_vpc_route
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_route
	p.AddResourceConfigurator("flexibleengine_vpc_route", func(r *config.Resource) {

		// route_table_id is the ID of the Route Table
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}

		// note(azrod): The nexthop field require multiple types of values
		// https://github.com/upbound/upjet/issues/95
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

		// peer_tenant_id is the ID of the tenant to which the peer VPC belongs.
		r.TerraformResource.Schema["peer_tenant_id"].Sensitive = true

	})

	// flexibleengine_vpc_peering_connection_accepter_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_peering_accepter_v2
	p.AddResourceConfigurator("flexibleengine_vpc_peering_connection_accepter_v2", func(r *config.Resource) {

		// peering_id is the ID of the VPC Peering Connection
		r.References["vpc_peering_connection_id"] = config.Reference{
			Type: "PeeringConnection",
		}

	})

	// flexibleengine_networking_secgroup_rule_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_secgroup_rule_v2
	p.AddResourceConfigurator("flexibleengine_networking_secgroup_rule_v2", func(r *config.Resource) {

		// remote_group_id is the ID of the security group to which this rule applies.
		r.References["remote_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	// flexibleengine_vpc_vip_associate_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_vip_associate_v2
	p.AddResourceConfigurator("flexibleengine_networking_vip_associate_v2", func(r *config.Resource) {
		// vip_id is the ID of the virtual IP address.
		r.References["vip_id"] = config.Reference{
			Type: "VIP",
		}
		// port_ids is the ID of the port to which the virtual IP address is associated.
		r.References["port_ids"] = config.Reference{
			Type: "Port",
		}
	})

	// flexibleengine_networking_port_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_port_v2
	p.AddResourceConfigurator("flexibleengine_networking_port_v2", func(r *config.Resource) {
		// fixed_ip.[].subnet_id is the ID of the subnet to which this port belongs.
		r.References["fixed_ip.subnet_id"] = config.Reference{
			Type: "VPCSubnet",
		}
	})

	// flexibleengine_networking_vip_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_vip_v2
	p.AddResourceConfigurator("flexibleengine_networking_vip_v2", func(r *config.Resource) {
		// subnet_id is DEPRECATED and generate conflict with ip_version
		config.MoveToStatus(r.TerraformResource, "subnet_id")
	})
}
