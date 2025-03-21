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

// ECRAuthorizationTokenSpec uses the GetAuthorizationToken API to retrieve an
// authorization token.
// The authorization token is valid for 12 hours.
// The authorizationToken returned is a base64 encoded string that can be decoded
// and used in a docker login command to authenticate to a registry.
// For more information, see Registry authentication (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth) in the Amazon Elastic Container Registry User Guide.
type ECRAuthorizationToken struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput         `pulumi:"metadata"`
	Spec     ECRAuthorizationTokenSpecOutput `pulumi:"spec"`
}

// NewECRAuthorizationToken registers a new resource with the given unique name, arguments, and options.
func NewECRAuthorizationToken(ctx *pulumi.Context,
	name string, args *ECRAuthorizationTokenArgs, opts ...pulumi.ResourceOption) (*ECRAuthorizationToken, error) {
	if args == nil {
		args = &ECRAuthorizationTokenArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ECRAuthorizationToken")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ECRAuthorizationToken
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:ECRAuthorizationToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetECRAuthorizationToken gets an existing ECRAuthorizationToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetECRAuthorizationToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ECRAuthorizationTokenState, opts ...pulumi.ResourceOption) (*ECRAuthorizationToken, error) {
	var resource ECRAuthorizationToken
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:ECRAuthorizationToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ECRAuthorizationToken resources.
type ecrauthorizationTokenState struct {
}

type ECRAuthorizationTokenState struct {
}

func (ECRAuthorizationTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecrauthorizationTokenState)(nil)).Elem()
}

type ecrauthorizationTokenArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta         `pulumi:"metadata"`
	Spec     *ECRAuthorizationTokenSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ECRAuthorizationToken resource.
type ECRAuthorizationTokenArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     ECRAuthorizationTokenSpecPtrInput
}

func (ECRAuthorizationTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecrauthorizationTokenArgs)(nil)).Elem()
}

type ECRAuthorizationTokenInput interface {
	pulumi.Input

	ToECRAuthorizationTokenOutput() ECRAuthorizationTokenOutput
	ToECRAuthorizationTokenOutputWithContext(ctx context.Context) ECRAuthorizationTokenOutput
}

func (*ECRAuthorizationToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ECRAuthorizationToken)(nil)).Elem()
}

func (i *ECRAuthorizationToken) ToECRAuthorizationTokenOutput() ECRAuthorizationTokenOutput {
	return i.ToECRAuthorizationTokenOutputWithContext(context.Background())
}

func (i *ECRAuthorizationToken) ToECRAuthorizationTokenOutputWithContext(ctx context.Context) ECRAuthorizationTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECRAuthorizationTokenOutput)
}

// ECRAuthorizationTokenArrayInput is an input type that accepts ECRAuthorizationTokenArray and ECRAuthorizationTokenArrayOutput values.
// You can construct a concrete instance of `ECRAuthorizationTokenArrayInput` via:
//
//	ECRAuthorizationTokenArray{ ECRAuthorizationTokenArgs{...} }
type ECRAuthorizationTokenArrayInput interface {
	pulumi.Input

	ToECRAuthorizationTokenArrayOutput() ECRAuthorizationTokenArrayOutput
	ToECRAuthorizationTokenArrayOutputWithContext(context.Context) ECRAuthorizationTokenArrayOutput
}

type ECRAuthorizationTokenArray []ECRAuthorizationTokenInput

func (ECRAuthorizationTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ECRAuthorizationToken)(nil)).Elem()
}

func (i ECRAuthorizationTokenArray) ToECRAuthorizationTokenArrayOutput() ECRAuthorizationTokenArrayOutput {
	return i.ToECRAuthorizationTokenArrayOutputWithContext(context.Background())
}

func (i ECRAuthorizationTokenArray) ToECRAuthorizationTokenArrayOutputWithContext(ctx context.Context) ECRAuthorizationTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECRAuthorizationTokenArrayOutput)
}

