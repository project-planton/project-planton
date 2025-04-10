package cronjobkubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestCronJobKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CronJobKubernetes Suite")
}

var _ = Describe("CronJobKubernetes Custom Validation Tests", func() {
	var input *CronJobKubernetes

	BeforeEach(func() {
		input = &CronJobKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "CronJobKubernetes",
			Metadata: &shared.ApiResourceMetadata{
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
				ConcurrencyPolicy: "Forbid",
				RestartPolicy:     "Never",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("cron_job_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
