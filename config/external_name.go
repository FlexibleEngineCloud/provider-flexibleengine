/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{

	/*
		> Document Database Service (DDS)
	*/

	// flexibleengine_dds_instance_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dds_instance_v3
	"flexibleengine_dds_instance_v3": config.IdentifierFromProvider,

	// flexibleengine_dds_database_user - Imported using the Template (instance-id/db-name/username)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dds_database_user
	"flexibleengine_dds_database_user": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.db_name }}/{{ .external_name }}"),

	// flexibleengine_dds_database_role - Imported using the Template (instance-id/db-name/role-name)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dds_database_role
	"flexibleengine_dds_database_role": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.db_name }}/{{ .external_name }}"),

	/*
		> Software Repository for Container (SWR)
	*/

	// flexibleengine_swr_organization - Imported using the Name
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/swr_organization
	"flexibleengine_swr_organization": config.NameAsIdentifier,

	// flexibleengine_swr_organization_users - Imported using the ID (ID is name of resource 0_o )
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/swr_organization_users
	"flexibleengine_swr_organization_users": config.IdentifierFromProvider,

	// flexibleengine_swr_repository - Imported using the Template (org-name/repo-name)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/swr_repository
	"flexibleengine_swr_repository": TemplatedStringAsIdentifierWithNoName("{{ .parameters.organization }}/{{ .external_name }}"),

	// flexibleengine_swr_repository_sharing - Imported using the Template (org-name/repo-name/sharing-account)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/swr_repository_sharing
	"flexibleengine_swr_repository_sharing": TemplatedStringAsIdentifierWithNoName("{{ .parameters.organization }}/{{ .parameters.repository }}/{{ .external_name }}"),

	/*
		> Bare Metal Server (BMS)
	*/

	// flexibleengine_compute_bms_server_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_bms_server_v2
	"flexibleengine_compute_bms_server_v2": config.IdentifierFromProvider,

	/*
		> Elastic Volume Service (EVS)
	*/

	// flexibleengine_blockstorage_volume_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/blockstorage_volume_v2
	"flexibleengine_blockstorage_volume_v2": config.IdentifierFromProvider,

	/*
		> Volume Backup Service (VBS)
	*/

	// flexibleengine_vbs_backup_policy_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vbs_backup_policy_v2
	"flexibleengine_vbs_backup_policy_v2": config.IdentifierFromProvider,

	// flexibleengine_vbs_backup_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vbs_backup_v2
	"flexibleengine_vbs_backup_v2": config.IdentifierFromProvider,

	/*
		> Web Application Firewall (WAF)
	*/

	// flexibleengine_waf_certificate - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_certificate
	"flexibleengine_waf_certificate": config.IdentifierFromProvider,

	// flexibleengine_waf_dedicated_certificate - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_dedicated_certificate
	"flexibleengine_waf_dedicated_certificate": config.IdentifierFromProvider,

	// flexibleengine_waf_dedicated_domain - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_dedicated_domain
	"flexibleengine_waf_dedicated_domain": config.IdentifierFromProvider,

	// flexibleengine_waf_dedicated_instance - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_dedicated_instance
	"flexibleengine_waf_dedicated_instance": config.IdentifierFromProvider,

	// flexibleengine_waf_dedicated_policy - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_dedicated_policy
	"flexibleengine_waf_dedicated_policy": config.IdentifierFromProvider,

	// flexibleengine_waf_domain - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_domain
	"flexibleengine_waf_domain": config.IdentifierFromProvider,

	// flexibleengine_waf_policy - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_policy
	"flexibleengine_waf_policy": config.IdentifierFromProvider,

	// flexibleengine_waf_rule_alarm_masking - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_rule_alarm_masking
	"flexibleengine_waf_rule_alarm_masking": TemplatedStringAsIdentifierWithNoName("{{ .parameters.policy_id }}/{{ .external_name }}"),

	// flexibleengine_waf_rule_blacklist - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_rule_blacklist
	"flexibleengine_waf_rule_blacklist": TemplatedStringAsIdentifierWithNoName("{{ .parameters.policy_id }}/{{ .external_name }}"),

	// flexibleengine_waf_rule_cc_protection - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_rule_cc_protection
	"flexibleengine_waf_rule_cc_protection": TemplatedStringAsIdentifierWithNoName("{{ .parameters.policy_id }}/{{ .external_name }}"),

	// flexibleengine_waf_rule_data_masking - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_rule_data_masking
	"flexibleengine_waf_rule_data_masking": TemplatedStringAsIdentifierWithNoName("{{ .parameters.policy_id }}/{{ .external_name }}"),

	// flexibleengine_waf_rule_precise_protection - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_rule_precise_protection
	"flexibleengine_waf_rule_precise_protection": TemplatedStringAsIdentifierWithNoName("{{ .parameters.policy_id }}/{{ .external_name }}"),

	// flexibleengine_waf_rule_web_tamper_protection - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_rule_web_tamper_protection
	"flexibleengine_waf_rule_web_tamper_protection": TemplatedStringAsIdentifierWithNoName("{{ .parameters.policy_id }}/{{ .external_name }}"),

	/*
		> Domain Name Service (DNS)
	*/

	// flexibleengine_dns_zone_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dns_zone_v2
	"flexibleengine_dns_zone_v2": config.IdentifierFromProvider,

	// flexibleengine_dns_recordset_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dns_recordset_v2
	"flexibleengine_dns_recordset_v2": TemplatedStringAsIdentifierWithNoName("{{ .parameters.zone_id }}/{{ .parameters.recordset_id }}"),

	// flexibleengine_dns_ptrrecord_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dns_ptrrecord_v2
	// "flexibleengine_dns_ptrrecord_v2": FormattedIdentifierFromProvider(":", "flexibleengine_vpc_eip"),
	/*
		> Cloud Container Engine (CCE)
	*/
	"flexibleengine_cce_addon_v3":     TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .external_name }}"),
	"flexibleengine_cce_cluster_v3":   config.IdentifierFromProvider,
	"flexibleengine_cce_namespace":    TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .parameters.name }}"),
	"flexibleengine_cce_node_pool_v3": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .external_name }}"),
	"flexibleengine_cce_pvc":          TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .parameters.namespace }}/{{ .parameters.name }}"),

	// No import documented
	"flexibleengine_cce_node_v3": config.IdentifierFromProvider,
	/*
		> Elastic Cloud Server (ECS)
	*/

	// Imported with the following format : {floating_ip}/{instance_id}/{fixed_ip}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_floatingip_associate_v2
	"flexibleengine_compute_floatingip_associate_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.floating_ip}}/{{.parameters.instance_id}}/{{.parameters.fixed_ip}}"),
	// Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_instance_v2
	"flexibleengine_compute_instance_v2": config.IdentifierFromProvider,
	// Imported with the following format : {instance_id}/{port_id}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_interface_attach_v2
	"flexibleengine_compute_interface_attach_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.instance_id}}/{{.parameters.port_id}}"),
	// Imported using name
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_keypair_v2
	"flexibleengine_compute_keypair_v2": config.NameAsIdentifier,
	// Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_secgroup_v2
	"flexibleengine_compute_servergroup_v2": config.IdentifierFromProvider,
	// Imported with the following format : {instance_id}/{volume_id}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_volume_attach_v2
	"flexibleengine_compute_volume_attach_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.instance_id}}/{{.parameters.volume_id}}"),

	/*
		> Identity and Access Management (IAM)
	*/

	// Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_agency_v3
	"flexibleengine_identity_agency_v3": config.IdentifierFromProvider,
	// No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_agency_token_v3
	"flexibleengine_identity_group_membership_v3": config.IdentifierFromProvider,
	// Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_group_v3
	"flexibleengine_identity_group_v3": config.IdentifierFromProvider,
	// Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_project_v3
	"flexibleengine_identity_project_v3": config.IdentifierFromProvider,
	// Imported using name
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_provider_v3
	"flexibleengine_identity_provider": config.NameAsIdentifier,
	// Imported using provider_id
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_provider_conversion
	"flexibleengine_identity_provider_conversion": TemplatedStringAsIdentifierWithNoName("{{.parameters.provider_id}}"),
	// No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_role_assignment_v3
	"flexibleengine_identity_role_assignment_v3": config.IdentifierFromProvider,
	// Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_role_v3
	"flexibleengine_identity_role_v3": config.IdentifierFromProvider,
	// Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_user_v3
	"flexibleengine_identity_user_v3": config.IdentifierFromProvider,

	/*
		> Image Management Service (IMS)
	*/

	// flexibleengine_images_image_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/images_image_v2
	"flexibleengine_images_image_v2": config.IdentifierFromProvider,

	/*
		> Key Management Service (KMS)
	*/
	// flexibleengine_kms_key_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/kms_key_v1
	"flexibleengine_kms_key_v1": config.IdentifierFromProvider,

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

	/*
		> Object Storage Service (OSS)
	*/

	// flexibleengine_obs_bucket - Imported using the Name
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket
	"flexibleengine_obs_bucket": config.NameAsIdentifier,

	// flexibleengine_obs_bucket_object - Imported using ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_object
	"flexibleengine_obs_bucket_object": config.IdentifierFromProvider,

	// flexibleengine_obs_bucket_replication - Imported using the Name
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_replication
	"flexibleengine_obs_bucket_replication": config.NameAsIdentifier,

	// flexibleengine_s3_bucket - Imported using the Name
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket
	"flexibleengine_s3_bucket": config.NameAsIdentifier,

	// flexibleengine_s3_bucket_object - Imported using ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket_object
	"flexibleengine_s3_bucket_object": config.IdentifierFromProvider,

	// flexibleengine_s3_bucket_policy - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket_policy
	"flexibleengine_s3_bucket_policy": config.IdentifierFromProvider,

	/*
	  > NAT Gateway (NAT)
	*/

	// flexibleengine_nat_gateway_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/nat_gateway_v2
	"flexibleengine_nat_gateway_v2": config.IdentifierFromProvider,

	// flexibleengine_nat_dnat_rule_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/nat_dnat_rule_v2
	"flexibleengine_nat_dnat_rule_v2": config.IdentifierFromProvider,
	"flexibleengine_nat_snat_rule_v2": config.IdentifierFromProvider,
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

// FormattedIdentifierFromProvider is a helper function to construct Terraform
// IDs that use elements from the parameters in a certain string format.
// It should be used in cases where all information in the ID is gathered from
// the spec and not user defined like name. For example, zone_id:vpc_id.
func FormattedIdentifierFromProvider(separator string, keys ...string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys))
		for i, key := range keys {
			val, ok := parameters[key]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", key)
			}
			s, ok := val.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be string", key)
			}
			vals[i] = s
		}
		return strings.Join(vals, separator), nil
	}
	return e
}
