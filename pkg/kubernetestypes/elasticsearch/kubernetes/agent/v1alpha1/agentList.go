// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AgentList is a list of Agent
type AgentList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of agents. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items AgentTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewAgentList registers a new resource with the given unique name, arguments, and options.
func NewAgentList(ctx *pulumi.Context,
	name string, args *AgentListArgs, opts ...pulumi.ResourceOption) (*AgentList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("agent.k8s.elastic.co/v1alpha1")
	args.Kind = pulumi.StringPtr("AgentList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AgentList
	err := ctx.RegisterResource("kubernetes:agent.k8s.elastic.co/v1alpha1:AgentList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentList gets an existing AgentList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentListState, opts ...pulumi.ResourceOption) (*AgentList, error) {
	var resource AgentList
	err := ctx.ReadResource("kubernetes:agent.k8s.elastic.co/v1alpha1:AgentList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentList resources.
type agentListState struct {
}

type AgentListState struct {
}

func (AgentListState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentListState)(nil)).Elem()
}

type agentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of agents. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []AgentType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a AgentList resource.
type AgentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of agents. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items AgentTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (AgentListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentListArgs)(nil)).Elem()
}

type AgentListInput interface {
	pulumi.Input

	ToAgentListOutput() AgentListOutput
	ToAgentListOutputWithContext(ctx context.Context) AgentListOutput
}

func (*AgentList) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentList)(nil)).Elem()
}

func (i *AgentList) ToAgentListOutput() AgentListOutput {
	return i.ToAgentListOutputWithContext(context.Background())
}

func (i *AgentList) ToAgentListOutputWithContext(ctx context.Context) AgentListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentListOutput)
}

// AgentListArrayInput is an input type that accepts AgentListArray and AgentListArrayOutput values.
// You can construct a concrete instance of `AgentListArrayInput` via:
//
//	AgentListArray{ AgentListArgs{...} }
type AgentListArrayInput interface {
	pulumi.Input

	ToAgentListArrayOutput() AgentListArrayOutput
	ToAgentListArrayOutputWithContext(context.Context) AgentListArrayOutput
}

type AgentListArray []AgentListInput

func (AgentListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AgentList)(nil)).Elem()
}

func (i AgentListArray) ToAgentListArrayOutput() AgentListArrayOutput {
	return i.ToAgentListArrayOutputWithContext(context.Background())
}

func (i AgentListArray) ToAgentListArrayOutputWithContext(ctx context.Context) AgentListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentListArrayOutput)
}

// AgentListMapInput is an input type that accepts AgentListMap and AgentListMapOutput values.
// You can construct a concrete instance of `AgentListMapInput` via:
//
//	AgentListMap{ "key": AgentListArgs{...} }
type AgentListMapInput interface {
	pulumi.Input

	ToAgentListMapOutput() AgentListMapOutput
	ToAgentListMapOutputWithContext(context.Context) AgentListMapOutput
}

type AgentListMap map[string]AgentListInput

func (AgentListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AgentList)(nil)).Elem()
}

func (i AgentListMap) ToAgentListMapOutput() AgentListMapOutput {
	return i.ToAgentListMapOutputWithContext(context.Background())
}

func (i AgentListMap) ToAgentListMapOutputWithContext(ctx context.Context) AgentListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentListMapOutput)
}

type AgentListOutput struct{ *pulumi.OutputState }

func (AgentListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentList)(nil)).Elem()
}

func (o AgentListOutput) ToAgentListOutput() AgentListOutput {
	return o
}

func (o AgentListOutput) ToAgentListOutputWithContext(ctx context.Context) AgentListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o AgentListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of agents. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o AgentListOutput) Items() AgentTypeArrayOutput {
	return o.ApplyT(func(v *AgentList) AgentTypeArrayOutput { return v.Items }).(AgentTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o AgentListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o AgentListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *AgentList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type AgentListArrayOutput struct{ *pulumi.OutputState }

func (AgentListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AgentList)(nil)).Elem()
}

func (o AgentListArrayOutput) ToAgentListArrayOutput() AgentListArrayOutput {
	return o
}

func (o AgentListArrayOutput) ToAgentListArrayOutputWithContext(ctx context.Context) AgentListArrayOutput {
	return o
}

func (o AgentListArrayOutput) Index(i pulumi.IntInput) AgentListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AgentList {
		return vs[0].([]*AgentList)[vs[1].(int)]
	}).(AgentListOutput)
}

type AgentListMapOutput struct{ *pulumi.OutputState }

func (AgentListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AgentList)(nil)).Elem()
}

func (o AgentListMapOutput) ToAgentListMapOutput() AgentListMapOutput {
	return o
}

func (o AgentListMapOutput) ToAgentListMapOutputWithContext(ctx context.Context) AgentListMapOutput {
	return o
}

func (o AgentListMapOutput) MapIndex(k pulumi.StringInput) AgentListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AgentList {
		return vs[0].(map[string]*AgentList)[vs[1].(string)]
	}).(AgentListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgentListInput)(nil)).Elem(), &AgentList{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentListArrayInput)(nil)).Elem(), AgentListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentListMapInput)(nil)).Elem(), AgentListMap{})
	pulumi.RegisterOutputType(AgentListOutput{})
	pulumi.RegisterOutputType(AgentListArrayOutput{})
	pulumi.RegisterOutputType(AgentListMapOutput{})
}
