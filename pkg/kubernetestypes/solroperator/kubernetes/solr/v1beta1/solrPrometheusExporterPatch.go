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

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// SolrPrometheusExporter is the Schema for the solrprometheusexporters API
type SolrPrometheusExporterPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput            `pulumi:"metadata"`
	Spec     SolrPrometheusExporterSpecPatchPtrOutput   `pulumi:"spec"`
	Status   SolrPrometheusExporterStatusPatchPtrOutput `pulumi:"status"`
}

// NewSolrPrometheusExporterPatch registers a new resource with the given unique name, arguments, and options.
func NewSolrPrometheusExporterPatch(ctx *pulumi.Context,
	name string, args *SolrPrometheusExporterPatchArgs, opts ...pulumi.ResourceOption) (*SolrPrometheusExporterPatch, error) {
	if args == nil {
		args = &SolrPrometheusExporterPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("solr.apache.org/v1beta1")
	args.Kind = pulumi.StringPtr("SolrPrometheusExporter")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SolrPrometheusExporterPatch
	err := ctx.RegisterResource("kubernetes:solr.apache.org/v1beta1:SolrPrometheusExporterPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSolrPrometheusExporterPatch gets an existing SolrPrometheusExporterPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSolrPrometheusExporterPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SolrPrometheusExporterPatchState, opts ...pulumi.ResourceOption) (*SolrPrometheusExporterPatch, error) {
	var resource SolrPrometheusExporterPatch
	err := ctx.ReadResource("kubernetes:solr.apache.org/v1beta1:SolrPrometheusExporterPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SolrPrometheusExporterPatch resources.
type solrPrometheusExporterPatchState struct {
}

type SolrPrometheusExporterPatchState struct {
}

func (SolrPrometheusExporterPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*solrPrometheusExporterPatchState)(nil)).Elem()
}

type solrPrometheusExporterPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch          `pulumi:"metadata"`
	Spec     *SolrPrometheusExporterSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a SolrPrometheusExporterPatch resource.
type SolrPrometheusExporterPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     SolrPrometheusExporterSpecPatchPtrInput
}

func (SolrPrometheusExporterPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*solrPrometheusExporterPatchArgs)(nil)).Elem()
}

type SolrPrometheusExporterPatchInput interface {
	pulumi.Input

	ToSolrPrometheusExporterPatchOutput() SolrPrometheusExporterPatchOutput
	ToSolrPrometheusExporterPatchOutputWithContext(ctx context.Context) SolrPrometheusExporterPatchOutput
}

func (*SolrPrometheusExporterPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**SolrPrometheusExporterPatch)(nil)).Elem()
}

func (i *SolrPrometheusExporterPatch) ToSolrPrometheusExporterPatchOutput() SolrPrometheusExporterPatchOutput {
	return i.ToSolrPrometheusExporterPatchOutputWithContext(context.Background())
}

func (i *SolrPrometheusExporterPatch) ToSolrPrometheusExporterPatchOutputWithContext(ctx context.Context) SolrPrometheusExporterPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolrPrometheusExporterPatchOutput)
}

// SolrPrometheusExporterPatchArrayInput is an input type that accepts SolrPrometheusExporterPatchArray and SolrPrometheusExporterPatchArrayOutput values.
// You can construct a concrete instance of `SolrPrometheusExporterPatchArrayInput` via:
//
//	SolrPrometheusExporterPatchArray{ SolrPrometheusExporterPatchArgs{...} }
type SolrPrometheusExporterPatchArrayInput interface {
	pulumi.Input

	ToSolrPrometheusExporterPatchArrayOutput() SolrPrometheusExporterPatchArrayOutput
	ToSolrPrometheusExporterPatchArrayOutputWithContext(context.Context) SolrPrometheusExporterPatchArrayOutput
}

type SolrPrometheusExporterPatchArray []SolrPrometheusExporterPatchInput

func (SolrPrometheusExporterPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SolrPrometheusExporterPatch)(nil)).Elem()
}

func (i SolrPrometheusExporterPatchArray) ToSolrPrometheusExporterPatchArrayOutput() SolrPrometheusExporterPatchArrayOutput {
	return i.ToSolrPrometheusExporterPatchArrayOutputWithContext(context.Background())
}

func (i SolrPrometheusExporterPatchArray) ToSolrPrometheusExporterPatchArrayOutputWithContext(ctx context.Context) SolrPrometheusExporterPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolrPrometheusExporterPatchArrayOutput)
}

// SolrPrometheusExporterPatchMapInput is an input type that accepts SolrPrometheusExporterPatchMap and SolrPrometheusExporterPatchMapOutput values.
// You can construct a concrete instance of `SolrPrometheusExporterPatchMapInput` via:
//
//	SolrPrometheusExporterPatchMap{ "key": SolrPrometheusExporterPatchArgs{...} }
type SolrPrometheusExporterPatchMapInput interface {
	pulumi.Input

	ToSolrPrometheusExporterPatchMapOutput() SolrPrometheusExporterPatchMapOutput
	ToSolrPrometheusExporterPatchMapOutputWithContext(context.Context) SolrPrometheusExporterPatchMapOutput
}

