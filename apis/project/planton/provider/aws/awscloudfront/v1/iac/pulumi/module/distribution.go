package module

import (
	"fmt"

	"github.com/pkg/errors"
	awscloudfrontv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awscloudfront/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// distribution creates a CloudFront distribution based on the Spec.
type DistributionRefs struct {
	Distribution *cloudfront.Distribution
	DomainName   pulumi.StringOutput
	HostedZoneId pulumi.StringOutput
}

func distribution(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) (*DistributionRefs, error) {
	// Build Origins
	originArgs := cloudfront.DistributionOriginArray{}
	for _, o := range locals.Spec.Origins {
		originArgs = append(originArgs, cloudfront.DistributionOriginArgs{
			DomainName: pulumi.String(o.DomainName),
			OriginId:   pulumi.String(o.Id),
		})
	}

	// Default Cache Behavior
	defaultBehavior := locals.Spec.DefaultCacheBehavior
	if defaultBehavior.OriginId == "" {
		return nil, errors.New("default_cache_behavior.origin_id must be set")
	}

	viewerProtocolPolicy := pulumi.String("allow-all")
	switch defaultBehavior.ViewerProtocolPolicy {
	case awscloudfrontv1.AwsCloudFrontSpec_DefaultCacheBehavior_VIEWER_PROTOCOL_POLICY_UNSPECIFIED:
		viewerProtocolPolicy = pulumi.String("allow-all")
	case awscloudfrontv1.AwsCloudFrontSpec_DefaultCacheBehavior_ALLOW_ALL:
		viewerProtocolPolicy = pulumi.String("allow-all")
	case awscloudfrontv1.AwsCloudFrontSpec_DefaultCacheBehavior_HTTPS_ONLY:
		viewerProtocolPolicy = pulumi.String("https-only")
	case awscloudfrontv1.AwsCloudFrontSpec_DefaultCacheBehavior_REDIRECT_TO_HTTPS:
		viewerProtocolPolicy = pulumi.String("redirect-to-https")
	}

	methods := pulumi.ToStringArray([]string{"GET", "HEAD"})
	switch defaultBehavior.AllowedMethods {
	case awscloudfrontv1.AwsCloudFrontSpec_DefaultCacheBehavior_ALL:
		methods = pulumi.ToStringArray([]string{"GET", "HEAD", "OPTIONS", "PUT", "POST", "PATCH", "DELETE"})
	case awscloudfrontv1.AwsCloudFrontSpec_DefaultCacheBehavior_GET_HEAD_OPTIONS:
		methods = pulumi.ToStringArray([]string{"GET", "HEAD", "OPTIONS"})
	}

	defaultCacheBehavior := &cloudfront.DistributionDefaultCacheBehaviorArgs{
		TargetOriginId:       pulumi.String(defaultBehavior.OriginId),
		ViewerProtocolPolicy: viewerProtocolPolicy,
		Compress:             pulumi.Bool(defaultBehavior.Compress),
		CachedMethods:        pulumi.ToStringArray([]string{"GET", "HEAD"}),
		AllowedMethods:       methods,
	}
	if defaultBehavior.CachePolicyId != "" {
		defaultCacheBehavior.CachePolicyId = pulumi.String(defaultBehavior.CachePolicyId)
	}

	// Price class mapping
	priceClass := pulumi.String("PriceClass_100")
	switch locals.Spec.PriceClass {
	case awscloudfrontv1.AwsCloudFrontSpec_PRICE_CLASS_ALL:
		priceClass = pulumi.String("PriceClass_All")
	case awscloudfrontv1.AwsCloudFrontSpec_PRICE_CLASS_200:
		priceClass = pulumi.String("PriceClass_200")
	}

	// Distribution args
	args := &cloudfront.DistributionArgs{
		Enabled:              pulumi.Bool(true),
		Origins:              originArgs,
		DefaultCacheBehavior: defaultCacheBehavior,
		PriceClass:           priceClass,
	}

	if locals.Spec.CertificateArn != "" {
		args.ViewerCertificate = &cloudfront.DistributionViewerCertificateArgs{
			AcmCertificateArn:      pulumi.String(locals.Spec.CertificateArn),
			SslSupportMethod:       pulumi.String("sni-only"),
			MinimumProtocolVersion: pulumi.String("TLSv1.2_2021"),
		}
	}

	if len(locals.Spec.Aliases) > 0 {
		args.Aliases = pulumi.ToStringArray(locals.Spec.Aliases)
	}

	if locals.Spec.Logging.GetEnabled() {
		lc := &cloudfront.DistributionLoggingConfigArgs{
			Bucket:         pulumi.String(fmt.Sprintf("%s.s3.amazonaws.com", locals.Spec.Logging.BucketName)),
			IncludeCookies: pulumi.Bool(false),
		}
		if locals.Spec.Logging.Prefix != "" {
			lc.Prefix = pulumi.String(locals.Spec.Logging.Prefix)
		}
		args.LoggingConfig = lc
	}

	if locals.Spec.WebAclArn != "" {
		args.WebAclId = pulumi.String(locals.Spec.WebAclArn)
	}

	name := fmt.Sprintf("%s-dist", locals.AwsCloudFront.Metadata.Name)
	cdn, err := cloudfront.NewDistribution(ctx, name, args, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}
	return &DistributionRefs{
		Distribution: cdn,
		DomainName:   cdn.DomainName,
		HostedZoneId: cdn.HostedZoneId,
	}, nil
}
