package upgrade

import "testing"

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name     string
		current  string
		latest   string
		expected bool
	}{
		// Current is higher semver - no upgrade
		{"higher_semver_plain", "v0.3.17", "v0.3.15-cli.20260113.1", false},
		{"higher_semver_plain_2", "v0.3.17", "v0.3.16-cli.20260113.1", false},

		// Current is lower semver - upgrade needed
		{"lower_semver_plain", "v0.3.14", "v0.3.15-cli.20260113.1", true},
		{"lower_semver_plain_2", "v0.3.15", "v0.3.16-cli.20260113.1", true},

		// Same semver (current is plain, latest is CLI format)
		{"same_semver_mixed", "v0.3.15", "v0.3.15-cli.20260113.1", false},

		// Both CLI format - proper comparison
		{"cli_newer_revision", "v0.3.15-cli.20260113.0", "v0.3.15-cli.20260113.1", true},
		{"cli_older_revision", "v0.3.15-cli.20260113.1", "v0.3.15-cli.20260113.0", false},

		// Dev version - always upgrade
		{"dev_version", "dev", "v0.3.15-cli.20260113.1", true},

		// Same version
		{"same_exact", "v0.3.15-cli.20260113.1", "v0.3.15-cli.20260113.1", false},

		// Plain semver comparisons (new release format)
		{"plain_semver_upgrade", "v0.3.15", "v0.3.16", true},
		{"plain_semver_no_upgrade", "v0.3.16", "v0.3.15", false},
		{"plain_semver_same", "v0.3.16", "v0.3.16", false},

		// CLI format vs plain semver (latest is plain)
		{"cli_vs_plain_upgrade", "v0.3.15-cli.20260113.1", "v0.3.16", true},
		{"cli_vs_plain_no_upgrade", "v0.3.16-cli.20260113.1", "v0.3.15", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareVersions(tt.current, tt.latest)
			if result != tt.expected {
				t.Errorf("CompareVersions(%q, %q) = %v, expected %v", tt.current, tt.latest, result, tt.expected)
			}
		})
	}
}
