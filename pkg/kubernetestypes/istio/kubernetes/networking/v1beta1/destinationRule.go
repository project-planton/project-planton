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

type DestinationRule struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput   `pulumi:"metadata"`
	Spec     DestinationRuleSpecOutput `pulumi:"spec"`
	Status   pulumi.MapOutput          `pulumi:"status"`
}

// NewDestinationRule registers a new resource with the given unique name, arguments, and options.
func NewDestinationRule(ctx *pulumi.Context,
	name string, args *DestinationRuleArgs, opts ...pulumi.ResourceOption) (*DestinationRule, error) {
	if args == nil {
		args = &DestinationRuleArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1beta1")
	args.Kind = pulumi.StringPtr("DestinationRule")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:networking.istio.io/v1:DestinationRule"),
		},
		{
			Type: pulumi.String("kubernetes:networking.istio.io/v1alpha3:DestinationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DestinationRule
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1beta1:DestinationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationRule gets an existing DestinationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationRuleState, opts ...pulumi.ResourceOption) (*DestinationRule, error) {
	var resource DestinationRule
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1beta1:DestinationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationRule resources.
type destinationRuleState struct {
}

type DestinationRuleState struct {
}

func (DestinationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationRuleState)(nil)).Elem()
}

type destinationRuleArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta   `pulumi:"metadata"`
	Spec     *DestinationRuleSpec `pulumi:"spec"`
}

// The set of arguments for constructing a DestinationRule resource.
type DestinationRuleArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     DestinationRuleSpecPtrInput
}

func (DestinationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationRuleArgs)(nil)).Elem()
}

type DestinationRuleInput interface {
	pulumi.Input

	ToDestinationRuleOutput() DestinationRuleOutput
	ToDestinationRuleOutputWithContext(ctx context.Context) DestinationRuleOutput
}

func (*DestinationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationRule)(nil)).Elem()
}

func (i *DestinationRule) ToDestinationRuleOutput() DestinationRuleOutput {
	return i.ToDestinationRuleOutputWithContext(context.Background())
}

func (i *DestinationRule) ToDestinationRuleOutputWithContext(ctx context.Context) DestinationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationRuleOutput)
}

// DestinationRuleArrayInput is an input type that accepts DestinationRuleArray and DestinationRuleArrayOutput values.
// You can construct a concrete instance of `DestinationRuleArrayInput` via:
//
//	DestinationRuleArray{ DestinationRuleArgs{...} }
type DestinationRuleArrayInput interface {
	pulumi.Input

	ToDestinationRuleArrayOutput() DestinationRuleArrayOutput
	ToDestinationRuleArrayOutputWithContext(context.Context) DestinationRuleArrayOutput
}

type DestinationRuleArray []DestinationRuleInput

func (DestinationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationRule)(nil)).Elem()
}

func (i DestinationRuleArray) ToDestinationRuleArrayOutput() DestinationRuleArrayOutput {
	return i.ToDestinationRuleArrayOutputWithContext(context.Background())
}

func (i DestinationRuleArray) ToDestinationRuleArrayOutputWithContext(ctx context.Context) DestinationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationRuleArrayOutput)
}

// DestinationRuleMapInput is an input type that accepts DestinationRuleMap and DestinationRuleMapOutput values.
// You can construct a concrete instance of `DestinationRuleMapInput` via:
//
//	DestinationRuleMap{ "key": DestinationRuleArgs{...} }
type DestinationRuleMapInput interface {
	pulumi.Input

	ToDestinationRuleMapOutput() DestinationRuleMapOutput
	ToDestinationRuleMapOutputWithContext(context.Context) DestinationRuleMapOutput
}

type DestinationRuleMap map[string]DestinationRuleInput

func (DestinationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationRule)(nil)).Elem()
}

func (i DestinationRuleMap) ToDestinationRuleMapOutput() DestinationRuleMapOutput {
	return i.ToDestinationRuleMapOutputWithContext(context.Background())
}

func (i DestinationRuleMap) ToDestinationRuleMapOutputWithContext(ctx context.Context) DestinationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationRuleMapOutput)
}

type DestinationRuleOutput struct{ *pulumi.OutputState }

func (DestinationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationRule)(nil)).Elem()
}

func (o DestinationRuleOutput) ToDestinationRuleOutput() DestinationRuleOutput {
	return o
}

func (o DestinationRuleOutput) ToDestinationRuleOutputWithContext(ctx context.Context) DestinationRuleOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o DestinationRuleOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationRule) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o DestinationRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o DestinationRuleOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *DestinationRule) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o DestinationRuleOutput) Spec() DestinationRuleSpecOutput {
	return o.ApplyT(func(v *DestinationRule) DestinationRuleSpecOutput { return v.Spec }).(DestinationRuleSpecOutput)
}

func (o DestinationRuleOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *DestinationRule) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

type DestinationRuleArrayOutput struct{ *pulumi.OutputState }

func (DestinationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationRule)(nil)).Elem()
}

func (o DestinationRuleArrayOutput) ToDestinationRuleArrayOutput() DestinationRuleArrayOutput {
	return o
}

func (o DestinationRuleArrayOutput) ToDestinationRuleArrayOutputWithContext(ctx context.Context) DestinationRuleArrayOutput {
	return o
}

func (o DestinationRuleArrayOutput) Index(i pulumi.IntInput) DestinationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationRule {
		return vs[0].([]*DestinationRule)[vs[1].(int)]
	}).(DestinationRuleOutput)
}

type DestinationRuleMapOutput struct{ *pulumi.OutputState }

func (DestinationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationRule)(nil)).Elem()
}

func (o DestinationRuleMapOutput) ToDestinationRuleMapOutput() DestinationRuleMapOutput {
	return o
}

func (o DestinationRuleMapOutput) ToDestinationRuleMapOutputWithContext(ctx context.Context) DestinationRuleMapOutput {
	return o
}

func (o DestinationRuleMapOutput) MapIndex(k pulumi.StringInput) DestinationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationRule {
		return vs[0].(map[string]*DestinationRule)[vs[1].(string)]
	}).(DestinationRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationRuleInput)(nil)).Elem(), &DestinationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationRuleArrayInput)(nil)).Elem(), DestinationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationRuleMapInput)(nil)).Elem(), DestinationRuleMap{})
	pulumi.RegisterOutputType(DestinationRuleOutput{})
	pulumi.RegisterOutputType(DestinationRuleArrayOutput{})
	pulumi.RegisterOutputType(DestinationRuleMapOutput{})
}