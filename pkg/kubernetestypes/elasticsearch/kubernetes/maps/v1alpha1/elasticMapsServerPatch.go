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

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// ElasticMapsServer represents an Elastic Map Server resource in a Kubernetes cluster.
type ElasticMapsServerPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput       `pulumi:"metadata"`
	Spec     ElasticMapsServerSpecPatchPtrOutput   `pulumi:"spec"`
	Status   ElasticMapsServerStatusPatchPtrOutput `pulumi:"status"`
}

// NewElasticMapsServerPatch registers a new resource with the given unique name, arguments, and options.
func NewElasticMapsServerPatch(ctx *pulumi.Context,
	name string, args *ElasticMapsServerPatchArgs, opts ...pulumi.ResourceOption) (*ElasticMapsServerPatch, error) {
	if args == nil {
		args = &ElasticMapsServerPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("maps.k8s.elastic.co/v1alpha1")
	args.Kind = pulumi.StringPtr("ElasticMapsServer")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ElasticMapsServerPatch
	err := ctx.RegisterResource("kubernetes:maps.k8s.elastic.co/v1alpha1:ElasticMapsServerPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticMapsServerPatch gets an existing ElasticMapsServerPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticMapsServerPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticMapsServerPatchState, opts ...pulumi.ResourceOption) (*ElasticMapsServerPatch, error) {
	var resource ElasticMapsServerPatch
	err := ctx.ReadResource("kubernetes:maps.k8s.elastic.co/v1alpha1:ElasticMapsServerPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticMapsServerPatch resources.
type elasticMapsServerPatchState struct {
}

type ElasticMapsServerPatchState struct {
}

func (ElasticMapsServerPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticMapsServerPatchState)(nil)).Elem()
}

type elasticMapsServerPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch     `pulumi:"metadata"`
	Spec     *ElasticMapsServerSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a ElasticMapsServerPatch resource.
type ElasticMapsServerPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     ElasticMapsServerSpecPatchPtrInput
}

func (ElasticMapsServerPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticMapsServerPatchArgs)(nil)).Elem()
}

type ElasticMapsServerPatchInput interface {
	pulumi.Input

	ToElasticMapsServerPatchOutput() ElasticMapsServerPatchOutput
	ToElasticMapsServerPatchOutputWithContext(ctx context.Context) ElasticMapsServerPatchOutput
}

func (*ElasticMapsServerPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticMapsServerPatch)(nil)).Elem()
}

func (i *ElasticMapsServerPatch) ToElasticMapsServerPatchOutput() ElasticMapsServerPatchOutput {
	return i.ToElasticMapsServerPatchOutputWithContext(context.Background())
}

func (i *ElasticMapsServerPatch) ToElasticMapsServerPatchOutputWithContext(ctx context.Context) ElasticMapsServerPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticMapsServerPatchOutput)
}

// ElasticMapsServerPatchArrayInput is an input type that accepts ElasticMapsServerPatchArray and ElasticMapsServerPatchArrayOutput values.
// You can construct a concrete instance of `ElasticMapsServerPatchArrayInput` via:
//
//	ElasticMapsServerPatchArray{ ElasticMapsServerPatchArgs{...} }
type ElasticMapsServerPatchArrayInput interface {
	pulumi.Input

	ToElasticMapsServerPatchArrayOutput() ElasticMapsServerPatchArrayOutput
	ToElasticMapsServerPatchArrayOutputWithContext(context.Context) ElasticMapsServerPatchArrayOutput
}

type ElasticMapsServerPatchArray []ElasticMapsServerPatchInput

func (ElasticMapsServerPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticMapsServerPatch)(nil)).Elem()
}

func (i ElasticMapsServerPatchArray) ToElasticMapsServerPatchArrayOutput() ElasticMapsServerPatchArrayOutput {
	return i.ToElasticMapsServerPatchArrayOutputWithContext(context.Background())
}

func (i ElasticMapsServerPatchArray) ToElasticMapsServerPatchArrayOutputWithContext(ctx context.Context) ElasticMapsServerPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticMapsServerPatchArrayOutput)
}

