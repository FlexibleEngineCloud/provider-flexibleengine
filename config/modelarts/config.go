package modelsarts

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// TODO flexibleengine_modelarts_dataset Require MultiType

	// flexibleengine_modelarts_dataset_version
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/modelarts_dataset_version
	p.AddResourceConfigurator("flexibleengine_modelarts_dataset_version", func(r *config.Resource) {

		// dataset_id
		r.References["dataset_id"] = config.Reference{
			Type: "DataSet",
		}

		r.UseAsync = true

	})

}
