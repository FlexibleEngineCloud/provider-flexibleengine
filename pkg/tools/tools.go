package tools

import (
	"fmt"
	"regexp"
	"strings"
)

// VersionV1Beta1 is used to signify that the resource has been tested and external name configured
const (
	versionV1Beta1 = "v1beta1"

	// Version is the version of the API
	Version = versionV1Beta1

	// ResourcePrefix is the prefix of the resource
	ResourcePrefix = "flexibleengine"

	// ModulePath is the path of the module
	ModulePath = "github.com/FrangipaneTeam/provider-flexibleengine"
)

// GenerateType generates the type name for a given module and type.
// For example, GenerateType("vpc", "VPC") will return
// "github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC"
func GenerateType(module, _type string) string {
	return fmt.Sprintf("%s/apis/%s/%s.%s", ModulePath, module, Version, _type)
}

// RemoveVersion removes the version from the resource name
// For example, RemoveVersion("flexibleengine_vpc_subnet_v1") will return
// "flexibleengine_vpc_subnet"
func RemoveVersion(name string) []string {
	// Replace string with regex

	rg := regexp.MustCompile(`_v[0-9]+`)
	y := rg.ReplaceAllString(name, "")

	return strings.Split(RemovePrefix(y), "_")

}

// RemovePrefix removes the prefix from the resource name
// For example, RemovePrefix("flexibleengine_vpc_subnet") will return
// "vpc_subnet"
func RemovePrefix(name string) string {
	return strings.TrimPrefix(name, fmt.Sprintf("%s_", ResourcePrefix))
}

// RemoveGroup removes the group from the resource name
// For example, RemoveGroup("vpc_subnet") will return
// "subnet"
func RemoveGroup(name []string) string {
	return strings.Join(name[1:], "_")
}
