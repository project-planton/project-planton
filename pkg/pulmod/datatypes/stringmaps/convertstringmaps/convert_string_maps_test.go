package convertstringmaps

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
)

func TestConvertGoMapToPulumiMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected pulumi.Map
	}{
		{
			name:     "Empty map",
			input:    map[string]string{},
			expected: pulumi.Map{},
		},
		{
			name: "Single element map",
			input: map[string]string{
				"key1": "value1",
			},
			expected: pulumi.Map{
				"key1": pulumi.String("value1"),
			},
		},
		{
			name: "Multiple element map",
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expected: pulumi.Map{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("value2"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertGoStringMapToPulumiStringMap(tt.input)
			for key, expectedValue := range tt.expected {
				actualValue, ok := result[key]
				if !ok {
					t.Errorf("expected key %s not found in result", key)
				}
				assert.Equal(t, expectedValue, actualValue)
			}
		})
	}
}
