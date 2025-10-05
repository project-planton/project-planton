package cloudflarecredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CloudflareCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("CloudflareCredentialSpec Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("with api_token auth scheme", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloudflare-cred",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with R2 credentials", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloudflare-cred-with-r2",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
						R2: &CloudflareCredentialsR2Spec{
							AccessKeyId:     "12345678901234567890",
							SecretAccessKey: "12345678901234567890abcdef",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with R2 credentials and endpoint", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloudflare-cred-r2-endpoint",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
						R2: &CloudflareCredentialsR2Spec{
							AccessKeyId:     "12345678901234567890",
							SecretAccessKey: "12345678901234567890abcdef",
							Endpoint:        "https://custom.r2.cloudflarestorage.com",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with legacy_api_key auth scheme", func() {

			ginkgo.It("should not return a validation error for valid legacy auth", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloudflare-legacy",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_legacy_api_key,
						ApiKey:     "12345678901234567890",
						Email:      "user@example.com",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with R2 credentials", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloudflare-legacy-r2",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_legacy_api_key,
						ApiKey:     "12345678901234567890",
						Email:      "user@example.com",
						AccountId:  "00000000000000000000000000000000",
						R2: &CloudflareCredentialsR2Spec{
							AccessKeyId:     "12345678901234567890",
							SecretAccessKey: "12345678901234567890abcdef",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("auth_scheme validation", func() {

			ginkgo.It("should return error if auth_scheme is unspecified", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-invalid-auth",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_cloudflare_auth_scheme_unspecified,
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("account_id validation", func() {

			ginkgo.It("should return error if account_id is missing", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-no-account",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if account_id is not 32 characters", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-short-account",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if account_id contains non-hex characters", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-invalid-hex",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("api_token auth scheme validation", func() {

			ginkgo.It("should return error if api_token is missing", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-no-token",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if api_token is too short", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-short-token",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "short",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if api_key is provided with api_token scheme", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-mixed-auth",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						ApiKey:     "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if email is provided with api_token scheme", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-mixed-email",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						Email:      "user@example.com",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("legacy_api_key auth scheme validation", func() {

			ginkgo.It("should return error if api_key is missing", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-no-api-key",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_legacy_api_key,
						Email:      "user@example.com",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if email is missing", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-no-email",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_legacy_api_key,
						ApiKey:     "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if api_key is too short", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-short-key",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_legacy_api_key,
						ApiKey:     "short",
						Email:      "user@example.com",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if email format is invalid", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-invalid-email",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_legacy_api_key,
						ApiKey:     "12345678901234567890",
						Email:      "not-an-email",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if api_token is provided with legacy scheme", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-mixed-legacy",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_legacy_api_key,
						ApiKey:     "12345678901234567890",
						ApiToken:   "12345678901234567890",
						Email:      "user@example.com",
						AccountId:  "00000000000000000000000000000000",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("R2 credentials validation", func() {

			ginkgo.It("should return error if R2 access_key_id is too short", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-r2-short-key",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
						R2: &CloudflareCredentialsR2Spec{
							AccessKeyId:     "short",
							SecretAccessKey: "12345678901234567890abcdef",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if R2 secret_access_key is too short", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-r2-short-secret",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
						R2: &CloudflareCredentialsR2Spec{
							AccessKeyId:     "12345678901234567890",
							SecretAccessKey: "short",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if R2 access_key_id is missing", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-r2-no-key",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
						R2: &CloudflareCredentialsR2Spec{
							SecretAccessKey: "12345678901234567890abcdef",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if R2 secret_access_key is missing", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-r2-no-secret",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
						R2: &CloudflareCredentialsR2Spec{
							AccessKeyId: "12345678901234567890",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if R2 endpoint is not a valid URI", func() {
				input := &CloudflareCredential{
					ApiVersion: "credential.project-planton.org/v1",
					Kind:       "CloudflareCredential",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-r2-invalid-endpoint",
					},
					Spec: &CloudflareCredentialSpec{
						AuthScheme: CloudflareAuthScheme_api_token,
						ApiToken:   "12345678901234567890",
						AccountId:  "00000000000000000000000000000000",
						R2: &CloudflareCredentialsR2Spec{
							AccessKeyId:     "12345678901234567890",
							SecretAccessKey: "12345678901234567890abcdef",
							Endpoint:        "not-a-valid-url",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})

