package tofumodule

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/plantonhq/planton/pkg/iac/provisioner"
)

// A committed lock is honoured exactly by OpenTofu; without one, or under the
// terraform binary (whose registry the lock does not name), init resolves
// providers itself.
func TestInitHonoursACommittedLockOnlyUnderOpenTofu(t *testing.T) {
	locked := t.TempDir()
	if err := os.WriteFile(filepath.Join(locked, ".terraform.lock.hcl"), []byte("# lock\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	unlocked := t.TempDir()

	cases := []struct {
		name   string
		engine provisioner.ProvisionerType
		dir    string
		want   []string
	}{
		{"tofu with a lock", provisioner.ProvisionerTypeTofu, locked, []string{"-lockfile=readonly"}},
		{"tofu without a lock", provisioner.ProvisionerTypeTofu, unlocked, nil},
		{"terraform with a lock", provisioner.ProvisionerTypeTerraform, locked, nil},
	}
	for _, tc := range cases {
		if got := lockFileArgs(tc.engine, tc.dir); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: lockFileArgs = %v, want %v", tc.name, got, tc.want)
		}
	}
}
