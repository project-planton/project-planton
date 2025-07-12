package awscloudfrontv1

import (
    "testing"
    "time"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "google.golang.org/protobuf/types/known/durationpb"
)

func TestAwsCloudFrontSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var _ = Describe("AwsCloudFrontSpec validation", func() {
    var validator protovalidate.Validator

    BeforeEach(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    Context("with a valid spec", func() {
        It("passes validation", func() {
            spec := &AwsCloudFrontSpec{
                Enabled:           true,
                Aliases:           []string{"example.com"},
                Comment:           "my distribution",
                DefaultRootObject: "index.html",
                PriceClass:        "PriceClass_100",
                IsIpv6Enabled:     true,
                Origins: []*AwsCloudFrontSpec_Origin{
                    {
                        Id:         "origin1",
                        DomainName: "mybucket.s3.amazonaws.com",
                        OriginPath: "/assets",
                        CustomHeaders: []*AwsCloudFrontSpec_CustomHeader{{
                            Name:  "X-Test",
                            Value: "value",
                        }},
                        S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                            OriginAccessIdentity: "origin-access-identity/cloudfront/ABCDEFG1234567",
                        },
                        CustomOriginConfig: &AwsCloudFrontSpec_CustomOriginConfig{
                            OriginProtocolPolicy: "https-only",
                            HttpPort:             80,
                            HttpsPort:            443,
                            OriginSslProtocols:   []string{"TLSv1.2"},
                        },
                    },
                },
                DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                    TargetOriginId:       "origin1",
                    AllowedMethods:       []string{"GET", "HEAD"},
                    CachedMethods:        []string{"GET", "HEAD"},
                    ViewerProtocolPolicy: "redirect-to-https",
                    Compress:            true,
                    MinTtl:              durationpb.New(0 * time.Second),
                    DefaultTtl:          durationpb.New(60 * time.Second),
                    MaxTtl:              durationpb.New(300 * time.Second),
                    ForwardedValues: &AwsCloudFrontSpec_ForwardedValues{
                        QueryString: true,
                        Headers:     []string{"Authorization"},
                        Cookies: &AwsCloudFrontSpec_Cookies{
                            Forward: "all",
                        },
                    },
                },
                Logging: &AwsCloudFrontSpec_LoggingConfig{
                    Bucket:         "logs.example.com",
                    Prefix:         "cf/",
                    IncludeCookies: true,
                },
                ViewerCertificate: &AwsCloudFrontSpec_ViewerCertificate{
                    AcmCertificateArn:           "arn:aws:acm:us-east-1:123456789012:certificate/123e4567-e89b-12d3-a456-426614174000",
                    CloudfrontDefaultCertificate: false,
                    SslSupportMethod:            "sni-only",
                    MinimumProtocolVersion:      "TLSv1.2_2018",
                },
                Restrictions: &AwsCloudFrontSpec_Restrictions{
                    GeoRestriction: &AwsCloudFrontSpec_GeoRestriction{
                        RestrictionType: "whitelist",
                        Locations:       []string{"US", "CA"},
                    },
                },
                WebAclId: "arn:aws:wafv2:us-east-1:123456789012:regional/webacl/sample/12345678-1234-1234-1234-123456789012",
                Tags:     map[string]string{"env": "prod"},
            }

            err := validator.Validate(spec)
            Expect(err).NotTo(HaveOccurred())
        })
    })

    Context("with an invalid spec", func() {
        It("fails when default_cache_behavior is missing", func() {
            spec := &AwsCloudFrontSpec{}
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })

        It("fails when aliases contain empty string", func() {
            spec := minimalValidSpec()
            spec.Aliases = []string{""}
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })

        It("fails when price_class is out of allowed set", func() {
            spec := minimalValidSpec()
            spec.PriceClass = "PriceClass_999"
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })

        It("fails when web_acl_id has wrong pattern", func() {
            spec := minimalValidSpec()
            spec.WebAclId = "invalid-arn"
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })

        It("fails when origins are empty", func() {
            spec := minimalValidSpec()
            spec.Origins = []*AwsCloudFrontSpec_Origin{}
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })
    })
})

func minimalValidSpec() *AwsCloudFrontSpec {
    return &AwsCloudFrontSpec{
        Enabled:           true,
        Aliases:           []string{"example.com"},
        DefaultRootObject: "index.html",
        PriceClass:        "PriceClass_100",
        Origins: []*AwsCloudFrontSpec_Origin{{
            Id:         "origin1",
            DomainName: "mybucket.s3.amazonaws.com",
        }},
        DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
            TargetOriginId:       "origin1",
            AllowedMethods:       []string{"GET"},
            ViewerProtocolPolicy: "allow-all",
        },
    }
}