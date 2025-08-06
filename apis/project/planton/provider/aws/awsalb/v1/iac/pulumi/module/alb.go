package module

import (
	"fmt"
	"github.com/project-planton/project-planton/internal/valuefrom"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	awsalbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsalb/v1"
)

// alb creates an AWS Application Load Balancer and, if SSL is enabled, two listeners:
//  1. HTTP (port 80) → auto-redirect to HTTPS
//  2. HTTPS (port 443) with the supplied certificate ARN
func alb(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) (*lb.LoadBalancer, error) {
	createdLoadBalancer, err := lb.NewLoadBalancer(ctx, locals.AwsAlb.Metadata.Name, &lb.LoadBalancerArgs{
		Name:                     pulumi.String(locals.AwsAlb.Metadata.Name),
		LoadBalancerType:         pulumi.String("application"),
		SecurityGroups:           pulumi.ToStringArray(valuefrom.ToStringArray(locals.AwsAlb.Spec.SecurityGroups)),
		Subnets:                  pulumi.ToStringArray(valuefrom.ToStringArray(locals.AwsAlb.Spec.Subnets)),
		Internal:                 pulumi.Bool(locals.AwsAlb.Spec.Internal),
		IpAddressType:            pulumi.String("ipv4"),
		EnableDeletionProtection: pulumi.Bool(locals.AwsAlb.Spec.DeleteProtectionEnabled),
		IdleTimeout:              pulumi.Int(int(locals.AwsAlb.Spec.IdleTimeoutSeconds)),
		Tags:                     pulumi.ToStringMap(locals.AwsTags),
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, errors.Wrap(err, "unable to create AWS ALB")
	}

	// If SSL is enabled, create the typical HTTP->HTTPS + HTTPS listeners
	if locals.AwsAlb.Spec.Ssl != nil && locals.AwsAlb.Spec.Ssl.Enabled {
		if locals.AwsAlb.Spec.Ssl.CertificateArn == nil || locals.AwsAlb.Spec.Ssl.CertificateArn.GetValue() == "" {
			return nil, fmt.Errorf("ssl.enabled is true, but ssl.certificate_arn is not provided")
		}
		if err := sslListeners(ctx, createdLoadBalancer, locals.AwsAlb.Spec.Ssl, provider, locals.AwsAlb.Metadata.Name); err != nil {
			return nil, errors.Wrap(err, "unable to create SSL listeners")
		}
	}

	// Export key ALB outputs
	ctx.Export(OpAlbArn, createdLoadBalancer.Arn)
	ctx.Export(OpAlbName, createdLoadBalancer.Name)
	ctx.Export(OpAlbDnsName, createdLoadBalancer.DnsName)
	ctx.Export(OpAlbHostedZoneId, createdLoadBalancer.ZoneId)

	return createdLoadBalancer, nil
}

// sslListeners implements the simple "SSL enabled" approach.
// It creates:
//  1. HTTP listener on port 80 → redirect to HTTPS:443
//  2. HTTPS listener on port 443 with the user-supplied cert.
func sslListeners(
	ctx *pulumi.Context,
	albResource *lb.LoadBalancer,
	sslSpec *awsalbv1.AwsAlbSsl,
	provider *aws.Provider,
	baseName string,
) error {
	// 1) HTTP :80 => redirect => :443 (HTTPS)
	httpListenerName := fmt.Sprintf("%s-http-redirect", baseName)
	_, err := lb.NewListener(ctx, httpListenerName, &lb.ListenerArgs{
		LoadBalancerArn: albResource.Arn,
		Port:            pulumi.Int(80),
		Protocol:        pulumi.String("HTTP"),
		DefaultActions: lb.ListenerDefaultActionArray{
			&lb.ListenerDefaultActionArgs{
				Type: pulumi.String("redirect"),
				Redirect: &lb.ListenerDefaultActionRedirectArgs{
					Port:       pulumi.String("443"),
					Protocol:   pulumi.String("HTTPS"),
					StatusCode: pulumi.String("HTTP_301"),
				},
			},
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create HTTP->HTTPS redirect listener")
	}

	// 2) HTTPS :443 => use the user’s certificate ARN
	httpsListenerName := fmt.Sprintf("%s-https", baseName)
	_, err = lb.NewListener(ctx, httpsListenerName, &lb.ListenerArgs{
		LoadBalancerArn: albResource.Arn,
		Port:            pulumi.Int(443),
		Protocol:        pulumi.String("HTTPS"),
		CertificateArn:  pulumi.String(sslSpec.CertificateArn.GetValue()),
		SslPolicy:       pulumi.String("ELBSecurityPolicy-2016-08"), // Hard-coded 80/20
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
		return errors.Wrap(err, "unable to create HTTPS listener")
	}

	return nil
}
