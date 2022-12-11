package config

import (
	"regexp"
	"strings"

	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/upjet/pkg/types/name"
)

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}

// KindRemoveVersion removes the version from the kind of the resource.
func KindRemoveVersion() config.ResourceOption {
	return func(r *config.Resource) {
		matchedName, _ := regexp.MatchString(`_v[0-9]+`, r.Name)
		matchedKind, _ := regexp.MatchString(`V[0-9]+`, r.Kind)

		if (r.Kind == "" && matchedName) || matchedKind {

			r.Kind = name.NewFromSnake(tools.RemoveGroup(tools.RemoveVersion(r.Name))).Camel

		}
	}
}

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		// words := strings.Split(strings.TrimPrefix(resource, "flexibleengine_"), "_")
		words := tools.RemoveVersion(resource)
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// GroupMap contains all overrides we'd like to make to the default group search.
// It's written with data from TF Provider AWS repo service grouping in here:
// https://github.com/hashicorp/terraform-provider-aws/tree/main/internal/service
//
// At the end, all of them are based on grouping of the AWS Go SDK.
// The initial grouping is calculated based on folder grouping of AWS TF Provider
// which is based on Go SDK. Here is the script used to fetch that list:
// https://gist.github.com/muvaf/8d61365ffc1df7757864422ba16d7819
// ! Warning do not consider _v[0-9]
// ? Examples
// ? "flexibleengine_vpc_subnet_v1": ReplaceGroupWords("vpc", 0), => Output: Group: vpc, Kind: Subnet
// ? "flexibleengine_networking_subnet_v2": ReplaceGroupWords("vpc", 0), => Output: Group: vpc, Kind: NetWorkingSubnet
var GroupMap = map[string]GroupKindCalculator{
	// IAM
	"flexibleengine_identity_agency_v3":           ReplaceGroupWords("iam", 1), // Group: iam, Kind: Agency
	"flexibleengine_identity_group_membership_v3": ReplaceGroupWords("iam", 1), // Group: iam, Kind: GroupMembership
	"flexibleengine_identity_group_v3":            ReplaceGroupWords("iam", 1), // Group: iam, Kind: Group
	"flexibleengine_identity_project_v3":          ReplaceGroupWords("iam", 1), // Group: iam, Kind: Project
	"flexibleengine_identity_provider":            ReplaceGroupWords("iam", 1), // Group: iam, Kind: Provider
	"flexibleengine_identity_provider_conversion": ReplaceGroupWords("iam", 1), // Group: iam, Kind: ProviderConversion
	"flexibleengine_identity_role_assignment_v3":  ReplaceGroupWords("iam", 1), // Group: iam, Kind: RoleAssignment
	"flexibleengine_identity_role_v3":             ReplaceGroupWords("iam", 1), // Group: iam, Kind: Role
	"flexibleengine_identity_user_v3":             ReplaceGroupWords("iam", 1), // Group: iam, Kind: User
	// IMS
	"flexibleengine_images_image_v2": ReplaceGroupWords("ims", 1), // Group: ims, Kind: Image
	// ECS
	"flexibleengine_compute_floatingip_associate_v2": ReplaceGroupWords("ecs", 1), // Group: ecs, Kind: FloatingipAssociate (! Rewrite in KingMap)
	"flexibleengine_compute_instance_v2":             ReplaceGroupWords("ecs", 1), // Group: ecs, Kind: Instance
	"flexibleengine_compute_interface_attach_v2":     ReplaceGroupWords("ecs", 1), // Group: ecs, Kind: InterfaceAttach
	"flexibleengine_compute_keypair_v2":              ReplaceGroupWords("ecs", 1), // Group: ecs, Kind: Keypair (! Rewrite in KingMap)
	"flexibleengine_compute_servergroup_v2":          ReplaceGroupWords("ecs", 1), // Group: ecs, Kind: Servergroup (! Rewrite in KingMap)
	"flexibleengine_compute_volume_attach_v2":        ReplaceGroupWords("ecs", 1), // Group: ecs, Kind: VolumeAttach
	// VPC
	"flexibleengine_vpc_v1":                         ReplaceGroupWords("vpc", 0), // Group: vpc, Kind: VPC
	"flexibleengine_vpc_subnet_v1":                  ReplaceGroupWords("vpc", 0), // Group: vpc, Kind: VPCSubnet
	"flexibleengine_networking_network_v2":          ReplaceGroupWords("vpc", 1), // Group: vpc, Kind: Network
	"flexibleengine_networking_port_v2":             ReplaceGroupWords("vpc", 1), // Group: vpc, Kind: Port
	"flexibleengine_networking_router_interface_v2": ReplaceGroupWords("vpc", 1), // Group: vpc, Kind: RouterInterface
	"flexibleengine_networking_router_v2":           ReplaceGroupWords("vpc", 1), // Group: vpc, Kind: Router
	"flexibleengine_networking_secgroup_rule_v2":    ReplaceGroupWords("vpc", 1), // Group: vpc, Kind: SecurityGroupRule
	"flexibleengine_networking_secgroup_v2":         ReplaceGroupWords("vpc", 1), // Group: vpc, Kind: SecurityGroup
	"flexibleengine_networking_subnet_v2":           ReplaceGroupWords("vpc", 0), // Group: vpc, Kind: NetworkingSubnet
	"flexibleengine_networking_vip_associate_v2":    ReplaceGroupWords("vpc", 1), // Group: vpc, Kind: VipAssociate
	"flexibleengine_networking_vip_v2":              ReplaceGroupWords("vpc", 1), // Group: vpc, Kind: Vip
	// EIP
	"flexibleengine_vpc_eip":           ReplaceGroupWords("eip", 1), // Group: eip, Kind: EIP
	"flexibleengine_vpc_eip_associate": ReplaceGroupWords("eip", 1), // Group: eip, Kind: Associate
	// Dedicated ELB
	"flexibleengine_elb_certificate":    ReplaceGroupWords("dedicatedelb", 1), // Group: dedicatedelb, Kind: Certificate
	"flexibleengine_elb_ipgroup":        ReplaceGroupWords("dedicatedelb", 1), // Group: dedicatedelb, Kind: Ipgroup (! Rewrite in KingMap)
	"flexibleengine_lb_listener_v3":     ReplaceGroupWords("dedicatedelb", 1), // Group: dedicatedelb, Kind: Listener
	"flexibleengine_lb_loadbalancer_v3": ReplaceGroupWords("dedicatedelb", 1), // Group: dedicatedelb, Kind: Loadbalancer (! Rewrite in KingMap)
	"flexibleengine_lb_member_v3":       ReplaceGroupWords("dedicatedelb", 1), // Group: dedicatedelb, Kind: Member
	"flexibleengine_lb_monitor_v3":      ReplaceGroupWords("dedicatedelb", 1), // Group: dedicatedelb, Kind: Monitor
	"flexibleengine_lb_pool_v3":         ReplaceGroupWords("dedicatedelb", 1), // Group: dedicatedelb, Kind: Pool
	// ELB
	"flexibleengine_lb_l7policy_v2":     ReplaceGroupWords("elb", 1), // Group: elb, Kind: L7policy (! Rewrite in KingMap)
	"flexibleengine_lb_l7rule_v2":       ReplaceGroupWords("elb", 1), // Group: elb, Kind: L7rule (! Rewrite in KingMap)
	"flexibleengine_lb_listener_v2":     ReplaceGroupWords("elb", 1), // Group: elb, Kind: Listener
	"flexibleengine_lb_loadbalancer_v2": ReplaceGroupWords("elb", 1), // Group: elb, Kind: Loadbalancer (! Rewrite in KingMap)
	"flexibleengine_lb_member_v2":       ReplaceGroupWords("elb", 1), // Group: elb, Kind: Member
	"flexibleengine_lb_monitor_v2":      ReplaceGroupWords("elb", 1), // Group: elb, Kind: Monitor
	"flexibleengine_lb_pool_v2":         ReplaceGroupWords("elb", 1), // Group: elb, Kind: Pool
	"flexibleengine_lb_whitelist_v2":    ReplaceGroupWords("elb", 1), // Group: elb, Kind: Whitelist
	// OSS
	"flexibleengine_obs_bucket":             ReplaceGroupWords("oss", 0), // Group: oss, Kind: OBSBucket
	"flexibleengine_obs_bucket_object":      ReplaceGroupWords("oss", 0), // Group: oss, Kind: OBSBucketObject
	"flexibleengine_obs_bucket_replication": ReplaceGroupWords("oss", 0), // Group: oss, Kind: OBSBucketReplication
	"flexibleengine_s3_bucket":              ReplaceGroupWords("oss", 0), // Group: oss, Kind: S3Bucket
	"flexibleengine_s3_bucket_object":       ReplaceGroupWords("oss", 0), // Group: oss, Kind: S3Object
	"flexibleengine_s3_bucket_policy":       ReplaceGroupWords("oss", 0), // Group: oss, Kind: S3BucketPolicy

	// EVS
	"flexibleengine_blockstorage_volume_v2": ReplaceGroupWords("evs", 0), // Group: evs, Kind: Volume

	// BMS
	"flexibleengine_compute_bms_server_v2": ReplaceGroupWords("bms", 2), // Group: bms, Kind: Server

	// AG
	"flexibleengine_api_gateway_api":   ReplaceGroupWords("ag", 2), // Group: ag, Kind: API
	"flexibleengine_api_gateway_group": ReplaceGroupWords("ag", 2), // Group: ag, Kind: Group

	// AGD
	"flexibleengine_apig_api":                         ReplaceGroupWords("agd", 1), // Group: agd, Kind: API
	"flexibleengine_apig_throttling_policy_associate": ReplaceGroupWords("agd", 1), // Group: agd, Kind: ThrottlingPolicyAssociate
	"flexibleengine_apig_environment":                 ReplaceGroupWords("agd", 1), // Group: agd, Kind: Environment
	"flexibleengine_apig_vpc_channel":                 ReplaceGroupWords("agd", 1), // Group: agd, Kind: VpcChannel
	"flexibleengine_apig_application":                 ReplaceGroupWords("agd", 1), // Group: agd, Kind: Application
	"flexibleengine_apig_response":                    ReplaceGroupWords("agd", 1), // Group: agd, Kind: Response
	"flexibleengine_apig_throttling_policy":           ReplaceGroupWords("agd", 1), // Group: agd, Kind: ThrottlingPolicy
	"flexibleengine_apig_instance":                    ReplaceGroupWords("agd", 1), // Group: agd, Kind: Instance
	"flexibleengine_apig_group":                       ReplaceGroupWords("agd", 1), // Group: agd, Kind: Group
	"flexibleengine_apig_custom_authorizer":           ReplaceGroupWords("agd", 1), // Group: agd, Kind: CustomAuthorizer
	"flexibleengine_apig_api_publishment":             ReplaceGroupWords("agd", 1), // Group: agd, Kind: ApiPublishment

	// EPS
	"flexibleengine_enterprise_project": ReplaceGroupWords("eps", 1), // Group: eps, Kind: Project
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	"flexibleengine_compute_keypair_v2":              "KeyPair",
	"flexibleengine_compute_servergroup_v2":          "ServerGroup",
	"flexibleengine_compute_floatingip_associate_v2": "FloatingIpAssociate",
	"flexibleengine_networking_secgroup_v2":          "SecGroup",
	"flexibleengine_networking_secgroup_rule_v2":     "SecGroupRule",
	"flexibleengine_lb_loadbalancer_v3":              "LoadBalancer",
	"flexibleengine_elb_ipgroup":                     "IPGroup",
	"flexibleengine_lb_loadbalancer_v2":              "LoadBalancer",
	"flexibleengine_lb_l7policy_v2":                  "L7Policy",
	"flexibleengine_lb_l7rule_v2":                    "L7Rule",
	"flexibleengine_obs_bucket":                      "OBSBucket",
	"flexibleengine_obs_bucket_object":               "OBSBucketObject",
	"flexibleengine_obs_bucket_replication":          "OBSBucketReplication",
	"flexibleengine_waf_rule_blacklist":              "RuleBlacklist",
	"flexibleengine_blockstorage_volume_v2":          "BlockStorageVolume",
}