type SolrPrometheusExporterPatchMap map[string]SolrPrometheusExporterPatchInput

func (SolrPrometheusExporterPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SolrPrometheusExporterPatch)(nil)).Elem()
}

func (i SolrPrometheusExporterPatchMap) ToSolrPrometheusExporterPatchMapOutput() SolrPrometheusExporterPatchMapOutput {
	return i.ToSolrPrometheusExporterPatchMapOutputWithContext(context.Background())
}

func (i SolrPrometheusExporterPatchMap) ToSolrPrometheusExporterPatchMapOutputWithContext(ctx context.Context) SolrPrometheusExporterPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolrPrometheusExporterPatchMapOutput)
}

type SolrPrometheusExporterPatchOutput struct{ *pulumi.OutputState }

func (SolrPrometheusExporterPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolrPrometheusExporterPatch)(nil)).Elem()
}

func (o SolrPrometheusExporterPatchOutput) ToSolrPrometheusExporterPatchOutput() SolrPrometheusExporterPatchOutput {
	return o
}

func (o SolrPrometheusExporterPatchOutput) ToSolrPrometheusExporterPatchOutputWithContext(ctx context.Context) SolrPrometheusExporterPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o SolrPrometheusExporterPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolrPrometheusExporterPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o SolrPrometheusExporterPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolrPrometheusExporterPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o SolrPrometheusExporterPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *SolrPrometheusExporterPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o SolrPrometheusExporterPatchOutput) Spec() SolrPrometheusExporterSpecPatchPtrOutput {
	return o.ApplyT(func(v *SolrPrometheusExporterPatch) SolrPrometheusExporterSpecPatchPtrOutput { return v.Spec }).(SolrPrometheusExporterSpecPatchPtrOutput)
}

func (o SolrPrometheusExporterPatchOutput) Status() SolrPrometheusExporterStatusPatchPtrOutput {
	return o.ApplyT(func(v *SolrPrometheusExporterPatch) SolrPrometheusExporterStatusPatchPtrOutput { return v.Status }).(SolrPrometheusExporterStatusPatchPtrOutput)
}

type SolrPrometheusExporterPatchArrayOutput struct{ *pulumi.OutputState }

func (SolrPrometheusExporterPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SolrPrometheusExporterPatch)(nil)).Elem()
}

func (o SolrPrometheusExporterPatchArrayOutput) ToSolrPrometheusExporterPatchArrayOutput() SolrPrometheusExporterPatchArrayOutput {
	return o
}

func (o SolrPrometheusExporterPatchArrayOutput) ToSolrPrometheusExporterPatchArrayOutputWithContext(ctx context.Context) SolrPrometheusExporterPatchArrayOutput {
	return o
}

func (o SolrPrometheusExporterPatchArrayOutput) Index(i pulumi.IntInput) SolrPrometheusExporterPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SolrPrometheusExporterPatch {
		return vs[0].([]*SolrPrometheusExporterPatch)[vs[1].(int)]
	}).(SolrPrometheusExporterPatchOutput)
}

type SolrPrometheusExporterPatchMapOutput struct{ *pulumi.OutputState }

func (SolrPrometheusExporterPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SolrPrometheusExporterPatch)(nil)).Elem()
}

func (o SolrPrometheusExporterPatchMapOutput) ToSolrPrometheusExporterPatchMapOutput() SolrPrometheusExporterPatchMapOutput {
	return o
}

func (o SolrPrometheusExporterPatchMapOutput) ToSolrPrometheusExporterPatchMapOutputWithContext(ctx context.Context) SolrPrometheusExporterPatchMapOutput {
	return o
}

func (o SolrPrometheusExporterPatchMapOutput) MapIndex(k pulumi.StringInput) SolrPrometheusExporterPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SolrPrometheusExporterPatch {
		return vs[0].(map[string]*SolrPrometheusExporterPatch)[vs[1].(string)]
	}).(SolrPrometheusExporterPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SolrPrometheusExporterPatchInput)(nil)).Elem(), &SolrPrometheusExporterPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolrPrometheusExporterPatchArrayInput)(nil)).Elem(), SolrPrometheusExporterPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolrPrometheusExporterPatchMapInput)(nil)).Elem(), SolrPrometheusExporterPatchMap{})
	pulumi.RegisterOutputType(SolrPrometheusExporterPatchOutput{})
	pulumi.RegisterOutputType(SolrPrometheusExporterPatchArrayOutput{})
	pulumi.RegisterOutputType(SolrPrometheusExporterPatchMapOutput{})
}
