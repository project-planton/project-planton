package azurenatgatewayv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureNatGateway(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AzureNatGateway Suite")
}

var _ = Describe("AzureNatGateway Custom Validation Tests", func() {

	var input *AzureNatGateway

	BeforeEach(func() {
		input = &AzureNatGateway{
			ApiVersion: "azure.project-planton.org/v1",
			Kind:       "AzureNatGateway",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-aks-cluster",
			},
			Spec: &AzureNatGatewaySpec{
				Region:        "eastus",
				ResourceGroup: "my-rg",
				SubnetId:      "subnet-1",
				NodeVmSize:    "Standard_B2s",
				MinNodeCount:  1,
				MaxNodeCount:  3,
				DnsPrefix:     "aks-valid",
				SshPublicKey:  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCvalidkey",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("azure_aks", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	Context("DNS Prefix Pattern Validation", func() {
		It("should accept a valid dns_prefix", func() {
			input.Spec.DnsPrefix = "abc-123"
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})

		It("should reject a prefix that starts with a non-letter", func() {
			input.Spec.DnsPrefix = "1abc"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should reject a prefix with invalid characters", func() {
			input.Spec.DnsPrefix = "ab@d!"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should reject a prefix that is too short", func() {
			input.Spec.DnsPrefix = "a"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should reject a prefix that exceeds the maximum length", func() {
			input.Spec.DnsPrefix = "abcdefghijklmnopqrstuvwxyzabcdef"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("SSH Public Key Pattern Validation", func() {
		It("should accept a valid ssh-rsa key", func() {
			input.Spec.SshPublicKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCUvalid"
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})

		It("should accept a valid ssh-ed25519 key", func() {
			input.Spec.SshPublicKey = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIvalid"
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})

		It("should reject a key without the proper prefix", func() {
			input.Spec.SshPublicKey = "invalid-rsa AAAAB3NzaC1yc2E"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should reject a key missing the base64 part", func() {
			input.Spec.SshPublicKey = "ssh-rsa "
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should reject a key with invalid characters in the base64 portion", func() {
			input.Spec.SshPublicKey = "ssh-rsa AAAAB3NzaC1yc2E@@invalid"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})
})
