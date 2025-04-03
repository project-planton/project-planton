package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	awsalbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsalb/v1"
)

// alb creates an AWS Application Load Balancer based on AwsAlbSpec fields and any optional listeners.
func alb(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	spec := locals.AwsAlb.Spec

	var isInternal pulumi.BoolInput
	if spec.Scheme == "internal" {
		isInternal = pulumi.Bool(true)
	} else {
		isInternal = pulumi.Bool(false)
	}

	// Create the ALB
	albResource, err := lb.NewLoadBalancer(ctx, locals.AwsAlb.Metadata.Name, &lb.LoadBalancerArgs{
		Name:                     pulumi.String(locals.AwsAlb.Metadata.Name),
		LoadBalancerType:         pulumi.String("application"),
		SecurityGroups:           pulumi.ToStringArray(spec.SecurityGroups),
		Subnets:                  pulumi.ToStringArray(spec.Subnets),
		Internal:                 isInternal,
		IpAddressType:            pulumi.String(spec.IpAddressType),
		EnableDeletionProtection: pulumi.Bool(spec.EnableDeletionProtection),
		IdleTimeout:              pulumi.Int(int(spec.IdleTimeoutSeconds)),
		Tags:                     pulumi.ToStringMap(locals.AwsTags),
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create AWS ALB")
	}

	// Create ALB listeners (if any)
	if len(spec.Listeners) > 0 {
		if err := albListeners(ctx, albResource, spec.Listeners, provider, locals.AwsAlb.Metadata.Name); err != nil {
			return errors.Wrap(err, "unable to create ALB listeners")
		}
	}

	// Export stack outputs
	ctx.Export(OpAlbArn, albResource.Arn)
	ctx.Export(OpAlbName, albResource.Name)
	ctx.Export(OpAlbDnsName, albResource.DnsName)
	ctx.Export(OpAlbHostedZoneId, albResource.ZoneId)

	return nil
}

// albListeners creates one or more listeners for the given ALB using the repeated AwsAlbListener field.
func albListeners(
	ctx *pulumi.Context,
	albResource *lb.LoadBalancer,
	listenerSpecs []*awsalbv1.AwsAlbListener,
	provider *aws.Provider,
	baseName string,
) error {
	for i, spec := range listenerSpecs {
		if spec.Protocol == "HTTPS" && spec.CertificateArn == "" {
			return fmt.Errorf("listener %d: certificateArn is required when using HTTPS protocol", i)
		}

		listenerName := fmt.Sprintf("%s-listener-%d", baseName, i)
		_, err := lb.NewListener(ctx, listenerName, &lb.ListenerArgs{
			LoadBalancerArn: albResource.Arn,
			Port:            pulumi.Int(int(spec.Port)),
			Protocol:        pulumi.String(spec.Protocol),
			SslPolicy:       pulumi.String(spec.SslPolicy),
			CertificateArn:  pulumi.String(spec.CertificateArn),
			// Basic default action that returns a 200 OK response.
			DefaultActions: lb.ListenerDefaultActionArray{
				&lb.ListenerDefaultActionArgs{
					Type: pulumi.String("fixed-response"),
					FixedResponse: &lb.ListenerDefaultActionFixedResponseArgs{
						ContentType: pulumi.String("text/plain"),
						StatusCode:  pulumi.String("200"),
						MessageBody: pulumi.String("OK"),
					},
				},
			},
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrapf(err, "unable to create ALB listener %d", i)
		}
	}
	return nil
}
