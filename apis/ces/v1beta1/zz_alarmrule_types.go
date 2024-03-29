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

type AlarmActionsObservation struct {
}

type AlarmActionsParameters struct {

	// specifies the list of objects to be notified
	// if the alarm status changes, the maximum length is 5.
	// if type is set to notification, the value of notification_list cannot be empty.
	// if type is set to autoscaling, the value of notification_list must be []
	// and the value of namespace must be SYS.AS.
	// Note: to enable the autoscaling alarm rules take effect, you must bind scaling
	// policies.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/smn/v1beta1.Topic
	// +kubebuilder:validation:Optional
	NotificationList []*string `json:"notificationList,omitempty" tf:"notification_list,omitempty"`

	// References to Topic in smn to populate notificationList.
	// +kubebuilder:validation:Optional
	NotificationListRefs []v1.Reference `json:"notificationListRefs,omitempty" tf:"-"`

	// Selector for a list of Topic in smn to populate notificationList.
	// +kubebuilder:validation:Optional
	NotificationListSelector *v1.Selector `json:"notificationListSelector,omitempty" tf:"-"`

	// specifies the type of action triggered by an alarm. the
	// value can be notification or autoscaling.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type AlarmRuleObservation struct {

	// Indicates the alarm status. The value can be:
	AlarmState *string `json:"alarmState,omitempty" tf:"alarm_state,omitempty"`

	// Specifies the alarm triggering condition. The structure
	// is described below.
	// +kubebuilder:validation:Required
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Indicates the alarm rule ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the time when the alarm status changed.
	// The value is a UNIX timestamp and the unit is ms.
	UpdateTime *float64 `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type AlarmRuleParameters struct {

	// Specifies whether to enable the action
	// to be triggered by an alarm. The default value is true.
	// +kubebuilder:validation:Optional
	AlarmActionEnabled *bool `json:"alarmActionEnabled,omitempty" tf:"alarm_action_enabled,omitempty"`

	// Specifies the action triggered by an alarm. The
	// structure is described below.
	// +kubebuilder:validation:Optional
	AlarmActions []AlarmActionsParameters `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// The value can be a string of 0 to 256 characters.
	// +kubebuilder:validation:Optional
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// Specifies whether to enable the alarm. The default
	// value is true.
	// +kubebuilder:validation:Optional
	AlarmEnabled *bool `json:"alarmEnabled,omitempty" tf:"alarm_enabled,omitempty"`

	// Specifies the alarm severity. The value can be
	// 1, 2, 3 or 4, which indicates critical, major, minor, and informational, respectively.
	// The default value is 2. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the name of an alarm rule. The value can
	// be a string of 1 to 128 characters that can consist of numbers, lowercase letters,
	// uppercase letters, underscores (_), or hyphens (-).
	// +kubebuilder:validation:Required
	AlarmName *string `json:"alarmName" tf:"alarm_name,omitempty"`

	// Specifies the alarm triggering condition. The structure
	// is described below.
	// +kubebuilder:validation:Required
	Condition []ConditionParameters `json:"condition" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	InsufficientdataActions []InsufficientdataActionsParameters `json:"insufficientdataActions,omitempty" tf:"insufficientdata_actions,omitempty"`

	// Specifies the alarm metrics. The structure is described
	// below.
	// +kubebuilder:validation:Required
	Metric []MetricParameters `json:"metric" tf:"metric,omitempty"`

	// Specifies the action triggered by the clearing of
	// an alarm. The structure is described below.
	// +kubebuilder:validation:Optional
	OkActions []OkActionsParameters `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ConditionObservation struct {
	SuppressDuration *float64 `json:"suppressDuration,omitempty" tf:"suppress_duration,omitempty"`
}

type ConditionParameters struct {

	// Specifies the comparison condition of alarm
	// thresholds. The value can be >, =, <, >=, or <=.
	// +kubebuilder:validation:Required
	ComparisonOperator *string `json:"comparisonOperator" tf:"comparison_operator,omitempty"`

	// Specifies the number of consecutive occurrence times.
	// The value ranges from 1 to 5.
	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// Specifies the data rollup methods. The value can be
	// max, min, average, sum, and vaiance.
	// +kubebuilder:validation:Required
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// Specifies the alarm checking period in seconds. The
	// value can be 1, 300, 1200, 3600, 14400, and 86400.
	// Note: If period is set to 1, the raw metric data is used to determine
	// whether to generate an alarm.
	// +kubebuilder:validation:Required
	Period *float64 `json:"period" tf:"period,omitempty"`

	// Specifies the data unit.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Specifies the alarm threshold. The value ranges from
	// 0 to Number of 1.7976931348623157e+308.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type DimensionsObservation struct {
}

type DimensionsParameters struct {

	// Specifies the dimension name. The value can be a string
	// of 1 to 32 characters that must start with a letter and can consists of uppercase
	// letters, lowercase letters, numbers, underscores (_), or hyphens (-).
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the dimension value. The value can be a string
	// of 1 to 64 characters that must start with a letter or a number and can consists
	// of uppercase letters, lowercase letters, numbers, underscores (_), or hyphens
	// (-).
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type InsufficientdataActionsObservation struct {
}

type InsufficientdataActionsParameters struct {

	// specifies the list of objects to be notified
	// if the alarm status changes, the maximum length is 5.
	// +kubebuilder:validation:Required
	NotificationList []*string `json:"notificationList" tf:"notification_list,omitempty"`

	// specifies the type of action triggered by an alarm. the
	// value is notification.
	// notification: indicates that a notification will be sent to the user.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type MetricObservation struct {
}

type MetricParameters struct {

	// Specifies the list of metric dimensions. Currently,
	// the maximum length of the dimesion list that are supported is 3. The structure
	// is described below.
	// +kubebuilder:validation:Required
	Dimensions []DimensionsParameters `json:"dimensions" tf:"dimensions,omitempty"`

	// Specifies the metric name. The value can be a string
	// of 1 to 64 characters that must start with a letter and can consists of uppercase
	// letters, lowercase letters, numbers, or underscores (_).
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// Specifies the namespace in service.item format.
	// service.item can be a string of 3 to 32 characters that must start with a letter and
	// can consists of uppercase letters, lowercase letters, numbers, or underscores (_).
	// For details, see Services Interconnected with Cloud Eye.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type OkActionsObservation struct {
}

type OkActionsParameters struct {

	// specifies the list of objects to be notified
	// if the alarm status changes, the maximum length is 5.
	// +kubebuilder:validation:Required
	NotificationList []*string `json:"notificationList" tf:"notification_list,omitempty"`

	// specifies the type of action triggered by an alarm. the
	// value is notification.
	// notification: indicates that a notification will be sent to the user.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// AlarmRuleSpec defines the desired state of AlarmRule
type AlarmRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlarmRuleParameters `json:"forProvider"`
}

// AlarmRuleStatus defines the observed state of AlarmRule.
type AlarmRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlarmRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlarmRule is the Schema for the AlarmRules API. ""page_title: "flexibleengine_ces_alarmrule"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type AlarmRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlarmRuleSpec   `json:"spec"`
	Status            AlarmRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlarmRuleList contains a list of AlarmRules
type AlarmRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlarmRule `json:"items"`
}

// Repository type metadata.
var (
	AlarmRule_Kind             = "AlarmRule"
	AlarmRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlarmRule_Kind}.String()
	AlarmRule_KindAPIVersion   = AlarmRule_Kind + "." + CRDGroupVersion.String()
	AlarmRule_GroupVersionKind = CRDGroupVersion.WithKind(AlarmRule_Kind)
)

func init() {
	SchemeBuilder.Register(&AlarmRule{}, &AlarmRuleList{})
}
