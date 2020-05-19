package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedHatterSpec defines the desired state of RedHatter
type RedHatterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	EmployeeName string `json:"employeename,omitempty"`
	EmployeeStatus string `json:"employeestatus,omitempty"`
	IsCovidThere bool `json:"iscovidthere,omitempty"`
}

// RedHatterStatus defines the observed state of RedHatter
type RedHatterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedHatter is the Schema for the redhatters API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=redhatters,scope=Namespaced
type RedHatter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedHatterSpec   `json:"spec,omitempty"`
	Status RedHatterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedHatterList contains a list of RedHatter
type RedHatterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedHatter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedHatter{}, &RedHatterList{})
}
