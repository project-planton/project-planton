package kubernetesgrafanav1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestKubernetesGrafana(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesGrafana Suite")
}

var _ = ginkgo.Describe("KubernetesGrafana Custom Validation Tests", func() {
	var input *KubernetesGrafana

	ginkgo.BeforeEach(func() {
		input = &KubernetesGrafana{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesGrafana",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-grafana",
			},
			Spec: &KubernetesGrafanaSpec{
				Container: &KubernetesGrafanaSpecContainer{},
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
