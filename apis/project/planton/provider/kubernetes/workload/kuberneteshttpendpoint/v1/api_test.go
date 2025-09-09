package kuberneteshttpendpointv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestKubernetesHttpEndpoint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KubernetesHttpEndpoint Suite")
}

var _ = Describe("KubernetesHttpEndpoint Custom Validation Tests", func() {
	var input *KubernetesHttpEndpoint

	BeforeEach(func() {
		input = &KubernetesHttpEndpoint{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesHttpEndpoint",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-endpoint",
			},
			Spec: &KubernetesHttpEndpointSpec{
				IsTlsEnabled:          true,
				CertClusterIssuerName: "my-cluster-issuer",
				IsGrpcWebCompatible:   true,
				RoutingRules: []*KubernetesHttpEndpointRoutingRule{
					{
						UrlPathPrefix: "/api",
						BackendService: &KubernetesHttpEndpointRoutingRuleBackendService{
							Name:      "backend-svc",
							Namespace: "default",
							Port:      8080,
						},
					},
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("kubernetes_http_endpoint", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
