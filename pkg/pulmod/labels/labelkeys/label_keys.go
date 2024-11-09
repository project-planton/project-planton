package labelkeys

import (
	"fmt"
	"strings"
)

const (
	PlantonCloudDomain           = "planton.cloud"
	PlantonCloudDomainNormalized = "planton-cloud"
)

func WithDomainPrefix(label string) string {
	return fmt.Sprintf("%s/%s", PlantonCloudDomain, label)
}

// WithNormalizedDomainPrefix uses prefix that does not contain dots or slashes which
// are not accepted for label keys on gcp projects.
// underscore is used in place of slash.
func WithNormalizedDomainPrefix(labelKey string) string {
	return fmt.Sprintf("%s_%s", PlantonCloudDomainNormalized, labelKey)
}

// WithoutPrefix returns the label without the prefix.
// prometheus label rules explained in https://stackoverflow.com/a/71507356 are also processed while removing the prefix.
func WithoutPrefix(label string) string {
	l := strings.TrimPrefix(label, PlantonCloudDomain)
	//prometheus replaces dots with underscores
	l = strings.TrimPrefix(label, strings.ReplaceAll(PlantonCloudDomain, ".", "_"))
	l = strings.TrimPrefix(label, "/")
	//prometheus replaces slashes with underscores
	l = strings.TrimPrefix(label, "_")
	return l
}

// WithPrometheusFormat returns the label with prometheus transformation applied.
// prometheus label transformation rules explained in https://stackoverflow.com/a/71507356.
// ex: planton.cloud/company label gets transformed to planton_cloud_company
// rules: replace dots, slashes and hyphens with underscores.
func WithPrometheusFormat(label string) string {
	//replace all dots with underscores
	l := strings.ReplaceAll(label, ".", "_")
	//replace all slashes with underscores
	l = strings.ReplaceAll(l, "/", "_")
	//replace all hyphens with underscores
	l = strings.ReplaceAll(l, "-", "_")
	return l
}

//WARNING: logic is incorrect here. fix it before using it.
//// FromPrometheusFormat returns the label by undoing the prometheus transformation rules.
//// prometheus label transformation rules explained in https://stackoverflow.com/a/71507356.
//// ex: planton_cloud_company will be converted to planton.cloud/company
//// rules: replace "planton_cloud_" with "planton.cloud/"
//func FromPrometheusFormat(label string) string {
//	prometheusFormattedPrefix := strings.ReplaceAll(PlantonCloudDomain, ".", "_")
//	//replace prometheus formatted prefix with normal prefix
//	l := strings.ReplaceAll(label, prometheusFormattedPrefix, PlantonCloudDomain)
//	//replace all other underscores with slashes
//	l = strings.ReplaceAll(l, "_", "/")
//	return l
//}
