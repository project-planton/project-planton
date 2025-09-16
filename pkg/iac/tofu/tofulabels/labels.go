package tofulabels

const (
	// BackendTypeLabelKey specifies the backend type (e.g., "s3", "gcs", "azurerm")
	BackendTypeLabelKey = "terraform.project-planton.org/backend.type"

	// BackendObjectLabelKey specifies the backend object path
	// For S3: "bucket-name/path/to/state"
	// For GCS: "bucket-name/path/to/state"
	// For Azure: "container-name/path/to/state"
	BackendObjectLabelKey = "terraform.project-planton.org/backend.object"
)
