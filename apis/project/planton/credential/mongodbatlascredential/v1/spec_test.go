package mongodbatlascredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestMongodbAtlasCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "MongodbAtlasCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("MongodbAtlasCredentialSpec Validation Tests", func() {
	var input *MongodbAtlasCredential

	ginkgo.BeforeEach(func() {
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("with valid credentials", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("missing required fields", func() {

			ginkgo.It("should return error if public_key is missing", func() {
				input.Spec.PublicKey = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if private_key is missing", func() {
				input.Spec.PrivateKey = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
