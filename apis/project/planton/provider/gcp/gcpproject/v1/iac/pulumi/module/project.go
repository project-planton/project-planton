package module

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	gcpprojectv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpproject/v1"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func project(ctx *pulumi.Context, locals *Locals) (*organizations.Project, error) {
	// Create a random 3-char suffix
	createdRand, err := random.NewRandomString(ctx,
		fmt.Sprintf("%s-suffix", locals.GcpProject.Metadata.Name),
		&random.RandomStringArgs{
			Length:  pulumi.Int(3),
			Special: pulumi.Bool(false),
			Numeric: pulumi.Bool(false),
			Upper:   pulumi.Bool(false),
			Lower:   pulumi.Bool(true),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate random suffix for projectId")
	}

	// Construct the final projectId from metadata.name + "-" + 3-char suffix
	// Apply transformations to meet GCP's length and character constraints.
	projectId := pulumi.All(createdRand.Result).ApplyT(func(args []interface{}) (string, error) {
		suffix := args[0].(string)
		safeName := makeProjectIdSafe(locals.GcpProject.Metadata.Name)
		finalId := fmt.Sprintf("%s-%s", safeName, suffix)
		if len(finalId) > 30 {
			finalId = finalId[:30]
		}
		// Remove trailing hyphens if any (GCP disallows ending with '-')
		finalId = strings.TrimRight(finalId, "-")
		// Ensure at least 6 chars remain
		for len(finalId) < 6 {
			finalId = finalId + "x"
		}
		return finalId, nil
	}).(pulumi.StringOutput)

	projectArgs := &organizations.ProjectArgs{
		Name:              pulumi.String(locals.GcpProject.Metadata.Name),
		ProjectId:         projectId,
		BillingAccount:    pulumi.String(locals.GcpProject.Spec.BillingAccountId),
		Labels:            pulumi.ToStringMap(locals.GcpLabels),
		AutoCreateNetwork: pulumi.Bool(!locals.GcpProject.Spec.GetDisableDefaultNetwork()),
	}

	if locals.GcpProject.Spec.ParentType == gcpprojectv1.GcpProjectParentType_organization {
		projectArgs.OrgId = pulumi.String(locals.GcpProject.Spec.ParentId)
	}
	if locals.GcpProject.Spec.ParentType == gcpprojectv1.GcpProjectParentType_folder {
		projectArgs.FolderId = pulumi.String(locals.GcpProject.Spec.ParentId)
	}

	// Create the GCP project using the generated projectId
	createdProject, err := organizations.NewProject(ctx, locals.GcpProject.Metadata.Name, projectArgs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create GCP project")
	}

	return createdProject, nil
}

// makeProjectIdSafe transforms the given name into a GCP-valid prefix:
//   - Lowercase letters, digits, hyphens only
//   - Must start with a letter
//   - Trim or replace invalid chars
//   - We'll do a best-effort approach, then the random suffix ensures uniqueness.
func makeProjectIdSafe(input string) string {
	if input == "" {
		input = "proj"
	}
	// Lowercase
	safe := strings.ToLower(input)

	// Replace any non letter/digit/hyphen with '-'
	out := make([]rune, 0, len(safe))
	for _, c := range safe {
		if (c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') ||
			(c == '-') {
			out = append(out, c)
		} else {
			out = append(out, '-')
		}
	}
	safe = string(out)

	// Ensure starts with letter
	if safe[0] < 'a' || safe[0] > 'z' {
		safe = "p" + safe
	}
	// Remove leading hyphens if any remain
	safe = strings.TrimLeft(safe, "-")

	// Trim if too long so we have room for suffix
	// We'll leave at most 27 chars, leaving space for "-xyz"
	if len(safe) > 27 {
		safe = safe[:27]
	}

	// Remove trailing hyphens if any
	safe = strings.TrimRight(safe, "-")

	// If empty or too short, pad
	if safe == "" {
		safe = "proj"
	}
	for len(safe) < 3 {
		safe += "x"
	}

	return safe
}