// ElasticMapsServerPatchMapInput is an input type that accepts ElasticMapsServerPatchMap and ElasticMapsServerPatchMapOutput values.
// You can construct a concrete instance of `ElasticMapsServerPatchMapInput` via:
//
//	ElasticMapsServerPatchMap{ "key": ElasticMapsServerPatchArgs{...} }
type ElasticMapsServerPatchMapInput interface {
	pulumi.Input

	ToElasticMapsServerPatchMapOutput() ElasticMapsServerPatchMapOutput
	ToElasticMapsServerPatchMapOutputWithContext(context.Context) ElasticMapsServerPatchMapOutput
}

type ElasticMapsServerPatchMap map[string]ElasticMapsServerPatchInput

func (ElasticMapsServerPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticMapsServerPatch)(nil)).Elem()
}

func (i ElasticMapsServerPatchMap) ToElasticMapsServerPatchMapOutput() ElasticMapsServerPatchMapOutput {
	return i.ToElasticMapsServerPatchMapOutputWithContext(context.Background())
}

func (i ElasticMapsServerPatchMap) ToElasticMapsServerPatchMapOutputWithContext(ctx context.Context) ElasticMapsServerPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticMapsServerPatchMapOutput)
}

type ElasticMapsServerPatchOutput struct{ *pulumi.OutputState }

func (ElasticMapsServerPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticMapsServerPatch)(nil)).Elem()
}

func (o ElasticMapsServerPatchOutput) ToElasticMapsServerPatchOutput() ElasticMapsServerPatchOutput {
	return o
}

func (o ElasticMapsServerPatchOutput) ToElasticMapsServerPatchOutputWithContext(ctx context.Context) ElasticMapsServerPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ElasticMapsServerPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticMapsServerPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ElasticMapsServerPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticMapsServerPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ElasticMapsServerPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ElasticMapsServerPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o ElasticMapsServerPatchOutput) Spec() ElasticMapsServerSpecPatchPtrOutput {
	return o.ApplyT(func(v *ElasticMapsServerPatch) ElasticMapsServerSpecPatchPtrOutput { return v.Spec }).(ElasticMapsServerSpecPatchPtrOutput)
}

func (o ElasticMapsServerPatchOutput) Status() ElasticMapsServerStatusPatchPtrOutput {
	return o.ApplyT(func(v *ElasticMapsServerPatch) ElasticMapsServerStatusPatchPtrOutput { return v.Status }).(ElasticMapsServerStatusPatchPtrOutput)
}

type ElasticMapsServerPatchArrayOutput struct{ *pulumi.OutputState }

func (ElasticMapsServerPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticMapsServerPatch)(nil)).Elem()
}

func (o ElasticMapsServerPatchArrayOutput) ToElasticMapsServerPatchArrayOutput() ElasticMapsServerPatchArrayOutput {
	return o
}

func (o ElasticMapsServerPatchArrayOutput) ToElasticMapsServerPatchArrayOutputWithContext(ctx context.Context) ElasticMapsServerPatchArrayOutput {
	return o
}

func (o ElasticMapsServerPatchArrayOutput) Index(i pulumi.IntInput) ElasticMapsServerPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElasticMapsServerPatch {
		return vs[0].([]*ElasticMapsServerPatch)[vs[1].(int)]
	}).(ElasticMapsServerPatchOutput)
}

type ElasticMapsServerPatchMapOutput struct{ *pulumi.OutputState }

func (ElasticMapsServerPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticMapsServerPatch)(nil)).Elem()
}

func (o ElasticMapsServerPatchMapOutput) ToElasticMapsServerPatchMapOutput() ElasticMapsServerPatchMapOutput {
	return o
}

func (o ElasticMapsServerPatchMapOutput) ToElasticMapsServerPatchMapOutputWithContext(ctx context.Context) ElasticMapsServerPatchMapOutput {
	return o
}

func (o ElasticMapsServerPatchMapOutput) MapIndex(k pulumi.StringInput) ElasticMapsServerPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElasticMapsServerPatch {
		return vs[0].(map[string]*ElasticMapsServerPatch)[vs[1].(string)]
	}).(ElasticMapsServerPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticMapsServerPatchInput)(nil)).Elem(), &ElasticMapsServerPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticMapsServerPatchArrayInput)(nil)).Elem(), ElasticMapsServerPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticMapsServerPatchMapInput)(nil)).Elem(), ElasticMapsServerPatchMap{})
	pulumi.RegisterOutputType(ElasticMapsServerPatchOutput{})
	pulumi.RegisterOutputType(ElasticMapsServerPatchArrayOutput{})
	pulumi.RegisterOutputType(ElasticMapsServerPatchMapOutput{})
}