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
	if diags != nil && diags.HasErrors() {
		t.Fatalf("Failed to parse variables.tf: %s", diags.Error())
	}

	if file == nil {
		t.Fatalf("Failed to parse variables.tf: file is nil without diagnostics")
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
	if diags != nil && diags.HasErrors() {
		t.Fatalf("Failed to get content: %s", diags.Error())
	}

	if content == nil {
		t.Fatalf("Failed to get content: content is nil but no diagnostics returned")
	}

	// Expecting two variable blocks: "metadata" and "spec"
	if len(content.Blocks) != 2 {
		t.Fatalf("Expected 2 variables, got %d", len(content.Blocks))
	}

	for _, block := range content.Blocks {
		varName := block.Labels[0]
		attrs, diags := block.Body.JustAttributes()
		if diags != nil && diags.HasErrors() {
			t.Fatalf("Failed to read attributes: %s", diags.Error())
		}

		// Check that "description" attribute exists and is a string
		descAttr := attrs["description"]
		if descAttr == nil {
			t.Errorf("variable %s missing description attribute", varName)
			continue
		}
		descVal, diags := descAttr.Expr.Value(nil)
		if diags != nil && diags.HasErrors() {
			t.Errorf("Failed to get description value for %s: %s", varName, diags.Error())
			continue
		}

		if descVal.Type() != cty.String {
			t.Errorf("Expected %s description to be string, got %s", varName, descVal.Type().FriendlyName())
		}

		// Check that "type" attribute exists
		typeAttr := attrs["type"]
		if typeAttr == nil {
			t.Errorf("variable %s missing type attribute", varName)
			continue
		}

		// Instead of evaluating the type (which would require Terraform functions),
		// we directly inspect the source code for the 'type' attribute.
		typeRange := typeAttr.Expr.Range()
		exprText := variablesTF[typeRange.Start.Byte:typeRange.End.Byte]

		// Check that it contains "object("
		if !strings.Contains(exprText, "object(") {
			t.Errorf("Expected %s to have an object type, got %q", varName, exprText)
		}
	}
}
