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
// the IAM policy for an existing Yandex Resource Manager cloud.
//
// > **Note:** Roles controlled by `ResourcemanagerCloudIamBinding`
//    should not be assigned using `ResourcemanagerCloudIamMember`.
//
// > **Note:** When you delete `ResourcemanagerCloudIamBinding` resource,
//    the roles can be deleted from other users within the cloud as well. Be careful!
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
// 		opt0 := "Department 1"
// 		department1, err := yandex.GetResourcemanagerCloud(ctx, &GetResourcemanagerCloudArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewResourcemanagerCloudIamMember(ctx, "admin", &yandex.ResourcemanagerCloudIamMemberArgs{
// 			CloudId: pulumi.String(department1.Id),
// 			Member:  pulumi.String("userAccount:user_id"),
// 			Role:    pulumi.String("editor"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the `cloud id`, role, and account, e.g.
//
// ```sh
//  $ pulumi import yandex:index/resourcemanagerCloudIamMember:ResourcemanagerCloudIamMember my_project "cloud_id viewer foo@example.com"
// ```
type ResourcemanagerCloudIamMember struct {
	pulumi.CustomResourceState

	// ID of the cloud to attach a policy to.
	CloudId pulumi.StringOutput `pulumi:"cloudId"`
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be assigned.
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
}

// NewResourcemanagerCloudIamMember registers a new resource with the given unique name, arguments, and options.
func NewResourcemanagerCloudIamMember(ctx *pulumi.Context,
	name string, args *ResourcemanagerCloudIamMemberArgs, opts ...pulumi.ResourceOption) (*ResourcemanagerCloudIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudId == nil {
		return nil, errors.New("invalid value for required argument 'CloudId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ResourcemanagerCloudIamMember
	err := ctx.RegisterResource("yandex:index/resourcemanagerCloudIamMember:ResourcemanagerCloudIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcemanagerCloudIamMember gets an existing ResourcemanagerCloudIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcemanagerCloudIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcemanagerCloudIamMemberState, opts ...pulumi.ResourceOption) (*ResourcemanagerCloudIamMember, error) {
	var resource ResourcemanagerCloudIamMember
	err := ctx.ReadResource("yandex:index/resourcemanagerCloudIamMember:ResourcemanagerCloudIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcemanagerCloudIamMember resources.
type resourcemanagerCloudIamMemberState struct {
	// ID of the cloud to attach a policy to.
	CloudId *string `pulumi:"cloudId"`
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member *string `pulumi:"member"`
	// The role that should be assigned.
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
}

type ResourcemanagerCloudIamMemberState struct {
	// ID of the cloud to attach a policy to.
	CloudId pulumi.StringPtrInput
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member pulumi.StringPtrInput
	// The role that should be assigned.
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
}

func (ResourcemanagerCloudIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcemanagerCloudIamMemberState)(nil)).Elem()
}

type resourcemanagerCloudIamMemberArgs struct {
	// ID of the cloud to attach a policy to.
	CloudId string `pulumi:"cloudId"`
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member string `pulumi:"member"`
	// The role that should be assigned.
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
}

// The set of arguments for constructing a ResourcemanagerCloudIamMember resource.
type ResourcemanagerCloudIamMemberArgs struct {
	// ID of the cloud to attach a policy to.
	CloudId pulumi.StringInput
	// The identity that will be granted the privilege that is specified in the `role` field.
	// This field can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	Member pulumi.StringInput
	// The role that should be assigned.
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
}

func (ResourcemanagerCloudIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcemanagerCloudIamMemberArgs)(nil)).Elem()
}

type ResourcemanagerCloudIamMemberInput interface {
	pulumi.Input

	ToResourcemanagerCloudIamMemberOutput() ResourcemanagerCloudIamMemberOutput
	ToResourcemanagerCloudIamMemberOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberOutput
}

func (*ResourcemanagerCloudIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcemanagerCloudIamMember)(nil))
}

func (i *ResourcemanagerCloudIamMember) ToResourcemanagerCloudIamMemberOutput() ResourcemanagerCloudIamMemberOutput {
	return i.ToResourcemanagerCloudIamMemberOutputWithContext(context.Background())
}

func (i *ResourcemanagerCloudIamMember) ToResourcemanagerCloudIamMemberOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerCloudIamMemberOutput)
}

