// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows creation and management of a single binding within IAM policy for
// an existing Yandex.Cloud Organization Manager organization.
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
// 		_, err := yandex.NewOrganizationManagerOrganizationIamBinding(ctx, "editor", &yandex.OrganizationManagerOrganizationIamBindingArgs{
// 			Members: pulumi.StringArray{
// 				pulumi.String("userAccount:some_user_id"),
// 			},
// 			OrganizationId: pulumi.String("some_organization_id"),
// 			Role:           pulumi.String("editor"),
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
// IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the `organization_id` and role, e.g.
//
// ```sh
//  $ pulumi import yandex:index/organizationManagerOrganizationIamBinding:OrganizationManagerOrganizationIamBinding viewer "organization_id viewer"
// ```
type OrganizationManagerOrganizationIamBinding struct {
	pulumi.CustomResourceState

	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}**: A unique federated user ID.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// ID of the organization to attach the policy to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The role that should be assigned. Only one
	// `OrganizationManagerOrganizationIamBinding` can be used per role.
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
}

// NewOrganizationManagerOrganizationIamBinding registers a new resource with the given unique name, arguments, and options.
func NewOrganizationManagerOrganizationIamBinding(ctx *pulumi.Context,
	name string, args *OrganizationManagerOrganizationIamBindingArgs, opts ...pulumi.ResourceOption) (*OrganizationManagerOrganizationIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource OrganizationManagerOrganizationIamBinding
	err := ctx.RegisterResource("yandex:index/organizationManagerOrganizationIamBinding:OrganizationManagerOrganizationIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationManagerOrganizationIamBinding gets an existing OrganizationManagerOrganizationIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationManagerOrganizationIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationManagerOrganizationIamBindingState, opts ...pulumi.ResourceOption) (*OrganizationManagerOrganizationIamBinding, error) {
	var resource OrganizationManagerOrganizationIamBinding
	err := ctx.ReadResource("yandex:index/organizationManagerOrganizationIamBinding:OrganizationManagerOrganizationIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationManagerOrganizationIamBinding resources.
type organizationManagerOrganizationIamBindingState struct {
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}**: A unique federated user ID.
	Members []string `pulumi:"members"`
	// ID of the organization to attach the policy to.
	OrganizationId *string `pulumi:"organizationId"`
	// The role that should be assigned. Only one
	// `OrganizationManagerOrganizationIamBinding` can be used per role.
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
}

type OrganizationManagerOrganizationIamBindingState struct {
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}**: A unique federated user ID.
	Members pulumi.StringArrayInput
	// ID of the organization to attach the policy to.
	OrganizationId pulumi.StringPtrInput
	// The role that should be assigned. Only one
	// `OrganizationManagerOrganizationIamBinding` can be used per role.
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
}

func (OrganizationManagerOrganizationIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationManagerOrganizationIamBindingState)(nil)).Elem()
}

type organizationManagerOrganizationIamBindingArgs struct {
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}**: A unique federated user ID.
	Members []string `pulumi:"members"`
	// ID of the organization to attach the policy to.
	OrganizationId string `pulumi:"organizationId"`
	// The role that should be assigned. Only one
	// `OrganizationManagerOrganizationIamBinding` can be used per role.
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
}

// The set of arguments for constructing a OrganizationManagerOrganizationIamBinding resource.
type OrganizationManagerOrganizationIamBindingArgs struct {
	// An array of identities that will be granted the privilege in the `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **federatedUser:{federated_user_id}**: A unique federated user ID.
	Members pulumi.StringArrayInput
	// ID of the organization to attach the policy to.
	OrganizationId pulumi.StringInput
	// The role that should be assigned. Only one
	// `OrganizationManagerOrganizationIamBinding` can be used per role.
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
}

func (OrganizationManagerOrganizationIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationManagerOrganizationIamBindingArgs)(nil)).Elem()
}

type OrganizationManagerOrganizationIamBindingInput interface {
	pulumi.Input

	ToOrganizationManagerOrganizationIamBindingOutput() OrganizationManagerOrganizationIamBindingOutput
	ToOrganizationManagerOrganizationIamBindingOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingOutput
}

func (*OrganizationManagerOrganizationIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationManagerOrganizationIamBinding)(nil))
}

func (i *OrganizationManagerOrganizationIamBinding) ToOrganizationManagerOrganizationIamBindingOutput() OrganizationManagerOrganizationIamBindingOutput {
	return i.ToOrganizationManagerOrganizationIamBindingOutputWithContext(context.Background())
}

