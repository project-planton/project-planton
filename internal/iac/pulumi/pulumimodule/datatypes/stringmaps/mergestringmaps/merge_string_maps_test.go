package mergestringmaps

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"reflect"
	"testing"
)

// Utility function to simplify pulumi.String creation in test cases
func ps(s string) pulumi.String {
	return pulumi.String(s)
}

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name     string
		baseMap  pulumi.Map
		newMap   map[string]string
		expected pulumi.Map
	}{
		{
			name: "simple key update",
			baseMap: pulumi.Map{
				"key1": ps("value1"),
			},
			newMap: map[string]string{
				"key1": "updatedValue",
			},
			expected: pulumi.Map{
				"key1": ps("updatedValue"),
			},
		},
		{
			name: "nested key update",
			baseMap: pulumi.Map{
				"nested": pulumi.Map{
					"key1": ps("value1"),
				},
			},
			newMap: map[string]string{
				"nested.key1": "updatedValue",
			},
			expected: pulumi.Map{
				"nested": pulumi.Map{
					"key1": ps("updatedValue"),
				},
			},
		},
		{
			name: "add new top-level key",
			baseMap: pulumi.Map{
				"existingKey": ps("existingValue"),
			},
			newMap: map[string]string{
				"newKey": "newValue",
			},
			expected: pulumi.Map{
				"existingKey": ps("existingValue"),
				"newKey":      ps("newValue"),
			},
		},
		{
			name: "add new nested key",
			baseMap: pulumi.Map{
				"existingKey": pulumi.Map{
					"nestedKey": ps("existingValue"),
				},
			},
			newMap: map[string]string{
				"existingKey.newNestedKey": "newValue",
			},
			expected: pulumi.Map{
				"existingKey": pulumi.Map{
					"nestedKey":    ps("existingValue"),
					"newNestedKey": ps("newValue"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeMapToPulumiMap(tt.baseMap, tt.newMap)
			if !reflect.DeepEqual(tt.baseMap, tt.expected) {
				t.Errorf("mergeHelmValuesMap() = %#v, want %#v", tt.baseMap, tt.expected)
			}
		})
	}
}
