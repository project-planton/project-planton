package natskubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestNatsKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NatsKubernetes Suite")
}

var _ = Describe("NatsKubernetes Custom Validation Tests", func() {
	var input *NatsKubernetes

	BeforeEach(func() {
		input = &NatsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "NatsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-solr",
			},
			Spec: &NatsKubernetesSpec{},
		}
	})

	Describe("When valid input is passed", func() {
		Context("nats_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
