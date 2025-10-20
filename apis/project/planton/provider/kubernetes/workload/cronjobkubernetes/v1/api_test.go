package cronjobkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	"google.golang.org/protobuf/proto"
)

func TestCronJobKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CronJobKubernetes Suite")
}

var _ = ginkgo.Describe("CronJobKubernetes Custom Validation Tests", func() {
	var input *CronJobKubernetes

	ginkgo.BeforeEach(func() {
		input = &CronJobKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "CronJobKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "my-cron-job",
			},
			Spec: &CronJobKubernetesSpec{
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
				Env: &CronJobKubernetesContainerAppEnv{
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
