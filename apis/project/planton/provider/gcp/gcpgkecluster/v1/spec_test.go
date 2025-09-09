package gcpgkeclusterv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpGkeClusterSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpGkeClusterSpec Custom Validation Tests")
}

var _ = Describe("GcpGkeClusterSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_gke_cluster", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpGkeCluster{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpGkeCluster",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-gke-cluster",
					},
					Spec: &GcpGkeClusterSpec{
						ClusterProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
