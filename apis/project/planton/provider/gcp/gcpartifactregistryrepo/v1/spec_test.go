package gcpartifactregistryrepov1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpArtifactRegistryRepo(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpArtifactRegistryRepo Suite")
}

var _ = ginkgo.Describe("GcpArtifactRegistryRepo Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_artifact_registry_repo", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
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
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with public access enabled", func() {
				input := &GcpArtifactRegistryRepo{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpArtifactRegistryRepo",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-repo",
					},
					Spec: &GcpArtifactRegistryRepoSpec{
						RepoFormat:         GcpArtifactRegistryRepoFormat_DOCKER,
						ProjectId:          "some-project-id",
						Region:             "us-west2",
						EnablePublicAccess: true,
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
