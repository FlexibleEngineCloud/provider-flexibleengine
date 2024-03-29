/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OBSBucketReplicationObservation struct {

	// The name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A configuration of object cross-region replication management. The object supports the following:
	// +kubebuilder:validation:Optional
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type OBSBucketReplicationParameters struct {

	// Specifies the IAM agency applied to the cross-region replication function.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/iam/v1beta1.Agency
	// +crossplane:generate:reference:extractor=github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools.ExtractorParamPathfunc(false, "name")
	// +kubebuilder:validation:Optional
	Agency *string `json:"agency,omitempty" tf:"agency,omitempty"`

	// Reference to a Agency in iam to populate agency.
	// +kubebuilder:validation:Optional
	AgencyRef *v1.Reference `json:"agencyRef,omitempty" tf:"-"`

	// Selector for a Agency in iam to populate agency.
	// +kubebuilder:validation:Optional
	AgencySelector *v1.Selector `json:"agencySelector,omitempty" tf:"-"`

	// Specifies the name of the source bucket. Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=OBSBucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a OBSBucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a OBSBucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Specifies the name of the destination bucket.
	// +crossplane:generate:reference:type=OBSBucket
	// +kubebuilder:validation:Optional
	DestinationBucket *string `json:"destinationBucket,omitempty" tf:"destination_bucket,omitempty"`

	// Reference to a OBSBucket to populate destinationBucket.
	// +kubebuilder:validation:Optional
	DestinationBucketRef *v1.Reference `json:"destinationBucketRef,omitempty" tf:"-"`

	// Selector for a OBSBucket to populate destinationBucket.
	// +kubebuilder:validation:Optional
	DestinationBucketSelector *v1.Selector `json:"destinationBucketSelector,omitempty" tf:"-"`

	// A configuration of object cross-region replication management. The object supports the following:
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RuleObservation struct {

	// The name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleParameters struct {

	// Specifies cross-region replication rule status. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the object key prefix identifying one or more objects to which the rule applies and
	// duplicated prefixes are not supported. If omitted, all objects in the bucket will be managed by the lifecycle rule.
	// To copy a folder, end the prefix with a slash (/), for example, imgs/.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies the storage class for replicated objects. Valid values are "STANDARD",
	// "WARM" (Infrequent Access) and "COLD" (Archive).
	// If omitted, the storage class of object copies is the same as that of objects in the source bucket.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

// OBSBucketReplicationSpec defines the desired state of OBSBucketReplication
type OBSBucketReplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OBSBucketReplicationParameters `json:"forProvider"`
}

// OBSBucketReplicationStatus defines the observed state of OBSBucketReplication.
type OBSBucketReplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OBSBucketReplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OBSBucketReplication is the Schema for the OBSBucketReplications API. ""page_title: "flexibleengine_obs_bucket_replication"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type OBSBucketReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OBSBucketReplicationSpec   `json:"spec"`
	Status            OBSBucketReplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OBSBucketReplicationList contains a list of OBSBucketReplications
type OBSBucketReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OBSBucketReplication `json:"items"`
}

// Repository type metadata.
var (
	OBSBucketReplication_Kind             = "OBSBucketReplication"
	OBSBucketReplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OBSBucketReplication_Kind}.String()
	OBSBucketReplication_KindAPIVersion   = OBSBucketReplication_Kind + "." + CRDGroupVersion.String()
	OBSBucketReplication_GroupVersionKind = CRDGroupVersion.WithKind(OBSBucketReplication_Kind)
)

func init() {
	SchemeBuilder.Register(&OBSBucketReplication{}, &OBSBucketReplicationList{})
}
