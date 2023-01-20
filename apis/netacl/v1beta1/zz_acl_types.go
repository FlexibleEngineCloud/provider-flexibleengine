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

type ACLObservation struct {

	// The ID of the network ACL.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the ingress firewall policy for the network ACL.
	InboundPolicyID *string `json:"inboundPolicyId,omitempty" tf:"inbound_policy_id,omitempty"`

	// The ID of the egress firewall policy for the network ACL.
	OutboundPolicyID *string `json:"outboundPolicyId,omitempty" tf:"outbound_policy_id,omitempty"`

	// A list of the port IDs of the subnet gateway.
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The status of the network ACL.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ACLParameters struct {

	// Specifies the supplementary information about the network ACL.
	// This parameter can contain a maximum of 255 characters and cannot contain angle brackets (< or >).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// References to ACLRule in netacl to populate inboundRules.
	// +kubebuilder:validation:Optional
	InboundRuleRefs []v1.Reference `json:"inboundRuleRefs,omitempty" tf:"-"`

	// Selector for a list of ACLRule in netacl to populate inboundRules.
	// +kubebuilder:validation:Optional
	InboundRuleSelector *v1.Selector `json:"inboundRuleSelector,omitempty" tf:"-"`

	// A list of the IDs of ingress rules associated with the network ACL.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/netacl/v1beta1.ACLRule
	// +crossplane:generate:reference:refFieldName=InboundRuleRefs
	// +crossplane:generate:reference:selectorFieldName=InboundRuleSelector
	// +kubebuilder:validation:Optional
	InboundRules []*string `json:"inboundRules,omitempty" tf:"inbound_rules,omitempty"`

	// Specifies the network ACL name. This parameter can contain a maximum of 64 characters,
	// which may consist of letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// References to ACLRule in netacl to populate outboundRules.
	// +kubebuilder:validation:Optional
	OutboundRuleRefs []v1.Reference `json:"outboundRuleRefs,omitempty" tf:"-"`

	// Selector for a list of ACLRule in netacl to populate outboundRules.
	// +kubebuilder:validation:Optional
	OutboundRuleSelector *v1.Selector `json:"outboundRuleSelector,omitempty" tf:"-"`

	// A list of the IDs of egress rules associated with the network ACL.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/netacl/v1beta1.ACLRule
	// +crossplane:generate:reference:refFieldName=OutboundRuleRefs
	// +crossplane:generate:reference:selectorFieldName=OutboundRuleSelector
	// +kubebuilder:validation:Optional
	OutboundRules []*string `json:"outboundRules,omitempty" tf:"outbound_rules,omitempty"`

	// References to VPCSubnet in vpc to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of VPCSubnet in vpc to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of the IDs of networks associated with the network ACL.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools.ExtractorParamPathfunc(true, "id")
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`
}

// ACLSpec defines the desired state of ACL
type ACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ACLParameters `json:"forProvider"`
}

// ACLStatus defines the observed state of ACL.
type ACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ACL is the Schema for the ACLs API. ""page_title: "flexibleengine_network_acl"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type ACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ACLSpec   `json:"spec"`
	Status            ACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ACLList contains a list of ACLs
type ACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ACL `json:"items"`
}

// Repository type metadata.
var (
	ACL_Kind             = "ACL"
	ACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ACL_Kind}.String()
	ACL_KindAPIVersion   = ACL_Kind + "." + CRDGroupVersion.String()
	ACL_GroupVersionKind = CRDGroupVersion.WithKind(ACL_Kind)
)

func init() {
	SchemeBuilder.Register(&ACL{}, &ACLList{})
}
