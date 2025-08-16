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

var _ = Describe("AwsCloudFrontSpec validations", func() {
	var spec *AwsCloudFrontSpec

	BeforeEach(func() {
		spec = &AwsCloudFrontSpec{
			Enabled:        true,
			Aliases:        []string{"cdn.example.com"},
			CertificateArn: "arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
			PriceClass:     AwsCloudFrontSpec_PRICE_CLASS_100,
			Origins: []*AwsCloudFrontSpec_Origin{
				{
					DomainName: "my-bucket.s3.amazonaws.com",
					OriginPath: "/assets",
					IsDefault:  true,
				},
				{
					DomainName: "api.example.com",
					OriginPath: "",
					IsDefault:  false,
				},
			},
			DefaultRootObject: "index.html",
		}
	})

	It("accepts a valid CloudFront spec", func() {
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts a valid spec without aliases and certificate", func() {
		spec.Aliases = []string{}
		spec.CertificateArn = ""
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts a valid spec with single origin", func() {
		spec.Origins = []*AwsCloudFrontSpec_Origin{
			{
				DomainName: "my-bucket.s3.amazonaws.com",
				OriginPath: "",
				IsDefault:  true,
			},
		}
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	// Field validation tests
	It("fails when aliases contain invalid domain names", func() {
		spec.Aliases = []string{"invalid-domain"}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when aliases contain duplicate values", func() {
		spec.Aliases = []string{"cdn.example.com", "cdn.example.com"}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when certificate_arn has invalid format", func() {
		spec.CertificateArn = "invalid-arn"
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when price_class is undefined", func() {
		spec.PriceClass = AwsCloudFrontSpec_PriceClass(99)
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when origins is empty", func() {
		spec.Origins = []*AwsCloudFrontSpec_Origin{}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when origin domain_name is empty", func() {
		spec.Origins[0].DomainName = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when origin domain_name has invalid format", func() {
		spec.Origins[0].DomainName = "invalid-domain"
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when origin_path doesn't start with /", func() {
		spec.Origins[0].OriginPath = "assets"
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when default_root_object has invalid characters", func() {
		spec.DefaultRootObject = "index.html!"
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// CEL expression tests
	It("fails when aliases are provided but certificate_arn is empty (aliases_require_cert)", func() {
		spec.Aliases = []string{"cdn.example.com"}
		spec.CertificateArn = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("passes when aliases are empty and certificate_arn is empty (aliases_require_cert)", func() {
		spec.Aliases = []string{}
		spec.CertificateArn = ""
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("passes when aliases are provided and certificate_arn is set (aliases_require_cert)", func() {
		spec.Aliases = []string{"cdn.example.com"}
		spec.CertificateArn = "arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("fails when no origin is marked as default (exactly_one_default_origin)", func() {
		spec.Origins[0].IsDefault = false
		spec.Origins[1].IsDefault = false
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when multiple origins are marked as default (exactly_one_default_origin)", func() {
		spec.Origins[0].IsDefault = true
		spec.Origins[1].IsDefault = true
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("passes when exactly one origin is marked as default (exactly_one_default_origin)", func() {
		spec.Origins[0].IsDefault = true
		spec.Origins[1].IsDefault = false
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("passes when exactly one origin is marked as default in different position (exactly_one_default_origin)", func() {
		spec.Origins[0].IsDefault = false
		spec.Origins[1].IsDefault = true
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	// Edge cases
	It("accepts valid domain names with hyphens", func() {
		spec.Aliases = []string{"my-cdn.example.com"}
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts valid domain names with numbers", func() {
		spec.Aliases = []string{"cdn1.example.com"}
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts valid origin paths with underscores and hyphens", func() {
		spec.Origins[0].OriginPath = "/assets/images"
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts empty origin_path", func() {
		spec.Origins[0].OriginPath = ""
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts valid default_root_object with underscores", func() {
		spec.DefaultRootObject = "index_main.html"
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts empty default_root_object", func() {
		spec.DefaultRootObject = ""
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})
})
