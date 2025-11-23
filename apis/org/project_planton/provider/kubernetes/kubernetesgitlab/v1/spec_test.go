package kubernetesgitlabv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project_planton/shared/foreignkey/v1"
)

func TestKubernetesGitlab(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesGitlab Suite")
}

var _ = ginkgo.Describe("KubernetesGitlab Custom Validation Tests", func() {
	var input *KubernetesGitlab

	ginkgo.BeforeEach(func() {
		input = &KubernetesGitlab{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesGitlab",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-gitlab",
			},
			Spec: &KubernetesGitlabSpec{
				TargetCluster: &kubernetes.KubernetesClusterSelector{
					ClusterKind: 607, // GKE cluster kind
					ClusterName: "test-cluster",
				},
				Namespace: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "test-namespace",
					},
				},
				Container: &KubernetesGitlabSpecContainer{},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gitlab_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
