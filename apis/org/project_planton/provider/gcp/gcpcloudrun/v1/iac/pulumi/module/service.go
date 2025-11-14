package module

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"            // provider
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/cloudrunv2" // Cloud Run v2
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// service creates a Cloud Run v2 Service and applies optional public access.
func service(
	ctx *pulumi.Context,
	locals *Locals,
	gcpProvider *gcp.Provider,
) (*cloudrunv2.Service, error) {

	image := fmt.Sprintf("%s:%s",
		locals.GcpCloudRun.Spec.Container.Image.Repo,
		locals.GcpCloudRun.Spec.Container.Image.Tag)

	memory := fmt.Sprintf("%dMi", locals.GcpCloudRun.Spec.Container.Memory)
	cpu := strconv.Itoa(int(locals.GcpCloudRun.Spec.Container.Cpu))

	port := int32(8080)
	if locals.GcpCloudRun.Spec.Container.Port != 0 {
		port = locals.GcpCloudRun.Spec.Container.Port
	}

	// Determine service name: use spec.service_name if provided, otherwise metadata.name
	serviceName := locals.GcpCloudRun.Metadata.Name
	if locals.GcpCloudRun.Spec.ServiceName != "" {
		serviceName = locals.GcpCloudRun.Spec.ServiceName
	}

	createdService, err := cloudrunv2.NewService(ctx,
		locals.GcpCloudRun.Metadata.Name,
		&cloudrunv2.ServiceArgs{
			Project:  pulumi.String(locals.GcpCloudRun.Spec.ProjectId),
			Location: pulumi.String(locals.GcpCloudRun.Spec.Region),

			Name: pulumi.String(serviceName),

			// Public access → disable Invoker IAM check
			InvokerIamDisabled: pulumi.BoolPtr(locals.GcpCloudRun.Spec.AllowUnauthenticated),

			Template: &cloudrunv2.ServiceTemplateArgs{
				Containers: cloudrunv2.ServiceTemplateContainerArray{
					&cloudrunv2.ServiceTemplateContainerArgs{
						Image: pulumi.String(image),

						Ports: &cloudrunv2.ServiceTemplateContainerPortsArgs{
							ContainerPort: pulumi.Int(port),
						},

						Resources: &cloudrunv2.ServiceTemplateContainerResourcesArgs{
							Limits: pulumi.StringMap{
								"memory": pulumi.String(memory),
								"cpu":    pulumi.String(cpu),
							},
						},

						Envs: toEnvArray(locals),
					},
				},

				// Min / Max instances map to Scaling in v2.
				Scaling: &cloudrunv2.ServiceTemplateScalingArgs{
					MinInstanceCount: pulumi.Int(int(locals.GcpCloudRun.Spec.Container.Replicas.Min)),
					MaxInstanceCount: pulumi.Int(int(locals.GcpCloudRun.Spec.Container.Replicas.Max)),
				},
			},

			// attach useful labels
			Labels: pulumi.ToStringMap(locals.GcpLabels),
		},
		pulumi.Provider(gcpProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Cloud Run v2 service")
	}

	return createdService, nil
}

// toEnvArray converts proto env / secret maps into the v2 container-env slice.
func toEnvArray(locals *Locals) cloudrunv2.ServiceTemplateContainerEnvArray {
	envs := cloudrunv2.ServiceTemplateContainerEnvArray{}

	for k, v := range locals.GcpCloudRun.Spec.Container.Env.Variables {
		envs = append(envs, &cloudrunv2.ServiceTemplateContainerEnvArgs{
			Name:  pulumi.String(k),
			Value: pulumi.String(v),
		})
	}

	for k, v := range locals.GcpCloudRun.Spec.Container.Env.Secrets {
		envs = append(envs, &cloudrunv2.ServiceTemplateContainerEnvArgs{
			Name: pulumi.String(k),
			ValueSource: &cloudrunv2.ServiceTemplateContainerEnvValueSourceArgs{
				SecretKeyRef: &cloudrunv2.ServiceTemplateContainerEnvValueSourceSecretKeyRefArgs{
					Secret: pulumi.String(v), // projects/*/secrets/*:version
				},
			},
		})
	}

	return envs
}
