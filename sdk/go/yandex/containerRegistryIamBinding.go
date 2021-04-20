// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## yandex\_container\_registry\_iam\_binding
//
// Allows creation and management of a single binding within IAM policy for
// an existing Yandex Container Registry.
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
// 		_, err := yandex.NewContainerRegistry(ctx, "your_registry", &yandex.ContainerRegistryArgs{
// 			FolderId: pulumi.String("your-folder-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewContainerRegistryIamBinding(ctx, "puller", &yandex.ContainerRegistryIamBindingArgs{
// 			RegistryId: your_registry.ID(),
// 			Role:       pulumi.String("container-registry.images.puller"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("system:allUsers"),
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
// IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the `registry_id` and role, e.g.
//
// ```sh
//  $ pulumi import yandex:index/containerRegistryIamBinding:ContainerRegistryIamBinding puller "registry_id container-registry.images.puller"
// ```
type ContainerRegistryIamBinding struct {
	pulumi.CustomResourceState

	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The [Yandex Container Registry](https://cloud.yandex.com/docs/container-registry/) ID to apply a binding to.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/container-registry/security/).
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
}

// NewContainerRegistryIamBinding registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistryIamBinding(ctx *pulumi.Context,
	name string, args *ContainerRegistryIamBindingArgs, opts ...pulumi.ResourceOption) (*ContainerRegistryIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ContainerRegistryIamBinding
	err := ctx.RegisterResource("yandex:index/containerRegistryIamBinding:ContainerRegistryIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistryIamBinding gets an existing ContainerRegistryIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistryIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryIamBindingState, opts ...pulumi.ResourceOption) (*ContainerRegistryIamBinding, error) {
	var resource ContainerRegistryIamBinding
	err := ctx.ReadResource("yandex:index/containerRegistryIamBinding:ContainerRegistryIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistryIamBinding resources.
type containerRegistryIamBindingState struct {
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Members []string `pulumi:"members"`
	// The [Yandex Container Registry](https://cloud.yandex.com/docs/container-registry/) ID to apply a binding to.
	RegistryId *string `pulumi:"registryId"`
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/container-registry/security/).
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
}

type ContainerRegistryIamBindingState struct {
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Members pulumi.StringArrayInput
	// The [Yandex Container Registry](https://cloud.yandex.com/docs/container-registry/) ID to apply a binding to.
	RegistryId pulumi.StringPtrInput
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/container-registry/security/).
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
}

func (ContainerRegistryIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryIamBindingState)(nil)).Elem()
}

type containerRegistryIamBindingArgs struct {
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Members []string `pulumi:"members"`
	// The [Yandex Container Registry](https://cloud.yandex.com/docs/container-registry/) ID to apply a binding to.
	RegistryId string `pulumi:"registryId"`
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/container-registry/security/).
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
}

// The set of arguments for constructing a ContainerRegistryIamBinding resource.
type ContainerRegistryIamBindingArgs struct {
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
	Members pulumi.StringArrayInput
	// The [Yandex Container Registry](https://cloud.yandex.com/docs/container-registry/) ID to apply a binding to.
	RegistryId pulumi.StringInput
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/container-registry/security/).
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
}

func (ContainerRegistryIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryIamBindingArgs)(nil)).Elem()
}

type ContainerRegistryIamBindingInput interface {
	pulumi.Input

	ToContainerRegistryIamBindingOutput() ContainerRegistryIamBindingOutput
	ToContainerRegistryIamBindingOutputWithContext(ctx context.Context) ContainerRegistryIamBindingOutput
}

func (*ContainerRegistryIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryIamBinding)(nil))
}

func (i *ContainerRegistryIamBinding) ToContainerRegistryIamBindingOutput() ContainerRegistryIamBindingOutput {
	return i.ToContainerRegistryIamBindingOutputWithContext(context.Background())
}

func (i *ContainerRegistryIamBinding) ToContainerRegistryIamBindingOutputWithContext(ctx context.Context) ContainerRegistryIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIamBindingOutput)
}

func (i *ContainerRegistryIamBinding) ToContainerRegistryIamBindingPtrOutput() ContainerRegistryIamBindingPtrOutput {
	return i.ToContainerRegistryIamBindingPtrOutputWithContext(context.Background())
}

func (i *ContainerRegistryIamBinding) ToContainerRegistryIamBindingPtrOutputWithContext(ctx context.Context) ContainerRegistryIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIamBindingPtrOutput)
}

type ContainerRegistryIamBindingPtrInput interface {
	pulumi.Input

	ToContainerRegistryIamBindingPtrOutput() ContainerRegistryIamBindingPtrOutput
	ToContainerRegistryIamBindingPtrOutputWithContext(ctx context.Context) ContainerRegistryIamBindingPtrOutput
}

type containerRegistryIamBindingPtrType ContainerRegistryIamBindingArgs

func (*containerRegistryIamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryIamBinding)(nil))
}

func (i *containerRegistryIamBindingPtrType) ToContainerRegistryIamBindingPtrOutput() ContainerRegistryIamBindingPtrOutput {
	return i.ToContainerRegistryIamBindingPtrOutputWithContext(context.Background())
}

func (i *containerRegistryIamBindingPtrType) ToContainerRegistryIamBindingPtrOutputWithContext(ctx context.Context) ContainerRegistryIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIamBindingPtrOutput)
}