func (i *OrganizationManagerOrganizationIamBinding) ToOrganizationManagerOrganizationIamBindingOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagerOrganizationIamBindingOutput)
}

func (i *OrganizationManagerOrganizationIamBinding) ToOrganizationManagerOrganizationIamBindingPtrOutput() OrganizationManagerOrganizationIamBindingPtrOutput {
	return i.ToOrganizationManagerOrganizationIamBindingPtrOutputWithContext(context.Background())
}

func (i *OrganizationManagerOrganizationIamBinding) ToOrganizationManagerOrganizationIamBindingPtrOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagerOrganizationIamBindingPtrOutput)
}

type OrganizationManagerOrganizationIamBindingPtrInput interface {
	pulumi.Input

	ToOrganizationManagerOrganizationIamBindingPtrOutput() OrganizationManagerOrganizationIamBindingPtrOutput
	ToOrganizationManagerOrganizationIamBindingPtrOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingPtrOutput
}

type organizationManagerOrganizationIamBindingPtrType OrganizationManagerOrganizationIamBindingArgs

func (*organizationManagerOrganizationIamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationManagerOrganizationIamBinding)(nil))
}

func (i *organizationManagerOrganizationIamBindingPtrType) ToOrganizationManagerOrganizationIamBindingPtrOutput() OrganizationManagerOrganizationIamBindingPtrOutput {
	return i.ToOrganizationManagerOrganizationIamBindingPtrOutputWithContext(context.Background())
}

func (i *organizationManagerOrganizationIamBindingPtrType) ToOrganizationManagerOrganizationIamBindingPtrOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagerOrganizationIamBindingPtrOutput)
}

// OrganizationManagerOrganizationIamBindingArrayInput is an input type that accepts OrganizationManagerOrganizationIamBindingArray and OrganizationManagerOrganizationIamBindingArrayOutput values.
// You can construct a concrete instance of `OrganizationManagerOrganizationIamBindingArrayInput` via:
//
//          OrganizationManagerOrganizationIamBindingArray{ OrganizationManagerOrganizationIamBindingArgs{...} }
type OrganizationManagerOrganizationIamBindingArrayInput interface {
	pulumi.Input

	ToOrganizationManagerOrganizationIamBindingArrayOutput() OrganizationManagerOrganizationIamBindingArrayOutput
	ToOrganizationManagerOrganizationIamBindingArrayOutputWithContext(context.Context) OrganizationManagerOrganizationIamBindingArrayOutput
}

type OrganizationManagerOrganizationIamBindingArray []OrganizationManagerOrganizationIamBindingInput

func (OrganizationManagerOrganizationIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*OrganizationManagerOrganizationIamBinding)(nil))
}

func (i OrganizationManagerOrganizationIamBindingArray) ToOrganizationManagerOrganizationIamBindingArrayOutput() OrganizationManagerOrganizationIamBindingArrayOutput {
	return i.ToOrganizationManagerOrganizationIamBindingArrayOutputWithContext(context.Background())
}

func (i OrganizationManagerOrganizationIamBindingArray) ToOrganizationManagerOrganizationIamBindingArrayOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagerOrganizationIamBindingArrayOutput)
}

// OrganizationManagerOrganizationIamBindingMapInput is an input type that accepts OrganizationManagerOrganizationIamBindingMap and OrganizationManagerOrganizationIamBindingMapOutput values.
// You can construct a concrete instance of `OrganizationManagerOrganizationIamBindingMapInput` via:
//
//          OrganizationManagerOrganizationIamBindingMap{ "key": OrganizationManagerOrganizationIamBindingArgs{...} }
type OrganizationManagerOrganizationIamBindingMapInput interface {
	pulumi.Input

	ToOrganizationManagerOrganizationIamBindingMapOutput() OrganizationManagerOrganizationIamBindingMapOutput
	ToOrganizationManagerOrganizationIamBindingMapOutputWithContext(context.Context) OrganizationManagerOrganizationIamBindingMapOutput
}

type OrganizationManagerOrganizationIamBindingMap map[string]OrganizationManagerOrganizationIamBindingInput

func (OrganizationManagerOrganizationIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*OrganizationManagerOrganizationIamBinding)(nil))
}

func (i OrganizationManagerOrganizationIamBindingMap) ToOrganizationManagerOrganizationIamBindingMapOutput() OrganizationManagerOrganizationIamBindingMapOutput {
	return i.ToOrganizationManagerOrganizationIamBindingMapOutputWithContext(context.Background())
}

