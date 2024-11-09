package kuberneteslabelkeys

import (
	"github.com/project-planton/project-planton/pkg/pulmod/labels/labelkeys"
)

var (
	Resource     = labelkeys.WithDomainPrefix("resource")
	Organization = labelkeys.WithDomainPrefix("organization")
	Environment  = labelkeys.WithDomainPrefix("environment")
	ResourceKind = labelkeys.WithDomainPrefix("resource-kind")
	ResourceId   = labelkeys.WithDomainPrefix("resource-id")
)
