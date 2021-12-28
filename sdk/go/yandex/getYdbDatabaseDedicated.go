// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Database (dedicated) cluster.
// For more information, see [the official documentation](https://cloud.yandex.com/en/docs/ydb/concepts/serverless_and_dedicated).
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
// 		opt0 := "some_ydb_dedicated_database_id"
// 		myDatabase, err := yandex.LookupYdbDatabaseDedicated(ctx, &GetYdbDatabaseDedicatedArgs{
// 			DatabaseId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ydbApiEndpoint", myDatabase.YdbApiEndpoint)
// 		return nil
// 	})
// }
// ```
func LookupYdbDatabaseDedicated(ctx *pulumi.Context, args *LookupYdbDatabaseDedicatedArgs, opts ...pulumi.InvokeOption) (*LookupYdbDatabaseDedicatedResult, error) {
	var rv LookupYdbDatabaseDedicatedResult
	err := ctx.Invoke("yandex:index/getYdbDatabaseDedicated:getYdbDatabaseDedicated", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getYdbDatabaseDedicated.
type LookupYdbDatabaseDedicatedArgs struct {
	// ID of the Yandex Database cluster.
	DatabaseId *string `pulumi:"databaseId"`
	// ID of the folder that the Yandex Database cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderId *string `pulumi:"folderId"`
	// Name of the Yandex Database cluster.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getYdbDatabaseDedicated.
type LookupYdbDatabaseDedicatedResult struct {
	// Whether public IP addresses are assigned to the Yandex Database cluster.
	AssignPublicIps bool `pulumi:"assignPublicIps"`
	// The Yandex Database cluster creation timestamp.
	CreatedAt  string  `pulumi:"createdAt"`
	DatabaseId *string `pulumi:"databaseId"`
	// Full database path of the Yandex Database cluster.
	// Useful for SDK configuration.
	DatabasePath string `pulumi:"databasePath"`
	// A description of the Yandex Database cluster.
	Description string  `pulumi:"description"`
	FolderId    *string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of key/value label pairs assigned to the Yandex Database cluster.
	Labels map[string]string `pulumi:"labels"`
	// Location ID of the Yandex Database cluster.
	LocationId string `pulumi:"locationId"`
	// Location of the Yandex Database cluster.
	// The structure is documented below.
	Locations []GetYdbDatabaseDedicatedLocation `pulumi:"locations"`
	Name      *string                           `pulumi:"name"`
	// ID of the network the Yandex Database cluster is attached to.
	NetworkId string `pulumi:"networkId"`
	// The Yandex Database cluster preset.
	ResourcePresetId string `pulumi:"resourcePresetId"`
	// Scaling policy of the Yandex Database cluster.
	// The structure is documented below.
	ScalePolicies []GetYdbDatabaseDedicatedScalePolicy `pulumi:"scalePolicies"`
	// Status of the Yandex Database cluster.
	Status string `pulumi:"status"`
	// A list of storage configuration options of the Yandex Database cluster.
	// The structure is documented below.
	StorageConfigs []GetYdbDatabaseDedicatedStorageConfig `pulumi:"storageConfigs"`
	// List of subnet IDs the Yandex Database cluster is attached to.
	SubnetIds []string `pulumi:"subnetIds"`
	// Whether TLS is enabled for the Yandex Database cluster.
	// Useful for SDK configuration.
	TlsEnabled bool `pulumi:"tlsEnabled"`
	// API endpoint of the Yandex Database cluster.
	// Useful for SDK configuration.
	YdbApiEndpoint string `pulumi:"ydbApiEndpoint"`
	// Full endpoint of the Yandex Database cluster.
	YdbFullEndpoint string `pulumi:"ydbFullEndpoint"`
}

func LookupYdbDatabaseDedicatedOutput(ctx *pulumi.Context, args LookupYdbDatabaseDedicatedOutputArgs, opts ...pulumi.InvokeOption) LookupYdbDatabaseDedicatedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupYdbDatabaseDedicatedResult, error) {
			args := v.(LookupYdbDatabaseDedicatedArgs)
			r, err := LookupYdbDatabaseDedicated(ctx, &args, opts...)
			return *r, err
		}).(LookupYdbDatabaseDedicatedResultOutput)
}

// A collection of arguments for invoking getYdbDatabaseDedicated.
type LookupYdbDatabaseDedicatedOutputArgs struct {
	// ID of the Yandex Database cluster.
	DatabaseId pulumi.StringPtrInput `pulumi:"databaseId"`
	// ID of the folder that the Yandex Database cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// Name of the Yandex Database cluster.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupYdbDatabaseDedicatedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupYdbDatabaseDedicatedArgs)(nil)).Elem()
}

