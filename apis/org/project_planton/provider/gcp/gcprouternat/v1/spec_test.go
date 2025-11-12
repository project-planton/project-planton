package gcprouternatv1

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project_planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestGcpRouterNatSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpRouterNatSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpRouterNatSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_router_nat", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpRouterNat{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpRouterNat",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-router-nat",
					},
					Spec: &GcpRouterNatSpec{
						VpcSelfLink: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "projects/test-project-123/global/networks/test-vpc"},
						},
						Region: "us-central1",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
