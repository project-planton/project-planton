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

// GatewayClassList is a list of GatewayClass
type GatewayClassList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of gatewayclasses. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items GatewayClassTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewGatewayClassList registers a new resource with the given unique name, arguments, and options.
func NewGatewayClassList(ctx *pulumi.Context,
	name string, args *GatewayClassListArgs, opts ...pulumi.ResourceOption) (*GatewayClassList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("gateway.networking.k8s.io/v1")
	args.Kind = pulumi.StringPtr("GatewayClassList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GatewayClassList
	err := ctx.RegisterResource("kubernetes:gateway.networking.k8s.io/v1:GatewayClassList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayClassList gets an existing GatewayClassList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayClassList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayClassListState, opts ...pulumi.ResourceOption) (*GatewayClassList, error) {
	var resource GatewayClassList
	err := ctx.ReadResource("kubernetes:gateway.networking.k8s.io/v1:GatewayClassList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayClassList resources.
type gatewayClassListState struct {
}

type GatewayClassListState struct {
}

func (GatewayClassListState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayClassListState)(nil)).Elem()
}

type gatewayClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of gatewayclasses. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []GatewayClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a GatewayClassList resource.
type GatewayClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of gatewayclasses. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items GatewayClassTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (GatewayClassListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayClassListArgs)(nil)).Elem()
}

type GatewayClassListInput interface {
	pulumi.Input

	ToGatewayClassListOutput() GatewayClassListOutput
	ToGatewayClassListOutputWithContext(ctx context.Context) GatewayClassListOutput
}

func (*GatewayClassList) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayClassList)(nil)).Elem()
}

func (i *GatewayClassList) ToGatewayClassListOutput() GatewayClassListOutput {
	return i.ToGatewayClassListOutputWithContext(context.Background())
}

func (i *GatewayClassList) ToGatewayClassListOutputWithContext(ctx context.Context) GatewayClassListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayClassListOutput)
}

// GatewayClassListArrayInput is an input type that accepts GatewayClassListArray and GatewayClassListArrayOutput values.
// You can construct a concrete instance of `GatewayClassListArrayInput` via:
//
//	GatewayClassListArray{ GatewayClassListArgs{...} }
type GatewayClassListArrayInput interface {
	pulumi.Input

	ToGatewayClassListArrayOutput() GatewayClassListArrayOutput
	ToGatewayClassListArrayOutputWithContext(context.Context) GatewayClassListArrayOutput
}

type GatewayClassListArray []GatewayClassListInput

func (GatewayClassListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayClassList)(nil)).Elem()
}

func (i GatewayClassListArray) ToGatewayClassListArrayOutput() GatewayClassListArrayOutput {
	return i.ToGatewayClassListArrayOutputWithContext(context.Background())
}

func (i GatewayClassListArray) ToGatewayClassListArrayOutputWithContext(ctx context.Context) GatewayClassListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayClassListArrayOutput)
}

// GatewayClassListMapInput is an input type that accepts GatewayClassListMap and GatewayClassListMapOutput values.
// You can construct a concrete instance of `GatewayClassListMapInput` via:
//
//	GatewayClassListMap{ "key": GatewayClassListArgs{...} }
type GatewayClassListMapInput interface {
	pulumi.Input

	ToGatewayClassListMapOutput() GatewayClassListMapOutput
	ToGatewayClassListMapOutputWithContext(context.Context) GatewayClassListMapOutput
}

type GatewayClassListMap map[string]GatewayClassListInput

func (GatewayClassListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayClassList)(nil)).Elem()
}

func (i GatewayClassListMap) ToGatewayClassListMapOutput() GatewayClassListMapOutput {
	return i.ToGatewayClassListMapOutputWithContext(context.Background())
}

func (i GatewayClassListMap) ToGatewayClassListMapOutputWithContext(ctx context.Context) GatewayClassListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayClassListMapOutput)
}

type GatewayClassListOutput struct{ *pulumi.OutputState }

func (GatewayClassListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayClassList)(nil)).Elem()
}

func (o GatewayClassListOutput) ToGatewayClassListOutput() GatewayClassListOutput {
	return o
}

func (o GatewayClassListOutput) ToGatewayClassListOutputWithContext(ctx context.Context) GatewayClassListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o GatewayClassListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayClassList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of gatewayclasses. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o GatewayClassListOutput) Items() GatewayClassTypeArrayOutput {
	return o.ApplyT(func(v *GatewayClassList) GatewayClassTypeArrayOutput { return v.Items }).(GatewayClassTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o GatewayClassListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayClassList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o GatewayClassListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *GatewayClassList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type GatewayClassListArrayOutput struct{ *pulumi.OutputState }

func (GatewayClassListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayClassList)(nil)).Elem()
}

func (o GatewayClassListArrayOutput) ToGatewayClassListArrayOutput() GatewayClassListArrayOutput {
	return o
}

func (o GatewayClassListArrayOutput) ToGatewayClassListArrayOutputWithContext(ctx context.Context) GatewayClassListArrayOutput {
	return o
}

func (o GatewayClassListArrayOutput) Index(i pulumi.IntInput) GatewayClassListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayClassList {
		return vs[0].([]*GatewayClassList)[vs[1].(int)]
	}).(GatewayClassListOutput)
}

type GatewayClassListMapOutput struct{ *pulumi.OutputState }

func (GatewayClassListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayClassList)(nil)).Elem()
}

func (o GatewayClassListMapOutput) ToGatewayClassListMapOutput() GatewayClassListMapOutput {
	return o
}

func (o GatewayClassListMapOutput) ToGatewayClassListMapOutputWithContext(ctx context.Context) GatewayClassListMapOutput {
	return o
}

func (o GatewayClassListMapOutput) MapIndex(k pulumi.StringInput) GatewayClassListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayClassList {
		return vs[0].(map[string]*GatewayClassList)[vs[1].(string)]
	}).(GatewayClassListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayClassListInput)(nil)).Elem(), &GatewayClassList{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayClassListArrayInput)(nil)).Elem(), GatewayClassListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayClassListMapInput)(nil)).Elem(), GatewayClassListMap{})
	pulumi.RegisterOutputType(GatewayClassListOutput{})
	pulumi.RegisterOutputType(GatewayClassListArrayOutput{})
	pulumi.RegisterOutputType(GatewayClassListMapOutput{})
}
