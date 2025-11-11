package gcpgkeaddonbundlev1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestGcpGkeAddonBundleSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpGkeAddonBundleSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpGkeAddonBundleSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_gke_addon_bundle", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpGkeAddonBundle{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpGkeAddonBundle",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-gke-addon-bundle",
					},
					Spec: &GcpGkeAddonBundleSpec{
						ClusterProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
