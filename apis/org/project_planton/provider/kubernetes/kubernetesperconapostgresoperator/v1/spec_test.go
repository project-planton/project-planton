package kubernetesperconapostgresoperatorv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesPerconaPostgresOperator(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesPerconaPostgresOperator Suite")
}

var _ = ginkgo.Describe("KubernetesPerconaPostgresOperator Validation Tests", func() {
	var input *KubernetesPerconaPostgresOperator

	ginkgo.BeforeEach(func() {
		input = &KubernetesPerconaPostgresOperator{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesPerconaPostgresOperator",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-percona-postgres-operator",
			},
			Spec: &KubernetesPerconaPostgresOperatorSpec{
				TargetCluster: &kubernetes.KubernetesAddonTargetCluster{},
				Namespace:     "percona-postgres-operator",
				Container: &KubernetesPerconaPostgresOperatorSpecContainer{
					Resources: &kubernetes.ContainerResources{
						Limits: &kubernetes.CpuMemory{
							Cpu:    "1000m",
							Memory: "1Gi",
						},
						Requests: &kubernetes.CpuMemory{
							Cpu:    "100m",
							Memory: "256Mi",
						},
					},
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("percona_postgres_operator", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When namespace pattern validation is tested", func() {
		ginkgo.Context("with valid lowercase namespace", func() {
			ginkgo.It("should pass validation", func() {
				input.Spec.Namespace = "my-namespace-123"
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with invalid uppercase namespace", func() {
			ginkgo.It("should fail validation", func() {
				input.Spec.Namespace = "MyNamespace"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("with invalid special characters in namespace", func() {
			ginkgo.It("should fail validation", func() {
				input.Spec.Namespace = "my_namespace"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When container resources are required", func() {
		ginkgo.Context("with missing container spec", func() {
			ginkgo.It("should fail validation", func() {
				input.Spec.Container = nil
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
