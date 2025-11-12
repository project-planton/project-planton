package manifestprotobuf

import (
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/workload/rediskubernetes/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestSetProtoField(t *testing.T) {
	tests := []struct {
		name      string
		message   proto.Message
		fieldPath string
		value     interface{}
		expected  proto.Message
		expectErr bool
	}{
		{
			name: "Set existing string field in snake case",
			message: &rediskubernetesv1.RedisKubernetes{Spec: &rediskubernetesv1.RedisKubernetesSpec{
				Container: &rediskubernetesv1.RedisKubernetesContainer{
					DiskSize: "1Gi",
				},
			}},
			fieldPath: "spec.container.disk_size",
			value:     "2Gi",
			expected: &rediskubernetesv1.RedisKubernetes{Spec: &rediskubernetesv1.RedisKubernetesSpec{
				Container: &rediskubernetesv1.RedisKubernetesContainer{
					DiskSize: "2Gi",
				},
			}},
			expectErr: false,
		},
		{
			name: "Set existing string field in camel case",
			message: &rediskubernetesv1.RedisKubernetes{Spec: &rediskubernetesv1.RedisKubernetesSpec{
				Container: &rediskubernetesv1.RedisKubernetesContainer{
					DiskSize: "1Gi",
				},
			}},
			fieldPath: "spec.container.diskSize",
			value:     "2Gi",
			expected: &rediskubernetesv1.RedisKubernetes{Spec: &rediskubernetesv1.RedisKubernetesSpec{
				Container: &rediskubernetesv1.RedisKubernetesContainer{
					DiskSize: "2Gi",
				},
			}},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SetProtoField(tt.message, tt.fieldPath, tt.value)
			if tt.expectErr {
				assert.Error(t, err, "Expected error but got none")
			} else {
				assert.NoError(t, err, "Unexpected error: %v", err)
				assert.True(t, proto.Equal(tt.expected, result), "Expected %v but got %v", tt.expected, result)
			}
		})
	}
}
