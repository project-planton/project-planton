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

// ClusterExternalSecretList is a list of ClusterExternalSecret
type ClusterExternalSecretList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of clusterexternalsecrets. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ClusterExternalSecretTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewClusterExternalSecretList registers a new resource with the given unique name, arguments, and options.
func NewClusterExternalSecretList(ctx *pulumi.Context,
	name string, args *ClusterExternalSecretListArgs, opts ...pulumi.ResourceOption) (*ClusterExternalSecretList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("external-secrets.io/v1beta1")
	args.Kind = pulumi.StringPtr("ClusterExternalSecretList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ClusterExternalSecretList
	err := ctx.RegisterResource("kubernetes:external-secrets.io/v1beta1:ClusterExternalSecretList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterExternalSecretList gets an existing ClusterExternalSecretList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterExternalSecretList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterExternalSecretListState, opts ...pulumi.ResourceOption) (*ClusterExternalSecretList, error) {
	var resource ClusterExternalSecretList
	err := ctx.ReadResource("kubernetes:external-secrets.io/v1beta1:ClusterExternalSecretList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterExternalSecretList resources.
type clusterExternalSecretListState struct {
}

type ClusterExternalSecretListState struct {
}

func (ClusterExternalSecretListState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterExternalSecretListState)(nil)).Elem()
}

type clusterExternalSecretListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of clusterexternalsecrets. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ClusterExternalSecretType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ClusterExternalSecretList resource.
type ClusterExternalSecretListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of clusterexternalsecrets. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ClusterExternalSecretTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ClusterExternalSecretListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterExternalSecretListArgs)(nil)).Elem()
}

type ClusterExternalSecretListInput interface {
	pulumi.Input

	ToClusterExternalSecretListOutput() ClusterExternalSecretListOutput
	ToClusterExternalSecretListOutputWithContext(ctx context.Context) ClusterExternalSecretListOutput
}

func (*ClusterExternalSecretList) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterExternalSecretList)(nil)).Elem()
}

func (i *ClusterExternalSecretList) ToClusterExternalSecretListOutput() ClusterExternalSecretListOutput {
	return i.ToClusterExternalSecretListOutputWithContext(context.Background())
}

func (i *ClusterExternalSecretList) ToClusterExternalSecretListOutputWithContext(ctx context.Context) ClusterExternalSecretListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterExternalSecretListOutput)
}

// ClusterExternalSecretListArrayInput is an input type that accepts ClusterExternalSecretListArray and ClusterExternalSecretListArrayOutput values.
// You can construct a concrete instance of `ClusterExternalSecretListArrayInput` via:
//
//	ClusterExternalSecretListArray{ ClusterExternalSecretListArgs{...} }
type ClusterExternalSecretListArrayInput interface {
	pulumi.Input

	ToClusterExternalSecretListArrayOutput() ClusterExternalSecretListArrayOutput
	ToClusterExternalSecretListArrayOutputWithContext(context.Context) ClusterExternalSecretListArrayOutput
}

type ClusterExternalSecretListArray []ClusterExternalSecretListInput

func (ClusterExternalSecretListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterExternalSecretList)(nil)).Elem()
}

func (i ClusterExternalSecretListArray) ToClusterExternalSecretListArrayOutput() ClusterExternalSecretListArrayOutput {
	return i.ToClusterExternalSecretListArrayOutputWithContext(context.Background())
}

func (i ClusterExternalSecretListArray) ToClusterExternalSecretListArrayOutputWithContext(ctx context.Context) ClusterExternalSecretListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterExternalSecretListArrayOutput)
}

// ClusterExternalSecretListMapInput is an input type that accepts ClusterExternalSecretListMap and ClusterExternalSecretListMapOutput values.
// You can construct a concrete instance of `ClusterExternalSecretListMapInput` via:
//
//	ClusterExternalSecretListMap{ "key": ClusterExternalSecretListArgs{...} }
type ClusterExternalSecretListMapInput interface {
	pulumi.Input

	ToClusterExternalSecretListMapOutput() ClusterExternalSecretListMapOutput
	ToClusterExternalSecretListMapOutputWithContext(context.Context) ClusterExternalSecretListMapOutput
}

