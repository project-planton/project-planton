package gcpgkeclustercorev1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpGkeClusterCoreSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpGkeClusterCoreSpec Custom Validation Tests")
}

var _ = Describe("GcpGkeClusterCoreSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_gke_cluster_core", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpGkeClusterCore{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpGkeClusterCore",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-gke-cluster-core",
					},
					Spec: &GcpGkeClusterCoreSpec{
						ProjectId: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-project-123"},
						},
						Location: "us-central1",
						SubnetworkSelfLink: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "projects/test-project-123/regions/us-central1/subnetworks/test-subnet"},
						},
						ClusterSecondaryRangeName: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "pods-range"},
						},
						ServicesSecondaryRangeName: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "services-range"},
						},
						MasterIpv4CidrBlock: "10.0.0.0/28",
						RouterNatName: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-nat-gateway"},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
