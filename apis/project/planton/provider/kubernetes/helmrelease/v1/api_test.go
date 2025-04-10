package helmreleasev1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
)

func TestHelmReleaseSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HelmReleaseSpec Suite")
}

var _ = Describe("HelmReleaseSpec", func() {
	Context("with a valid spec", func() {
		It("should not return any validation errors", func() {
			spec := &HelmReleaseSpec{
				Repo:    "https://charts.helm.sh/stable",
				Name:    "nginx-ingress",
				Version: "1.41.3",
				Values: map[string]string{
					"controller.replicaCount": "2",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors")
		})
	})

	Context("when the repo field is missing", func() {
		It("should return an error mentioning 'repo'", func() {
			spec := &HelmReleaseSpec{
				// Repo missing
				Name:    "nginx-ingress",
				Version: "1.41.3",
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for missing repo")
			Expect(err.Error()).To(ContainSubstring("repo"), "expected error mentioning 'repo' field")
		})
	})

	Context("when the name field is missing", func() {
		It("should return an error mentioning 'name'", func() {
			spec := &HelmReleaseSpec{
				Repo: "https://charts.helm.sh/stable",
				// Name missing
				Version: "1.41.3",
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for missing name")
			Expect(err.Error()).To(ContainSubstring("name"), "expected error mentioning 'name' field")
		})
	})

	Context("when the version field is missing", func() {
		It("should return an error mentioning 'version'", func() {
			spec := &HelmReleaseSpec{
				Repo: "https://charts.helm.sh/stable",
				Name: "nginx-ingress",
				// Version missing
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for missing version")
			Expect(err.Error()).To(ContainSubstring("version"), "expected error mentioning 'version' field")
		})
	})
})
