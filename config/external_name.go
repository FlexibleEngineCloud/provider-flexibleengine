// Package config contains the configuration for FlexibleEngine external name resources.
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
		> Anti-DDoS (antiddos)
	*/

	// flexibleengine_antiddos_v1 - Imported using the Template (floating_ip_id)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/antiddos_v1
	"flexibleengine_antiddos_v1": TemplatedStringAsIdentifierWithNoName("{{ .parameters.floating_ip_id }}"),

	/*
		> Relational Database Service (RDS)
	*/

	// flexibleengine_rds_instance_v3 - Imported using the Instance ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_instance_v3
	"flexibleengine_rds_instance_v3": config.IdentifierFromProvider,

	// flexibleengine_rds_account - Imported using the Template (instance_id/username)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_account
	"flexibleengine_rds_account": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.name }}"),

	// flexibleengine_rds_read_replica_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_read_replica_v3
	"flexibleengine_rds_read_replica_v3": config.IdentifierFromProvider,

	// flexibleengine_rds_database - Imported using the Template (instance_id/db_name)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_database
	"flexibleengine_rds_database": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.name }}"),

	// flexibleengine_rds_parametergroup_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_parametergroup_v3
	"flexibleengine_rds_parametergroup_v3": config.IdentifierFromProvider,

	// flexibleengine_rds_database_privilege - Imported using the Template (instance_id/db_name)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_database_privilege
	"flexibleengine_rds_database_privilege": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.db_name }}"),

	/*
		> AI Development Platform (ModelArts)
	*/

	// flexibleengine_modelarts_dataset - Imported using the Dataset ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/modelarts_dataset
	"flexibleengine_modelarts_dataset": config.IdentifierFromProvider,

	// flexibleengine_modelarts_dataset_version - Imported using the Dataset Version ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/modelarts_dataset_version
	"flexibleengine_modelarts_dataset_version": config.IdentifierFromProvider,

	/*
		> Enterprise Project Management Service (EPS)
	*/

	// flexibleengine_enterprise_project - Imported using the Enterprise Project ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/enterprise_project
	"flexibleengine_enterprise_project": config.IdentifierFromProvider,

	/*
		> Data Replication Service (DRS)
	*/

	// flexibleengine_drs_job - Imported using the Job ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/drs_job
	"flexibleengine_drs_job": config.IdentifierFromProvider,

	/*
		> Distributed Cache Service (DCS)
	*/

	// flexibleengine_dcs_instance_v1 Imported using the Id
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dcs_instance_v1
	"flexibleengine_dcs_instance_v1": config.IdentifierFromProvider,

	/*
		> API Gateway Dedicated (AGD)
	*/

	// flexibleengine_apig_api - Imported using the Template (name/id)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_api
	"flexibleengine_apig_api": TemplatedStringAsIdentifierWithNoName("{{ .parameters.name }}/{{ .parameters.instance_id }}"),

	// flexibleengine_apig_throttling_policy_associate - Imported using the Template (<instance id>/<policy_id>)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_throttling_policy_associate
	"flexibleengine_apig_throttling_policy_associate": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.policy_id }}"),

	// flexibleengine_apig_environment - Imported using the Template (<instance id>/<environment id>)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_environment
	"flexibleengine_apig_environment": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .external_name }}"),

	// flexibleengine_apig_vpc_channel - Imported using the Template (<instance id>/<name>)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_vpc_channel
	"flexibleengine_apig_vpc_channel": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.name }}"),

	// flexibleengine_apig_application - Imported using the Template (<instance id>/<id>)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_application
	"flexibleengine_apig_application": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .external_name }}"),

	// flexibleengine_apig_response - Imported using the Template (<instance id>/<api id>/<response id>)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_response
	"flexibleengine_apig_response": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.group_id }}/{{ .external_name }}"),

	// flexibleengine_apig_throttling_policy - Imported using the Template (<instance id>/<policy id>)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_throttling_policy
	"flexibleengine_apig_throttling_policy": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .external_name }}"),

	// flexibleengine_apig_instance - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_instance
	"flexibleengine_apig_instance": config.IdentifierFromProvider,

	// flexibleengine_apig_group - Imported using the Template (<instance id>/<group id>)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_group
	"flexibleengine_apig_group": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .external_name }}"),

	// flexibleengine_apig_custom_authorizer - Imported using the Template (<instance id>/name)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_custom_authorizer
	"flexibleengine_apig_custom_authorizer": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.name }}"),

	// flexibleengine_apig_api_publishment - Imported using the Template (<instance id>/<env id>/<api id>)
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_api_publishment
	"flexibleengine_apig_api_publishment": TemplatedStringAsIdentifierWithNoName("{{ .parameters.instance_id }}/{{ .parameters.env_id }}/{{ .parameters.api_id }}"),

	/*
		> API Gateway (AG)
	*/

	// flexibleengine_api_gateway_api - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/api_gateway_api
	"flexibleengine_api_gateway_api": config.IdentifierFromProvider,

	// flexibleengine_api_gateway_group - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/api_gateway_group
	"flexibleengine_api_gateway_group": config.IdentifierFromProvider,

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

	// flexibleengine_dns_recordset_v2 - Imported using {zone_id}/{recordset_id}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dns_recordset_v2
	"flexibleengine_dns_recordset_v2": TemplatedStringAsIdentifierWithNoName("{{ .parameters.zone_id }}/{{ .external_name }}"),

	// flexibleengine_dns_ptrrecord_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dns_ptrrecord_v2
	"flexibleengine_dns_ptrrecord_v2": TemplatedStringAsIdentifierWithNoName("{{ .parameters.region }}:{{ .external_name }}"),

	/*
		> Cloud Container Engine (CCE)
	*/
	// flexibleengine_cce_addon_v3 - Imported using {cluster_id}/{addon_id}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_addon_v3
	"flexibleengine_cce_addon_v3": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .external_name }}"),

	// flexibleengine_cce_cluster_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_cluster_v3
	"flexibleengine_cce_cluster_v3": config.IdentifierFromProvider,

	// flexibleengine_cce_namespace - Imported using {cluster_id}/{namespace}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_namespace
	"flexibleengine_cce_namespace": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .parameters.name }}"),

	// flexibleengine_cce_node_pool_v3 - Imported using {cluster_id}/{node_pool_id}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_node_pool_v3
	"flexibleengine_cce_node_pool_v3": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .external_name }}"),

	// flexibleengine_cce_pvc - Imported using {cluster_id}/{namespace}/{name}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_pvc
	"flexibleengine_cce_pvc": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .parameters.namespace }}/{{ .external_name }}"),

	// flexibleengine_cce_node_v3 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_node_v3
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
	"flexibleengine_vpc_subnet_v1": {
		SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
		GetExternalNameFn: func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["ipv4_subnet_id"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find id in tfstate")
		},
		GetIDFn:                config.ExternalNameAsID,
		DisableNameInitializer: true,
	},

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

	// flexibleengine_networking_port_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_port_v2
	"flexibleengine_networking_port_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_secgroup_rule_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_secgroup_rule_v2
	"flexibleengine_networking_secgroup_rule_v2": config.IdentifierFromProvider,

	// flexibleengine_networking_secgroup_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_secgroup_v2
	"flexibleengine_networking_secgroup_v2": config.IdentifierFromProvider,

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
	// flexibleengine_lb_l7policy_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_l7policy_v2
	"flexibleengine_lb_l7policy_v2": config.IdentifierFromProvider,
	// flexibleengine_lb_l7rule_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_l7rule_v2
	"flexibleengine_lb_l7rule_v2": config.IdentifierFromProvider,
	// flexibleengine_lb_listener_v2 - No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_listener_v2
	"flexibleengine_lb_listener_v2": config.IdentifierFromProvider,
	// flexibleengine_lb_loadbalancer_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_loadbalancer_v2
	"flexibleengine_lb_loadbalancer_v2": config.IdentifierFromProvider,
	// flexibleengine_lb_member_v2 - No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_member_v2
	"flexibleengine_lb_member_v2": config.IdentifierFromProvider,
	// flexibleengine_lb_monitor_v2 - No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_monitor_v2
	"flexibleengine_lb_monitor_v2": config.IdentifierFromProvider,
	// flexibleengine_lb_pool_v2 - No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_pool_v2
	"flexibleengine_lb_pool_v2": config.IdentifierFromProvider,
	// flexibleengine_lb_whitelist_v2 - No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lb_whitelist_v2
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

	// flexibleengine_obs_bucket - Imported using ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket
	"flexibleengine_obs_bucket": config.IdentifierFromProvider,

	// flexibleengine_obs_bucket_object - No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_object
	"flexibleengine_obs_bucket_object": config.IdentifierFromProvider,

	// flexibleengine_obs_bucket_replication - Imported using the source bucket name
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_replication
	"flexibleengine_obs_bucket_replication": TemplatedStringAsIdentifierWithNoName("{{.parameters.bucket}}"),

	// flexibleengine_s3_bucket - Imported using ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket
	"flexibleengine_s3_bucket": config.IdentifierFromProvider,

	// flexibleengine_s3_bucket_object - No import documented
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket_object
	"flexibleengine_s3_bucket_object": config.IdentifierFromProvider,

	// flexibleengine_s3_bucket_policy - No import documented
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

	/*
	  > Storage Disaster Recovery Service (SDRS)
	*/
	// flexibleengine_sdrs_drill_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sdrs_drill_v1
	"flexibleengine_sdrs_drill_v1": config.IdentifierFromProvider,
	// flexibleengine_sdrs_protectedinstance_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sdrs_protectedinstance_v1
	"flexibleengine_sdrs_protectedinstance_v1": config.IdentifierFromProvider,
	// flexibleengine_sdrs_protectiongroup_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sdrs_protectiongroup_v1
	"flexibleengine_sdrs_protectiongroup_v1": config.IdentifierFromProvider,
	// flexibleengine_sdrs_replication_attach_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sdrs_replication_attach_v1
	"flexibleengine_sdrs_replication_attach_v1": config.IdentifierFromProvider,
	// flexibleengine_sdrs_replication_pair_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sdrs_replication_pair_v1
	"flexibleengine_sdrs_replication_pair_v1": config.IdentifierFromProvider,

	/*
		> Cloud Service Engine (CSE)
	*/
	// flexibleengine_cse_microservice_engine - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cse_microservice_engine
	"flexibleengine_cse_microservice_engine": config.IdentifierFromProvider,

	// flexibleengine_cse_microservice - Imported using the Template connect_address
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cse_microservice
	"flexibleengine_cse_microservice": TemplatedStringAsIdentifierWithNoName("{{.parameters.connect_address}}/{{.external_name}}"),

	// flexibleengine_cse_microservice_instance - Imported using the Template connect_address, microservice_id and their id
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cse_microservice_instance
	"flexibleengine_cse_microservice_instance": TemplatedStringAsIdentifierWithNoName("{{.parameters.connect_address}}/{{.parameters.microservice_id}}/{{.external_name}}"),

	/*
		> Cloud Eye (CES)
	*/

	// flexibleengine_ces_alarmrule - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/ces_alarmrule
	"flexibleengine_ces_alarmrule": config.IdentifierFromProvider,

	/*
		> Log Tanker Service (LTS)
	*/

	// flexibleengine_lts_topic - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lts_topic
	"flexibleengine_lts_topic": TemplatedStringAsIdentifierWithNoName("{{.parameters.group_id}}/{{.external_name}}"),

	// flexibleengine_lts_group - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lts_group
	"flexibleengine_lts_group": config.IdentifierFromProvider,

	/*
		> Ressource Template Service (RTS)
	*/
	// flexibleengine_rts_software_config_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rts_software_config_v1
	"flexibleengine_rts_software_config_v1": config.IdentifierFromProvider,

	// flexibleengine_rts_stack_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rts_stack_v1
	"flexibleengine_rts_stack_v1": TemplatedStringAsIdentifierWithNoName("{{.parameters.name}}"),

	/*
	  > Distributed Message Service (DMS)
	*/

	// flexibleengine_dms_kafka_instance - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dms_kafka_instance
	"flexibleengine_dms_kafka_instance": config.IdentifierFromProvider,

	// flexibleengine_dms_kafka_topic - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dms_kafka_topic
	"flexibleengine_dms_kafka_topic": config.IdentifierFromProvider,

	// flexibleengine_dms_kafka_user  - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dms_kafka_user
	"flexibleengine_dms_kafka_user": config.IdentifierFromProvider,

	/*
	  > Cloud Search Service (CSS)
	*/
	// flexibleengine_css_cluster_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/css_cluster_v1
	"flexibleengine_css_cluster_v1": config.IdentifierFromProvider,
	// flexibleengine_css_snapshot_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/css_snapshot_v1
	"flexibleengine_css_snapshot_v1": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .external_name }}"),

	/*
	  > Cloud  Server Backup Service (CSBS)
	*/

	// flexibleengine_csbs_backup_policy_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/csbs_backup_policy_v1
	"flexibleengine_csbs_backup_policy_v1": config.IdentifierFromProvider,

	// flexibleengine_csbs_backup_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/csbs_backup_v1
	"flexibleengine_csbs_backup_v1": config.IdentifierFromProvider,

	/*
		> Scalable File Service (SFS)
	*/
	// flexibleengine_sfs_access_rule_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sfs_access_rule_v2
	"flexibleengine_sfs_access_rule_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.sfs_id}}/{{.external_name}}"),
	// flexibleengine_sfs_file_system_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sfs_file_system_v2
	"flexibleengine_sfs_file_system_v2": config.IdentifierFromProvider,
	// flexibleengine_sfs_share_access_rule_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sfs_share_access_rule_v2
	"flexibleengine_sfs_turbo": config.IdentifierFromProvider,

	/*
	  > Machine Learning Service (MLS)
	*/

	// flexibleengine_mls_instance_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/csbs_backup_policy_v1
	"flexibleengine_mls_instance_v1": config.IdentifierFromProvider,

	/*
	  >  Cloud Backup and Recovery (CBR)
	*/

	// flexibleengine_cbr_vault - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cbr_vault
	"flexibleengine_cbr_vault": config.IdentifierFromProvider,

	// flexibleengine_cbr_policy - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cbr_policy
	"flexibleengine_cbr_policy": config.IdentifierFromProvider,

	/*
	  >  Auto Scaling (AS)
	*/
	// flexibleengine_as_configuration_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/as_configuration_v1
	"flexibleengine_as_configuration_v1": config.IdentifierFromProvider,
	// flexibleengine_as_group_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/as_group_v1
	"flexibleengine_as_group_v1": config.IdentifierFromProvider,
	// flexibleengine_as_policy_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/as_policy_v1
	"flexibleengine_as_policy_v1": config.IdentifierFromProvider,
	// flexibleengine_as_lifecycle_hook_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/as_lifecycle_hook_v1
	"flexibleengine_as_lifecycle_hook_v1": config.TemplatedStringAsIdentifier("name", "{{.parameters.scaling_group_id}}/{{.external_name}}"),

	/*
		> MapReduce Service (MRS)
	*/
	// flexibleengine_mrs_cluster_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_cluster_v1
	"flexibleengine_mrs_cluster_v1": config.IdentifierFromProvider,

	// flexibleengine_mrs_cluster_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_cluster_v2
	"flexibleengine_mrs_cluster_v2": config.IdentifierFromProvider,

	// flexibleengine_mrs_hybrid_cluster_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_hybrid_cluster_v1
	"flexibleengine_mrs_hybrid_cluster_v1": config.IdentifierFromProvider,

	// flexibleengine_mrs_job_v1 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_job_v1
	"flexibleengine_mrs_job_v1": config.IdentifierFromProvider,

	// flexibleengine_mrs_job_v2 - Imported using template
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_job_v2
	"flexibleengine_mrs_job_v2": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .external_name }}"),

	/*
		> Tag Management Service (TMS)
	*/
	// flexibleengine_tms_tags - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/tms_tags
	"flexibleengine_tms_tags": config.IdentifierFromProvider,

	/*
		> Data Lake Insight (DLI)
	*/
	// flexibleengine_dli_queue - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dli_queue_v1
	"flexibleengine_dli_queue": config.IdentifierFromProvider,
	// flexibleengine_dli_package - No import documentation
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dli_package
	"flexibleengine_dli_package": config.IdentifierFromProvider,
	// flexibleengine_dli_database - No import documentation
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dli_database
	"flexibleengine_dli_database": config.IdentifierFromProvider,
	// flexibleengine_dli_flinksql_job - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dli_flinksql_job
	"flexibleengine_dli_flinksql_job": config.IdentifierFromProvider,
	// flexibleengine_dli_spark_job - No import documentation
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dli_spark_job
	"flexibleengine_dli_spark_job": config.IdentifierFromProvider,
	// flexibleengine_dli_table - Imported using {database_name}/{table_name}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dli_table
	"flexibleengine_dli_table": TemplatedStringAsIdentifierWithNoName("{{.parameters.database_name}}/{{.parameters.name}}"),

	/*
	 > Simple Message Notification (SMN)
	*/
	// flexibleengine_smn_topic_v2 - No import documentation
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/smn_topic_v2
	"flexibleengine_smn_topic_v2": config.IdentifierFromProvider,
	// flexibleengine_smn_subscription_v2 - No import documentation
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/smn_subscription_v2
	"flexibleengine_smn_subscription_v2": config.IdentifierFromProvider,

	/*
		> Network ACL
	*/
	// flexibleengine_fw_firewall_group_v2- Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fw_firewall_group_v2
	"flexibleengine_fw_firewall_group_v2": config.IdentifierFromProvider,

	// flexibleengine_fw_policy_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fw_policy_v2
	"flexibleengine_fw_policy_v2": config.IdentifierFromProvider,

	// flexibleengine_fw_rule_v2 - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fw_rule_v2
	"flexibleengine_fw_rule_v2": config.IdentifierFromProvider,

	// flexibleengine_network_acl - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/network_acl
	"flexibleengine_network_acl": config.IdentifierFromProvider,

	// flexibleengine_network_acl_rule - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/network_acl_rule
	"flexibleengine_network_acl_rule": config.IdentifierFromProvider,

	/*
		>  Data Ingestion Service (DIS)
	*/
	// flexibleengine_dis_stream - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dis_stream
	"flexibleengine_dis_stream": config.IdentifierFromProvider,

	/*
		> FunctionGraph
	*/

	// flexibleengine_fgs_function - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fgs_function
	"flexibleengine_fgs_function": config.IdentifierFromProvider,

	// flexibleengine_fgs_dependency - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fgs_dependency
	"flexibleengine_fgs_dependency": config.IdentifierFromProvider,

	// flexibleengine_fgs_trigger - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fgs_trigger
	"flexibleengine_fgs_trigger": config.IdentifierFromProvider,

	/*
		> Data Warehouse (DWS)
	*/
	// flexibleengine_dws_cluster_v1 - No import documentation
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dws_cluster_v1
	"flexibleengine_dws_cluster_v1": config.IdentifierFromProvider,
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
