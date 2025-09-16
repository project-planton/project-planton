package keycloakkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestKeycloakKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KeycloakKubernetes Suite")
}

var _ = ginkgo.Describe("KeycloakKubernetes Custom Validation Tests", func() {
	var input *KeycloakKubernetes

	ginkgo.BeforeEach(func() {
		input = &KeycloakKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KeycloakKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-keycloak",
			},
			Spec: &KeycloakKubernetesSpec{
				Container: &KeycloakKubernetesContainer{
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
