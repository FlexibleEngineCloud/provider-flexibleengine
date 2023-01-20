package config

import (
	"strings"

	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/references"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() config.ResourceOption { //nolint:gocyclo
	return func(r *config.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			switch k {

			case "vpc_id":
				r.References[k] = references.TypeVPCID().Get()

			case "subnet_id":
				r.References[k] = references.TypeVPCSubnetID().Get()

			case "subnet_ids", "subnets":
				r.References[k] = references.TypeVPCSubnetID().WithCustomRefsSelectors("SubnetIDRefs", "SubnetIDSelector").Get()

			case "network_id":
				r.References[k] = references.TypeVPCSubnetID().WithCustomRefsSelectors("NetworkIDRef", "NetworkIDSelector").Get()

			case "security_group_id", "security_group":
				r.References[k] = references.TypeSecurityGroupID().Get()

			case "security_group_ids":
				r.References[k] = references.TypeSecurityGroupID().WithCustomRefsSelectors("SecurityGroupIDRefs", "SecurityGroupIDSelector").Get()

			case "ipv4_eip_id", "eip_id":
				r.References[k] = references.TypeEIPID().Get()

			case "port_id":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("vpc", "Port"),
				}

			case "tenant_id":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("iam", "Project"),
				}

			case "instance_id":
				switch r.ShortGroup {
				case "dms":
					if strings.Contains(r.MetaResource.Name, "kafka") {
						r.References[k] = config.Reference{
							Type: tools.GenerateType("dms", "KafkaInstance"),
						}
					} else {
						r.References[k] = config.Reference{
							Type: tools.GenerateType("dms", "RocketMQInstance"),
						}
					}
				default:
					r.References[k] = config.Reference{
						Type: tools.GenerateType("ecs", "Instance"),
					}
				}

			case "image_name":
				r.References[k] = config.Reference{
					Type:      tools.GenerateType("ims", "Image"),
					Extractor: tools.GenerateExtractor(false, "image_name"),
				}

			case "image_id":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("ims", "Image"),
				}

			case "pool_id":
				switch r.ShortGroup {
				case "elb":
					r.References[k] = config.Reference{
						Type: tools.GenerateType("elb", "Pool"),
					}
				case "dedicatedelb":
					r.References[k] = config.Reference{
						Type: tools.GenerateType("dedicatedelb", "Pool"),
					}
				}

			case "policy_id":
				if r.ShortGroup == "waf" {
					r.References[k] = config.Reference{
						Type: tools.GenerateType("waf", "Policy"),
					}
				}
			}
		}
	}
}

func defaultVersion() config.ResourceOption {
	return func(r *config.Resource) {
		r.Version = tools.Version
	}
}