func (i OrganizationManagerOrganizationIamBindingMap) ToOrganizationManagerOrganizationIamBindingMapOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagerOrganizationIamBindingMapOutput)
}

type OrganizationManagerOrganizationIamBindingOutput struct {
	*pulumi.OutputState
}

func (OrganizationManagerOrganizationIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationManagerOrganizationIamBinding)(nil))
}

func (o OrganizationManagerOrganizationIamBindingOutput) ToOrganizationManagerOrganizationIamBindingOutput() OrganizationManagerOrganizationIamBindingOutput {
	return o
}

func (o OrganizationManagerOrganizationIamBindingOutput) ToOrganizationManagerOrganizationIamBindingOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingOutput {
	return o
}

func (o OrganizationManagerOrganizationIamBindingOutput) ToOrganizationManagerOrganizationIamBindingPtrOutput() OrganizationManagerOrganizationIamBindingPtrOutput {
	return o.ToOrganizationManagerOrganizationIamBindingPtrOutputWithContext(context.Background())
}

func (o OrganizationManagerOrganizationIamBindingOutput) ToOrganizationManagerOrganizationIamBindingPtrOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingPtrOutput {
	return o.ApplyT(func(v OrganizationManagerOrganizationIamBinding) *OrganizationManagerOrganizationIamBinding {
		return &v
	}).(OrganizationManagerOrganizationIamBindingPtrOutput)
}

type OrganizationManagerOrganizationIamBindingPtrOutput struct {
	*pulumi.OutputState
}

func (OrganizationManagerOrganizationIamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationManagerOrganizationIamBinding)(nil))
}

func (o OrganizationManagerOrganizationIamBindingPtrOutput) ToOrganizationManagerOrganizationIamBindingPtrOutput() OrganizationManagerOrganizationIamBindingPtrOutput {
	return o
}

func (o OrganizationManagerOrganizationIamBindingPtrOutput) ToOrganizationManagerOrganizationIamBindingPtrOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingPtrOutput {
	return o
}

type OrganizationManagerOrganizationIamBindingArrayOutput struct{ *pulumi.OutputState }

func (OrganizationManagerOrganizationIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrganizationManagerOrganizationIamBinding)(nil))
}

func (o OrganizationManagerOrganizationIamBindingArrayOutput) ToOrganizationManagerOrganizationIamBindingArrayOutput() OrganizationManagerOrganizationIamBindingArrayOutput {
	return o
}

func (o OrganizationManagerOrganizationIamBindingArrayOutput) ToOrganizationManagerOrganizationIamBindingArrayOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingArrayOutput {
	return o
}

func (o OrganizationManagerOrganizationIamBindingArrayOutput) Index(i pulumi.IntInput) OrganizationManagerOrganizationIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OrganizationManagerOrganizationIamBinding {
		return vs[0].([]OrganizationManagerOrganizationIamBinding)[vs[1].(int)]
	}).(OrganizationManagerOrganizationIamBindingOutput)
}

type OrganizationManagerOrganizationIamBindingMapOutput struct{ *pulumi.OutputState }

func (OrganizationManagerOrganizationIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OrganizationManagerOrganizationIamBinding)(nil))
}

func (o OrganizationManagerOrganizationIamBindingMapOutput) ToOrganizationManagerOrganizationIamBindingMapOutput() OrganizationManagerOrganizationIamBindingMapOutput {
	return o
}

func (o OrganizationManagerOrganizationIamBindingMapOutput) ToOrganizationManagerOrganizationIamBindingMapOutputWithContext(ctx context.Context) OrganizationManagerOrganizationIamBindingMapOutput {
	return o
}

func (o OrganizationManagerOrganizationIamBindingMapOutput) MapIndex(k pulumi.StringInput) OrganizationManagerOrganizationIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OrganizationManagerOrganizationIamBinding {
		return vs[0].(map[string]OrganizationManagerOrganizationIamBinding)[vs[1].(string)]
	}).(OrganizationManagerOrganizationIamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationManagerOrganizationIamBindingOutput{})
	pulumi.RegisterOutputType(OrganizationManagerOrganizationIamBindingPtrOutput{})
	pulumi.RegisterOutputType(OrganizationManagerOrganizationIamBindingArrayOutput{})
	pulumi.RegisterOutputType(OrganizationManagerOrganizationIamBindingMapOutput{})
}