// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ApmServerList is a list of ApmServer
type ApmServerList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of apmservers. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ApmServerTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewApmServerList registers a new resource with the given unique name, arguments, and options.
func NewApmServerList(ctx *pulumi.Context,
	name string, args *ApmServerListArgs, opts ...pulumi.ResourceOption) (*ApmServerList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apm.k8s.elastic.co/v1")
	args.Kind = pulumi.StringPtr("ApmServerList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApmServerList
	err := ctx.RegisterResource("kubernetes:apm.k8s.elastic.co/v1:ApmServerList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApmServerList gets an existing ApmServerList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApmServerList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApmServerListState, opts ...pulumi.ResourceOption) (*ApmServerList, error) {
	var resource ApmServerList
	err := ctx.ReadResource("kubernetes:apm.k8s.elastic.co/v1:ApmServerList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApmServerList resources.
type apmServerListState struct {
}

type ApmServerListState struct {
}

func (ApmServerListState) ElementType() reflect.Type {
	return reflect.TypeOf((*apmServerListState)(nil)).Elem()
}

type apmServerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of apmservers. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ApmServerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ApmServerList resource.
type ApmServerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of apmservers. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ApmServerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ApmServerListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apmServerListArgs)(nil)).Elem()
}

type ApmServerListInput interface {
	pulumi.Input

	ToApmServerListOutput() ApmServerListOutput
	ToApmServerListOutputWithContext(ctx context.Context) ApmServerListOutput
}

func (*ApmServerList) ElementType() reflect.Type {
	return reflect.TypeOf((**ApmServerList)(nil)).Elem()
}

func (i *ApmServerList) ToApmServerListOutput() ApmServerListOutput {
	return i.ToApmServerListOutputWithContext(context.Background())
}

func (i *ApmServerList) ToApmServerListOutputWithContext(ctx context.Context) ApmServerListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmServerListOutput)
}

// ApmServerListArrayInput is an input type that accepts ApmServerListArray and ApmServerListArrayOutput values.
// You can construct a concrete instance of `ApmServerListArrayInput` via:
//
//	ApmServerListArray{ ApmServerListArgs{...} }
type ApmServerListArrayInput interface {
	pulumi.Input

	ToApmServerListArrayOutput() ApmServerListArrayOutput
	ToApmServerListArrayOutputWithContext(context.Context) ApmServerListArrayOutput
}

type ApmServerListArray []ApmServerListInput

func (ApmServerListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApmServerList)(nil)).Elem()
}

func (i ApmServerListArray) ToApmServerListArrayOutput() ApmServerListArrayOutput {
	return i.ToApmServerListArrayOutputWithContext(context.Background())
}

func (i ApmServerListArray) ToApmServerListArrayOutputWithContext(ctx context.Context) ApmServerListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmServerListArrayOutput)
}

// ApmServerListMapInput is an input type that accepts ApmServerListMap and ApmServerListMapOutput values.
// You can construct a concrete instance of `ApmServerListMapInput` via:
//
//	ApmServerListMap{ "key": ApmServerListArgs{...} }
type ApmServerListMapInput interface {
	pulumi.Input

	ToApmServerListMapOutput() ApmServerListMapOutput
	ToApmServerListMapOutputWithContext(context.Context) ApmServerListMapOutput
}

type ApmServerListMap map[string]ApmServerListInput

func (ApmServerListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApmServerList)(nil)).Elem()
}

func (i ApmServerListMap) ToApmServerListMapOutput() ApmServerListMapOutput {
	return i.ToApmServerListMapOutputWithContext(context.Background())
}

func (i ApmServerListMap) ToApmServerListMapOutputWithContext(ctx context.Context) ApmServerListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmServerListMapOutput)
}

type ApmServerListOutput struct{ *pulumi.OutputState }

func (ApmServerListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApmServerList)(nil)).Elem()
}

func (o ApmServerListOutput) ToApmServerListOutput() ApmServerListOutput {
	return o
}

func (o ApmServerListOutput) ToApmServerListOutputWithContext(ctx context.Context) ApmServerListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ApmServerListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ApmServerList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of apmservers. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ApmServerListOutput) Items() ApmServerTypeArrayOutput {
	return o.ApplyT(func(v *ApmServerList) ApmServerTypeArrayOutput { return v.Items }).(ApmServerTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ApmServerListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ApmServerList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ApmServerListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ApmServerList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ApmServerListArrayOutput struct{ *pulumi.OutputState }

func (ApmServerListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApmServerList)(nil)).Elem()
}

func (o ApmServerListArrayOutput) ToApmServerListArrayOutput() ApmServerListArrayOutput {
	return o
}

func (o ApmServerListArrayOutput) ToApmServerListArrayOutputWithContext(ctx context.Context) ApmServerListArrayOutput {
	return o
}

func (o ApmServerListArrayOutput) Index(i pulumi.IntInput) ApmServerListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApmServerList {
		return vs[0].([]*ApmServerList)[vs[1].(int)]
	}).(ApmServerListOutput)
}

type ApmServerListMapOutput struct{ *pulumi.OutputState }

func (ApmServerListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApmServerList)(nil)).Elem()
}

func (o ApmServerListMapOutput) ToApmServerListMapOutput() ApmServerListMapOutput {
	return o
}

func (o ApmServerListMapOutput) ToApmServerListMapOutputWithContext(ctx context.Context) ApmServerListMapOutput {
	return o
}

func (o ApmServerListMapOutput) MapIndex(k pulumi.StringInput) ApmServerListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApmServerList {
		return vs[0].(map[string]*ApmServerList)[vs[1].(string)]
	}).(ApmServerListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApmServerListInput)(nil)).Elem(), &ApmServerList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApmServerListArrayInput)(nil)).Elem(), ApmServerListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApmServerListMapInput)(nil)).Elem(), ApmServerListMap{})
	pulumi.RegisterOutputType(ApmServerListOutput{})
	pulumi.RegisterOutputType(ApmServerListArrayOutput{})
	pulumi.RegisterOutputType(ApmServerListMapOutput{})
}