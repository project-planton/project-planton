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

// GatewayList is a list of Gateway
type GatewayList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of gateways. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items GatewayTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewGatewayList registers a new resource with the given unique name, arguments, and options.
func NewGatewayList(ctx *pulumi.Context,
	name string, args *GatewayListArgs, opts ...pulumi.ResourceOption) (*GatewayList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1")
	args.Kind = pulumi.StringPtr("GatewayList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GatewayList
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1:GatewayList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayList gets an existing GatewayList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayListState, opts ...pulumi.ResourceOption) (*GatewayList, error) {
	var resource GatewayList
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1:GatewayList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayList resources.
type gatewayListState struct {
}

type GatewayListState struct {
}

func (GatewayListState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayListState)(nil)).Elem()
}

type gatewayListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of gateways. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []GatewayType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a GatewayList resource.
type GatewayListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of gateways. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items GatewayTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (GatewayListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayListArgs)(nil)).Elem()
}

type GatewayListInput interface {
	pulumi.Input

	ToGatewayListOutput() GatewayListOutput
	ToGatewayListOutputWithContext(ctx context.Context) GatewayListOutput
}

func (*GatewayList) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayList)(nil)).Elem()
}

func (i *GatewayList) ToGatewayListOutput() GatewayListOutput {
	return i.ToGatewayListOutputWithContext(context.Background())
}

func (i *GatewayList) ToGatewayListOutputWithContext(ctx context.Context) GatewayListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayListOutput)
}

// GatewayListArrayInput is an input type that accepts GatewayListArray and GatewayListArrayOutput values.
// You can construct a concrete instance of `GatewayListArrayInput` via:
//
//	GatewayListArray{ GatewayListArgs{...} }
type GatewayListArrayInput interface {
	pulumi.Input

	ToGatewayListArrayOutput() GatewayListArrayOutput
	ToGatewayListArrayOutputWithContext(context.Context) GatewayListArrayOutput
}

type GatewayListArray []GatewayListInput

func (GatewayListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayList)(nil)).Elem()
}

func (i GatewayListArray) ToGatewayListArrayOutput() GatewayListArrayOutput {
	return i.ToGatewayListArrayOutputWithContext(context.Background())
}

func (i GatewayListArray) ToGatewayListArrayOutputWithContext(ctx context.Context) GatewayListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayListArrayOutput)
}

// GatewayListMapInput is an input type that accepts GatewayListMap and GatewayListMapOutput values.
// You can construct a concrete instance of `GatewayListMapInput` via:
//
//	GatewayListMap{ "key": GatewayListArgs{...} }
type GatewayListMapInput interface {
	pulumi.Input

	ToGatewayListMapOutput() GatewayListMapOutput
	ToGatewayListMapOutputWithContext(context.Context) GatewayListMapOutput
}

type GatewayListMap map[string]GatewayListInput

func (GatewayListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayList)(nil)).Elem()
}

func (i GatewayListMap) ToGatewayListMapOutput() GatewayListMapOutput {
	return i.ToGatewayListMapOutputWithContext(context.Background())
}

func (i GatewayListMap) ToGatewayListMapOutputWithContext(ctx context.Context) GatewayListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayListMapOutput)
}

type GatewayListOutput struct{ *pulumi.OutputState }

func (GatewayListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayList)(nil)).Elem()
}

func (o GatewayListOutput) ToGatewayListOutput() GatewayListOutput {
	return o
}

func (o GatewayListOutput) ToGatewayListOutputWithContext(ctx context.Context) GatewayListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o GatewayListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of gateways. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o GatewayListOutput) Items() GatewayTypeArrayOutput {
	return o.ApplyT(func(v *GatewayList) GatewayTypeArrayOutput { return v.Items }).(GatewayTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o GatewayListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o GatewayListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *GatewayList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type GatewayListArrayOutput struct{ *pulumi.OutputState }

func (GatewayListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayList)(nil)).Elem()
}

func (o GatewayListArrayOutput) ToGatewayListArrayOutput() GatewayListArrayOutput {
	return o
}

func (o GatewayListArrayOutput) ToGatewayListArrayOutputWithContext(ctx context.Context) GatewayListArrayOutput {
	return o
}

func (o GatewayListArrayOutput) Index(i pulumi.IntInput) GatewayListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayList {
		return vs[0].([]*GatewayList)[vs[1].(int)]
	}).(GatewayListOutput)
}

type GatewayListMapOutput struct{ *pulumi.OutputState }

func (GatewayListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayList)(nil)).Elem()
}

func (o GatewayListMapOutput) ToGatewayListMapOutput() GatewayListMapOutput {
	return o
}

func (o GatewayListMapOutput) ToGatewayListMapOutputWithContext(ctx context.Context) GatewayListMapOutput {
	return o
}

func (o GatewayListMapOutput) MapIndex(k pulumi.StringInput) GatewayListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayList {
		return vs[0].(map[string]*GatewayList)[vs[1].(string)]
	}).(GatewayListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayListInput)(nil)).Elem(), &GatewayList{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayListArrayInput)(nil)).Elem(), GatewayListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayListMapInput)(nil)).Elem(), GatewayListMap{})
	pulumi.RegisterOutputType(GatewayListOutput{})
	pulumi.RegisterOutputType(GatewayListArrayOutput{})
	pulumi.RegisterOutputType(GatewayListMapOutput{})
}
