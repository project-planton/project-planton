// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// KafkaMirrorMaker2List is a list of KafkaMirrorMaker2
type KafkaMirrorMaker2List struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of kafkamirrormaker2s. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items KafkaMirrorMaker2TypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewKafkaMirrorMaker2List registers a new resource with the given unique name, arguments, and options.
func NewKafkaMirrorMaker2List(ctx *pulumi.Context,
	name string, args *KafkaMirrorMaker2ListArgs, opts ...pulumi.ResourceOption) (*KafkaMirrorMaker2List, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("kafka.strimzi.io/v1beta2")
	args.Kind = pulumi.StringPtr("KafkaMirrorMaker2List")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KafkaMirrorMaker2List
	err := ctx.RegisterResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaMirrorMaker2List", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaMirrorMaker2List gets an existing KafkaMirrorMaker2List resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaMirrorMaker2List(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaMirrorMaker2ListState, opts ...pulumi.ResourceOption) (*KafkaMirrorMaker2List, error) {
	var resource KafkaMirrorMaker2List
	err := ctx.ReadResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaMirrorMaker2List", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaMirrorMaker2List resources.
type kafkaMirrorMaker2ListState struct {
}

type KafkaMirrorMaker2ListState struct {
}

func (KafkaMirrorMaker2ListState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaMirrorMaker2ListState)(nil)).Elem()
}

