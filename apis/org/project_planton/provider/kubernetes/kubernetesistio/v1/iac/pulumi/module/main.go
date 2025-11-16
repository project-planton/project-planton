package module

import (
	"github.com/pkg/errors"
	kubernetesistiov1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesistio/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources wires up Istio (base, control‑plane & ingress gateway) on the target
// Kubernetes cluster.  It honours spec.container.resources for the Istiod
// deployment, allowing callers to tune CPU / memory without touching Helm YAML.
func Resources(ctx *pulumi.Context, in *kubernetesistiov1.KubernetesIstioStackInput) error {
	// Kubernetes provider from cluster‑credential
	kubernetesProviderConfig, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, in.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	spec := in.Target.Spec

	// ---- pick chart version ----
	// (No release_channel field yet – stick to default stable.)
	chartVersion := vars.DefaultStableVersion

	// ---- namespaces ----
	sysNS, err := corev1.NewNamespace(ctx, vars.SystemNamespace,
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String(vars.SystemNamespace),
			},
		},
		pulumi.Provider(kubernetesProviderConfig))
	if err != nil {
		return errors.Wrap(err, "failed to create istio-system namespace")
	}

	gwNS, err := corev1.NewNamespace(ctx, vars.GatewayNamespace,
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String(vars.GatewayNamespace),
			},
		},
		pulumi.Provider(kubernetesProviderConfig))
	if err != nil {
		return errors.Wrap(err, "failed to create istio-ingress namespace")
	}

	// convenience for repeated repo opts
	repo := helm.RepositoryOptsArgs{Repo: pulumi.String(vars.HelmRepo)}

	// ---- istio/base ----
	_, err = helm.NewRelease(ctx, "istio-base",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.BaseChart),
			Namespace:       sysNS.Metadata.Name(),
			Chart:           pulumi.String(vars.BaseChart),
			Version:         pulumi.String(chartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			RepositoryOpts:  repo,
		},
		pulumi.Provider(kubernetesProviderConfig),
		pulumi.Parent(sysNS))
	if err != nil {
		return errors.Wrap(err, "installing istio/base")
	}

	// ---- istiod (control‑plane) ----
	istiodValues := pulumi.Map{}
	if res := spec.GetContainer(); res != nil && res.Resources != nil {
		// map protobuf fields -> Helm values: pilot.resources
		limits := res.Resources.GetLimits()
		requests := res.Resources.GetRequests()
		istiodValues["pilot"] = pulumi.Map{
			"resources": pulumi.Map{
				"limits": pulumi.Map{
					"cpu":    pulumi.String(limits.GetCpu()),
					"memory": pulumi.String(limits.GetMemory()),
				},
				"requests": pulumi.Map{
					"cpu":    pulumi.String(requests.GetCpu()),
					"memory": pulumi.String(requests.GetMemory()),
				},
			},
		}
	}

	_, err = helm.NewRelease(ctx, "istiod",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.IstiodChart),
			Namespace:       sysNS.Metadata.Name(),
			Chart:           pulumi.String(vars.IstiodChart),
			Version:         pulumi.String(chartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values:          istiodValues,
			RepositoryOpts:  repo,
		},
		pulumi.Provider(kubernetesProviderConfig),
		pulumi.Parent(sysNS))
	if err != nil {
		return errors.Wrap(err, "installing istiod control‑plane")
	}

	// ---- ingress‑gateway ----
	_, err = helm.NewRelease(ctx, "istio-gateway",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.GatewayChart),
			Namespace:       gwNS.Metadata.Name(),
			Chart:           pulumi.String(vars.GatewayChart),
			Version:         pulumi.String(chartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			RepositoryOpts:  repo,
			Values: pulumi.Map{
				"service": pulumi.Map{
					"type": pulumi.String("ClusterIP"),
				},
			},
		},
		pulumi.Provider(kubernetesProviderConfig),
		pulumi.Parent(gwNS))
	if err != nil {
		return errors.Wrap(err, "installing istio ingress‑gateway")
	}

	// ---- stack outputs ----
	ctx.Export(OpNamespace, sysNS.Metadata.Name())
	ctx.Export(OpService, pulumi.String("istiod"))
	ctx.Export(OpPortForwardCommand, pulumi.Sprintf("kubectl port-forward -n %s svc/istiod 15014:15014", sysNS.Metadata.Name()))
	ctx.Export(OpKubeEndpoint, pulumi.Sprintf("istiod.%s.svc.cluster.local:15012", sysNS.Metadata.Name()))
	ctx.Export(OpIngressEndpoint, pulumi.Sprintf("istio-gateway.%s.svc.cluster.local:80", gwNS.Metadata.Name()))

	return nil
}
