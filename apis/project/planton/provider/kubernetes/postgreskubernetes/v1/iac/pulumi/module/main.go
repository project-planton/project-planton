package module

import (
	"github.com/pkg/errors"
	postgreskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/postgreskubernetes/v1"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	zalandov1 "github.com/project-planton/project-planton/pkg/kubernetestypes/zalandooperator/kubernetes/acid/v1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *postgreskubernetesv1.PostgresKubernetesStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.KubernetesCluster, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	//create namespace resource
	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		locals.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name:   pulumi.String(locals.Namespace),
				Labels: pulumi.ToStringMap(locals.Labels),
			}),
		}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "5s", Update: "5s", Delete: "5s"}),
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	//create zalando postgresql resource
	_, err = zalandov1.NewPostgresql(ctx,
		"database",
		&zalandov1.PostgresqlArgs{
			Metadata: metav1.ObjectMetaArgs{
				// for zolando operator the name is required to be always prefixed by teamId
				// a kubernetes service with the same name is created by the operator
				Name:      pulumi.Sprintf("%s-%s", vars.TeamId, locals.PostgresKubernetes.Metadata.Id),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: zalandov1.PostgresqlSpecArgs{
				NumberOfInstances: pulumi.Int(locals.PostgresKubernetes.Spec.Container.Replicas),
				Patroni:           zalandov1.PostgresqlSpecPatroniArgs{},
				PodAnnotations: pulumi.ToStringMap(map[string]string{
					"postgres-cluster-id": locals.PostgresKubernetes.Metadata.Id,
				}),
				Postgresql: zalandov1.PostgresqlSpecPostgresqlArgs{
					Version: pulumi.String(vars.PostgresVersion),
					Parameters: pulumi.StringMap{
						"max_connections": pulumi.String("200"),
					},
				},
				Resources: zalandov1.PostgresqlSpecResourcesArgs{
					Limits: zalandov1.PostgresqlSpecResourcesLimitsArgs{
						Cpu:    pulumi.String(locals.PostgresKubernetes.Spec.Container.Resources.Limits.Cpu),
						Memory: pulumi.String(locals.PostgresKubernetes.Spec.Container.Resources.Limits.Memory),
					},
					Requests: zalandov1.PostgresqlSpecResourcesRequestsArgs{
						Cpu:    pulumi.String(locals.PostgresKubernetes.Spec.Container.Resources.Requests.Cpu),
						Memory: pulumi.String(locals.PostgresKubernetes.Spec.Container.Resources.Requests.Memory),
					},
				},
				TeamId: pulumi.String(vars.TeamId),
				Volume: zalandov1.PostgresqlSpecVolumeArgs{
					Size: pulumi.String(locals.PostgresKubernetes.Spec.Container.DiskSize),
				},
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create postgresql")
	}

	if locals.PostgresKubernetes.Spec.Ingress == nil ||
		!locals.PostgresKubernetes.Spec.Ingress.IsEnabled ||
		locals.PostgresKubernetes.Spec.Ingress.DnsDomain == "" {
		//if ingress is not enabled, no load-balancer resource is required. so just exit the function.
		return nil
	}

	//create kubernetes-service of type load-balancer(external)
	//this load-balancer can be used by postgres clients outside the kubernetes cluster.
	_, err = kubernetescorev1.NewService(ctx,
		"ingress-external-lb",
		&kubernetescorev1.ServiceArgs{
			Metadata: &kubernetesmetav1.ObjectMetaArgs{
				Name:      pulumi.String("ingress-external-lb"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    createdNamespace.Metadata.Labels(),
				Annotations: pulumi.StringMap{
					"external-dns.alpha.kubernetes.io/hostname": pulumi.String(locals.IngressExternalHostname),
				},
			},
			Spec: &kubernetescorev1.ServiceSpecArgs{
				Type: pulumi.String("LoadBalancer"),
				Ports: kubernetescorev1.ServicePortArray{
					&kubernetescorev1.ServicePortArgs{
						Name:       pulumi.String("postgres"),
						Protocol:   pulumi.String("TCP"),
						Port:       pulumi.Int(5432),
						TargetPort: pulumi.Int(5432),
					},
				},
				Selector: pulumi.ToStringMap(locals.PostgresPodSectorLabels),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrapf(err, "failed to create external load balancer service")
	}

	//create kubernetes-service of type load-balancer(internal)
	//this load-balancer can be used by postgres clients outside the kubernetes cluster
	//but are running in the same vpc network ex: a google cloud function/aws fargate service deployed in the same vpc.
	_, err = kubernetescorev1.NewService(ctx,
		"ingress-internal-lb",
		&kubernetescorev1.ServiceArgs{
			Metadata: &kubernetesmetav1.ObjectMetaArgs{
				Name:      pulumi.String("ingress-internal-lb"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    createdNamespace.Metadata.Labels(),
				Annotations: pulumi.StringMap{
					"cloud.google.com/load-balancer-type":       pulumi.String("Internal"),
					"external-dns.alpha.kubernetes.io/hostname": pulumi.String(locals.IngressInternalHostname),
				},
			},
			Spec: &kubernetescorev1.ServiceSpecArgs{
				Type: pulumi.String("LoadBalancer"),
				Ports: kubernetescorev1.ServicePortArray{
					&kubernetescorev1.ServicePortArgs{
						Name:       pulumi.String("postgres"),
						Protocol:   pulumi.String("TCP"),
						Port:       pulumi.Int(5432),
						TargetPort: pulumi.Int(5432),
					},
				},
				Selector: pulumi.ToStringMap(locals.PostgresPodSectorLabels),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrapf(err, "failed to create external load balancer service")
	}

	return nil
}