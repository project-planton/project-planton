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

type RequestAuthentication struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput         `pulumi:"metadata"`
	Spec     RequestAuthenticationSpecOutput `pulumi:"spec"`
	Status   pulumi.MapOutput                `pulumi:"status"`
}

// NewRequestAuthentication registers a new resource with the given unique name, arguments, and options.
func NewRequestAuthentication(ctx *pulumi.Context,
	name string, args *RequestAuthenticationArgs, opts ...pulumi.ResourceOption) (*RequestAuthentication, error) {
	if args == nil {
		args = &RequestAuthenticationArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("security.istio.io/v1beta1")
	args.Kind = pulumi.StringPtr("RequestAuthentication")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:security.istio.io/v1:RequestAuthentication"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RequestAuthentication
	err := ctx.RegisterResource("kubernetes:security.istio.io/v1beta1:RequestAuthentication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRequestAuthentication gets an existing RequestAuthentication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequestAuthentication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RequestAuthenticationState, opts ...pulumi.ResourceOption) (*RequestAuthentication, error) {
	var resource RequestAuthentication
	err := ctx.ReadResource("kubernetes:security.istio.io/v1beta1:RequestAuthentication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RequestAuthentication resources.
type requestAuthenticationState struct {
}

type RequestAuthenticationState struct {
}

func (RequestAuthenticationState) ElementType() reflect.Type {
	return reflect.TypeOf((*requestAuthenticationState)(nil)).Elem()
}

type requestAuthenticationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta         `pulumi:"metadata"`
	Spec     *RequestAuthenticationSpec `pulumi:"spec"`
}

// The set of arguments for constructing a RequestAuthentication resource.
type RequestAuthenticationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     RequestAuthenticationSpecPtrInput
}

func (RequestAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*requestAuthenticationArgs)(nil)).Elem()
}

type RequestAuthenticationInput interface {
	pulumi.Input

	ToRequestAuthenticationOutput() RequestAuthenticationOutput
	ToRequestAuthenticationOutputWithContext(ctx context.Context) RequestAuthenticationOutput
}

func (*RequestAuthentication) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestAuthentication)(nil)).Elem()
}

func (i *RequestAuthentication) ToRequestAuthenticationOutput() RequestAuthenticationOutput {
	return i.ToRequestAuthenticationOutputWithContext(context.Background())
}

func (i *RequestAuthentication) ToRequestAuthenticationOutputWithContext(ctx context.Context) RequestAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestAuthenticationOutput)
}

// RequestAuthenticationArrayInput is an input type that accepts RequestAuthenticationArray and RequestAuthenticationArrayOutput values.
// You can construct a concrete instance of `RequestAuthenticationArrayInput` via:
//
//	RequestAuthenticationArray{ RequestAuthenticationArgs{...} }
type RequestAuthenticationArrayInput interface {
	pulumi.Input

	ToRequestAuthenticationArrayOutput() RequestAuthenticationArrayOutput
	ToRequestAuthenticationArrayOutputWithContext(context.Context) RequestAuthenticationArrayOutput
}

type RequestAuthenticationArray []RequestAuthenticationInput

func (RequestAuthenticationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RequestAuthentication)(nil)).Elem()
}

func (i RequestAuthenticationArray) ToRequestAuthenticationArrayOutput() RequestAuthenticationArrayOutput {
	return i.ToRequestAuthenticationArrayOutputWithContext(context.Background())
}

func (i RequestAuthenticationArray) ToRequestAuthenticationArrayOutputWithContext(ctx context.Context) RequestAuthenticationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestAuthenticationArrayOutput)
}

// RequestAuthenticationMapInput is an input type that accepts RequestAuthenticationMap and RequestAuthenticationMapOutput values.
// You can construct a concrete instance of `RequestAuthenticationMapInput` via:
//
//	RequestAuthenticationMap{ "key": RequestAuthenticationArgs{...} }
type RequestAuthenticationMapInput interface {
	pulumi.Input

	ToRequestAuthenticationMapOutput() RequestAuthenticationMapOutput
	ToRequestAuthenticationMapOutputWithContext(context.Context) RequestAuthenticationMapOutput
}

type RequestAuthenticationMap map[string]RequestAuthenticationInput

func (RequestAuthenticationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RequestAuthentication)(nil)).Elem()
}

func (i RequestAuthenticationMap) ToRequestAuthenticationMapOutput() RequestAuthenticationMapOutput {
	return i.ToRequestAuthenticationMapOutputWithContext(context.Background())
}

func (i RequestAuthenticationMap) ToRequestAuthenticationMapOutputWithContext(ctx context.Context) RequestAuthenticationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestAuthenticationMapOutput)
}

type RequestAuthenticationOutput struct{ *pulumi.OutputState }

func (RequestAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestAuthentication)(nil)).Elem()
}

func (o RequestAuthenticationOutput) ToRequestAuthenticationOutput() RequestAuthenticationOutput {
	return o
}

func (o RequestAuthenticationOutput) ToRequestAuthenticationOutputWithContext(ctx context.Context) RequestAuthenticationOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o RequestAuthenticationOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RequestAuthentication) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o RequestAuthenticationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *RequestAuthentication) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o RequestAuthenticationOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *RequestAuthentication) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o RequestAuthenticationOutput) Spec() RequestAuthenticationSpecOutput {
	return o.ApplyT(func(v *RequestAuthentication) RequestAuthenticationSpecOutput { return v.Spec }).(RequestAuthenticationSpecOutput)
}

func (o RequestAuthenticationOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *RequestAuthentication) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

type RequestAuthenticationArrayOutput struct{ *pulumi.OutputState }

func (RequestAuthenticationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RequestAuthentication)(nil)).Elem()
}

func (o RequestAuthenticationArrayOutput) ToRequestAuthenticationArrayOutput() RequestAuthenticationArrayOutput {
	return o
}

func (o RequestAuthenticationArrayOutput) ToRequestAuthenticationArrayOutputWithContext(ctx context.Context) RequestAuthenticationArrayOutput {
	return o
}

func (o RequestAuthenticationArrayOutput) Index(i pulumi.IntInput) RequestAuthenticationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RequestAuthentication {
		return vs[0].([]*RequestAuthentication)[vs[1].(int)]
	}).(RequestAuthenticationOutput)
}

type RequestAuthenticationMapOutput struct{ *pulumi.OutputState }

func (RequestAuthenticationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RequestAuthentication)(nil)).Elem()
}

func (o RequestAuthenticationMapOutput) ToRequestAuthenticationMapOutput() RequestAuthenticationMapOutput {
	return o
}

func (o RequestAuthenticationMapOutput) ToRequestAuthenticationMapOutputWithContext(ctx context.Context) RequestAuthenticationMapOutput {
	return o
}

func (o RequestAuthenticationMapOutput) MapIndex(k pulumi.StringInput) RequestAuthenticationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RequestAuthentication {
		return vs[0].(map[string]*RequestAuthentication)[vs[1].(string)]
	}).(RequestAuthenticationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RequestAuthenticationInput)(nil)).Elem(), &RequestAuthentication{})
	pulumi.RegisterInputType(reflect.TypeOf((*RequestAuthenticationArrayInput)(nil)).Elem(), RequestAuthenticationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RequestAuthenticationMapInput)(nil)).Elem(), RequestAuthenticationMap{})
	pulumi.RegisterOutputType(RequestAuthenticationOutput{})
	pulumi.RegisterOutputType(RequestAuthenticationArrayOutput{})
	pulumi.RegisterOutputType(RequestAuthenticationMapOutput{})
}