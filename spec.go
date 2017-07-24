package awstpr

import (
	"github.com/giantswarm/awstpr/spec"
	"github.com/giantswarm/clustertpr"
)

type Spec struct {
	Cluster clustertpr.Cluster `json:"cluster" yaml:"cluster"`
	AWS     spec.AWS           `json:"aws" yaml:"aws"`
}
