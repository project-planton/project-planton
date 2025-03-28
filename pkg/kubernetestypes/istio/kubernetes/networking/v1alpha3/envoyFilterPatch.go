// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha3

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
type EnvoyFilterPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     EnvoyFilterSpecPatchPtrOutput   `pulumi:"spec"`
	Status   pulumi.MapOutput                `pulumi:"status"`
}

// NewEnvoyFilterPatch registers a new resource with the given unique name, arguments, and options.
func NewEnvoyFilterPatch(ctx *pulumi.Context,
	name string, args *EnvoyFilterPatchArgs, opts ...pulumi.ResourceOption) (*EnvoyFilterPatch, error) {
	if args == nil {
		args = &EnvoyFilterPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1alpha3")
	args.Kind = pulumi.StringPtr("EnvoyFilter")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnvoyFilterPatch
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1alpha3:EnvoyFilterPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvoyFilterPatch gets an existing EnvoyFilterPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvoyFilterPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvoyFilterPatchState, opts ...pulumi.ResourceOption) (*EnvoyFilterPatch, error) {
	var resource EnvoyFilterPatch
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1alpha3:EnvoyFilterPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvoyFilterPatch resources.
type envoyFilterPatchState struct {
}

type EnvoyFilterPatchState struct {
}

func (EnvoyFilterPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*envoyFilterPatchState)(nil)).Elem()
}

type envoyFilterPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *EnvoyFilterSpecPatch   `pulumi:"spec"`
}

// The set of arguments for constructing a EnvoyFilterPatch resource.
type EnvoyFilterPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     EnvoyFilterSpecPatchPtrInput
}

func (EnvoyFilterPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*envoyFilterPatchArgs)(nil)).Elem()
}

type EnvoyFilterPatchInput interface {
	pulumi.Input

	ToEnvoyFilterPatchOutput() EnvoyFilterPatchOutput
	ToEnvoyFilterPatchOutputWithContext(ctx context.Context) EnvoyFilterPatchOutput
}

func (*EnvoyFilterPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvoyFilterPatch)(nil)).Elem()
}

func (i *EnvoyFilterPatch) ToEnvoyFilterPatchOutput() EnvoyFilterPatchOutput {
	return i.ToEnvoyFilterPatchOutputWithContext(context.Background())
}

func (i *EnvoyFilterPatch) ToEnvoyFilterPatchOutputWithContext(ctx context.Context) EnvoyFilterPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvoyFilterPatchOutput)
}

// EnvoyFilterPatchArrayInput is an input type that accepts EnvoyFilterPatchArray and EnvoyFilterPatchArrayOutput values.
// You can construct a concrete instance of `EnvoyFilterPatchArrayInput` via:
//
//	EnvoyFilterPatchArray{ EnvoyFilterPatchArgs{...} }
type EnvoyFilterPatchArrayInput interface {
	pulumi.Input

	ToEnvoyFilterPatchArrayOutput() EnvoyFilterPatchArrayOutput
	ToEnvoyFilterPatchArrayOutputWithContext(context.Context) EnvoyFilterPatchArrayOutput
}

type EnvoyFilterPatchArray []EnvoyFilterPatchInput

func (EnvoyFilterPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvoyFilterPatch)(nil)).Elem()
}

func (i EnvoyFilterPatchArray) ToEnvoyFilterPatchArrayOutput() EnvoyFilterPatchArrayOutput {
	return i.ToEnvoyFilterPatchArrayOutputWithContext(context.Background())
}

func (i EnvoyFilterPatchArray) ToEnvoyFilterPatchArrayOutputWithContext(ctx context.Context) EnvoyFilterPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvoyFilterPatchArrayOutput)
}

// EnvoyFilterPatchMapInput is an input type that accepts EnvoyFilterPatchMap and EnvoyFilterPatchMapOutput values.
// You can construct a concrete instance of `EnvoyFilterPatchMapInput` via:
//
//	EnvoyFilterPatchMap{ "key": EnvoyFilterPatchArgs{...} }
type EnvoyFilterPatchMapInput interface {
	pulumi.Input

	ToEnvoyFilterPatchMapOutput() EnvoyFilterPatchMapOutput
	ToEnvoyFilterPatchMapOutputWithContext(context.Context) EnvoyFilterPatchMapOutput
}

