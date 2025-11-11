package grafanakubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestGrafanaKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GrafanaKubernetes Suite")
}

var _ = ginkgo.Describe("GrafanaKubernetes Custom Validation Tests", func() {
	var input *GrafanaKubernetes

	ginkgo.BeforeEach(func() {
		input = &GrafanaKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "GrafanaKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-grafana",
			},
			Spec: &GrafanaKubernetesSpec{
				Container: &GrafanaKubernetesSpecContainer{},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("grafana_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
