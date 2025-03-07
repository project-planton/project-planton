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
// Kibana represents a Kibana resource in a Kubernetes cluster.
type KibanaPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     KibanaSpecPatchPtrOutput        `pulumi:"spec"`
	Status   KibanaStatusPatchPtrOutput      `pulumi:"status"`
}

// NewKibanaPatch registers a new resource with the given unique name, arguments, and options.
func NewKibanaPatch(ctx *pulumi.Context,
	name string, args *KibanaPatchArgs, opts ...pulumi.ResourceOption) (*KibanaPatch, error) {
	if args == nil {
		args = &KibanaPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("kibana.k8s.elastic.co/v1beta1")
	args.Kind = pulumi.StringPtr("Kibana")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:kibana.k8s.elastic.co/v1:KibanaPatch"),
		},
		{
			Type: pulumi.String("kubernetes:kibana.k8s.elastic.co/v1alpha1:KibanaPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KibanaPatch
	err := ctx.RegisterResource("kubernetes:kibana.k8s.elastic.co/v1beta1:KibanaPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKibanaPatch gets an existing KibanaPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKibanaPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KibanaPatchState, opts ...pulumi.ResourceOption) (*KibanaPatch, error) {
	var resource KibanaPatch
	err := ctx.ReadResource("kubernetes:kibana.k8s.elastic.co/v1beta1:KibanaPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KibanaPatch resources.
type kibanaPatchState struct {
}

type KibanaPatchState struct {
}

func (KibanaPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*kibanaPatchState)(nil)).Elem()
}

type kibanaPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *KibanaSpecPatch        `pulumi:"spec"`
}

// The set of arguments for constructing a KibanaPatch resource.
type KibanaPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     KibanaSpecPatchPtrInput
}

func (KibanaPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kibanaPatchArgs)(nil)).Elem()
}

type KibanaPatchInput interface {
	pulumi.Input

	ToKibanaPatchOutput() KibanaPatchOutput
	ToKibanaPatchOutputWithContext(ctx context.Context) KibanaPatchOutput
}

func (*KibanaPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**KibanaPatch)(nil)).Elem()
}

func (i *KibanaPatch) ToKibanaPatchOutput() KibanaPatchOutput {
	return i.ToKibanaPatchOutputWithContext(context.Background())
}

func (i *KibanaPatch) ToKibanaPatchOutputWithContext(ctx context.Context) KibanaPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KibanaPatchOutput)
}

// KibanaPatchArrayInput is an input type that accepts KibanaPatchArray and KibanaPatchArrayOutput values.
// You can construct a concrete instance of `KibanaPatchArrayInput` via:
//
//	KibanaPatchArray{ KibanaPatchArgs{...} }
type KibanaPatchArrayInput interface {
	pulumi.Input

	ToKibanaPatchArrayOutput() KibanaPatchArrayOutput
	ToKibanaPatchArrayOutputWithContext(context.Context) KibanaPatchArrayOutput
}

type KibanaPatchArray []KibanaPatchInput

func (KibanaPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KibanaPatch)(nil)).Elem()
}

func (i KibanaPatchArray) ToKibanaPatchArrayOutput() KibanaPatchArrayOutput {
	return i.ToKibanaPatchArrayOutputWithContext(context.Background())
}

func (i KibanaPatchArray) ToKibanaPatchArrayOutputWithContext(ctx context.Context) KibanaPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KibanaPatchArrayOutput)
}

// KibanaPatchMapInput is an input type that accepts KibanaPatchMap and KibanaPatchMapOutput values.
// You can construct a concrete instance of `KibanaPatchMapInput` via:
//
//	KibanaPatchMap{ "key": KibanaPatchArgs{...} }
type KibanaPatchMapInput interface {
	pulumi.Input

	ToKibanaPatchMapOutput() KibanaPatchMapOutput
	ToKibanaPatchMapOutputWithContext(context.Context) KibanaPatchMapOutput
}

type KibanaPatchMap map[string]KibanaPatchInput

func (KibanaPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KibanaPatch)(nil)).Elem()
}

func (i KibanaPatchMap) ToKibanaPatchMapOutput() KibanaPatchMapOutput {
	return i.ToKibanaPatchMapOutputWithContext(context.Background())
}

func (i KibanaPatchMap) ToKibanaPatchMapOutputWithContext(ctx context.Context) KibanaPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KibanaPatchMapOutput)
}

type KibanaPatchOutput struct{ *pulumi.OutputState }

func (KibanaPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KibanaPatch)(nil)).Elem()
}

func (o KibanaPatchOutput) ToKibanaPatchOutput() KibanaPatchOutput {
	return o
}

func (o KibanaPatchOutput) ToKibanaPatchOutputWithContext(ctx context.Context) KibanaPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o KibanaPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KibanaPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KibanaPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KibanaPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o KibanaPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *KibanaPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o KibanaPatchOutput) Spec() KibanaSpecPatchPtrOutput {
	return o.ApplyT(func(v *KibanaPatch) KibanaSpecPatchPtrOutput { return v.Spec }).(KibanaSpecPatchPtrOutput)
}

func (o KibanaPatchOutput) Status() KibanaStatusPatchPtrOutput {
	return o.ApplyT(func(v *KibanaPatch) KibanaStatusPatchPtrOutput { return v.Status }).(KibanaStatusPatchPtrOutput)
}

type KibanaPatchArrayOutput struct{ *pulumi.OutputState }

func (KibanaPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KibanaPatch)(nil)).Elem()
}

func (o KibanaPatchArrayOutput) ToKibanaPatchArrayOutput() KibanaPatchArrayOutput {
	return o
}

func (o KibanaPatchArrayOutput) ToKibanaPatchArrayOutputWithContext(ctx context.Context) KibanaPatchArrayOutput {
	return o
}

func (o KibanaPatchArrayOutput) Index(i pulumi.IntInput) KibanaPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KibanaPatch {
		return vs[0].([]*KibanaPatch)[vs[1].(int)]
	}).(KibanaPatchOutput)
}

type KibanaPatchMapOutput struct{ *pulumi.OutputState }

func (KibanaPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KibanaPatch)(nil)).Elem()
}

func (o KibanaPatchMapOutput) ToKibanaPatchMapOutput() KibanaPatchMapOutput {
	return o
}

func (o KibanaPatchMapOutput) ToKibanaPatchMapOutputWithContext(ctx context.Context) KibanaPatchMapOutput {
	return o
}

func (o KibanaPatchMapOutput) MapIndex(k pulumi.StringInput) KibanaPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KibanaPatch {
		return vs[0].(map[string]*KibanaPatch)[vs[1].(string)]
	}).(KibanaPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KibanaPatchInput)(nil)).Elem(), &KibanaPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*KibanaPatchArrayInput)(nil)).Elem(), KibanaPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KibanaPatchMapInput)(nil)).Elem(), KibanaPatchMap{})
	pulumi.RegisterOutputType(KibanaPatchOutput{})
	pulumi.RegisterOutputType(KibanaPatchArrayOutput{})
	pulumi.RegisterOutputType(KibanaPatchMapOutput{})
}
