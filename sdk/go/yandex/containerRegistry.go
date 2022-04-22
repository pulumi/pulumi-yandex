// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new container registry. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/container-registry/concepts/registry)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := yandex.NewContainerRegistry(ctx, "default", &yandex.ContainerRegistryArgs{
// 			FolderId: pulumi.String("test_folder_id"),
// 			Labels: pulumi.StringMap{
// 				"my-label": pulumi.String("my-label-value"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// A registry can be imported using the `id` of the resource, e.g.
//
// ```sh
//  $ pulumi import yandex:index/containerRegistry:ContainerRegistry default registry_id
// ```
type ContainerRegistry struct {
	pulumi.CustomResourceState

	// Creation timestamp of the registry.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the registry.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A name of the registry.
	Name pulumi.StringOutput `pulumi:"name"`
	// Status of the registry.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewContainerRegistry registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistry(ctx *pulumi.Context,
	name string, args *ContainerRegistryArgs, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	if args == nil {
		args = &ContainerRegistryArgs{}
	}

	var resource ContainerRegistry
	err := ctx.RegisterResource("yandex:index/containerRegistry:ContainerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistry gets an existing ContainerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryState, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	var resource ContainerRegistry
	err := ctx.ReadResource("yandex:index/containerRegistry:ContainerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistry resources.
type containerRegistryState struct {
	// Creation timestamp of the registry.
	CreatedAt *string `pulumi:"createdAt"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the registry.
	Labels map[string]string `pulumi:"labels"`
	// A name of the registry.
	Name *string `pulumi:"name"`
	// Status of the registry.
	Status *string `pulumi:"status"`
}

type ContainerRegistryState struct {
	// Creation timestamp of the registry.
	CreatedAt pulumi.StringPtrInput
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the registry.
	Labels pulumi.StringMapInput
	// A name of the registry.
	Name pulumi.StringPtrInput
	// Status of the registry.
	Status pulumi.StringPtrInput
}

func (ContainerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryState)(nil)).Elem()
}

type containerRegistryArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the registry.
	Labels map[string]string `pulumi:"labels"`
	// A name of the registry.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ContainerRegistry resource.
type ContainerRegistryArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the registry.
	Labels pulumi.StringMapInput
	// A name of the registry.
	Name pulumi.StringPtrInput
}

func (ContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryArgs)(nil)).Elem()
}

type ContainerRegistryInput interface {
	pulumi.Input

	ToContainerRegistryOutput() ContainerRegistryOutput
	ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput
}

func (*ContainerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (i *ContainerRegistry) ToContainerRegistryOutput() ContainerRegistryOutput {
	return i.ToContainerRegistryOutputWithContext(context.Background())
}

func (i *ContainerRegistry) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOutput)
}

// ContainerRegistryArrayInput is an input type that accepts ContainerRegistryArray and ContainerRegistryArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryArrayInput` via:
//
//          ContainerRegistryArray{ ContainerRegistryArgs{...} }
type ContainerRegistryArrayInput interface {
	pulumi.Input

	ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput
	ToContainerRegistryArrayOutputWithContext(context.Context) ContainerRegistryArrayOutput
}

type ContainerRegistryArray []ContainerRegistryInput

func (ContainerRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistry)(nil)).Elem()
}

func (i ContainerRegistryArray) ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput {
	return i.ToContainerRegistryArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryArray) ToContainerRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryArrayOutput)
}

// ContainerRegistryMapInput is an input type that accepts ContainerRegistryMap and ContainerRegistryMapOutput values.
// You can construct a concrete instance of `ContainerRegistryMapInput` via:
//
//          ContainerRegistryMap{ "key": ContainerRegistryArgs{...} }
type ContainerRegistryMapInput interface {
	pulumi.Input

	ToContainerRegistryMapOutput() ContainerRegistryMapOutput
	ToContainerRegistryMapOutputWithContext(context.Context) ContainerRegistryMapOutput
}

type ContainerRegistryMap map[string]ContainerRegistryInput

func (ContainerRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistry)(nil)).Elem()
}

func (i ContainerRegistryMap) ToContainerRegistryMapOutput() ContainerRegistryMapOutput {
	return i.ToContainerRegistryMapOutputWithContext(context.Background())
}

func (i ContainerRegistryMap) ToContainerRegistryMapOutputWithContext(ctx context.Context) ContainerRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryMapOutput)
}

type ContainerRegistryOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryOutput) ToContainerRegistryOutput() ContainerRegistryOutput {
	return o
}

func (o ContainerRegistryOutput) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return o
}

type ContainerRegistryArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryArrayOutput) ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput {
	return o
}

func (o ContainerRegistryArrayOutput) ToContainerRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryArrayOutput {
	return o
}

func (o ContainerRegistryArrayOutput) Index(i pulumi.IntInput) ContainerRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRegistry {
		return vs[0].([]*ContainerRegistry)[vs[1].(int)]
	}).(ContainerRegistryOutput)
}

type ContainerRegistryMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryMapOutput) ToContainerRegistryMapOutput() ContainerRegistryMapOutput {
	return o
}

func (o ContainerRegistryMapOutput) ToContainerRegistryMapOutputWithContext(ctx context.Context) ContainerRegistryMapOutput {
	return o
}

func (o ContainerRegistryMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRegistry {
		return vs[0].(map[string]*ContainerRegistry)[vs[1].(string)]
	}).(ContainerRegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryInput)(nil)).Elem(), &ContainerRegistry{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryArrayInput)(nil)).Elem(), ContainerRegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryMapInput)(nil)).Elem(), ContainerRegistryMap{})
	pulumi.RegisterOutputType(ContainerRegistryOutput{})
	pulumi.RegisterOutputType(ContainerRegistryArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryMapOutput{})
}
