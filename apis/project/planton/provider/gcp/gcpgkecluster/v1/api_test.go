package gcpgkeclusterv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpGkeCluster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpGkeCluster Suite")
}

var _ = Describe("GcpGkeCluster Custom Validation Tests", func() {
	var input *GcpGkeCluster

	BeforeEach(func() {
		input = &GcpGkeCluster{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpGkeCluster",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-gke-cluster",
			},
			Spec: &GcpGkeClusterSpec{
				ClusterProjectId: "some-gcp-project-id",
				SharedVpcConfig: &GcpGkeClusterSharedVpcConfig{
					IsEnabled:    true,
					VpcProjectId: "some-shared-vpc-project-id",
				},
				IsWorkloadLogsEnabled: true,
				ClusterAutoscalingConfig: &GcpGkeClusterAutoscalingConfig{
					IsEnabled:   false,
					CpuMinCores: 0,
					CpuMaxCores: 8,
					MemoryMinGb: 0,
					MemoryMaxGb: 32,
				},
				NodePools: []*GcpGkeClusterNodePool{
					{
						Name:         "pool1",
						MachineType:  "n2-custom-8-16234",
						MinNodeCount: 1,
						MaxNodeCount: 3,
					},
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("gcp_gke", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
