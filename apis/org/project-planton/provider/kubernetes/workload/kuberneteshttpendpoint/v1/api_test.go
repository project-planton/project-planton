package kuberneteshttpendpointv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestKubernetesHttpEndpoint(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesHttpEndpoint Suite")
}

var _ = ginkgo.Describe("KubernetesHttpEndpoint Custom Validation Tests", func() {
	var input *KubernetesHttpEndpoint

	ginkgo.BeforeEach(func() {
		input = &KubernetesHttpEndpoint{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesHttpEndpoint",
			Metadata: &shared.CloudResourceMetadata{
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("kubernetes_http_endpoint", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
