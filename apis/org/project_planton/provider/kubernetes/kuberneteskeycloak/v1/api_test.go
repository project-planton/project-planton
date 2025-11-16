package kuberneteskeycloakv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesKeycloak(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesKeycloak Suite")
}

var _ = ginkgo.Describe("KubernetesKeycloak Custom Validation Tests", func() {
	var input *KubernetesKeycloak

	ginkgo.BeforeEach(func() {
		input = &KubernetesKeycloak{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesKeycloak",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-keycloak",
			},
			Spec: &KubernetesKeycloakSpec{
				Container: &KubernetesKeycloakContainer{
					Resources: &kubernetes.ContainerResources{
						Limits: &kubernetes.CpuMemory{
							Cpu:    "1000m",
							Memory: "1Gi",
						},
						Requests: &kubernetes.CpuMemory{
							Cpu:    "50m",
							Memory: "100Mi",
						},
					},
				},
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "keycloak.example.com",
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("keycloak_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
