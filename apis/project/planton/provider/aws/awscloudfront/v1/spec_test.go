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
			CertificateArn: "arn:aws:acm:us-east-1:123456789012:certificate/abc",
			PriceClass:     AwsCloudFrontSpec_PRICE_CLASS_100,
			Origins: []*AwsCloudFrontSpec_Origin{{
				Id:         "origin-1",
				DomainName: "bucket.s3.amazonaws.com",
			}},
			DefaultOriginId:   "origin-1",
			DefaultRootObject: "index.html",
		}
	})

	It("accepts a valid spec", func() {
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("fails when aliases are set but certificate_arn is empty (CEL)", func() {
		spec.CertificateArn = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when aliases contain duplicates (unique)", func() {
		spec.Aliases = []string{"cdn.example.com", "cdn.example.com"}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when origins are empty (min_items)", func() {
		spec.Origins = nil
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when default_origin_id does not match any origin (CEL)", func() {
		spec.DefaultOriginId = "missing"
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when default_origin_id is empty (min_len)", func() {
		spec.DefaultOriginId = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})
})
