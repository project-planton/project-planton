// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

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
type KafkaConnectPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput  `pulumi:"metadata"`
	Spec     KafkaConnectSpecPatchPtrOutput   `pulumi:"spec"`
	Status   KafkaConnectStatusPatchPtrOutput `pulumi:"status"`
}

// NewKafkaConnectPatch registers a new resource with the given unique name, arguments, and options.
func NewKafkaConnectPatch(ctx *pulumi.Context,
	name string, args *KafkaConnectPatchArgs, opts ...pulumi.ResourceOption) (*KafkaConnectPatch, error) {
	if args == nil {
		args = &KafkaConnectPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("kafka.strimzi.io/v1beta2")
	args.Kind = pulumi.StringPtr("KafkaConnect")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KafkaConnectPatch
	err := ctx.RegisterResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaConnectPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaConnectPatch gets an existing KafkaConnectPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaConnectPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaConnectPatchState, opts ...pulumi.ResourceOption) (*KafkaConnectPatch, error) {
	var resource KafkaConnectPatch
	err := ctx.ReadResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaConnectPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaConnectPatch resources.
type kafkaConnectPatchState struct {
}

type KafkaConnectPatchState struct {
}

func (KafkaConnectPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaConnectPatchState)(nil)).Elem()
}

type kafkaConnectPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *KafkaConnectSpecPatch  `pulumi:"spec"`
}

// The set of arguments for constructing a KafkaConnectPatch resource.
type KafkaConnectPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     KafkaConnectSpecPatchPtrInput
}

func (KafkaConnectPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaConnectPatchArgs)(nil)).Elem()
}

type KafkaConnectPatchInput interface {
	pulumi.Input

	ToKafkaConnectPatchOutput() KafkaConnectPatchOutput
	ToKafkaConnectPatchOutputWithContext(ctx context.Context) KafkaConnectPatchOutput
}

func (*KafkaConnectPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaConnectPatch)(nil)).Elem()
}

func (i *KafkaConnectPatch) ToKafkaConnectPatchOutput() KafkaConnectPatchOutput {
	return i.ToKafkaConnectPatchOutputWithContext(context.Background())
}

func (i *KafkaConnectPatch) ToKafkaConnectPatchOutputWithContext(ctx context.Context) KafkaConnectPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaConnectPatchOutput)
}

// KafkaConnectPatchArrayInput is an input type that accepts KafkaConnectPatchArray and KafkaConnectPatchArrayOutput values.
// You can construct a concrete instance of `KafkaConnectPatchArrayInput` via:
//
//	KafkaConnectPatchArray{ KafkaConnectPatchArgs{...} }
type KafkaConnectPatchArrayInput interface {
	pulumi.Input

	ToKafkaConnectPatchArrayOutput() KafkaConnectPatchArrayOutput
	ToKafkaConnectPatchArrayOutputWithContext(context.Context) KafkaConnectPatchArrayOutput
}

type KafkaConnectPatchArray []KafkaConnectPatchInput

func (KafkaConnectPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaConnectPatch)(nil)).Elem()
}

func (i KafkaConnectPatchArray) ToKafkaConnectPatchArrayOutput() KafkaConnectPatchArrayOutput {
	return i.ToKafkaConnectPatchArrayOutputWithContext(context.Background())
}

func (i KafkaConnectPatchArray) ToKafkaConnectPatchArrayOutputWithContext(ctx context.Context) KafkaConnectPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaConnectPatchArrayOutput)
}

// KafkaConnectPatchMapInput is an input type that accepts KafkaConnectPatchMap and KafkaConnectPatchMapOutput values.
// You can construct a concrete instance of `KafkaConnectPatchMapInput` via:
//
//	KafkaConnectPatchMap{ "key": KafkaConnectPatchArgs{...} }
type KafkaConnectPatchMapInput interface {
	pulumi.Input

	ToKafkaConnectPatchMapOutput() KafkaConnectPatchMapOutput
	ToKafkaConnectPatchMapOutputWithContext(context.Context) KafkaConnectPatchMapOutput
}

