// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

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
// Order is a type to represent an Order with an ACME server
type OrderPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     OrderSpecPatchPtrOutput         `pulumi:"spec"`
	Status   OrderStatusPatchPtrOutput       `pulumi:"status"`
}

// NewOrderPatch registers a new resource with the given unique name, arguments, and options.
func NewOrderPatch(ctx *pulumi.Context,
	name string, args *OrderPatchArgs, opts ...pulumi.ResourceOption) (*OrderPatch, error) {
	if args == nil {
		args = &OrderPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("acme.cert-manager.io/v1")
	args.Kind = pulumi.StringPtr("Order")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OrderPatch
	err := ctx.RegisterResource("kubernetes:acme.cert-manager.io/v1:OrderPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrderPatch gets an existing OrderPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrderPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderPatchState, opts ...pulumi.ResourceOption) (*OrderPatch, error) {
	var resource OrderPatch
	err := ctx.ReadResource("kubernetes:acme.cert-manager.io/v1:OrderPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrderPatch resources.
type orderPatchState struct {
}

type OrderPatchState struct {
}

func (OrderPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderPatchState)(nil)).Elem()
}

type orderPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *OrderSpecPatch         `pulumi:"spec"`
}

// The set of arguments for constructing a OrderPatch resource.
type OrderPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     OrderSpecPatchPtrInput
}

func (OrderPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderPatchArgs)(nil)).Elem()
}

type OrderPatchInput interface {
	pulumi.Input

	ToOrderPatchOutput() OrderPatchOutput
	ToOrderPatchOutputWithContext(ctx context.Context) OrderPatchOutput
}

func (*OrderPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderPatch)(nil)).Elem()
}

func (i *OrderPatch) ToOrderPatchOutput() OrderPatchOutput {
	return i.ToOrderPatchOutputWithContext(context.Background())
}

func (i *OrderPatch) ToOrderPatchOutputWithContext(ctx context.Context) OrderPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderPatchOutput)
}

// OrderPatchArrayInput is an input type that accepts OrderPatchArray and OrderPatchArrayOutput values.
// You can construct a concrete instance of `OrderPatchArrayInput` via:
//
//	OrderPatchArray{ OrderPatchArgs{...} }
type OrderPatchArrayInput interface {
	pulumi.Input

	ToOrderPatchArrayOutput() OrderPatchArrayOutput
	ToOrderPatchArrayOutputWithContext(context.Context) OrderPatchArrayOutput
}

type OrderPatchArray []OrderPatchInput

func (OrderPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrderPatch)(nil)).Elem()
}

func (i OrderPatchArray) ToOrderPatchArrayOutput() OrderPatchArrayOutput {
	return i.ToOrderPatchArrayOutputWithContext(context.Background())
}

func (i OrderPatchArray) ToOrderPatchArrayOutputWithContext(ctx context.Context) OrderPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderPatchArrayOutput)
}

// OrderPatchMapInput is an input type that accepts OrderPatchMap and OrderPatchMapOutput values.
// You can construct a concrete instance of `OrderPatchMapInput` via:
//
//	OrderPatchMap{ "key": OrderPatchArgs{...} }
type OrderPatchMapInput interface {
	pulumi.Input

	ToOrderPatchMapOutput() OrderPatchMapOutput
	ToOrderPatchMapOutputWithContext(context.Context) OrderPatchMapOutput
}

type OrderPatchMap map[string]OrderPatchInput

func (OrderPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrderPatch)(nil)).Elem()
}

func (i OrderPatchMap) ToOrderPatchMapOutput() OrderPatchMapOutput {
	return i.ToOrderPatchMapOutputWithContext(context.Background())
}

func (i OrderPatchMap) ToOrderPatchMapOutputWithContext(ctx context.Context) OrderPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderPatchMapOutput)
}

type OrderPatchOutput struct{ *pulumi.OutputState }

func (OrderPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderPatch)(nil)).Elem()
}

func (o OrderPatchOutput) ToOrderPatchOutput() OrderPatchOutput {
	return o
}

func (o OrderPatchOutput) ToOrderPatchOutputWithContext(ctx context.Context) OrderPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o OrderPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o OrderPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o OrderPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *OrderPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o OrderPatchOutput) Spec() OrderSpecPatchPtrOutput {
	return o.ApplyT(func(v *OrderPatch) OrderSpecPatchPtrOutput { return v.Spec }).(OrderSpecPatchPtrOutput)
}

func (o OrderPatchOutput) Status() OrderStatusPatchPtrOutput {
	return o.ApplyT(func(v *OrderPatch) OrderStatusPatchPtrOutput { return v.Status }).(OrderStatusPatchPtrOutput)
}

type OrderPatchArrayOutput struct{ *pulumi.OutputState }

func (OrderPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrderPatch)(nil)).Elem()
}

func (o OrderPatchArrayOutput) ToOrderPatchArrayOutput() OrderPatchArrayOutput {
	return o
}

func (o OrderPatchArrayOutput) ToOrderPatchArrayOutputWithContext(ctx context.Context) OrderPatchArrayOutput {
	return o
}

func (o OrderPatchArrayOutput) Index(i pulumi.IntInput) OrderPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrderPatch {
		return vs[0].([]*OrderPatch)[vs[1].(int)]
	}).(OrderPatchOutput)
}

type OrderPatchMapOutput struct{ *pulumi.OutputState }

func (OrderPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrderPatch)(nil)).Elem()
}

func (o OrderPatchMapOutput) ToOrderPatchMapOutput() OrderPatchMapOutput {
	return o
}

func (o OrderPatchMapOutput) ToOrderPatchMapOutputWithContext(ctx context.Context) OrderPatchMapOutput {
	return o
}

func (o OrderPatchMapOutput) MapIndex(k pulumi.StringInput) OrderPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrderPatch {
		return vs[0].(map[string]*OrderPatch)[vs[1].(string)]
	}).(OrderPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrderPatchInput)(nil)).Elem(), &OrderPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderPatchArrayInput)(nil)).Elem(), OrderPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderPatchMapInput)(nil)).Elem(), OrderPatchMap{})
	pulumi.RegisterOutputType(OrderPatchOutput{})
	pulumi.RegisterOutputType(OrderPatchArrayOutput{})
	pulumi.RegisterOutputType(OrderPatchMapOutput{})
}
