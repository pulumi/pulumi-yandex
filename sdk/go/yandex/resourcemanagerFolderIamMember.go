// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows creation and management of a single member for a single binding within
// the IAM policy for an existing Yandex Resource Manager folder.
//
// > **Note:** This resource _must not_ be used in conjunction with
//    `ResourcemanagerFolderIamPolicy` or they will conflict over what your policy should be. Similarly, roles controlled by `ResourcemanagerFolderIamBinding`
//    should not be assigned using `ResourcemanagerFolderIamMember`.
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
// 		opt0 := "some_folder_id"
// 		_, err := yandex.LookupResourcemanagerFolder(ctx, &yandex.LookupResourcemanagerFolderArgs{
// 			FolderId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewResourcemanagerFolderIamMember(ctx, "admin", &yandex.ResourcemanagerFolderIamMemberArgs{
// 			FolderId: pulumi.Any(data.Yandex_resourcemanager.Department1.Name),
// 			Member:   pulumi.String("userAccount:user_id"),
// 			Role:     pulumi.String("editor"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the `folder id`, role, and account, e.g.
//
// ```sh
//  $ pulumi import yandex:index/resourcemanagerFolderIamMember:ResourcemanagerFolderIamMember my_project "folder_id viewer foo@example.com"
// ```
type ResourcemanagerFolderIamMember struct {
	pulumi.CustomResourceState

	// ID of the folder to attach a policy to.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be assigned.
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
}

// NewResourcemanagerFolderIamMember registers a new resource with the given unique name, arguments, and options.
func NewResourcemanagerFolderIamMember(ctx *pulumi.Context,
	name string, args *ResourcemanagerFolderIamMemberArgs, opts ...pulumi.ResourceOption) (*ResourcemanagerFolderIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ResourcemanagerFolderIamMember
	err := ctx.RegisterResource("yandex:index/resourcemanagerFolderIamMember:ResourcemanagerFolderIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcemanagerFolderIamMember gets an existing ResourcemanagerFolderIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcemanagerFolderIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcemanagerFolderIamMemberState, opts ...pulumi.ResourceOption) (*ResourcemanagerFolderIamMember, error) {
	var resource ResourcemanagerFolderIamMember
	err := ctx.ReadResource("yandex:index/resourcemanagerFolderIamMember:ResourcemanagerFolderIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcemanagerFolderIamMember resources.
type resourcemanagerFolderIamMemberState struct {
	// ID of the folder to attach a policy to.
	FolderId *string `pulumi:"folderId"`
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member *string `pulumi:"member"`
	// The role that should be assigned.
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
}

type ResourcemanagerFolderIamMemberState struct {
	// ID of the folder to attach a policy to.
	FolderId pulumi.StringPtrInput
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member pulumi.StringPtrInput
	// The role that should be assigned.
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
}

func (ResourcemanagerFolderIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcemanagerFolderIamMemberState)(nil)).Elem()
}

type resourcemanagerFolderIamMemberArgs struct {
	// ID of the folder to attach a policy to.
	FolderId string `pulumi:"folderId"`
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member string `pulumi:"member"`
	// The role that should be assigned.
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
}

// The set of arguments for constructing a ResourcemanagerFolderIamMember resource.
type ResourcemanagerFolderIamMemberArgs struct {
	// ID of the folder to attach a policy to.
	FolderId pulumi.StringInput
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member pulumi.StringInput
	// The role that should be assigned.
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
}

func (ResourcemanagerFolderIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcemanagerFolderIamMemberArgs)(nil)).Elem()
}

type ResourcemanagerFolderIamMemberInput interface {
	pulumi.Input

	ToResourcemanagerFolderIamMemberOutput() ResourcemanagerFolderIamMemberOutput
	ToResourcemanagerFolderIamMemberOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberOutput
}

func (*ResourcemanagerFolderIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcemanagerFolderIamMember)(nil))
}

func (i *ResourcemanagerFolderIamMember) ToResourcemanagerFolderIamMemberOutput() ResourcemanagerFolderIamMemberOutput {
	return i.ToResourcemanagerFolderIamMemberOutputWithContext(context.Background())
}

