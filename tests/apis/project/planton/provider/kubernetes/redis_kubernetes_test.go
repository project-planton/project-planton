package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/rediskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestRedisKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &rediskubernetesv1.RedisKubernetesSpec{
		Container: &rediskubernetesv1.RedisKubernetesContainer{
			Replicas:             1,
			IsPersistenceEnabled: true,
			DiskSize:             "10Gi", // valid format
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
		},
		Ingress: &kubernetes.IngressSpec{
			DnsDomain: "redis.example.com",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

func TestRedisKubernetesSpec_PersistenceEnabledNoDiskSize(t *testing.T) {
	spec := &rediskubernetesv1.RedisKubernetesSpec{
		Container: &rediskubernetesv1.RedisKubernetesContainer{
			Replicas:             1,
			IsPersistenceEnabled: true,
			// No disk_size provided
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
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected a validation error for missing disk_size when persistence is enabled, got none")
	} else {
		if !strings.Contains(err.Error(), "[spec.container.disk_size.required]") {
			t.Errorf("expected validation error with constraint id `spec.container.disk_size.required`, got: %v", err)
		}
	}
}

func TestRedisKubernetesSpec_InvalidDiskSizeFormat(t *testing.T) {
	spec := &rediskubernetesv1.RedisKubernetesSpec{
		Container: &rediskubernetesv1.RedisKubernetesContainer{
			Replicas:             1,
			IsPersistenceEnabled: true,
			DiskSize:             "abc", // invalid format
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
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected a validation error for invalid disk_size format, got none")
	} else {
		if !strings.Contains(err.Error(), "[spec.container.disk_size.required]") {
			t.Errorf("expected validation error with constraint id `spec.container.disk_size.required`, got: %v", err)
		}
	}
}

func TestRedisKubernetesSpec_PersistenceDisabledNoDiskSize(t *testing.T) {
	spec := &rediskubernetesv1.RedisKubernetesSpec{
		Container: &rediskubernetesv1.RedisKubernetesContainer{
			Replicas:             1,
			IsPersistenceEnabled: false,
			// disk_size is empty, but persistence is disabled, so no error expected
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
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("did not expect a validation error when persistence is disabled and disk_size is empty: %v", err)
	}
}
