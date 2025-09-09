package mongodbatlascredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestMongodbAtlasCredential(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MongodbAtlasCredential Suite")
}

var _ = Describe("MongodbAtlasCredentialSpec Custom Validation Tests", func() {
	var input *MongodbAtlasCredential

	BeforeEach(func() {
		input = &MongodbAtlasCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "MongodbAtlasCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-mongodb-cred",
			},
			Spec: &MongodbAtlasCredentialSpec{
				PublicKey:  "dummyPublicKey",
				PrivateKey: "dummyPrivateKey",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("with a valid MongodbAtlasCredentialSpec", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
