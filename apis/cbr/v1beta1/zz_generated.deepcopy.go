//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupCycleObservation) DeepCopyInto(out *BackupCycleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupCycleObservation.
func (in *BackupCycleObservation) DeepCopy() *BackupCycleObservation {
	if in == nil {
		return nil
	}
	out := new(BackupCycleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupCycleParameters) DeepCopyInto(out *BackupCycleParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(string)
		**out = **in
	}
	if in.ExecutionTimes != nil {
		in, out := &in.ExecutionTimes, &out.ExecutionTimes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupCycleParameters.
func (in *BackupCycleParameters) DeepCopy() *BackupCycleParameters {
	if in == nil {
		return nil
	}
	out := new(BackupCycleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LongTermRetentionObservation) DeepCopyInto(out *LongTermRetentionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LongTermRetentionObservation.
func (in *LongTermRetentionObservation) DeepCopy() *LongTermRetentionObservation {
	if in == nil {
		return nil
	}
	out := new(LongTermRetentionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LongTermRetentionParameters) DeepCopyInto(out *LongTermRetentionParameters) {
	*out = *in
	if in.Daily != nil {
		in, out := &in.Daily, &out.Daily
		*out = new(float64)
		**out = **in
	}
	if in.Monthly != nil {
		in, out := &in.Monthly, &out.Monthly
		*out = new(float64)
		**out = **in
	}
	if in.Weekly != nil {
		in, out := &in.Weekly, &out.Weekly
		*out = new(float64)
		**out = **in
	}
	if in.Yearly != nil {
		in, out := &in.Yearly, &out.Yearly
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LongTermRetentionParameters.
func (in *LongTermRetentionParameters) DeepCopy() *LongTermRetentionParameters {
	if in == nil {
		return nil
	}
	out := new(LongTermRetentionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyObservation) DeepCopyInto(out *PolicyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyObservation.
func (in *PolicyObservation) DeepCopy() *PolicyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyParameters) DeepCopyInto(out *PolicyParameters) {
	*out = *in
	if in.BackupCycle != nil {
		in, out := &in.BackupCycle, &out.BackupCycle
		*out = make([]BackupCycleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackupQuantity != nil {
		in, out := &in.BackupQuantity, &out.BackupQuantity
		*out = new(float64)
		**out = **in
	}
	if in.DestinationProjectID != nil {
		in, out := &in.DestinationProjectID, &out.DestinationProjectID
		*out = new(string)
		**out = **in
	}
	if in.DestinationRegion != nil {
		in, out := &in.DestinationRegion, &out.DestinationRegion
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LongTermRetention != nil {
		in, out := &in.LongTermRetention, &out.LongTermRetention
		*out = make([]LongTermRetentionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TimePeriod != nil {
		in, out := &in.TimePeriod, &out.TimePeriod
		*out = new(float64)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyParameters.
func (in *PolicyParameters) DeepCopy() *PolicyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesObservation) DeepCopyInto(out *ResourcesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesObservation.
func (in *ResourcesObservation) DeepCopy() *ResourcesObservation {
	if in == nil {
		return nil
	}
	out := new(ResourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesParameters) DeepCopyInto(out *ResourcesParameters) {
	*out = *in
	if in.Excludes != nil {
		in, out := &in.Excludes, &out.Excludes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesParameters.
func (in *ResourcesParameters) DeepCopy() *ResourcesParameters {
	if in == nil {
		return nil
	}
	out := new(ResourcesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vault) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultList) DeepCopyInto(out *VaultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vault, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultList.
func (in *VaultList) DeepCopy() *VaultList {
	if in == nil {
		return nil
	}
	out := new(VaultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultObservation) DeepCopyInto(out *VaultObservation) {
	*out = *in
	if in.Allocated != nil {
		in, out := &in.Allocated, &out.Allocated
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SpecCode != nil {
		in, out := &in.SpecCode, &out.SpecCode
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultObservation.
func (in *VaultObservation) DeepCopy() *VaultObservation {
	if in == nil {
		return nil
	}
	out := new(VaultObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultParameters) DeepCopyInto(out *VaultParameters) {
	*out = *in
	if in.AutoBind != nil {
		in, out := &in.AutoBind, &out.AutoBind
		*out = new(bool)
		**out = **in
	}
	if in.AutoExpand != nil {
		in, out := &in.AutoExpand, &out.AutoExpand
		*out = new(bool)
		**out = **in
	}
	if in.AutoPay != nil {
		in, out := &in.AutoPay, &out.AutoPay
		*out = new(string)
		**out = **in
	}
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(string)
		**out = **in
	}
	if in.BindRules != nil {
		in, out := &in.BindRules, &out.BindRules
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ChargingMode != nil {
		in, out := &in.ChargingMode, &out.ChargingMode
		*out = new(string)
		**out = **in
	}
	if in.ConsistentLevel != nil {
		in, out := &in.ConsistentLevel, &out.ConsistentLevel
		*out = new(string)
		**out = **in
	}
	if in.EnterpriseProjectID != nil {
		in, out := &in.EnterpriseProjectID, &out.EnterpriseProjectID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.PeriodUnit != nil {
		in, out := &in.PeriodUnit, &out.PeriodUnit
		*out = new(string)
		**out = **in
	}
	if in.PolicyID != nil {
		in, out := &in.PolicyID, &out.PolicyID
		*out = new(string)
		**out = **in
	}
	if in.ProtectionType != nil {
		in, out := &in.ProtectionType, &out.ProtectionType
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourcesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultParameters.
func (in *VaultParameters) DeepCopy() *VaultParameters {
	if in == nil {
		return nil
	}
	out := new(VaultParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSpec) DeepCopyInto(out *VaultSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSpec.
func (in *VaultSpec) DeepCopy() *VaultSpec {
	if in == nil {
		return nil
	}
	out := new(VaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultStatus) DeepCopyInto(out *VaultStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultStatus.
func (in *VaultStatus) DeepCopy() *VaultStatus {
	if in == nil {
		return nil
	}
	out := new(VaultStatus)
	in.DeepCopyInto(out)
	return out
}
