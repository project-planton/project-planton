package credential

import (
	"github.com/bufbuild/protovalidate-go"
	pulumibackendcredentialv1 "github.com/project-planton/project-planton/apis/go/project/planton/credential/pulumibackendcredential/v1"
	"strings"
	"testing"
)

func TestInitiate_PulumiBackendSpec_Without_LocalFileSystemDetails(t *testing.T) {
	pulumiBackendCredentialSpec := &pulumibackendcredentialv1.PulumiBackendCredentialSpec{
		PulumiBackendType: pulumibackendcredentialv1.PulumiBackendType_local_file_system,
	}
	if err := protovalidate.Validate(pulumiBackendCredentialSpec); err != nil {
		if !strings.Contains(err.Error(), "[local_file_system.mandatory]") {
			t.Errorf("test failed - validation error of constraint id `local_file_system.mandatory` is expected")
		}
	} else {
		t.Errorf("test failed - validation error expected")
	}
}

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
		PulumiBackendType: pulumibackendcredentialv1.PulumiBackendType_aws_s3,
	}
	if err := protovalidate.Validate(pulumiBackendCredentialSpec); err != nil {
		if !strings.Contains(err.Error(), "[aws_s3.mandatory]") {
			t.Errorf("test failed - validation error of constraint id `aws_s3.mandatory` is expected")
		}
	} else {
		t.Errorf("test failed - validation error expected")
	}
}

func TestInitiate_PulumiBackendSpec_Without_GoogleCloudStorageDetails(t *testing.T) {
	pulumiBackendCredentialSpec := &pulumibackendcredentialv1.PulumiBackendCredentialSpec{
		PulumiBackendType: pulumibackendcredentialv1.PulumiBackendType_google_cloud_storage,
	}
	if err := protovalidate.Validate(pulumiBackendCredentialSpec); err != nil {
		if !strings.Contains(err.Error(), "[google_cloud_storage.mandatory]") {
			t.Errorf("test failed - validation error of constraint id `google_cloud_storage.mandatory` is expected")
		}
	} else {
		t.Errorf("test failed - validation error expected")
	}
}

func TestInitiate_PulumiBackendSpec_Without_AzureBlobStorageDetails(t *testing.T) {
	pulumiBackendCredentialSpec := &pulumibackendcredentialv1.PulumiBackendCredentialSpec{
		PulumiBackendType: pulumibackendcredentialv1.PulumiBackendType_azure_blob_storage,
	}
	if err := protovalidate.Validate(pulumiBackendCredentialSpec); err != nil {
		if !strings.Contains(err.Error(), "[azure_blob_storage.mandatory]") {
			t.Errorf("test failed - validation error of constraint id `azure_blob_storage.mandatory` is expected")
		}
	} else {
		t.Errorf("test failed - validation error expected")
	}
}
