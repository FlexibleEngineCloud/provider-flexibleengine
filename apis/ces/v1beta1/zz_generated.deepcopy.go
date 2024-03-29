//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmActionsObservation) DeepCopyInto(out *AlarmActionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmActionsObservation.
func (in *AlarmActionsObservation) DeepCopy() *AlarmActionsObservation {
	if in == nil {
		return nil
	}
	out := new(AlarmActionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmActionsParameters) DeepCopyInto(out *AlarmActionsParameters) {
	*out = *in
	if in.NotificationList != nil {
		in, out := &in.NotificationList, &out.NotificationList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NotificationListRefs != nil {
		in, out := &in.NotificationListRefs, &out.NotificationListRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NotificationListSelector != nil {
		in, out := &in.NotificationListSelector, &out.NotificationListSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmActionsParameters.
func (in *AlarmActionsParameters) DeepCopy() *AlarmActionsParameters {
	if in == nil {
		return nil
	}
	out := new(AlarmActionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmRule) DeepCopyInto(out *AlarmRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmRule.
func (in *AlarmRule) DeepCopy() *AlarmRule {
	if in == nil {
		return nil
	}
	out := new(AlarmRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlarmRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmRuleList) DeepCopyInto(out *AlarmRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlarmRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmRuleList.
func (in *AlarmRuleList) DeepCopy() *AlarmRuleList {
	if in == nil {
		return nil
	}
	out := new(AlarmRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlarmRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmRuleObservation) DeepCopyInto(out *AlarmRuleObservation) {
	*out = *in
	if in.AlarmState != nil {
		in, out := &in.AlarmState, &out.AlarmState
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]ConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmRuleObservation.
func (in *AlarmRuleObservation) DeepCopy() *AlarmRuleObservation {
	if in == nil {
		return nil
	}
	out := new(AlarmRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmRuleParameters) DeepCopyInto(out *AlarmRuleParameters) {
	*out = *in
	if in.AlarmActionEnabled != nil {
		in, out := &in.AlarmActionEnabled, &out.AlarmActionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]AlarmActionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.AlarmEnabled != nil {
		in, out := &in.AlarmEnabled, &out.AlarmEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AlarmLevel != nil {
		in, out := &in.AlarmLevel, &out.AlarmLevel
		*out = new(float64)
		**out = **in
	}
	if in.AlarmName != nil {
		in, out := &in.AlarmName, &out.AlarmName
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]ConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InsufficientdataActions != nil {
		in, out := &in.InsufficientdataActions, &out.InsufficientdataActions
		*out = make([]InsufficientdataActionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = make([]MetricParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OkActions != nil {
		in, out := &in.OkActions, &out.OkActions
		*out = make([]OkActionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmRuleParameters.
func (in *AlarmRuleParameters) DeepCopy() *AlarmRuleParameters {
	if in == nil {
		return nil
	}
	out := new(AlarmRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmRuleSpec) DeepCopyInto(out *AlarmRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmRuleSpec.
func (in *AlarmRuleSpec) DeepCopy() *AlarmRuleSpec {
	if in == nil {
		return nil
	}
	out := new(AlarmRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmRuleStatus) DeepCopyInto(out *AlarmRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmRuleStatus.
func (in *AlarmRuleStatus) DeepCopy() *AlarmRuleStatus {
	if in == nil {
		return nil
	}
	out := new(AlarmRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionObservation) DeepCopyInto(out *ConditionObservation) {
	*out = *in
	if in.SuppressDuration != nil {
		in, out := &in.SuppressDuration, &out.SuppressDuration
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionObservation.
func (in *ConditionObservation) DeepCopy() *ConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionParameters) DeepCopyInto(out *ConditionParameters) {
	*out = *in
	if in.ComparisonOperator != nil {
		in, out := &in.ComparisonOperator, &out.ComparisonOperator
		*out = new(string)
		**out = **in
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(float64)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionParameters.
func (in *ConditionParameters) DeepCopy() *ConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsObservation) DeepCopyInto(out *DimensionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsObservation.
func (in *DimensionsObservation) DeepCopy() *DimensionsObservation {
	if in == nil {
		return nil
	}
	out := new(DimensionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsParameters) DeepCopyInto(out *DimensionsParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsParameters.
func (in *DimensionsParameters) DeepCopy() *DimensionsParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsufficientdataActionsObservation) DeepCopyInto(out *InsufficientdataActionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsufficientdataActionsObservation.
func (in *InsufficientdataActionsObservation) DeepCopy() *InsufficientdataActionsObservation {
	if in == nil {
		return nil
	}
	out := new(InsufficientdataActionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsufficientdataActionsParameters) DeepCopyInto(out *InsufficientdataActionsParameters) {
	*out = *in
	if in.NotificationList != nil {
		in, out := &in.NotificationList, &out.NotificationList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsufficientdataActionsParameters.
func (in *InsufficientdataActionsParameters) DeepCopy() *InsufficientdataActionsParameters {
	if in == nil {
		return nil
	}
	out := new(InsufficientdataActionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricObservation) DeepCopyInto(out *MetricObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricObservation.
func (in *MetricObservation) DeepCopy() *MetricObservation {
	if in == nil {
		return nil
	}
	out := new(MetricObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricParameters) DeepCopyInto(out *MetricParameters) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]DimensionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricParameters.
func (in *MetricParameters) DeepCopy() *MetricParameters {
	if in == nil {
		return nil
	}
	out := new(MetricParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OkActionsObservation) DeepCopyInto(out *OkActionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OkActionsObservation.
func (in *OkActionsObservation) DeepCopy() *OkActionsObservation {
	if in == nil {
		return nil
	}
	out := new(OkActionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OkActionsParameters) DeepCopyInto(out *OkActionsParameters) {
	*out = *in
	if in.NotificationList != nil {
		in, out := &in.NotificationList, &out.NotificationList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OkActionsParameters.
func (in *OkActionsParameters) DeepCopy() *OkActionsParameters {
	if in == nil {
		return nil
	}
	out := new(OkActionsParameters)
	in.DeepCopyInto(out)
	return out
}
