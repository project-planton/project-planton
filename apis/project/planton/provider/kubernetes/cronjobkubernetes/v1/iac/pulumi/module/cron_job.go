package module

import (
	"fmt"

	"github.com/pkg/errors"
	batchv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/batch/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/sortstringmap"
)

func cronJob(ctx *pulumi.Context, locals *Locals, createdNamespace *corev1.Namespace) (*batchv1.CronJob, error) {
	target := locals.CronJobKubernetes

	envVarInputs := make([]corev1.EnvVarInput, 0)

	envVarInputs = append(envVarInputs, corev1.EnvVarInput(corev1.EnvVarArgs{
		Name: pulumi.String("HOSTNAME"),
		ValueFrom: &corev1.EnvVarSourceArgs{
			FieldRef: &corev1.ObjectFieldSelectorArgs{
				FieldPath: pulumi.String("status.podIP"),
			},
		},
	}))

	envVarInputs = append(envVarInputs, corev1.EnvVarInput(corev1.EnvVarArgs{
		Name: pulumi.String("K8S_POD_ID"),
		ValueFrom: &corev1.EnvVarSourceArgs{
			FieldRef: &corev1.ObjectFieldSelectorArgs{
				ApiVersion: pulumi.String("v1"),
				FieldPath:  pulumi.String("metadata.name"),
			},
		},
	}))

	if target.Spec.Env != nil {
		if target.Spec.Env.Variables != nil {
			sortedVarKeys := sortstringmap.SortMap(target.Spec.Env.Variables)
			for _, key := range sortedVarKeys {
				envVarInputs = append(envVarInputs, corev1.EnvVarInput(corev1.EnvVarArgs{
					Name:  pulumi.String(key),
					Value: pulumi.String(target.Spec.Env.Variables[key]),
				}))
			}
		}

		if target.Spec.Env.Secrets != nil {
			sortedSecretKeys := sortstringmap.SortMap(target.Spec.Env.Secrets)
			for _, secretKey := range sortedSecretKeys {
				envVarInputs = append(envVarInputs, corev1.EnvVarInput(corev1.EnvVarArgs{
					Name: pulumi.String(secretKey),
					ValueFrom: &corev1.EnvVarSourceArgs{
						SecretKeyRef: &corev1.SecretKeySelectorArgs{
							Name: pulumi.String("main"),
							Key:  pulumi.String(secretKey),
						},
					},
				}))
			}
		}
	}

	mainContainer := &corev1.ContainerArgs{
		Name: pulumi.String("cronjob-container"),
		Image: pulumi.String(fmt.Sprintf("%s:%s",
			target.Spec.Image.Repo,
			target.Spec.Image.Tag)),
		Env: corev1.EnvVarArray(envVarInputs),
		Resources: corev1.ResourceRequirementsArgs{
			Limits: pulumi.ToStringMap(map[string]string{
				"cpu":    target.Spec.Resources.Limits.Cpu,
				"memory": target.Spec.Resources.Limits.Memory,
			}),
			Requests: pulumi.ToStringMap(map[string]string{
				"cpu":    target.Spec.Resources.Requests.Cpu,
				"memory": target.Spec.Resources.Requests.Memory,
			}),
		},
	}

	if len(target.Spec.Command) > 0 {
		mainContainer.Command = pulumi.ToStringArray(target.Spec.Command)
	}
	if len(target.Spec.Args) > 0 {
		mainContainer.Args = pulumi.ToStringArray(target.Spec.Args)
	}

	podSpecArgs := &corev1.PodSpecArgs{
		RestartPolicy: pulumi.String(target.Spec.RestartPolicy),
		Containers: corev1.ContainerArray{
			mainContainer,
		},
	}

	if locals.ImagePullSecretData != nil {
		podSpecArgs.ImagePullSecrets = corev1.LocalObjectReferenceArray{
			corev1.LocalObjectReferenceArgs{
				Name: pulumi.String("image-pull-secret"),
			},
		}
	}

	cronJobSpec := &batchv1.CronJobSpecArgs{
		Schedule:                   pulumi.String(target.Spec.Schedule),
		ConcurrencyPolicy:          pulumi.String(target.Spec.ConcurrencyPolicy),
		Suspend:                    pulumi.BoolPtr(target.Spec.Suspend),
		SuccessfulJobsHistoryLimit: pulumi.IntPtr(int(target.Spec.SuccessfulJobsHistoryLimit)),
		FailedJobsHistoryLimit:     pulumi.IntPtr(int(target.Spec.FailedJobsHistoryLimit)),
		JobTemplate: &batchv1.JobTemplateSpecArgs{
			Spec: &batchv1.JobSpecArgs{
				BackoffLimit: pulumi.IntPtr(int(target.Spec.BackoffLimit)),
				Template: &corev1.PodTemplateSpecArgs{
					Spec: podSpecArgs,
				},
			},
		},
	}

	if target.Spec.StartingDeadlineSeconds > 0 {
		cronJobSpec.StartingDeadlineSeconds = pulumi.IntPtr(int(target.Spec.StartingDeadlineSeconds))
	}

	createdCronJob, err := batchv1.NewCronJob(ctx,
		target.Metadata.Name,
		&batchv1.CronJobArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(target.Metadata.Name),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: cronJobSpec,
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cronjob")
	}

	return createdCronJob, nil
}
