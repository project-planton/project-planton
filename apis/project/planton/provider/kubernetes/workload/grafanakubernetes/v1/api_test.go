package grafanakubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGrafanaKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GrafanaKubernetes Suite")
}

var _ = Describe("GrafanaKubernetes Custom Validation Tests", func() {
	var input *GrafanaKubernetes

	BeforeEach(func() {
		input = &GrafanaKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "GrafanaKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-grafana",
			},
			Spec: &GrafanaKubernetesSpec{
				Container: &GrafanaKubernetesSpecContainer{},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("grafana_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
