package certmanagerv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestCertManager(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CertManager Suite")
}

var _ = ginkgo.Describe("CertManager Custom Validation Tests", func() {
	var input *CertManager

	ginkgo.BeforeEach(func() {
		input = &CertManager{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "CertManager",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-cert-manager",
			},
			Spec: &CertManagerSpec{
				Acme: &AcmeConfig{
					Email: "admin@example.com",
				},
				DnsProviders: []*DnsProviderConfig{
					{
						Name:     "cloudflare-test",
						DnsZones: []string{"example.com"},
						Provider: &DnsProviderConfig_Cloudflare{
							Cloudflare: &CloudflareProvider{
								ApiToken: "test-token",
							},
						},
					},
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("cert_manager", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