func (i *ResourcemanagerCloudIamMember) ToResourcemanagerCloudIamMemberPtrOutput() ResourcemanagerCloudIamMemberPtrOutput {
	return i.ToResourcemanagerCloudIamMemberPtrOutputWithContext(context.Background())
}

func (i *ResourcemanagerCloudIamMember) ToResourcemanagerCloudIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerCloudIamMemberPtrOutput)
}

type ResourcemanagerCloudIamMemberPtrInput interface {
	pulumi.Input

	ToResourcemanagerCloudIamMemberPtrOutput() ResourcemanagerCloudIamMemberPtrOutput
	ToResourcemanagerCloudIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberPtrOutput
}

type resourcemanagerCloudIamMemberPtrType ResourcemanagerCloudIamMemberArgs

func (*resourcemanagerCloudIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcemanagerCloudIamMember)(nil))
}

func (i *resourcemanagerCloudIamMemberPtrType) ToResourcemanagerCloudIamMemberPtrOutput() ResourcemanagerCloudIamMemberPtrOutput {
	return i.ToResourcemanagerCloudIamMemberPtrOutputWithContext(context.Background())
}

func (i *resourcemanagerCloudIamMemberPtrType) ToResourcemanagerCloudIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerCloudIamMemberPtrOutput)
}

// ResourcemanagerCloudIamMemberArrayInput is an input type that accepts ResourcemanagerCloudIamMemberArray and ResourcemanagerCloudIamMemberArrayOutput values.
// You can construct a concrete instance of `ResourcemanagerCloudIamMemberArrayInput` via:
//
//          ResourcemanagerCloudIamMemberArray{ ResourcemanagerCloudIamMemberArgs{...} }
type ResourcemanagerCloudIamMemberArrayInput interface {
	pulumi.Input

	ToResourcemanagerCloudIamMemberArrayOutput() ResourcemanagerCloudIamMemberArrayOutput
	ToResourcemanagerCloudIamMemberArrayOutputWithContext(context.Context) ResourcemanagerCloudIamMemberArrayOutput
}

type ResourcemanagerCloudIamMemberArray []ResourcemanagerCloudIamMemberInput

func (ResourcemanagerCloudIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcemanagerCloudIamMember)(nil)).Elem()
}

func (i ResourcemanagerCloudIamMemberArray) ToResourcemanagerCloudIamMemberArrayOutput() ResourcemanagerCloudIamMemberArrayOutput {
	return i.ToResourcemanagerCloudIamMemberArrayOutputWithContext(context.Background())
}

func (i ResourcemanagerCloudIamMemberArray) ToResourcemanagerCloudIamMemberArrayOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerCloudIamMemberArrayOutput)
}

// ResourcemanagerCloudIamMemberMapInput is an input type that accepts ResourcemanagerCloudIamMemberMap and ResourcemanagerCloudIamMemberMapOutput values.
// You can construct a concrete instance of `ResourcemanagerCloudIamMemberMapInput` via:
//
//          ResourcemanagerCloudIamMemberMap{ "key": ResourcemanagerCloudIamMemberArgs{...} }
type ResourcemanagerCloudIamMemberMapInput interface {
	pulumi.Input

	ToResourcemanagerCloudIamMemberMapOutput() ResourcemanagerCloudIamMemberMapOutput
	ToResourcemanagerCloudIamMemberMapOutputWithContext(context.Context) ResourcemanagerCloudIamMemberMapOutput
}

type ResourcemanagerCloudIamMemberMap map[string]ResourcemanagerCloudIamMemberInput

func (ResourcemanagerCloudIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcemanagerCloudIamMember)(nil)).Elem()
}

func (i ResourcemanagerCloudIamMemberMap) ToResourcemanagerCloudIamMemberMapOutput() ResourcemanagerCloudIamMemberMapOutput {
	return i.ToResourcemanagerCloudIamMemberMapOutputWithContext(context.Background())
}

func (i ResourcemanagerCloudIamMemberMap) ToResourcemanagerCloudIamMemberMapOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerCloudIamMemberMapOutput)
}

type ResourcemanagerCloudIamMemberOutput struct{ *pulumi.OutputState }

func (ResourcemanagerCloudIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcemanagerCloudIamMember)(nil))
}

func (o ResourcemanagerCloudIamMemberOutput) ToResourcemanagerCloudIamMemberOutput() ResourcemanagerCloudIamMemberOutput {
	return o
}

