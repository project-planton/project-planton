package certmanagerkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestCertManagerKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CertManagerKubernetes Suite")
}

var _ = ginkgo.Describe("CertManagerKubernetes Custom Validation Tests", func() {
	var input *CertManagerKubernetes

	ginkgo.BeforeEach(func() {
		input = &CertManagerKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "CertManagerKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-cert-manager",
			},
			Spec: &CertManagerKubernetesSpec{
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
		ginkgo.Context("cert_manager_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
