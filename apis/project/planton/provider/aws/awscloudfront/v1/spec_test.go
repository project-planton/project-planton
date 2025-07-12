package awscloudfrontv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var validator protovalidate.Validator

func TestAwsCloudFrontSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var _ = BeforeSuite(func() {
    var err error
    validator, err = protovalidate.New()
    Expect(err).NotTo(HaveOccurred())
})

// validSpec returns a fully-populated message that should satisfy all
// buf.validate constraints declared for AwsCloudFrontSpec.
func validSpec() *AwsCloudFrontSpec {
    return &AwsCloudFrontSpec{
        Aliases:           []string{"example.com"},
        Enabled:           true,
        Comment:           "example distribution",
        PriceClass:        "PriceClass_100",
        DefaultRootObject: "index.html",
        Origins: []*Origin{
            {
                Id:         "origin1",
                DomainName: "mybucket.s3.amazonaws.com",
            },
        },
        DefaultCacheBehavior: &DefaultCacheBehavior{
            TargetOriginId:       "origin1",
            AllowedMethods:      []string{"GET", "HEAD"},
            CachedMethods:       []string{"GET", "HEAD"},
            ViewerProtocolPolicy: ViewerProtocolPolicy_ALLOW_ALL,
            ForwardedValues: &ForwardedValues{
                QueryString: false,
                Cookies: &Cookies{
                    ForwardAll: false,
                },
            },
            Compress: false,
        },
        IsIpv6Enabled:     true,
        WaitForDeployment: false,
    }
}

var _ = Describe("AwsCloudFrontSpec validation", func() {
    It("accepts a fully valid spec", func() {
        spec := validSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    DescribeTable("invalid specs should be rejected",
        func(modify func(*AwsCloudFrontSpec)) {
            spec := validSpec()
            modify(spec)
            Expect(validator.Validate(spec)).NotTo(Succeed())
        },
        Entry("duplicate aliases", func(s *AwsCloudFrontSpec) {
            s.Aliases = append(s.Aliases, s.Aliases[0])
        }),
        Entry("invalid price class", func(s *AwsCloudFrontSpec) {
            s.PriceClass = "PriceClass_999"
        }),
        Entry("default root object contains whitespace", func(s *AwsCloudFrontSpec) {
            s.DefaultRootObject = "index file.html"
        }),
        Entry("no origins present", func(s *AwsCloudFrontSpec) {
            s.Origins = nil
        }),
        Entry("missing default cache behavior", func(s *AwsCloudFrontSpec) {
            s.DefaultCacheBehavior = nil
        }),
    )
})