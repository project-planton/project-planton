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
// ECRAuthorizationTokenSpec uses the GetAuthorizationToken API to retrieve an
// authorization token.
// The authorization token is valid for 12 hours.
// The authorizationToken returned is a base64 encoded string that can be decoded
// and used in a docker login command to authenticate to a registry.
// For more information, see Registry authentication (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth) in the Amazon Elastic Container Registry User Guide.
type ECRAuthorizationTokenPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput         `pulumi:"metadata"`
	Spec     ECRAuthorizationTokenSpecPatchPtrOutput `pulumi:"spec"`
}

// NewECRAuthorizationTokenPatch registers a new resource with the given unique name, arguments, and options.
func NewECRAuthorizationTokenPatch(ctx *pulumi.Context,
	name string, args *ECRAuthorizationTokenPatchArgs, opts ...pulumi.ResourceOption) (*ECRAuthorizationTokenPatch, error) {
	if args == nil {
		args = &ECRAuthorizationTokenPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ECRAuthorizationToken")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ECRAuthorizationTokenPatch
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:ECRAuthorizationTokenPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetECRAuthorizationTokenPatch gets an existing ECRAuthorizationTokenPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetECRAuthorizationTokenPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ECRAuthorizationTokenPatchState, opts ...pulumi.ResourceOption) (*ECRAuthorizationTokenPatch, error) {
	var resource ECRAuthorizationTokenPatch
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:ECRAuthorizationTokenPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ECRAuthorizationTokenPatch resources.
type ecrauthorizationTokenPatchState struct {
}

type ECRAuthorizationTokenPatchState struct {
}

func (ECRAuthorizationTokenPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecrauthorizationTokenPatchState)(nil)).Elem()
}

type ecrauthorizationTokenPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch         `pulumi:"metadata"`
	Spec     *ECRAuthorizationTokenSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a ECRAuthorizationTokenPatch resource.
type ECRAuthorizationTokenPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     ECRAuthorizationTokenSpecPatchPtrInput
}

func (ECRAuthorizationTokenPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecrauthorizationTokenPatchArgs)(nil)).Elem()
}

type ECRAuthorizationTokenPatchInput interface {
	pulumi.Input

	ToECRAuthorizationTokenPatchOutput() ECRAuthorizationTokenPatchOutput
	ToECRAuthorizationTokenPatchOutputWithContext(ctx context.Context) ECRAuthorizationTokenPatchOutput
}

func (*ECRAuthorizationTokenPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ECRAuthorizationTokenPatch)(nil)).Elem()
}

func (i *ECRAuthorizationTokenPatch) ToECRAuthorizationTokenPatchOutput() ECRAuthorizationTokenPatchOutput {
	return i.ToECRAuthorizationTokenPatchOutputWithContext(context.Background())
}

func (i *ECRAuthorizationTokenPatch) ToECRAuthorizationTokenPatchOutputWithContext(ctx context.Context) ECRAuthorizationTokenPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECRAuthorizationTokenPatchOutput)
}

// ECRAuthorizationTokenPatchArrayInput is an input type that accepts ECRAuthorizationTokenPatchArray and ECRAuthorizationTokenPatchArrayOutput values.
// You can construct a concrete instance of `ECRAuthorizationTokenPatchArrayInput` via:
//
//	ECRAuthorizationTokenPatchArray{ ECRAuthorizationTokenPatchArgs{...} }
type ECRAuthorizationTokenPatchArrayInput interface {
	pulumi.Input

	ToECRAuthorizationTokenPatchArrayOutput() ECRAuthorizationTokenPatchArrayOutput
	ToECRAuthorizationTokenPatchArrayOutputWithContext(context.Context) ECRAuthorizationTokenPatchArrayOutput
}

type ECRAuthorizationTokenPatchArray []ECRAuthorizationTokenPatchInput

func (ECRAuthorizationTokenPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ECRAuthorizationTokenPatch)(nil)).Elem()
}

func (i ECRAuthorizationTokenPatchArray) ToECRAuthorizationTokenPatchArrayOutput() ECRAuthorizationTokenPatchArrayOutput {
	return i.ToECRAuthorizationTokenPatchArrayOutputWithContext(context.Background())
}

func (i ECRAuthorizationTokenPatchArray) ToECRAuthorizationTokenPatchArrayOutputWithContext(ctx context.Context) ECRAuthorizationTokenPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECRAuthorizationTokenPatchArrayOutput)
}

