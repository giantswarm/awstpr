package awsclusterspec

import (
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/api/v1"
)

type CustomObject struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty"`
	Spec                 Spec `json:"spec"`
}
