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
func (in *ColumnsObservation) DeepCopyInto(out *ColumnsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnsObservation.
func (in *ColumnsObservation) DeepCopy() *ColumnsObservation {
	if in == nil {
		return nil
	}
	out := new(ColumnsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColumnsParameters) DeepCopyInto(out *ColumnsParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IsPartition != nil {
		in, out := &in.IsPartition, &out.IsPartition
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnsParameters.
func (in *ColumnsParameters) DeepCopy() *ColumnsParameters {
	if in == nil {
		return nil
	}
	out := new(ColumnsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLIPackage) DeepCopyInto(out *DLIPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLIPackage.
func (in *DLIPackage) DeepCopy() *DLIPackage {
	if in == nil {
		return nil
	}
	out := new(DLIPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DLIPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLIPackageList) DeepCopyInto(out *DLIPackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DLIPackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLIPackageList.
func (in *DLIPackageList) DeepCopy() *DLIPackageList {
	if in == nil {
		return nil
	}
	out := new(DLIPackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DLIPackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLIPackageObservation) DeepCopyInto(out *DLIPackageObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ObjectName != nil {
		in, out := &in.ObjectName, &out.ObjectName
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLIPackageObservation.
func (in *DLIPackageObservation) DeepCopy() *DLIPackageObservation {
	if in == nil {
		return nil
	}
	out := new(DLIPackageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLIPackageParameters) DeepCopyInto(out *DLIPackageParameters) {
	*out = *in
	if in.GroupName != nil {
		in, out := &in.GroupName, &out.GroupName
		*out = new(string)
		**out = **in
	}
	if in.IsAsync != nil {
		in, out := &in.IsAsync, &out.IsAsync
		*out = new(bool)
		**out = **in
	}
	if in.ObjectPath != nil {
		in, out := &in.ObjectPath, &out.ObjectPath
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLIPackageParameters.
func (in *DLIPackageParameters) DeepCopy() *DLIPackageParameters {
	if in == nil {
		return nil
	}
	out := new(DLIPackageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLIPackageSpec) DeepCopyInto(out *DLIPackageSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLIPackageSpec.
func (in *DLIPackageSpec) DeepCopy() *DLIPackageSpec {
	if in == nil {
		return nil
	}
	out := new(DLIPackageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DLIPackageStatus) DeepCopyInto(out *DLIPackageStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DLIPackageStatus.
func (in *DLIPackageStatus) DeepCopy() *DLIPackageStatus {
	if in == nil {
		return nil
	}
	out := new(DLIPackageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Database) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseList) DeepCopyInto(out *DatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Database, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseList.
func (in *DatabaseList) DeepCopy() *DatabaseList {
	if in == nil {
		return nil
	}
	out := new(DatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseObservation) DeepCopyInto(out *DatabaseObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseObservation.
func (in *DatabaseObservation) DeepCopy() *DatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(DatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseParameters) DeepCopyInto(out *DatabaseParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseParameters.
func (in *DatabaseParameters) DeepCopy() *DatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(DatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseStatus) DeepCopyInto(out *DatabaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseStatus.
func (in *DatabaseStatus) DeepCopy() *DatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependentPackagesObservation) DeepCopyInto(out *DependentPackagesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependentPackagesObservation.
func (in *DependentPackagesObservation) DeepCopy() *DependentPackagesObservation {
	if in == nil {
		return nil
	}
	out := new(DependentPackagesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependentPackagesParameters) DeepCopyInto(out *DependentPackagesParameters) {
	*out = *in
	if in.GroupName != nil {
		in, out := &in.GroupName, &out.GroupName
		*out = new(string)
		**out = **in
	}
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]PackagesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependentPackagesParameters.
func (in *DependentPackagesParameters) DeepCopy() *DependentPackagesParameters {
	if in == nil {
		return nil
	}
	out := new(DependentPackagesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinksqlJob) DeepCopyInto(out *FlinksqlJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinksqlJob.
func (in *FlinksqlJob) DeepCopy() *FlinksqlJob {
	if in == nil {
		return nil
	}
	out := new(FlinksqlJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinksqlJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinksqlJobList) DeepCopyInto(out *FlinksqlJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlinksqlJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinksqlJobList.
func (in *FlinksqlJobList) DeepCopy() *FlinksqlJobList {
	if in == nil {
		return nil
	}
	out := new(FlinksqlJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinksqlJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinksqlJobObservation) DeepCopyInto(out *FlinksqlJobObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinksqlJobObservation.
func (in *FlinksqlJobObservation) DeepCopy() *FlinksqlJobObservation {
	if in == nil {
		return nil
	}
	out := new(FlinksqlJobObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinksqlJobParameters) DeepCopyInto(out *FlinksqlJobParameters) {
	*out = *in
	if in.CheckpointEnabled != nil {
		in, out := &in.CheckpointEnabled, &out.CheckpointEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CheckpointInterval != nil {
		in, out := &in.CheckpointInterval, &out.CheckpointInterval
		*out = new(float64)
		**out = **in
	}
	if in.CheckpointMode != nil {
		in, out := &in.CheckpointMode, &out.CheckpointMode
		*out = new(string)
		**out = **in
	}
	if in.CuNumber != nil {
		in, out := &in.CuNumber, &out.CuNumber
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DirtyDataStrategy != nil {
		in, out := &in.DirtyDataStrategy, &out.DirtyDataStrategy
		*out = new(string)
		**out = **in
	}
	if in.EdgeGroupIds != nil {
		in, out := &in.EdgeGroupIds, &out.EdgeGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IdleStateRetention != nil {
		in, out := &in.IdleStateRetention, &out.IdleStateRetention
		*out = new(float64)
		**out = **in
	}
	if in.LogEnabled != nil {
		in, out := &in.LogEnabled, &out.LogEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ManagerCuNumber != nil {
		in, out := &in.ManagerCuNumber, &out.ManagerCuNumber
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObsBucket != nil {
		in, out := &in.ObsBucket, &out.ObsBucket
		*out = new(string)
		**out = **in
	}
	if in.ParallelNumber != nil {
		in, out := &in.ParallelNumber, &out.ParallelNumber
		*out = new(float64)
		**out = **in
	}
	if in.QueueName != nil {
		in, out := &in.QueueName, &out.QueueName
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RestartWhenException != nil {
		in, out := &in.RestartWhenException, &out.RestartWhenException
		*out = new(bool)
		**out = **in
	}
	if in.ResumeCheckpoint != nil {
		in, out := &in.ResumeCheckpoint, &out.ResumeCheckpoint
		*out = new(bool)
		**out = **in
	}
	if in.ResumeMaxNum != nil {
		in, out := &in.ResumeMaxNum, &out.ResumeMaxNum
		*out = new(float64)
		**out = **in
	}
	if in.RunMode != nil {
		in, out := &in.RunMode, &out.RunMode
		*out = new(string)
		**out = **in
	}
	if in.RuntimeConfig != nil {
		in, out := &in.RuntimeConfig, &out.RuntimeConfig
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
	if in.SQL != nil {
		in, out := &in.SQL, &out.SQL
		*out = new(string)
		**out = **in
	}
	if in.SmnTopic != nil {
		in, out := &in.SmnTopic, &out.SmnTopic
		*out = new(string)
		**out = **in
	}
	if in.SmnTopicRef != nil {
		in, out := &in.SmnTopicRef, &out.SmnTopicRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SmnTopicSelector != nil {
		in, out := &in.SmnTopicSelector, &out.SmnTopicSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
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
	if in.TmCus != nil {
		in, out := &in.TmCus, &out.TmCus
		*out = new(float64)
		**out = **in
	}
	if in.TmSlotNum != nil {
		in, out := &in.TmSlotNum, &out.TmSlotNum
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UdfJarURL != nil {
		in, out := &in.UdfJarURL, &out.UdfJarURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinksqlJobParameters.
func (in *FlinksqlJobParameters) DeepCopy() *FlinksqlJobParameters {
	if in == nil {
		return nil
	}
	out := new(FlinksqlJobParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinksqlJobSpec) DeepCopyInto(out *FlinksqlJobSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinksqlJobSpec.
func (in *FlinksqlJobSpec) DeepCopy() *FlinksqlJobSpec {
	if in == nil {
		return nil
	}
	out := new(FlinksqlJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinksqlJobStatus) DeepCopyInto(out *FlinksqlJobStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinksqlJobStatus.
func (in *FlinksqlJobStatus) DeepCopy() *FlinksqlJobStatus {
	if in == nil {
		return nil
	}
	out := new(FlinksqlJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackagesObservation) DeepCopyInto(out *PackagesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackagesObservation.
func (in *PackagesObservation) DeepCopy() *PackagesObservation {
	if in == nil {
		return nil
	}
	out := new(PackagesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackagesParameters) DeepCopyInto(out *PackagesParameters) {
	*out = *in
	if in.PackageName != nil {
		in, out := &in.PackageName, &out.PackageName
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackagesParameters.
func (in *PackagesParameters) DeepCopy() *PackagesParameters {
	if in == nil {
		return nil
	}
	out := new(PackagesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Queue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueList) DeepCopyInto(out *QueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Queue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueList.
func (in *QueueList) DeepCopy() *QueueList {
	if in == nil {
		return nil
	}
	out := new(QueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueObservation) DeepCopyInto(out *QueueObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueObservation.
func (in *QueueObservation) DeepCopy() *QueueObservation {
	if in == nil {
		return nil
	}
	out := new(QueueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueParameters) DeepCopyInto(out *QueueParameters) {
	*out = *in
	if in.CuCount != nil {
		in, out := &in.CuCount, &out.CuCount
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.QueueType != nil {
		in, out := &in.QueueType, &out.QueueType
		*out = new(string)
		**out = **in
	}
	if in.ResourceMode != nil {
		in, out := &in.ResourceMode, &out.ResourceMode
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueParameters.
func (in *QueueParameters) DeepCopy() *QueueParameters {
	if in == nil {
		return nil
	}
	out := new(QueueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpec) DeepCopyInto(out *QueueSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpec.
func (in *QueueSpec) DeepCopy() *QueueSpec {
	if in == nil {
		return nil
	}
	out := new(QueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatus) DeepCopyInto(out *QueueStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatus.
func (in *QueueStatus) DeepCopy() *QueueStatus {
	if in == nil {
		return nil
	}
	out := new(QueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkJob) DeepCopyInto(out *SparkJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkJob.
func (in *SparkJob) DeepCopy() *SparkJob {
	if in == nil {
		return nil
	}
	out := new(SparkJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkJobList) DeepCopyInto(out *SparkJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SparkJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkJobList.
func (in *SparkJobList) DeepCopy() *SparkJobList {
	if in == nil {
		return nil
	}
	out := new(SparkJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkJobObservation) DeepCopyInto(out *SparkJobObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkJobObservation.
func (in *SparkJobObservation) DeepCopy() *SparkJobObservation {
	if in == nil {
		return nil
	}
	out := new(SparkJobObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkJobParameters) DeepCopyInto(out *SparkJobParameters) {
	*out = *in
	if in.AppName != nil {
		in, out := &in.AppName, &out.AppName
		*out = new(string)
		**out = **in
	}
	if in.AppParameters != nil {
		in, out := &in.AppParameters, &out.AppParameters
		*out = new(string)
		**out = **in
	}
	if in.Configurations != nil {
		in, out := &in.Configurations, &out.Configurations
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
	if in.DependentPackages != nil {
		in, out := &in.DependentPackages, &out.DependentPackages
		*out = make([]DependentPackagesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DriverCores != nil {
		in, out := &in.DriverCores, &out.DriverCores
		*out = new(float64)
		**out = **in
	}
	if in.DriverMemory != nil {
		in, out := &in.DriverMemory, &out.DriverMemory
		*out = new(string)
		**out = **in
	}
	if in.ExecutorCores != nil {
		in, out := &in.ExecutorCores, &out.ExecutorCores
		*out = new(float64)
		**out = **in
	}
	if in.ExecutorMemory != nil {
		in, out := &in.ExecutorMemory, &out.ExecutorMemory
		*out = new(string)
		**out = **in
	}
	if in.Executors != nil {
		in, out := &in.Executors, &out.Executors
		*out = new(float64)
		**out = **in
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Jars != nil {
		in, out := &in.Jars, &out.Jars
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MainClass != nil {
		in, out := &in.MainClass, &out.MainClass
		*out = new(string)
		**out = **in
	}
	if in.MaxRetries != nil {
		in, out := &in.MaxRetries, &out.MaxRetries
		*out = new(float64)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PythonFiles != nil {
		in, out := &in.PythonFiles, &out.PythonFiles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.QueueName != nil {
		in, out := &in.QueueName, &out.QueueName
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Specification != nil {
		in, out := &in.Specification, &out.Specification
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkJobParameters.
func (in *SparkJobParameters) DeepCopy() *SparkJobParameters {
	if in == nil {
		return nil
	}
	out := new(SparkJobParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkJobSpec) DeepCopyInto(out *SparkJobSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkJobSpec.
func (in *SparkJobSpec) DeepCopy() *SparkJobSpec {
	if in == nil {
		return nil
	}
	out := new(SparkJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkJobStatus) DeepCopyInto(out *SparkJobStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkJobStatus.
func (in *SparkJobStatus) DeepCopy() *SparkJobStatus {
	if in == nil {
		return nil
	}
	out := new(SparkJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Table) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableList) DeepCopyInto(out *TableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Table, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableList.
func (in *TableList) DeepCopy() *TableList {
	if in == nil {
		return nil
	}
	out := new(TableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableObservation) DeepCopyInto(out *TableObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableObservation.
func (in *TableObservation) DeepCopy() *TableObservation {
	if in == nil {
		return nil
	}
	out := new(TableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableParameters) DeepCopyInto(out *TableParameters) {
	*out = *in
	if in.BucketLocation != nil {
		in, out := &in.BucketLocation, &out.BucketLocation
		*out = new(string)
		**out = **in
	}
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]ColumnsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataFormat != nil {
		in, out := &in.DataFormat, &out.DataFormat
		*out = new(string)
		**out = **in
	}
	if in.DataLocation != nil {
		in, out := &in.DataLocation, &out.DataLocation
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.DatabaseNameRef != nil {
		in, out := &in.DatabaseNameRef, &out.DatabaseNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseNameSelector != nil {
		in, out := &in.DatabaseNameSelector, &out.DatabaseNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DateFormat != nil {
		in, out := &in.DateFormat, &out.DateFormat
		*out = new(string)
		**out = **in
	}
	if in.Delimiter != nil {
		in, out := &in.Delimiter, &out.Delimiter
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EscapeChar != nil {
		in, out := &in.EscapeChar, &out.EscapeChar
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.QuoteChar != nil {
		in, out := &in.QuoteChar, &out.QuoteChar
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TimestampFormat != nil {
		in, out := &in.TimestampFormat, &out.TimestampFormat
		*out = new(string)
		**out = **in
	}
	if in.WithColumnHeader != nil {
		in, out := &in.WithColumnHeader, &out.WithColumnHeader
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableParameters.
func (in *TableParameters) DeepCopy() *TableParameters {
	if in == nil {
		return nil
	}
	out := new(TableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpec) DeepCopyInto(out *TableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpec.
func (in *TableSpec) DeepCopy() *TableSpec {
	if in == nil {
		return nil
	}
	out := new(TableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableStatus) DeepCopyInto(out *TableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableStatus.
func (in *TableStatus) DeepCopy() *TableStatus {
	if in == nil {
		return nil
	}
	out := new(TableStatus)
	in.DeepCopyInto(out)
	return out
}