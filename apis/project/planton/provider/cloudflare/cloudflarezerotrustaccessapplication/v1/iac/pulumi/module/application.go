package module

import (
	"fmt"

	"github.com/pkg/errors"
	cloudflarezerotrustaccessapplicationv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflarezerotrustaccessapplication/v1"
	"github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// application provisions the Access Application and its default policy.
func application(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudflare.Provider,
) (*cloudflare.AccessApplication, error) {

	// --- Access Application --------------------------------------------------
	appArgs := &cloudflare.AccessApplicationArgs{
		Name:   pulumi.String(locals.CloudflareZeroTrustAccessApplication.Spec.ApplicationName),
		ZoneId: pulumi.String(locals.CloudflareZeroTrustAccessApplication.Spec.ZoneId),
		Domain: pulumi.String(locals.CloudflareZeroTrustAccessApplication.Spec.Hostname),
		Type:   pulumi.StringPtr("self_hosted"),
	}

	if locals.CloudflareZeroTrustAccessApplication.Spec.SessionDurationMinutes > 0 {
		appArgs.SessionDuration = pulumi.StringPtr(
			fmt.Sprintf("%dm", locals.CloudflareZeroTrustAccessApplication.Spec.SessionDurationMinutes),
		)
	}

	createdAccessApplication, err := cloudflare.NewAccessApplication(
		ctx,
		"access_application",
		appArgs,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create access application")
	}

	// --- Access Policy -------------------------------------------------------
	var includeBlocks cloudflare.AccessPolicyIncludeArray
	if len(locals.CloudflareZeroTrustAccessApplication.Spec.AllowedEmails) > 0 {
		var emails pulumi.StringArray
		for _, e := range locals.CloudflareZeroTrustAccessApplication.Spec.AllowedEmails {
			emails = append(emails, pulumi.String(e))
		}
		includeBlocks = append(includeBlocks, &cloudflare.AccessPolicyIncludeArgs{
			Emails: emails,
		})
	}

	if len(locals.CloudflareZeroTrustAccessApplication.Spec.AllowedGoogleGroups) > 0 {
		var groups pulumi.StringArray
		for _, g := range locals.CloudflareZeroTrustAccessApplication.Spec.AllowedGoogleGroups {
			groups = append(groups, pulumi.String(g))
		}
		includeBlocks = append(includeBlocks, &cloudflare.AccessPolicyIncludeArgs{
			Groups: groups,
		})
	}

	var requireBlocks cloudflare.AccessPolicyRequireArray
	if locals.CloudflareZeroTrustAccessApplication.Spec.RequireMfa {
		requireBlocks = append(requireBlocks, &cloudflare.AccessPolicyRequireArgs{})
	}

	decision := "allow"
	if locals.CloudflareZeroTrustAccessApplication.Spec.PolicyType ==
		cloudflarezerotrustaccessapplicationv1.CloudflareZeroTrustPolicyType_BLOCK {
		decision = "deny"
	}

	createdAccessPolicy, err := cloudflare.NewAccessPolicy(
		ctx,
		"access_policy",
		&cloudflare.AccessPolicyArgs{
			ApplicationId: createdAccessApplication.ID(),
			ZoneId:        pulumi.String(locals.CloudflareZeroTrustAccessApplication.Spec.ZoneId),
			Name:          pulumi.String("default-policy"),
			Decision:      pulumi.String(decision),
			Precedence:    pulumi.IntPtr(1),
		},
		pulumi.Provider(cloudflareProvider),
		pulumi.DependsOn([]pulumi.Resource{createdAccessApplication}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create access policy")
	}

	// --- Stack Outputs -------------------------------------------------------
	ctx.Export(OpApplicationId, createdAccessApplication.ID())
	ctx.Export(OpPublicHostname, pulumi.String(locals.CloudflareZeroTrustAccessApplication.Spec.Hostname))
	ctx.Export(OpPolicyId, createdAccessPolicy.ID())

	return createdAccessApplication, nil
}