func (i *ResourcemanagerFolderIamMember) ToResourcemanagerFolderIamMemberOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderIamMemberOutput)
}

func (i *ResourcemanagerFolderIamMember) ToResourcemanagerFolderIamMemberPtrOutput() ResourcemanagerFolderIamMemberPtrOutput {
	return i.ToResourcemanagerFolderIamMemberPtrOutputWithContext(context.Background())
}

func (i *ResourcemanagerFolderIamMember) ToResourcemanagerFolderIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderIamMemberPtrOutput)
}

type ResourcemanagerFolderIamMemberPtrInput interface {
	pulumi.Input

	ToResourcemanagerFolderIamMemberPtrOutput() ResourcemanagerFolderIamMemberPtrOutput
	ToResourcemanagerFolderIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberPtrOutput
}

type resourcemanagerFolderIamMemberPtrType ResourcemanagerFolderIamMemberArgs

func (*resourcemanagerFolderIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcemanagerFolderIamMember)(nil))
}

func (i *resourcemanagerFolderIamMemberPtrType) ToResourcemanagerFolderIamMemberPtrOutput() ResourcemanagerFolderIamMemberPtrOutput {
	return i.ToResourcemanagerFolderIamMemberPtrOutputWithContext(context.Background())
}

func (i *resourcemanagerFolderIamMemberPtrType) ToResourcemanagerFolderIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderIamMemberPtrOutput)
}

// ResourcemanagerFolderIamMemberArrayInput is an input type that accepts ResourcemanagerFolderIamMemberArray and ResourcemanagerFolderIamMemberArrayOutput values.
// You can construct a concrete instance of `ResourcemanagerFolderIamMemberArrayInput` via:
//
//          ResourcemanagerFolderIamMemberArray{ ResourcemanagerFolderIamMemberArgs{...} }
type ResourcemanagerFolderIamMemberArrayInput interface {
	pulumi.Input

	ToResourcemanagerFolderIamMemberArrayOutput() ResourcemanagerFolderIamMemberArrayOutput
	ToResourcemanagerFolderIamMemberArrayOutputWithContext(context.Context) ResourcemanagerFolderIamMemberArrayOutput
}

type ResourcemanagerFolderIamMemberArray []ResourcemanagerFolderIamMemberInput

func (ResourcemanagerFolderIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ResourcemanagerFolderIamMember)(nil))
}

func (i ResourcemanagerFolderIamMemberArray) ToResourcemanagerFolderIamMemberArrayOutput() ResourcemanagerFolderIamMemberArrayOutput {
	return i.ToResourcemanagerFolderIamMemberArrayOutputWithContext(context.Background())
}

func (i ResourcemanagerFolderIamMemberArray) ToResourcemanagerFolderIamMemberArrayOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderIamMemberArrayOutput)
}

// ResourcemanagerFolderIamMemberMapInput is an input type that accepts ResourcemanagerFolderIamMemberMap and ResourcemanagerFolderIamMemberMapOutput values.
// You can construct a concrete instance of `ResourcemanagerFolderIamMemberMapInput` via:
//
//          ResourcemanagerFolderIamMemberMap{ "key": ResourcemanagerFolderIamMemberArgs{...} }
type ResourcemanagerFolderIamMemberMapInput interface {
	pulumi.Input

	ToResourcemanagerFolderIamMemberMapOutput() ResourcemanagerFolderIamMemberMapOutput
	ToResourcemanagerFolderIamMemberMapOutputWithContext(context.Context) ResourcemanagerFolderIamMemberMapOutput
}

type ResourcemanagerFolderIamMemberMap map[string]ResourcemanagerFolderIamMemberInput

func (ResourcemanagerFolderIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ResourcemanagerFolderIamMember)(nil))
}

func (i ResourcemanagerFolderIamMemberMap) ToResourcemanagerFolderIamMemberMapOutput() ResourcemanagerFolderIamMemberMapOutput {
	return i.ToResourcemanagerFolderIamMemberMapOutputWithContext(context.Background())
}

