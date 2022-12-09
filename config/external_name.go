/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{

	/*
		> Elastic Cloud Server (ECS)
	*/

	// Imported with the following format : {floating_ip}/{instance_id}/{fixed_ip}
	"flexibleengine_compute_floatingip_associate_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.floating_ip}}/{{.parameters.instance_id}}/{{.parameters.fixed_ip}}"),
	// Imported using the ID
	"flexibleengine_compute_instance_v2": config.IdentifierFromProvider,
	// Imported with the following format : {instance_id}/{port_id}
	"flexibleengine_compute_interface_attach_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.instance_id}}/{{.parameters.port_id}}"),
	// Imported using name
	"flexibleengine_compute_keypair_v2": config.NameAsIdentifier,
	// Imported using the ID
	"flexibleengine_compute_servergroup_v2": config.IdentifierFromProvider,
	// Imported with the following format : {instance_id}/{volume_id}
	"flexibleengine_compute_volume_attach_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.instance_id}}/{{.parameters.volume_id}}"),

	/*
		> Identity and Access Management (IAM)
	*/

	// Imported using the ID
	"flexibleengine_identity_agency_v3": config.IdentifierFromProvider,
	// No import documented
	"flexibleengine_identity_group_membership_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_group_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_project_v3": config.IdentifierFromProvider,
	// Imported using name
	"flexibleengine_identity_provider": config.NameAsIdentifier,
	// Imported using provider_id
	"flexibleengine_identity_provider_conversion": TemplatedStringAsIdentifierWithNoName("{{.parameters.provider_id}}"),
	// No import documented
	"flexibleengine_identity_role_assignment_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_role_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_user_v3": config.IdentifierFromProvider,

	/*
		> Virtual Private Cloud (VPC)
	*/

	// flexibleengine_vpc_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_v1
	"flexibleengine_vpc_v1": config.IdentifierFromProvider,

	// flexibleengine_vpc_subnet_v1 - Imported using the subnet id
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_subnet_v1
	"flexibleengine_vpc_subnet_v1": config.IdentifierFromProvider,

	// flexibleengine_vpc_route_table - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_route_table
	"flexibleengine_vpc_route_table": config.IdentifierFromProvider,

	// flexibleengine_vpc_route - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_route
	"flexibleengine_vpc_route": config.IdentifierFromProvider,

	// flexibleengine_vpc_peering_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_peering_v2
	"flexibleengine_vpc_peering_connection_v2": config.IdentifierFromProvider,

	// flexibleengine_vpc_peering_connection_accepter_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_peering_accepter_v2
	"flexibleengine_vpc_peering_connection_accepter_v2": config.IdentifierFromProvider,

	// flexibleengine_vpc_flow_log_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_flow_log_v1
	"flexibleengine_vpc_flow_log_v1": config.IdentifierFromProvider,

	// ! THIS IS DEPRECATED
	// ! Now use flexibleengine_vpc_eip
	// flexibleengine_networking_floatingip_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_floatingip_v2
	// "flexibleengine_networking_floatingip_v2": config.IdentifierFromProvider,

	// ! THIS IS NOT ANNOUNCED DEPRECATED BUT A PARENT RESOURCE IS DEPRECATED
	// ! Now use flexibleengine_vpc_eip_associate
	// flexibleengine_networking_floatingip_associate_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_floatingip_associate_v2
	// "flexibleengine_networking_floatingip_associate_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_network_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_network_v2
	"flexibleengine_networking_network_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_port_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_port_v2
	"flexibleengine_networking_port_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_router_interface_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_router_interface_v2
	"flexibleengine_networking_router_interface_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_router_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_router_v2
	"flexibleengine_networking_router_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_secgroup_rule_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_secgroup_rule_v2
	"flexibleengine_networking_secgroup_rule_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_secgroup_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_secgroup_v2
	"flexibleengine_networking_secgroup_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_subnet_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_subnet_v2
	"flexibleengine_networking_subnet_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_vip_associate_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_vip_associate_v2
	"flexibleengine_networking_vip_associate_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_vip_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_vip_v2
	"flexibleengine_networking_vip_v2": config.IdentifierFromProvider,

	/*
		> Elastic IP (EIP)
	*/

	// flexibleengine_vpc_eip - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_eip
	"flexibleengine_vpc_eip": config.IdentifierFromProvider,

	// flexibleengine_vpc_eip_associate - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_eip_associate
	"flexibleengine_vpc_eip_associate": config.IdentifierFromProvider,

	/*
		> Elastic Load Balance (Dedicated ELB)
	*/

	// flexibleengine_elb_certificate - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/elb_certificate
	"flexibleengine_elb_certificate": config.IdentifierFromProvider,

	// flexibleengine_elb_ipgroup - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/elb_ipgroup
	"flexibleengine_elb_ipgroup": config.IdentifierFromProvider,

	// flexibleengine_lb_listener_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_listener_v3
	"flexibleengine_lb_listener_v3": config.IdentifierFromProvider,

	// flexibleengine_lb_loadbalancer_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_loadbalancer_v3
	"flexibleengine_lb_loadbalancer_v3": config.IdentifierFromProvider,

	// flexibleengine_lb_member_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_member_v3
	"flexibleengine_lb_member_v3": config.IdentifierFromProvider,

	// flexibleengine_lb_monitor_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_monitor_v3
	"flexibleengine_lb_monitor_v3": config.IdentifierFromProvider,

	// flexibleengine_lb_pool_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_pool_v3
	"flexibleengine_lb_pool_v3": config.IdentifierFromProvider,

	/*
		> Elastic Load Balance (ELB)
	*/

	"flexibleengine_lb_l7policy_v2": config.IdentifierFromProvider,
	"flexibleengine_lb_l7rule_v2":   config.IdentifierFromProvider,
	// No import documented
	"flexibleengine_lb_listener_v2":     config.IdentifierFromProvider,
	"flexibleengine_lb_loadbalancer_v2": config.IdentifierFromProvider,
	// No import documented
	"flexibleengine_lb_member_v2": config.IdentifierFromProvider,
	// No import documented
	"flexibleengine_lb_monitor_v2": config.IdentifierFromProvider,
	// No import documented
	"flexibleengine_lb_pool_v2": config.IdentifierFromProvider,
	// No import documented
	"flexibleengine_lb_whitelist_v2": config.IdentifierFromProvider,

	/*
		> VPC Endpoint (VPCEP)
	*/

	// flexibleengine_vpcep_service - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_service
	"flexibleengine_vpcep_service": config.IdentifierFromProvider,

	// flexibleengine_vpcep_endpoint - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_endpoint
	"flexibleengine_vpcep_endpoint": config.IdentifierFromProvider,

	// flexibleengine_vpcep_approval - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_approval
	"flexibleengine_vpcep_approval": config.IdentifierFromProvider,
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {

		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
