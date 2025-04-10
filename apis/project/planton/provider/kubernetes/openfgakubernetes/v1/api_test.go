package openfgakubernetesv1

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestOpenFgaKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OpenFgaKubernetesSpec Suite")
}

var _ = Describe("OpenFgaKubernetesSpec", func() {
	Context("with a fully valid spec", func() {
		It("should pass validation without errors", func() {
			spec := &OpenFgaKubernetesSpec{
				Container: &OpenFgaKubernetesContainer{
					Replicas: 1,
					Resources: &kubernetes.ContainerResources{
						Limits: &kubernetes.CpuMemory{
							Cpu:    "1000m",
							Memory: "1Gi",
						},
						Requests: &kubernetes.CpuMemory{
							Cpu:    "50m",
							Memory: "100Mi",
						},
					},
				},
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "openfga.example.com",
				},
				Datastore: &OpenFgaKubernetesDataStore{
					Engine: "postgres",
					Uri:    "postgres://user:pass@localhost:5432/mydb",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors")
		})
	})

	Context("when the datastore engine is invalid", func() {
		It("should fail validation, expecting an error about allowed engine values", func() {
			spec := &OpenFgaKubernetesSpec{
				Container: &OpenFgaKubernetesContainer{
					Replicas: 1,
					Resources: &kubernetes.ContainerResources{
						Limits: &kubernetes.CpuMemory{
							Cpu:    "1000m",
							Memory: "1Gi",
						},
						Requests: &kubernetes.CpuMemory{
							Cpu:    "50m",
							Memory: "100Mi",
						},
					},
				},
				Datastore: &OpenFgaKubernetesDataStore{
					Engine: "sqlite", // Invalid
					Uri:    "sqlite://path/to/db",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for invalid engine")
			Expect(strings.Contains(err.Error(), "The datastore engine must be one of \"postgres\" and \"mysql\".")).
				To(BeTrue(), "expected error about allowed engine values")
		})
	})

	Context("when ingress is omitted", func() {
		It("should pass validation if ingress is optional", func() {
			spec := &OpenFgaKubernetesSpec{
				Container: &OpenFgaKubernetesContainer{
					Replicas: 1,
					Resources: &kubernetes.ContainerResources{
						Limits: &kubernetes.CpuMemory{
							Cpu:    "1000m",
							Memory: "1Gi",
						},
						Requests: &kubernetes.CpuMemory{
							Cpu:    "50m",
							Memory: "100Mi",
						},
					},
				},
				Datastore: &OpenFgaKubernetesDataStore{
					Engine: "postgres",
					Uri:    "postgres://user:pass@localhost:5432/mydb",
				},
				// No ingress provided
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors for omitted optional ingress")
		})
	})
})
