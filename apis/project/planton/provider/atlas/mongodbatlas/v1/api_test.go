package mongodbatlasv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestMongodbAtlas(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "MongodbAtlas Suite")
}

var _ = ginkgo.Describe("KubernetesProviderConfig Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("mongodb_atlas", func() {
			var input *MongodbAtlas

			ginkgo.BeforeEach(func() {
				input = &MongodbAtlas{
					ApiVersion: "atlas.project-planton.org/v1",
					Kind:       "MongodbAtlas",
					Metadata: &shared.CloudResourceMetadata{
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

			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
