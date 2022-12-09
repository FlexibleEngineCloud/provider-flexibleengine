package tools

import (
	"fmt"
)

// VersionV1Beta1 is used to signify that the resource has been tested and external name configured
const (
	VersionV1Beta1 = "v1beta1"
	Version        = VersionV1Beta1
	ResourcePrefix = "flexibleengine"                               //golint:ignore
	ModulePath     = "github.com/gaetanars/provider-flexibleengine" //golint:ignore
)

// GenerateType generates the type name for a given module and type.
// For example, GenerateType("vpc", "VPC") will return
// "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.VPC"
func GenerateType(module, _type string) string {
	return fmt.Sprintf("%s/apis/%s/%s.%s", ModulePath, module, Version, _type)
}
