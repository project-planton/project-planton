// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha1

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
type KeycloakRealmImportPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput         `pulumi:"metadata"`
	Spec     KeycloakRealmImportSpecPatchPtrOutput   `pulumi:"spec"`
	Status   KeycloakRealmImportStatusPatchPtrOutput `pulumi:"status"`
}

// NewKeycloakRealmImportPatch registers a new resource with the given unique name, arguments, and options.
func NewKeycloakRealmImportPatch(ctx *pulumi.Context,
	name string, args *KeycloakRealmImportPatchArgs, opts ...pulumi.ResourceOption) (*KeycloakRealmImportPatch, error) {
	if args == nil {
		args = &KeycloakRealmImportPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("k8s.keycloak.org/v2alpha1")
	args.Kind = pulumi.StringPtr("KeycloakRealmImport")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KeycloakRealmImportPatch
	err := ctx.RegisterResource("kubernetes:k8s.keycloak.org/v2alpha1:KeycloakRealmImportPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeycloakRealmImportPatch gets an existing KeycloakRealmImportPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeycloakRealmImportPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeycloakRealmImportPatchState, opts ...pulumi.ResourceOption) (*KeycloakRealmImportPatch, error) {
	var resource KeycloakRealmImportPatch
	err := ctx.ReadResource("kubernetes:k8s.keycloak.org/v2alpha1:KeycloakRealmImportPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeycloakRealmImportPatch resources.
type keycloakRealmImportPatchState struct {
}

type KeycloakRealmImportPatchState struct {
}

func (KeycloakRealmImportPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*keycloakRealmImportPatchState)(nil)).Elem()
}

type keycloakRealmImportPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch       `pulumi:"metadata"`
	Spec     *KeycloakRealmImportSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a KeycloakRealmImportPatch resource.
type KeycloakRealmImportPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     KeycloakRealmImportSpecPatchPtrInput
}

func (KeycloakRealmImportPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keycloakRealmImportPatchArgs)(nil)).Elem()
}

type KeycloakRealmImportPatchInput interface {
	pulumi.Input

	ToKeycloakRealmImportPatchOutput() KeycloakRealmImportPatchOutput
	ToKeycloakRealmImportPatchOutputWithContext(ctx context.Context) KeycloakRealmImportPatchOutput
}

func (*KeycloakRealmImportPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**KeycloakRealmImportPatch)(nil)).Elem()
}

func (i *KeycloakRealmImportPatch) ToKeycloakRealmImportPatchOutput() KeycloakRealmImportPatchOutput {
	return i.ToKeycloakRealmImportPatchOutputWithContext(context.Background())
}

func (i *KeycloakRealmImportPatch) ToKeycloakRealmImportPatchOutputWithContext(ctx context.Context) KeycloakRealmImportPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeycloakRealmImportPatchOutput)
}

// KeycloakRealmImportPatchArrayInput is an input type that accepts KeycloakRealmImportPatchArray and KeycloakRealmImportPatchArrayOutput values.
// You can construct a concrete instance of `KeycloakRealmImportPatchArrayInput` via:
//
//	KeycloakRealmImportPatchArray{ KeycloakRealmImportPatchArgs{...} }
type KeycloakRealmImportPatchArrayInput interface {
	pulumi.Input

	ToKeycloakRealmImportPatchArrayOutput() KeycloakRealmImportPatchArrayOutput
	ToKeycloakRealmImportPatchArrayOutputWithContext(context.Context) KeycloakRealmImportPatchArrayOutput
}

type KeycloakRealmImportPatchArray []KeycloakRealmImportPatchInput

func (KeycloakRealmImportPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeycloakRealmImportPatch)(nil)).Elem()
}

func (i KeycloakRealmImportPatchArray) ToKeycloakRealmImportPatchArrayOutput() KeycloakRealmImportPatchArrayOutput {
	return i.ToKeycloakRealmImportPatchArrayOutputWithContext(context.Background())
}

func (i KeycloakRealmImportPatchArray) ToKeycloakRealmImportPatchArrayOutputWithContext(ctx context.Context) KeycloakRealmImportPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeycloakRealmImportPatchArrayOutput)
}

// KeycloakRealmImportPatchMapInput is an input type that accepts KeycloakRealmImportPatchMap and KeycloakRealmImportPatchMapOutput values.
// You can construct a concrete instance of `KeycloakRealmImportPatchMapInput` via:
//
//	KeycloakRealmImportPatchMap{ "key": KeycloakRealmImportPatchArgs{...} }
type KeycloakRealmImportPatchMapInput interface {
	pulumi.Input

	ToKeycloakRealmImportPatchMapOutput() KeycloakRealmImportPatchMapOutput
	ToKeycloakRealmImportPatchMapOutputWithContext(context.Context) KeycloakRealmImportPatchMapOutput
}

