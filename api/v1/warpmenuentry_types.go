/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ConditionReady indicates that the warp menu entry has been rendered successfully.
	ConditionReady = "Ready"
	// ReasonEntryRendered indicates that the warp menu entry has been rendered successfully.
	ReasonEntryRendered = "EntryRendered"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DisplayName defines a display name for the WarpMenuEntry.

// +kubebuilder:validation:MinProperties=1
type DisplayName struct {
	// de is the display name for the WarpMenuEntry in German.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	DE string `json:"de,omitzero"`

	// en is the display name for the WarpMenuEntry in English.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	EN string `json:"en,omitzero"`
}

// WarpMenuEntrySpec defines the desired state of WarpMenuEntry
type WarpMenuEntrySpec struct {

	// DisplayName defines the name to display for this WarpMenuEntry.
	// +required
	DisplayName *DisplayName `json:"displayName"`

	// DisplayName defines the categorie-Key in Warp-Menü for this WarpMenuEntry.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Category string `json:"category"`

	// Path is the URL path under which the application should be reachable for this WarpMenuEntry.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=^/.*$
	Path string `json:"path"`

	// Disabled flag suppresses the display of the WarpMenuEntry in Warp-Menü if true
	// +optional
	Disabled bool `json:"disabled"`
}

// WarpMenuEntryStatus defines the observed state of WarpMenuEntry.
type WarpMenuEntryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// For Kubernetes API conventions, see:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	// conditions represent the current state of the WarpMenuEntry resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=warp
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type == 'Ready')].status",description="Warp menu entry has been rendered successfully."
// +kubebuilder:printcolumn:name="Path",type="string",JSONPath=".spec.path",description="The URL path under which the application should be reachable for the Warp Menu Entry"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the resource"

// WarpMenuEntry is the Schema for the warpmenuentries API
type WarpMenuEntry struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of WarpMenuEntry
	// +required
	Spec WarpMenuEntrySpec `json:"spec"`

	// status defines the observed state of WarpMenuEntry
	// +optional
	Status WarpMenuEntryStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// WarpMenuEntryList contains a list of WarpMenuEntry
type WarpMenuEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []WarpMenuEntry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WarpMenuEntry{}, &WarpMenuEntryList{})
}
