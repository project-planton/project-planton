// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// GithubAccessToken generates ghs_ accessToken
type GithubAccessTokenPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput     `pulumi:"metadata"`
	Spec     GithubAccessTokenSpecPatchPtrOutput `pulumi:"spec"`
}

// NewGithubAccessTokenPatch registers a new resource with the given unique name, arguments, and options.
func NewGithubAccessTokenPatch(ctx *pulumi.Context,
	name string, args *GithubAccessTokenPatchArgs, opts ...pulumi.ResourceOption) (*GithubAccessTokenPatch, error) {
	if args == nil {
		args = &GithubAccessTokenPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("GithubAccessToken")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GithubAccessTokenPatch
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:GithubAccessTokenPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGithubAccessTokenPatch gets an existing GithubAccessTokenPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGithubAccessTokenPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GithubAccessTokenPatchState, opts ...pulumi.ResourceOption) (*GithubAccessTokenPatch, error) {
	var resource GithubAccessTokenPatch
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:GithubAccessTokenPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GithubAccessTokenPatch resources.
type githubAccessTokenPatchState struct {
}

type GithubAccessTokenPatchState struct {
}

func (GithubAccessTokenPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*githubAccessTokenPatchState)(nil)).Elem()
}

type githubAccessTokenPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch     `pulumi:"metadata"`
	Spec     *GithubAccessTokenSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a GithubAccessTokenPatch resource.
type GithubAccessTokenPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     GithubAccessTokenSpecPatchPtrInput
}

func (GithubAccessTokenPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*githubAccessTokenPatchArgs)(nil)).Elem()
}

type GithubAccessTokenPatchInput interface {
	pulumi.Input

	ToGithubAccessTokenPatchOutput() GithubAccessTokenPatchOutput
	ToGithubAccessTokenPatchOutputWithContext(ctx context.Context) GithubAccessTokenPatchOutput
}

func (*GithubAccessTokenPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**GithubAccessTokenPatch)(nil)).Elem()
}

func (i *GithubAccessTokenPatch) ToGithubAccessTokenPatchOutput() GithubAccessTokenPatchOutput {
	return i.ToGithubAccessTokenPatchOutputWithContext(context.Background())
}

func (i *GithubAccessTokenPatch) ToGithubAccessTokenPatchOutputWithContext(ctx context.Context) GithubAccessTokenPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubAccessTokenPatchOutput)
}

// GithubAccessTokenPatchArrayInput is an input type that accepts GithubAccessTokenPatchArray and GithubAccessTokenPatchArrayOutput values.
// You can construct a concrete instance of `GithubAccessTokenPatchArrayInput` via:
//
//	GithubAccessTokenPatchArray{ GithubAccessTokenPatchArgs{...} }
type GithubAccessTokenPatchArrayInput interface {
	pulumi.Input

	ToGithubAccessTokenPatchArrayOutput() GithubAccessTokenPatchArrayOutput
	ToGithubAccessTokenPatchArrayOutputWithContext(context.Context) GithubAccessTokenPatchArrayOutput
}

type GithubAccessTokenPatchArray []GithubAccessTokenPatchInput

func (GithubAccessTokenPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GithubAccessTokenPatch)(nil)).Elem()
}

func (i GithubAccessTokenPatchArray) ToGithubAccessTokenPatchArrayOutput() GithubAccessTokenPatchArrayOutput {
	return i.ToGithubAccessTokenPatchArrayOutputWithContext(context.Background())
}

func (i GithubAccessTokenPatchArray) ToGithubAccessTokenPatchArrayOutputWithContext(ctx context.Context) GithubAccessTokenPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubAccessTokenPatchArrayOutput)
}

// GithubAccessTokenPatchMapInput is an input type that accepts GithubAccessTokenPatchMap and GithubAccessTokenPatchMapOutput values.
// You can construct a concrete instance of `GithubAccessTokenPatchMapInput` via:
//
//	GithubAccessTokenPatchMap{ "key": GithubAccessTokenPatchArgs{...} }
type GithubAccessTokenPatchMapInput interface {
	pulumi.Input

	ToGithubAccessTokenPatchMapOutput() GithubAccessTokenPatchMapOutput
	ToGithubAccessTokenPatchMapOutputWithContext(context.Context) GithubAccessTokenPatchMapOutput
}

