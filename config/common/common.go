package common

import (
	"fmt"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/upbound/upjet/pkg/resource"
)

var (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = fmt.Sprintf("%s/config/common", tools.ModulePath)

	// PathARNExtractor is the golang path to ARNExtractor function
	// in this package.
	PathARNExtractor = SelfPackagePath + ".ARNExtractor()"

	// PathTerraformIDExtractor is the golang path to TerraformID extractor
	// function in this package.
	PathTerraformIDExtractor = SelfPackagePath + ".TerraformID()"

	// PathRegionExtractor is the golang path to RegionExtractor function
	PathRegionExtractor = SelfPackagePath + ".RegionExtractor()"

	// PathImageNameExtractor is the golang path to ImageNameExtractor function
	PathImageNameExtractor = SelfPackagePath + ".ImageNameExtractor()"

	// PathDBNameExtractor is the golang path to DBNameExtractor function
	PathDBNameExtractor = SelfPackagePath + ".DBNameExtractor()"

	// PathNameExtractor is the golang path to NameExtractor function
	PathNameExtractor = SelfPackagePath + ".NameExtractor()"

	// PathBucketNameExtractor is the golang path to BucketNameExtractor function
	PathBucketNameExtractor = SelfPackagePath + ".BucketNameExtractor()"
)

// RegionExtractor extracts region from "spec.forProvider.region" which
func RegionExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		r, err := paved.GetString("spec.forProvider.region")
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		return r
	}
}

// TerraformID returns the Terraform ID string of the resource without any
// manipulation.
func TerraformID() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		return tr.GetID()
	}
}

// ImageNameExtractor extracts image name from "spec.forProvider.image_name"
func ImageNameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		r, err := paved.GetString("spec.forProvider.image_name")
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		return r
	}
}

// DBNameExtractor extracts database name from "spec.forProvider.db_name"
func DBNameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		r, err := paved.GetString("spec.forProvider.db_name")
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		return r
	}
}

// NameExtractor extracts name from "spec.forProvider.name"
func NameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		r, err := paved.GetString("spec.forProvider.name")
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		return r
	}
}

// BucketNameExtractor extracts bucket name from "spec.forProvider.bucket"
func BucketNameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		r, err := paved.GetString("spec.forProvider.bucket")
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		return r
	}
}
