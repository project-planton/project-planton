package snowflakedatabasev1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestSnowflakeDatabase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SnowflakeDatabase Suite")
}

var _ = Describe("SnowflakeDatabase Custom Validation Tests", func() {
	var input *SnowflakeDatabase

	BeforeEach(func() {
		input = &SnowflakeDatabase{
			ApiVersion: "snowflake.project-planton.org/v1",
			Kind:       "SnowflakeDatabase",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-snowflake-db",
			},
			Spec: &SnowflakeDatabaseSpec{
				Name: "my_database_name",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("snowflake_database", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