type kafkaMirrorMaker2ListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of kafkamirrormaker2s. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []KafkaMirrorMaker2Type `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a KafkaMirrorMaker2List resource.
type KafkaMirrorMaker2ListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of kafkamirrormaker2s. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items KafkaMirrorMaker2TypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (KafkaMirrorMaker2ListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaMirrorMaker2ListArgs)(nil)).Elem()
}

type KafkaMirrorMaker2ListInput interface {
	pulumi.Input

	ToKafkaMirrorMaker2ListOutput() KafkaMirrorMaker2ListOutput
	ToKafkaMirrorMaker2ListOutputWithContext(ctx context.Context) KafkaMirrorMaker2ListOutput
}

func (*KafkaMirrorMaker2List) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaMirrorMaker2List)(nil)).Elem()
}

func (i *KafkaMirrorMaker2List) ToKafkaMirrorMaker2ListOutput() KafkaMirrorMaker2ListOutput {
	return i.ToKafkaMirrorMaker2ListOutputWithContext(context.Background())
}

func (i *KafkaMirrorMaker2List) ToKafkaMirrorMaker2ListOutputWithContext(ctx context.Context) KafkaMirrorMaker2ListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaMirrorMaker2ListOutput)
}

// KafkaMirrorMaker2ListArrayInput is an input type that accepts KafkaMirrorMaker2ListArray and KafkaMirrorMaker2ListArrayOutput values.
// You can construct a concrete instance of `KafkaMirrorMaker2ListArrayInput` via:
//
//	KafkaMirrorMaker2ListArray{ KafkaMirrorMaker2ListArgs{...} }
type KafkaMirrorMaker2ListArrayInput interface {
	pulumi.Input

	ToKafkaMirrorMaker2ListArrayOutput() KafkaMirrorMaker2ListArrayOutput
	ToKafkaMirrorMaker2ListArrayOutputWithContext(context.Context) KafkaMirrorMaker2ListArrayOutput
}

type KafkaMirrorMaker2ListArray []KafkaMirrorMaker2ListInput

func (KafkaMirrorMaker2ListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaMirrorMaker2List)(nil)).Elem()
}

func (i KafkaMirrorMaker2ListArray) ToKafkaMirrorMaker2ListArrayOutput() KafkaMirrorMaker2ListArrayOutput {
	return i.ToKafkaMirrorMaker2ListArrayOutputWithContext(context.Background())
}

func (i KafkaMirrorMaker2ListArray) ToKafkaMirrorMaker2ListArrayOutputWithContext(ctx context.Context) KafkaMirrorMaker2ListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaMirrorMaker2ListArrayOutput)
}

// KafkaMirrorMaker2ListMapInput is an input type that accepts KafkaMirrorMaker2ListMap and KafkaMirrorMaker2ListMapOutput values.
// You can construct a concrete instance of `KafkaMirrorMaker2ListMapInput` via:
//
//	KafkaMirrorMaker2ListMap{ "key": KafkaMirrorMaker2ListArgs{...} }
type KafkaMirrorMaker2ListMapInput interface {
	pulumi.Input

	ToKafkaMirrorMaker2ListMapOutput() KafkaMirrorMaker2ListMapOutput
	ToKafkaMirrorMaker2ListMapOutputWithContext(context.Context) KafkaMirrorMaker2ListMapOutput
}

type KafkaMirrorMaker2ListMap map[string]KafkaMirrorMaker2ListInput

func (KafkaMirrorMaker2ListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaMirrorMaker2List)(nil)).Elem()
}

func (i KafkaMirrorMaker2ListMap) ToKafkaMirrorMaker2ListMapOutput() KafkaMirrorMaker2ListMapOutput {
	return i.ToKafkaMirrorMaker2ListMapOutputWithContext(context.Background())
}

func (i KafkaMirrorMaker2ListMap) ToKafkaMirrorMaker2ListMapOutputWithContext(ctx context.Context) KafkaMirrorMaker2ListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaMirrorMaker2ListMapOutput)
}

type KafkaMirrorMaker2ListOutput struct{ *pulumi.OutputState }

func (KafkaMirrorMaker2ListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaMirrorMaker2List)(nil)).Elem()
}

func (o KafkaMirrorMaker2ListOutput) ToKafkaMirrorMaker2ListOutput() KafkaMirrorMaker2ListOutput {
	return o
}

func (o KafkaMirrorMaker2ListOutput) ToKafkaMirrorMaker2ListOutputWithContext(ctx context.Context) KafkaMirrorMaker2ListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o KafkaMirrorMaker2ListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2List) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of kafkamirrormaker2s. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o KafkaMirrorMaker2ListOutput) Items() KafkaMirrorMaker2TypeArrayOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2List) KafkaMirrorMaker2TypeArrayOutput { return v.Items }).(KafkaMirrorMaker2TypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KafkaMirrorMaker2ListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2List) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KafkaMirrorMaker2ListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2List) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type KafkaMirrorMaker2ListArrayOutput struct{ *pulumi.OutputState }

func (KafkaMirrorMaker2ListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaMirrorMaker2List)(nil)).Elem()
}

func (o KafkaMirrorMaker2ListArrayOutput) ToKafkaMirrorMaker2ListArrayOutput() KafkaMirrorMaker2ListArrayOutput {
	return o
}

func (o KafkaMirrorMaker2ListArrayOutput) ToKafkaMirrorMaker2ListArrayOutputWithContext(ctx context.Context) KafkaMirrorMaker2ListArrayOutput {
	return o
}

func (o KafkaMirrorMaker2ListArrayOutput) Index(i pulumi.IntInput) KafkaMirrorMaker2ListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaMirrorMaker2List {
		return vs[0].([]*KafkaMirrorMaker2List)[vs[1].(int)]
	}).(KafkaMirrorMaker2ListOutput)
}

type KafkaMirrorMaker2ListMapOutput struct{ *pulumi.OutputState }

func (KafkaMirrorMaker2ListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaMirrorMaker2List)(nil)).Elem()
}

func (o KafkaMirrorMaker2ListMapOutput) ToKafkaMirrorMaker2ListMapOutput() KafkaMirrorMaker2ListMapOutput {
	return o
}

func (o KafkaMirrorMaker2ListMapOutput) ToKafkaMirrorMaker2ListMapOutputWithContext(ctx context.Context) KafkaMirrorMaker2ListMapOutput {
	return o
}

func (o KafkaMirrorMaker2ListMapOutput) MapIndex(k pulumi.StringInput) KafkaMirrorMaker2ListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaMirrorMaker2List {
		return vs[0].(map[string]*KafkaMirrorMaker2List)[vs[1].(string)]
	}).(KafkaMirrorMaker2ListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaMirrorMaker2ListInput)(nil)).Elem(), &KafkaMirrorMaker2List{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaMirrorMaker2ListArrayInput)(nil)).Elem(), KafkaMirrorMaker2ListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaMirrorMaker2ListMapInput)(nil)).Elem(), KafkaMirrorMaker2ListMap{})
	pulumi.RegisterOutputType(KafkaMirrorMaker2ListOutput{})
	pulumi.RegisterOutputType(KafkaMirrorMaker2ListArrayOutput{})
	pulumi.RegisterOutputType(KafkaMirrorMaker2ListMapOutput{})
}
