// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ReferenceGrantList is a list of ReferenceGrant
type ReferenceGrantList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of referencegrants. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ReferenceGrantTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewReferenceGrantList registers a new resource with the given unique name, arguments, and options.
func NewReferenceGrantList(ctx *pulumi.Context,
	name string, args *ReferenceGrantListArgs, opts ...pulumi.ResourceOption) (*ReferenceGrantList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("gateway.networking.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("ReferenceGrantList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReferenceGrantList
	err := ctx.RegisterResource("kubernetes:gateway.networking.k8s.io/v1alpha2:ReferenceGrantList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReferenceGrantList gets an existing ReferenceGrantList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReferenceGrantList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReferenceGrantListState, opts ...pulumi.ResourceOption) (*ReferenceGrantList, error) {
	var resource ReferenceGrantList
	err := ctx.ReadResource("kubernetes:gateway.networking.k8s.io/v1alpha2:ReferenceGrantList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReferenceGrantList resources.
type referenceGrantListState struct {
}

type ReferenceGrantListState struct {
}

func (ReferenceGrantListState) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceGrantListState)(nil)).Elem()
}

type referenceGrantListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of referencegrants. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ReferenceGrantType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ReferenceGrantList resource.
type ReferenceGrantListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of referencegrants. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ReferenceGrantTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ReferenceGrantListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceGrantListArgs)(nil)).Elem()
}

type ReferenceGrantListInput interface {
	pulumi.Input

	ToReferenceGrantListOutput() ReferenceGrantListOutput
	ToReferenceGrantListOutputWithContext(ctx context.Context) ReferenceGrantListOutput
}

func (*ReferenceGrantList) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceGrantList)(nil)).Elem()
}

func (i *ReferenceGrantList) ToReferenceGrantListOutput() ReferenceGrantListOutput {
	return i.ToReferenceGrantListOutputWithContext(context.Background())
}

func (i *ReferenceGrantList) ToReferenceGrantListOutputWithContext(ctx context.Context) ReferenceGrantListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceGrantListOutput)
}

// ReferenceGrantListArrayInput is an input type that accepts ReferenceGrantListArray and ReferenceGrantListArrayOutput values.
// You can construct a concrete instance of `ReferenceGrantListArrayInput` via:
//
//	ReferenceGrantListArray{ ReferenceGrantListArgs{...} }
type ReferenceGrantListArrayInput interface {
	pulumi.Input

	ToReferenceGrantListArrayOutput() ReferenceGrantListArrayOutput
	ToReferenceGrantListArrayOutputWithContext(context.Context) ReferenceGrantListArrayOutput
}

type ReferenceGrantListArray []ReferenceGrantListInput

func (ReferenceGrantListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReferenceGrantList)(nil)).Elem()
}

func (i ReferenceGrantListArray) ToReferenceGrantListArrayOutput() ReferenceGrantListArrayOutput {
	return i.ToReferenceGrantListArrayOutputWithContext(context.Background())
}

func (i ReferenceGrantListArray) ToReferenceGrantListArrayOutputWithContext(ctx context.Context) ReferenceGrantListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceGrantListArrayOutput)
}

// ReferenceGrantListMapInput is an input type that accepts ReferenceGrantListMap and ReferenceGrantListMapOutput values.
// You can construct a concrete instance of `ReferenceGrantListMapInput` via:
//
//	ReferenceGrantListMap{ "key": ReferenceGrantListArgs{...} }
type ReferenceGrantListMapInput interface {
	pulumi.Input

	ToReferenceGrantListMapOutput() ReferenceGrantListMapOutput
	ToReferenceGrantListMapOutputWithContext(context.Context) ReferenceGrantListMapOutput
}

type ReferenceGrantListMap map[string]ReferenceGrantListInput

func (ReferenceGrantListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReferenceGrantList)(nil)).Elem()
}

func (i ReferenceGrantListMap) ToReferenceGrantListMapOutput() ReferenceGrantListMapOutput {
	return i.ToReferenceGrantListMapOutputWithContext(context.Background())
}

func (i ReferenceGrantListMap) ToReferenceGrantListMapOutputWithContext(ctx context.Context) ReferenceGrantListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceGrantListMapOutput)
}

type ReferenceGrantListOutput struct{ *pulumi.OutputState }

func (ReferenceGrantListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceGrantList)(nil)).Elem()
}

func (o ReferenceGrantListOutput) ToReferenceGrantListOutput() ReferenceGrantListOutput {
	return o
}

func (o ReferenceGrantListOutput) ToReferenceGrantListOutputWithContext(ctx context.Context) ReferenceGrantListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ReferenceGrantListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceGrantList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of referencegrants. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ReferenceGrantListOutput) Items() ReferenceGrantTypeArrayOutput {
	return o.ApplyT(func(v *ReferenceGrantList) ReferenceGrantTypeArrayOutput { return v.Items }).(ReferenceGrantTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReferenceGrantListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceGrantList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReferenceGrantListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ReferenceGrantList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ReferenceGrantListArrayOutput struct{ *pulumi.OutputState }

func (ReferenceGrantListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReferenceGrantList)(nil)).Elem()
}

func (o ReferenceGrantListArrayOutput) ToReferenceGrantListArrayOutput() ReferenceGrantListArrayOutput {
	return o
}

func (o ReferenceGrantListArrayOutput) ToReferenceGrantListArrayOutputWithContext(ctx context.Context) ReferenceGrantListArrayOutput {
	return o
}

func (o ReferenceGrantListArrayOutput) Index(i pulumi.IntInput) ReferenceGrantListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReferenceGrantList {
		return vs[0].([]*ReferenceGrantList)[vs[1].(int)]
	}).(ReferenceGrantListOutput)
}

type ReferenceGrantListMapOutput struct{ *pulumi.OutputState }

func (ReferenceGrantListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReferenceGrantList)(nil)).Elem()
}

func (o ReferenceGrantListMapOutput) ToReferenceGrantListMapOutput() ReferenceGrantListMapOutput {
	return o
}

func (o ReferenceGrantListMapOutput) ToReferenceGrantListMapOutputWithContext(ctx context.Context) ReferenceGrantListMapOutput {
	return o
}

func (o ReferenceGrantListMapOutput) MapIndex(k pulumi.StringInput) ReferenceGrantListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReferenceGrantList {
		return vs[0].(map[string]*ReferenceGrantList)[vs[1].(string)]
	}).(ReferenceGrantListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceGrantListInput)(nil)).Elem(), &ReferenceGrantList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceGrantListArrayInput)(nil)).Elem(), ReferenceGrantListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceGrantListMapInput)(nil)).Elem(), ReferenceGrantListMap{})
	pulumi.RegisterOutputType(ReferenceGrantListOutput{})
	pulumi.RegisterOutputType(ReferenceGrantListArrayOutput{})
	pulumi.RegisterOutputType(ReferenceGrantListMapOutput{})
}