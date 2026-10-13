package gcpgkeclusterv1alpha1

import (
	"encoding/json"
	"reflect"
	"testing"

	"buf.build/go/protovalidate"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// A present empty restriction is a policy, not an unconfigured form section.
// Exercise both serialization boundaries so edits cannot erase that intent.
func TestAuthorizedNetworksPolicyRoundTrip(t *testing.T) {
	cases := []struct {
		name, policy string
	}{
		{"omitted", `{}`},
		{"empty restriction", `{"masterAuthorizedNetworks":{}}`},
		{"explicit unrestricted", `{"unrestrictedNetworks":{}}`},
		{"explicit false", `{"masterAuthorizedNetworks":{"gcpPublicCidrsAccessEnabled":false,"privateEndpointEnforcementEnabled":false}}`},
		{"populated restriction", `{"masterAuthorizedNetworks":{"cidrBlocks":[{"cidrBlock":"203.0.113.0/24","displayName":"office"}],"privateEndpointEnforcementEnabled":true}}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			baseline, err := protojson.Marshal(minimalSpec())
			if err != nil {
				t.Fatal(err)
			}
			payload := jsonObject(t, baseline)
			policy := jsonObject(t, []byte(tc.policy))
			for key, value := range policy {
				payload[key] = value
			}
			input, err := json.Marshal(payload)
			if err != nil {
				t.Fatal(err)
			}
			var spec GcpGkeClusterSpec
			if err := protojson.Unmarshal(input, &spec); err != nil {
				t.Fatal(err)
			}
			if err := protovalidate.Validate(newCluster(&spec)); err != nil {
				t.Fatalf("valid policy rejected: %v", err)
			}
			wire, err := proto.Marshal(&spec)
			if err != nil {
				t.Fatal(err)
			}
			var restored GcpGkeClusterSpec
			if err := proto.Unmarshal(wire, &restored); err != nil {
				t.Fatal(err)
			}
			output, err := protojson.Marshal(&restored)
			if err != nil {
				t.Fatal(err)
			}
			actual := jsonObject(t, output)
			for _, key := range []string{"masterAuthorizedNetworks", "unrestrictedNetworks"} {
				want, wantPresent := policy[key]
				got, gotPresent := actual[key]
				if wantPresent != gotPresent || !reflect.DeepEqual(want, got) {
					t.Errorf("%s changed through JSON/binary round trip: want %#v (present %t), got %#v (present %t)", key, want, wantPresent, got, gotPresent)
				}
			}
		})
	}
}

func TestAuthorizedNetworksPolicyRejectsContradictoryJSON(t *testing.T) {
	for _, input := range []string{
		`{"masterAuthorizedNetworks":{},"unrestrictedNetworks":{}}`,
		`{"unrestrictedNetworks":{},"masterAuthorizedNetworks":{}}`,
	} {
		var spec GcpGkeClusterSpec
		if err := protojson.Unmarshal([]byte(input), &spec); err == nil {
			t.Errorf("conflicting network policies must not silently pick a winner: %s", input)
		}
	}
}

func TestAuthorizedNetworksRetainsExistingFieldIdentity(t *testing.T) {
	field := (&GcpGkeClusterSpec{}).ProtoReflect().Descriptor().Fields().ByName("master_authorized_networks")
	if field.Number() != 28 || field.JSONName() != "masterAuthorizedNetworks" {
		t.Fatalf("existing policy field identity changed: number %d, JSON name %q", field.Number(), field.JSONName())
	}
	if field.ContainingOneof() == nil || field.ContainingOneof().Name() != "authorized_networks_policy" {
		t.Fatal("a present empty policy must select the real authorized-networks choice")
	}
}

func jsonObject(t *testing.T, data []byte) map[string]any {
	t.Helper()
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}
	return result
}
