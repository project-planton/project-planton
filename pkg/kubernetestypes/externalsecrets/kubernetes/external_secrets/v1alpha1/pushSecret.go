// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PushSecret struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput   `pulumi:"metadata"`
	Spec     PushSecretSpecOutput      `pulumi:"spec"`
	Status   PushSecretStatusPtrOutput `pulumi:"status"`
}

// NewPushSecret registers a new resource with the given unique name, arguments, and options.
func NewPushSecret(ctx *pulumi.Context,
	name string, args *PushSecretArgs, opts ...pulumi.ResourceOption) (*PushSecret, error) {
	if args == nil {
		args = &PushSecretArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("PushSecret")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PushSecret
	err := ctx.RegisterResource("kubernetes:external-secrets.io/v1alpha1:PushSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPushSecret gets an existing PushSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPushSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PushSecretState, opts ...pulumi.ResourceOption) (*PushSecret, error) {
	var resource PushSecret
	err := ctx.ReadResource("kubernetes:external-secrets.io/v1alpha1:PushSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PushSecret resources.
type pushSecretState struct {
}

type PushSecretState struct {
}

func (PushSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*pushSecretState)(nil)).Elem()
}

type pushSecretArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *PushSecretSpec    `pulumi:"spec"`
}

// The set of arguments for constructing a PushSecret resource.
type PushSecretArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     PushSecretSpecPtrInput
}

func (PushSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pushSecretArgs)(nil)).Elem()
}

type PushSecretInput interface {
	pulumi.Input

	ToPushSecretOutput() PushSecretOutput
	ToPushSecretOutputWithContext(ctx context.Context) PushSecretOutput
}

func (*PushSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSecret)(nil)).Elem()
}

func (i *PushSecret) ToPushSecretOutput() PushSecretOutput {
	return i.ToPushSecretOutputWithContext(context.Background())
}

func (i *PushSecret) ToPushSecretOutputWithContext(ctx context.Context) PushSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSecretOutput)
}

// PushSecretArrayInput is an input type that accepts PushSecretArray and PushSecretArrayOutput values.
// You can construct a concrete instance of `PushSecretArrayInput` via:
//
//	PushSecretArray{ PushSecretArgs{...} }
type PushSecretArrayInput interface {
	pulumi.Input

	ToPushSecretArrayOutput() PushSecretArrayOutput
	ToPushSecretArrayOutputWithContext(context.Context) PushSecretArrayOutput
}

type PushSecretArray []PushSecretInput

func (PushSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PushSecret)(nil)).Elem()
}

func (i PushSecretArray) ToPushSecretArrayOutput() PushSecretArrayOutput {
	return i.ToPushSecretArrayOutputWithContext(context.Background())
}

func (i PushSecretArray) ToPushSecretArrayOutputWithContext(ctx context.Context) PushSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSecretArrayOutput)
}

// PushSecretMapInput is an input type that accepts PushSecretMap and PushSecretMapOutput values.
// You can construct a concrete instance of `PushSecretMapInput` via:
//
//	PushSecretMap{ "key": PushSecretArgs{...} }
type PushSecretMapInput interface {
	pulumi.Input

	ToPushSecretMapOutput() PushSecretMapOutput
	ToPushSecretMapOutputWithContext(context.Context) PushSecretMapOutput
}

type PushSecretMap map[string]PushSecretInput

func (PushSecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PushSecret)(nil)).Elem()
}

func (i PushSecretMap) ToPushSecretMapOutput() PushSecretMapOutput {
	return i.ToPushSecretMapOutputWithContext(context.Background())
}

func (i PushSecretMap) ToPushSecretMapOutputWithContext(ctx context.Context) PushSecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSecretMapOutput)
}

type PushSecretOutput struct{ *pulumi.OutputState }

func (PushSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSecret)(nil)).Elem()
}

func (o PushSecretOutput) ToPushSecretOutput() PushSecretOutput {
	return o
}

func (o PushSecretOutput) ToPushSecretOutputWithContext(ctx context.Context) PushSecretOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PushSecretOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PushSecret) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PushSecretOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PushSecret) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PushSecretOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *PushSecret) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o PushSecretOutput) Spec() PushSecretSpecOutput {
	return o.ApplyT(func(v *PushSecret) PushSecretSpecOutput { return v.Spec }).(PushSecretSpecOutput)
}

func (o PushSecretOutput) Status() PushSecretStatusPtrOutput {
	return o.ApplyT(func(v *PushSecret) PushSecretStatusPtrOutput { return v.Status }).(PushSecretStatusPtrOutput)
}

type PushSecretArrayOutput struct{ *pulumi.OutputState }

func (PushSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PushSecret)(nil)).Elem()
}

func (o PushSecretArrayOutput) ToPushSecretArrayOutput() PushSecretArrayOutput {
	return o
}

func (o PushSecretArrayOutput) ToPushSecretArrayOutputWithContext(ctx context.Context) PushSecretArrayOutput {
	return o
}

func (o PushSecretArrayOutput) Index(i pulumi.IntInput) PushSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PushSecret {
		return vs[0].([]*PushSecret)[vs[1].(int)]
	}).(PushSecretOutput)
}

type PushSecretMapOutput struct{ *pulumi.OutputState }

func (PushSecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PushSecret)(nil)).Elem()
}

func (o PushSecretMapOutput) ToPushSecretMapOutput() PushSecretMapOutput {
	return o
}

func (o PushSecretMapOutput) ToPushSecretMapOutputWithContext(ctx context.Context) PushSecretMapOutput {
	return o
}

func (o PushSecretMapOutput) MapIndex(k pulumi.StringInput) PushSecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PushSecret {
		return vs[0].(map[string]*PushSecret)[vs[1].(string)]
	}).(PushSecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PushSecretInput)(nil)).Elem(), &PushSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*PushSecretArrayInput)(nil)).Elem(), PushSecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PushSecretMapInput)(nil)).Elem(), PushSecretMap{})
	pulumi.RegisterOutputType(PushSecretOutput{})
	pulumi.RegisterOutputType(PushSecretArrayOutput{})
	pulumi.RegisterOutputType(PushSecretMapOutput{})
}