type EnvoyFilterPatchMap map[string]EnvoyFilterPatchInput

func (EnvoyFilterPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvoyFilterPatch)(nil)).Elem()
}

func (i EnvoyFilterPatchMap) ToEnvoyFilterPatchMapOutput() EnvoyFilterPatchMapOutput {
	return i.ToEnvoyFilterPatchMapOutputWithContext(context.Background())
}

func (i EnvoyFilterPatchMap) ToEnvoyFilterPatchMapOutputWithContext(ctx context.Context) EnvoyFilterPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvoyFilterPatchMapOutput)
}

type EnvoyFilterPatchOutput struct{ *pulumi.OutputState }

func (EnvoyFilterPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvoyFilterPatch)(nil)).Elem()
}

func (o EnvoyFilterPatchOutput) ToEnvoyFilterPatchOutput() EnvoyFilterPatchOutput {
	return o
}

func (o EnvoyFilterPatchOutput) ToEnvoyFilterPatchOutputWithContext(ctx context.Context) EnvoyFilterPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EnvoyFilterPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvoyFilterPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EnvoyFilterPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvoyFilterPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o EnvoyFilterPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *EnvoyFilterPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o EnvoyFilterPatchOutput) Spec() EnvoyFilterSpecPatchPtrOutput {
	return o.ApplyT(func(v *EnvoyFilterPatch) EnvoyFilterSpecPatchPtrOutput { return v.Spec }).(EnvoyFilterSpecPatchPtrOutput)
}

func (o EnvoyFilterPatchOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *EnvoyFilterPatch) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

type EnvoyFilterPatchArrayOutput struct{ *pulumi.OutputState }

func (EnvoyFilterPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvoyFilterPatch)(nil)).Elem()
}

func (o EnvoyFilterPatchArrayOutput) ToEnvoyFilterPatchArrayOutput() EnvoyFilterPatchArrayOutput {
	return o
}

func (o EnvoyFilterPatchArrayOutput) ToEnvoyFilterPatchArrayOutputWithContext(ctx context.Context) EnvoyFilterPatchArrayOutput {
	return o
}

func (o EnvoyFilterPatchArrayOutput) Index(i pulumi.IntInput) EnvoyFilterPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnvoyFilterPatch {
		return vs[0].([]*EnvoyFilterPatch)[vs[1].(int)]
	}).(EnvoyFilterPatchOutput)
}

type EnvoyFilterPatchMapOutput struct{ *pulumi.OutputState }

func (EnvoyFilterPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvoyFilterPatch)(nil)).Elem()
}

func (o EnvoyFilterPatchMapOutput) ToEnvoyFilterPatchMapOutput() EnvoyFilterPatchMapOutput {
	return o
}

func (o EnvoyFilterPatchMapOutput) ToEnvoyFilterPatchMapOutputWithContext(ctx context.Context) EnvoyFilterPatchMapOutput {
	return o
}

func (o EnvoyFilterPatchMapOutput) MapIndex(k pulumi.StringInput) EnvoyFilterPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnvoyFilterPatch {
		return vs[0].(map[string]*EnvoyFilterPatch)[vs[1].(string)]
	}).(EnvoyFilterPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvoyFilterPatchInput)(nil)).Elem(), &EnvoyFilterPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvoyFilterPatchArrayInput)(nil)).Elem(), EnvoyFilterPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvoyFilterPatchMapInput)(nil)).Elem(), EnvoyFilterPatchMap{})
	pulumi.RegisterOutputType(EnvoyFilterPatchOutput{})
	pulumi.RegisterOutputType(EnvoyFilterPatchArrayOutput{})
	pulumi.RegisterOutputType(EnvoyFilterPatchMapOutput{})
}
