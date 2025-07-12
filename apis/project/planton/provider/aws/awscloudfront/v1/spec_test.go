package awscloudfrontv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/durationpb"
)

// -----------------------------------------------------------------------------
//  Ginkgo test runner
// -----------------------------------------------------------------------------
func TestAwsCloudFrontSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var _ = Describe("AwsCloudFrontSpec buf.validate rules", func() {
    var (
        validator protovalidate.Validator
    )

    BeforeEach(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    // ---------------------------------------------------------------------
    //  Positive case – the spec should pass validation
    // ---------------------------------------------------------------------
    It("accepts a fully valid specification", func() {
        spec := validSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    // ---------------------------------------------------------------------
    //  Negative cases – each should fail validation for the indicated reason
    // ---------------------------------------------------------------------
    It("rejects an empty alias entry", func() {
        spec := proto.Clone(validSpec()).(*AwsCloudFrontSpec)
        spec.Aliases = []string{""}
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects an invalid price class", func() {
        spec := proto.Clone(validSpec()).(*AwsCloudFrontSpec)
        spec.PriceClass = "PriceClass_999"
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("requires at least one origin", func() {
        spec := proto.Clone(validSpec()).(*AwsCloudFrontSpec)
        spec.Origins = nil
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("requires default_cache_behavior to be present", func() {
        spec := proto.Clone(validSpec()).(*AwsCloudFrontSpec)
        spec.DefaultCacheBehavior = nil
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects an origin with a non-hostname domain name", func() {
        spec := proto.Clone(validSpec()).(*AwsCloudFrontSpec)
        spec.Origins[0].DomainName = "not a host name"
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })
})

// -----------------------------------------------------------------------------
//  Helper – construct a valid AwsCloudFrontSpec instance
// -----------------------------------------------------------------------------
func validSpec() *AwsCloudFrontSpec {
    return &AwsCloudFrontSpec{
        Enabled:           true,
        Aliases:           []string{"example.com"},
        Comment:           "My CloudFront distribution",
        DefaultRootObject: "index.html",
        PriceClass:        "PriceClass_All",
        IsIpv6Enabled:     true,
        Origins: []*AwsCloudFrontSpec_Origin{
            {
                Id:         "myS3Origin",
                DomainName:  "mybucket.s3.amazonaws.com",
                OriginPath:  "/content",
                CustomHeaders: []*AwsCloudFrontSpec_CustomHeader{
                    {Name: "X-Test", Value: "true"},
                },
                S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                    OriginAccessIdentity: "origin-access-identity/cloudfront/ABCDEFG1234567",
                },
            },
        },
        DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
            TargetOriginId:       "myS3Origin",
            AllowedMethods:       []string{"GET", "HEAD"},
            CachedMethods:        []string{"GET", "HEAD"},
            ViewerProtocolPolicy: "redirect-to-https",
            Compress:             true,
            MinTtl:               durationpb.New(0),
            DefaultTtl:           durationpb.New(0),
            MaxTtl:               durationpb.New(0),
            ForwardedValues: &AwsCloudFrontSpec_ForwardedValues{
                QueryString: true,
                Headers:     []string{"Authorization"},
                Cookies: &AwsCloudFrontSpec_Cookies{
                    Forward: "all",
                },
            },
        },
        OrderedCacheBehaviors: nil,
        Logging: &AwsCloudFrontSpec_LoggingConfig{
            Bucket:         "logs.example.com",
            Prefix:         "cdn/",
            IncludeCookies: true,
        },
        ViewerCertificate: &AwsCloudFrontSpec_ViewerCertificate{
            AcmCertificateArn:            "arn:aws:acm:us-east-1:123456789012:certificate/123e4567-e89b-12d3-a456-426655440000",
            CloudfrontDefaultCertificate: false,
            SslSupportMethod:             "sni-only",
            MinimumProtocolVersion:       "TLSv1.2_2021",
        },
        Restrictions: &AwsCloudFrontSpec_Restrictions{
            GeoRestriction: &AwsCloudFrontSpec_GeoRestriction{
                RestrictionType: "none",
                Locations:       nil,
            },
        },
        WebAclId: "arn:aws:wafv2:us-east-1:123456789012:regional/webacl/myWebAcl/12345678-1234-1234-1234-123456789012",
        Tags: map[string]string{
            "env": "prod",
        },
    }
}
