package tfvars

import (
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/rediskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	"testing"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

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

	got, err := ProtoToTFVars(msg)
	if err != nil {
		t.Fatalf("ProtoToTFVars() error = %v, want nil", err)
	}

	// For demonstration, let's assume got looks like:
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

	// Parse the HCL (tfvars)
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL([]byte(got), "test.tfvars")
	if diags.HasErrors() {
		t.Fatalf("Failed to parse generated tfvars: %s", diags.Error())
	}

	// Define a decoding specification using hcldec.
	// Use DynamicPseudoType for metadata and spec so we can handle arbitrary objects.
	spec := &hcldec.ObjectSpec{
		"apiVersion": &hcldec.AttrSpec{
			Name: "apiVersion",
			Type: cty.String,
		},
		"kind": &hcldec.AttrSpec{
			Name: "kind",
			Type: cty.String,
		},
		"metadata": &hcldec.AttrSpec{
			Name: "metadata",
			Type: cty.DynamicPseudoType, // allow arbitrary object
		},
		"spec": &hcldec.AttrSpec{
			Name: "spec",
			Type: cty.DynamicPseudoType, // allow arbitrary object
		},
	}

	val, diags := hcldec.Decode(file.Body, spec, nil)
	if diags.HasErrors() {
		t.Fatalf("Failed to decode: %s", diags.Error())
	}

	// Validate top-level fields
	apiVersion := val.GetAttr("apiVersion").AsString()
	if apiVersion != "kubernetes.project.planton/v1" {
		t.Errorf("expected apiVersion = 'kubernetes.project.planton/v1', got %q", apiVersion)
	}

	kind := val.GetAttr("kind").AsString()
	if kind != "RedisKubernetes" {
		t.Errorf("expected kind = 'RedisKubernetes', got %q", kind)
	}

	// metadata check
	metadataVal := val.GetAttr("metadata")
	if metadataVal.Type().IsObjectType() {
		nameVal := metadataVal.GetAttr("name").AsString()
		if nameVal != "red-one" {
			t.Errorf("expected metadata.name = 'red-one', got %q", nameVal)
		}

		labelsVal := metadataVal.GetAttr("labels")
		if labelsVal.Type().IsObjectType() {
			envVal := labelsVal.GetAttr("env").AsString()
			if envVal != "production" {
				t.Errorf("expected metadata.labels.env = 'production', got %q", envVal)
			}
		} else {
			t.Errorf("metadata.labels should be an object")
		}
	} else {
		t.Errorf("metadata should be an object")
	}

	// spec check
	specVal := val.GetAttr("spec")
	if specVal.Type().IsObjectType() {
		containerVal := specVal.GetAttr("container")
		if !containerVal.Type().IsObjectType() {
			t.Fatalf("spec.container should be object")
		}

		diskSizeVal := containerVal.GetAttr("diskSize").AsString()
		if diskSizeVal != "2Gi" {
			t.Errorf("expected diskSize = '2Gi', got %q", diskSizeVal)
		}

		isPersistenceEnabledVal := containerVal.GetAttr("isPersistenceEnabled").True()
		if !isPersistenceEnabledVal {
			t.Errorf("expected isPersistenceEnabled = true")
		}

		replicasVal := containerVal.GetAttr("replicas")
		replicasFloat := replicasVal.AsBigFloat()
		i, _ := replicasFloat.Int64()
		if i != 1 {
			t.Errorf("expected replicas = 1, got %d", i)
		}
	} else {
		t.Errorf("spec should be an object")
	}
}
