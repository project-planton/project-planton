package kubernetesargocdv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestKubernetesArgocd(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesArgocd Suite")
}

var _ = ginkgo.Describe("KubernetesArgocd Custom Validation Tests", func() {
	var input *KubernetesArgocd

	ginkgo.BeforeEach(func() {
		input = &KubernetesArgocd{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesArgocd",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-argocd",
			},
			Spec: &KubernetesArgocdSpec{
				Container: &KubernetesArgocdContainer{},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("argocd_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
