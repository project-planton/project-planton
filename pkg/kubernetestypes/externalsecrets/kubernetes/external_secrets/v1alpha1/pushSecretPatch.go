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
type PushSecretPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     PushSecretSpecPatchPtrOutput    `pulumi:"spec"`
	Status   PushSecretStatusPatchPtrOutput  `pulumi:"status"`
}

// NewPushSecretPatch registers a new resource with the given unique name, arguments, and options.
func NewPushSecretPatch(ctx *pulumi.Context,
	name string, args *PushSecretPatchArgs, opts ...pulumi.ResourceOption) (*PushSecretPatch, error) {
	if args == nil {
		args = &PushSecretPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("PushSecret")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PushSecretPatch
	err := ctx.RegisterResource("kubernetes:external-secrets.io/v1alpha1:PushSecretPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPushSecretPatch gets an existing PushSecretPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPushSecretPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PushSecretPatchState, opts ...pulumi.ResourceOption) (*PushSecretPatch, error) {
	var resource PushSecretPatch
	err := ctx.ReadResource("kubernetes:external-secrets.io/v1alpha1:PushSecretPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PushSecretPatch resources.
type pushSecretPatchState struct {
}

type PushSecretPatchState struct {
}

func (PushSecretPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*pushSecretPatchState)(nil)).Elem()
}

type pushSecretPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *PushSecretSpecPatch    `pulumi:"spec"`
}

// The set of arguments for constructing a PushSecretPatch resource.
type PushSecretPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     PushSecretSpecPatchPtrInput
}

func (PushSecretPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pushSecretPatchArgs)(nil)).Elem()
}

type PushSecretPatchInput interface {
	pulumi.Input

	ToPushSecretPatchOutput() PushSecretPatchOutput
	ToPushSecretPatchOutputWithContext(ctx context.Context) PushSecretPatchOutput
}

func (*PushSecretPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSecretPatch)(nil)).Elem()
}

func (i *PushSecretPatch) ToPushSecretPatchOutput() PushSecretPatchOutput {
	return i.ToPushSecretPatchOutputWithContext(context.Background())
}

func (i *PushSecretPatch) ToPushSecretPatchOutputWithContext(ctx context.Context) PushSecretPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSecretPatchOutput)
}

// PushSecretPatchArrayInput is an input type that accepts PushSecretPatchArray and PushSecretPatchArrayOutput values.
// You can construct a concrete instance of `PushSecretPatchArrayInput` via:
//
//	PushSecretPatchArray{ PushSecretPatchArgs{...} }
type PushSecretPatchArrayInput interface {
	pulumi.Input

	ToPushSecretPatchArrayOutput() PushSecretPatchArrayOutput
	ToPushSecretPatchArrayOutputWithContext(context.Context) PushSecretPatchArrayOutput
}

type PushSecretPatchArray []PushSecretPatchInput

func (PushSecretPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PushSecretPatch)(nil)).Elem()
}

func (i PushSecretPatchArray) ToPushSecretPatchArrayOutput() PushSecretPatchArrayOutput {
	return i.ToPushSecretPatchArrayOutputWithContext(context.Background())
}

func (i PushSecretPatchArray) ToPushSecretPatchArrayOutputWithContext(ctx context.Context) PushSecretPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSecretPatchArrayOutput)
}

// PushSecretPatchMapInput is an input type that accepts PushSecretPatchMap and PushSecretPatchMapOutput values.
// You can construct a concrete instance of `PushSecretPatchMapInput` via:
//
//	PushSecretPatchMap{ "key": PushSecretPatchArgs{...} }
type PushSecretPatchMapInput interface {
	pulumi.Input

	ToPushSecretPatchMapOutput() PushSecretPatchMapOutput
	ToPushSecretPatchMapOutputWithContext(context.Context) PushSecretPatchMapOutput
}

type PushSecretPatchMap map[string]PushSecretPatchInput

func (PushSecretPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PushSecretPatch)(nil)).Elem()
}

func (i PushSecretPatchMap) ToPushSecretPatchMapOutput() PushSecretPatchMapOutput {
	return i.ToPushSecretPatchMapOutputWithContext(context.Background())
}

func (i PushSecretPatchMap) ToPushSecretPatchMapOutputWithContext(ctx context.Context) PushSecretPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSecretPatchMapOutput)
}

type PushSecretPatchOutput struct{ *pulumi.OutputState }

func (PushSecretPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSecretPatch)(nil)).Elem()
}

func (o PushSecretPatchOutput) ToPushSecretPatchOutput() PushSecretPatchOutput {
	return o
}

func (o PushSecretPatchOutput) ToPushSecretPatchOutputWithContext(ctx context.Context) PushSecretPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PushSecretPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSecretPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PushSecretPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSecretPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PushSecretPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *PushSecretPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o PushSecretPatchOutput) Spec() PushSecretSpecPatchPtrOutput {
	return o.ApplyT(func(v *PushSecretPatch) PushSecretSpecPatchPtrOutput { return v.Spec }).(PushSecretSpecPatchPtrOutput)
}

func (o PushSecretPatchOutput) Status() PushSecretStatusPatchPtrOutput {
	return o.ApplyT(func(v *PushSecretPatch) PushSecretStatusPatchPtrOutput { return v.Status }).(PushSecretStatusPatchPtrOutput)
}

type PushSecretPatchArrayOutput struct{ *pulumi.OutputState }

func (PushSecretPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PushSecretPatch)(nil)).Elem()
}

func (o PushSecretPatchArrayOutput) ToPushSecretPatchArrayOutput() PushSecretPatchArrayOutput {
	return o
}

func (o PushSecretPatchArrayOutput) ToPushSecretPatchArrayOutputWithContext(ctx context.Context) PushSecretPatchArrayOutput {
	return o
}

func (o PushSecretPatchArrayOutput) Index(i pulumi.IntInput) PushSecretPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PushSecretPatch {
		return vs[0].([]*PushSecretPatch)[vs[1].(int)]
	}).(PushSecretPatchOutput)
}

type PushSecretPatchMapOutput struct{ *pulumi.OutputState }

func (PushSecretPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PushSecretPatch)(nil)).Elem()
}

func (o PushSecretPatchMapOutput) ToPushSecretPatchMapOutput() PushSecretPatchMapOutput {
	return o
}

func (o PushSecretPatchMapOutput) ToPushSecretPatchMapOutputWithContext(ctx context.Context) PushSecretPatchMapOutput {
	return o
}

func (o PushSecretPatchMapOutput) MapIndex(k pulumi.StringInput) PushSecretPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PushSecretPatch {
		return vs[0].(map[string]*PushSecretPatch)[vs[1].(string)]
	}).(PushSecretPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PushSecretPatchInput)(nil)).Elem(), &PushSecretPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PushSecretPatchArrayInput)(nil)).Elem(), PushSecretPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PushSecretPatchMapInput)(nil)).Elem(), PushSecretPatchMap{})
	pulumi.RegisterOutputType(PushSecretPatchOutput{})
	pulumi.RegisterOutputType(PushSecretPatchArrayOutput{})
	pulumi.RegisterOutputType(PushSecretPatchMapOutput{})
}