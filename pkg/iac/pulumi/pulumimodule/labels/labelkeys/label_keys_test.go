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
			testName:                "project-planton org label should be converted to prometheus label",
			inputLabel:              "project-planton.org/org",
			expectedPrometheusLabel: "project_planton_org_org",
		},
		{
			testName:                "project-planton service label should be converted to prometheus label",
			inputLabel:              "project-planton.org/service",
			expectedPrometheusLabel: "project_planton_org_service",
		},
		{
			testName:                "project-planton service-env label should be converted to prometheus label",
			inputLabel:              "project-planton.org/env",
			expectedPrometheusLabel: "project_planton_org_env",
		},
		{
			testName:                "project-planton kind label should be converted to prometheus label",
			inputLabel:              "project-planton.org/kind",
			expectedPrometheusLabel: "project_planton_org_kind",
		},
		{
			testName:                "project-planton id label should be converted to prometheus label",
			inputLabel:              "project-planton.org/id",
			expectedPrometheusLabel: "project_planton_org_id",
		},
	}
	t.Run("test project-planton label conversion to prometheus format labels", func(t *testing.T) {
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
