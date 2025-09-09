package cloudflareloadbalancerv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareLoadBalancerSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudflareLoadBalancerSpec Custom Validation Tests")
}

var _ = Describe("CloudflareLoadBalancerSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("cloudflare_load_balancer", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareLoadBalancer{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareLoadBalancer",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-load-balancer",
					},
					Spec: &CloudflareLoadBalancerSpec{
						Hostname: "lb.example.com",
						ZoneId: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-zone-123"},
						},
						Origins: []*CloudflareLoadBalancerOrigin{
							{
								Name:    "origin-1",
								Address: "192.168.1.1",
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
