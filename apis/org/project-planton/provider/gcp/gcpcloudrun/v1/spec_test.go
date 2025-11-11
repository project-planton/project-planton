package gcpcloudrunv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestGcpCloudRunSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpCloudRunSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpCloudRunSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_cloud_run", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpCloudRun{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudRun",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-cloud-run",
					},
					Spec: &GcpCloudRunSpec{
						ProjectId: "test-project-123",
						Region:    "us-central1",
						Container: &GcpCloudRunContainer{
							Image: &GcpCloudRunContainerImage{
								Repo: "us-docker.pkg.dev/prj/registry/app",
								Tag:  "1.0.0",
							},
							Port:   8080, // Default port for Cloud Run
							Cpu:    1,
							Memory: 512,
							Replicas: &GcpCloudRunContainerReplicas{
								Min: 0,
								Max: 10,
							},
						},
						MaxConcurrency: 80, // Default max concurrency
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
