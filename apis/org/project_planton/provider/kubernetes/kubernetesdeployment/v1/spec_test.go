package kubernetesdeploymentv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project_planton/shared/foreignkey/v1"
)

func TestKubernetesDeployment(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesDeployment Suite")
}

var _ = ginkgo.Describe("KubernetesDeployment Custom Validation Tests", func() {
	var input *KubernetesDeployment

	ginkgo.BeforeEach(func() {
		input = &KubernetesDeployment{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesDeployment",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-deployment",
			},
			Spec: &KubernetesDeploymentSpec{
				TargetCluster: &kubernetes.KubernetesClusterSelector{
					ClusterKind: cloudresourcekind.CloudResourceKind_GcpGkeCluster,
					ClusterName: "test-cluster",
				},
				Namespace: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "test-namespace",
					},
				},
				Version: "main",
				Container: &KubernetesDeploymentContainer{
					App: &KubernetesDeploymentContainerApp{
						Image: &kubernetes.ContainerImage{
							Repo: "nginx",
							Tag:  "latest",
						},
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
						Ports: []*KubernetesDeploymentContainerAppPort{
							{
								Name:            "http",
								ContainerPort:   8080,
								ServicePort:     80,
								NetworkProtocol: "TCP",
								AppProtocol:     "http",
								IsIngressPort:   true,
							},
						},
					},
				},
				Ingress: &KubernetesDeploymentIngress{
					Enabled:  true,
					Hostname: "myapp.example.com",
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("deployment_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("Ingress validation", func() {
		ginkgo.Context("When ingress is enabled without hostname", func() {
			ginkgo.It("should return a validation error", func() {
				input.Spec.Ingress.Hostname = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})

		ginkgo.Context("When ingress is disabled", func() {
			ginkgo.It("should not require hostname", func() {
				input.Spec.Ingress.Enabled = false
				input.Spec.Ingress.Hostname = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("Version validation", func() {
		ginkgo.Context("When version contains uppercase", func() {
			ginkgo.It("should return a validation error", func() {
				input.Spec.Version = "Main"
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})

		ginkgo.Context("When version ends with hyphen", func() {
			ginkgo.It("should return a validation error", func() {
				input.Spec.Version = "main-"
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("Namespace creation flag", func() {
		ginkgo.Context("When create_namespace is true", func() {
			ginkgo.It("should pass validation", func() {
				input.Spec.CreateNamespace = true
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("When create_namespace is false", func() {
			ginkgo.It("should pass validation", func() {
				input.Spec.CreateNamespace = false
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
