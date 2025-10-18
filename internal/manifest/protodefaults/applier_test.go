package protodefaults

import (
	"testing"

	externaldnskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/addon/externaldnskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestApplyDefaults_ExternalDnsKubernetes(t *testing.T) {
	t.Run("applies defaults to unset fields", func(t *testing.T) {
		// Create a message with minimal required fields, leaving fields with defaults unset
		msg := &externaldnskubernetesv1.ExternalDnsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ExternalDnsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-external-dns",
			},
			Spec: &externaldnskubernetesv1.ExternalDnsKubernetesSpec{
				TargetCluster: &kubernetes.KubernetesAddonTargetCluster{
					CredentialSource: &kubernetes.KubernetesAddonTargetCluster_KubernetesClusterCredentialId{
						KubernetesClusterCredentialId: "test-cred-id",
					},
				},
				// namespace, external_dns_version, and helm_chart_version are left unset
			},
		}

		// Apply defaults
		err := ApplyDefaults(msg)
		require.NoError(t, err)

		// Verify defaults were applied
		assert.Equal(t, "external-dns", msg.Spec.Namespace, "namespace should have default value")
		assert.Equal(t, "v0.19.0", msg.Spec.ExternalDnsVersion, "external_dns_version should have default value")
		assert.Equal(t, "1.19.0", msg.Spec.HelmChartVersion, "helm_chart_version should have default value")
	})

	t.Run("preserves existing values when field is already set", func(t *testing.T) {
		// Create a message with custom values
		msg := &externaldnskubernetesv1.ExternalDnsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ExternalDnsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-external-dns",
			},
			Spec: &externaldnskubernetesv1.ExternalDnsKubernetesSpec{
				TargetCluster: &kubernetes.KubernetesAddonTargetCluster{
					CredentialSource: &kubernetes.KubernetesAddonTargetCluster_KubernetesClusterCredentialId{
						KubernetesClusterCredentialId: "test-cred-id",
					},
				},
				Namespace:          "custom-namespace",
				ExternalDnsVersion: "v0.20.0",
				HelmChartVersion:   "2.0.0",
			},
		}

		// Apply defaults
		err := ApplyDefaults(msg)
		require.NoError(t, err)

		// Verify existing values were preserved
		assert.Equal(t, "custom-namespace", msg.Spec.Namespace)
		assert.Equal(t, "v0.20.0", msg.Spec.ExternalDnsVersion)
		assert.Equal(t, "2.0.0", msg.Spec.HelmChartVersion)
	})

	t.Run("handles partial values - some set, some unset", func(t *testing.T) {
		msg := &externaldnskubernetesv1.ExternalDnsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ExternalDnsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-external-dns",
			},
			Spec: &externaldnskubernetesv1.ExternalDnsKubernetesSpec{
				TargetCluster: &kubernetes.KubernetesAddonTargetCluster{
					CredentialSource: &kubernetes.KubernetesAddonTargetCluster_KubernetesClusterCredentialId{
						KubernetesClusterCredentialId: "test-cred-id",
					},
				},
				Namespace: "custom-namespace",
				// external_dns_version and helm_chart_version are left unset
			},
		}

		err := ApplyDefaults(msg)
		require.NoError(t, err)

		// Verify: custom value preserved, defaults applied to unset fields
		assert.Equal(t, "custom-namespace", msg.Spec.Namespace)
		assert.Equal(t, "v0.19.0", msg.Spec.ExternalDnsVersion)
		assert.Equal(t, "1.19.0", msg.Spec.HelmChartVersion)
	})

	t.Run("handles nil message gracefully", func(t *testing.T) {
		err := ApplyDefaults(nil)
		assert.NoError(t, err)
	})

	t.Run("handles nil spec gracefully", func(t *testing.T) {
		msg := &externaldnskubernetesv1.ExternalDnsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ExternalDnsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-external-dns",
			},
			Spec: nil,
		}

		err := ApplyDefaults(msg)
		assert.NoError(t, err)
	})
}

func TestApplyDefaults_NestedMessages(t *testing.T) {
	t.Run("applies defaults recursively to nested messages", func(t *testing.T) {
		// Create message with nested structure
		msg := &externaldnskubernetesv1.ExternalDnsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ExternalDnsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-external-dns",
			},
			Spec: &externaldnskubernetesv1.ExternalDnsKubernetesSpec{
				TargetCluster: &kubernetes.KubernetesAddonTargetCluster{
					CredentialSource: &kubernetes.KubernetesAddonTargetCluster_KubernetesClusterCredentialId{
						KubernetesClusterCredentialId: "test-cred-id",
					},
				},
				// Leave defaults unset at spec level
			},
		}

		err := ApplyDefaults(msg)
		require.NoError(t, err)

		// Verify defaults were applied at all levels
		assert.Equal(t, "external-dns", msg.Spec.Namespace)
		assert.Equal(t, "v0.19.0", msg.Spec.ExternalDnsVersion)
		assert.Equal(t, "1.19.0", msg.Spec.HelmChartVersion)
	})
}

func TestApplyDefaults_FieldsWithoutDefaults(t *testing.T) {
	t.Run("leaves fields without defaults unchanged", func(t *testing.T) {
		msg := &externaldnskubernetesv1.ExternalDnsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ExternalDnsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-external-dns",
			},
			Spec: &externaldnskubernetesv1.ExternalDnsKubernetesSpec{
				TargetCluster: &kubernetes.KubernetesAddonTargetCluster{
					CredentialSource: &kubernetes.KubernetesAddonTargetCluster_KubernetesClusterCredentialId{
						KubernetesClusterCredentialId: "test-cred-id",
					},
				},
				// namespace, version fields unset
				// provider_config (oneof) also unset - no defaults defined
			},
		}

		err := ApplyDefaults(msg)
		require.NoError(t, err)

		// Fields with defaults should be set
		assert.Equal(t, "external-dns", msg.Spec.Namespace)
		assert.Equal(t, "v0.19.0", msg.Spec.ExternalDnsVersion)

		// Provider config (no defaults) should remain nil
		assert.Nil(t, msg.Spec.GetGke())
		assert.Nil(t, msg.Spec.GetEks())
		assert.Nil(t, msg.Spec.GetAks())
		assert.Nil(t, msg.Spec.GetCloudflare())
	})
}

func TestApplyDefaults_Idempotency(t *testing.T) {
	t.Run("applying defaults multiple times is idempotent", func(t *testing.T) {
		msg := &externaldnskubernetesv1.ExternalDnsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ExternalDnsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-external-dns",
			},
			Spec: &externaldnskubernetesv1.ExternalDnsKubernetesSpec{
				TargetCluster: &kubernetes.KubernetesAddonTargetCluster{
					CredentialSource: &kubernetes.KubernetesAddonTargetCluster_KubernetesClusterCredentialId{
						KubernetesClusterCredentialId: "test-cred-id",
					},
				},
			},
		}

		// Apply defaults first time
		err := ApplyDefaults(msg)
		require.NoError(t, err)
		firstResult := proto.Clone(msg)

		// Apply defaults second time
		err = ApplyDefaults(msg)
		require.NoError(t, err)
		secondResult := msg

		// Results should be identical
		assert.True(t, proto.Equal(firstResult, secondResult),
			"Applying defaults multiple times should produce identical results")
	})
}

