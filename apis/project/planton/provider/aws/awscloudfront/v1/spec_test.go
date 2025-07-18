package awscloudfrontv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestAwsCloudFrontSpec(t *testing.T) {
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

    It("accepts a valid specification", func() {
        spec := &AwsCloudFrontSpec{
            Name:    "my-distribution",
            Enabled: true,
            PriceClass: AwsCloudFrontSpec_PriceClass_PRICE_CLASS_ALL,
            Origins: []*AwsCloudFrontSpec_Origin{
                {
                    Id:         "origin1",
                    DomainName:  "example.com",
                    S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                        OriginAccessControlEnabled: false,
                        OriginAccessIdentity:      "origin-access-identity/cloudfront/ABCDEFG",
                    },
                },
            },
            DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                TargetOriginId:       "origin1",
                ViewerProtocolPolicy: AwsCloudFrontSpec_ViewerProtocolPolicy_ALLOW_ALL,
                AllowedMethods:       []string{"GET"},
            },
        }

        Expect(validator.Validate(spec)).To(Succeed())
    })

    It("rejects an invalid specification (empty name)", func() {
        spec := &AwsCloudFrontSpec{
            Name:    "", // invalid: name is required and must have min_len 1
            Enabled: true,
            PriceClass: AwsCloudFrontSpec_PriceClass_PRICE_CLASS_ALL,
            Origins: []*AwsCloudFrontSpec_Origin{
                {
                    Id:         "origin1",
                    DomainName:  "example.com",
                    S3OriginConfig: &AwsCloudFrontSpec_S3OriginConfig{
                        OriginAccessControlEnabled: false,
                        OriginAccessIdentity:      "origin-access-identity/cloudfront/ABCDEFG",
                    },
                },
            },
            DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                TargetOriginId:       "origin1",
                ViewerProtocolPolicy: AwsCloudFrontSpec_ViewerProtocolPolicy_ALLOW_ALL,
                AllowedMethods:       []string{"GET"},
            },
        }

        Expect(validator.Validate(spec)).NotTo(Succeed())
    })
})
