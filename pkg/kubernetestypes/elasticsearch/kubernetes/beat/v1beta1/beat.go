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

// Beat is the Schema for the Beats API.
type Beat struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	Spec     BeatSpecOutput          `pulumi:"spec"`
	Status   BeatStatusPtrOutput     `pulumi:"status"`
}

// NewBeat registers a new resource with the given unique name, arguments, and options.
func NewBeat(ctx *pulumi.Context,
	name string, args *BeatArgs, opts ...pulumi.ResourceOption) (*Beat, error) {
	if args == nil {
		args = &BeatArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("beat.k8s.elastic.co/v1beta1")
	args.Kind = pulumi.StringPtr("Beat")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Beat
	err := ctx.RegisterResource("kubernetes:beat.k8s.elastic.co/v1beta1:Beat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBeat gets an existing Beat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBeat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BeatState, opts ...pulumi.ResourceOption) (*Beat, error) {
	var resource Beat
	err := ctx.ReadResource("kubernetes:beat.k8s.elastic.co/v1beta1:Beat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Beat resources.
type beatState struct {
}

type BeatState struct {
}

func (BeatState) ElementType() reflect.Type {
	return reflect.TypeOf((*beatState)(nil)).Elem()
}

type beatArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *BeatSpec          `pulumi:"spec"`
}

// The set of arguments for constructing a Beat resource.
type BeatArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     BeatSpecPtrInput
}

func (BeatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*beatArgs)(nil)).Elem()
}

type BeatInput interface {
	pulumi.Input

	ToBeatOutput() BeatOutput
	ToBeatOutputWithContext(ctx context.Context) BeatOutput
}

func (*Beat) ElementType() reflect.Type {
	return reflect.TypeOf((**Beat)(nil)).Elem()
}

func (i *Beat) ToBeatOutput() BeatOutput {
	return i.ToBeatOutputWithContext(context.Background())
}

func (i *Beat) ToBeatOutputWithContext(ctx context.Context) BeatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BeatOutput)
}

// BeatArrayInput is an input type that accepts BeatArray and BeatArrayOutput values.
// You can construct a concrete instance of `BeatArrayInput` via:
//
//	BeatArray{ BeatArgs{...} }
type BeatArrayInput interface {
	pulumi.Input

	ToBeatArrayOutput() BeatArrayOutput
	ToBeatArrayOutputWithContext(context.Context) BeatArrayOutput
}

type BeatArray []BeatInput

func (BeatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Beat)(nil)).Elem()
}

func (i BeatArray) ToBeatArrayOutput() BeatArrayOutput {
	return i.ToBeatArrayOutputWithContext(context.Background())
}

func (i BeatArray) ToBeatArrayOutputWithContext(ctx context.Context) BeatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BeatArrayOutput)
}

// BeatMapInput is an input type that accepts BeatMap and BeatMapOutput values.
// You can construct a concrete instance of `BeatMapInput` via:
//
//	BeatMap{ "key": BeatArgs{...} }
type BeatMapInput interface {
	pulumi.Input

	ToBeatMapOutput() BeatMapOutput
	ToBeatMapOutputWithContext(context.Context) BeatMapOutput
}

type BeatMap map[string]BeatInput

func (BeatMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Beat)(nil)).Elem()
}

func (i BeatMap) ToBeatMapOutput() BeatMapOutput {
	return i.ToBeatMapOutputWithContext(context.Background())
}

func (i BeatMap) ToBeatMapOutputWithContext(ctx context.Context) BeatMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BeatMapOutput)
}

type BeatOutput struct{ *pulumi.OutputState }

func (BeatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Beat)(nil)).Elem()
}

func (o BeatOutput) ToBeatOutput() BeatOutput {
	return o
}

func (o BeatOutput) ToBeatOutputWithContext(ctx context.Context) BeatOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o BeatOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Beat) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o BeatOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Beat) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o BeatOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *Beat) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o BeatOutput) Spec() BeatSpecOutput {
	return o.ApplyT(func(v *Beat) BeatSpecOutput { return v.Spec }).(BeatSpecOutput)
}

func (o BeatOutput) Status() BeatStatusPtrOutput {
	return o.ApplyT(func(v *Beat) BeatStatusPtrOutput { return v.Status }).(BeatStatusPtrOutput)
}

type BeatArrayOutput struct{ *pulumi.OutputState }

func (BeatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Beat)(nil)).Elem()
}

func (o BeatArrayOutput) ToBeatArrayOutput() BeatArrayOutput {
	return o
}

func (o BeatArrayOutput) ToBeatArrayOutputWithContext(ctx context.Context) BeatArrayOutput {
	return o
}

func (o BeatArrayOutput) Index(i pulumi.IntInput) BeatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Beat {
		return vs[0].([]*Beat)[vs[1].(int)]
	}).(BeatOutput)
}

type BeatMapOutput struct{ *pulumi.OutputState }

func (BeatMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Beat)(nil)).Elem()
}

func (o BeatMapOutput) ToBeatMapOutput() BeatMapOutput {
	return o
}

func (o BeatMapOutput) ToBeatMapOutputWithContext(ctx context.Context) BeatMapOutput {
	return o
}

func (o BeatMapOutput) MapIndex(k pulumi.StringInput) BeatOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Beat {
		return vs[0].(map[string]*Beat)[vs[1].(string)]
	}).(BeatOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BeatInput)(nil)).Elem(), &Beat{})
	pulumi.RegisterInputType(reflect.TypeOf((*BeatArrayInput)(nil)).Elem(), BeatArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BeatMapInput)(nil)).Elem(), BeatMap{})
	pulumi.RegisterOutputType(BeatOutput{})
	pulumi.RegisterOutputType(BeatArrayOutput{})
	pulumi.RegisterOutputType(BeatMapOutput{})
}