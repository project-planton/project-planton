package labelkeys

import (
	"testing"
)

func TestLabelConversionPrometheusFormat(t *testing.T) {
	testCases := []struct {
		testName                string
		inputLabel              string
		expectedPrometheusLabel string
	}{
		{
			testName:                "planton-cloud company label should be converted to prometheus label",
			inputLabel:              "planton.cloud/company",
			expectedPrometheusLabel: "planton_cloud_company",
		},
		{
			testName:                "planton-cloud product label should be converted to prometheus label",
			inputLabel:              "planton.cloud/product",
			expectedPrometheusLabel: "planton_cloud_product",
		},
		{
			testName:                "planton-cloud product-env label should be converted to prometheus label",
			inputLabel:              "planton.cloud/env",
			expectedPrometheusLabel: "planton_cloud_env",
		},
		{
			testName:                "planton-cloud resource-type label should be converted to prometheus label",
			inputLabel:              "planton.cloud/resource-type",
			expectedPrometheusLabel: "planton_cloud_resource_type",
		},
		{
			testName:                "planton-cloud resource-id label should be converted to prometheus label",
			inputLabel:              "planton.cloud/resource-id",
			expectedPrometheusLabel: "planton_cloud_resource_id",
		},
	}
	t.Run("test planton-cloud label conversion to prometheus format labels", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				r := WithPrometheusFormat(tc.inputLabel)
				if r != tc.expectedPrometheusLabel {
					t.Errorf("expected: %s, got: %s", tc.expectedPrometheusLabel, r)
				}
			})
		}
	})
}
