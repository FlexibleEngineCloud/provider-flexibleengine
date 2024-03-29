/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/iam/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this OrganizationUsers.
func (mg *OrganizationUsers) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Organization),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Organization")
	}
	mg.Spec.ForProvider.Organization = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Users); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Users[i3].UserID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Users[i3].UserIDRef,
			Selector:     mg.Spec.ForProvider.Users[i3].UserIDSelector,
			To: reference.To{
				List:    &v1beta1.UserList{},
				Managed: &v1beta1.User{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Users[i3].UserID")
		}
		mg.Spec.ForProvider.Users[i3].UserID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Users[i3].UserIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Repository.
func (mg *Repository) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Organization),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Organization")
	}
	mg.Spec.ForProvider.Organization = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RepositorySharing.
func (mg *RepositorySharing) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Organization),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Organization")
	}
	mg.Spec.ForProvider.Organization = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}