func (i ResourcemanagerFolderIamMemberMap) ToResourcemanagerFolderIamMemberMapOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderIamMemberMapOutput)
}

type ResourcemanagerFolderIamMemberOutput struct {
	*pulumi.OutputState
}

func (ResourcemanagerFolderIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcemanagerFolderIamMember)(nil))
}

func (o ResourcemanagerFolderIamMemberOutput) ToResourcemanagerFolderIamMemberOutput() ResourcemanagerFolderIamMemberOutput {
	return o
}

func (o ResourcemanagerFolderIamMemberOutput) ToResourcemanagerFolderIamMemberOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberOutput {
	return o
}

func (o ResourcemanagerFolderIamMemberOutput) ToResourcemanagerFolderIamMemberPtrOutput() ResourcemanagerFolderIamMemberPtrOutput {
	return o.ToResourcemanagerFolderIamMemberPtrOutputWithContext(context.Background())
}

func (o ResourcemanagerFolderIamMemberOutput) ToResourcemanagerFolderIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberPtrOutput {
	return o.ApplyT(func(v ResourcemanagerFolderIamMember) *ResourcemanagerFolderIamMember {
		return &v
	}).(ResourcemanagerFolderIamMemberPtrOutput)
}

type ResourcemanagerFolderIamMemberPtrOutput struct {
	*pulumi.OutputState
}

func (ResourcemanagerFolderIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcemanagerFolderIamMember)(nil))
}

func (o ResourcemanagerFolderIamMemberPtrOutput) ToResourcemanagerFolderIamMemberPtrOutput() ResourcemanagerFolderIamMemberPtrOutput {
	return o
}

func (o ResourcemanagerFolderIamMemberPtrOutput) ToResourcemanagerFolderIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberPtrOutput {
	return o
}

type ResourcemanagerFolderIamMemberArrayOutput struct{ *pulumi.OutputState }

func (ResourcemanagerFolderIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourcemanagerFolderIamMember)(nil))
}

func (o ResourcemanagerFolderIamMemberArrayOutput) ToResourcemanagerFolderIamMemberArrayOutput() ResourcemanagerFolderIamMemberArrayOutput {
	return o
}

func (o ResourcemanagerFolderIamMemberArrayOutput) ToResourcemanagerFolderIamMemberArrayOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberArrayOutput {
	return o
}

func (o ResourcemanagerFolderIamMemberArrayOutput) Index(i pulumi.IntInput) ResourcemanagerFolderIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourcemanagerFolderIamMember {
		return vs[0].([]ResourcemanagerFolderIamMember)[vs[1].(int)]
	}).(ResourcemanagerFolderIamMemberOutput)
}

type ResourcemanagerFolderIamMemberMapOutput struct{ *pulumi.OutputState }

func (ResourcemanagerFolderIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourcemanagerFolderIamMember)(nil))
}

func (o ResourcemanagerFolderIamMemberMapOutput) ToResourcemanagerFolderIamMemberMapOutput() ResourcemanagerFolderIamMemberMapOutput {
	return o
}

func (o ResourcemanagerFolderIamMemberMapOutput) ToResourcemanagerFolderIamMemberMapOutputWithContext(ctx context.Context) ResourcemanagerFolderIamMemberMapOutput {
	return o
}

func (o ResourcemanagerFolderIamMemberMapOutput) MapIndex(k pulumi.StringInput) ResourcemanagerFolderIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourcemanagerFolderIamMember {
		return vs[0].(map[string]ResourcemanagerFolderIamMember)[vs[1].(string)]
	}).(ResourcemanagerFolderIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourcemanagerFolderIamMemberOutput{})
	pulumi.RegisterOutputType(ResourcemanagerFolderIamMemberPtrOutput{})
	pulumi.RegisterOutputType(ResourcemanagerFolderIamMemberArrayOutput{})
	pulumi.RegisterOutputType(ResourcemanagerFolderIamMemberMapOutput{})
}
