package awscertmanagercertv1

import (
	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"testing"
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
				Route53HostedZoneId: "a-route53-hosted-zone-id",
				ValidationMethod:    "DNS",
			},
		}
	})

	Context("when valid input is passed", func() {
		It("should not return a validation error", func() {
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})
})
