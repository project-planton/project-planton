// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:gateway.networking.k8s.io/v1alpha2:GRPCRoute":
		r = &GRPCRoute{}
	case "kubernetes:gateway.networking.k8s.io/v1alpha2:GRPCRouteList":
		r = &GRPCRouteList{}
	case "kubernetes:gateway.networking.k8s.io/v1alpha2:GRPCRoutePatch":
		r = &GRPCRoutePatch{}
	case "kubernetes:gateway.networking.k8s.io/v1alpha2:ReferenceGrant":
		r = &ReferenceGrant{}
	case "kubernetes:gateway.networking.k8s.io/v1alpha2:ReferenceGrantList":
		r = &ReferenceGrantList{}
	case "kubernetes:gateway.networking.k8s.io/v1alpha2:ReferenceGrantPatch":
		r = &ReferenceGrantPatch{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"crds",
		"gateway.networking.k8s.io/v1alpha2",
		&module{version},
	)
}