package stackjobrunnerkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestStackJobRunnerKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "StackJobRunnerKubernetes Suite")
}

var _ = ginkgo.Describe("StackJobRunnerKubernetes Custom Validation Tests", func() {
	var input *StackJobRunnerKubernetes

	ginkgo.BeforeEach(func() {
		input = &StackJobRunnerKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "StackJobRunnerKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-stack-job-runner",
			},
			Spec: &StackJobRunnerKubernetesSpec{},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("stack_job_runner_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
