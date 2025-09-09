package mongodbatlasv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestMongodbAtlas(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MongodbAtlas Suite")
}

var _ = Describe("KubernetesClusterCredentialSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("mongodb_atlas", func() {
			var input *MongodbAtlas

			BeforeEach(func() {
				input = &MongodbAtlas{
					ApiVersion: "atlas.project-planton.org/v1",
					Kind:       "MongodbAtlas",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-atlas-resource",
					},
					Spec: &MongodbAtlasSpec{
						ClusterConfig: &MongodbAtlasClusterConfig{
							ProjectId:                "some-project-id",
							ClusterType:              "REPLICASET",
							ElectableNodes:           3,
							Priority:                 7,
							ReadOnlyNodes:            0,
							CloudBackup:              true,
							AutoScalingDiskGbEnabled: true,
							MongoDbMajorVersion:      "7.0",
							ProviderName:             "AWS",
							ProviderInstanceSizeName: "M10",
						},
					},
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
