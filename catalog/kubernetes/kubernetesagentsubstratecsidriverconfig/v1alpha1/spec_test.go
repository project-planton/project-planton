package kubernetesagentsubstratecsidriverconfigv1alpha1

import (
	"errors"
	"strings"
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/plantonhq/planton/shared"
)

// The validation suite pins what the spec refuses before the cluster sees the
// CSIDriverConfig: a name equal to the driver's name (Substrate finds a
// driver's config by its object name), upstream's driver-name grammar, a
// tcp:// or dns:// Controller endpoint, a unix:// node socket, and TLS that
// is either off or pod identity with a server name.

func TestKubernetesAgentSubstrateCsiDriverConfig(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesAgentSubstrateCsiDriverConfig Suite")
}

func ptr[T any](v T) *T { return &v }

func expectViolation(err error, ruleID string) {
	gomega.Expect(err).NotTo(gomega.BeNil())
	var validationErr *protovalidate.ValidationError
	gomega.Expect(errors.As(err, &validationErr)).To(gomega.BeTrue(), "not a validation error: %v", err)
	var ids []string
	for _, violation := range validationErr.Violations {
		ids = append(ids, violation.Proto.GetRuleId())
	}
	gomega.Expect(ids).To(gomega.ContainElement(ruleID), "violations: %v", err)
}

var _ = ginkgo.Describe("KubernetesAgentSubstrateCsiDriverConfig Validation Tests", func() {
	var input *KubernetesAgentSubstrateCsiDriverConfig

	ginkgo.BeforeEach(func() {
		input = &KubernetesAgentSubstrateCsiDriverConfig{
			ApiVersion: "kubernetes.planton.dev/v1alpha1",
			Kind:       "KubernetesAgentSubstrateCsiDriverConfig",
			Metadata:   &shared.CatalogObjectMetadata{Name: "hostpath"},
			Spec: &KubernetesAgentSubstrateCsiDriverConfigSpec{
				Name:               "hostpath.csi.k8s.io",
				DriverName:         "hostpath.csi.k8s.io",
				ControllerEndpoint: "tcp://csi-hostpath-controller.default.svc.cluster.local:50051",
				NodeSocketOverride: "unix:///var/lib/kubelet/plugins/csi-hostpath/csi.sock",
				Tls: &KubernetesAgentSubstrateCsiDriverConfigTls{
					Enabled:        ptr(true),
					UsePodIdentity: true,
					ServerName:     "csi-hostpath-controller.default.svc",
				},
			},
		}
	})

	ginkgo.It("accepts a config a person would write", func() {
		gomega.Expect(protovalidate.Validate(input)).To(gomega.BeNil())
	})

	ginkgo.It("accepts a plaintext config with a dns endpoint and the kubelet's socket", func() {
		input.Spec.ControllerEndpoint = "dns:///csi-nfs-controller.kube-system.svc:50052"
		input.Spec.NodeSocketOverride = ""
		input.Spec.Tls = &KubernetesAgentSubstrateCsiDriverConfigTls{Enabled: ptr(false)}
		gomega.Expect(protovalidate.Validate(input)).To(gomega.BeNil())
		input.Spec.Tls = nil
		gomega.Expect(protovalidate.Validate(input)).To(gomega.BeNil())
	})

	ginkgo.It("refuses a name that differs from the driver's name", func() {
		input.Spec.Name = "hostpath"
		expectViolation(protovalidate.Validate(input), "spec.name.is_driver_name")
	})

	ginkgo.It("refuses a driver name upstream's grammar refuses", func() {
		input.Spec.DriverName = "Hostpath_CSI"
		input.Spec.Name = "hostpath-csi"
		expectViolation(protovalidate.Validate(input), "string.pattern")
	})

	ginkgo.It("bounds the driver's name to 63 characters", func() {
		long := strings.Repeat("d", 64)
		input.Spec.Name, input.Spec.DriverName = long, long
		expectViolation(protovalidate.Validate(input), "string.max_len")
	})

	ginkgo.It("refuses a name that is not a DNS subdomain", func() {
		input.Spec.Name, input.Spec.DriverName = "-csi.example.com", "-csi.example.com"
		expectViolation(protovalidate.Validate(input), "spec.name.dns_subdomain")
	})

	ginkgo.It("refuses label keys and values outside Kubernetes' grammar", func() {
		input.Spec.Labels = map[string]string{"storage/": "nfs"}
		expectViolation(protovalidate.Validate(input), "string.pattern")
		input.Spec.Labels = map[string]string{"storage": "n f s"}
		expectViolation(protovalidate.Validate(input), "string.pattern")
	})

	ginkgo.It("refuses a substrate.io/ driver, whose config no object name could carry", func() {
		input.Spec.DriverName = "substrate.io/hostpath"
		expectViolation(protovalidate.Validate(input), "spec.name.is_driver_name")
	})

	ginkgo.It("requires a tcp:// or dns:// Controller endpoint", func() {
		input.Spec.ControllerEndpoint = ""
		expectViolation(protovalidate.Validate(input), "required")
		input.Spec.ControllerEndpoint = "https://csi:443"
		expectViolation(protovalidate.Validate(input), "string.pattern")
	})

	ginkgo.It("refuses a node socket that is not unix://", func() {
		input.Spec.NodeSocketOverride = "/var/lib/kubelet/plugins/x/csi.sock"
		expectViolation(protovalidate.Validate(input), "spec.node_socket_override.unix")
	})

	ginkgo.It("requires TLS's enabled flag, so false is written", func() {
		input.Spec.Tls.Enabled = nil
		expectViolation(protovalidate.Validate(input), "required")
	})

	ginkgo.It("requires pod identity with TLS", func() {
		input.Spec.Tls.UsePodIdentity = false
		expectViolation(protovalidate.Validate(input), "tls.enabled_needs_pod_identity")
	})

	ginkgo.It("requires a server name with TLS", func() {
		input.Spec.Tls.ServerName = ""
		expectViolation(protovalidate.Validate(input), "tls.enabled_needs_server_name")
	})
})