type KafkaConnectPatchMap map[string]KafkaConnectPatchInput

func (KafkaConnectPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaConnectPatch)(nil)).Elem()
}

func (i KafkaConnectPatchMap) ToKafkaConnectPatchMapOutput() KafkaConnectPatchMapOutput {
	return i.ToKafkaConnectPatchMapOutputWithContext(context.Background())
}

func (i KafkaConnectPatchMap) ToKafkaConnectPatchMapOutputWithContext(ctx context.Context) KafkaConnectPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaConnectPatchMapOutput)
}

type KafkaConnectPatchOutput struct{ *pulumi.OutputState }

func (KafkaConnectPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaConnectPatch)(nil)).Elem()
}

func (o KafkaConnectPatchOutput) ToKafkaConnectPatchOutput() KafkaConnectPatchOutput {
	return o
}

func (o KafkaConnectPatchOutput) ToKafkaConnectPatchOutputWithContext(ctx context.Context) KafkaConnectPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o KafkaConnectPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaConnectPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KafkaConnectPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaConnectPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o KafkaConnectPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *KafkaConnectPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o KafkaConnectPatchOutput) Spec() KafkaConnectSpecPatchPtrOutput {
	return o.ApplyT(func(v *KafkaConnectPatch) KafkaConnectSpecPatchPtrOutput { return v.Spec }).(KafkaConnectSpecPatchPtrOutput)
}

func (o KafkaConnectPatchOutput) Status() KafkaConnectStatusPatchPtrOutput {
	return o.ApplyT(func(v *KafkaConnectPatch) KafkaConnectStatusPatchPtrOutput { return v.Status }).(KafkaConnectStatusPatchPtrOutput)
}

type KafkaConnectPatchArrayOutput struct{ *pulumi.OutputState }

func (KafkaConnectPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaConnectPatch)(nil)).Elem()
}

func (o KafkaConnectPatchArrayOutput) ToKafkaConnectPatchArrayOutput() KafkaConnectPatchArrayOutput {
	return o
}

func (o KafkaConnectPatchArrayOutput) ToKafkaConnectPatchArrayOutputWithContext(ctx context.Context) KafkaConnectPatchArrayOutput {
	return o
}

func (o KafkaConnectPatchArrayOutput) Index(i pulumi.IntInput) KafkaConnectPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaConnectPatch {
		return vs[0].([]*KafkaConnectPatch)[vs[1].(int)]
	}).(KafkaConnectPatchOutput)
}

type KafkaConnectPatchMapOutput struct{ *pulumi.OutputState }

func (KafkaConnectPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaConnectPatch)(nil)).Elem()
}

func (o KafkaConnectPatchMapOutput) ToKafkaConnectPatchMapOutput() KafkaConnectPatchMapOutput {
	return o
}

func (o KafkaConnectPatchMapOutput) ToKafkaConnectPatchMapOutputWithContext(ctx context.Context) KafkaConnectPatchMapOutput {
	return o
}

func (o KafkaConnectPatchMapOutput) MapIndex(k pulumi.StringInput) KafkaConnectPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaConnectPatch {
		return vs[0].(map[string]*KafkaConnectPatch)[vs[1].(string)]
	}).(KafkaConnectPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaConnectPatchInput)(nil)).Elem(), &KafkaConnectPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaConnectPatchArrayInput)(nil)).Elem(), KafkaConnectPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaConnectPatchMapInput)(nil)).Elem(), KafkaConnectPatchMap{})
	pulumi.RegisterOutputType(KafkaConnectPatchOutput{})
	pulumi.RegisterOutputType(KafkaConnectPatchArrayOutput{})
	pulumi.RegisterOutputType(KafkaConnectPatchMapOutput{})
}