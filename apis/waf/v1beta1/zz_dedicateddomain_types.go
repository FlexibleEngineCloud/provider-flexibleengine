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

type DedicatedDomainObservation struct {

	// Whether a domain name is connected to WAF. Valid values are:
	AccessStatus *float64 `json:"accessStatus,omitempty" tf:"access_status,omitempty"`

	// The alarm page of domain. Valid values are:
	AlarmPage map[string]*string `json:"alarmPage,omitempty" tf:"alarm_page,omitempty"`

	// The name of the certificate used by the domain name.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// The compliance certifications of the domain, values are:
	ComplianceCertification map[string]*bool `json:"complianceCertification,omitempty" tf:"compliance_certification,omitempty"`

	// ID of the domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The protocol type of the client. The options are HTTP and HTTPS.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The traffic identifier of domain. Valid values are:
	TrafficIdentifier map[string]*string `json:"trafficIdentifier,omitempty" tf:"traffic_identifier,omitempty"`
}

type DedicatedDomainParameters struct {

	// Specifies the certificate ID. This parameter is mandatory when client_protocol
	// is set to HTTPS.
	// +crossplane:generate:reference:type=DedicatedCertificate
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Reference to a DedicatedCertificate to populate certificateId.
	// +kubebuilder:validation:Optional
	CertificateIDRef *v1.Reference `json:"certificateIdRef,omitempty" tf:"-"`

	// Selector for a DedicatedCertificate to populate certificateId.
	// +kubebuilder:validation:Optional
	CertificateIDSelector *v1.Selector `json:"certificateIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Cipher *string `json:"cipher,omitempty" tf:"cipher,omitempty"`

	// Specifies the domain name to be protected. For example, www.example.com or
	// *.example.com. Changing this creates a new domain.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	KeepPolicy *bool `json:"keepPolicy,omitempty" tf:"keep_policy,omitempty"`

	// The status of domain PCI 3DS, true: enabled, false: disabled.
	// +kubebuilder:validation:Optional
	Pci3Ds *bool `json:"pci3Ds,omitempty" tf:"pci_3ds,omitempty"`

	// The status of domain PCI DSS, true: enabled, false: disabled.
	// +kubebuilder:validation:Optional
	PciDss *bool `json:"pciDss,omitempty" tf:"pci_dss,omitempty"`

	// Specifies the policy ID associated with the domain. If not specified,
	// a new policy will be created automatically. Changing this creates a new domain.
	// +crossplane:generate:reference:type=DedicatedPolicy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a DedicatedPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a DedicatedPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`

	// The protection status of domain, 0: suspended, 1: enabled.
	// Default value is 1.
	// +kubebuilder:validation:Optional
	ProtectStatus *float64 `json:"protectStatus,omitempty" tf:"protect_status,omitempty"`

	// Specifies whether a proxy is configured. Default value is false.
	// +kubebuilder:validation:Optional
	Proxy *bool `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The object structure is documented below.
	// +kubebuilder:validation:Required
	Server []ServerParameters `json:"server" tf:"server,omitempty"`

	// The TLS configuration of domain.
	// +kubebuilder:validation:Optional
	TLS *string `json:"tls,omitempty" tf:"tls,omitempty"`
}

type ServerObservation struct {
}

type ServerParameters struct {

	// IP address or domain name of the web server that the client accesses. For
	// example, 192.168.1.1 or www.example.com. Changing this creates a new service.
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// Protocol type of the client. The options include HTTP and HTTPS.
	// Changing this creates a new service.
	// +kubebuilder:validation:Required
	ClientProtocol *string `json:"clientProtocol" tf:"client_protocol,omitempty"`

	// Port number used by the web server. The value ranges from 0 to 65535. Changing this
	// creates a new service.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Protocol used by WAF to forward client requests to the server. The
	// options include HTTP and HTTPS. Changing this creates a new service.
	// +kubebuilder:validation:Required
	ServerProtocol *string `json:"serverProtocol" tf:"server_protocol,omitempty"`

	// Server network type, IPv4 or IPv6. Valid values are: ipv4 and ipv6. Changing
	// this creates a new service.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The id of the vpc used by the server. Changing this creates a service.
	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

// DedicatedDomainSpec defines the desired state of DedicatedDomain
type DedicatedDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedDomainParameters `json:"forProvider"`
}

// DedicatedDomainStatus defines the observed state of DedicatedDomain.
type DedicatedDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedDomain is the Schema for the DedicatedDomains API. ""page_title: "flexibleengine_waf_dedicated_domain"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type DedicatedDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedDomainSpec   `json:"spec"`
	Status            DedicatedDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedDomainList contains a list of DedicatedDomains
type DedicatedDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedDomain `json:"items"`
}

// Repository type metadata.
var (
	DedicatedDomain_Kind             = "DedicatedDomain"
	DedicatedDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedDomain_Kind}.String()
	DedicatedDomain_KindAPIVersion   = DedicatedDomain_Kind + "." + CRDGroupVersion.String()
	DedicatedDomain_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedDomain{}, &DedicatedDomainList{})
}
