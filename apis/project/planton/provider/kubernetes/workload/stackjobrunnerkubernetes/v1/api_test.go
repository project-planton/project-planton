package stackjobrunnerkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestStackJobRunnerKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StackJobRunnerKubernetes Suite")
}

var _ = Describe("StackJobRunnerKubernetes Custom Validation Tests", func() {
	var input *StackJobRunnerKubernetes

	BeforeEach(func() {
		input = &StackJobRunnerKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "StackJobRunnerKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-stack-job-runner",
			},
			Spec: &StackJobRunnerKubernetesSpec{},
		}
	})

	Describe("When valid input is passed", func() {
		Context("stack_job_runner_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
