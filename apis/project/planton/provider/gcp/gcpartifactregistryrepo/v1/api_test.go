package gcpartifactregistryrepov1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpArtifactRegistryRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpArtifactRegistryRepo Suite")
}

var _ = Describe("GcpArtifactRegistryRepo Custom Validation Tests", func() {
	var input *GcpArtifactRegistryRepo

	BeforeEach(func() {
		input = &GcpArtifactRegistryRepo{
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
	})

	Describe("When valid input is passed", func() {
		Context("gcp_artifact_registry_repo", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
