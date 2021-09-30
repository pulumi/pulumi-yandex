// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
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
// 		myDatabase, err := yandex.LookupYdbDatabaseDedicated(ctx, &yandex.LookupYdbDatabaseDedicatedArgs{
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
	// Location of the Yandex Database cluster.
	// The structure is documented below.
	Location GetYdbDatabaseDedicatedLocation `pulumi:"location"`
	// Location ID of the Yandex Database cluster.
	LocationId string  `pulumi:"locationId"`
	Name       *string `pulumi:"name"`
	// ID of the network the Yandex Database cluster is attached to.
	NetworkId string `pulumi:"networkId"`
	// The Yandex Database cluster preset.
	ResourcePresetId string `pulumi:"resourcePresetId"`
	// Scaling policy of the Yandex Database cluster.
	// The structure is documented below.
	ScalePolicy GetYdbDatabaseDedicatedScalePolicy `pulumi:"scalePolicy"`
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
