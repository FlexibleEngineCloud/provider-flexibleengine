// Package dns contains custom ResourceConfigurators for the dns package.
package dns

import (
	"context"
	"fmt"

	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("flexibleengine_dns_ptrrecord_v2", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider

		r.References["region_floatingip"] = config.Reference{
			TerraformName: "flexibleengine_vpc_eip",
			Extractor:     common.PathRegionExtractor,
		}

		r.ExternalName.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {

			ErrFmtNoAttribute := "no %s attribute found in parameters"

			// floating_ip_id is the ID of the floating IP address
			floatingIPID, ok := parameters["floating_ip_id"]
			if !ok {
				return "", errors.Errorf(ErrFmtNoAttribute, "floating_ip_id")
			}

			regionFloatingIP, ok := providerConfig["region_floatingip"]
			if !ok {
				return "", errors.Errorf(ErrFmtNoAttribute, "region_floatingip")
			}

			return fmt.Sprintf("%s:%s", regionFloatingIP, floatingIPID), nil
		}

	})

}
