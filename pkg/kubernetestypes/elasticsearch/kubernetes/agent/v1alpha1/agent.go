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

// Agent is the Schema for the Agents API.
type Agent struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	Spec     AgentSpecOutput         `pulumi:"spec"`
	Status   AgentStatusPtrOutput    `pulumi:"status"`
}

// NewAgent registers a new resource with the given unique name, arguments, and options.
func NewAgent(ctx *pulumi.Context,
	name string, args *AgentArgs, opts ...pulumi.ResourceOption) (*Agent, error) {
	if args == nil {
		args = &AgentArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("agent.k8s.elastic.co/v1alpha1")
	args.Kind = pulumi.StringPtr("Agent")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Agent
	err := ctx.RegisterResource("kubernetes:agent.k8s.elastic.co/v1alpha1:Agent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgent gets an existing Agent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentState, opts ...pulumi.ResourceOption) (*Agent, error) {
	var resource Agent
	err := ctx.ReadResource("kubernetes:agent.k8s.elastic.co/v1alpha1:Agent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Agent resources.
type agentState struct {
}

type AgentState struct {
}

func (AgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentState)(nil)).Elem()
}

type agentArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *AgentSpec         `pulumi:"spec"`
}

// The set of arguments for constructing a Agent resource.
type AgentArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     AgentSpecPtrInput
}

func (AgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentArgs)(nil)).Elem()
}

type AgentInput interface {
	pulumi.Input

	ToAgentOutput() AgentOutput
	ToAgentOutputWithContext(ctx context.Context) AgentOutput
}

func (*Agent) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil)).Elem()
}

func (i *Agent) ToAgentOutput() AgentOutput {
	return i.ToAgentOutputWithContext(context.Background())
}

func (i *Agent) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentOutput)
}

// AgentArrayInput is an input type that accepts AgentArray and AgentArrayOutput values.
// You can construct a concrete instance of `AgentArrayInput` via:
//
//	AgentArray{ AgentArgs{...} }
type AgentArrayInput interface {
	pulumi.Input

	ToAgentArrayOutput() AgentArrayOutput
	ToAgentArrayOutputWithContext(context.Context) AgentArrayOutput
}

type AgentArray []AgentInput

func (AgentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Agent)(nil)).Elem()
}

func (i AgentArray) ToAgentArrayOutput() AgentArrayOutput {
	return i.ToAgentArrayOutputWithContext(context.Background())
}

func (i AgentArray) ToAgentArrayOutputWithContext(ctx context.Context) AgentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentArrayOutput)
}

// AgentMapInput is an input type that accepts AgentMap and AgentMapOutput values.
// You can construct a concrete instance of `AgentMapInput` via:
//
//	AgentMap{ "key": AgentArgs{...} }
type AgentMapInput interface {
	pulumi.Input

	ToAgentMapOutput() AgentMapOutput
	ToAgentMapOutputWithContext(context.Context) AgentMapOutput
}

type AgentMap map[string]AgentInput

func (AgentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Agent)(nil)).Elem()
}

func (i AgentMap) ToAgentMapOutput() AgentMapOutput {
	return i.ToAgentMapOutputWithContext(context.Background())
}

func (i AgentMap) ToAgentMapOutputWithContext(ctx context.Context) AgentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentMapOutput)
}

type AgentOutput struct{ *pulumi.OutputState }

func (AgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil)).Elem()
}

func (o AgentOutput) ToAgentOutput() AgentOutput {
	return o
}

func (o AgentOutput) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o AgentOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o AgentOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o AgentOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *Agent) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o AgentOutput) Spec() AgentSpecOutput {
	return o.ApplyT(func(v *Agent) AgentSpecOutput { return v.Spec }).(AgentSpecOutput)
}

func (o AgentOutput) Status() AgentStatusPtrOutput {
	return o.ApplyT(func(v *Agent) AgentStatusPtrOutput { return v.Status }).(AgentStatusPtrOutput)
}

type AgentArrayOutput struct{ *pulumi.OutputState }

func (AgentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Agent)(nil)).Elem()
}

func (o AgentArrayOutput) ToAgentArrayOutput() AgentArrayOutput {
	return o
}

func (o AgentArrayOutput) ToAgentArrayOutputWithContext(ctx context.Context) AgentArrayOutput {
	return o
}

func (o AgentArrayOutput) Index(i pulumi.IntInput) AgentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Agent {
		return vs[0].([]*Agent)[vs[1].(int)]
	}).(AgentOutput)
}

type AgentMapOutput struct{ *pulumi.OutputState }

func (AgentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Agent)(nil)).Elem()
}

func (o AgentMapOutput) ToAgentMapOutput() AgentMapOutput {
	return o
}

func (o AgentMapOutput) ToAgentMapOutputWithContext(ctx context.Context) AgentMapOutput {
	return o
}

func (o AgentMapOutput) MapIndex(k pulumi.StringInput) AgentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Agent {
		return vs[0].(map[string]*Agent)[vs[1].(string)]
	}).(AgentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgentInput)(nil)).Elem(), &Agent{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentArrayInput)(nil)).Elem(), AgentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentMapInput)(nil)).Elem(), AgentMap{})
	pulumi.RegisterOutputType(AgentOutput{})
	pulumi.RegisterOutputType(AgentArrayOutput{})
	pulumi.RegisterOutputType(AgentMapOutput{})
}
