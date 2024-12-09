package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	kafkakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/kafkakubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

// TestKafkaKubernetesSpec_ValidSpec ensures a fully valid spec passes validation.
func TestKafkaKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		KafkaTopics: []*kafkakubernetesv1.KafkaTopic{
			{
				Name:       "validTopicName",
				Partitions: 3,
				Replicas:   3,
				Config: map[string]string{
					"cleanup.policy": "compact",
				},
			},
		},
		BrokerContainer: &kafkakubernetesv1.KafkaKubernetesBrokerContainer{
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
			DiskSize: "10Gi",
		},
		ZookeeperContainer: &kafkakubernetesv1.KafkaKubernetesZookeeperContainer{
			Replicas: 3,
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
			DiskSize: "10Gi",
		},
		SchemaRegistryContainer: &kafkakubernetesv1.KafkaKubernetesSchemaRegistryContainer{
			IsEnabled: true,
			Replicas:  1,
			Resources: &kubernetes.ContainerResources{
				Limits: &kubernetes.CpuMemory{
					Cpu:    "500m",
					Memory: "512Mi",
				},
				Requests: &kubernetes.CpuMemory{
					Cpu:    "100m",
					Memory: "256Mi",
				},
			},
		},
		Ingress: &kubernetes.IngressSpec{
			DnsDomain: "kafka.example.com",
		},
		IsDeployKafkaUi: true,
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

// TestKafkaKubernetesSpec_InvalidBrokerDiskSize checks that an invalid broker disk size fails validation.
func TestKafkaKubernetesSpec_InvalidBrokerDiskSize(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		BrokerContainer: &kafkakubernetesv1.KafkaKubernetesBrokerContainer{
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
			DiskSize: "invalidDiskSize", // Invalid format
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid broker disk size, got none")
	} else {
		if !strings.Contains(err.Error(), "[spec.broker_container.disk_size.format]") {
			t.Errorf("expected validation error with constraint id `spec.broker_container.disk_size.format`, got: %v", err)
		}
	}
}

// TestKafkaKubernetesSpec_InvalidZookeeperDiskSize checks that an invalid Zookeeper disk size fails validation.
func TestKafkaKubernetesSpec_InvalidZookeeperDiskSize(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		ZookeeperContainer: &kafkakubernetesv1.KafkaKubernetesZookeeperContainer{
			Replicas: 3,
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
			DiskSize: "123abc", // Invalid format
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid zookeeper disk size, got none")
	} else {
		if !strings.Contains(err.Error(), "[spec.broker_container.disk_size.format]") {
			t.Errorf("expected validation error with constraint id `spec.broker_container.disk_size.format`, got: %v", err)
		}
	}
}

// TestKafkaKubernetesSpec_InvalidTopicName checks that an invalid topic name fails validation.
// For example, topic that starts with a non-alphanumeric character.
func TestKafkaKubernetesSpec_InvalidTopicName(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		KafkaTopics: []*kafkakubernetesv1.KafkaTopic{
			{
				Name:       ".invalidName", // Starts with a non-alphanumeric character
				Partitions: 1,
				Replicas:   1,
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid topic name, got none")
	} else {
		if !strings.Contains(err.Error(), "Should start with an alphanumeric character") {
			t.Errorf("expected error message about topic name start, got: %v", err)
		}
	}
}

// TestKafkaKubernetesSpec_InvalidTopicNameEndsWithNonAlphanumeric checks a topic that doesn't end with alphanumeric.
func TestKafkaKubernetesSpec_InvalidTopicNameEndsWithNonAlphanumeric(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		KafkaTopics: []*kafkakubernetesv1.KafkaTopic{
			{
				Name:       "validName-",
				Partitions: 1,
				Replicas:   1,
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for topic name ending, got none")
	} else {
		if !strings.Contains(err.Error(), "Should end with an alphanumeric character") {
			t.Errorf("expected error about ending with alphanumeric, got: %v", err)
		}
	}
}

// TestKafkaKubernetesSpec_InvalidTopicNameContainsNonASCII checks a topic with non-ASCII characters.
func TestKafkaKubernetesSpec_InvalidTopicNameContainsNonASCII(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		KafkaTopics: []*kafkakubernetesv1.KafkaTopic{
			{
				Name:       "invalidName✓",
				Partitions: 1,
				Replicas:   1,
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for non-ASCII character in topic name, got none")
	} else {
		if !strings.Contains(err.Error(), "Must not contain non-ASCII characters") {
			t.Errorf("expected error about non-ASCII, got: %v", err)
		}
	}
}

// TestKafkaKubernetesSpec_InvalidTopicNameContainsDotDot checks a topic that contains "..".
func TestKafkaKubernetesSpec_InvalidTopicNameContainsDotDot(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		KafkaTopics: []*kafkakubernetesv1.KafkaTopic{
			{
				Name:       "invalid..name",
				Partitions: 1,
				Replicas:   1,
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for '..' in topic name, got none")
	} else {
		if !strings.Contains(err.Error(), "Must not contain '..'") {
			t.Errorf("expected error about '..', got: %v", err)
		}
	}
}

// TestKafkaKubernetesSpec_InvalidTopicNameInvalidChars checks a topic that contains invalid characters.
func TestKafkaKubernetesSpec_InvalidTopicNameInvalidChars(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		KafkaTopics: []*kafkakubernetesv1.KafkaTopic{
			{
				Name:       "invalid#name",
				Partitions: 1,
				Replicas:   1,
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid characters in topic name, got none")
	} else {
		if !strings.Contains(err.Error(), "Only alphanumeric and ('.', '_' and '-') characters are allowed") {
			t.Errorf("expected error about allowed characters, got: %v", err)
		}
	}
}

// TestKafkaKubernetesSpec_DisabledSchemaRegistry checks that disabling the schema registry passes validation.
func TestKafkaKubernetesSpec_DisabledSchemaRegistry(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		BrokerContainer: &kafkakubernetesv1.KafkaKubernetesBrokerContainer{
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
		ZookeeperContainer: &kafkakubernetesv1.KafkaKubernetesZookeeperContainer{
			Replicas: 3,
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
		SchemaRegistryContainer: &kafkakubernetesv1.KafkaKubernetesSchemaRegistryContainer{
			IsEnabled: false,
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors when schema registry is disabled, got: %v", err)
	}
}

// TestKafkaKubernetesSpec_EmptyKafkaTopics checks that having no topics defined is still valid.
func TestKafkaKubernetesSpec_EmptyKafkaTopics(t *testing.T) {
	spec := &kafkakubernetesv1.KafkaKubernetesSpec{
		BrokerContainer: &kafkakubernetesv1.KafkaKubernetesBrokerContainer{
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
			DiskSize: "10Gi",
		},
		ZookeeperContainer: &kafkakubernetesv1.KafkaKubernetesZookeeperContainer{
			Replicas: 3,
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
			DiskSize: "10Gi",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors with empty KafkaTopics, got: %v", err)
	}
}
