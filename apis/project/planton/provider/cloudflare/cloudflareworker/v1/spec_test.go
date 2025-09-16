package cloudflareworkerv1

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareWorkerSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CloudflareWorkerSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("CloudflareWorkerSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("cloudflare_worker", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
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
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
