package module

import (
	"github.com/pkg/errors"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func service(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace, createdDeployment *appsv1.Deployment) error {

	portsArray := make(kubernetescorev1.ServicePortArray, 0)
	for _, p := range locals.MicroserviceKubernetes.Spec.Container.App.Ports {
		portsArray = append(portsArray, &kubernetescorev1.ServicePortArgs{
			Name:        pulumi.String(p.Name),
			Protocol:    pulumi.String(p.NetworkProtocol),
			Port:        pulumi.Int(p.ServicePort),
			TargetPort:  pulumi.Int(p.ContainerPort),
			AppProtocol: pulumi.String(p.AppProtocol),
		})
	}

	_, err := kubernetescorev1.NewService(ctx,
		locals.MicroserviceKubernetes.Spec.Version,
		&kubernetescorev1.ServiceArgs{
			Metadata: kubernetesmetav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.MicroserviceKubernetes.Spec.Version),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: &kubernetescorev1.ServiceSpecArgs{
				Type:     pulumi.String("ClusterIP"),
				Selector: pulumi.ToStringMap(locals.Labels),
				Ports:    portsArray,
			},
		}, pulumi.Parent(createdNamespace), pulumi.DependsOn([]pulumi.Resource{createdDeployment}))
	if err != nil {
		return errors.Wrap(err, "failed to add service")
	}
	return nil
}
