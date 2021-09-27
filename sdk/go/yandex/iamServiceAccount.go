// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of a Yandex.Cloud IAM [service account](https://cloud.yandex.com/docs/iam/concepts/users/service-accounts).
// To assign roles and permissions, use the yandex_iam_service_account_iam_binding,
// IamServiceAccountIamMember and
// IamServiceAccountIamPolicy resources.
//
// ## Example Usage
//
// This snippet creates a service account.
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
// 		_, err := yandex.NewIamServiceAccount(ctx, "sa", &yandex.IamServiceAccountArgs{
// 			Description: pulumi.String("service account to manage VMs"),
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
// A service account can be imported using the `id` of the resource, e.g.
//
// ```sh
//  $ pulumi import yandex:index/iamServiceAccount:IamServiceAccount sa account_id
// ```
type IamServiceAccount struct {
	pulumi.CustomResourceState

	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the service account.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ID of the folder that the service account will be created in.
	// Defaults to the provider folder configuration.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Name of the service account.
	// Can be updated without creating a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIamServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewIamServiceAccount(ctx *pulumi.Context,
	name string, args *IamServiceAccountArgs, opts ...pulumi.ResourceOption) (*IamServiceAccount, error) {
	if args == nil {
		args = &IamServiceAccountArgs{}
	}

	var resource IamServiceAccount
	err := ctx.RegisterResource("yandex:index/iamServiceAccount:IamServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamServiceAccount gets an existing IamServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamServiceAccountState, opts ...pulumi.ResourceOption) (*IamServiceAccount, error) {
	var resource IamServiceAccount
	err := ctx.ReadResource("yandex:index/iamServiceAccount:IamServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamServiceAccount resources.
type iamServiceAccountState struct {
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the service account.
	Description *string `pulumi:"description"`
	// ID of the folder that the service account will be created in.
	// Defaults to the provider folder configuration.
	FolderId *string `pulumi:"folderId"`
	// Name of the service account.
	// Can be updated without creating a new resource.
	Name *string `pulumi:"name"`
}

type IamServiceAccountState struct {
	CreatedAt pulumi.StringPtrInput
	// Description of the service account.
	Description pulumi.StringPtrInput
	// ID of the folder that the service account will be created in.
	// Defaults to the provider folder configuration.
	FolderId pulumi.StringPtrInput
	// Name of the service account.
	// Can be updated without creating a new resource.
	Name pulumi.StringPtrInput
}

func (IamServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamServiceAccountState)(nil)).Elem()
}

type iamServiceAccountArgs struct {
	// Description of the service account.
	Description *string `pulumi:"description"`
	// ID of the folder that the service account will be created in.
	// Defaults to the provider folder configuration.
	FolderId *string `pulumi:"folderId"`
	// Name of the service account.
	// Can be updated without creating a new resource.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a IamServiceAccount resource.
type IamServiceAccountArgs struct {
	// Description of the service account.
	Description pulumi.StringPtrInput
	// ID of the folder that the service account will be created in.
	// Defaults to the provider folder configuration.
	FolderId pulumi.StringPtrInput
	// Name of the service account.
	// Can be updated without creating a new resource.
	Name pulumi.StringPtrInput
}

func (IamServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamServiceAccountArgs)(nil)).Elem()
}

type IamServiceAccountInput interface {
	pulumi.Input

	ToIamServiceAccountOutput() IamServiceAccountOutput
	ToIamServiceAccountOutputWithContext(ctx context.Context) IamServiceAccountOutput
}

func (*IamServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*IamServiceAccount)(nil))
}

func (i *IamServiceAccount) ToIamServiceAccountOutput() IamServiceAccountOutput {
	return i.ToIamServiceAccountOutputWithContext(context.Background())
}

func (i *IamServiceAccount) ToIamServiceAccountOutputWithContext(ctx context.Context) IamServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountOutput)
}

func (i *IamServiceAccount) ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput {
	return i.ToIamServiceAccountPtrOutputWithContext(context.Background())
}

func (i *IamServiceAccount) ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountPtrOutput)
}

type IamServiceAccountPtrInput interface {
	pulumi.Input

	ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput
	ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput
}

type iamServiceAccountPtrType IamServiceAccountArgs

func (*iamServiceAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IamServiceAccount)(nil))
}

func (i *iamServiceAccountPtrType) ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput {
	return i.ToIamServiceAccountPtrOutputWithContext(context.Background())
}

func (i *iamServiceAccountPtrType) ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountPtrOutput)
}

