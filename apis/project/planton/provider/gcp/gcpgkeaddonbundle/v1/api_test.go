package gcpgkeaddonbundlev1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpGkeAddonBundle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpGkeAddonBundle Suite")
}

var _ = Describe("GcpGkeAddonBundle Custom Validation Tests", func() {
	var input *GcpGkeAddonBundle

	BeforeEach(func() {
		input = &GcpGkeAddonBundle{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpGkeAddonBundle",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-addon-bundle",
			},
			Spec: &GcpGkeAddonBundleSpec{
				ClusterProjectId: "my-test-project",
				Istio: &GcpGkeAddonBundleIstio{
					Enabled:            true,
					ClusterRegion:      "us-central1",
					SubNetworkSelfLink: "https://www.googleapis.com/compute/v1/projects/my-test-project/regions/us-central1/subnetworks/test-subnetwork",
				},
				InstallPostgresOperator: true,
				InstallKafkaOperator:    false,
				InstallSolrOperator:     true,
				InstallKubecost:         false,
				InstallIngressNginx:     true,
				InstallCertManager:      true,
				InstallExternalDns:      false,
				InstallExternalSecrets:  false,
				InstallElasticOperator:  true,
				InstallKeycloakOperator: false,
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("gcp_gke_addon_bundle", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
