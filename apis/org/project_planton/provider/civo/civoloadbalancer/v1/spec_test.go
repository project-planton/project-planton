package civoloadbalancerv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	civo "github.com/project-planton/project-planton/apis/org/project_planton/provider/civo"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project_planton/shared/foreignkey/v1"
)

func TestCivoLoadBalancerSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CivoLoadBalancerSpec Validation Suite")
}

var _ = ginkgo.Describe("CivoLoadBalancerSpec validations", func() {

	// Helper function to create a minimal valid spec
	makeValidSpec := func() *CivoLoadBalancerSpec {
		return &CivoLoadBalancerSpec{
			LoadBalancerName: "test-lb",
			Region:           civo.CivoRegion_lon1,
			Network: &foreignkeyv1.StringValueOrRef{
				LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
					Value: "net-12345",
				},
			},
			ForwardingRules: []*CivoLoadBalancerForwardingRule{
				{
					EntryPort:      80,
					EntryProtocol:  CivoLoadBalancerProtocol_http,
					TargetPort:     80,
					TargetProtocol: CivoLoadBalancerProtocol_http,
				},
			},
			InstanceIds: []*foreignkeyv1.StringValueOrRef{
				{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "inst-12345",
					},
				},
			},
		}
	}

	ginkgo.Context("Required fields", func() {
		ginkgo.It("accepts a minimal valid spec", func() {
			spec := makeValidSpec()
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("rejects spec with missing load_balancer_name", func() {
			spec := makeValidSpec()
			spec.LoadBalancerName = ""
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("rejects spec with missing region", func() {
			spec := makeValidSpec()
			spec.Region = civo.CivoRegion_civo_region_unspecified
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("rejects spec with missing network", func() {
			spec := makeValidSpec()
			spec.Network = nil
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("rejects spec with empty forwarding_rules", func() {
			spec := makeValidSpec()
			spec.ForwardingRules = nil
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
	})

	ginkgo.Context("load_balancer_name validation", func() {
		ginkgo.It("accepts valid names with lowercase and hyphens", func() {
			spec := makeValidSpec()
			spec.LoadBalancerName = "my-test-lb-01"
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("rejects name with uppercase letters", func() {
			spec := makeValidSpec()
			spec.LoadBalancerName = "MyTestLB"
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("rejects name with underscores", func() {
			spec := makeValidSpec()
			spec.LoadBalancerName = "my_test_lb"
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("rejects name exceeding 64 characters", func() {
			spec := makeValidSpec()
			spec.LoadBalancerName = "a-very-long-load-balancer-name-that-exceeds-the-maximum-64-char-limit-for-names"
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("accepts name at 64 character limit", func() {
			spec := makeValidSpec()
			spec.LoadBalancerName = "a-load-balancer-name-exactly-sixty-four-characters-long-test-1234"
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})

	ginkgo.Context("forwarding_rules validation", func() {
		ginkgo.It("accepts valid HTTP forwarding rule", func() {
			spec := makeValidSpec()
			spec.ForwardingRules = []*CivoLoadBalancerForwardingRule{
				{
					EntryPort:      80,
					EntryProtocol:  CivoLoadBalancerProtocol_http,
					TargetPort:     8080,
					TargetProtocol: CivoLoadBalancerProtocol_http,
				},
			}
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("accepts valid HTTPS forwarding rule", func() {
			spec := makeValidSpec()
			spec.ForwardingRules = []*CivoLoadBalancerForwardingRule{
				{
					EntryPort:      443,
					EntryProtocol:  CivoLoadBalancerProtocol_https,
					TargetPort:     443,
					TargetProtocol: CivoLoadBalancerProtocol_https,
				},
			}
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("accepts valid TCP forwarding rule", func() {
			spec := makeValidSpec()
			spec.ForwardingRules = []*CivoLoadBalancerForwardingRule{
				{
					EntryPort:      3306,
					EntryProtocol:  CivoLoadBalancerProtocol_tcp,
					TargetPort:     3306,
					TargetProtocol: CivoLoadBalancerProtocol_tcp,
				},
			}
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("accepts multiple forwarding rules", func() {
			spec := makeValidSpec()
			spec.ForwardingRules = []*CivoLoadBalancerForwardingRule{
				{
					EntryPort:      80,
					EntryProtocol:  CivoLoadBalancerProtocol_http,
					TargetPort:     8080,
					TargetProtocol: CivoLoadBalancerProtocol_http,
				},
				{
					EntryPort:      443,
					EntryProtocol:  CivoLoadBalancerProtocol_https,
					TargetPort:     8443,
					TargetProtocol: CivoLoadBalancerProtocol_https,
				},
			}
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("rejects port below valid range", func() {
			spec := makeValidSpec()
			spec.ForwardingRules[0].EntryPort = 0
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("rejects port above valid range", func() {
			spec := makeValidSpec()
			spec.ForwardingRules[0].TargetPort = 65536
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("rejects unspecified protocol", func() {
			spec := makeValidSpec()
			spec.ForwardingRules[0].EntryProtocol = CivoLoadBalancerProtocol_civo_load_balancer_protocol_unspecified
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
	})

	ginkgo.Context("health_check validation", func() {
		ginkgo.It("accepts valid HTTP health check with path", func() {
			spec := makeValidSpec()
			spec.HealthCheck = &CivoLoadBalancerHealthCheck{
				Port:     80,
				Protocol: CivoLoadBalancerProtocol_http,
				Path:     "/health",
			}
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("accepts valid TCP health check without path", func() {
			spec := makeValidSpec()
			spec.HealthCheck = &CivoLoadBalancerHealthCheck{
				Port:     3306,
				Protocol: CivoLoadBalancerProtocol_tcp,
			}
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("rejects health check with invalid port", func() {
			spec := makeValidSpec()
			spec.HealthCheck = &CivoLoadBalancerHealthCheck{
				Port:     0,
				Protocol: CivoLoadBalancerProtocol_http,
			}
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("accepts spec without health check (optional)", func() {
			spec := makeValidSpec()
			spec.HealthCheck = nil
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})

	ginkgo.Context("instance selection (mutually exclusive)", func() {
		ginkgo.It("accepts spec with instance_ids", func() {
			spec := makeValidSpec()
			spec.InstanceIds = []*foreignkeyv1.StringValueOrRef{
				{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "inst-1",
					},
				},
				{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "inst-2",
					},
				},
			}
			spec.InstanceTag = ""
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("accepts spec with instance_tag", func() {
			spec := makeValidSpec()
			spec.InstanceIds = nil
			spec.InstanceTag = "web-server"
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("accepts valid tag with alphanumeric and hyphens", func() {
			spec := makeValidSpec()
			spec.InstanceIds = nil
			spec.InstanceTag = "web-server-v1"
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("rejects tag exceeding 255 characters", func() {
			spec := makeValidSpec()
			spec.InstanceIds = nil
			spec.InstanceTag = "a-very-long-tag-name-that-exceeds-the-maximum-allowed-length-of-255-characters-for-instance-tags-this-should-fail-validation-because-its-way-too-long-and-would-cause-issues-with-the-civo-api-and-database-storage-constraints-we-need-to-keep-tags-reasonable-in-length"
			err := protovalidate.Validate(spec)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
	})

	ginkgo.Context("reserved IP configuration", func() {
		ginkgo.It("accepts spec with reserved IP", func() {
			spec := makeValidSpec()
			spec.ReservedIpId = &foreignkeyv1.StringValueOrRef{
				LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
					Value: "ip-12345",
				},
			}
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("accepts spec without reserved IP (optional)", func() {
			spec := makeValidSpec()
			spec.ReservedIpId = nil
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})

	ginkgo.Context("sticky sessions configuration", func() {
		ginkgo.It("accepts spec with sticky sessions enabled", func() {
			spec := makeValidSpec()
			spec.EnableStickySessions = true
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("accepts spec with sticky sessions disabled", func() {
			spec := makeValidSpec()
			spec.EnableStickySessions = false
			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})

	ginkgo.Context("complete production configuration", func() {
		ginkgo.It("accepts full production spec with all options", func() {
			spec := &CivoLoadBalancerSpec{
				LoadBalancerName: "prod-lb",
				Region:           civo.CivoRegion_lon1,
				Network: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "net-prod-12345",
					},
				},
				ForwardingRules: []*CivoLoadBalancerForwardingRule{
					{
						EntryPort:      80,
						EntryProtocol:  CivoLoadBalancerProtocol_http,
						TargetPort:     8080,
						TargetProtocol: CivoLoadBalancerProtocol_http,
					},
					{
						EntryPort:      443,
						EntryProtocol:  CivoLoadBalancerProtocol_https,
						TargetPort:     8443,
						TargetProtocol: CivoLoadBalancerProtocol_https,
					},
				},
				HealthCheck: &CivoLoadBalancerHealthCheck{
					Port:     8080,
					Protocol: CivoLoadBalancerProtocol_http,
					Path:     "/healthz",
				},
				InstanceIds: []*foreignkeyv1.StringValueOrRef{
					{
						LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
							Value: "inst-web-1",
						},
					},
					{
						LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
							Value: "inst-web-2",
						},
					},
				},
				ReservedIpId: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "ip-prod-12345",
					},
				},
				EnableStickySessions: false,
			}

			err := protovalidate.Validate(spec)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})
})

