/*
Copyright 2021 Upbound Inc.
*/

package common

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
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
