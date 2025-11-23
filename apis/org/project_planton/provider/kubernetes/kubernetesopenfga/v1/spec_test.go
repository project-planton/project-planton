package kubernetesopenfgav1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project_planton/shared/foreignkey/v1"
)

func TestKubernetesOpenFga(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesOpenFga Suite")
}

var _ = ginkgo.Describe("KubernetesOpenFga Custom Validation Tests", func() {
	var input *KubernetesOpenFga

	ginkgo.BeforeEach(func() {
		input = &KubernetesOpenFga{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesOpenFga",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-openfga",
			},
			Spec: &KubernetesOpenFgaSpec{
				TargetCluster: &kubernetes.KubernetesClusterSelector{
					ClusterKind: cloudresourcekind.CloudResourceKind_GcpGkeCluster,
					ClusterName: "test-cluster",
				},
				Namespace: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "test-namespace",
					},
				},
				Container: &KubernetesOpenFgaContainer{
					Replicas: 1,
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
				Ingress: &KubernetesOpenFgaIngress{
					Enabled:  true,
					Hostname: "test-openfga.example.com",
				},
				Datastore: &KubernetesOpenFgaDataStore{
					Engine: "postgres",
					Uri:    "postgres://user:pass@localhost:5432/testdb",
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("openfga_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
