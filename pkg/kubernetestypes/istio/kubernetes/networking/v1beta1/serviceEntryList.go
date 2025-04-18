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

// ServiceEntryList is a list of ServiceEntry
type ServiceEntryList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of serviceentries. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ServiceEntryTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewServiceEntryList registers a new resource with the given unique name, arguments, and options.
func NewServiceEntryList(ctx *pulumi.Context,
	name string, args *ServiceEntryListArgs, opts ...pulumi.ResourceOption) (*ServiceEntryList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1beta1")
	args.Kind = pulumi.StringPtr("ServiceEntryList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServiceEntryList
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1beta1:ServiceEntryList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEntryList gets an existing ServiceEntryList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEntryList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEntryListState, opts ...pulumi.ResourceOption) (*ServiceEntryList, error) {
	var resource ServiceEntryList
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1beta1:ServiceEntryList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEntryList resources.
type serviceEntryListState struct {
}

type ServiceEntryListState struct {
}

func (ServiceEntryListState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEntryListState)(nil)).Elem()
}

type serviceEntryListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of serviceentries. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ServiceEntryType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ServiceEntryList resource.
type ServiceEntryListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of serviceentries. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ServiceEntryTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ServiceEntryListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEntryListArgs)(nil)).Elem()
}

type ServiceEntryListInput interface {
	pulumi.Input

	ToServiceEntryListOutput() ServiceEntryListOutput
	ToServiceEntryListOutputWithContext(ctx context.Context) ServiceEntryListOutput
}

func (*ServiceEntryList) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEntryList)(nil)).Elem()
}

func (i *ServiceEntryList) ToServiceEntryListOutput() ServiceEntryListOutput {
	return i.ToServiceEntryListOutputWithContext(context.Background())
}

func (i *ServiceEntryList) ToServiceEntryListOutputWithContext(ctx context.Context) ServiceEntryListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEntryListOutput)
}

// ServiceEntryListArrayInput is an input type that accepts ServiceEntryListArray and ServiceEntryListArrayOutput values.
// You can construct a concrete instance of `ServiceEntryListArrayInput` via:
//
//	ServiceEntryListArray{ ServiceEntryListArgs{...} }
type ServiceEntryListArrayInput interface {
	pulumi.Input

	ToServiceEntryListArrayOutput() ServiceEntryListArrayOutput
	ToServiceEntryListArrayOutputWithContext(context.Context) ServiceEntryListArrayOutput
}

type ServiceEntryListArray []ServiceEntryListInput

func (ServiceEntryListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEntryList)(nil)).Elem()
}

func (i ServiceEntryListArray) ToServiceEntryListArrayOutput() ServiceEntryListArrayOutput {
	return i.ToServiceEntryListArrayOutputWithContext(context.Background())
}

func (i ServiceEntryListArray) ToServiceEntryListArrayOutputWithContext(ctx context.Context) ServiceEntryListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEntryListArrayOutput)
}

// ServiceEntryListMapInput is an input type that accepts ServiceEntryListMap and ServiceEntryListMapOutput values.
// You can construct a concrete instance of `ServiceEntryListMapInput` via:
//
//	ServiceEntryListMap{ "key": ServiceEntryListArgs{...} }
type ServiceEntryListMapInput interface {
	pulumi.Input

	ToServiceEntryListMapOutput() ServiceEntryListMapOutput
	ToServiceEntryListMapOutputWithContext(context.Context) ServiceEntryListMapOutput
}

type ServiceEntryListMap map[string]ServiceEntryListInput

func (ServiceEntryListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEntryList)(nil)).Elem()
}

func (i ServiceEntryListMap) ToServiceEntryListMapOutput() ServiceEntryListMapOutput {
	return i.ToServiceEntryListMapOutputWithContext(context.Background())
}

func (i ServiceEntryListMap) ToServiceEntryListMapOutputWithContext(ctx context.Context) ServiceEntryListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEntryListMapOutput)
}

type ServiceEntryListOutput struct{ *pulumi.OutputState }

func (ServiceEntryListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEntryList)(nil)).Elem()
}

func (o ServiceEntryListOutput) ToServiceEntryListOutput() ServiceEntryListOutput {
	return o
}

func (o ServiceEntryListOutput) ToServiceEntryListOutputWithContext(ctx context.Context) ServiceEntryListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ServiceEntryListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEntryList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of serviceentries. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ServiceEntryListOutput) Items() ServiceEntryTypeArrayOutput {
	return o.ApplyT(func(v *ServiceEntryList) ServiceEntryTypeArrayOutput { return v.Items }).(ServiceEntryTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ServiceEntryListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEntryList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ServiceEntryListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ServiceEntryList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ServiceEntryListArrayOutput struct{ *pulumi.OutputState }

func (ServiceEntryListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEntryList)(nil)).Elem()
}

func (o ServiceEntryListArrayOutput) ToServiceEntryListArrayOutput() ServiceEntryListArrayOutput {
	return o
}

func (o ServiceEntryListArrayOutput) ToServiceEntryListArrayOutputWithContext(ctx context.Context) ServiceEntryListArrayOutput {
	return o
}

func (o ServiceEntryListArrayOutput) Index(i pulumi.IntInput) ServiceEntryListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEntryList {
		return vs[0].([]*ServiceEntryList)[vs[1].(int)]
	}).(ServiceEntryListOutput)
}

type ServiceEntryListMapOutput struct{ *pulumi.OutputState }

func (ServiceEntryListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEntryList)(nil)).Elem()
}

func (o ServiceEntryListMapOutput) ToServiceEntryListMapOutput() ServiceEntryListMapOutput {
	return o
}

func (o ServiceEntryListMapOutput) ToServiceEntryListMapOutputWithContext(ctx context.Context) ServiceEntryListMapOutput {
	return o
}

func (o ServiceEntryListMapOutput) MapIndex(k pulumi.StringInput) ServiceEntryListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEntryList {
		return vs[0].(map[string]*ServiceEntryList)[vs[1].(string)]
	}).(ServiceEntryListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEntryListInput)(nil)).Elem(), &ServiceEntryList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEntryListArrayInput)(nil)).Elem(), ServiceEntryListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEntryListMapInput)(nil)).Elem(), ServiceEntryListMap{})
	pulumi.RegisterOutputType(ServiceEntryListOutput{})
	pulumi.RegisterOutputType(ServiceEntryListArrayOutput{})
	pulumi.RegisterOutputType(ServiceEntryListMapOutput{})
}
