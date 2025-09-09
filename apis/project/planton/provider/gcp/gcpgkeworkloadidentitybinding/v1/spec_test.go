package gcpgkeworkloadidentitybindingv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpGkeWorkloadIdentityBindingSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpGkeWorkloadIdentityBindingSpec Custom Validation Tests")
}

var _ = Describe("GcpGkeWorkloadIdentityBindingSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_gke_workload_identity_binding", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpGkeWorkloadIdentityBinding{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpGkeWorkloadIdentityBinding",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-workload-identity-binding",
					},
					Spec: &GcpGkeWorkloadIdentityBindingSpec{
						ProjectId: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-project-123"},
						},
						ServiceAccountEmail: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "cert-manager@test-project-123.iam.gserviceaccount.com"},
						},
						KsaNamespace: "cert-manager",
						KsaName:      "cert-manager",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
