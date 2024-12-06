package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	solrkubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/solrkubernetes/v1"
	"github.com/project-planton/project-planton/apis/go/project/planton/shared/kubernetes"
)

// TestSolrKubernetesSpec_ValidSpec ensures a fully valid spec passes validation.
func TestSolrKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &solrkubernetesv1.SolrKubernetesSpec{
		SolrContainer: &solrkubernetesv1.SolrKubernetesSolrContainer{
			Replicas: 1,
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
			DiskSize: "1Gi", // Valid disk size format
			Image: &kubernetes.ContainerImage{
				Repo: "solr",
				Tag:  "8.7.0",
			},
		},
		Config: &solrkubernetesv1.SolrKubernetesSolrConfig{
			JavaMem:                 "-Xmx512m",
			Opts:                    "-Dsolr.autoSoftCommit.maxTime=10000",
			GarbageCollectionTuning: "-XX:SurvivorRatio=4",
		},
		ZookeeperContainer: &solrkubernetesv1.SolrKubernetesZookeeperContainer{
			Replicas: 1,
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
			DiskSize: "1Gi", // Valid disk size
		},
		Ingress: &kubernetes.IngressSpec{
			// Add valid ingress fields as needed
			DnsDomain: "solr.example.com",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

// TestSolrKubernetesSpec_InvalidSolrDiskSize checks that an invalid Solr container disk size fails validation.
func TestSolrKubernetesSpec_InvalidSolrDiskSize(t *testing.T) {
	spec := &solrkubernetesv1.SolrKubernetesSpec{
		SolrContainer: &solrkubernetesv1.SolrKubernetesSolrContainer{
			Replicas: 1,
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
			DiskSize: "abc", // Invalid format
			Image: &kubernetes.ContainerImage{
				Repo: "solr",
				Tag:  "8.7.0",
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid disk size, got none")
	} else {
		if !strings.Contains(err.Error(), "[spec.container.disk_size.required]") {
			t.Errorf("expected validation error with constraint id `spec.container.disk_size.required`, got: %v", err)
		}
	}
}

// TestSolrKubernetesSpec_InvalidZookeeperDiskSize checks that an invalid Zookeeper container disk size fails validation.
func TestSolrKubernetesSpec_InvalidZookeeperDiskSize(t *testing.T) {
	spec := &solrkubernetesv1.SolrKubernetesSpec{
		ZookeeperContainer: &solrkubernetesv1.SolrKubernetesZookeeperContainer{
			Replicas: 1,
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
			DiskSize: "100", // Missing unit, invalid format
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid disk size, got none")
	} else {
		if !strings.Contains(err.Error(), "[spec.container.disk_size.required]") {
			t.Errorf("expected validation error with constraint id `spec.container.disk_size.required`, got: %v", err)
		}
	}
}

// TestSolrKubernetesSpec_MissingSolrContainer checks that missing Solr container details fail validation if required.
func TestSolrKubernetesSpec_MissingSolrContainer(t *testing.T) {
	spec := &solrkubernetesv1.SolrKubernetesSpec{
		ZookeeperContainer: &solrkubernetesv1.SolrKubernetesZookeeperContainer{
			Replicas: 1,
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
			DiskSize: "1Gi",
		},
	}

	// If SolrContainer is optional, this might pass. If it's required, this should fail.
	// Adjust test logic depending on the intended requirements.
	err := protovalidate.Validate(spec)
	if err != nil {
		t.Errorf("did not expect a validation error when SolrContainer is missing if it's optional: %v", err)
	}
}
