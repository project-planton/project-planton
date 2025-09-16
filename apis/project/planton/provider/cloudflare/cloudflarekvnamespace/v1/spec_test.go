package cloudflarekvnamespacev1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareKvNamespaceSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CloudflareKvNamespaceSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("CloudflareKvNamespaceSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("cloudflare_kv_namespace", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareKvNamespace{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareKvNamespace",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-kv-namespace",
					},
					Spec: &CloudflareKvNamespaceSpec{
						NamespaceName: "test-namespace",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
