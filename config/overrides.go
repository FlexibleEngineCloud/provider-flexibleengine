package config

import (
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
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
				// router_id is a reference to a Router resource
			case "router_id":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("vpc", "Router"),
				}
				// subnet_id is a reference to a Subnet resource
			case "subnet_id":
				if _, ok := r.References["network_id"]; ok {
					r.References[k] = config.Reference{
						Type: tools.GenerateType("vpc", "NetworkingSubnet"),
					}
				} else {
					r.References[k] = config.Reference{
						Type: tools.GenerateType("vpc", "VPCSubnet"),
					}
				}
				// network_id is a reference to a Network resource
			case "network_id":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("vpc", "Network"),
				}
				// security_group_id is a reference to a SecurityGroup resource
			case "security_group_id":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("vpc", "SecGroup"),
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
			}
		}
	}
}

func defaultVersion() config.ResourceOption {
	return func(r *config.Resource) {
		r.Version = tools.Version
	}
}
