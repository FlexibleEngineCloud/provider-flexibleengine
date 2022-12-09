package config

import (
	"fmt"

	"github.com/upbound/upjet/pkg/config"
)

// VersionV1Beta1 is used to signify that the resource has been tested and external name configured
const (
	VersionV1Beta1 = "v1beta1"
	Version        = VersionV1Beta1
)

// Generate Type
func GenerateType(Module, Type string) string {
	return fmt.Sprintf("%s/apis/%s.%s", ModulePath, Version, Type)
}

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
					Type: GenerateType("vpc", "VPC"),
				}
				// router_id is a reference to a Router resource
			case "router_id":
				r.References[k] = config.Reference{
					Type: GenerateType("vpc", "Router"),
				}
				// subnet_id is a reference to a Subnet resource
			case "subnet_id":
				if _, ok := r.References["network_id"]; ok {
					r.References[k] = config.Reference{
						Type: GenerateType("vpc", "NetworkingSubnet"),
					}
				} else {
					r.References[k] = config.Reference{
						Type: GenerateType("vpc", "VPCSubnet"),
					}
				}
				// network_id is a reference to a Network resource
			case "network_id":
				r.References[k] = config.Reference{
					Type: GenerateType("vpc", "Network"),
				}
				// security_group_id is a reference to a SecurityGroup resource
			case "security_group_id":
				r.References[k] = config.Reference{
					Type: GenerateType("vpc", "SecGroup"),
				}
				// port_id is a reference to a NetworkPort resource
			case "port_id":
				r.References[k] = config.Reference{
					Type: GenerateType("vpc", "NetworkPort"),
				}
			}
		}
	}
}

func defaultVersion() config.ResourceOption {
	return func(r *config.Resource) {
		r.Version = VersionV1Beta1
	}
}
