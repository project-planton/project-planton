// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ProxyConfigList is a list of ProxyConfig
type ProxyConfigList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of proxyconfigs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ProxyConfigTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewProxyConfigList registers a new resource with the given unique name, arguments, and options.
func NewProxyConfigList(ctx *pulumi.Context,
	name string, args *ProxyConfigListArgs, opts ...pulumi.ResourceOption) (*ProxyConfigList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1beta1")
	args.Kind = pulumi.StringPtr("ProxyConfigList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProxyConfigList
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1beta1:ProxyConfigList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProxyConfigList gets an existing ProxyConfigList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProxyConfigList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProxyConfigListState, opts ...pulumi.ResourceOption) (*ProxyConfigList, error) {
	var resource ProxyConfigList
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1beta1:ProxyConfigList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProxyConfigList resources.
type proxyConfigListState struct {
}

type ProxyConfigListState struct {
}

func (ProxyConfigListState) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyConfigListState)(nil)).Elem()
}

type proxyConfigListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of proxyconfigs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ProxyConfigType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ProxyConfigList resource.
type ProxyConfigListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of proxyconfigs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ProxyConfigTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ProxyConfigListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyConfigListArgs)(nil)).Elem()
}

type ProxyConfigListInput interface {
	pulumi.Input

	ToProxyConfigListOutput() ProxyConfigListOutput
	ToProxyConfigListOutputWithContext(ctx context.Context) ProxyConfigListOutput
}

func (*ProxyConfigList) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyConfigList)(nil)).Elem()
}

func (i *ProxyConfigList) ToProxyConfigListOutput() ProxyConfigListOutput {
	return i.ToProxyConfigListOutputWithContext(context.Background())
}

func (i *ProxyConfigList) ToProxyConfigListOutputWithContext(ctx context.Context) ProxyConfigListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyConfigListOutput)
}

// ProxyConfigListArrayInput is an input type that accepts ProxyConfigListArray and ProxyConfigListArrayOutput values.
// You can construct a concrete instance of `ProxyConfigListArrayInput` via:
//
//	ProxyConfigListArray{ ProxyConfigListArgs{...} }
type ProxyConfigListArrayInput interface {
	pulumi.Input

	ToProxyConfigListArrayOutput() ProxyConfigListArrayOutput
	ToProxyConfigListArrayOutputWithContext(context.Context) ProxyConfigListArrayOutput
}

type ProxyConfigListArray []ProxyConfigListInput

func (ProxyConfigListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProxyConfigList)(nil)).Elem()
}

func (i ProxyConfigListArray) ToProxyConfigListArrayOutput() ProxyConfigListArrayOutput {
	return i.ToProxyConfigListArrayOutputWithContext(context.Background())
}

func (i ProxyConfigListArray) ToProxyConfigListArrayOutputWithContext(ctx context.Context) ProxyConfigListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyConfigListArrayOutput)
}

// ProxyConfigListMapInput is an input type that accepts ProxyConfigListMap and ProxyConfigListMapOutput values.
// You can construct a concrete instance of `ProxyConfigListMapInput` via:
//
//	ProxyConfigListMap{ "key": ProxyConfigListArgs{...} }
type ProxyConfigListMapInput interface {
	pulumi.Input

	ToProxyConfigListMapOutput() ProxyConfigListMapOutput
	ToProxyConfigListMapOutputWithContext(context.Context) ProxyConfigListMapOutput
}

type ProxyConfigListMap map[string]ProxyConfigListInput

func (ProxyConfigListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProxyConfigList)(nil)).Elem()
}

func (i ProxyConfigListMap) ToProxyConfigListMapOutput() ProxyConfigListMapOutput {
	return i.ToProxyConfigListMapOutputWithContext(context.Background())
}

func (i ProxyConfigListMap) ToProxyConfigListMapOutputWithContext(ctx context.Context) ProxyConfigListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyConfigListMapOutput)
}

type ProxyConfigListOutput struct{ *pulumi.OutputState }

func (ProxyConfigListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyConfigList)(nil)).Elem()
}

func (o ProxyConfigListOutput) ToProxyConfigListOutput() ProxyConfigListOutput {
	return o
}

func (o ProxyConfigListOutput) ToProxyConfigListOutputWithContext(ctx context.Context) ProxyConfigListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ProxyConfigListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ProxyConfigList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of proxyconfigs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ProxyConfigListOutput) Items() ProxyConfigTypeArrayOutput {
	return o.ApplyT(func(v *ProxyConfigList) ProxyConfigTypeArrayOutput { return v.Items }).(ProxyConfigTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ProxyConfigListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ProxyConfigList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ProxyConfigListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ProxyConfigList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ProxyConfigListArrayOutput struct{ *pulumi.OutputState }

func (ProxyConfigListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProxyConfigList)(nil)).Elem()
}

func (o ProxyConfigListArrayOutput) ToProxyConfigListArrayOutput() ProxyConfigListArrayOutput {
	return o
}

func (o ProxyConfigListArrayOutput) ToProxyConfigListArrayOutputWithContext(ctx context.Context) ProxyConfigListArrayOutput {
	return o
}

func (o ProxyConfigListArrayOutput) Index(i pulumi.IntInput) ProxyConfigListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProxyConfigList {
		return vs[0].([]*ProxyConfigList)[vs[1].(int)]
	}).(ProxyConfigListOutput)
}

type ProxyConfigListMapOutput struct{ *pulumi.OutputState }

func (ProxyConfigListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProxyConfigList)(nil)).Elem()
}

func (o ProxyConfigListMapOutput) ToProxyConfigListMapOutput() ProxyConfigListMapOutput {
	return o
}

func (o ProxyConfigListMapOutput) ToProxyConfigListMapOutputWithContext(ctx context.Context) ProxyConfigListMapOutput {
	return o
}

func (o ProxyConfigListMapOutput) MapIndex(k pulumi.StringInput) ProxyConfigListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProxyConfigList {
		return vs[0].(map[string]*ProxyConfigList)[vs[1].(string)]
	}).(ProxyConfigListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyConfigListInput)(nil)).Elem(), &ProxyConfigList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyConfigListArrayInput)(nil)).Elem(), ProxyConfigListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyConfigListMapInput)(nil)).Elem(), ProxyConfigListMap{})
	pulumi.RegisterOutputType(ProxyConfigListOutput{})
	pulumi.RegisterOutputType(ProxyConfigListArrayOutput{})
	pulumi.RegisterOutputType(ProxyConfigListMapOutput{})
}