package variablestf

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"strings"
	"testing"
)

func TestVariablesTF(t *testing.T) {
	variablesTF := `
variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string
  })
}

variable "spec" {
  description = "Specification for the S3Bucket"
  type = object({
    is_public  = bool
    aws_region = string
  })
}
`

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL([]byte(variablesTF), "variables.tf")
	if diags.HasErrors() {
		t.Fatalf("Failed to parse variables.tf: %s", diags.Error())
	}

	// Define a block schema that expects variable blocks
	blockSchema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "variable",
				LabelNames: []string{"name"},
			},
		},
	}

	content, diags := file.Body.Content(blockSchema)
	if diags.HasErrors() {
		t.Fatalf("Failed to get content: %s", diags.Error())
	}

	// Expecting two variable blocks: "metadata" and "spec"
	if len(content.Blocks) != 2 {
		t.Fatalf("Expected 2 variables, got %d", len(content.Blocks))
	}

	for _, block := range content.Blocks {
		varName := block.Labels[0]
		attrs, diags := block.Body.JustAttributes()
		if diags.HasErrors() {
			t.Fatalf("Failed to read attributes: %s", diags.Error())
		}

		// Check that "description" attribute exists and is a string
		descAttr := attrs["description"]
		if descAttr == nil {
			t.Errorf("variable %s missing description attribute", varName)
			continue
		}
		descVal, diags := descAttr.Expr.Value(nil)
		if diags.HasErrors() {
			t.Errorf("Failed to get description value for %s: %s", varName, diags.Error())
			continue
		}
		if descVal.Type() != cty.String {
			t.Errorf("variable %s description should be string", varName)
		}

		// Check that "type" attribute exists
		typeAttr := attrs["type"]
		if typeAttr == nil {
			t.Errorf("variable %s missing type attribute", varName)
			continue
		}

		typeVal, diags := typeAttr.Expr.Value(nil)
		if diags.HasErrors() {
			t.Errorf("Failed to get type value for %s: %s", varName, diags.Error())
			continue
		}

		// typeVal is cty.String representing HCL's literal. For more detailed checking,
		// you'd need to parse the type expression (which is an HCL expression) and assert
		// it forms a correct Terraform type. Since `type` is a Terraform TypeSpec, you may
		// need to parse it from the HCL expression AST directly.

		// For illustration, let's just check it contains "object(".
		typeStr := typeVal.AsString()
		if !strings.Contains(typeStr, "object(") {
			t.Errorf("Expected %s to have an object type, got %q", varName, typeStr)
		}
	}
}