type KeycloakRealmImportPatchMap map[string]KeycloakRealmImportPatchInput

func (KeycloakRealmImportPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeycloakRealmImportPatch)(nil)).Elem()
}

func (i KeycloakRealmImportPatchMap) ToKeycloakRealmImportPatchMapOutput() KeycloakRealmImportPatchMapOutput {
	return i.ToKeycloakRealmImportPatchMapOutputWithContext(context.Background())
}

func (i KeycloakRealmImportPatchMap) ToKeycloakRealmImportPatchMapOutputWithContext(ctx context.Context) KeycloakRealmImportPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeycloakRealmImportPatchMapOutput)
}

type KeycloakRealmImportPatchOutput struct{ *pulumi.OutputState }

func (KeycloakRealmImportPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeycloakRealmImportPatch)(nil)).Elem()
}

func (o KeycloakRealmImportPatchOutput) ToKeycloakRealmImportPatchOutput() KeycloakRealmImportPatchOutput {
	return o
}

func (o KeycloakRealmImportPatchOutput) ToKeycloakRealmImportPatchOutputWithContext(ctx context.Context) KeycloakRealmImportPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o KeycloakRealmImportPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeycloakRealmImportPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KeycloakRealmImportPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeycloakRealmImportPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o KeycloakRealmImportPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *KeycloakRealmImportPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o KeycloakRealmImportPatchOutput) Spec() KeycloakRealmImportSpecPatchPtrOutput {
	return o.ApplyT(func(v *KeycloakRealmImportPatch) KeycloakRealmImportSpecPatchPtrOutput { return v.Spec }).(KeycloakRealmImportSpecPatchPtrOutput)
}

func (o KeycloakRealmImportPatchOutput) Status() KeycloakRealmImportStatusPatchPtrOutput {
	return o.ApplyT(func(v *KeycloakRealmImportPatch) KeycloakRealmImportStatusPatchPtrOutput { return v.Status }).(KeycloakRealmImportStatusPatchPtrOutput)
}

type KeycloakRealmImportPatchArrayOutput struct{ *pulumi.OutputState }

func (KeycloakRealmImportPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeycloakRealmImportPatch)(nil)).Elem()
}

func (o KeycloakRealmImportPatchArrayOutput) ToKeycloakRealmImportPatchArrayOutput() KeycloakRealmImportPatchArrayOutput {
	return o
}

func (o KeycloakRealmImportPatchArrayOutput) ToKeycloakRealmImportPatchArrayOutputWithContext(ctx context.Context) KeycloakRealmImportPatchArrayOutput {
	return o
}

func (o KeycloakRealmImportPatchArrayOutput) Index(i pulumi.IntInput) KeycloakRealmImportPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KeycloakRealmImportPatch {
		return vs[0].([]*KeycloakRealmImportPatch)[vs[1].(int)]
	}).(KeycloakRealmImportPatchOutput)
}

type KeycloakRealmImportPatchMapOutput struct{ *pulumi.OutputState }

func (KeycloakRealmImportPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeycloakRealmImportPatch)(nil)).Elem()
}

func (o KeycloakRealmImportPatchMapOutput) ToKeycloakRealmImportPatchMapOutput() KeycloakRealmImportPatchMapOutput {
	return o
}

func (o KeycloakRealmImportPatchMapOutput) ToKeycloakRealmImportPatchMapOutputWithContext(ctx context.Context) KeycloakRealmImportPatchMapOutput {
	return o
}

func (o KeycloakRealmImportPatchMapOutput) MapIndex(k pulumi.StringInput) KeycloakRealmImportPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KeycloakRealmImportPatch {
		return vs[0].(map[string]*KeycloakRealmImportPatch)[vs[1].(string)]
	}).(KeycloakRealmImportPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeycloakRealmImportPatchInput)(nil)).Elem(), &KeycloakRealmImportPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeycloakRealmImportPatchArrayInput)(nil)).Elem(), KeycloakRealmImportPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeycloakRealmImportPatchMapInput)(nil)).Elem(), KeycloakRealmImportPatchMap{})
	pulumi.RegisterOutputType(KeycloakRealmImportPatchOutput{})
	pulumi.RegisterOutputType(KeycloakRealmImportPatchArrayOutput{})
	pulumi.RegisterOutputType(KeycloakRealmImportPatchMapOutput{})
}
