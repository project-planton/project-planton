package module

import (
	"reflect"
	"sync"
	"testing"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/encoding/protojson"

	gcpgkeclusterv1alpha1 "github.com/plantonhq/planton/catalog/gcp/gcpgkecluster/v1alpha1"
)

type authorizedNetworksMocks struct {
	mu     sync.Mutex
	inputs resource.PropertyMap
}

func (m *authorizedNetworksMocks) Call(args pulumi.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}

func (m *authorizedNetworksMocks) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	outputs := args.Inputs.Copy()
	if args.TypeToken == "gcp:container/cluster:Cluster" {
		m.mu.Lock()
		m.inputs = args.Inputs.Copy()
		m.mu.Unlock()
		outputs["masterAuth"] = resource.NewObjectProperty(resource.PropertyMap{
			"clusterCaCertificate": resource.NewStringProperty("public-ca"),
		})
		outputs["fleet"] = resource.NewObjectProperty(resource.PropertyMap{
			"membership": resource.NewStringProperty(""),
		})
	}
	return args.Name + "-id", outputs, nil
}

// Inspect the inputs registered by the real cluster module, without a provider
// process or cloud credentials. Presence of the empty provider block matters.
func TestAuthorizedNetworksRegisteredInputs(t *testing.T) {
	cases := []struct {
		name, policy string
		want         map[string]any
	}{
		{"omitted", "", nil},
		{"unrestricted", `,"unrestrictedNetworks":{}`, nil},
		{"empty restriction", `,"masterAuthorizedNetworks":{}`, map[string]any{"cidrBlocks": []any{}}},
		{"explicit false", `,"masterAuthorizedNetworks":{"gcpPublicCidrsAccessEnabled":false,"privateEndpointEnforcementEnabled":false}`, map[string]any{
			"cidrBlocks": []any{}, "gcpPublicCidrsAccessEnabled": false, "privateEndpointEnforcementEnabled": false,
		}},
		{"populated restriction", `,"masterAuthorizedNetworks":{"cidrBlocks":[{"cidrBlock":"203.0.113.0/24","displayName":"office"}],"privateEndpointEnforcementEnabled":true}`, map[string]any{
			"cidrBlocks":                        []any{map[string]any{"cidrBlock": "203.0.113.0/24", "displayName": "office"}},
			"privateEndpointEnforcementEnabled": true,
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			inputs := registeredClusterInputs(t, tc.policy)
			value, present := inputs["masterAuthorizedNetworksConfig"]
			if tc.want == nil {
				if present && !value.IsNull() {
					t.Fatalf("unrestricted policy emitted an IP restriction: %v", value)
				}
				return
			}
			if !present || !value.IsObject() {
				t.Fatalf("configured policy lost its provider block: %v", inputs)
			}
			if got := value.ObjectValue().Mappable(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("registered authorized-networks config = %#v; want %#v", got, tc.want)
			}
		})
	}
}

func registeredClusterInputs(t *testing.T, policy string) resource.PropertyMap {
	t.Helper()
	var spec gcpgkeclusterv1alpha1.GcpGkeClusterSpec
	input := `{"projectId":{"value":"test-project"},"location":"us-central1","network":{"value":"test-vpc"},"subnetwork":{"value":"test-subnet"}` + policy + `}`
	if err := protojson.Unmarshal([]byte(input), &spec); err != nil {
		t.Fatal(err)
	}
	mocks := &authorizedNetworksMocks{}
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		provider, err := gcp.NewProvider(ctx, "gcp", &gcp.ProviderArgs{Project: pulumi.String("test-project")})
		if err != nil {
			return errors.Wrap(err, "create mocked GCP provider")
		}
		return cluster(ctx, &Locals{
			GcpGkeCluster: &gcpgkeclusterv1alpha1.GcpGkeCluster{Spec: &spec},
			ClusterName:   "policy-test", ReleaseChannel: "REGULAR", GcpLabels: map[string]string{},
		}, provider)
	}, pulumi.WithMocks("project", "stack", mocks))
	if err != nil {
		t.Fatalf("cluster module failed: %v", err)
	}
	if mocks.inputs == nil {
		t.Fatal("cluster module registered no cluster")
	}
	return mocks.inputs
}
