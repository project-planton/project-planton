package snowflakedatabasev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestSnowflakeDatabaseSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SnowflakeDatabaseSpec Custom Validation Tests")
}

var _ = Describe("SnowflakeDatabaseSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("snowflake_database", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &SnowflakeDatabase{
					ApiVersion: "snowflake.project-planton.org/v1",
					Kind:       "SnowflakeDatabase",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-snowflake-database",
					},
					Spec: &SnowflakeDatabaseSpec{
						// No required fields, all are optional
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
