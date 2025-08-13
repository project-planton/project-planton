package awscloudfrontv1

import (
	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AwsCloudFrontSpec validations", func() {
	var spec *AwsCloudFrontSpec

	BeforeEach(func() {
		spec = &AwsCloudFrontSpec{
			Aliases:        []string{"cdn.example.com"},
			CertificateArn: "arn:aws:acm:us-east-1:123456789012:certificate/abc",
			PriceClass:     AwsCloudFrontSpec_PRICE_CLASS_100,
			Origins: []*AwsCloudFrontSpec_Origin{{
				Id:         "origin-1",
				DomainName: "bucket.s3.amazonaws.com",
			}},
			DefaultCacheBehavior: &AwsCloudFrontSpec_DefaultCacheBehavior{
				OriginId:             "origin-1",
				ViewerProtocolPolicy: AwsCloudFrontSpec_DefaultCacheBehavior_HTTPS_ONLY,
			},
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

	It("fails when default behavior viewer_protocol_policy is unspecified (CEL)", func() {
		spec.DefaultCacheBehavior.ViewerProtocolPolicy = AwsCloudFrontSpec_DefaultCacheBehavior_VIEWER_PROTOCOL_POLICY_UNSPECIFIED
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when origins is empty (min_items)", func() {
		spec.Origins = nil
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when origin id is empty (min_len)", func() {
		spec.Origins[0].Id = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})
})
