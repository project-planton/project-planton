package kubernetesnatsv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/plantonhq/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared"
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	foreignkeyv1 "github.com/plantonhq/project-planton/apis/org/project_planton/shared/foreignkey/v1"
)

func TestKubernetesNats(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesNats Suite")
}

var _ = ginkgo.Describe("KubernetesNats Custom Validation Tests", func() {
	var input *KubernetesNats

	ginkgo.BeforeEach(func() {
		input = &KubernetesNats{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesNats",
			Metadata: &shared.CloudResourceMetadata{
				Name: "nats-demo",
			},
			Spec: &KubernetesNatsSpec{
				TargetCluster: &kubernetes.KubernetesClusterSelector{
					ClusterKind: cloudresourcekind.CloudResourceKind_GcpGkeCluster,
					ClusterName: "test-cluster",
				},
				Namespace: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "nats-demo",
					},
				},
				ServerContainer: &KubernetesNatsServerContainer{
					Replicas: 3,      // satisfies gt:0
					DiskSize: "10Gi", // required by proto but standard, so fine to include
				},
				DisableJetStream: false,
				TlsEnabled:       false,
				DisableNatsBox:   false,
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("with replicas greater than zero", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
