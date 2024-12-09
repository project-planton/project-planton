package tfvars

import (
	"github.com/google/go-cmp/cmp"
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/rediskubernetes/v1"
	"github.com/project-planton/project-planton/apis/go/project/planton/shared"
	"github.com/project-planton/project-planton/apis/go/project/planton/shared/kubernetes"
	"strings"

	"testing"
)

// ProtoToTerraformTFVars is assumed to be defined in yourpackage.
// Import it if it's in another package.
func TestProtoToTerraformTFVars_Basic(t *testing.T) {
	// Create a test proto message with some fields.
	msg := &rediskubernetesv1.RedisKubernetes{
		ApiVersion: "kubernetes.project.planton/v1",
		Kind:       "RedisKubernetes",
		Metadata: &shared.ApiResourceMetadata{
			Name: "red-one",
			Labels: map[string]string{
				"env": "production",
			},
		},
		Spec: &rediskubernetesv1.RedisKubernetesSpec{
			Container: &rediskubernetesv1.RedisKubernetesContainer{
				DiskSize:             "2Gi",
				IsPersistenceEnabled: true,
				Replicas:             1,
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
		},
	}

	got, err := ProtoToTerraformTFVars(msg)
	if err != nil {
		t.Fatalf("ProtoToTerraformTFVars() error = %v, want nil", err)
	}

	// Expected output: This depends on how your tfvars are structured.
	// The output should reflect the structure of the message.
	// Note that keys are in sorted order by map iteration in the code above,
	// so final ordering may differ. Adjust as needed or use cmp.Diff.
	want := `
apiVersion = "kubernetes.project.planton/v1"
kind = "RedisKubernetes"
metadata = {
  name = "red-one"
  labels = {
    env = "production"
  }
}
spec = {
  container = {
    resources = {
      limits = {
        cpu = "1000m"
        memory = "1Gi"
      }
      requests = {
        cpu = "50m"
        memory = "100Mi"
      }
    }
    isPersistenceEnabled = true
    diskSize = "2Gi"
    replicas = 1
  }
}
`
	// For example, you might expect something like:
	// Normalize whitespace because indentation might differ slightly
	gotTrimmed := strings.TrimSpace(got)
	wantTrimmed := strings.TrimSpace(want)

	if diff := cmp.Diff(wantTrimmed, gotTrimmed); diff != "" {
		t.Errorf("ProtoToTerraformTFVars() mismatch (-want +got):\n%s", diff)
	}
}

func TestProtoToTerraformTFVars_EmptyFields(t *testing.T) {
	// Test with minimal fields set
	msg := &rediskubernetesv1.RedisKubernetes{
		ApiVersion: "",
		Kind:       "",
		Metadata:   nil,
		Spec:       nil,
	}

	got, err := ProtoToTerraformTFVars(msg)
	if err != nil {
		t.Fatalf("ProtoToTerraformTFVars() error = %v, want nil", err)
	}

	// Expect tfvars with empty or null values
	want := `
apiVersion = ""
kind = ""
metadata = null
spec = null
`
	gotTrimmed := strings.TrimSpace(got)
	wantTrimmed := strings.TrimSpace(want)

	if diff := cmp.Diff(wantTrimmed, gotTrimmed); diff != "" {
		t.Errorf("ProtoToTerraformTFVars() mismatch (-want +got):\n%s", diff)
	}
}

func TestProtoToTerraformTFVars_NestedArrays(t *testing.T) {
	// Suppose the message can contain an array of strings or objects
	msg := &rediskubernetesv1.RedisKubernetes{
		Kind: "ExampleWithArray",
		Metadata: &shared.ApiResourceMetadata{
			Name: "resource-with-array",
		},
		// Assume we have a field repeated in proto (e.g., repeated string items)
		// For demonstration, let's say `Spec` can contain a field `AllowedZones []string`
		// Adjust this to your actual fields.
		Spec: &rediskubernetesv1.RedisKubernetesSpec{
			Container: &rediskubernetesv1.RedisKubernetesContainer{
				DiskSize: "5Gi",
			},
		},
	}

	got, err := ProtoToTerraformTFVars(msg)
	if err != nil {
		t.Fatalf("ProtoToTerraformTFVars() error = %v", err)
	}

	want := `
apiVersion = "kubernetes.project.planton/v1"
kind = "ExampleWithArray"
metadata = {
  name = "resource-with-array"
}
spec = {
  allowedZones = [
    "us-east-1a",
    "us-east-1b",
  ]
  container = {
    diskSize = "5Gi"
  }
}
`
	gotTrimmed := strings.TrimSpace(got)
	wantTrimmed := strings.TrimSpace(want)

	if diff := cmp.Diff(wantTrimmed, gotTrimmed); diff != "" {
		t.Errorf("ProtoToTerraformTFVars() mismatch (-want +got):\n%s", diff)
	}
}

//func TestProtoToTerraformTFVars_ComplexNestedStructures(t *testing.T) {
//	// Assume you have a dynamic field that stores arbitrary structured data in a google.protobuf.Struct
//	complexData, _ := structpb.NewStruct(map[string]interface{}{
//		"configMap": map[string]interface{}{
//			"max_connections": float64(100),
//			"enable_feature":  true,
//		},
//	})
//
//	msg := &rediskubernetesv1.RedisKubernetes{
//		ApiVersion: "kubernetes.project.planton/v1",
//		Kind:       "Complex",
//		Metadata: &shared.ApiResourceMetadata{
//			Name: "complex-resource",
//		},
//	}
//
//	got, err := ProtoToTerraformTFVars(msg)
//	if err != nil {
//		t.Fatalf("ProtoToTerraformTFVars() error = %v", err)
//	}
//
//	// Expect the struct fields to appear as nested maps
//	want := `
//apiVersion = "kubernetes.project.planton/v1"
//kind = "Complex"
//metadata = {
//  name = "complex-resource"
//}
//complexData = {
//  configMap = {
//    enable_feature = true
//    max_connections = 100
//  }
//}
//`
//	gotTrimmed := strings.TrimSpace(got)
//	wantTrimmed := strings.TrimSpace(want)
//
//	if diff := cmp.Diff(wantTrimmed, gotTrimmed); diff != "" {
//		t.Errorf("ProtoToTerraformTFVars() mismatch (-want +got):\n%s", diff)
//	}
//}

// Additional tests can check error conditions if you add error handling logic:
// For example, if unsupported types or invalid input occur, test that errors are returned.
func TestProtoToTerraformTFVars_UnsupportedType(t *testing.T) {
	// If you modify the code to fail on certain conditions,
	// you can test for expected errors here.
	// For now, we assume all types are supported.
	// This is a placeholder example.
	msg := &rediskubernetesv1.RedisKubernetes{}

	// If we had a code path that fails:
	_, err := ProtoToTerraformTFVars(msg)
	if err != nil {
		// Check if the error message matches what we expect
		t.Logf("Got expected error: %v", err)
	} else {
		t.Error("Expected an error but got none")
	}
}
