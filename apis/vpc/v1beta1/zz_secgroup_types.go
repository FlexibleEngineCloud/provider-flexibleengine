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

type SecGroupObservation struct {

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecGroupParameters struct {

	// Whether or not to delete the default
	// egress security rules. This is false by default. See the below note
	// for more information.
	// +kubebuilder:validation:Optional
	DeleteDefaultRules *bool `json:"deleteDefaultRules,omitempty" tf:"delete_default_rules,omitempty"`

	// A unique name for the security group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique name for the security group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// security group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The resource ID in UUID format.
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/iam/v1beta1.Project
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Reference to a Project in iam to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDRef *v1.Reference `json:"tenantIdRef,omitempty" tf:"-"`

	// Selector for a Project in iam to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDSelector *v1.Selector `json:"tenantIdSelector,omitempty" tf:"-"`
}

// SecGroupSpec defines the desired state of SecGroup
type SecGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecGroupParameters `json:"forProvider"`
}

// SecGroupStatus defines the observed state of SecGroup.
type SecGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecGroup is the Schema for the SecGroups API. ""page_title: "flexibleengine_networking_secgroup_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type SecGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecGroupSpec   `json:"spec"`
	Status            SecGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecGroupList contains a list of SecGroups
type SecGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecGroup `json:"items"`
}

// Repository type metadata.
var (
	SecGroup_Kind             = "SecGroup"
	SecGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecGroup_Kind}.String()
	SecGroup_KindAPIVersion   = SecGroup_Kind + "." + CRDGroupVersion.String()
	SecGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecGroup{}, &SecGroupList{})
}
