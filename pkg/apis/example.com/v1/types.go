package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Database describes a database.
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DatabaseSpec `json:"spec"`
}

// DatabaseSpec is the spec for a Foo resource
type DatabaseSpec struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Encoding string `json:"encoding,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatabaseList is a list of Database resources
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Database `json:"items"`
}
