// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows creation and management of Cloud Folders for an existing Yandex Cloud. See [the official documentation](https://cloud.yandex.com/docs/resource-manager/concepts/resources-hierarchy) for additional info.
// Note: deletion of folders may take up to 30 minutes as it requires a lot of communication between cloud services.
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
// 		_, err := yandex.NewResourcemanagerFolder(ctx, "folder1", &yandex.ResourcemanagerFolderArgs{
// 			CloudId: pulumi.String("my_cloud_id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ResourcemanagerFolder struct {
	pulumi.CustomResourceState

	// Cloud that the resource belongs to. If value is omitted, the default provider Cloud ID is used.
	CloudId   pulumi.StringOutput `pulumi:"cloudId"`
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A description of the Folder.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A set of key/value label pairs to assign to the Folder.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the Folder.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewResourcemanagerFolder registers a new resource with the given unique name, arguments, and options.
func NewResourcemanagerFolder(ctx *pulumi.Context,
	name string, args *ResourcemanagerFolderArgs, opts ...pulumi.ResourceOption) (*ResourcemanagerFolder, error) {
	if args == nil {
		args = &ResourcemanagerFolderArgs{}
	}

	var resource ResourcemanagerFolder
	err := ctx.RegisterResource("yandex:index/resourcemanagerFolder:ResourcemanagerFolder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcemanagerFolder gets an existing ResourcemanagerFolder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcemanagerFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcemanagerFolderState, opts ...pulumi.ResourceOption) (*ResourcemanagerFolder, error) {
	var resource ResourcemanagerFolder
	err := ctx.ReadResource("yandex:index/resourcemanagerFolder:ResourcemanagerFolder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcemanagerFolder resources.
type resourcemanagerFolderState struct {
	// Cloud that the resource belongs to. If value is omitted, the default provider Cloud ID is used.
	CloudId   *string `pulumi:"cloudId"`
	CreatedAt *string `pulumi:"createdAt"`
	// A description of the Folder.
	Description *string `pulumi:"description"`
	// A set of key/value label pairs to assign to the Folder.
	Labels map[string]string `pulumi:"labels"`
	// The name of the Folder.
	Name *string `pulumi:"name"`
}

type ResourcemanagerFolderState struct {
	// Cloud that the resource belongs to. If value is omitted, the default provider Cloud ID is used.
	CloudId   pulumi.StringPtrInput
	CreatedAt pulumi.StringPtrInput
	// A description of the Folder.
	Description pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Folder.
	Labels pulumi.StringMapInput
	// The name of the Folder.
	Name pulumi.StringPtrInput
}

func (ResourcemanagerFolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcemanagerFolderState)(nil)).Elem()
}

type resourcemanagerFolderArgs struct {
	// Cloud that the resource belongs to. If value is omitted, the default provider Cloud ID is used.
	CloudId *string `pulumi:"cloudId"`
	// A description of the Folder.
	Description *string `pulumi:"description"`
	// A set of key/value label pairs to assign to the Folder.
	Labels map[string]string `pulumi:"labels"`
	// The name of the Folder.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ResourcemanagerFolder resource.
type ResourcemanagerFolderArgs struct {
	// Cloud that the resource belongs to. If value is omitted, the default provider Cloud ID is used.
	CloudId pulumi.StringPtrInput
	// A description of the Folder.
	Description pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Folder.
	Labels pulumi.StringMapInput
	// The name of the Folder.
	Name pulumi.StringPtrInput
}

func (ResourcemanagerFolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcemanagerFolderArgs)(nil)).Elem()
}

type ResourcemanagerFolderInput interface {
	pulumi.Input

	ToResourcemanagerFolderOutput() ResourcemanagerFolderOutput
	ToResourcemanagerFolderOutputWithContext(ctx context.Context) ResourcemanagerFolderOutput
}

func (*ResourcemanagerFolder) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcemanagerFolder)(nil)).Elem()
}

func (i *ResourcemanagerFolder) ToResourcemanagerFolderOutput() ResourcemanagerFolderOutput {
	return i.ToResourcemanagerFolderOutputWithContext(context.Background())
}

func (i *ResourcemanagerFolder) ToResourcemanagerFolderOutputWithContext(ctx context.Context) ResourcemanagerFolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderOutput)
}

// ResourcemanagerFolderArrayInput is an input type that accepts ResourcemanagerFolderArray and ResourcemanagerFolderArrayOutput values.
// You can construct a concrete instance of `ResourcemanagerFolderArrayInput` via:
//
//          ResourcemanagerFolderArray{ ResourcemanagerFolderArgs{...} }
type ResourcemanagerFolderArrayInput interface {
	pulumi.Input

	ToResourcemanagerFolderArrayOutput() ResourcemanagerFolderArrayOutput
	ToResourcemanagerFolderArrayOutputWithContext(context.Context) ResourcemanagerFolderArrayOutput
}

