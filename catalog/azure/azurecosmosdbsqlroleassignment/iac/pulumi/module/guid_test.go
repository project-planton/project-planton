package module

import (
	"testing"

	"github.com/google/uuid"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type noMocks struct{}

func (noMocks) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	return args.Name + "-id", args.Inputs, nil
}

func (noMocks) Call(args pulumi.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}

// guidsIn derives the GUID for each metadata id inside one run of the stack.
func guidsIn(t *testing.T, stack string, metadataIds ...string) []string {
	t.Helper()
	var guids []string
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		for _, id := range metadataIds {
			guids = append(guids, stableGuid(ctx, id))
		}
		return nil
	}, pulumi.WithMocks("cosmos-role-assignment", stack, noMocks{}))
	if err != nil {
		t.Fatal(err)
	}
	return guids
}

// An unpinned assignment keeps one GUID from run to run: a GUID drawn fresh on
// each run would replace an unchanged assignment every time.
func TestStableGuid_TheSameIdentityDerivesTheSameGuidOnEveryRun(t *testing.T) {
	first := guidsIn(t, "prod", "azcdbsqlra-01", "")
	second := guidsIn(t, "prod", "azcdbsqlra-01", "")
	if first[0] != second[0] || first[1] != second[1] {
		t.Fatalf("an identity must derive one GUID: %v then %v", first, second)
	}
	for _, guid := range first {
		if _, err := uuid.Parse(guid); err != nil {
			t.Errorf("%q is not a GUID: %v", guid, err)
		}
	}
}

func TestStableGuid_DifferentIdentitiesDeriveDifferentGuids(t *testing.T) {
	ids := guidsIn(t, "prod", "azcdbsqlra-01", "azcdbsqlra-02")
	stacks := append(guidsIn(t, "prod", ""), guidsIn(t, "staging", "")...)
	if ids[0] == ids[1] {
		t.Errorf("two metadata ids derived one GUID %q", ids[0])
	}
	if stacks[0] == stacks[1] {
		t.Errorf("two stacks without a metadata id derived one GUID %q", stacks[0])
	}
}
