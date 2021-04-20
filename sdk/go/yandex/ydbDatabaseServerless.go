// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Yandex Database (serverless) resource. For more information, see
//     [the official documentation](https://cloud.yandex.com/en/docs/ydb/concepts/serverless_and_dedicated).
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
// 		_, err := yandex.NewYdbDatabaseServerless(ctx, "database1", &yandex.YdbDatabaseServerlessArgs{
// 			FolderId: pulumi.Any(data.Yandex_resourcemanager_folder.Test_folder.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type YdbDatabaseServerless struct {
	pulumi.CustomResourceState

	// The Yandex Database serverless cluster creation timestamp.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Full database path of the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	DatabasePath pulumi.StringOutput `pulumi:"databasePath"`
	// A description for the Yandex Database serverless cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Document API endpoint of the Yandex Database serverless cluster.
	DocumentApiEndpoint pulumi.StringOutput `pulumi:"documentApiEndpoint"`
	// ID of the folder that the Yandex Database serverless cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Database serverless cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location ID for the Yandex Database serverless cluster.
	LocationId pulumi.StringOutput `pulumi:"locationId"`
	// Name for the Yandex Database serverless cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Status of the Yandex Database serverless cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// Whether TLS is enabled for the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	TlsEnabled pulumi.BoolOutput `pulumi:"tlsEnabled"`
	// API endpoint of the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	YdbApiEndpoint pulumi.StringOutput `pulumi:"ydbApiEndpoint"`
	// Full endpoint of the Yandex Database serverless cluster.
	YdbFullEndpoint pulumi.StringOutput `pulumi:"ydbFullEndpoint"`
}

// NewYdbDatabaseServerless registers a new resource with the given unique name, arguments, and options.
func NewYdbDatabaseServerless(ctx *pulumi.Context,
	name string, args *YdbDatabaseServerlessArgs, opts ...pulumi.ResourceOption) (*YdbDatabaseServerless, error) {
	if args == nil {
		args = &YdbDatabaseServerlessArgs{}
	}

	var resource YdbDatabaseServerless
	err := ctx.RegisterResource("yandex:index/ydbDatabaseServerless:YdbDatabaseServerless", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetYdbDatabaseServerless gets an existing YdbDatabaseServerless resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetYdbDatabaseServerless(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *YdbDatabaseServerlessState, opts ...pulumi.ResourceOption) (*YdbDatabaseServerless, error) {
	var resource YdbDatabaseServerless
	err := ctx.ReadResource("yandex:index/ydbDatabaseServerless:YdbDatabaseServerless", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering YdbDatabaseServerless resources.
type ydbDatabaseServerlessState struct {
	// The Yandex Database serverless cluster creation timestamp.
	CreatedAt *string `pulumi:"createdAt"`
	// Full database path of the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	DatabasePath *string `pulumi:"databasePath"`
	// A description for the Yandex Database serverless cluster.
	Description *string `pulumi:"description"`
	// Document API endpoint of the Yandex Database serverless cluster.
	DocumentApiEndpoint *string `pulumi:"documentApiEndpoint"`
	// ID of the folder that the Yandex Database serverless cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Database serverless cluster.
	Labels map[string]string `pulumi:"labels"`
	// Location ID for the Yandex Database serverless cluster.
	LocationId *string `pulumi:"locationId"`
	// Name for the Yandex Database serverless cluster.
	Name *string `pulumi:"name"`
	// Status of the Yandex Database serverless cluster.
	Status *string `pulumi:"status"`
	// Whether TLS is enabled for the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	TlsEnabled *bool `pulumi:"tlsEnabled"`
	// API endpoint of the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	YdbApiEndpoint *string `pulumi:"ydbApiEndpoint"`
	// Full endpoint of the Yandex Database serverless cluster.
	YdbFullEndpoint *string `pulumi:"ydbFullEndpoint"`
}

type YdbDatabaseServerlessState struct {
	// The Yandex Database serverless cluster creation timestamp.
	CreatedAt pulumi.StringPtrInput
	// Full database path of the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	DatabasePath pulumi.StringPtrInput
	// A description for the Yandex Database serverless cluster.
	Description pulumi.StringPtrInput
	// Document API endpoint of the Yandex Database serverless cluster.
	DocumentApiEndpoint pulumi.StringPtrInput
	// ID of the folder that the Yandex Database serverless cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Yandex Database serverless cluster.
	Labels pulumi.StringMapInput
	// Location ID for the Yandex Database serverless cluster.
	LocationId pulumi.StringPtrInput
	// Name for the Yandex Database serverless cluster.
	Name pulumi.StringPtrInput
	// Status of the Yandex Database serverless cluster.
	Status pulumi.StringPtrInput
	// Whether TLS is enabled for the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	TlsEnabled pulumi.BoolPtrInput
	// API endpoint of the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	YdbApiEndpoint pulumi.StringPtrInput
	// Full endpoint of the Yandex Database serverless cluster.
	YdbFullEndpoint pulumi.StringPtrInput
}

func (YdbDatabaseServerlessState) ElementType() reflect.Type {
	return reflect.TypeOf((*ydbDatabaseServerlessState)(nil)).Elem()
}

type ydbDatabaseServerlessArgs struct {
	// A description for the Yandex Database serverless cluster.
	Description *string `pulumi:"description"`
	// ID of the folder that the Yandex Database serverless cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Database serverless cluster.
	Labels map[string]string `pulumi:"labels"`
	// Location ID for the Yandex Database serverless cluster.
	LocationId *string `pulumi:"locationId"`
	// Name for the Yandex Database serverless cluster.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a YdbDatabaseServerless resource.
type YdbDatabaseServerlessArgs struct {
	// A description for the Yandex Database serverless cluster.
	Description pulumi.StringPtrInput
	// ID of the folder that the Yandex Database serverless cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Yandex Database serverless cluster.
	Labels pulumi.StringMapInput
	// Location ID for the Yandex Database serverless cluster.
	LocationId pulumi.StringPtrInput
	// Name for the Yandex Database serverless cluster.
	Name pulumi.StringPtrInput
}

func (YdbDatabaseServerlessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ydbDatabaseServerlessArgs)(nil)).Elem()
}

type YdbDatabaseServerlessInput interface {
	pulumi.Input

	ToYdbDatabaseServerlessOutput() YdbDatabaseServerlessOutput
	ToYdbDatabaseServerlessOutputWithContext(ctx context.Context) YdbDatabaseServerlessOutput
}

func (*YdbDatabaseServerless) ElementType() reflect.Type {
	return reflect.TypeOf((*YdbDatabaseServerless)(nil))
}

func (i *YdbDatabaseServerless) ToYdbDatabaseServerlessOutput() YdbDatabaseServerlessOutput {
	return i.ToYdbDatabaseServerlessOutputWithContext(context.Background())
}

func (i *YdbDatabaseServerless) ToYdbDatabaseServerlessOutputWithContext(ctx context.Context) YdbDatabaseServerlessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YdbDatabaseServerlessOutput)
}

func (i *YdbDatabaseServerless) ToYdbDatabaseServerlessPtrOutput() YdbDatabaseServerlessPtrOutput {
	return i.ToYdbDatabaseServerlessPtrOutputWithContext(context.Background())
}

func (i *YdbDatabaseServerless) ToYdbDatabaseServerlessPtrOutputWithContext(ctx context.Context) YdbDatabaseServerlessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YdbDatabaseServerlessPtrOutput)
}