type ResourcemanagerFolderArray []ResourcemanagerFolderInput

func (ResourcemanagerFolderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcemanagerFolder)(nil)).Elem()
}

func (i ResourcemanagerFolderArray) ToResourcemanagerFolderArrayOutput() ResourcemanagerFolderArrayOutput {
	return i.ToResourcemanagerFolderArrayOutputWithContext(context.Background())
}

func (i ResourcemanagerFolderArray) ToResourcemanagerFolderArrayOutputWithContext(ctx context.Context) ResourcemanagerFolderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderArrayOutput)
}

// ResourcemanagerFolderMapInput is an input type that accepts ResourcemanagerFolderMap and ResourcemanagerFolderMapOutput values.
// You can construct a concrete instance of `ResourcemanagerFolderMapInput` via:
//
//          ResourcemanagerFolderMap{ "key": ResourcemanagerFolderArgs{...} }
type ResourcemanagerFolderMapInput interface {
	pulumi.Input

	ToResourcemanagerFolderMapOutput() ResourcemanagerFolderMapOutput
	ToResourcemanagerFolderMapOutputWithContext(context.Context) ResourcemanagerFolderMapOutput
}

type ResourcemanagerFolderMap map[string]ResourcemanagerFolderInput

func (ResourcemanagerFolderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcemanagerFolder)(nil)).Elem()
}

func (i ResourcemanagerFolderMap) ToResourcemanagerFolderMapOutput() ResourcemanagerFolderMapOutput {
	return i.ToResourcemanagerFolderMapOutputWithContext(context.Background())
}

func (i ResourcemanagerFolderMap) ToResourcemanagerFolderMapOutputWithContext(ctx context.Context) ResourcemanagerFolderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderMapOutput)
}

type ResourcemanagerFolderOutput struct{ *pulumi.OutputState }

func (ResourcemanagerFolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcemanagerFolder)(nil)).Elem()
}

func (o ResourcemanagerFolderOutput) ToResourcemanagerFolderOutput() ResourcemanagerFolderOutput {
	return o
}

func (o ResourcemanagerFolderOutput) ToResourcemanagerFolderOutputWithContext(ctx context.Context) ResourcemanagerFolderOutput {
	return o
}

type ResourcemanagerFolderArrayOutput struct{ *pulumi.OutputState }

func (ResourcemanagerFolderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcemanagerFolder)(nil)).Elem()
}

func (o ResourcemanagerFolderArrayOutput) ToResourcemanagerFolderArrayOutput() ResourcemanagerFolderArrayOutput {
	return o
}

func (o ResourcemanagerFolderArrayOutput) ToResourcemanagerFolderArrayOutputWithContext(ctx context.Context) ResourcemanagerFolderArrayOutput {
	return o
}

func (o ResourcemanagerFolderArrayOutput) Index(i pulumi.IntInput) ResourcemanagerFolderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourcemanagerFolder {
		return vs[0].([]*ResourcemanagerFolder)[vs[1].(int)]
	}).(ResourcemanagerFolderOutput)
}

type ResourcemanagerFolderMapOutput struct{ *pulumi.OutputState }

func (ResourcemanagerFolderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcemanagerFolder)(nil)).Elem()
}

func (o ResourcemanagerFolderMapOutput) ToResourcemanagerFolderMapOutput() ResourcemanagerFolderMapOutput {
	return o
}

func (o ResourcemanagerFolderMapOutput) ToResourcemanagerFolderMapOutputWithContext(ctx context.Context) ResourcemanagerFolderMapOutput {
	return o
}

func (o ResourcemanagerFolderMapOutput) MapIndex(k pulumi.StringInput) ResourcemanagerFolderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourcemanagerFolder {
		return vs[0].(map[string]*ResourcemanagerFolder)[vs[1].(string)]
	}).(ResourcemanagerFolderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerFolderInput)(nil)).Elem(), &ResourcemanagerFolder{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerFolderArrayInput)(nil)).Elem(), ResourcemanagerFolderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerFolderMapInput)(nil)).Elem(), ResourcemanagerFolderMap{})
	pulumi.RegisterOutputType(ResourcemanagerFolderOutput{})
	pulumi.RegisterOutputType(ResourcemanagerFolderArrayOutput{})
	pulumi.RegisterOutputType(ResourcemanagerFolderMapOutput{})
}
