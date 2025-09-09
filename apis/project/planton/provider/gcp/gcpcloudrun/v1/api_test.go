package gcpcloudrunv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCloudRun(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCloudRun Custom Validation Tests Suite")
}

var _ = Describe("GcpCloudRun Custom Validation Tests", func() {

	var input *GcpCloudRun

	BeforeEach(func() {
		input = &GcpCloudRun{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpCloudRun",
			Metadata: &shared.ApiResourceMetadata{
				Name: "demo-service",
			},
			Spec: &GcpCloudRunSpec{
				ProjectId: "myproj-1234",
				Region:    "us-central1",
				Container: &GcpCloudRunContainer{
					Image: &GcpCloudRunContainerImage{
						Repo: "us-docker.pkg.dev/myproj/registry/app",
						Tag:  "1.0.0",
					},
					Env: &GcpCloudRunContainerEnv{
						Variables: map[string]string{"ENV": "prod"},
					},
					Port:   8080,
					Cpu:    2,
					Memory: 512,
					Replicas: &GcpCloudRunContainerReplicas{
						Min: 1,
						Max: 5,
					},
				},
				MaxConcurrency:       100,
				AllowUnauthenticated: true,
				Dns: &GcpCloudRunDns{
					Enabled:     true,
					Hostnames:   []string{"app.example.com"},
					ManagedZone: "example-com",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("GCP Cloud Run", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	/* ---------- Project-ID Pattern ---------- */
	Context("ProjectId Pattern Validations", func() {
		It("should accept a valid project ID", func() {
			input.Spec.ProjectId = "validproj-1"
			Expect(protovalidate.Validate(input)).To(BeNil())
		})

		It("should reject an ID that starts with a digit", func() {
			input.Spec.ProjectId = "1invalidproj"
			Expect(protovalidate.Validate(input)).NotTo(BeNil())
		})

		It("should reject an ID that is too long", func() {
			input.Spec.ProjectId = "a123456789012345678901234567890"
			Expect(protovalidate.Validate(input)).NotTo(BeNil())
		})
	})

	/* ---------- Region Pattern ---------- */
	Context("Region Pattern Validations", func() {
		It("should accept a canonical region", func() {
			input.Spec.Region = "europe-west2"
			Expect(protovalidate.Validate(input)).To(BeNil())
		})

		It("should reject a region without dash", func() {
			input.Spec.Region = "europewest2"
			Expect(protovalidate.Validate(input)).NotTo(BeNil())
		})
	})

	/* ---------- CPU Allowed Values ---------- */
	Context("CPU Value Validations", func() {
		It("should accept an allowed CPU value (4)", func() {
			input.Spec.Container.Cpu = 4
			Expect(protovalidate.Validate(input)).To(BeNil())
		})

		It("should reject a disallowed CPU value (3)", func() {
			input.Spec.Container.Cpu = 3
			Expect(protovalidate.Validate(input)).NotTo(BeNil())
		})
	})

	/* ---------- DNS Hostname & CEL Rule ---------- */
	Context("DNS Configuration Validations", func() {
		It("should reject duplicate hostnames", func() {
			input.Spec.Dns.Hostnames = []string{"dup.example.com", "dup.example.com"}
			Expect(protovalidate.Validate(input)).NotTo(BeNil())
		})

		It("should reject an invalid hostname", func() {
			input.Spec.Dns.Hostnames = []string{"bad_hostname.example.com"}
			Expect(protovalidate.Validate(input)).NotTo(BeNil())
		})

		It("should reject when enabled=true but hostnames are empty", func() {
			input.Spec.Dns.Hostnames = []string{}
			Expect(protovalidate.Validate(input)).NotTo(BeNil())
		})

		It("should reject when enabled=true but managed_zone is empty", func() {
			input.Spec.Dns.Hostnames = []string{"app.example.com"}
			input.Spec.Dns.ManagedZone = ""
			Expect(protovalidate.Validate(input)).NotTo(BeNil())
		})

		It("should accept when enabled=false with no hostnames/zone", func() {
			input.Spec.Dns.Enabled = false
			input.Spec.Dns.Hostnames = nil
			input.Spec.Dns.ManagedZone = ""
			Expect(protovalidate.Validate(input)).To(BeNil())
		})
	})
})
