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

// ExternalSecretList is a list of ExternalSecret
type ExternalSecretList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of externalsecrets. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ExternalSecretTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewExternalSecretList registers a new resource with the given unique name, arguments, and options.
func NewExternalSecretList(ctx *pulumi.Context,
	name string, args *ExternalSecretListArgs, opts ...pulumi.ResourceOption) (*ExternalSecretList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("external-secrets.io/v1beta1")
	args.Kind = pulumi.StringPtr("ExternalSecretList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ExternalSecretList
	err := ctx.RegisterResource("kubernetes:external-secrets.io/v1beta1:ExternalSecretList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalSecretList gets an existing ExternalSecretList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalSecretList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalSecretListState, opts ...pulumi.ResourceOption) (*ExternalSecretList, error) {
	var resource ExternalSecretList
	err := ctx.ReadResource("kubernetes:external-secrets.io/v1beta1:ExternalSecretList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalSecretList resources.
type externalSecretListState struct {
}

type ExternalSecretListState struct {
}

func (ExternalSecretListState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalSecretListState)(nil)).Elem()
}

type externalSecretListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of externalsecrets. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ExternalSecretType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ExternalSecretList resource.
type ExternalSecretListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of externalsecrets. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ExternalSecretTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ExternalSecretListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalSecretListArgs)(nil)).Elem()
}

type ExternalSecretListInput interface {
	pulumi.Input

	ToExternalSecretListOutput() ExternalSecretListOutput
	ToExternalSecretListOutputWithContext(ctx context.Context) ExternalSecretListOutput
}

func (*ExternalSecretList) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalSecretList)(nil)).Elem()
}

func (i *ExternalSecretList) ToExternalSecretListOutput() ExternalSecretListOutput {
	return i.ToExternalSecretListOutputWithContext(context.Background())
}

func (i *ExternalSecretList) ToExternalSecretListOutputWithContext(ctx context.Context) ExternalSecretListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalSecretListOutput)
}

// ExternalSecretListArrayInput is an input type that accepts ExternalSecretListArray and ExternalSecretListArrayOutput values.
// You can construct a concrete instance of `ExternalSecretListArrayInput` via:
//
//	ExternalSecretListArray{ ExternalSecretListArgs{...} }
type ExternalSecretListArrayInput interface {
	pulumi.Input

	ToExternalSecretListArrayOutput() ExternalSecretListArrayOutput
	ToExternalSecretListArrayOutputWithContext(context.Context) ExternalSecretListArrayOutput
}

type ExternalSecretListArray []ExternalSecretListInput

func (ExternalSecretListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalSecretList)(nil)).Elem()
}

func (i ExternalSecretListArray) ToExternalSecretListArrayOutput() ExternalSecretListArrayOutput {
	return i.ToExternalSecretListArrayOutputWithContext(context.Background())
}

func (i ExternalSecretListArray) ToExternalSecretListArrayOutputWithContext(ctx context.Context) ExternalSecretListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalSecretListArrayOutput)
}

// ExternalSecretListMapInput is an input type that accepts ExternalSecretListMap and ExternalSecretListMapOutput values.
// You can construct a concrete instance of `ExternalSecretListMapInput` via:
//
//	ExternalSecretListMap{ "key": ExternalSecretListArgs{...} }
type ExternalSecretListMapInput interface {
	pulumi.Input

	ToExternalSecretListMapOutput() ExternalSecretListMapOutput
	ToExternalSecretListMapOutputWithContext(context.Context) ExternalSecretListMapOutput
}

type ExternalSecretListMap map[string]ExternalSecretListInput

func (ExternalSecretListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalSecretList)(nil)).Elem()
}

func (i ExternalSecretListMap) ToExternalSecretListMapOutput() ExternalSecretListMapOutput {
	return i.ToExternalSecretListMapOutputWithContext(context.Background())
}

func (i ExternalSecretListMap) ToExternalSecretListMapOutputWithContext(ctx context.Context) ExternalSecretListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalSecretListMapOutput)
}

type ExternalSecretListOutput struct{ *pulumi.OutputState }

func (ExternalSecretListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalSecretList)(nil)).Elem()
}

func (o ExternalSecretListOutput) ToExternalSecretListOutput() ExternalSecretListOutput {
	return o
}

func (o ExternalSecretListOutput) ToExternalSecretListOutputWithContext(ctx context.Context) ExternalSecretListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ExternalSecretListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalSecretList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of externalsecrets. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ExternalSecretListOutput) Items() ExternalSecretTypeArrayOutput {
	return o.ApplyT(func(v *ExternalSecretList) ExternalSecretTypeArrayOutput { return v.Items }).(ExternalSecretTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ExternalSecretListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalSecretList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ExternalSecretListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ExternalSecretList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ExternalSecretListArrayOutput struct{ *pulumi.OutputState }

func (ExternalSecretListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalSecretList)(nil)).Elem()
}

func (o ExternalSecretListArrayOutput) ToExternalSecretListArrayOutput() ExternalSecretListArrayOutput {
	return o
}

func (o ExternalSecretListArrayOutput) ToExternalSecretListArrayOutputWithContext(ctx context.Context) ExternalSecretListArrayOutput {
	return o
}

func (o ExternalSecretListArrayOutput) Index(i pulumi.IntInput) ExternalSecretListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalSecretList {
		return vs[0].([]*ExternalSecretList)[vs[1].(int)]
	}).(ExternalSecretListOutput)
}

type ExternalSecretListMapOutput struct{ *pulumi.OutputState }

func (ExternalSecretListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalSecretList)(nil)).Elem()
}

func (o ExternalSecretListMapOutput) ToExternalSecretListMapOutput() ExternalSecretListMapOutput {
	return o
}

func (o ExternalSecretListMapOutput) ToExternalSecretListMapOutputWithContext(ctx context.Context) ExternalSecretListMapOutput {
	return o
}

func (o ExternalSecretListMapOutput) MapIndex(k pulumi.StringInput) ExternalSecretListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalSecretList {
		return vs[0].(map[string]*ExternalSecretList)[vs[1].(string)]
	}).(ExternalSecretListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalSecretListInput)(nil)).Elem(), &ExternalSecretList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalSecretListArrayInput)(nil)).Elem(), ExternalSecretListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalSecretListMapInput)(nil)).Elem(), ExternalSecretListMap{})
	pulumi.RegisterOutputType(ExternalSecretListOutput{})
	pulumi.RegisterOutputType(ExternalSecretListArrayOutput{})
	pulumi.RegisterOutputType(ExternalSecretListMapOutput{})
}