type YdbDatabaseServerlessPtrInput interface {
	pulumi.Input

	ToYdbDatabaseServerlessPtrOutput() YdbDatabaseServerlessPtrOutput
	ToYdbDatabaseServerlessPtrOutputWithContext(ctx context.Context) YdbDatabaseServerlessPtrOutput
}

type ydbDatabaseServerlessPtrType YdbDatabaseServerlessArgs

func (*ydbDatabaseServerlessPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**YdbDatabaseServerless)(nil))
}

func (i *ydbDatabaseServerlessPtrType) ToYdbDatabaseServerlessPtrOutput() YdbDatabaseServerlessPtrOutput {
	return i.ToYdbDatabaseServerlessPtrOutputWithContext(context.Background())
}

func (i *ydbDatabaseServerlessPtrType) ToYdbDatabaseServerlessPtrOutputWithContext(ctx context.Context) YdbDatabaseServerlessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YdbDatabaseServerlessPtrOutput)
}

// YdbDatabaseServerlessArrayInput is an input type that accepts YdbDatabaseServerlessArray and YdbDatabaseServerlessArrayOutput values.
// You can construct a concrete instance of `YdbDatabaseServerlessArrayInput` via:
//
//          YdbDatabaseServerlessArray{ YdbDatabaseServerlessArgs{...} }
type YdbDatabaseServerlessArrayInput interface {
	pulumi.Input

	ToYdbDatabaseServerlessArrayOutput() YdbDatabaseServerlessArrayOutput
	ToYdbDatabaseServerlessArrayOutputWithContext(context.Context) YdbDatabaseServerlessArrayOutput
}

type YdbDatabaseServerlessArray []YdbDatabaseServerlessInput

func (YdbDatabaseServerlessArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*YdbDatabaseServerless)(nil))
}

func (i YdbDatabaseServerlessArray) ToYdbDatabaseServerlessArrayOutput() YdbDatabaseServerlessArrayOutput {
	return i.ToYdbDatabaseServerlessArrayOutputWithContext(context.Background())
}

func (i YdbDatabaseServerlessArray) ToYdbDatabaseServerlessArrayOutputWithContext(ctx context.Context) YdbDatabaseServerlessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YdbDatabaseServerlessArrayOutput)
}

