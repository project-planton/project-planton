package jenkinskubernetesv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestJenkinsKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JenkinsKubernetesSpec Suite")
}

var _ = Describe("JenkinsKubernetesSpec", func() {
	Context("with a valid spec", func() {
		It("should not return validation errors", func() {
			spec := &JenkinsKubernetesSpec{
				ContainerResources: &kubernetes.ContainerResources{
					Limits: &kubernetes.CpuMemory{
						Cpu:    "2000m",
						Memory: "2Gi",
					},
					Requests: &kubernetes.CpuMemory{
						Cpu:    "100m",
						Memory: "200Mi",
					},
				},
				HelmValues: map[string]string{
					"controller.tag": "lts",
				},
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "jenkins.example.com",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors")
		})
	})

	Context("when container_resources is not provided", func() {
		It("should still pass validation as defaults are applied", func() {
			spec := &JenkinsKubernetesSpec{
				// Not setting ContainerResources
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors with default container_resources")
		})
	})

	Context("when helm_values is empty", func() {
		It("should allow an empty helm_values map", func() {
			spec := &JenkinsKubernetesSpec{
				ContainerResources: &kubernetes.ContainerResources{
					Limits: &kubernetes.CpuMemory{
						Cpu:    "1000m",
						Memory: "1Gi",
					},
					Requests: &kubernetes.CpuMemory{
						Cpu:    "50m",
						Memory: "100Mi",
					},
				},
				HelmValues: make(map[string]string),
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors for empty helm_values")
		})
	})

	Context("when ingress is not provided", func() {
		It("should not return validation errors if ingress is optional", func() {
			spec := &JenkinsKubernetesSpec{
				ContainerResources: &kubernetes.ContainerResources{
					Limits: &kubernetes.CpuMemory{
						Cpu:    "1000m",
						Memory: "1Gi",
					},
					Requests: &kubernetes.CpuMemory{
						Cpu:    "50m",
						Memory: "100Mi",
					},
				},
				// No ingress
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors with no ingress provided")
		})
	})
})
