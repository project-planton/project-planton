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

// EnterpriseSearch is a Kubernetes CRD to represent Enterprise Search.
type EnterpriseSearch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput         `pulumi:"metadata"`
	Spec     EnterpriseSearchSpecOutput      `pulumi:"spec"`
	Status   EnterpriseSearchStatusPtrOutput `pulumi:"status"`
}

// NewEnterpriseSearch registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseSearch(ctx *pulumi.Context,
	name string, args *EnterpriseSearchArgs, opts ...pulumi.ResourceOption) (*EnterpriseSearch, error) {
	if args == nil {
		args = &EnterpriseSearchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("enterprisesearch.k8s.elastic.co/v1beta1")
	args.Kind = pulumi.StringPtr("EnterpriseSearch")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:enterprisesearch.k8s.elastic.co/v1:EnterpriseSearch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnterpriseSearch
	err := ctx.RegisterResource("kubernetes:enterprisesearch.k8s.elastic.co/v1beta1:EnterpriseSearch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseSearch gets an existing EnterpriseSearch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseSearch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseSearchState, opts ...pulumi.ResourceOption) (*EnterpriseSearch, error) {
	var resource EnterpriseSearch
	err := ctx.ReadResource("kubernetes:enterprisesearch.k8s.elastic.co/v1beta1:EnterpriseSearch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseSearch resources.
type enterpriseSearchState struct {
}

type EnterpriseSearchState struct {
}

func (EnterpriseSearchState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseSearchState)(nil)).Elem()
}

type enterpriseSearchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta    `pulumi:"metadata"`
	Spec     *EnterpriseSearchSpec `pulumi:"spec"`
}

// The set of arguments for constructing a EnterpriseSearch resource.
type EnterpriseSearchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     EnterpriseSearchSpecPtrInput
}

func (EnterpriseSearchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseSearchArgs)(nil)).Elem()
}

type EnterpriseSearchInput interface {
	pulumi.Input

	ToEnterpriseSearchOutput() EnterpriseSearchOutput
	ToEnterpriseSearchOutputWithContext(ctx context.Context) EnterpriseSearchOutput
}

func (*EnterpriseSearch) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseSearch)(nil)).Elem()
}

func (i *EnterpriseSearch) ToEnterpriseSearchOutput() EnterpriseSearchOutput {
	return i.ToEnterpriseSearchOutputWithContext(context.Background())
}

func (i *EnterpriseSearch) ToEnterpriseSearchOutputWithContext(ctx context.Context) EnterpriseSearchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSearchOutput)
}

// EnterpriseSearchArrayInput is an input type that accepts EnterpriseSearchArray and EnterpriseSearchArrayOutput values.
// You can construct a concrete instance of `EnterpriseSearchArrayInput` via:
//
//	EnterpriseSearchArray{ EnterpriseSearchArgs{...} }
type EnterpriseSearchArrayInput interface {
	pulumi.Input

	ToEnterpriseSearchArrayOutput() EnterpriseSearchArrayOutput
	ToEnterpriseSearchArrayOutputWithContext(context.Context) EnterpriseSearchArrayOutput
}

type EnterpriseSearchArray []EnterpriseSearchInput

func (EnterpriseSearchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseSearch)(nil)).Elem()
}

func (i EnterpriseSearchArray) ToEnterpriseSearchArrayOutput() EnterpriseSearchArrayOutput {
	return i.ToEnterpriseSearchArrayOutputWithContext(context.Background())
}

func (i EnterpriseSearchArray) ToEnterpriseSearchArrayOutputWithContext(ctx context.Context) EnterpriseSearchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSearchArrayOutput)
}

// EnterpriseSearchMapInput is an input type that accepts EnterpriseSearchMap and EnterpriseSearchMapOutput values.
// You can construct a concrete instance of `EnterpriseSearchMapInput` via:
//
//	EnterpriseSearchMap{ "key": EnterpriseSearchArgs{...} }
type EnterpriseSearchMapInput interface {
	pulumi.Input

	ToEnterpriseSearchMapOutput() EnterpriseSearchMapOutput
	ToEnterpriseSearchMapOutputWithContext(context.Context) EnterpriseSearchMapOutput
}

type EnterpriseSearchMap map[string]EnterpriseSearchInput

func (EnterpriseSearchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseSearch)(nil)).Elem()
}

func (i EnterpriseSearchMap) ToEnterpriseSearchMapOutput() EnterpriseSearchMapOutput {
	return i.ToEnterpriseSearchMapOutputWithContext(context.Background())
}

func (i EnterpriseSearchMap) ToEnterpriseSearchMapOutputWithContext(ctx context.Context) EnterpriseSearchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSearchMapOutput)
}

type EnterpriseSearchOutput struct{ *pulumi.OutputState }

func (EnterpriseSearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseSearch)(nil)).Elem()
}

func (o EnterpriseSearchOutput) ToEnterpriseSearchOutput() EnterpriseSearchOutput {
	return o
}

func (o EnterpriseSearchOutput) ToEnterpriseSearchOutputWithContext(ctx context.Context) EnterpriseSearchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EnterpriseSearchOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseSearch) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EnterpriseSearchOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseSearch) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o EnterpriseSearchOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *EnterpriseSearch) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o EnterpriseSearchOutput) Spec() EnterpriseSearchSpecOutput {
	return o.ApplyT(func(v *EnterpriseSearch) EnterpriseSearchSpecOutput { return v.Spec }).(EnterpriseSearchSpecOutput)
}

func (o EnterpriseSearchOutput) Status() EnterpriseSearchStatusPtrOutput {
	return o.ApplyT(func(v *EnterpriseSearch) EnterpriseSearchStatusPtrOutput { return v.Status }).(EnterpriseSearchStatusPtrOutput)
}

type EnterpriseSearchArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseSearchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseSearch)(nil)).Elem()
}

func (o EnterpriseSearchArrayOutput) ToEnterpriseSearchArrayOutput() EnterpriseSearchArrayOutput {
	return o
}

func (o EnterpriseSearchArrayOutput) ToEnterpriseSearchArrayOutputWithContext(ctx context.Context) EnterpriseSearchArrayOutput {
	return o
}

func (o EnterpriseSearchArrayOutput) Index(i pulumi.IntInput) EnterpriseSearchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseSearch {
		return vs[0].([]*EnterpriseSearch)[vs[1].(int)]
	}).(EnterpriseSearchOutput)
}

type EnterpriseSearchMapOutput struct{ *pulumi.OutputState }

func (EnterpriseSearchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseSearch)(nil)).Elem()
}

func (o EnterpriseSearchMapOutput) ToEnterpriseSearchMapOutput() EnterpriseSearchMapOutput {
	return o
}

func (o EnterpriseSearchMapOutput) ToEnterpriseSearchMapOutputWithContext(ctx context.Context) EnterpriseSearchMapOutput {
	return o
}

func (o EnterpriseSearchMapOutput) MapIndex(k pulumi.StringInput) EnterpriseSearchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseSearch {
		return vs[0].(map[string]*EnterpriseSearch)[vs[1].(string)]
	}).(EnterpriseSearchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSearchInput)(nil)).Elem(), &EnterpriseSearch{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSearchArrayInput)(nil)).Elem(), EnterpriseSearchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSearchMapInput)(nil)).Elem(), EnterpriseSearchMap{})
	pulumi.RegisterOutputType(EnterpriseSearchOutput{})
	pulumi.RegisterOutputType(EnterpriseSearchArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseSearchMapOutput{})
}