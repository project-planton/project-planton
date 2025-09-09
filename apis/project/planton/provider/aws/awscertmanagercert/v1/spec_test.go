package awscertmanagercertv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
)

func TestAwsCertManagerCert(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsCertManagerCert Suite")
}

var _ = Describe("AwsCertManagerCert", func() {

	var input *AwsCertManagerCert

	BeforeEach(func() {
		input = &AwsCertManagerCert{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsCertManagerCert",
			Metadata: &shared.ApiResourceMetadata{
				Name: "a-test-name",
			},
			Spec: &AwsCertManagerCertSpec{
				PrimaryDomainName: "example.com",
				AlternateDomainNames: []string{
					"www.example.com",
					"test.example.com",
				},
				ValidationMethod: "DNS",
				Route53HostedZoneId: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-zone-id"},
				},
			},
		}
	})

	Context("when valid input is passed", func() {
		It("should not return a validation error", func() {
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})

	Context("Domain Pattern Validations", func() {

		Context("PrimaryDomainName", func() {
			It("should accept a valid apex domain", func() {
				input.Spec.PrimaryDomainName = "example.com"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should accept a valid wildcard domain", func() {
				input.Spec.PrimaryDomainName = "*.example.com"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject a domain missing a TLD", func() {
				input.Spec.PrimaryDomainName = "example"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should reject multiple wildcard asterisks", func() {
				input.Spec.PrimaryDomainName = "**.example.com"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should reject a domain with invalid characters", func() {
				input.Spec.PrimaryDomainName = "exa@mple.com"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("AlternateDomainNames", func() {
			It("should accept multiple valid domains", func() {
				input.Spec.AlternateDomainNames = []string{"www.example.com", "*.foo.org"}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject if any domain is invalid", func() {
				input.Spec.AlternateDomainNames = []string{"www.example.com", "invalid@@domain"}
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