// ECRAuthorizationTokenMapInput is an input type that accepts ECRAuthorizationTokenMap and ECRAuthorizationTokenMapOutput values.
// You can construct a concrete instance of `ECRAuthorizationTokenMapInput` via:
//
//	ECRAuthorizationTokenMap{ "key": ECRAuthorizationTokenArgs{...} }
type ECRAuthorizationTokenMapInput interface {
	pulumi.Input

	ToECRAuthorizationTokenMapOutput() ECRAuthorizationTokenMapOutput
	ToECRAuthorizationTokenMapOutputWithContext(context.Context) ECRAuthorizationTokenMapOutput
}

type ECRAuthorizationTokenMap map[string]ECRAuthorizationTokenInput

func (ECRAuthorizationTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ECRAuthorizationToken)(nil)).Elem()
}

func (i ECRAuthorizationTokenMap) ToECRAuthorizationTokenMapOutput() ECRAuthorizationTokenMapOutput {
	return i.ToECRAuthorizationTokenMapOutputWithContext(context.Background())
}

func (i ECRAuthorizationTokenMap) ToECRAuthorizationTokenMapOutputWithContext(ctx context.Context) ECRAuthorizationTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECRAuthorizationTokenMapOutput)
}

type ECRAuthorizationTokenOutput struct{ *pulumi.OutputState }

func (ECRAuthorizationTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ECRAuthorizationToken)(nil)).Elem()
}

func (o ECRAuthorizationTokenOutput) ToECRAuthorizationTokenOutput() ECRAuthorizationTokenOutput {
	return o
}

func (o ECRAuthorizationTokenOutput) ToECRAuthorizationTokenOutputWithContext(ctx context.Context) ECRAuthorizationTokenOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ECRAuthorizationTokenOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ECRAuthorizationToken) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ECRAuthorizationTokenOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ECRAuthorizationToken) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ECRAuthorizationTokenOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ECRAuthorizationToken) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o ECRAuthorizationTokenOutput) Spec() ECRAuthorizationTokenSpecOutput {
	return o.ApplyT(func(v *ECRAuthorizationToken) ECRAuthorizationTokenSpecOutput { return v.Spec }).(ECRAuthorizationTokenSpecOutput)
}

type ECRAuthorizationTokenArrayOutput struct{ *pulumi.OutputState }

func (ECRAuthorizationTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ECRAuthorizationToken)(nil)).Elem()
}

func (o ECRAuthorizationTokenArrayOutput) ToECRAuthorizationTokenArrayOutput() ECRAuthorizationTokenArrayOutput {
	return o
}

func (o ECRAuthorizationTokenArrayOutput) ToECRAuthorizationTokenArrayOutputWithContext(ctx context.Context) ECRAuthorizationTokenArrayOutput {
	return o
}

func (o ECRAuthorizationTokenArrayOutput) Index(i pulumi.IntInput) ECRAuthorizationTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ECRAuthorizationToken {
		return vs[0].([]*ECRAuthorizationToken)[vs[1].(int)]
	}).(ECRAuthorizationTokenOutput)
}

type ECRAuthorizationTokenMapOutput struct{ *pulumi.OutputState }

func (ECRAuthorizationTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ECRAuthorizationToken)(nil)).Elem()
}

func (o ECRAuthorizationTokenMapOutput) ToECRAuthorizationTokenMapOutput() ECRAuthorizationTokenMapOutput {
	return o
}

func (o ECRAuthorizationTokenMapOutput) ToECRAuthorizationTokenMapOutputWithContext(ctx context.Context) ECRAuthorizationTokenMapOutput {
	return o
}

func (o ECRAuthorizationTokenMapOutput) MapIndex(k pulumi.StringInput) ECRAuthorizationTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ECRAuthorizationToken {
		return vs[0].(map[string]*ECRAuthorizationToken)[vs[1].(string)]
	}).(ECRAuthorizationTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ECRAuthorizationTokenInput)(nil)).Elem(), &ECRAuthorizationToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ECRAuthorizationTokenArrayInput)(nil)).Elem(), ECRAuthorizationTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ECRAuthorizationTokenMapInput)(nil)).Elem(), ECRAuthorizationTokenMap{})
	pulumi.RegisterOutputType(ECRAuthorizationTokenOutput{})
	pulumi.RegisterOutputType(ECRAuthorizationTokenArrayOutput{})
	pulumi.RegisterOutputType(ECRAuthorizationTokenMapOutput{})
}