// ECRAuthorizationTokenPatchMapInput is an input type that accepts ECRAuthorizationTokenPatchMap and ECRAuthorizationTokenPatchMapOutput values.
// You can construct a concrete instance of `ECRAuthorizationTokenPatchMapInput` via:
//
//	ECRAuthorizationTokenPatchMap{ "key": ECRAuthorizationTokenPatchArgs{...} }
type ECRAuthorizationTokenPatchMapInput interface {
	pulumi.Input

	ToECRAuthorizationTokenPatchMapOutput() ECRAuthorizationTokenPatchMapOutput
	ToECRAuthorizationTokenPatchMapOutputWithContext(context.Context) ECRAuthorizationTokenPatchMapOutput
}

type ECRAuthorizationTokenPatchMap map[string]ECRAuthorizationTokenPatchInput

func (ECRAuthorizationTokenPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ECRAuthorizationTokenPatch)(nil)).Elem()
}

func (i ECRAuthorizationTokenPatchMap) ToECRAuthorizationTokenPatchMapOutput() ECRAuthorizationTokenPatchMapOutput {
	return i.ToECRAuthorizationTokenPatchMapOutputWithContext(context.Background())
}

func (i ECRAuthorizationTokenPatchMap) ToECRAuthorizationTokenPatchMapOutputWithContext(ctx context.Context) ECRAuthorizationTokenPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECRAuthorizationTokenPatchMapOutput)
}

type ECRAuthorizationTokenPatchOutput struct{ *pulumi.OutputState }

func (ECRAuthorizationTokenPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ECRAuthorizationTokenPatch)(nil)).Elem()
}

func (o ECRAuthorizationTokenPatchOutput) ToECRAuthorizationTokenPatchOutput() ECRAuthorizationTokenPatchOutput {
	return o
}

func (o ECRAuthorizationTokenPatchOutput) ToECRAuthorizationTokenPatchOutputWithContext(ctx context.Context) ECRAuthorizationTokenPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ECRAuthorizationTokenPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ECRAuthorizationTokenPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ECRAuthorizationTokenPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ECRAuthorizationTokenPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ECRAuthorizationTokenPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ECRAuthorizationTokenPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o ECRAuthorizationTokenPatchOutput) Spec() ECRAuthorizationTokenSpecPatchPtrOutput {
	return o.ApplyT(func(v *ECRAuthorizationTokenPatch) ECRAuthorizationTokenSpecPatchPtrOutput { return v.Spec }).(ECRAuthorizationTokenSpecPatchPtrOutput)
}

type ECRAuthorizationTokenPatchArrayOutput struct{ *pulumi.OutputState }

func (ECRAuthorizationTokenPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ECRAuthorizationTokenPatch)(nil)).Elem()
}

func (o ECRAuthorizationTokenPatchArrayOutput) ToECRAuthorizationTokenPatchArrayOutput() ECRAuthorizationTokenPatchArrayOutput {
	return o
}

func (o ECRAuthorizationTokenPatchArrayOutput) ToECRAuthorizationTokenPatchArrayOutputWithContext(ctx context.Context) ECRAuthorizationTokenPatchArrayOutput {
	return o
}

func (o ECRAuthorizationTokenPatchArrayOutput) Index(i pulumi.IntInput) ECRAuthorizationTokenPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ECRAuthorizationTokenPatch {
		return vs[0].([]*ECRAuthorizationTokenPatch)[vs[1].(int)]
	}).(ECRAuthorizationTokenPatchOutput)
}

type ECRAuthorizationTokenPatchMapOutput struct{ *pulumi.OutputState }

func (ECRAuthorizationTokenPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ECRAuthorizationTokenPatch)(nil)).Elem()
}

func (o ECRAuthorizationTokenPatchMapOutput) ToECRAuthorizationTokenPatchMapOutput() ECRAuthorizationTokenPatchMapOutput {
	return o
}

func (o ECRAuthorizationTokenPatchMapOutput) ToECRAuthorizationTokenPatchMapOutputWithContext(ctx context.Context) ECRAuthorizationTokenPatchMapOutput {
	return o
}

func (o ECRAuthorizationTokenPatchMapOutput) MapIndex(k pulumi.StringInput) ECRAuthorizationTokenPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ECRAuthorizationTokenPatch {
		return vs[0].(map[string]*ECRAuthorizationTokenPatch)[vs[1].(string)]
	}).(ECRAuthorizationTokenPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ECRAuthorizationTokenPatchInput)(nil)).Elem(), &ECRAuthorizationTokenPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ECRAuthorizationTokenPatchArrayInput)(nil)).Elem(), ECRAuthorizationTokenPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ECRAuthorizationTokenPatchMapInput)(nil)).Elem(), ECRAuthorizationTokenPatchMap{})
	pulumi.RegisterOutputType(ECRAuthorizationTokenPatchOutput{})
	pulumi.RegisterOutputType(ECRAuthorizationTokenPatchArrayOutput{})
	pulumi.RegisterOutputType(ECRAuthorizationTokenPatchMapOutput{})
}