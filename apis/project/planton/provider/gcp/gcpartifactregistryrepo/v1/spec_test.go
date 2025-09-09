package gcpartifactregistryrepov1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpArtifactRegistryRepoSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpArtifactRegistryRepoSpec Custom Validation Tests")
}

var _ = Describe("GcpArtifactRegistryRepoSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_artifact_registry_repo", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpArtifactRegistryRepo{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpArtifactRegistryRepo",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-artifact-registry",
					},
					Spec: &GcpArtifactRegistryRepoSpec{
						RepoFormat: GcpArtifactRegistryRepoFormat_DOCKER,
						ProjectId:  "test-project-123",
						Region:     "us-central1",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
