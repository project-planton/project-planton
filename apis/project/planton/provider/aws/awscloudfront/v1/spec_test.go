package awscloudfrontv1

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "github.com/bufbuild/protovalidate-go"
    "google.golang.org/protobuf/types/known/durationpb"
)

func TestAwsCloudFrontSpec(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var _ = Describe("AwsCloudFrontSpec validation", func() {
    var validator *protovalidate.Validator

    BeforeEach(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    newValidSpec := func() *AwsCloudFrontSpec {
        return &AwsCloudFrontSpec{
            Enabled:           true,
            Aliases:           []string{"example.com"},
            Comment:           "My distribution",
            DefaultRootObject: "index.html",
            PriceClass:        "PriceClass_100",
            IsIpv6Enabled:     true,
            Origins: []*AwsCloudFrontSpec_Origin{
                {
                    Id:         "origin1",
                    DomainName: "example.com",
                    OriginPath: "/content",
                    S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                        OriginAccessIdentity: "origin-access-identity/cloudfront/ABCDEFG1234567",
                    },
                },
            },
            DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                TargetOriginId:       "origin1",
                AllowedMethods:       []string{"GET", "HEAD"},
                CachedMethods:        []string{"GET", "HEAD"},
                ViewerProtocolPolicy: "allow-all",
                Compress:            true,
                MinTtl:              &durationpb.Duration{Seconds: 0},
                DefaultTtl:          &durationpb.Duration{Seconds: 60},
                MaxTtl:              &durationpb.Duration{Seconds: 3600},
            },
            OrderedCacheBehaviors: []*AwsCloudFrontSpec_CacheBehavior{},
            Logging: &AwsCloudFrontSpec_LoggingConfig{
                Bucket:         "logs.example.com.s3.amazonaws.com",
                Prefix:         "prefix/",
                IncludeCookies: true,
            },
            ViewerCertificate: &AwsCloudFrontSpec_ViewerCertificate{
                AcmCertificateArn:            "arn:aws:acm:us-east-1:123456789012:certificate/abcdef12-3456-7890-abcd-ef1234567890",
                CloudfrontDefaultCertificate: false,
                SslSupportMethod:             "sni-only",
                MinimumProtocolVersion:       "TLSv1.2_2018",
            },
            Restrictions: &AwsCloudFrontSpec_Restrictions{
                GeoRestriction: &AwsCloudFrontSpec_GeoRestriction{
                    RestrictionType: "none",
                    Locations:       []string{},
                },
            },
            WebAclId: "arn:aws:wafv2:us-east-1:123456789012:regional/webacl/Example/12345678-1234-1234-1234-123456789012",
            Tags:     map[string]string{"env": "prod"},
        }
    }

    It("accepts a fully valid spec", func() {
        spec := newValidSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    DescribeTable("rejects invalid specifications",
        func(mutate func(spec *AwsCloudFrontSpec)) {
            spec := newValidSpec()
            mutate(spec)
            Expect(validator.Validate(spec)).NotTo(Succeed())
        },
        Entry("invalid price class", func(spec *AwsCloudFrontSpec) {
            spec.PriceClass = "PriceClass_Invalid"
        }),
        Entry("missing origins", func(spec *AwsCloudFrontSpec) {
            spec.Origins = nil
        }),
        Entry("empty default root object", func(spec *AwsCloudFrontSpec) {
            spec.DefaultRootObject = ""
        }),
        Entry("alias empty string", func(spec *AwsCloudFrontSpec) {
            spec.Aliases = []string{""}
        }),
        Entry("invalid web_acl_id pattern", func(spec *AwsCloudFrontSpec) {
            spec.WebAclId = "invalid-arn"
        }),
    )
})
