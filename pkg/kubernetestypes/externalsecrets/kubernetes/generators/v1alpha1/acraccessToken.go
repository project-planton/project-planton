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

// ACRAccessToken returns a Azure Container Registry token
// that can be used for pushing/pulling images.
// Note: by default it will return an ACR Refresh Token with full access
// (depending on the identity).
// This can be scoped down to the repository level using .spec.scope.
// In case scope is defined it will return an ACR Access Token.
//
// See docs: https://github.com/Azure/acr/blob/main/docs/AAD-OAuth.md
type ACRAccessToken struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput  `pulumi:"metadata"`
	Spec     ACRAccessTokenSpecOutput `pulumi:"spec"`
}

// NewACRAccessToken registers a new resource with the given unique name, arguments, and options.
func NewACRAccessToken(ctx *pulumi.Context,
	name string, args *ACRAccessTokenArgs, opts ...pulumi.ResourceOption) (*ACRAccessToken, error) {
	if args == nil {
		args = &ACRAccessTokenArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ACRAccessToken")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ACRAccessToken
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:ACRAccessToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetACRAccessToken gets an existing ACRAccessToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetACRAccessToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ACRAccessTokenState, opts ...pulumi.ResourceOption) (*ACRAccessToken, error) {
	var resource ACRAccessToken
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:ACRAccessToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ACRAccessToken resources.
type acraccessTokenState struct {
}

type ACRAccessTokenState struct {
}

func (ACRAccessTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*acraccessTokenState)(nil)).Elem()
}

type acraccessTokenArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta  `pulumi:"metadata"`
	Spec     *ACRAccessTokenSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ACRAccessToken resource.
type ACRAccessTokenArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     ACRAccessTokenSpecPtrInput
}

func (ACRAccessTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*acraccessTokenArgs)(nil)).Elem()
}

type ACRAccessTokenInput interface {
	pulumi.Input

	ToACRAccessTokenOutput() ACRAccessTokenOutput
	ToACRAccessTokenOutputWithContext(ctx context.Context) ACRAccessTokenOutput
}

func (*ACRAccessToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ACRAccessToken)(nil)).Elem()
}

func (i *ACRAccessToken) ToACRAccessTokenOutput() ACRAccessTokenOutput {
	return i.ToACRAccessTokenOutputWithContext(context.Background())
}

func (i *ACRAccessToken) ToACRAccessTokenOutputWithContext(ctx context.Context) ACRAccessTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACRAccessTokenOutput)
}

// ACRAccessTokenArrayInput is an input type that accepts ACRAccessTokenArray and ACRAccessTokenArrayOutput values.
// You can construct a concrete instance of `ACRAccessTokenArrayInput` via:
//
//	ACRAccessTokenArray{ ACRAccessTokenArgs{...} }
type ACRAccessTokenArrayInput interface {
	pulumi.Input

	ToACRAccessTokenArrayOutput() ACRAccessTokenArrayOutput
	ToACRAccessTokenArrayOutputWithContext(context.Context) ACRAccessTokenArrayOutput
}

type ACRAccessTokenArray []ACRAccessTokenInput

func (ACRAccessTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ACRAccessToken)(nil)).Elem()
}

func (i ACRAccessTokenArray) ToACRAccessTokenArrayOutput() ACRAccessTokenArrayOutput {
	return i.ToACRAccessTokenArrayOutputWithContext(context.Background())
}

func (i ACRAccessTokenArray) ToACRAccessTokenArrayOutputWithContext(ctx context.Context) ACRAccessTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACRAccessTokenArrayOutput)
}

// ACRAccessTokenMapInput is an input type that accepts ACRAccessTokenMap and ACRAccessTokenMapOutput values.
// You can construct a concrete instance of `ACRAccessTokenMapInput` via:
//
//	ACRAccessTokenMap{ "key": ACRAccessTokenArgs{...} }
type ACRAccessTokenMapInput interface {
	pulumi.Input

	ToACRAccessTokenMapOutput() ACRAccessTokenMapOutput
	ToACRAccessTokenMapOutputWithContext(context.Context) ACRAccessTokenMapOutput
}

type ACRAccessTokenMap map[string]ACRAccessTokenInput

func (ACRAccessTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ACRAccessToken)(nil)).Elem()
}

func (i ACRAccessTokenMap) ToACRAccessTokenMapOutput() ACRAccessTokenMapOutput {
	return i.ToACRAccessTokenMapOutputWithContext(context.Background())
}

func (i ACRAccessTokenMap) ToACRAccessTokenMapOutputWithContext(ctx context.Context) ACRAccessTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACRAccessTokenMapOutput)
}

type ACRAccessTokenOutput struct{ *pulumi.OutputState }

func (ACRAccessTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACRAccessToken)(nil)).Elem()
}

func (o ACRAccessTokenOutput) ToACRAccessTokenOutput() ACRAccessTokenOutput {
	return o
}

func (o ACRAccessTokenOutput) ToACRAccessTokenOutputWithContext(ctx context.Context) ACRAccessTokenOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ACRAccessTokenOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ACRAccessToken) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ACRAccessTokenOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ACRAccessToken) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ACRAccessTokenOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ACRAccessToken) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o ACRAccessTokenOutput) Spec() ACRAccessTokenSpecOutput {
	return o.ApplyT(func(v *ACRAccessToken) ACRAccessTokenSpecOutput { return v.Spec }).(ACRAccessTokenSpecOutput)
}

type ACRAccessTokenArrayOutput struct{ *pulumi.OutputState }

func (ACRAccessTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ACRAccessToken)(nil)).Elem()
}

func (o ACRAccessTokenArrayOutput) ToACRAccessTokenArrayOutput() ACRAccessTokenArrayOutput {
	return o
}

func (o ACRAccessTokenArrayOutput) ToACRAccessTokenArrayOutputWithContext(ctx context.Context) ACRAccessTokenArrayOutput {
	return o
}

func (o ACRAccessTokenArrayOutput) Index(i pulumi.IntInput) ACRAccessTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ACRAccessToken {
		return vs[0].([]*ACRAccessToken)[vs[1].(int)]
	}).(ACRAccessTokenOutput)
}

type ACRAccessTokenMapOutput struct{ *pulumi.OutputState }

func (ACRAccessTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ACRAccessToken)(nil)).Elem()
}

func (o ACRAccessTokenMapOutput) ToACRAccessTokenMapOutput() ACRAccessTokenMapOutput {
	return o
}

func (o ACRAccessTokenMapOutput) ToACRAccessTokenMapOutputWithContext(ctx context.Context) ACRAccessTokenMapOutput {
	return o
}

func (o ACRAccessTokenMapOutput) MapIndex(k pulumi.StringInput) ACRAccessTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ACRAccessToken {
		return vs[0].(map[string]*ACRAccessToken)[vs[1].(string)]
	}).(ACRAccessTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ACRAccessTokenInput)(nil)).Elem(), &ACRAccessToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACRAccessTokenArrayInput)(nil)).Elem(), ACRAccessTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACRAccessTokenMapInput)(nil)).Elem(), ACRAccessTokenMap{})
	pulumi.RegisterOutputType(ACRAccessTokenOutput{})
	pulumi.RegisterOutputType(ACRAccessTokenArrayOutput{})
	pulumi.RegisterOutputType(ACRAccessTokenMapOutput{})
}