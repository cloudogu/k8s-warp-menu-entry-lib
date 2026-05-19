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
type DisplayName struct {
	// de is the display name for the WarpMenuEntry in German.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=50
	DE string `json:"de"`

	// en is the display name for the WarpMenuEntry in English.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=50
	EN string `json:"en"`
}

// WarpMenuEntrySpec defines the desired state of WarpMenuEntry
type WarpMenuEntrySpec struct {

	// DisplayName defines the name to display for this WarpMenuEntry.
	// +required
	DisplayName *DisplayName `json:"displayName"`

	// Category defines the categorie-Key in Warp-Menü for this WarpMenuEntry.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=50
	Category string `json:"category"`

	// Path is the URL path under which the application should be reachable for this WarpMenuEntry.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=^/.*$
	Path string `json:"path"`

	// Disabled flag suppresses the display of the WarpMenuEntry in Warp-Menü if true
	// +optional
	Disabled bool `json:"disabled,omitempty"`
}

// WarpMenuEntryStatus defines the observed state of WarpMenuEntry.
type WarpMenuEntryStatus struct {

	// Conditions represent the current state of the WarpMenuEntry resource.
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
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
