package gcpsecretsmanagerv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpSecretsManagerSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpSecretsManagerSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpSecretsManagerSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_secrets_manager", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
