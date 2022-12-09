package config

import (
	"github.com/upbound/upjet/pkg/config"
)

// VersionV1Beta1 is used to signify that the resource has been tested and external name configured
const VersionV1Beta1 = "v1beta1"

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
				r.References["vpc_id"] = config.Reference{
					Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.VPC",
				}
				// router_id is a reference to a Router resource
			case "router_id":
				r.References["router_id"] = config.Reference{
					Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Router",
				}
				// subnet_id is a reference to a Subnet resource
			case "subnet_id":
				if _, ok := r.References["network_id"]; ok {
					r.References["subnet_id"] = config.Reference{
						Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.NetworkingSubnet",
					}
				} else {
					r.References["subnet_id"] = config.Reference{
						Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet",
					}
				}
				// network_id is a reference to a Network resource
			case "network_id":
				r.References["network_id"] = config.Reference{
					Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Network",
				}
				// security_group_id is a reference to a SecurityGroup resource
			case "security_group_id":
				r.References["security_group_id"] = config.Reference{
					Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.SecGroup",
				}
				// port_id is a reference to a NetworkPort resource
			case "port_id":
				r.References["port_id"] = config.Reference{
					Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Port",
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
