package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	helmreleasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/helmrelease/v1"
)

func TestHelmReleaseSpec_ValidSpec(t *testing.T) {
	spec := &helmreleasev1.HelmReleaseSpec{
		Repo:    "https://charts.helm.sh/stable",
		Name:    "nginx-ingress",
		Version: "1.41.3",
		Values: map[string]string{
			"controller.replicaCount": "2",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

func TestHelmReleaseSpec_MissingRepo(t *testing.T) {
	spec := &helmreleasev1.HelmReleaseSpec{
		// Repo missing
		Name:    "nginx-ingress",
		Version: "1.41.3",
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing repo, got none")
	} else {
		if !strings.Contains(err.Error(), "repo") {
			t.Errorf("expected error mentioning 'repo' field, got: %v", err)
		}
	}
}

func TestHelmReleaseSpec_MissingName(t *testing.T) {
	spec := &helmreleasev1.HelmReleaseSpec{
		Repo: "https://charts.helm.sh/stable",
		// Name missing
		Version: "1.41.3",
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing name, got none")
	} else {
		if !strings.Contains(err.Error(), "name") {
			t.Errorf("expected error mentioning 'name' field, got: %v", err)
		}
	}
}

func TestHelmReleaseSpec_MissingVersion(t *testing.T) {
	spec := &helmreleasev1.HelmReleaseSpec{
		Repo: "https://charts.helm.sh/stable",
		Name: "nginx-ingress",
		// Version missing
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing version, got none")
	} else {
		if !strings.Contains(err.Error(), "version") {
			t.Errorf("expected error mentioning 'version' field, got: %v", err)
		}
	}
}
