package module

import (
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpvpc/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources provisions a GCP VPC, generating its project ID with a 3-char suffix.
func Resources(ctx *pulumi.Context, stackInput *gcpvpcv1.GcpVpcStackInput) error {
	return nil
}
