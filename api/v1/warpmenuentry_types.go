package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ConditionReady indicates that the warp menu entry has been processed successfully.
	ConditionReady = "Ready"
	// ReasonEntryRendered indicates that the warp menu entry has been rendered successfully.
	ReasonEntryRendered = "EntryRendered"
)

// DisplayName defines a language-dependent display name for the WarpMenuEntry.
// Currently, we have mandatory localized versions for German and English, but
// more (optional) languages might be added in the future.
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
	DisplayName DisplayName `json:"displayName"`

	// Category defines the key of the category under which to place the entry in the Warp Menu. Should preferably be
	// one of the pre-defined categories, but any missing category is created on the fly using the key as its name.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=50
	Category string `json:"category"`

	// Path is the URL path to which the WarpMenuEntry should point.
	// Note that this must not include a domain or protocol.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=^/.*$
	Path string `json:"path"`

	// Disabled flag suppresses the display of the WarpMenuEntry in the Warp Menu if true
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
// +kubebuilder:printcolumn:name="Category",type="string",JSONPath=".spec.category",description="Category for the Warp Menu Entry"
// +kubebuilder:printcolumn:name="Disabled",type="string",JSONPath=".spec.disabled",description="Is the Warp Menu Entry disabled"
// +kubebuilder:printcolumn:name="DisplayName English",type="string",JSONPath=".spec.displayName.en",description="Display name in English for the Warp Menu Entry",priority=1
// +kubebuilder:printcolumn:name="DisplayName German",type="string",JSONPath=".spec.displayName.de",description="Display name in German for the Warp Menu Entry",priority=1

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
