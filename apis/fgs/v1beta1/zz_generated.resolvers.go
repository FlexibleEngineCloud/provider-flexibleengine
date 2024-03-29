/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/iam/v1beta1"
	v1beta11 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1"
	tools "github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Function.
func (mg *Function) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Agency),
		Extract:      tools.ExtractorParamPathfunc(false, "name"),
		Reference:    mg.Spec.ForProvider.AgencyRef,
		Selector:     mg.Spec.ForProvider.AgencySelector,
		To: reference.To{
			List:    &v1beta1.AgencyList{},
			Managed: &v1beta1.Agency{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Agency")
	}
	mg.Spec.ForProvider.Agency = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AgencyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      tools.ExtractorParamPathfunc(true, "id"),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCSubnetList{},
			Managed: &v1beta11.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCList{},
			Managed: &v1beta11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Trigger.
func (mg *Trigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionUrn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FunctionUrnRef,
		Selector:     mg.Spec.ForProvider.FunctionUrnSelector,
		To: reference.To{
			List:    &FunctionList{},
			Managed: &Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionUrn")
	}
	mg.Spec.ForProvider.FunctionUrn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionUrnRef = rsp.ResolvedReference

	return nil
}
