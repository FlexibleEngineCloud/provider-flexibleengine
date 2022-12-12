/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta13 "github.com/FrangipaneTeam/provider-flexibleengine/apis/ces/v1beta1"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1"
	v1beta11 "github.com/FrangipaneTeam/provider-flexibleengine/apis/elb/v1beta1"
	v1beta12 "github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Configuration.
func (mg *Configuration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.InstanceConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceConfig[i3].InstanceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceConfig[i3].InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceConfig[i3].InstanceIDSelector,
			To: reference.To{
				List:    &v1beta1.InstanceList{},
				Managed: &v1beta1.Instance{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.InstanceConfig[i3].InstanceID")
		}
		mg.Spec.ForProvider.InstanceConfig[i3].InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.InstanceConfig[i3].InstanceIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LbaasListeners); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LbaasListeners[i3].PoolID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LbaasListeners[i3].PoolIDRef,
			Selector:     mg.Spec.ForProvider.LbaasListeners[i3].PoolIDSelector,
			To: reference.To{
				List:    &v1beta11.PoolList{},
				Managed: &v1beta11.Pool{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LbaasListeners[i3].PoolID")
		}
		mg.Spec.ForProvider.LbaasListeners[i3].PoolID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LbaasListeners[i3].PoolIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Networks); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Networks[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Networks[i3].IDRef,
			Selector:     mg.Spec.ForProvider.Networks[i3].IDSelector,
			To: reference.To{
				List:    &v1beta12.NetworkList{},
				Managed: &v1beta12.Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Networks[i3].ID")
		}
		mg.Spec.ForProvider.Networks[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Networks[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalingConfigurationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScalingConfigurationIDRef,
		Selector:     mg.Spec.ForProvider.ScalingConfigurationIDSelector,
		To: reference.To{
			List:    &ConfigurationList{},
			Managed: &Configuration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalingConfigurationID")
	}
	mg.Spec.ForProvider.ScalingConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalingConfigurationIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SecurityGroups); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroups[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SecurityGroups[i3].IDRef,
			Selector:     mg.Spec.ForProvider.SecurityGroups[i3].IDSelector,
			To: reference.To{
				List:    &v1beta12.SecGroupList{},
				Managed: &v1beta12.SecGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups[i3].ID")
		}
		mg.Spec.ForProvider.SecurityGroups[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SecurityGroups[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta12.VPCList{},
			Managed: &v1beta12.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AlarmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AlarmIDRef,
		Selector:     mg.Spec.ForProvider.AlarmIDSelector,
		To: reference.To{
			List:    &v1beta13.AlarmRuleList{},
			Managed: &v1beta13.AlarmRule{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AlarmID")
	}
	mg.Spec.ForProvider.AlarmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AlarmIDRef = rsp.ResolvedReference

	return nil
}