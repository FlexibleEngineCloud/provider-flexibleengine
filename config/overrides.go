package config

import (
	"strings"

	"github.com/upbound/upjet/pkg/config"

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
			// vpc_id is a reference to a Vpc resource
			case "vpc_id":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("vpc", "VPC"),
				}
			// subnet_id and network_id is a reference to a Subnet resource
			case "subnet_id", "network_id":
				r.References[k] = config.Reference{
					Type:      tools.GenerateType("vpc", "VPCSubnet"),
					Extractor: tools.GenerateExtractor(true, "id"),
				}
				// security_group_id is a reference to a SecurityGroup resource
			case "security_group_id":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("vpc", "SecurityGroup"),
				}
				// port_id is a reference to a NetworkPort resource
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
					TerraformName: "flexibleengine_images_image_v2",
					Extractor:     tools.GenerateExtractor(false, "image_name"),
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
