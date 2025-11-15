package gcpcloudsqlv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGcpCloudSqlSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCloudSqlSpec Validation Suite")
}

var _ = Describe("GcpCloudSqlSpec validations", func() {

	// Helper function to create a minimal valid spec
	makeValidSpec := func() *GcpCloudSqlSpec {
		return &GcpCloudSqlSpec{
			ProjectId:       "my-gcp-project",
			Region:          "us-central1",
			DatabaseEngine:  GcpCloudSqlDatabaseEngine_MYSQL,
			DatabaseVersion: "MYSQL_8_0",
			Tier:            "db-n1-standard-1",
			StorageGb:       10,
		}
	}

	Context("Required fields", func() {
		It("accepts a minimal valid spec", func() {
			spec := makeValidSpec()
			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects spec with missing project_id", func() {
			spec := makeValidSpec()
			spec.ProjectId = ""
			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects spec with missing region", func() {
			spec := makeValidSpec()
			spec.Region = ""
			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects spec with missing database_engine", func() {
			spec := makeValidSpec()
			spec.DatabaseEngine = GcpCloudSqlDatabaseEngine_DATABASE_ENGINE_UNSPECIFIED
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects spec with missing database_version", func() {
			spec := makeValidSpec()
			spec.DatabaseVersion = ""
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects spec with missing tier", func() {
			spec := makeValidSpec()
			spec.Tier = ""
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects spec with storage_gb = 0", func() {
			spec := makeValidSpec()
			spec.StorageGb = 0
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Project ID validation", func() {
		It("accepts valid project ID format", func() {
			spec := makeValidSpec()
			spec.ProjectId = "my-project-123"
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects project ID starting with number", func() {
			spec := makeValidSpec()
			spec.ProjectId = "123-project"
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects project ID ending with hyphen", func() {
			spec := makeValidSpec()
			spec.ProjectId = "my-project-"
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects project ID that is too short", func() {
			spec := makeValidSpec()
			spec.ProjectId = "proj"
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects project ID with uppercase letters", func() {
			spec := makeValidSpec()
			spec.ProjectId = "My-Project"
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Region validation", func() {
		It("accepts valid region format", func() {
			spec := makeValidSpec()
			spec.Region = "us-west1"
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts valid multi-region format", func() {
			spec := makeValidSpec()
			spec.Region = "europe-west2"
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects invalid region format without number", func() {
			spec := makeValidSpec()
			spec.Region = "us-central"
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects invalid region format with uppercase", func() {
			spec := makeValidSpec()
			spec.Region = "US-CENTRAL1"
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Database engine validation", func() {
		It("accepts MYSQL engine", func() {
			spec := makeValidSpec()
			spec.DatabaseEngine = GcpCloudSqlDatabaseEngine_MYSQL
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts POSTGRESQL engine", func() {
			spec := makeValidSpec()
			spec.DatabaseEngine = GcpCloudSqlDatabaseEngine_POSTGRESQL
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects unspecified engine", func() {
			spec := makeValidSpec()
			spec.DatabaseEngine = GcpCloudSqlDatabaseEngine_DATABASE_ENGINE_UNSPECIFIED
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Storage GB validation", func() {
		It("accepts minimum storage (10 GB)", func() {
			spec := makeValidSpec()
			spec.StorageGb = 10
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts maximum storage (65536 GB)", func() {
			spec := makeValidSpec()
			spec.StorageGb = 65536
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts mid-range storage (500 GB)", func() {
			spec := makeValidSpec()
			spec.StorageGb = 500
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects storage below minimum (9 GB)", func() {
			spec := makeValidSpec()
			spec.StorageGb = 9
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects storage above maximum (65537 GB)", func() {
			spec := makeValidSpec()
			spec.StorageGb = 65537
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Edition validation", func() {
		It("accepts ENTERPRISE edition", func() {
			spec := makeValidSpec()
			spec.Edition = GcpCloudSqlEdition_ENTERPRISE
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts ENTERPRISE_PLUS edition", func() {
			spec := makeValidSpec()
			spec.Edition = GcpCloudSqlEdition_ENTERPRISE_PLUS
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts EDITION_UNSPECIFIED (defaults to ENTERPRISE)", func() {
			spec := makeValidSpec()
			spec.Edition = GcpCloudSqlEdition_EDITION_UNSPECIFIED
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})
	})

	Context("Root password validation", func() {
		It("accepts password with minimum length (8 chars)", func() {
			spec := makeValidSpec()
			spec.RootPassword = "Password123"
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects password shorter than 8 characters", func() {
			spec := makeValidSpec()
			spec.RootPassword = "Pass123"
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts spec without password (optional)", func() {
			spec := makeValidSpec()
			spec.RootPassword = ""
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})
	})

	Context("Network configuration - CEL validation", func() {
		It("accepts private IP with VPC ID", func() {
			spec := makeValidSpec()
			spec.Network = &GcpCloudSqlNetwork{
				VpcId:            "projects/my-project/global/networks/my-vpc",
				PrivateIpEnabled: true,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects private IP without VPC ID (CEL rule)", func() {
			spec := makeValidSpec()
			spec.Network = &GcpCloudSqlNetwork{
				VpcId:            "",
				PrivateIpEnabled: true,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts public IP enabled (ipv4_enabled)", func() {
			spec := makeValidSpec()
			spec.Network = &GcpCloudSqlNetwork{
				Ipv4Enabled: true,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts Smart Hybrid pattern (private + public IP)", func() {
			spec := makeValidSpec()
			spec.Network = &GcpCloudSqlNetwork{
				VpcId:            "projects/my-project/global/networks/my-vpc",
				PrivateIpEnabled: true,
				Ipv4Enabled:      true,
				AuthorizedNetworks: []string{},
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts valid CIDR in authorized_networks", func() {
			spec := makeValidSpec()
			spec.Network = &GcpCloudSqlNetwork{
				Ipv4Enabled: true,
				AuthorizedNetworks: []string{
					"10.0.0.0/24",
					"192.168.1.0/16",
				},
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects invalid CIDR format in authorized_networks", func() {
			spec := makeValidSpec()
			spec.Network = &GcpCloudSqlNetwork{
				Ipv4Enabled: true,
				AuthorizedNetworks: []string{
					"10.0.0.0",
				},
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects duplicate CIDR in authorized_networks", func() {
			spec := makeValidSpec()
			spec.Network = &GcpCloudSqlNetwork{
				Ipv4Enabled: true,
				AuthorizedNetworks: []string{
					"10.0.0.0/24",
					"10.0.0.0/24",
				},
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("High availability - CEL validation", func() {
		It("accepts HA enabled with zone specified", func() {
			spec := makeValidSpec()
			spec.HighAvailability = &GcpCloudSqlHighAvailability{
				Enabled: true,
				Zone:    "us-central1-a",
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects HA enabled without zone (CEL rule)", func() {
			spec := makeValidSpec()
			spec.HighAvailability = &GcpCloudSqlHighAvailability{
				Enabled: true,
				Zone:    "",
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts HA disabled without zone", func() {
			spec := makeValidSpec()
			spec.HighAvailability = &GcpCloudSqlHighAvailability{
				Enabled: false,
				Zone:    "",
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})
	})

	Context("Backup configuration - CEL validations", func() {
		It("accepts backup enabled with start_time and retention_days", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "02:00",
				RetentionDays: 7,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects backup enabled without start_time (CEL rule)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "",
				RetentionDays: 7,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects backup enabled without retention_days (CEL rule)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "02:00",
				RetentionDays: 0,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts valid start_time format (HH:MM)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "23:59",
				RetentionDays: 30,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects invalid start_time format (missing leading zero)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "2:00",
				RetentionDays: 7,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects start_time with invalid hour (24)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "24:00",
				RetentionDays: 7,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts retention_days at minimum (1 day)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "02:00",
				RetentionDays: 1,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts retention_days at maximum (365 days)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "02:00",
				RetentionDays: 365,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects retention_days above maximum (366 days)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:       true,
				StartTime:     "02:00",
				RetentionDays: 366,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts PITR enabled when backup is enabled", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:                     true,
				StartTime:                   "02:00",
				RetentionDays:               7,
				PointInTimeRecoveryEnabled:  true,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects PITR enabled when backup is disabled (CEL rule)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled:                     false,
				PointInTimeRecoveryEnabled:  true,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts backup disabled (no config required)", func() {
			spec := makeValidSpec()
			spec.Backup = &GcpCloudSqlBackup{
				Enabled: false,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})
	})

	Context("Maintenance window validation", func() {
		It("accepts valid maintenance window", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:         1, // Monday
				Hour:        3, // 3 AM UTC
				UpdateTrack: "stable",
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts day at minimum (1 = Monday)", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:  1,
				Hour: 0,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts day at maximum (7 = Sunday)", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:  7,
				Hour: 0,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects day below minimum (0)", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:  0,
				Hour: 3,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("rejects day above maximum (8)", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:  8,
				Hour: 3,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts hour at minimum (0)", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:  1,
				Hour: 0,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts hour at maximum (23)", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:  1,
				Hour: 23,
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects hour above maximum (24)", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:  1,
				Hour: 24,
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})

		It("accepts canary update track", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:         1,
				Hour:        3,
				UpdateTrack: "canary",
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts stable update track", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:         1,
				Hour:        3,
				UpdateTrack: "stable",
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts empty update track (defaults)", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:         1,
				Hour:        3,
				UpdateTrack: "",
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("rejects invalid update track", func() {
			spec := makeValidSpec()
			spec.MaintenanceWindow = &GcpCloudSqlMaintenanceWindow{
				Day:         1,
				Hour:        3,
				UpdateTrack: "invalid",
			}
			err := validator.Validate(spec)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Database flags", func() {
		It("accepts database flags", func() {
			spec := makeValidSpec()
			spec.DatabaseFlags = map[string]string{
				"slow_query_log": "on",
				"log_output":     "FILE",
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})

		It("accepts empty database flags", func() {
			spec := makeValidSpec()
			spec.DatabaseFlags = map[string]string{}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})
	})

	Context("Production-grade configuration example", func() {
		It("accepts a complete production spec with all options", func() {
			spec := &GcpCloudSqlSpec{
				ProjectId:            "prod-database-project",
				Region:               "us-central1",
				DatabaseEngine:       GcpCloudSqlDatabaseEngine_POSTGRESQL,
				DatabaseVersion:      "POSTGRES_15",
				Tier:                 "db-n1-standard-4",
				StorageGb:            100,
				DiskAutoresize:       true,
				Edition:              GcpCloudSqlEdition_ENTERPRISE,
				DeletionProtection:   true,
				QueryInsightsEnabled: true,
				MaintenanceWindow: &GcpCloudSqlMaintenanceWindow{
					Day:         7, // Sunday
					Hour:        2, // 2 AM UTC
					UpdateTrack: "stable",
				},
				Network: &GcpCloudSqlNetwork{
					VpcId:              "projects/prod-project/global/networks/prod-vpc",
					PrivateIpEnabled:   true,
					Ipv4Enabled:        true,
					AuthorizedNetworks: []string{},
				},
				HighAvailability: &GcpCloudSqlHighAvailability{
					Enabled: true,
					Zone:    "us-central1-b",
				},
				Backup: &GcpCloudSqlBackup{
					Enabled:                    true,
					StartTime:                  "03:00",
					RetentionDays:              30,
					PointInTimeRecoveryEnabled: true,
				},
				DatabaseFlags: map[string]string{
					"log_min_duration_statement": "1000",
					"log_statement":              "ddl",
				},
				RootPassword: "SuperSecurePassword123!",
			}
			err := validator.Validate(spec)
			Expect(err).To(BeNil())
		})
	})
})
