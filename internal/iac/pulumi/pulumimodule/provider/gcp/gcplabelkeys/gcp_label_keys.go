package gcplabelkeys

import (
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/labels/labelkeys"
)

var (
	Resource     = labelkeys.WithNormalizedDomainPrefix("resource")
	Organization = labelkeys.WithNormalizedDomainPrefix("organization")
	Environment  = labelkeys.WithNormalizedDomainPrefix("environment")
	ResourceKind = labelkeys.WithNormalizedDomainPrefix("resource-kind")
	ResourceId   = labelkeys.WithNormalizedDomainPrefix("resource-id")
)
