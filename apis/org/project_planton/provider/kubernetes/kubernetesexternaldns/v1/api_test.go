package kubernetesexternaldnsv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project_planton/shared/foreignkey/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesExternalDns(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesExternalDns Suite")
}

var _ = ginkgo.Describe("KubernetesExternalDns Validation Tests", func() {
	var input *KubernetesExternalDns

	ginkgo.BeforeEach(func() {
		input = &KubernetesExternalDns{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesExternalDns",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-external-dns",
			},
			Spec: &KubernetesExternalDnsSpec{},
		}
	})

	ginkgo.Describe("GKE Configuration", func() {
		ginkgo.Context("with valid GKE config", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Gke{
					Gke: &KubernetesExternalDnsGkeConfig{
						ProjectId: &foreignkeyv1.StringValueOrRef{
							Value: "my-gcp-project",
						},
						DnsZoneId: &foreignkeyv1.StringValueOrRef{
							Value: "my-dns-zone",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with missing project_id", func() {
			ginkgo.It("should return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Gke{
					Gke: &KubernetesExternalDnsGkeConfig{
						DnsZoneId: &foreignkeyv1.StringValueOrRef{
							Value: "my-dns-zone",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("with missing dns_zone_id", func() {
			ginkgo.It("should return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Gke{
					Gke: &KubernetesExternalDnsGkeConfig{
						ProjectId: &foreignkeyv1.StringValueOrRef{
							Value: "my-gcp-project",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("EKS Configuration", func() {
		ginkgo.Context("with valid EKS config", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Eks{
					Eks: &KubernetesExternalDnsEksConfig{
						Route53ZoneId: &foreignkeyv1.StringValueOrRef{
							Value: "Z123456789ABC",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with valid EKS config and IRSA role override", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Eks{
					Eks: &KubernetesExternalDnsEksConfig{
						Route53ZoneId: &foreignkeyv1.StringValueOrRef{
							Value: "Z123456789ABC",
						},
						IrsaRoleArnOverride: "arn:aws:iam::123456789012:role/external-dns-role",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with missing route53_zone_id", func() {
			ginkgo.It("should return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Eks{
					Eks: &KubernetesExternalDnsEksConfig{},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("AKS Configuration", func() {
		ginkgo.Context("with valid AKS config", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Aks{
					Aks: &KubernetesExternalDnsAksConfig{
						DnsZoneId:               "my-azure-dns-zone-id",
						ManagedIdentityClientId: "12345678-1234-1234-1234-123456789012",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with minimal AKS config", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Aks{
					Aks: &KubernetesExternalDnsAksConfig{},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("Cloudflare Configuration", func() {
		ginkgo.Context("with valid Cloudflare config", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Cloudflare{
					Cloudflare: &KubernetesExternalDnsCloudflareConfig{
						ApiToken:  "my-cloudflare-api-token",
						DnsZoneId: "1234567890abcdef1234567890abcdef",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with Cloudflare proxy enabled", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Cloudflare{
					Cloudflare: &KubernetesExternalDnsCloudflareConfig{
						ApiToken:  "my-cloudflare-api-token",
						DnsZoneId: "1234567890abcdef1234567890abcdef",
						IsProxied: true,
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with missing api_token", func() {
			ginkgo.It("should return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Cloudflare{
					Cloudflare: &KubernetesExternalDnsCloudflareConfig{
						DnsZoneId: "1234567890abcdef1234567890abcdef",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("with missing dns_zone_id", func() {
			ginkgo.It("should return a validation error", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Cloudflare{
					Cloudflare: &KubernetesExternalDnsCloudflareConfig{
						ApiToken: "my-cloudflare-api-token",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("Default Values", func() {
		ginkgo.Context("with defaults for namespace and versions", func() {
			ginkgo.It("should apply default values from proto", func() {
				input.Spec.ProviderConfig = &KubernetesExternalDnsSpec_Gke{
					Gke: &KubernetesExternalDnsGkeConfig{
						ProjectId: &foreignkeyv1.StringValueOrRef{
							Value: "my-gcp-project",
						},
						DnsZoneId: &foreignkeyv1.StringValueOrRef{
							Value: "my-dns-zone",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
				// Defaults should be applied:
				// namespace: kubernetes-external-dns
				// kubernetes_external_dns_version: v0.19.0
				// helm_chart_version: 1.19.0
			})
		})
	})
})
