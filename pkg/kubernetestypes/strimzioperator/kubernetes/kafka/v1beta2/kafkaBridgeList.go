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

// KafkaBridgeList is a list of KafkaBridge
type KafkaBridgeList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of kafkabridges. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items KafkaBridgeTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewKafkaBridgeList registers a new resource with the given unique name, arguments, and options.
func NewKafkaBridgeList(ctx *pulumi.Context,
	name string, args *KafkaBridgeListArgs, opts ...pulumi.ResourceOption) (*KafkaBridgeList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("kafka.strimzi.io/v1beta2")
	args.Kind = pulumi.StringPtr("KafkaBridgeList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KafkaBridgeList
	err := ctx.RegisterResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaBridgeList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaBridgeList gets an existing KafkaBridgeList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaBridgeList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaBridgeListState, opts ...pulumi.ResourceOption) (*KafkaBridgeList, error) {
	var resource KafkaBridgeList
	err := ctx.ReadResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaBridgeList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaBridgeList resources.
type kafkaBridgeListState struct {
}

type KafkaBridgeListState struct {
}

func (KafkaBridgeListState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaBridgeListState)(nil)).Elem()
}

type kafkaBridgeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of kafkabridges. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []KafkaBridgeType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a KafkaBridgeList resource.
type KafkaBridgeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of kafkabridges. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items KafkaBridgeTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (KafkaBridgeListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaBridgeListArgs)(nil)).Elem()
}

type KafkaBridgeListInput interface {
	pulumi.Input

	ToKafkaBridgeListOutput() KafkaBridgeListOutput
	ToKafkaBridgeListOutputWithContext(ctx context.Context) KafkaBridgeListOutput
}

func (*KafkaBridgeList) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaBridgeList)(nil)).Elem()
}

func (i *KafkaBridgeList) ToKafkaBridgeListOutput() KafkaBridgeListOutput {
	return i.ToKafkaBridgeListOutputWithContext(context.Background())
}

func (i *KafkaBridgeList) ToKafkaBridgeListOutputWithContext(ctx context.Context) KafkaBridgeListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaBridgeListOutput)
}

// KafkaBridgeListArrayInput is an input type that accepts KafkaBridgeListArray and KafkaBridgeListArrayOutput values.
// You can construct a concrete instance of `KafkaBridgeListArrayInput` via:
//
//	KafkaBridgeListArray{ KafkaBridgeListArgs{...} }
type KafkaBridgeListArrayInput interface {
	pulumi.Input

	ToKafkaBridgeListArrayOutput() KafkaBridgeListArrayOutput
	ToKafkaBridgeListArrayOutputWithContext(context.Context) KafkaBridgeListArrayOutput
}

type KafkaBridgeListArray []KafkaBridgeListInput

func (KafkaBridgeListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaBridgeList)(nil)).Elem()
}

func (i KafkaBridgeListArray) ToKafkaBridgeListArrayOutput() KafkaBridgeListArrayOutput {
	return i.ToKafkaBridgeListArrayOutputWithContext(context.Background())
}

func (i KafkaBridgeListArray) ToKafkaBridgeListArrayOutputWithContext(ctx context.Context) KafkaBridgeListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaBridgeListArrayOutput)
}

// KafkaBridgeListMapInput is an input type that accepts KafkaBridgeListMap and KafkaBridgeListMapOutput values.
// You can construct a concrete instance of `KafkaBridgeListMapInput` via:
//
//	KafkaBridgeListMap{ "key": KafkaBridgeListArgs{...} }
type KafkaBridgeListMapInput interface {
	pulumi.Input

	ToKafkaBridgeListMapOutput() KafkaBridgeListMapOutput
	ToKafkaBridgeListMapOutputWithContext(context.Context) KafkaBridgeListMapOutput
}

type KafkaBridgeListMap map[string]KafkaBridgeListInput

func (KafkaBridgeListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaBridgeList)(nil)).Elem()
}

func (i KafkaBridgeListMap) ToKafkaBridgeListMapOutput() KafkaBridgeListMapOutput {
	return i.ToKafkaBridgeListMapOutputWithContext(context.Background())
}

func (i KafkaBridgeListMap) ToKafkaBridgeListMapOutputWithContext(ctx context.Context) KafkaBridgeListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaBridgeListMapOutput)
}

type KafkaBridgeListOutput struct{ *pulumi.OutputState }

func (KafkaBridgeListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaBridgeList)(nil)).Elem()
}

func (o KafkaBridgeListOutput) ToKafkaBridgeListOutput() KafkaBridgeListOutput {
	return o
}

func (o KafkaBridgeListOutput) ToKafkaBridgeListOutputWithContext(ctx context.Context) KafkaBridgeListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o KafkaBridgeListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaBridgeList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of kafkabridges. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o KafkaBridgeListOutput) Items() KafkaBridgeTypeArrayOutput {
	return o.ApplyT(func(v *KafkaBridgeList) KafkaBridgeTypeArrayOutput { return v.Items }).(KafkaBridgeTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KafkaBridgeListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaBridgeList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KafkaBridgeListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *KafkaBridgeList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type KafkaBridgeListArrayOutput struct{ *pulumi.OutputState }

func (KafkaBridgeListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaBridgeList)(nil)).Elem()
}

func (o KafkaBridgeListArrayOutput) ToKafkaBridgeListArrayOutput() KafkaBridgeListArrayOutput {
	return o
}

func (o KafkaBridgeListArrayOutput) ToKafkaBridgeListArrayOutputWithContext(ctx context.Context) KafkaBridgeListArrayOutput {
	return o
}

func (o KafkaBridgeListArrayOutput) Index(i pulumi.IntInput) KafkaBridgeListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaBridgeList {
		return vs[0].([]*KafkaBridgeList)[vs[1].(int)]
	}).(KafkaBridgeListOutput)
}

type KafkaBridgeListMapOutput struct{ *pulumi.OutputState }

func (KafkaBridgeListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaBridgeList)(nil)).Elem()
}

func (o KafkaBridgeListMapOutput) ToKafkaBridgeListMapOutput() KafkaBridgeListMapOutput {
	return o
}

func (o KafkaBridgeListMapOutput) ToKafkaBridgeListMapOutputWithContext(ctx context.Context) KafkaBridgeListMapOutput {
	return o
}

func (o KafkaBridgeListMapOutput) MapIndex(k pulumi.StringInput) KafkaBridgeListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaBridgeList {
		return vs[0].(map[string]*KafkaBridgeList)[vs[1].(string)]
	}).(KafkaBridgeListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaBridgeListInput)(nil)).Elem(), &KafkaBridgeList{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaBridgeListArrayInput)(nil)).Elem(), KafkaBridgeListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaBridgeListMapInput)(nil)).Elem(), KafkaBridgeListMap{})
	pulumi.RegisterOutputType(KafkaBridgeListOutput{})
	pulumi.RegisterOutputType(KafkaBridgeListArrayOutput{})
	pulumi.RegisterOutputType(KafkaBridgeListMapOutput{})
}
