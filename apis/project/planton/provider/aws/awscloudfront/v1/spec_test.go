package awscloudfrontv1

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "github.com/bufbuild/protovalidate-go"
)

// Ensure the Ginkgo test runner is executed.
func TestAwsCloudFrontSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var validator protovalidate.Validator

var _ = BeforeSuite(func() {
    v, err := protovalidate.New()
    Expect(err).NotTo(HaveOccurred())
    validator = *v // store the value, not the pointer (per requirements)
})

var _ = Describe("AwsCloudFrontSpec", func() {
    Context("with a valid specification", func() {
        It("validates successfully", func() {
            spec := &AwsCloudFrontSpec{
                Name:    "MyDistribution",
                Enabled: true,
                Aliases: []string{"www.example.com"},
                Comment: "sample distribution",
                PriceClass: AwsCloudFrontSpec_PRICE_CLASS_ALL,
                Origins: []*AwsCloudFrontSpec_Origin{
                    {
                        Id:         "origin1",
                        DomainName: "origin.example.com",
                        S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                            OriginAccessControlEnabled: true,
                            OriginAccessIdentity:       "", // must be empty when OAC enabled
                        },
                    },
                },
                DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                    TargetOriginId:       "origin1",
                    ViewerProtocolPolicy: AwsCloudFrontSpec_ALLOW_ALL,
                    AllowedMethods:       []string{"GET"},
                },
                ViewerCertificate: &AwsCloudFrontSpec_ViewerCertificate{
                    CloudfrontDefaultCertificate: true,
                    MinimumProtocolVersion:       AwsCloudFrontSpec_TLSV1_2_2018,
                },
                Logging: &AwsCloudFrontSpec_LoggingConfig{
                    Enabled: false,
                    Bucket:  "", // empty because disabled
                },
                Restrictions: &AwsCloudFrontSpec_Restrictions{
                    GeoRestriction: &AwsCloudFrontSpec_GeoRestriction{
                        RestrictionType: AwsCloudFrontSpec_NONE,
                        Locations:       nil,
                    },
                },
                Tags: map[string]string{"env": "prod"},
            }

            Expect(validator.Validate(spec)).To(Succeed())
        })
    })

    Context("with invalid specifications", func() {
        It("fails when name is empty", func() {
            spec := &AwsCloudFrontSpec{
                Name:       "", // invalid – min_len = 1
                PriceClass: AwsCloudFrontSpec_PRICE_CLASS_ALL,
                Origins: []*AwsCloudFrontSpec_Origin{
                    {
                        Id:         "origin1",
                        DomainName: "origin.example.com",
                        S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                            OriginAccessControlEnabled: true,
                        },
                    },
                },
                DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                    TargetOriginId:       "origin1",
                    ViewerProtocolPolicy: AwsCloudFrontSpec_ALLOW_ALL,
                    AllowedMethods:       []string{"GET"},
                },
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when price class is unspecified (enum not_in rule)", func() {
            spec := &AwsCloudFrontSpec{
                Name:       "ValidName",
                PriceClass: AwsCloudFrontSpec_PRICE_CLASS_UNSPECIFIED,
                Origins: []*AwsCloudFrontSpec_Origin{
                    {
                        Id:         "origin1",
                        DomainName: "origin.example.com",
                        S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                            OriginAccessControlEnabled: true,
                        },
                    },
                },
                DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                    TargetOriginId:       "origin1",
                    ViewerProtocolPolicy: AwsCloudFrontSpec_ALLOW_ALL,
                    AllowedMethods:       []string{"GET"},
                },
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when no origins are supplied (min_items = 1)", func() {
            spec := &AwsCloudFrontSpec{
                Name:       "ValidName",
                PriceClass: AwsCloudFrontSpec_PRICE_CLASS_ALL,
                Origins:    []*AwsCloudFrontSpec_Origin{}, // none provided
                DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                    TargetOriginId:       "origin1",
                    ViewerProtocolPolicy: AwsCloudFrontSpec_ALLOW_ALL,
                    AllowedMethods:       []string{"GET"},
                },
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when an origin has both s3 and custom configs set (one_origin_type CEL)", func() {
            spec := &AwsCloudFrontSpec{
                Name:       "ValidName",
                PriceClass: AwsCloudFrontSpec_PRICE_CLASS_ALL,
                Origins: []*AwsCloudFrontSpec_Origin{
                    {
                        Id:         "origin1",
                        DomainName: "origin.example.com",
                        S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                            OriginAccessControlEnabled: true,
                        },
                        CustomOriginConfig: &AwsCloudFrontSpec_CustomOriginConfig{
                            ProtocolPolicy: "https-only",
                            HttpPort:       80,
                            HttpsPort:      443,
                        },
                    },
                },
                DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                    TargetOriginId:       "origin1",
                    ViewerProtocolPolicy: AwsCloudFrontSpec_ALLOW_ALL,
                    AllowedMethods:       []string{"GET"},
                },
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails S3OriginConfig OAC/OAI exclusivity rule", func() {
            spec := &AwsCloudFrontSpec{
                Name:       "ValidName",
                PriceClass: AwsCloudFrontSpec_PRICE_CLASS_ALL,
                Origins: []*AwsCloudFrontSpec_Origin{
                    {
                        Id:         "origin1",
                        DomainName: "origin.example.com",
                        S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                            OriginAccessControlEnabled: false, // disabled but identity missing
                            OriginAccessIdentity:       "",    // should be non-empty
                        },
                    },
                },
                DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                    TargetOriginId:       "origin1",
                    ViewerProtocolPolicy: AwsCloudFrontSpec_ALLOW_ALL,
                    AllowedMethods:       []string{"GET"},
                },
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails viewer certificate mutual exclusivity rule", func() {
            spec := &AwsCloudFrontSpec{
                Name:       "ValidName",
                PriceClass: AwsCloudFrontSpec_PRICE_CLASS_ALL,
                Origins: []*AwsCloudFrontSpec_Origin{
                    {
                        Id:         "origin1",
                        DomainName: "origin.example.com",
                        S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                            OriginAccessControlEnabled: true,
                        },
                    },
                },
                DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                    TargetOriginId:       "origin1",
                    ViewerProtocolPolicy: AwsCloudFrontSpec_ALLOW_ALL,
                    AllowedMethods:       []string{"GET"},
                },
                ViewerCertificate: &AwsCloudFrontSpec_ViewerCertificate{
                    CloudfrontDefaultCertificate: false,
                    AcmCertificateArn:            "arn:aws:acm:us-east-1:123456789012:certificate/abcdef", // both ACM & IAM provided -> invalid
                    IamCertificateId:             "certificate123",
                    MinimumProtocolVersion:       AwsCloudFrontSpec_TLSV1_2_2018,
                },
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })
})
