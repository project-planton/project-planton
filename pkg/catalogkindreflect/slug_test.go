package catalogkindreflect

import (
	"testing"

	"github.com/plantonhq/planton/shared/catalogkind"
)

// The slug rule is ONE rule for every kind -- the provider stays an atomic
// word, the rest kebabs. These pins hold the derivation still.
func TestSlugOf(t *testing.T) {
	cases := []struct {
		kindName, providerDir, want string
	}{
		{"AwsS3Bucket", "aws", "aws-s3-bucket"},
		{"AwsKmsKey", "aws", "aws-kms-key"},
		{"AwsMemorydbCluster", "aws", "aws-memorydb-cluster"},
		{"KubernetesDeployment", "kubernetes", "kubernetes-deployment"},
		{"DigitalOceanDroplet", "digitalocean", "digitalocean-droplet"},
		{"CloudflareDnsZone", "cloudflare", "cloudflare-dns-zone"},
		// A kind not prefixed by its provider kebabs whole.
		{"GkeCluster", "gcp", "gke-cluster"},
	}
	for _, c := range cases {
		if got := SlugOf(c.kindName, c.providerDir); got != c.want {
			t.Errorf("SlugOf(%s, %s) = %q, want %q", c.kindName, c.providerDir, got, c.want)
		}
	}
}

// A kind resolves its own provider directory: the providers whose names are
// two words in the enum keep them one word in the address.
func TestSlug(t *testing.T) {
	cases := map[catalogkind.CatalogKind]string{
		catalogkind.CatalogKind_AwsS3Bucket:     "aws-s3-bucket",
		catalogkind.CatalogKind_GcpGkeCluster:   "gcp-gke-cluster",
		catalogkind.CatalogKind_DigitalOceanApp: "digitalocean-app",
		catalogkind.CatalogKind_OpenFgaStore:    "openfga-store",
	}
	for kind, want := range cases {
		if got := Slug(kind); got != want {
			t.Errorf("Slug(%s) = %q, want %q", kind, got, want)
		}
	}
}
