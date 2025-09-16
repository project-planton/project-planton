package pulumilabels

const (
	// StackFqdnLabelKey is the primary label that takes precedence over individual components
	// Format: "organization/project/stack"
	StackFqdnLabelKey = "pulumi.project-planton.org/stack.fqdn"

	// OrganizationLabelKey is used when stack.fqdn is not present
	OrganizationLabelKey = "pulumi.project-planton.org/organization"

	// ProjectLabelKey is used when stack.fqdn is not present
	ProjectLabelKey = "pulumi.project-planton.org/project"

	// StackNameLabelKey is used when stack.fqdn is not present
	StackNameLabelKey = "pulumi.project-planton.org/stack.name"
)
