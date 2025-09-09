package cloudflareworkerv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareWorkerSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudflareWorkerSpec Custom Validation Tests")
}

var _ = Describe("CloudflareWorkerSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("cloudflare_worker", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-worker",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName: "test-worker-script",
						ScriptSource: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/path/to/worker.js"},
						},
						CompatibilityDate: "2024-01-01",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