// ContainerRegistryIamBindingArrayInput is an input type that accepts ContainerRegistryIamBindingArray and ContainerRegistryIamBindingArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryIamBindingArrayInput` via:
//
//          ContainerRegistryIamBindingArray{ ContainerRegistryIamBindingArgs{...} }
type ContainerRegistryIamBindingArrayInput interface {
	pulumi.Input

	ToContainerRegistryIamBindingArrayOutput() ContainerRegistryIamBindingArrayOutput
	ToContainerRegistryIamBindingArrayOutputWithContext(context.Context) ContainerRegistryIamBindingArrayOutput
}

type ContainerRegistryIamBindingArray []ContainerRegistryIamBindingInput

func (ContainerRegistryIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ContainerRegistryIamBinding)(nil))
}

func (i ContainerRegistryIamBindingArray) ToContainerRegistryIamBindingArrayOutput() ContainerRegistryIamBindingArrayOutput {
	return i.ToContainerRegistryIamBindingArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryIamBindingArray) ToContainerRegistryIamBindingArrayOutputWithContext(ctx context.Context) ContainerRegistryIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIamBindingArrayOutput)
}

// ContainerRegistryIamBindingMapInput is an input type that accepts ContainerRegistryIamBindingMap and ContainerRegistryIamBindingMapOutput values.
// You can construct a concrete instance of `ContainerRegistryIamBindingMapInput` via:
//
//          ContainerRegistryIamBindingMap{ "key": ContainerRegistryIamBindingArgs{...} }
type ContainerRegistryIamBindingMapInput interface {
	pulumi.Input

	ToContainerRegistryIamBindingMapOutput() ContainerRegistryIamBindingMapOutput
	ToContainerRegistryIamBindingMapOutputWithContext(context.Context) ContainerRegistryIamBindingMapOutput
}

type ContainerRegistryIamBindingMap map[string]ContainerRegistryIamBindingInput

func (ContainerRegistryIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ContainerRegistryIamBinding)(nil))
}

func (i ContainerRegistryIamBindingMap) ToContainerRegistryIamBindingMapOutput() ContainerRegistryIamBindingMapOutput {
	return i.ToContainerRegistryIamBindingMapOutputWithContext(context.Background())
}

func (i ContainerRegistryIamBindingMap) ToContainerRegistryIamBindingMapOutputWithContext(ctx context.Context) ContainerRegistryIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIamBindingMapOutput)
}

type ContainerRegistryIamBindingOutput struct {
	*pulumi.OutputState
}

func (ContainerRegistryIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryIamBinding)(nil))
}

func (o ContainerRegistryIamBindingOutput) ToContainerRegistryIamBindingOutput() ContainerRegistryIamBindingOutput {
	return o
}

func (o ContainerRegistryIamBindingOutput) ToContainerRegistryIamBindingOutputWithContext(ctx context.Context) ContainerRegistryIamBindingOutput {
	return o
}

func (o ContainerRegistryIamBindingOutput) ToContainerRegistryIamBindingPtrOutput() ContainerRegistryIamBindingPtrOutput {
	return o.ToContainerRegistryIamBindingPtrOutputWithContext(context.Background())
}

func (o ContainerRegistryIamBindingOutput) ToContainerRegistryIamBindingPtrOutputWithContext(ctx context.Context) ContainerRegistryIamBindingPtrOutput {
	return o.ApplyT(func(v ContainerRegistryIamBinding) *ContainerRegistryIamBinding {
		return &v
	}).(ContainerRegistryIamBindingPtrOutput)
}

type ContainerRegistryIamBindingPtrOutput struct {
	*pulumi.OutputState
}

func (ContainerRegistryIamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryIamBinding)(nil))
}

func (o ContainerRegistryIamBindingPtrOutput) ToContainerRegistryIamBindingPtrOutput() ContainerRegistryIamBindingPtrOutput {
	return o
}

func (o ContainerRegistryIamBindingPtrOutput) ToContainerRegistryIamBindingPtrOutputWithContext(ctx context.Context) ContainerRegistryIamBindingPtrOutput {
	return o
}

type ContainerRegistryIamBindingArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerRegistryIamBinding)(nil))
}

func (o ContainerRegistryIamBindingArrayOutput) ToContainerRegistryIamBindingArrayOutput() ContainerRegistryIamBindingArrayOutput {
	return o
}

func (o ContainerRegistryIamBindingArrayOutput) ToContainerRegistryIamBindingArrayOutputWithContext(ctx context.Context) ContainerRegistryIamBindingArrayOutput {
	return o
}

func (o ContainerRegistryIamBindingArrayOutput) Index(i pulumi.IntInput) ContainerRegistryIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerRegistryIamBinding {
		return vs[0].([]ContainerRegistryIamBinding)[vs[1].(int)]
	}).(ContainerRegistryIamBindingOutput)
}

type ContainerRegistryIamBindingMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ContainerRegistryIamBinding)(nil))
}

func (o ContainerRegistryIamBindingMapOutput) ToContainerRegistryIamBindingMapOutput() ContainerRegistryIamBindingMapOutput {
	return o
}

func (o ContainerRegistryIamBindingMapOutput) ToContainerRegistryIamBindingMapOutputWithContext(ctx context.Context) ContainerRegistryIamBindingMapOutput {
	return o
}

func (o ContainerRegistryIamBindingMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ContainerRegistryIamBinding {
		return vs[0].(map[string]ContainerRegistryIamBinding)[vs[1].(string)]
	}).(ContainerRegistryIamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerRegistryIamBindingOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIamBindingPtrOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIamBindingArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIamBindingMapOutput{})
}
