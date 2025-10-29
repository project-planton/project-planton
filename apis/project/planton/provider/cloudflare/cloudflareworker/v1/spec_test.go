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
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-worker",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName: "test-worker-script",
						ScriptSource: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/path/to/worker.js"},
						},
						AccountId:         "00000000000000000000000000000000",
						CompatibilityDate: "2024-01-01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with environment variables and secrets", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-worker-env",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName: "test-worker-with-env",
						ScriptSource: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/path/to/worker.js"},
						},
						AccountId:         "00000000000000000000000000000000",
						CompatibilityDate: "2024-01-01",
						Env: &CloudflareWorkerEnv{
							Variables: map[string]string{
								"LOG_LEVEL": "info",
								"ENV":       "production",
							},
							Secrets: map[string]string{
								"API_KEY": "test-secret-key",
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with route pattern", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-worker-route",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName: "test-worker-with-route",
						ScriptSource: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/path/to/worker.js"},
						},
						AccountId:         "00000000000000000000000000000000",
						RoutePattern:      "https://example.com/*",
						CompatibilityDate: "2024-01-01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("account_id validation", func() {

			ginkgo.It("should return error if account_id is missing", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-no-account",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName: "test-worker",
						ScriptSource: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/path/to/worker.js"},
						},
						CompatibilityDate: "2024-01-01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if account_id is not 32 characters", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-short-account",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName: "test-worker",
						ScriptSource: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/path/to/worker.js"},
						},
						AccountId:         "123",
						CompatibilityDate: "2024-01-01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if account_id contains non-hex characters", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-invalid-hex",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName: "test-worker",
						ScriptSource: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/path/to/worker.js"},
						},
						AccountId:         "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ",
						CompatibilityDate: "2024-01-01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("script_source validation", func() {

			ginkgo.It("should return error if script_source is missing", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-no-source",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName:        "test-worker",
						AccountId:         "00000000000000000000000000000000",
						CompatibilityDate: "2024-01-01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("compatibility_date validation", func() {

			ginkgo.It("should return error if compatibility_date has invalid format", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-invalid-date",
					},
					Spec: &CloudflareWorkerSpec{
						ScriptName: "test-worker",
						ScriptSource: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/path/to/worker.js"},
						},
						AccountId:         "00000000000000000000000000000000",
						CompatibilityDate: "2024/01/01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
