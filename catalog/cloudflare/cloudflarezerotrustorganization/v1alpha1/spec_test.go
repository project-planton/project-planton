package cloudflarezerotrustorganizationv1alpha1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/plantonhq/planton/shared"
	foreignkeyv1 "github.com/plantonhq/planton/shared/foreignkey/v1"
	"google.golang.org/protobuf/proto"
)

func TestCloudflareZeroTrustOrganizationSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CloudflareZeroTrustOrganizationSpec Custom Validation Tests")
}

const testAccountID = "0da42c8d2132a9ddaf714f9e7c920711"

func literal(v string) *foreignkeyv1.StringValueOrRef {
	return &foreignkeyv1.StringValueOrRef{
		LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: v},
	}
}

func validOrg(spec *CloudflareZeroTrustOrganizationSpec) *CloudflareZeroTrustOrganization {
	return &CloudflareZeroTrustOrganization{
		ApiVersion: "cloudflare.planton.dev/v1alpha1",
		Kind:       "CloudflareZeroTrustOrganization",
		Metadata: &shared.CatalogObjectMetadata{
			Name: "test-org",
		},
		Spec: spec,
	}
}

func baseSpec() *CloudflareZeroTrustOrganizationSpec {
	return &CloudflareZeroTrustOrganizationSpec{
		AccountId:  testAccountID,
		AuthDomain: "acme-test",
		Name:       "Acme Zero Trust",
	}
}

var _ = ginkgo.Describe("CloudflareZeroTrustOrganizationSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {

		ginkgo.It("should accept an account-scoped organization", func() {
			gomega.Expect(protovalidate.Validate(validOrg(baseSpec()))).To(gomega.BeNil())
		})

		ginkgo.It("should accept a zone-scoped organization", func() {
			spec := baseSpec()
			spec.AccountId = ""
			spec.ZoneId = literal("9a7806061c88ada191ed06f989cc3dac")
			gomega.Expect(protovalidate.Validate(validOrg(spec))).To(gomega.BeNil())
		})

		ginkgo.It("should accept key rotation on an account-scoped organization", func() {
			spec := baseSpec()
			spec.KeyRotationIntervalDays = proto.Int32(30)
			gomega.Expect(protovalidate.Validate(validOrg(spec))).To(gomega.BeNil())
		})

		ginkgo.It("should accept a full MFA policy", func() {
			spec := baseSpec()
			spec.MfaConfig = &CloudflareZeroTrustOrganizationMfaConfig{
				AllowedAuthenticators: []string{"totp", "security_key"},
				SessionDuration:       "12h",
			}
			spec.MfaPivKeyRequirements = &CloudflareZeroTrustOrganizationMfaPivKeyRequirements{
				PinPolicy:   "once",
				TouchPolicy: "cached",
				SshKeyType:  []string{"ed25519", "ecdsa"},
				SshKeySize:  []int64{256, 4096},
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).To(gomega.BeNil())
		})

		ginkgo.It("should accept PIV and FIDO2 SSH keys as authenticators", func() {
			spec := baseSpec()
			spec.MfaConfig = &CloudflareZeroTrustOrganizationMfaConfig{
				AllowedAuthenticators: []string{"piv_key", "ssh_fido2_key", "totp"},
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).To(gomega.BeNil())
		})

		ginkgo.It("should accept service-token enforcement", func() {
			spec := baseSpec()
			spec.StrictServiceTokenAuth = proto.Bool(true)
			spec.WarpAuthNonBrowser_401 = proto.Bool(true)
			spec.ServiceTokenInactivity = &CloudflareZeroTrustOrganizationServiceTokenInactivity{
				Action:                  "disable",
				Enabled:                 true,
				InactivityThresholdDays: 90,
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).To(gomega.BeNil())
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {

		ginkgo.It("should reject the retired ssh_piv_key authenticator", func() {
			spec := baseSpec()
			spec.MfaConfig = &CloudflareZeroTrustOrganizationMfaConfig{
				AllowedAuthenticators: []string{"ssh_piv_key"},
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject an unknown inactivity action", func() {
			spec := baseSpec()
			spec.ServiceTokenInactivity = &CloudflareZeroTrustOrganizationServiceTokenInactivity{
				Action:                  "archive",
				InactivityThresholdDays: 90,
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject an inactivity threshold under 30 days", func() {
			spec := baseSpec()
			spec.ServiceTokenInactivity = &CloudflareZeroTrustOrganizationServiceTokenInactivity{
				Action:                  "delete",
				InactivityThresholdDays: 7,
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject both account_id and zone_id", func() {
			spec := baseSpec()
			spec.ZoneId = literal("9a7806061c88ada191ed06f989cc3dac")
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject neither account_id nor zone_id", func() {
			spec := baseSpec()
			spec.AccountId = ""
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject a malformed account_id", func() {
			spec := baseSpec()
			spec.AccountId = "not-hex"
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject key rotation on a zone-scoped organization -- the cadence is account-level at Cloudflare", func() {
			spec := baseSpec()
			spec.AccountId = ""
			spec.ZoneId = literal("9a7806061c88ada191ed06f989cc3dac")
			spec.KeyRotationIntervalDays = proto.Int32(30)
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject a rotation cadence below 21 days", func() {
			spec := baseSpec()
			spec.KeyRotationIntervalDays = proto.Int32(20)
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject a rotation cadence above 365 days", func() {
			spec := baseSpec()
			spec.KeyRotationIntervalDays = proto.Int32(366)
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject an unknown authenticator", func() {
			spec := baseSpec()
			spec.MfaConfig = &CloudflareZeroTrustOrganizationMfaConfig{
				AllowedAuthenticators: []string{"sms"},
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject an unknown pin_policy", func() {
			spec := baseSpec()
			spec.MfaPivKeyRequirements = &CloudflareZeroTrustOrganizationMfaPivKeyRequirements{
				PinPolicy: "sometimes",
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject an unknown touch_policy", func() {
			spec := baseSpec()
			spec.MfaPivKeyRequirements = &CloudflareZeroTrustOrganizationMfaPivKeyRequirements{
				TouchPolicy: "once",
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject an unsupported ssh key size", func() {
			spec := baseSpec()
			spec.MfaPivKeyRequirements = &CloudflareZeroTrustOrganizationMfaPivKeyRequirements{
				SshKeySize: []int64{1024},
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})

		ginkgo.It("should reject an unsupported ssh key type", func() {
			spec := baseSpec()
			spec.MfaPivKeyRequirements = &CloudflareZeroTrustOrganizationMfaPivKeyRequirements{
				SshKeyType: []string{"dsa"},
			}
			gomega.Expect(protovalidate.Validate(validOrg(spec))).NotTo(gomega.BeNil())
		})
	})
})
