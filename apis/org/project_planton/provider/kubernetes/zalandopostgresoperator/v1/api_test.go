package zalandopostgresoperatorv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestZalandoPostgresOperator(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ZalandoPostgresOperator Suite")
}

var _ = ginkgo.Describe("ZalandoPostgresOperator Custom Validation Tests", func() {
	var input *ZalandoPostgresOperator

	ginkgo.BeforeEach(func() {
		input = &ZalandoPostgresOperator{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ZalandoPostgresOperator",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-zalando-postgres-operator",
			},
			Spec: &ZalandoPostgresOperatorSpec{
				Container: &ZalandoPostgresOperatorSpecContainer{},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("zalando_postgres_operator", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When backup_config is provided with valid r2_config", func() {
		ginkgo.Context("zalando_postgres_operator with backup", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.BackupConfig = &ZalandoPostgresOperatorBackupConfig{
					R2Config: &ZalandoPostgresOperatorBackupR2Config{
						CloudflareAccountId: "test-account-id",
						BucketName:          "test-bucket",
						AccessKeyId:         "test-access-key",
						SecretAccessKey:     "test-secret-key",
					},
					BackupSchedule: "0 2 * * *",
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
