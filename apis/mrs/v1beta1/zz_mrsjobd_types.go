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

type MrsJobDObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	JobState *string `json:"jobState,omitempty" tf:"job_state,omitempty"`
}

type MrsJobDParameters struct {

	// Key parameter for program execution. The parameter
	// is specified by the function of the user's program. MRS is only responsible
	// for loading the parameter. The parameter contains a maximum of 2047 characters,
	// excluding special characters such as ;|&>'<$, and can be empty.
	// +kubebuilder:validation:Optional
	Arguments *string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// Cluster ID
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// SQL program path This parameter is needed
	// by Spark Script and Hive Script jobs only and must meet the following requirements:
	// Contains a maximum of 1023 characters, excluding special characters such as
	// ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://.
	// Ends with .sql. sql is case-insensitive.
	// +kubebuilder:validation:Optional
	HiveScriptPath *string `json:"hiveScriptPath,omitempty" tf:"hive_script_path,omitempty"`

	// Path for inputting data, which must start with / or s3a://.
	// A correct OBS path is required. The parameter contains a maximum of 1023 characters,
	// excluding special characters such as ;|&>'<$, and can be empty.
	// +kubebuilder:validation:Optional
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// Whether a job is protected true false The current
	// version does not support this function.
	// +kubebuilder:validation:Optional
	IsProtected *bool `json:"isProtected,omitempty" tf:"is_protected,omitempty"`

	// Whether a job is public true false The current version
	// does not support this function.
	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// Path of the .jar package or .sql file for program
	// execution The parameter must meet the following requirements: Contains a maximum
	// of 1023 characters, excluding special characters such as ;|&><'$. The address
	// cannot be empty or full of spaces. Starts with / or s3a://. Spark Script must
	// end with .sql; while MapReduce and Spark Jar must end with .jar. sql and jar
	// are case-insensitive.
	// +kubebuilder:validation:Required
	JarPath *string `json:"jarPath" tf:"jar_path,omitempty"`

	// Path for storing job logs that record job running status.
	// This path must start with / or s3a://. A correct OBS path is required. The parameter
	// contains a maximum of 1023 characters, excluding special characters such as
	// ;|&>'<$, and can be empty.
	// +kubebuilder:validation:Optional
	JobLog *string `json:"jobLog,omitempty" tf:"job_log,omitempty"`

	// Job name Contains only 1 to 64 letters, digits, hyphens
	// (-), and underscores (_). NOTE: Identical job names are allowed but not recommended.
	// +kubebuilder:validation:Required
	JobName *string `json:"jobName" tf:"job_name,omitempty"`

	// Job type 1: MapReduce 2: Spark 3: Hive Script 4: HiveQL
	// (not supported currently) 5: DistCp, importing and exporting data.  6: Spark
	// Script 7: Spark SQL, submitting Spark SQL statements. (not supported in this
	// APIcurrently) NOTE: Spark and Hive jobs can be added to only clusters including
	// Spark and Hive components.
	// +kubebuilder:validation:Required
	JobType *float64 `json:"jobType" tf:"job_type,omitempty"`

	// Path for outputting data, which must start with / or
	// s3a://. A correct OBS path is required. If the path does not exist, the system
	// automatically creates it. The parameter contains a maximum of 1023 characters,
	// excluding special characters such as ;|&>'<$, and can be empty.
	// +kubebuilder:validation:Optional
	Output *string `json:"output,omitempty" tf:"output,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// MrsJobDSpec defines the desired state of MrsJobD
type MrsJobDSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MrsJobDParameters `json:"forProvider"`
}

// MrsJobDStatus defines the observed state of MrsJobD.
type MrsJobDStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MrsJobDObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MrsJobD is the Schema for the MrsJobDs API. ""page_title: "flexibleengine_mrs_job_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type MrsJobD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MrsJobDSpec   `json:"spec"`
	Status            MrsJobDStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MrsJobDList contains a list of MrsJobDs
type MrsJobDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MrsJobD `json:"items"`
}

// Repository type metadata.
var (
	MrsJobD_Kind             = "MrsJobD"
	MrsJobD_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MrsJobD_Kind}.String()
	MrsJobD_KindAPIVersion   = MrsJobD_Kind + "." + CRDGroupVersion.String()
	MrsJobD_GroupVersionKind = CRDGroupVersion.WithKind(MrsJobD_Kind)
)

func init() {
	SchemeBuilder.Register(&MrsJobD{}, &MrsJobDList{})
}