// YdbDatabaseServerlessMapInput is an input type that accepts YdbDatabaseServerlessMap and YdbDatabaseServerlessMapOutput values.
// You can construct a concrete instance of `YdbDatabaseServerlessMapInput` via:
//
//          YdbDatabaseServerlessMap{ "key": YdbDatabaseServerlessArgs{...} }
type YdbDatabaseServerlessMapInput interface {
	pulumi.Input

	ToYdbDatabaseServerlessMapOutput() YdbDatabaseServerlessMapOutput
	ToYdbDatabaseServerlessMapOutputWithContext(context.Context) YdbDatabaseServerlessMapOutput
}

type YdbDatabaseServerlessMap map[string]YdbDatabaseServerlessInput

func (YdbDatabaseServerlessMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*YdbDatabaseServerless)(nil))
}

func (i YdbDatabaseServerlessMap) ToYdbDatabaseServerlessMapOutput() YdbDatabaseServerlessMapOutput {
	return i.ToYdbDatabaseServerlessMapOutputWithContext(context.Background())
}

func (i YdbDatabaseServerlessMap) ToYdbDatabaseServerlessMapOutputWithContext(ctx context.Context) YdbDatabaseServerlessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YdbDatabaseServerlessMapOutput)
}

type YdbDatabaseServerlessOutput struct {
	*pulumi.OutputState
}

func (YdbDatabaseServerlessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*YdbDatabaseServerless)(nil))
}

func (o YdbDatabaseServerlessOutput) ToYdbDatabaseServerlessOutput() YdbDatabaseServerlessOutput {
	return o
}

func (o YdbDatabaseServerlessOutput) ToYdbDatabaseServerlessOutputWithContext(ctx context.Context) YdbDatabaseServerlessOutput {
	return o
}

func (o YdbDatabaseServerlessOutput) ToYdbDatabaseServerlessPtrOutput() YdbDatabaseServerlessPtrOutput {
	return o.ToYdbDatabaseServerlessPtrOutputWithContext(context.Background())
}

func (o YdbDatabaseServerlessOutput) ToYdbDatabaseServerlessPtrOutputWithContext(ctx context.Context) YdbDatabaseServerlessPtrOutput {
	return o.ApplyT(func(v YdbDatabaseServerless) *YdbDatabaseServerless {
		return &v
	}).(YdbDatabaseServerlessPtrOutput)
}

type YdbDatabaseServerlessPtrOutput struct {
	*pulumi.OutputState
}

func (YdbDatabaseServerlessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**YdbDatabaseServerless)(nil))
}

func (o YdbDatabaseServerlessPtrOutput) ToYdbDatabaseServerlessPtrOutput() YdbDatabaseServerlessPtrOutput {
	return o
}

func (o YdbDatabaseServerlessPtrOutput) ToYdbDatabaseServerlessPtrOutputWithContext(ctx context.Context) YdbDatabaseServerlessPtrOutput {
	return o
}

type YdbDatabaseServerlessArrayOutput struct{ *pulumi.OutputState }

func (YdbDatabaseServerlessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]YdbDatabaseServerless)(nil))
}

func (o YdbDatabaseServerlessArrayOutput) ToYdbDatabaseServerlessArrayOutput() YdbDatabaseServerlessArrayOutput {
	return o
}

func (o YdbDatabaseServerlessArrayOutput) ToYdbDatabaseServerlessArrayOutputWithContext(ctx context.Context) YdbDatabaseServerlessArrayOutput {
	return o
}

func (o YdbDatabaseServerlessArrayOutput) Index(i pulumi.IntInput) YdbDatabaseServerlessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) YdbDatabaseServerless {
		return vs[0].([]YdbDatabaseServerless)[vs[1].(int)]
	}).(YdbDatabaseServerlessOutput)
}

type YdbDatabaseServerlessMapOutput struct{ *pulumi.OutputState }

func (YdbDatabaseServerlessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]YdbDatabaseServerless)(nil))
}

func (o YdbDatabaseServerlessMapOutput) ToYdbDatabaseServerlessMapOutput() YdbDatabaseServerlessMapOutput {
	return o
}

func (o YdbDatabaseServerlessMapOutput) ToYdbDatabaseServerlessMapOutputWithContext(ctx context.Context) YdbDatabaseServerlessMapOutput {
	return o
}

func (o YdbDatabaseServerlessMapOutput) MapIndex(k pulumi.StringInput) YdbDatabaseServerlessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) YdbDatabaseServerless {
		return vs[0].(map[string]YdbDatabaseServerless)[vs[1].(string)]
	}).(YdbDatabaseServerlessOutput)
}

func init() {
	pulumi.RegisterOutputType(YdbDatabaseServerlessOutput{})
	pulumi.RegisterOutputType(YdbDatabaseServerlessPtrOutput{})
	pulumi.RegisterOutputType(YdbDatabaseServerlessArrayOutput{})
	pulumi.RegisterOutputType(YdbDatabaseServerlessMapOutput{})
}
