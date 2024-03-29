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

type RecordsetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RecordsetParameters struct {

	// A description of the record set. Max length is 255 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the record set. Note the . at the end of the name.
	// Changing this creates a new DNS record set.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// An array of DNS records.
	// +kubebuilder:validation:Required
	Records []*string `json:"records" tf:"records,omitempty"`

	// The region in which to create the DNS record set.
	// If omitted, the region argument of the provider is used.
	// Changing this creates a new DNS record set.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The time to live (TTL) of the record set (in seconds). The value
	// range is 300–2147483647. The default value is 300.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The key/value pairs to associate with the record set.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of record set. The options include A, AAAA, MX,
	// CNAME, TXT, NS, SRV, CAA, and PTR.
	// Changing this creates a new DNS record set.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Map of additional options.
	// Changing this creates a new DNS record set.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the zone in which to create the record set.
	// Changing this creates a new DNS record set.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// RecordsetSpec defines the desired state of Recordset
type RecordsetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordsetParameters `json:"forProvider"`
}

// RecordsetStatus defines the observed state of Recordset.
type RecordsetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordsetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Recordset is the Schema for the Recordsets API. ""page_title: "flexibleengine_dns_recordset_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Recordset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordsetSpec   `json:"spec"`
	Status            RecordsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordsetList contains a list of Recordsets
type RecordsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Recordset `json:"items"`
}

// Repository type metadata.
var (
	Recordset_Kind             = "Recordset"
	Recordset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Recordset_Kind}.String()
	Recordset_KindAPIVersion   = Recordset_Kind + "." + CRDGroupVersion.String()
	Recordset_GroupVersionKind = CRDGroupVersion.WithKind(Recordset_Kind)
)

func init() {
	SchemeBuilder.Register(&Recordset{}, &RecordsetList{})
}
