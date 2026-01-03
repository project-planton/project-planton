package kubernetesdeploymentv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/plantonhq/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared"
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	foreignkeyv1 "github.com/plantonhq/project-planton/apis/org/project_planton/shared/foreignkey/v1"
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

	ginkgo.Describe("Environment secrets validation", func() {
		ginkgo.Context("When secrets have direct string values", func() {
			ginkgo.It("should pass validation", func() {
				input.Spec.Container.App.Env = &KubernetesDeploymentContainerAppEnv{
					Secrets: map[string]*kubernetes.KubernetesSensitiveValue{
						"DATABASE_PASSWORD": {
							SensitiveValue: &kubernetes.KubernetesSensitiveValue_Value{
								Value: "my-password",
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("When secrets have Kubernetes Secret references", func() {
			ginkgo.It("should pass validation with valid secret ref", func() {
				input.Spec.Container.App.Env = &KubernetesDeploymentContainerAppEnv{
					Secrets: map[string]*kubernetes.KubernetesSensitiveValue{
						"DATABASE_PASSWORD": {
							SensitiveValue: &kubernetes.KubernetesSensitiveValue_SecretRef{
								SecretRef: &kubernetes.KubernetesSecretKeyRef{
									Name: "my-app-secrets",
									Key:  "db-password",
								},
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("When secrets have mixed types", func() {
			ginkgo.It("should pass validation with both string values and secret refs", func() {
				input.Spec.Container.App.Env = &KubernetesDeploymentContainerAppEnv{
					Variables: map[string]string{
						"DATABASE_NAME": "mydb",
					},
					Secrets: map[string]*kubernetes.KubernetesSensitiveValue{
						"DATABASE_PASSWORD": {
							SensitiveValue: &kubernetes.KubernetesSensitiveValue_Value{
								Value: "my-password",
							},
						},
						"API_KEY": {
							SensitiveValue: &kubernetes.KubernetesSensitiveValue_SecretRef{
								SecretRef: &kubernetes.KubernetesSecretKeyRef{
									Name: "external-secrets",
									Key:  "api-key",
								},
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("When secret ref is missing required fields", func() {
			ginkgo.It("should fail validation when name is missing", func() {
				input.Spec.Container.App.Env = &KubernetesDeploymentContainerAppEnv{
					Secrets: map[string]*kubernetes.KubernetesSensitiveValue{
						"DATABASE_PASSWORD": {
							SensitiveValue: &kubernetes.KubernetesSensitiveValue_SecretRef{
								SecretRef: &kubernetes.KubernetesSecretKeyRef{
									Name: "",
									Key:  "db-password",
								},
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
			})

			ginkgo.It("should fail validation when key is missing", func() {
				input.Spec.Container.App.Env = &KubernetesDeploymentContainerAppEnv{
					Secrets: map[string]*kubernetes.KubernetesSensitiveValue{
						"DATABASE_PASSWORD": {
							SensitiveValue: &kubernetes.KubernetesSensitiveValue_SecretRef{
								SecretRef: &kubernetes.KubernetesSecretKeyRef{
									Name: "my-secrets",
									Key:  "",
								},
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
})
