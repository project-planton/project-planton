package digitaloceankubernetesclusterv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestDigitalOceanKubernetesClusterSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DigitalOceanKubernetesClusterSpec Custom Validation Tests")
}

var _ = Describe("DigitalOceanKubernetesClusterSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("digitalocean_kubernetes_cluster", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &DigitalOceanKubernetesCluster{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanKubernetesCluster",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-k8s-cluster",
					},
					Spec: &DigitalOceanKubernetesClusterSpec{
						ClusterName:       "test-cluster",
						Region:            digitalocean.DigitalOceanRegion_nyc3,
						KubernetesVersion: "1.26.3",
						Vpc: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-vpc"},
						},
						DefaultNodePool: &DigitalOceanKubernetesClusterDefaultNodePool{
							Size:      "s-2vcpu-4gb",
							NodeCount: 3,
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
