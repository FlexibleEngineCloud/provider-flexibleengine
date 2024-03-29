/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/eip/v1beta1"
	v1beta11 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1"
	tools "github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PublicIP); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PublicIP[i3].EIPID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PublicIP[i3].EIPIDRef,
			Selector:     mg.Spec.ForProvider.PublicIP[i3].EIPIDSelector,
			To: reference.To{
				List:    &v1beta1.EIPList{},
				Managed: &v1beta1.EIP{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PublicIP[i3].EIPID")
		}
		mg.Spec.ForProvider.PublicIP[i3].EIPID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PublicIP[i3].EIPIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      tools.ExtractorParamPathfunc(true, "id"),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCSubnetList{},
			Managed: &v1beta11.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

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