type ClusterExternalSecretListMap map[string]ClusterExternalSecretListInput

func (ClusterExternalSecretListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterExternalSecretList)(nil)).Elem()
}

func (i ClusterExternalSecretListMap) ToClusterExternalSecretListMapOutput() ClusterExternalSecretListMapOutput {
	return i.ToClusterExternalSecretListMapOutputWithContext(context.Background())
}

func (i ClusterExternalSecretListMap) ToClusterExternalSecretListMapOutputWithContext(ctx context.Context) ClusterExternalSecretListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterExternalSecretListMapOutput)
}

type ClusterExternalSecretListOutput struct{ *pulumi.OutputState }

func (ClusterExternalSecretListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterExternalSecretList)(nil)).Elem()
}

func (o ClusterExternalSecretListOutput) ToClusterExternalSecretListOutput() ClusterExternalSecretListOutput {
	return o
}

func (o ClusterExternalSecretListOutput) ToClusterExternalSecretListOutputWithContext(ctx context.Context) ClusterExternalSecretListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ClusterExternalSecretListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterExternalSecretList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of clusterexternalsecrets. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ClusterExternalSecretListOutput) Items() ClusterExternalSecretTypeArrayOutput {
	return o.ApplyT(func(v *ClusterExternalSecretList) ClusterExternalSecretTypeArrayOutput { return v.Items }).(ClusterExternalSecretTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterExternalSecretListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterExternalSecretList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterExternalSecretListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ClusterExternalSecretList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ClusterExternalSecretListArrayOutput struct{ *pulumi.OutputState }

func (ClusterExternalSecretListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterExternalSecretList)(nil)).Elem()
}

func (o ClusterExternalSecretListArrayOutput) ToClusterExternalSecretListArrayOutput() ClusterExternalSecretListArrayOutput {
	return o
}

func (o ClusterExternalSecretListArrayOutput) ToClusterExternalSecretListArrayOutputWithContext(ctx context.Context) ClusterExternalSecretListArrayOutput {
	return o
}

func (o ClusterExternalSecretListArrayOutput) Index(i pulumi.IntInput) ClusterExternalSecretListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterExternalSecretList {
		return vs[0].([]*ClusterExternalSecretList)[vs[1].(int)]
	}).(ClusterExternalSecretListOutput)
}

type ClusterExternalSecretListMapOutput struct{ *pulumi.OutputState }

func (ClusterExternalSecretListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterExternalSecretList)(nil)).Elem()
}

func (o ClusterExternalSecretListMapOutput) ToClusterExternalSecretListMapOutput() ClusterExternalSecretListMapOutput {
	return o
}

func (o ClusterExternalSecretListMapOutput) ToClusterExternalSecretListMapOutputWithContext(ctx context.Context) ClusterExternalSecretListMapOutput {
	return o
}

func (o ClusterExternalSecretListMapOutput) MapIndex(k pulumi.StringInput) ClusterExternalSecretListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterExternalSecretList {
		return vs[0].(map[string]*ClusterExternalSecretList)[vs[1].(string)]
	}).(ClusterExternalSecretListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterExternalSecretListInput)(nil)).Elem(), &ClusterExternalSecretList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterExternalSecretListArrayInput)(nil)).Elem(), ClusterExternalSecretListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterExternalSecretListMapInput)(nil)).Elem(), ClusterExternalSecretListMap{})
	pulumi.RegisterOutputType(ClusterExternalSecretListOutput{})
	pulumi.RegisterOutputType(ClusterExternalSecretListArrayOutput{})
	pulumi.RegisterOutputType(ClusterExternalSecretListMapOutput{})
}
