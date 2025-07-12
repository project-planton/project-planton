package awscloudfrontv1

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/bufbuild/protovalidate-go"
    "google.golang.org/protobuf/types/known/durationpb"
)

// NOTE: this file intentionally uses the same package name as the generated
// *.pb.go files. Do NOT add a _test suffix, otherwise the package would import
// itself which is disallowed by the exercise constraints.

func TestAwsCloudFrontSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var _ = Describe("AwsCloudFrontSpec validation", func() {
    var (
        validator *protovalidate.Validator
        err       error
    )

    BeforeEach(func() {
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    It("accepts a fully valid spec", func() {
        spec := makeValidSpec()
        Expect(validator.Validate(spec)).NotTo(HaveOccurred())
    })

    It("rejects an empty alias entry", func() {
        spec := makeValidSpec()
        spec.Aliases = append(spec.Aliases, "")
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects an unsupported price class", func() {
        spec := makeValidSpec()
        spec.PriceClass = "PriceClass_999"
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects when default cache behavior is missing", func() {
        spec := makeValidSpec()
        spec.DefaultCacheBehavior = nil
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects when no origins are configured", func() {
        spec := makeValidSpec()
        spec.Origins = nil
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })
})

// makeValidSpec builds a spec that satisfies all buf.validate constraints. The
// helpers are kept small to focus on validation logic rather than business
// semantics.
func makeValidSpec() *AwsCloudFrontSpec {
    return &AwsCloudFrontSpec{
        Enabled:           true,
        Aliases:           []string{"example.com"},
        Comment:           "test distribution",
        DefaultRootObject: "index.html",
        PriceClass:        "PriceClass_100",
        IsIpv6Enabled:     true,
        Origins: []*AwsCloudFrontSpec_Origin{
            {
                Id:         "origin1",
                DomainName: "mybucket.s3.amazonaws.com",
                OriginPath: "/content",
                CustomHeaders: []*AwsCloudFrontSpec_CustomHeader{{
                    Name:  "X-Test",
                    Value: "value",
                }},
                S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                    OriginAccessIdentity: "origin-access-identity/cloudfront/ABC123DEF456",
                },
            },
        },
        DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
            TargetOriginId:        "origin1",
            AllowedMethods:        []string{"GET", "HEAD"},
            CachedMethods:         []string{"GET", "HEAD"},
            ViewerProtocolPolicy:  "allow-all",
            Compress:             true,
            MinTtl:               durationpb.New(0),
            DefaultTtl:           durationpb.New(0),
            MaxTtl:               durationpb.New(0),
            ForwardedValues: &AwsCloudFrontSpec_ForwardedValues{
                QueryString: true,
                Headers:     []string{"Authorization"},
                Cookies: &AwsCloudFrontSpec_Cookies{
                    Forward: "none",
                },
            },
        },
        OrderedCacheBehaviors: []*AwsCloudFrontSpec_CacheBehavior{
            {
                PathPattern:         "/images/*",
                TargetOriginId:      "origin1",
                AllowedMethods:      []string{"GET"},
                CachedMethods:       []string{"GET"},
                ViewerProtocolPolicy: "redirect-to-https",
                Compress:           true,
                MinTtl:             durationpb.New(0),
            },
        },
        Logging: &AwsCloudFrontSpec_LoggingConfig{
            Bucket:         "logs.example.com.s3.amazonaws.com",
            Prefix:         "cf/",
            IncludeCookies: false,
        },
        ViewerCertificate: &AwsCloudFrontSpec_ViewerCertificate{
            AcmCertificateArn:         "arn:aws:acm:us-east-1:123456789012:certificate/abcdef12-3456-7890-abcd-ef1234567890",
            CloudfrontDefaultCertificate: false,
            SslSupportMethod:            "sni-only",
            MinimumProtocolVersion:      "TLSv1.2_2018",
        },
        Restrictions: &AwsCloudFrontSpec_Restrictions{
            GeoRestriction: &AwsCloudFrontSpec_GeoRestriction{
                RestrictionType: "none",
            },
        },
        WebAclId: "arn:aws:wafv2:us-east-1:123456789012:regional/webacl/name/12345678-1234-1234-1234-123456789012",
        Tags: map[string]string{
            "env": "test",
        },
    }
}