// A collection of values returned by getYdbDatabaseDedicated.
type LookupYdbDatabaseDedicatedResultOutput struct{ *pulumi.OutputState }

func (LookupYdbDatabaseDedicatedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupYdbDatabaseDedicatedResult)(nil)).Elem()
}

func (o LookupYdbDatabaseDedicatedResultOutput) ToLookupYdbDatabaseDedicatedResultOutput() LookupYdbDatabaseDedicatedResultOutput {
	return o
}

func (o LookupYdbDatabaseDedicatedResultOutput) ToLookupYdbDatabaseDedicatedResultOutputWithContext(ctx context.Context) LookupYdbDatabaseDedicatedResultOutput {
	return o
}

// Whether public IP addresses are assigned to the Yandex Database cluster.
func (o LookupYdbDatabaseDedicatedResultOutput) AssignPublicIps() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) bool { return v.AssignPublicIps }).(pulumi.BoolOutput)
}

// The Yandex Database cluster creation timestamp.
func (o LookupYdbDatabaseDedicatedResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupYdbDatabaseDedicatedResultOutput) DatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) *string { return v.DatabaseId }).(pulumi.StringPtrOutput)
}

// Full database path of the Yandex Database cluster.
// Useful for SDK configuration.
func (o LookupYdbDatabaseDedicatedResultOutput) DatabasePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.DatabasePath }).(pulumi.StringOutput)
}

// A description of the Yandex Database cluster.
func (o LookupYdbDatabaseDedicatedResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupYdbDatabaseDedicatedResultOutput) FolderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) *string { return v.FolderId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupYdbDatabaseDedicatedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of key/value label pairs assigned to the Yandex Database cluster.
func (o LookupYdbDatabaseDedicatedResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Location ID of the Yandex Database cluster.
func (o LookupYdbDatabaseDedicatedResultOutput) LocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.LocationId }).(pulumi.StringOutput)
}

// Location of the Yandex Database cluster.
// The structure is documented below.
func (o LookupYdbDatabaseDedicatedResultOutput) Locations() GetYdbDatabaseDedicatedLocationArrayOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) []GetYdbDatabaseDedicatedLocation { return v.Locations }).(GetYdbDatabaseDedicatedLocationArrayOutput)
}

func (o LookupYdbDatabaseDedicatedResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// ID of the network the Yandex Database cluster is attached to.
func (o LookupYdbDatabaseDedicatedResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// The Yandex Database cluster preset.
func (o LookupYdbDatabaseDedicatedResultOutput) ResourcePresetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.ResourcePresetId }).(pulumi.StringOutput)
}

// Scaling policy of the Yandex Database cluster.
// The structure is documented below.
func (o LookupYdbDatabaseDedicatedResultOutput) ScalePolicies() GetYdbDatabaseDedicatedScalePolicyArrayOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) []GetYdbDatabaseDedicatedScalePolicy { return v.ScalePolicies }).(GetYdbDatabaseDedicatedScalePolicyArrayOutput)
}

// Status of the Yandex Database cluster.
func (o LookupYdbDatabaseDedicatedResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.Status }).(pulumi.StringOutput)
}

// A list of storage configuration options of the Yandex Database cluster.
// The structure is documented below.
func (o LookupYdbDatabaseDedicatedResultOutput) StorageConfigs() GetYdbDatabaseDedicatedStorageConfigArrayOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) []GetYdbDatabaseDedicatedStorageConfig {
		return v.StorageConfigs
	}).(GetYdbDatabaseDedicatedStorageConfigArrayOutput)
}

// List of subnet IDs the Yandex Database cluster is attached to.
func (o LookupYdbDatabaseDedicatedResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Whether TLS is enabled for the Yandex Database cluster.
// Useful for SDK configuration.
func (o LookupYdbDatabaseDedicatedResultOutput) TlsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) bool { return v.TlsEnabled }).(pulumi.BoolOutput)
}

// API endpoint of the Yandex Database cluster.
// Useful for SDK configuration.
func (o LookupYdbDatabaseDedicatedResultOutput) YdbApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.YdbApiEndpoint }).(pulumi.StringOutput)
}

// Full endpoint of the Yandex Database cluster.
func (o LookupYdbDatabaseDedicatedResultOutput) YdbFullEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupYdbDatabaseDedicatedResult) string { return v.YdbFullEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupYdbDatabaseDedicatedResultOutput{})
}
