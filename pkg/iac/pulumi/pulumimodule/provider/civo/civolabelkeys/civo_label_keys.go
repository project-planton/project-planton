package civolabelkeys

import (
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/labels/labelkeys"
)

var (
	Resource     = labelkeys.WithNormalizedDomainPrefix("resource")
	Organization = labelkeys.WithNormalizedDomainPrefix("organization")
	Environment  = labelkeys.WithNormalizedDomainPrefix("environment")
	ResourceKind = labelkeys.WithNormalizedDomainPrefix("kind")
	ResourceId   = labelkeys.WithNormalizedDomainPrefix("id")
	ResourceName = labelkeys.WithNormalizedDomainPrefix("name")
)
