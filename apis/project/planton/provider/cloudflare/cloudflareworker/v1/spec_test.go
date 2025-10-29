package cloudflareworkerv1

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

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
						AccountId: "00000000000000000000000000000000",
						Script: &CloudflareWorkerScript{
							Name: "test-worker-script",
							Bundle: &CloudflareWorkerScriptBundleR2Object{
								Bucket: "test-bucket",
								Path:   "test/script.js",
							},
						},
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
						AccountId: "00000000000000000000000000000000",
						Script: &CloudflareWorkerScript{
							Name: "test-worker-with-env",
							Bundle: &CloudflareWorkerScriptBundleR2Object{
								Bucket: "test-bucket",
								Path:   "test/script-with-env.js",
							},
						},
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
						AccountId: "00000000000000000000000000000000",
						Script: &CloudflareWorkerScript{
							Name: "test-worker-with-route",
							Bundle: &CloudflareWorkerScriptBundleR2Object{
								Bucket: "test-bucket",
								Path:   "test/script-with-route.js",
							},
						},
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
						Script: &CloudflareWorkerScript{
							Name: "test-worker",
							Bundle: &CloudflareWorkerScriptBundleR2Object{
								Bucket: "test-bucket",
								Path:   "test/script.js",
							},
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
						AccountId: "123",
						Script: &CloudflareWorkerScript{
							Name: "test-worker",
							Bundle: &CloudflareWorkerScriptBundleR2Object{
								Bucket: "test-bucket",
								Path:   "test/script.js",
							},
						},
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
						AccountId: "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ",
						Script: &CloudflareWorkerScript{
							Name: "test-worker",
							Bundle: &CloudflareWorkerScriptBundleR2Object{
								Bucket: "test-bucket",
								Path:   "test/script.js",
							},
						},
						CompatibilityDate: "2024-01-01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("script validation", func() {

			ginkgo.It("should return error if script is missing", func() {
				input := &CloudflareWorker{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareWorker",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-no-source",
					},
					Spec: &CloudflareWorkerSpec{
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
						AccountId: "00000000000000000000000000000000",
						Script: &CloudflareWorkerScript{
							Name: "test-worker",
							Bundle: &CloudflareWorkerScriptBundleR2Object{
								Bucket: "test-bucket",
								Path:   "test/script.js",
							},
						},
						CompatibilityDate: "2024/01/01",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
