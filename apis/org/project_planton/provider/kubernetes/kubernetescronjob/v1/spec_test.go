package kubernetescronjobv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project_planton/shared/foreignkey/v1"
	"google.golang.org/protobuf/proto"
)

func TestKubernetesCronJob(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesCronJob Suite")
}

var _ = ginkgo.Describe("KubernetesCronJob Custom Validation Tests", func() {
	var input *KubernetesCronJob

	ginkgo.BeforeEach(func() {
		input = &KubernetesCronJob{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesCronJob",
			Metadata: &shared.CloudResourceMetadata{
				Name: "my-cron-job",
			},
			Spec: &KubernetesCronJobSpec{
				TargetCluster: &kubernetes.KubernetesClusterSelector{
					ClusterKind: cloudresourcekind.CloudResourceKind_GcpGkeCluster,
					ClusterName: "test-cluster",
				},
				Namespace: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "test-namespace",
					},
				},
				Image: &kubernetes.ContainerImage{
					Repo: "busybox",
					Tag:  "latest",
				},
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
				Env: &KubernetesCronJobContainerAppEnv{
					Variables: map[string]string{
						"ENV_VAR": "example",
					},
					Secrets: map[string]string{
						"SECRET_NAME": "secret_value",
					},
				},
				Schedule:          "0 0 * * *",
				ConcurrencyPolicy: proto.String("Forbid"),
				RestartPolicy:     proto.String("Never"),
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("cron_job_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
