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

// ApmServer represents an APM Server resource in a Kubernetes cluster.
type ApmServer struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput  `pulumi:"metadata"`
	Spec     ApmServerSpecOutput      `pulumi:"spec"`
	Status   ApmServerStatusPtrOutput `pulumi:"status"`
}

// NewApmServer registers a new resource with the given unique name, arguments, and options.
func NewApmServer(ctx *pulumi.Context,
	name string, args *ApmServerArgs, opts ...pulumi.ResourceOption) (*ApmServer, error) {
	if args == nil {
		args = &ApmServerArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("apm.k8s.elastic.co/v1")
	args.Kind = pulumi.StringPtr("ApmServer")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apm.k8s.elastic.co/v1alpha1:ApmServer"),
		},
		{
			Type: pulumi.String("kubernetes:apm.k8s.elastic.co/v1beta1:ApmServer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApmServer
	err := ctx.RegisterResource("kubernetes:apm.k8s.elastic.co/v1:ApmServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApmServer gets an existing ApmServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApmServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApmServerState, opts ...pulumi.ResourceOption) (*ApmServer, error) {
	var resource ApmServer
	err := ctx.ReadResource("kubernetes:apm.k8s.elastic.co/v1:ApmServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApmServer resources.
type apmServerState struct {
}

type ApmServerState struct {
}

func (ApmServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*apmServerState)(nil)).Elem()
}

type apmServerArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *ApmServerSpec     `pulumi:"spec"`
}

// The set of arguments for constructing a ApmServer resource.
type ApmServerArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     ApmServerSpecPtrInput
}

func (ApmServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apmServerArgs)(nil)).Elem()
}

type ApmServerInput interface {
	pulumi.Input

	ToApmServerOutput() ApmServerOutput
	ToApmServerOutputWithContext(ctx context.Context) ApmServerOutput
}

func (*ApmServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ApmServer)(nil)).Elem()
}

func (i *ApmServer) ToApmServerOutput() ApmServerOutput {
	return i.ToApmServerOutputWithContext(context.Background())
}

func (i *ApmServer) ToApmServerOutputWithContext(ctx context.Context) ApmServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmServerOutput)
}

// ApmServerArrayInput is an input type that accepts ApmServerArray and ApmServerArrayOutput values.
// You can construct a concrete instance of `ApmServerArrayInput` via:
//
//	ApmServerArray{ ApmServerArgs{...} }
type ApmServerArrayInput interface {
	pulumi.Input

	ToApmServerArrayOutput() ApmServerArrayOutput
	ToApmServerArrayOutputWithContext(context.Context) ApmServerArrayOutput
}

type ApmServerArray []ApmServerInput

func (ApmServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApmServer)(nil)).Elem()
}

func (i ApmServerArray) ToApmServerArrayOutput() ApmServerArrayOutput {
	return i.ToApmServerArrayOutputWithContext(context.Background())
}

func (i ApmServerArray) ToApmServerArrayOutputWithContext(ctx context.Context) ApmServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmServerArrayOutput)
}

// ApmServerMapInput is an input type that accepts ApmServerMap and ApmServerMapOutput values.
// You can construct a concrete instance of `ApmServerMapInput` via:
//
//	ApmServerMap{ "key": ApmServerArgs{...} }
type ApmServerMapInput interface {
	pulumi.Input

	ToApmServerMapOutput() ApmServerMapOutput
	ToApmServerMapOutputWithContext(context.Context) ApmServerMapOutput
}

type ApmServerMap map[string]ApmServerInput

func (ApmServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApmServer)(nil)).Elem()
}

func (i ApmServerMap) ToApmServerMapOutput() ApmServerMapOutput {
	return i.ToApmServerMapOutputWithContext(context.Background())
}

func (i ApmServerMap) ToApmServerMapOutputWithContext(ctx context.Context) ApmServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmServerMapOutput)
}

type ApmServerOutput struct{ *pulumi.OutputState }

func (ApmServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApmServer)(nil)).Elem()
}

func (o ApmServerOutput) ToApmServerOutput() ApmServerOutput {
	return o
}

func (o ApmServerOutput) ToApmServerOutputWithContext(ctx context.Context) ApmServerOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ApmServerOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ApmServer) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ApmServerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ApmServer) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ApmServerOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ApmServer) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o ApmServerOutput) Spec() ApmServerSpecOutput {
	return o.ApplyT(func(v *ApmServer) ApmServerSpecOutput { return v.Spec }).(ApmServerSpecOutput)
}

func (o ApmServerOutput) Status() ApmServerStatusPtrOutput {
	return o.ApplyT(func(v *ApmServer) ApmServerStatusPtrOutput { return v.Status }).(ApmServerStatusPtrOutput)
}

type ApmServerArrayOutput struct{ *pulumi.OutputState }

func (ApmServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApmServer)(nil)).Elem()
}

func (o ApmServerArrayOutput) ToApmServerArrayOutput() ApmServerArrayOutput {
	return o
}

func (o ApmServerArrayOutput) ToApmServerArrayOutputWithContext(ctx context.Context) ApmServerArrayOutput {
	return o
}

func (o ApmServerArrayOutput) Index(i pulumi.IntInput) ApmServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApmServer {
		return vs[0].([]*ApmServer)[vs[1].(int)]
	}).(ApmServerOutput)
}

type ApmServerMapOutput struct{ *pulumi.OutputState }

func (ApmServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApmServer)(nil)).Elem()
}

func (o ApmServerMapOutput) ToApmServerMapOutput() ApmServerMapOutput {
	return o
}

func (o ApmServerMapOutput) ToApmServerMapOutputWithContext(ctx context.Context) ApmServerMapOutput {
	return o
}

func (o ApmServerMapOutput) MapIndex(k pulumi.StringInput) ApmServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApmServer {
		return vs[0].(map[string]*ApmServer)[vs[1].(string)]
	}).(ApmServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApmServerInput)(nil)).Elem(), &ApmServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApmServerArrayInput)(nil)).Elem(), ApmServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApmServerMapInput)(nil)).Elem(), ApmServerMap{})
	pulumi.RegisterOutputType(ApmServerOutput{})
	pulumi.RegisterOutputType(ApmServerArrayOutput{})
	pulumi.RegisterOutputType(ApmServerMapOutput{})
}
