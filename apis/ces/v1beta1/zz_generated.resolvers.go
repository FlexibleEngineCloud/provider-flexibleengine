/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/smn/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AlarmRule.
func (mg *AlarmRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AlarmActions); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.AlarmActions[i3].NotificationList),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.AlarmActions[i3].NotificationListRefs,
			Selector:      mg.Spec.ForProvider.AlarmActions[i3].NotificationListSelector,
			To: reference.To{
				List:    &v1beta1.TopicList{},
				Managed: &v1beta1.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AlarmActions[i3].NotificationList")
		}
		mg.Spec.ForProvider.AlarmActions[i3].NotificationList = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.AlarmActions[i3].NotificationListRefs = mrsp.ResolvedReferences

	}

	return nil
}