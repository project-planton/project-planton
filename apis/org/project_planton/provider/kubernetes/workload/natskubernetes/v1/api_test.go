package natskubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestNatsKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "NatsKubernetes Suite")
}

var _ = ginkgo.Describe("NatsKubernetes Custom Validation Tests", func() {
	var input *NatsKubernetes

	ginkgo.BeforeEach(func() {
		input = &NatsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "NatsKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "nats-demo",
			},
			Spec: &NatsKubernetesSpec{
				ServerContainer: &NatsKubernetesServerContainer{
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
