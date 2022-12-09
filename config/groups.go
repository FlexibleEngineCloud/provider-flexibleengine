package config

import (
	"strings"

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

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		words := strings.Split(strings.TrimPrefix(resource, "flexibleengine_"), "_")
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
var GroupMap = map[string]GroupKindCalculator{
	"flexibleengine_identity_agency_v3":              ReplaceGroupWords("iam", 1),
	"flexibleengine_identity_group_membership_v3":    ReplaceGroupWords("iam", 1),
	"flexibleengine_identity_group_v3":               ReplaceGroupWords("iam", 1),
	"flexibleengine_identity_project_v3":             ReplaceGroupWords("iam", 1),
	"flexibleengine_identity_provider":               ReplaceGroupWords("iam", 1),
	"flexibleengine_identity_provider_conversion":    ReplaceGroupWords("iam", 1),
	"flexibleengine_identity_role_assignment_v3":     ReplaceGroupWords("iam", 1),
	"flexibleengine_identity_role_v3":                ReplaceGroupWords("iam", 1),
	"flexibleengine_identity_user_v3":                ReplaceGroupWords("iam", 1),
	"flexibleengine_compute_floatingip_associate_v2": ReplaceGroupWords("ecs", 1),
	"flexibleengine_compute_instance_v2":             ReplaceGroupWords("ecs", 1),
	"flexibleengine_compute_interface_attach_v2":     ReplaceGroupWords("ecs", 1),
	"flexibleengine_compute_keypair_v2":              ReplaceGroupWords("ecs", 1),
	"flexibleengine_compute_servergroup_v2":          ReplaceGroupWords("ecs", 1),
	"flexibleengine_compute_volume_attach_v2":        ReplaceGroupWords("ecs", 1),
	"flexibleengine_vpc_v1":                          ReplaceGroupWords("vpc", 0),
	"flexibleengine_vpc_subnet_v1":                   ReplaceGroupWords("vpc", 0),
	"flexibleengine_networking_network_v2":           ReplaceGroupWords("vpc", 1),
	"flexibleengine_networking_port_v2":              ReplaceGroupWords("vpc", 1),
	"flexibleengine_networking_router_interface_v2":  ReplaceGroupWords("vpc", 1),
	"flexibleengine_networking_router_v2":            ReplaceGroupWords("vpc", 1),
	"flexibleengine_networking_secgroup_rule_v2":     ReplaceGroupWords("vpc", 1),
	"flexibleengine_networking_secgroup_v2":          ReplaceGroupWords("vpc", 1),
	"flexibleengine_networking_subnet_v2":            ReplaceGroupWords("vpc", 0),
	"flexibleengine_networking_vip_associate_v2":     ReplaceGroupWords("vpc", 1),
	"flexibleengine_networking_vip_v2":               ReplaceGroupWords("vpc", 1),
	"flexibleengine_vpc_eip":                         ReplaceGroupWords("eip", 1),
	"flexibleengine_vpc_eip_associate":               ReplaceGroupWords("eip", 1),
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	"flexibleengine_compute_keypair_v2":              "KeyPair",
	"flexibleengine_compute_servergroup_v2":          "ServerGroup",
	"flexibleengine_compute_floatingip_associate_v2": "FloatingIpAssociate",
}
