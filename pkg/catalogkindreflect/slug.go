package catalogkindreflect

import (
	"strings"
	"unicode"

	"github.com/plantonhq/planton/shared/catalogkind"
)

// Slug returns the kebab-case URL segment every console addresses a kind by
// (AwsS3Bucket -> aws-s3-bucket, DigitalOceanDroplet -> digitalocean-droplet).
// It is the one address rule: the release bundle stamps it on each catalog
// entry, and consoles and the CLI derive it here rather than kebab-casing the
// kind name themselves, which splits a provider's own name (digital-ocean)
// and points at a page that does not exist.
func Slug(kind catalogkind.CatalogKind) string {
	return SlugOf(ExtractKindNameByKind(kind), ProviderDirName(GetProvider(kind)))
}

// SlugOf derives a kind's slug from its PascalCase name and its provider's
// catalog directory: the provider kept as one atomic word, then the rest of
// the name kebab-cased. A kind whose name does not start with its provider
// kebabs whole. Uniqueness across the catalog is gated at conformance.
func SlugOf(kindName, providerDir string) string {
	if rest, ok := strings.CutPrefix(strings.ToLower(kindName), providerDir); ok && rest != "" {
		return providerDir + "-" + kebabCase(kindName[len(kindName)-len(rest):])
	}
	return kebabCase(kindName)
}

// kebabCase converts a CamelCase name (S3Bucket) to kebab (s3-bucket): a
// hyphen lands before every uppercase rune that starts a new word -- after a
// lowercase or digit, or before an Upper->lower transition inside an acronym
// run.
func kebabCase(kindName string) string {
	runes := []rune(kindName)
	var b strings.Builder
	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) {
			prev := runes[i-1]
			nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) || (unicode.IsUpper(prev) && nextLower) {
				b.WriteRune('-')
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
