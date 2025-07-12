package awscloudfrontv1

import (
    "testing"
    "time"

    pv "github.com/bufbuild/protovalidate-go"
    "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "google.golang.org/protobuf/types/known/durationpb"
)

func TestAwsCloudFrontSpecValidation(t *testing.T) {
    RegisterFailHandler(ginkgo.Fail)
    ginkgo.RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var _ = ginkgo.Describe("AwsCloudFrontSpec", func() {
    var (
        validator *pv.Validator
        err       error
    )

    ginkgo.BeforeSuite(func() {
        validator, err = pv.New()
        Expect(err).ToNot(HaveOccurred())
    })

    // Helper that returns a fully-valid spec instance.
    buildValidSpec := func() *AwsCloudFrontSpec {
        return &AwsCloudFrontSpec{
            Enabled:           true,
            Aliases:           []string{"example.com"},
            Comment:           "distribution for example.com",
            DefaultRootObject: "index.html",
            PriceClass:        "PriceClass_100",
            IsIpv6Enabled:     true,
            Origins: []*AwsCloudFrontSpec_Origin{
                {
                    Id:         "origin1",
                    DomainName: "origin.example.com",
                },
            },
            DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
                TargetOriginId:      "origin1",
                AllowedMethods:      []string{"GET"},
                CachedMethods:       []string{"GET"},
                ViewerProtocolPolicy: "allow-all",
                Compress:           true,
                MinTtl:             durationpb.New(0),
                DefaultTtl:         durationpb.New(10 * time.Second),
                MaxTtl:             durationpb.New(60 * time.Second),
            },
            WebAclId: "arn:aws:wafv2:us-east-1:123456789012:regional/webacl/mywebacl/abcd1234",
            Tags: map[string]string{
                "env": "test",
            },
        }
    }

    ginkgo.It("accepts a valid spec", func() {
        spec := buildValidSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    ginkgo.DescribeTable("rejects invalid specs",
        func(mutator func(*AwsCloudFrontSpec)) {
            spec := buildValidSpec()
            mutator(spec)
            Expect(validator.Validate(spec)).ToNot(Succeed())
        },
        ginkgo.Entry("invalid price_class", func(s *AwsCloudFrontSpec) {
            s.PriceClass = "PriceClass_999"
        }),
        ginkgo.Entry("blank alias", func(s *AwsCloudFrontSpec) {
            s.Aliases = append(s.Aliases, "")
        }),
        ginkgo.Entry("no origins", func(s *AwsCloudFrontSpec) {
            s.Origins = nil
        }),
        ginkgo.Entry("missing default_cache_behavior", func(s *AwsCloudFrontSpec) {
            s.DefaultCacheBehavior = nil
        }),
        ginkgo.Entry("invalid web_acl_id", func(s *AwsCloudFrontSpec) {
            s.WebAclId = "not-an-arn"
        }),
    )
})
