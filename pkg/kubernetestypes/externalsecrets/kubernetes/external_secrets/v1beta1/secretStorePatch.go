// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
// SecretStore represents a secure external location for storing secrets, which can be referenced as part of `storeRef` fields.
type SecretStorePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     SecretStoreSpecPatchPtrOutput   `pulumi:"spec"`
	Status   SecretStoreStatusPatchPtrOutput `pulumi:"status"`
}

// NewSecretStorePatch registers a new resource with the given unique name, arguments, and options.
func NewSecretStorePatch(ctx *pulumi.Context,
	name string, args *SecretStorePatchArgs, opts ...pulumi.ResourceOption) (*SecretStorePatch, error) {
	if args == nil {
		args = &SecretStorePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("external-secrets.io/v1beta1")
	args.Kind = pulumi.StringPtr("SecretStore")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:external-secrets.io/v1alpha1:SecretStorePatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecretStorePatch
	err := ctx.RegisterResource("kubernetes:external-secrets.io/v1beta1:SecretStorePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretStorePatch gets an existing SecretStorePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretStorePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretStorePatchState, opts ...pulumi.ResourceOption) (*SecretStorePatch, error) {
	var resource SecretStorePatch
	err := ctx.ReadResource("kubernetes:external-secrets.io/v1beta1:SecretStorePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretStorePatch resources.
type secretStorePatchState struct {
}

type SecretStorePatchState struct {
}

func (SecretStorePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretStorePatchState)(nil)).Elem()
}

type secretStorePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *SecretStoreSpecPatch   `pulumi:"spec"`
}

// The set of arguments for constructing a SecretStorePatch resource.
type SecretStorePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     SecretStoreSpecPatchPtrInput
}

func (SecretStorePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretStorePatchArgs)(nil)).Elem()
}

type SecretStorePatchInput interface {
	pulumi.Input

	ToSecretStorePatchOutput() SecretStorePatchOutput
	ToSecretStorePatchOutputWithContext(ctx context.Context) SecretStorePatchOutput
}

func (*SecretStorePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStorePatch)(nil)).Elem()
}

func (i *SecretStorePatch) ToSecretStorePatchOutput() SecretStorePatchOutput {
	return i.ToSecretStorePatchOutputWithContext(context.Background())
}

func (i *SecretStorePatch) ToSecretStorePatchOutputWithContext(ctx context.Context) SecretStorePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStorePatchOutput)
}

// SecretStorePatchArrayInput is an input type that accepts SecretStorePatchArray and SecretStorePatchArrayOutput values.
// You can construct a concrete instance of `SecretStorePatchArrayInput` via:
//
//	SecretStorePatchArray{ SecretStorePatchArgs{...} }
type SecretStorePatchArrayInput interface {
	pulumi.Input

	ToSecretStorePatchArrayOutput() SecretStorePatchArrayOutput
	ToSecretStorePatchArrayOutputWithContext(context.Context) SecretStorePatchArrayOutput
}

type SecretStorePatchArray []SecretStorePatchInput

func (SecretStorePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretStorePatch)(nil)).Elem()
}

func (i SecretStorePatchArray) ToSecretStorePatchArrayOutput() SecretStorePatchArrayOutput {
	return i.ToSecretStorePatchArrayOutputWithContext(context.Background())
}

func (i SecretStorePatchArray) ToSecretStorePatchArrayOutputWithContext(ctx context.Context) SecretStorePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStorePatchArrayOutput)
}

// SecretStorePatchMapInput is an input type that accepts SecretStorePatchMap and SecretStorePatchMapOutput values.
// You can construct a concrete instance of `SecretStorePatchMapInput` via:
//
//	SecretStorePatchMap{ "key": SecretStorePatchArgs{...} }
type SecretStorePatchMapInput interface {
	pulumi.Input

	ToSecretStorePatchMapOutput() SecretStorePatchMapOutput
	ToSecretStorePatchMapOutputWithContext(context.Context) SecretStorePatchMapOutput
}

type SecretStorePatchMap map[string]SecretStorePatchInput

func (SecretStorePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretStorePatch)(nil)).Elem()
}

func (i SecretStorePatchMap) ToSecretStorePatchMapOutput() SecretStorePatchMapOutput {
	return i.ToSecretStorePatchMapOutputWithContext(context.Background())
}

func (i SecretStorePatchMap) ToSecretStorePatchMapOutputWithContext(ctx context.Context) SecretStorePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStorePatchMapOutput)
}

type SecretStorePatchOutput struct{ *pulumi.OutputState }

func (SecretStorePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStorePatch)(nil)).Elem()
}

func (o SecretStorePatchOutput) ToSecretStorePatchOutput() SecretStorePatchOutput {
	return o
}

func (o SecretStorePatchOutput) ToSecretStorePatchOutputWithContext(ctx context.Context) SecretStorePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o SecretStorePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStorePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o SecretStorePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStorePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o SecretStorePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *SecretStorePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o SecretStorePatchOutput) Spec() SecretStoreSpecPatchPtrOutput {
	return o.ApplyT(func(v *SecretStorePatch) SecretStoreSpecPatchPtrOutput { return v.Spec }).(SecretStoreSpecPatchPtrOutput)
}

func (o SecretStorePatchOutput) Status() SecretStoreStatusPatchPtrOutput {
	return o.ApplyT(func(v *SecretStorePatch) SecretStoreStatusPatchPtrOutput { return v.Status }).(SecretStoreStatusPatchPtrOutput)
}

type SecretStorePatchArrayOutput struct{ *pulumi.OutputState }

func (SecretStorePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretStorePatch)(nil)).Elem()
}

func (o SecretStorePatchArrayOutput) ToSecretStorePatchArrayOutput() SecretStorePatchArrayOutput {
	return o
}

func (o SecretStorePatchArrayOutput) ToSecretStorePatchArrayOutputWithContext(ctx context.Context) SecretStorePatchArrayOutput {
	return o
}

func (o SecretStorePatchArrayOutput) Index(i pulumi.IntInput) SecretStorePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretStorePatch {
		return vs[0].([]*SecretStorePatch)[vs[1].(int)]
	}).(SecretStorePatchOutput)
}

type SecretStorePatchMapOutput struct{ *pulumi.OutputState }

func (SecretStorePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretStorePatch)(nil)).Elem()
}

func (o SecretStorePatchMapOutput) ToSecretStorePatchMapOutput() SecretStorePatchMapOutput {
	return o
}

func (o SecretStorePatchMapOutput) ToSecretStorePatchMapOutputWithContext(ctx context.Context) SecretStorePatchMapOutput {
	return o
}

func (o SecretStorePatchMapOutput) MapIndex(k pulumi.StringInput) SecretStorePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretStorePatch {
		return vs[0].(map[string]*SecretStorePatch)[vs[1].(string)]
	}).(SecretStorePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretStorePatchInput)(nil)).Elem(), &SecretStorePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretStorePatchArrayInput)(nil)).Elem(), SecretStorePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretStorePatchMapInput)(nil)).Elem(), SecretStorePatchMap{})
	pulumi.RegisterOutputType(SecretStorePatchOutput{})
	pulumi.RegisterOutputType(SecretStorePatchArrayOutput{})
	pulumi.RegisterOutputType(SecretStorePatchMapOutput{})
}
