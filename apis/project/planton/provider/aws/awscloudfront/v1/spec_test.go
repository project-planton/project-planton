package awscloudfrontv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    durationpb "google.golang.org/protobuf/types/known/durationpb"
)

func TestAwsCloudFrontSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var _ = Describe("AwsCloudFrontSpec", func() {
    var validator protovalidate.Validator

    BeforeSuite(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    validSpec := func() *AwsCloudFrontSpec {
        return &AwsCloudFrontSpec{
            Enabled:           true,
            Aliases:           []string{"example.com"},
            Comment:           "valid comment",
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
                MinTtl:              durationpb.New(0),
                DefaultTtl:          durationpb.New(0),
                MaxTtl:              durationpb.New(0),
                ForwardedValues: &AwsCloudFrontSpec_ForwardedValues{
                    QueryString: true,
                    Headers:     []string{"Authorization"},
                    Cookies: &AwsCloudFrontSpec_Cookies{
                        Forward: "none",
                    },
                },
            },
            OrderedCacheBehaviors: []*AwsCloudFrontSpec_CacheBehavior{},
            Logging: &AwsCloudFrontSpec_LoggingConfig{
                Bucket:         "logs.example.com.s3.amazonaws.com",
                Prefix:         "prefix/",
                IncludeCookies: true,
            },
            ViewerCertificate: &AwsCloudFrontSpec_ViewerCertificate{
                AcmCertificateArn:           "arn:aws:acm:us-east-1:123456789012:certificate/abcdef12-abcd-abcd-abcd-abcdef123456",
                CloudfrontDefaultCertificate: false,
                SslSupportMethod:             "sni-only",
                MinimumProtocolVersion:       "TLSv1.2_2021",
            },
            Restrictions: &AwsCloudFrontSpec_Restrictions{
                GeoRestriction: &AwsCloudFrontSpec_GeoRestriction{
                    RestrictionType: "none",
                },
            },
            WebAclId: "arn:aws:wafv2:us-east-1:123456789012:regional/webacl/name/ID",
            Tags: map[string]string{
                "env": "prod",
            },
        }
    }

    It("accepts a fully valid spec", func() {
        spec := validSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    DescribeTable("invalid specs",
        func(mutate func(spec *AwsCloudFrontSpec)) {
            spec := validSpec()
            mutate(spec)
            Expect(validator.Validate(spec)).ToNot(Succeed())
        },
        Entry("alias with empty string", func(s *AwsCloudFrontSpec) {
            s.Aliases = []string{""}
        }),
        Entry("invalid price class", func(s *AwsCloudFrontSpec) {
            s.PriceClass = "PriceClass_300"
        }),
        Entry("origin with invalid domain name", func(s *AwsCloudFrontSpec) {
            s.Origins[0].DomainName = "not a hostname"
        }),
        Entry("missing default cache behavior", func(s *AwsCloudFrontSpec) {
            s.DefaultCacheBehavior = nil
        }),
        Entry("invalid web_acl_id pattern", func(s *AwsCloudFrontSpec) {
            s.WebAclId = "invalid-arn"
        }),
        Entry("empty tag key", func(s *AwsCloudFrontSpec) {
            s.Tags[""] = "value"
        }),
    )
})
