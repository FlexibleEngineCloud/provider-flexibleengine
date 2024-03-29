/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta11 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/iam/v1beta1"
	v1beta1 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/kms/v1beta1"
	tools "github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this OBSBucket.
func (mg *OBSBucket) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OBSBucketObject.
func (mg *OBSBucketObject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &OBSBucketList{},
			Managed: &OBSBucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OBSBucketReplication.
func (mg *OBSBucketReplication) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Agency),
		Extract:      tools.ExtractorParamPathfunc(false, "name"),
		Reference:    mg.Spec.ForProvider.AgencyRef,
		Selector:     mg.Spec.ForProvider.AgencySelector,
		To: reference.To{
			List:    &v1beta11.AgencyList{},
			Managed: &v1beta11.Agency{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Agency")
	}
	mg.Spec.ForProvider.Agency = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AgencyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &OBSBucketList{},
			Managed: &OBSBucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationBucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DestinationBucketRef,
		Selector:     mg.Spec.ForProvider.DestinationBucketSelector,
		To: reference.To{
			List:    &OBSBucketList{},
			Managed: &OBSBucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DestinationBucket")
	}
	mg.Spec.ForProvider.DestinationBucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DestinationBucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this S3BucketObject.
func (mg *S3BucketObject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &S3BucketList{},
			Managed: &S3Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this S3BucketPolicy.
func (mg *S3BucketPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &S3BucketList{},
			Managed: &S3Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}
