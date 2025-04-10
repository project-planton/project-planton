package gcpdnszonev1

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/networking/enums/dnsrecordtype"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpDnsZone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpDnsZone Suite")
}

var _ = Describe("GcpDnsZone Custom Validation Tests", func() {
	var input *GcpDnsZone

	BeforeEach(func() {
		input = &GcpDnsZone{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpDnsZone",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-zone",
			},
			Spec: &GcpDnsZoneSpec{
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
		Context("gcp_dns_zone", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
