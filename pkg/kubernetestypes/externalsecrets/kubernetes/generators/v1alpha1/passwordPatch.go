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
// Password generates a random password based on the
// configuration parameters in spec.
// You can specify the length, characterset and other attributes.
type PasswordPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     PasswordSpecPatchPtrOutput      `pulumi:"spec"`
}

// NewPasswordPatch registers a new resource with the given unique name, arguments, and options.
func NewPasswordPatch(ctx *pulumi.Context,
	name string, args *PasswordPatchArgs, opts ...pulumi.ResourceOption) (*PasswordPatch, error) {
	if args == nil {
		args = &PasswordPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("Password")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PasswordPatch
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:PasswordPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPasswordPatch gets an existing PasswordPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPasswordPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PasswordPatchState, opts ...pulumi.ResourceOption) (*PasswordPatch, error) {
	var resource PasswordPatch
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:PasswordPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PasswordPatch resources.
type passwordPatchState struct {
}

type PasswordPatchState struct {
}

func (PasswordPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordPatchState)(nil)).Elem()
}

type passwordPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *PasswordSpecPatch      `pulumi:"spec"`
}

// The set of arguments for constructing a PasswordPatch resource.
type PasswordPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     PasswordSpecPatchPtrInput
}

func (PasswordPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordPatchArgs)(nil)).Elem()
}

type PasswordPatchInput interface {
	pulumi.Input

	ToPasswordPatchOutput() PasswordPatchOutput
	ToPasswordPatchOutputWithContext(ctx context.Context) PasswordPatchOutput
}

func (*PasswordPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordPatch)(nil)).Elem()
}

func (i *PasswordPatch) ToPasswordPatchOutput() PasswordPatchOutput {
	return i.ToPasswordPatchOutputWithContext(context.Background())
}

func (i *PasswordPatch) ToPasswordPatchOutputWithContext(ctx context.Context) PasswordPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordPatchOutput)
}

// PasswordPatchArrayInput is an input type that accepts PasswordPatchArray and PasswordPatchArrayOutput values.
// You can construct a concrete instance of `PasswordPatchArrayInput` via:
//
//	PasswordPatchArray{ PasswordPatchArgs{...} }
type PasswordPatchArrayInput interface {
	pulumi.Input

	ToPasswordPatchArrayOutput() PasswordPatchArrayOutput
	ToPasswordPatchArrayOutputWithContext(context.Context) PasswordPatchArrayOutput
}

type PasswordPatchArray []PasswordPatchInput

func (PasswordPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordPatch)(nil)).Elem()
}

func (i PasswordPatchArray) ToPasswordPatchArrayOutput() PasswordPatchArrayOutput {
	return i.ToPasswordPatchArrayOutputWithContext(context.Background())
}

func (i PasswordPatchArray) ToPasswordPatchArrayOutputWithContext(ctx context.Context) PasswordPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordPatchArrayOutput)
}

// PasswordPatchMapInput is an input type that accepts PasswordPatchMap and PasswordPatchMapOutput values.
// You can construct a concrete instance of `PasswordPatchMapInput` via:
//
//	PasswordPatchMap{ "key": PasswordPatchArgs{...} }
type PasswordPatchMapInput interface {
	pulumi.Input

	ToPasswordPatchMapOutput() PasswordPatchMapOutput
	ToPasswordPatchMapOutputWithContext(context.Context) PasswordPatchMapOutput
}

type PasswordPatchMap map[string]PasswordPatchInput

func (PasswordPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordPatch)(nil)).Elem()
}

func (i PasswordPatchMap) ToPasswordPatchMapOutput() PasswordPatchMapOutput {
	return i.ToPasswordPatchMapOutputWithContext(context.Background())
}

func (i PasswordPatchMap) ToPasswordPatchMapOutputWithContext(ctx context.Context) PasswordPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordPatchMapOutput)
}

type PasswordPatchOutput struct{ *pulumi.OutputState }

func (PasswordPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordPatch)(nil)).Elem()
}

func (o PasswordPatchOutput) ToPasswordPatchOutput() PasswordPatchOutput {
	return o
}

func (o PasswordPatchOutput) ToPasswordPatchOutputWithContext(ctx context.Context) PasswordPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PasswordPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PasswordPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PasswordPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PasswordPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PasswordPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *PasswordPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o PasswordPatchOutput) Spec() PasswordSpecPatchPtrOutput {
	return o.ApplyT(func(v *PasswordPatch) PasswordSpecPatchPtrOutput { return v.Spec }).(PasswordSpecPatchPtrOutput)
}

type PasswordPatchArrayOutput struct{ *pulumi.OutputState }

func (PasswordPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordPatch)(nil)).Elem()
}

func (o PasswordPatchArrayOutput) ToPasswordPatchArrayOutput() PasswordPatchArrayOutput {
	return o
}

func (o PasswordPatchArrayOutput) ToPasswordPatchArrayOutputWithContext(ctx context.Context) PasswordPatchArrayOutput {
	return o
}

func (o PasswordPatchArrayOutput) Index(i pulumi.IntInput) PasswordPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PasswordPatch {
		return vs[0].([]*PasswordPatch)[vs[1].(int)]
	}).(PasswordPatchOutput)
}

type PasswordPatchMapOutput struct{ *pulumi.OutputState }

func (PasswordPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordPatch)(nil)).Elem()
}

func (o PasswordPatchMapOutput) ToPasswordPatchMapOutput() PasswordPatchMapOutput {
	return o
}

func (o PasswordPatchMapOutput) ToPasswordPatchMapOutputWithContext(ctx context.Context) PasswordPatchMapOutput {
	return o
}

func (o PasswordPatchMapOutput) MapIndex(k pulumi.StringInput) PasswordPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PasswordPatch {
		return vs[0].(map[string]*PasswordPatch)[vs[1].(string)]
	}).(PasswordPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordPatchInput)(nil)).Elem(), &PasswordPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordPatchArrayInput)(nil)).Elem(), PasswordPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordPatchMapInput)(nil)).Elem(), PasswordPatchMap{})
	pulumi.RegisterOutputType(PasswordPatchOutput{})
	pulumi.RegisterOutputType(PasswordPatchArrayOutput{})
	pulumi.RegisterOutputType(PasswordPatchMapOutput{})
}
