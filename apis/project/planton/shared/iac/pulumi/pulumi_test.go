package pulumi

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
)

func TestPulumiBackend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PulumiBackend Suite")
}

var _ = Describe("PulumiBackend", func() {
	When("http details are missing", func() {
		It("should return an error containing `[http.required]`", func() {
			pulumiBackendCredentialSpec := &PulumiBackend{
				Type: PulumiBackendType_http,
			}
			err := protovalidate.Validate(pulumiBackendCredentialSpec)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("[http.required]"))
		})
	})

	When("AWS S3 details are missing", func() {
		It("should return an error containing `[s3.required]`", func() {
			pulumiBackendCredentialSpec := &PulumiBackend{
				Type: PulumiBackendType_s3,
			}
			err := protovalidate.Validate(pulumiBackendCredentialSpec)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("[s3.required]"))
		})
	})

	When("Google Cloud Storage details are missing", func() {
		It("should return an error containing `[gcs.required]`", func() {
			pulumiBackendCredentialSpec := &PulumiBackend{
				Type: PulumiBackendType_gcs,
			}
			err := protovalidate.Validate(pulumiBackendCredentialSpec)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("[gcs.required]"))
		})
	})

	When("Azure Blob Storage details are missing", func() {
		It("should return an error containing `[azurerm.required]`", func() {
			pulumiBackendCredentialSpec := &PulumiBackend{
				Type: PulumiBackendType_azurerm,
			}
			err := protovalidate.Validate(pulumiBackendCredentialSpec)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("[azurerm.required]"))
		})
	})
})