type GithubAccessTokenPatchMap map[string]GithubAccessTokenPatchInput

func (GithubAccessTokenPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GithubAccessTokenPatch)(nil)).Elem()
}

func (i GithubAccessTokenPatchMap) ToGithubAccessTokenPatchMapOutput() GithubAccessTokenPatchMapOutput {
	return i.ToGithubAccessTokenPatchMapOutputWithContext(context.Background())
}

func (i GithubAccessTokenPatchMap) ToGithubAccessTokenPatchMapOutputWithContext(ctx context.Context) GithubAccessTokenPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubAccessTokenPatchMapOutput)
}

type GithubAccessTokenPatchOutput struct{ *pulumi.OutputState }

func (GithubAccessTokenPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GithubAccessTokenPatch)(nil)).Elem()
}

func (o GithubAccessTokenPatchOutput) ToGithubAccessTokenPatchOutput() GithubAccessTokenPatchOutput {
	return o
}

func (o GithubAccessTokenPatchOutput) ToGithubAccessTokenPatchOutputWithContext(ctx context.Context) GithubAccessTokenPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o GithubAccessTokenPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubAccessTokenPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o GithubAccessTokenPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubAccessTokenPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o GithubAccessTokenPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *GithubAccessTokenPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o GithubAccessTokenPatchOutput) Spec() GithubAccessTokenSpecPatchPtrOutput {
	return o.ApplyT(func(v *GithubAccessTokenPatch) GithubAccessTokenSpecPatchPtrOutput { return v.Spec }).(GithubAccessTokenSpecPatchPtrOutput)
}

type GithubAccessTokenPatchArrayOutput struct{ *pulumi.OutputState }

func (GithubAccessTokenPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GithubAccessTokenPatch)(nil)).Elem()
}

func (o GithubAccessTokenPatchArrayOutput) ToGithubAccessTokenPatchArrayOutput() GithubAccessTokenPatchArrayOutput {
	return o
}

func (o GithubAccessTokenPatchArrayOutput) ToGithubAccessTokenPatchArrayOutputWithContext(ctx context.Context) GithubAccessTokenPatchArrayOutput {
	return o
}

func (o GithubAccessTokenPatchArrayOutput) Index(i pulumi.IntInput) GithubAccessTokenPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GithubAccessTokenPatch {
		return vs[0].([]*GithubAccessTokenPatch)[vs[1].(int)]
	}).(GithubAccessTokenPatchOutput)
}

type GithubAccessTokenPatchMapOutput struct{ *pulumi.OutputState }

func (GithubAccessTokenPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GithubAccessTokenPatch)(nil)).Elem()
}

func (o GithubAccessTokenPatchMapOutput) ToGithubAccessTokenPatchMapOutput() GithubAccessTokenPatchMapOutput {
	return o
}

func (o GithubAccessTokenPatchMapOutput) ToGithubAccessTokenPatchMapOutputWithContext(ctx context.Context) GithubAccessTokenPatchMapOutput {
	return o
}

func (o GithubAccessTokenPatchMapOutput) MapIndex(k pulumi.StringInput) GithubAccessTokenPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GithubAccessTokenPatch {
		return vs[0].(map[string]*GithubAccessTokenPatch)[vs[1].(string)]
	}).(GithubAccessTokenPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GithubAccessTokenPatchInput)(nil)).Elem(), &GithubAccessTokenPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*GithubAccessTokenPatchArrayInput)(nil)).Elem(), GithubAccessTokenPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GithubAccessTokenPatchMapInput)(nil)).Elem(), GithubAccessTokenPatchMap{})
	pulumi.RegisterOutputType(GithubAccessTokenPatchOutput{})
	pulumi.RegisterOutputType(GithubAccessTokenPatchArrayOutput{})
	pulumi.RegisterOutputType(GithubAccessTokenPatchMapOutput{})
}