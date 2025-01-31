package credential

import (
	"github.com/bufbuild/protovalidate-go"
	pulumibackendcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/pulumibackendcredential/v1"
	"strings"
	"testing"
)

func TestInitiate_PulumiBackendSpec_Without_HttpDetails(t *testing.T) {
	pulumiBackendCredentialSpec := &pulumibackendcredentialv1.PulumiBackendCredentialSpec{
		PulumiBackendType: pulumibackendcredentialv1.PulumiBackendType_http,
	}
	if err := protovalidate.Validate(pulumiBackendCredentialSpec); err != nil {
		if !strings.Contains(err.Error(), "[http.mandatory]") {
			t.Errorf("test failed - validation error of constraint id `http.mandatory` is expected")
		}
	} else {
		t.Errorf("test failed - validation error expected")
	}
}

func TestInitiate_PulumiBackendSpec_Without_AwsS3Details(t *testing.T) {
	pulumiBackendCredentialSpec := &pulumibackendcredentialv1.PulumiBackendCredentialSpec{
		PulumiBackendType: pulumibackendcredentialv1.PulumiBackendType_s3,
	}
	if err := protovalidate.Validate(pulumiBackendCredentialSpec); err != nil {
		if !strings.Contains(err.Error(), "[s3.mandatory]") {
			t.Errorf("test failed - validation error of constraint id `aws_s3.mandatory` is expected")
		}
	} else {
		t.Errorf("test failed - validation error expected")
	}
}

func TestInitiate_PulumiBackendSpec_Without_GoogleCloudStorageDetails(t *testing.T) {
	pulumiBackendCredentialSpec := &pulumibackendcredentialv1.PulumiBackendCredentialSpec{
		PulumiBackendType: pulumibackendcredentialv1.PulumiBackendType_gcs,
	}
	if err := protovalidate.Validate(pulumiBackendCredentialSpec); err != nil {
		if !strings.Contains(err.Error(), "[gcs.mandatory]") {
			t.Errorf("test failed - validation error of constraint id `google_cloud_storage.mandatory` is expected")
		}
	} else {
		t.Errorf("test failed - validation error expected")
	}
}

func TestInitiate_PulumiBackendSpec_Without_AzureBlobStorageDetails(t *testing.T) {
	pulumiBackendCredentialSpec := &pulumibackendcredentialv1.PulumiBackendCredentialSpec{
		PulumiBackendType: pulumibackendcredentialv1.PulumiBackendType_azurerm,
	}
	if err := protovalidate.Validate(pulumiBackendCredentialSpec); err != nil {
		if !strings.Contains(err.Error(), "[azurerm.mandatory]") {
			t.Errorf("test failed - validation error of constraint id `azure_blob_storage.mandatory` is expected")
		}
	} else {
		t.Errorf("test failed - validation error expected")
	}
}
