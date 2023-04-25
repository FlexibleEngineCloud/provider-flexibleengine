package references

import (
	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// TypeVPCSubnet returns a Reference to a VPCSubnet resource.
// Type is vpc/VPCSubnet, Extractor(true, "id"), RefFieldName: SubnetIDRef, SelectorFieldName: SubnetIDSelector
func TypeVPCSubnetID() Reference {
	return Reference{
		Type:              tools.GenerateType("vpc", "VPCSubnet"),
		Extractor:         tools.GenerateExtractor(true, "id"),
		RefFieldName:      "SubnetIDRef",
		SelectorFieldName: "SubnetIDSelector",
	}
}

// TypeVPC returns a Reference to a VPC resource.
// Type is vpc/VPC
func TypeVPCID() Reference {
	return Reference{
		Type: tools.GenerateType("vpc", "VPC"),
	}
}

// TypeSecurityGroup returns a Reference to a SecurityGroup resource.
// Type is vpc/SecurityGroup, RefFieldName: SecurityGroupIDRef, SelectorFieldName: SecurityGroupIDSelector
func TypeSecurityGroupID() Reference {
	return Reference{
		Type:              tools.GenerateType("vpc", "SecurityGroup"),
		RefFieldName:      "SecurityGroupIDRef",
		SelectorFieldName: "SecurityGroupIDSelector",
	}
}

// TypeVPCSubnetIDIPV4 returns a Reference to a VPCSubnet resource with extractor for ipv4_subnet_id.
// Type is vpc/VPCSubnet, Extractor(true, "ipv4_subnet_id")
func TypeVPCSubnetIDIPV4() Reference {
	return Reference{
		Type:      tools.GenerateType("vpc", "VPCSubnet"),
		Extractor: tools.GenerateExtractor(true, "ipv4_subnet_id"),
	}
}

// TypeEIPID returns a Reference to a EIP resource.
// Type is vpc/EIP
func TypeEIPID() Reference {
	return Reference{
		Type: tools.GenerateType("eip", "EIP"),
	}
}
