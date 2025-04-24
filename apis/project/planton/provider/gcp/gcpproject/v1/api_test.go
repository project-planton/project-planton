package gcpprojectv1

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/networking/enums/dnsrecordtype"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpProject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpProject Suite")
}

var _ = Describe("GcpProject Custom Validation Tests", func() {
	var input *GcpProject

	BeforeEach(func() {
		input = &GcpProject{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpProject",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-zone",
			},
			Spec: &GcpProjectSpec{
				ProjectId: "my-gcp-project",
				IamServiceAccounts: []string{
					"some-iam@my-gcp-project.iam.gserviceaccount.com",
				},
				Records: []*GcpDnsRecord{
					{
						RecordType: dnsrecordtype.DnsRecordType_A,
						Name:       "example.com.",
						Values:     []string{"1.2.3.4"},
						TtlSeconds: 60,
					},
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("gcp_project", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
