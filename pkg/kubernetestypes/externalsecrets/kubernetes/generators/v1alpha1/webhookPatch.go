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
// Webhook connects to a third party API server to handle the secrets generation
// configuration parameters in spec.
// You can specify the server, the token, and additional body parameters.
// See documentation for the full API specification for requests and responses.
type WebhookPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     WebhookSpecPatchPtrOutput       `pulumi:"spec"`
}

// NewWebhookPatch registers a new resource with the given unique name, arguments, and options.
func NewWebhookPatch(ctx *pulumi.Context,
	name string, args *WebhookPatchArgs, opts ...pulumi.ResourceOption) (*WebhookPatch, error) {
	if args == nil {
		args = &WebhookPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("Webhook")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebhookPatch
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:WebhookPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhookPatch gets an existing WebhookPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhookPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookPatchState, opts ...pulumi.ResourceOption) (*WebhookPatch, error) {
	var resource WebhookPatch
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:WebhookPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebhookPatch resources.
type webhookPatchState struct {
}

type WebhookPatchState struct {
}

func (WebhookPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookPatchState)(nil)).Elem()
}

type webhookPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *WebhookSpecPatch       `pulumi:"spec"`
}

// The set of arguments for constructing a WebhookPatch resource.
type WebhookPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     WebhookSpecPatchPtrInput
}

func (WebhookPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookPatchArgs)(nil)).Elem()
}

type WebhookPatchInput interface {
	pulumi.Input

	ToWebhookPatchOutput() WebhookPatchOutput
	ToWebhookPatchOutputWithContext(ctx context.Context) WebhookPatchOutput
}

func (*WebhookPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookPatch)(nil)).Elem()
}

func (i *WebhookPatch) ToWebhookPatchOutput() WebhookPatchOutput {
	return i.ToWebhookPatchOutputWithContext(context.Background())
}

func (i *WebhookPatch) ToWebhookPatchOutputWithContext(ctx context.Context) WebhookPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPatchOutput)
}

// WebhookPatchArrayInput is an input type that accepts WebhookPatchArray and WebhookPatchArrayOutput values.
// You can construct a concrete instance of `WebhookPatchArrayInput` via:
//
//	WebhookPatchArray{ WebhookPatchArgs{...} }
type WebhookPatchArrayInput interface {
	pulumi.Input

	ToWebhookPatchArrayOutput() WebhookPatchArrayOutput
	ToWebhookPatchArrayOutputWithContext(context.Context) WebhookPatchArrayOutput
}

type WebhookPatchArray []WebhookPatchInput

func (WebhookPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebhookPatch)(nil)).Elem()
}

func (i WebhookPatchArray) ToWebhookPatchArrayOutput() WebhookPatchArrayOutput {
	return i.ToWebhookPatchArrayOutputWithContext(context.Background())
}

func (i WebhookPatchArray) ToWebhookPatchArrayOutputWithContext(ctx context.Context) WebhookPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPatchArrayOutput)
}

// WebhookPatchMapInput is an input type that accepts WebhookPatchMap and WebhookPatchMapOutput values.
// You can construct a concrete instance of `WebhookPatchMapInput` via:
//
//	WebhookPatchMap{ "key": WebhookPatchArgs{...} }
type WebhookPatchMapInput interface {
	pulumi.Input

	ToWebhookPatchMapOutput() WebhookPatchMapOutput
	ToWebhookPatchMapOutputWithContext(context.Context) WebhookPatchMapOutput
}

type WebhookPatchMap map[string]WebhookPatchInput

func (WebhookPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebhookPatch)(nil)).Elem()
}

func (i WebhookPatchMap) ToWebhookPatchMapOutput() WebhookPatchMapOutput {
	return i.ToWebhookPatchMapOutputWithContext(context.Background())
}

func (i WebhookPatchMap) ToWebhookPatchMapOutputWithContext(ctx context.Context) WebhookPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPatchMapOutput)
}

type WebhookPatchOutput struct{ *pulumi.OutputState }

func (WebhookPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookPatch)(nil)).Elem()
}

func (o WebhookPatchOutput) ToWebhookPatchOutput() WebhookPatchOutput {
	return o
}

func (o WebhookPatchOutput) ToWebhookPatchOutputWithContext(ctx context.Context) WebhookPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o WebhookPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o WebhookPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o WebhookPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *WebhookPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o WebhookPatchOutput) Spec() WebhookSpecPatchPtrOutput {
	return o.ApplyT(func(v *WebhookPatch) WebhookSpecPatchPtrOutput { return v.Spec }).(WebhookSpecPatchPtrOutput)
}

type WebhookPatchArrayOutput struct{ *pulumi.OutputState }

func (WebhookPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebhookPatch)(nil)).Elem()
}

func (o WebhookPatchArrayOutput) ToWebhookPatchArrayOutput() WebhookPatchArrayOutput {
	return o
}

func (o WebhookPatchArrayOutput) ToWebhookPatchArrayOutputWithContext(ctx context.Context) WebhookPatchArrayOutput {
	return o
}

func (o WebhookPatchArrayOutput) Index(i pulumi.IntInput) WebhookPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebhookPatch {
		return vs[0].([]*WebhookPatch)[vs[1].(int)]
	}).(WebhookPatchOutput)
}

type WebhookPatchMapOutput struct{ *pulumi.OutputState }

func (WebhookPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebhookPatch)(nil)).Elem()
}

func (o WebhookPatchMapOutput) ToWebhookPatchMapOutput() WebhookPatchMapOutput {
	return o
}

func (o WebhookPatchMapOutput) ToWebhookPatchMapOutputWithContext(ctx context.Context) WebhookPatchMapOutput {
	return o
}

func (o WebhookPatchMapOutput) MapIndex(k pulumi.StringInput) WebhookPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebhookPatch {
		return vs[0].(map[string]*WebhookPatch)[vs[1].(string)]
	}).(WebhookPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookPatchInput)(nil)).Elem(), &WebhookPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookPatchArrayInput)(nil)).Elem(), WebhookPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookPatchMapInput)(nil)).Elem(), WebhookPatchMap{})
	pulumi.RegisterOutputType(WebhookPatchOutput{})
	pulumi.RegisterOutputType(WebhookPatchArrayOutput{})
	pulumi.RegisterOutputType(WebhookPatchMapOutput{})
}