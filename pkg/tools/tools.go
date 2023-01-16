package tools

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
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

// GenerateExtractor generates the extractor path.
// Extracts the value of `sourceAttrs`
// from `spec.forProvider` allowing nested parameters.
// If `isStatus` is set, then referenced param
// is retrieved from the status, if not, it's extracted
// from the spec.
// An example argument to GenerateExtractor is
// `key`, if `spec.forProvider.key` is to be extracted
// from the referred resource.
// Other examples are `"network", "0", "id"` is converted to
// `spec.forProvider.network[0].id`
func GenerateExtractor(isStatus bool, sourceAttrs ...string) string {

	var sourceAttr string

	if len(sourceAttrs) > 1 {
		listOfSourceAttr := []string{}
		for _, sourceAttr := range sourceAttrs {
			listOfSourceAttr = append(listOfSourceAttr, fmt.Sprintf("\"%s\"", sourceAttr))
		}
		sourceAttr = strings.Join(listOfSourceAttr, ", ")
	} else {
		sourceAttr = fmt.Sprintf("\"%s\"", sourceAttrs[0])
	}
	// Resultat is ExtractorParamPathfunc(true, "network", "0", "id")
	return fmt.Sprintf("%s/pkg/tools.ExtractorParamPathfunc(%t, %s)", ModulePath, isStatus, sourceAttr)
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

// ExtractorParamPathfunc extracts the value of `sourceAttrs`
func ExtractorParamPathfunc(isStatus bool, sourceAttrs ...string) xpref.ExtractValueFn {
	return func(mg xpresource.Managed) string {

		var (
			err error
			r   string
		)

		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}

		sourceAttr := strings.Join(sourceAttrs, ".")

		// if sourceAttr match this regex network.[0-9].id transform to network[0].id
		rg := regexp.MustCompile(`(.*)\.([0-9])\.(.*)`)
		if rg.MatchString(sourceAttr) {
			sourceAttr = rg.ReplaceAllString(sourceAttr, "$1[$2].$3")
		}

		if isStatus {
			r, err = paved.GetString(fmt.Sprintf("status.atProvider.%s", sourceAttr))
			if err != nil {
				return ""
			}
		} else {
			r, err = paved.GetString(fmt.Sprintf("spec.forProvider.%s", sourceAttr))
			if err != nil {
				return ""
			}
		}
		return r

	}
}