func (o ResourcemanagerCloudIamMemberOutput) ToResourcemanagerCloudIamMemberOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberOutput {
	return o
}

func (o ResourcemanagerCloudIamMemberOutput) ToResourcemanagerCloudIamMemberPtrOutput() ResourcemanagerCloudIamMemberPtrOutput {
	return o.ToResourcemanagerCloudIamMemberPtrOutputWithContext(context.Background())
}

func (o ResourcemanagerCloudIamMemberOutput) ToResourcemanagerCloudIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourcemanagerCloudIamMember) *ResourcemanagerCloudIamMember {
		return &v
	}).(ResourcemanagerCloudIamMemberPtrOutput)
}

type ResourcemanagerCloudIamMemberPtrOutput struct{ *pulumi.OutputState }

func (ResourcemanagerCloudIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcemanagerCloudIamMember)(nil))
}

func (o ResourcemanagerCloudIamMemberPtrOutput) ToResourcemanagerCloudIamMemberPtrOutput() ResourcemanagerCloudIamMemberPtrOutput {
	return o
}

func (o ResourcemanagerCloudIamMemberPtrOutput) ToResourcemanagerCloudIamMemberPtrOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberPtrOutput {
	return o
}

func (o ResourcemanagerCloudIamMemberPtrOutput) Elem() ResourcemanagerCloudIamMemberOutput {
	return o.ApplyT(func(v *ResourcemanagerCloudIamMember) ResourcemanagerCloudIamMember {
		if v != nil {
			return *v
		}
		var ret ResourcemanagerCloudIamMember
		return ret
	}).(ResourcemanagerCloudIamMemberOutput)
}

type ResourcemanagerCloudIamMemberArrayOutput struct{ *pulumi.OutputState }

func (ResourcemanagerCloudIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourcemanagerCloudIamMember)(nil))
}

func (o ResourcemanagerCloudIamMemberArrayOutput) ToResourcemanagerCloudIamMemberArrayOutput() ResourcemanagerCloudIamMemberArrayOutput {
	return o
}

func (o ResourcemanagerCloudIamMemberArrayOutput) ToResourcemanagerCloudIamMemberArrayOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberArrayOutput {
	return o
}

func (o ResourcemanagerCloudIamMemberArrayOutput) Index(i pulumi.IntInput) ResourcemanagerCloudIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourcemanagerCloudIamMember {
		return vs[0].([]ResourcemanagerCloudIamMember)[vs[1].(int)]
	}).(ResourcemanagerCloudIamMemberOutput)
}

type ResourcemanagerCloudIamMemberMapOutput struct{ *pulumi.OutputState }

func (ResourcemanagerCloudIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourcemanagerCloudIamMember)(nil))
}

func (o ResourcemanagerCloudIamMemberMapOutput) ToResourcemanagerCloudIamMemberMapOutput() ResourcemanagerCloudIamMemberMapOutput {
	return o
}

func (o ResourcemanagerCloudIamMemberMapOutput) ToResourcemanagerCloudIamMemberMapOutputWithContext(ctx context.Context) ResourcemanagerCloudIamMemberMapOutput {
	return o
}

func (o ResourcemanagerCloudIamMemberMapOutput) MapIndex(k pulumi.StringInput) ResourcemanagerCloudIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourcemanagerCloudIamMember {
		return vs[0].(map[string]ResourcemanagerCloudIamMember)[vs[1].(string)]
	}).(ResourcemanagerCloudIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerCloudIamMemberInput)(nil)).Elem(), &ResourcemanagerCloudIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerCloudIamMemberPtrInput)(nil)).Elem(), &ResourcemanagerCloudIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerCloudIamMemberArrayInput)(nil)).Elem(), ResourcemanagerCloudIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerCloudIamMemberMapInput)(nil)).Elem(), ResourcemanagerCloudIamMemberMap{})
	pulumi.RegisterOutputType(ResourcemanagerCloudIamMemberOutput{})
	pulumi.RegisterOutputType(ResourcemanagerCloudIamMemberPtrOutput{})
	pulumi.RegisterOutputType(ResourcemanagerCloudIamMemberArrayOutput{})
	pulumi.RegisterOutputType(ResourcemanagerCloudIamMemberMapOutput{})
}