// IamServiceAccountArrayInput is an input type that accepts IamServiceAccountArray and IamServiceAccountArrayOutput values.
// You can construct a concrete instance of `IamServiceAccountArrayInput` via:
//
//          IamServiceAccountArray{ IamServiceAccountArgs{...} }
type IamServiceAccountArrayInput interface {
	pulumi.Input

	ToIamServiceAccountArrayOutput() IamServiceAccountArrayOutput
	ToIamServiceAccountArrayOutputWithContext(context.Context) IamServiceAccountArrayOutput
}

type IamServiceAccountArray []IamServiceAccountInput

func (IamServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IamServiceAccount)(nil))
}

func (i IamServiceAccountArray) ToIamServiceAccountArrayOutput() IamServiceAccountArrayOutput {
	return i.ToIamServiceAccountArrayOutputWithContext(context.Background())
}

func (i IamServiceAccountArray) ToIamServiceAccountArrayOutputWithContext(ctx context.Context) IamServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountArrayOutput)
}

// IamServiceAccountMapInput is an input type that accepts IamServiceAccountMap and IamServiceAccountMapOutput values.
// You can construct a concrete instance of `IamServiceAccountMapInput` via:
//
//          IamServiceAccountMap{ "key": IamServiceAccountArgs{...} }
type IamServiceAccountMapInput interface {
	pulumi.Input

	ToIamServiceAccountMapOutput() IamServiceAccountMapOutput
	ToIamServiceAccountMapOutputWithContext(context.Context) IamServiceAccountMapOutput
}

type IamServiceAccountMap map[string]IamServiceAccountInput

func (IamServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IamServiceAccount)(nil))
}

func (i IamServiceAccountMap) ToIamServiceAccountMapOutput() IamServiceAccountMapOutput {
	return i.ToIamServiceAccountMapOutputWithContext(context.Background())
}

func (i IamServiceAccountMap) ToIamServiceAccountMapOutputWithContext(ctx context.Context) IamServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountMapOutput)
}

type IamServiceAccountOutput struct {
	*pulumi.OutputState
}

func (IamServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamServiceAccount)(nil))
}

func (o IamServiceAccountOutput) ToIamServiceAccountOutput() IamServiceAccountOutput {
	return o
}

func (o IamServiceAccountOutput) ToIamServiceAccountOutputWithContext(ctx context.Context) IamServiceAccountOutput {
	return o
}

func (o IamServiceAccountOutput) ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput {
	return o.ToIamServiceAccountPtrOutputWithContext(context.Background())
}

func (o IamServiceAccountOutput) ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput {
	return o.ApplyT(func(v IamServiceAccount) *IamServiceAccount {
		return &v
	}).(IamServiceAccountPtrOutput)
}

type IamServiceAccountPtrOutput struct {
	*pulumi.OutputState
}

func (IamServiceAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamServiceAccount)(nil))
}

func (o IamServiceAccountPtrOutput) ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput {
	return o
}

func (o IamServiceAccountPtrOutput) ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput {
	return o
}

type IamServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (IamServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamServiceAccount)(nil))
}

func (o IamServiceAccountArrayOutput) ToIamServiceAccountArrayOutput() IamServiceAccountArrayOutput {
	return o
}

func (o IamServiceAccountArrayOutput) ToIamServiceAccountArrayOutputWithContext(ctx context.Context) IamServiceAccountArrayOutput {
	return o
}

func (o IamServiceAccountArrayOutput) Index(i pulumi.IntInput) IamServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IamServiceAccount {
		return vs[0].([]IamServiceAccount)[vs[1].(int)]
	}).(IamServiceAccountOutput)
}

type IamServiceAccountMapOutput struct{ *pulumi.OutputState }

func (IamServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamServiceAccount)(nil))
}

func (o IamServiceAccountMapOutput) ToIamServiceAccountMapOutput() IamServiceAccountMapOutput {
	return o
}

func (o IamServiceAccountMapOutput) ToIamServiceAccountMapOutputWithContext(ctx context.Context) IamServiceAccountMapOutput {
	return o
}

func (o IamServiceAccountMapOutput) MapIndex(k pulumi.StringInput) IamServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IamServiceAccount {
		return vs[0].(map[string]IamServiceAccount)[vs[1].(string)]
	}).(IamServiceAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(IamServiceAccountOutput{})
	pulumi.RegisterOutputType(IamServiceAccountPtrOutput{})
	pulumi.RegisterOutputType(IamServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(IamServiceAccountMapOutput{})
}
