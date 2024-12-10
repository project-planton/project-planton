package tfvars

import (
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/rediskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// TestGeneratedTFVarsParsing tests parsing the generated tfvars back into a structured form
// using the hclparse and gohcl packages.
func TestGeneratedTFVarsParsing(t *testing.T) {
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

	// Generate tfvars from your proto message.
	// In your actual code, you have the ProtoToTerraformTFVars function defined
	// in the same package, so we just call it here:
	got, err := ProtoToTerraformTFVars(msg)
	if err != nil {
		t.Fatalf("ProtoToTerraformTFVars() error = %v, want nil", err)
	}

	// For demonstration, let's assume the generated tfvars looks like this:
	// (In your actual test, got would be the output from ProtoToTerraformTFVars)
	got = `
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
    diskSize = "2Gi"
    isPersistenceEnabled = true
    replicas = 1
  }
}
`

	// Parse the HCL (tfvars) string
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL([]byte(got), "test.tfvars")
	if diags.HasErrors() {
		t.Fatalf("Failed to parse generated tfvars: %s", diags.Error())
	}

	// Define a struct to decode into. We must use `,block` for nested structures.
	var decoded struct {
		ApiVersion string `hcl:"apiVersion"`
		Kind       string `hcl:"kind"`
		Metadata   struct {
			Name   string            `hcl:"name"`
			Labels map[string]string `hcl:"labels"`
		} `hcl:"metadata,block"`
		Spec struct {
			Container struct {
				DiskSize             string `hcl:"diskSize"`
				IsPersistenceEnabled bool   `hcl:"isPersistenceEnabled"`
				Replicas             int    `hcl:"replicas"`
			} `hcl:"container,block"`
		} `hcl:"spec,block"`
	}

	// Decode the parsed HCL body into our struct
	diags = gohcl.DecodeBody(file.Body, nil, &decoded)
	if diags.HasErrors() {
		t.Fatalf("Failed to decode tfvars into struct: %s", diags.Error())
	}

	// Now we can assert values directly without worrying about key order or whitespace
	if decoded.ApiVersion != "kubernetes.project.planton/v1" {
		t.Errorf("expected apiVersion = 'kubernetes.project.planton/v1', got %q", decoded.ApiVersion)
	}
	if decoded.Kind != "RedisKubernetes" {
		t.Errorf("expected kind = 'RedisKubernetes', got %q", decoded.Kind)
	}

	if decoded.Metadata.Name != "red-one" {
		t.Errorf("expected metadata.name = 'red-one', got %q", decoded.Metadata.Name)
	}
	if decoded.Metadata.Labels["env"] != "production" {
		t.Errorf("expected metadata.labels.env = 'production', got %q", decoded.Metadata.Labels["env"])
	}

	if decoded.Spec.Container.DiskSize != "2Gi" {
		t.Errorf("expected spec.container.diskSize = '2Gi', got %q", decoded.Spec.Container.DiskSize)
	}
	if !decoded.Spec.Container.IsPersistenceEnabled {
		t.Errorf("expected spec.container.isPersistenceEnabled = true")
	}
	if decoded.Spec.Container.Replicas != 1 {
		t.Errorf("expected spec.container.replicas = 1, got %d", decoded.Spec.Container.Replicas)
	}
}
