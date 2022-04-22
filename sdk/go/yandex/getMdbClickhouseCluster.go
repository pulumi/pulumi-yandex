// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Managed ClickHouse cluster. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).
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
// 		foo, err := yandex.LookupMdbClickhouseCluster(ctx, &GetMdbClickhouseClusterArgs{
// 			Name: pulumi.StringRef("test"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("networkId", foo.NetworkId)
// 		return nil
// 	})
// }
// ```
func LookupMdbClickhouseCluster(ctx *pulumi.Context, args *LookupMdbClickhouseClusterArgs, opts ...pulumi.InvokeOption) (*LookupMdbClickhouseClusterResult, error) {
	var rv LookupMdbClickhouseClusterResult
	err := ctx.Invoke("yandex:index/getMdbClickhouseCluster:getMdbClickhouseCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbClickhouseCluster.
type LookupMdbClickhouseClusterArgs struct {
	CloudStorage *GetMdbClickhouseClusterCloudStorage `pulumi:"cloudStorage"`
	// The ID of the ClickHouse cluster.
	ClusterId          *string `pulumi:"clusterId"`
	DeletionProtection *bool   `pulumi:"deletionProtection"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// The name of the ClickHouse cluster.
	Name             *string `pulumi:"name"`
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

// A collection of values returned by getMdbClickhouseCluster.
type LookupMdbClickhouseClusterResult struct {
	// Access policy to the ClickHouse cluster. The structure is documented below.
	Accesses []GetMdbClickhouseClusterAccess `pulumi:"accesses"`
	// Time to start the daily backup, in the UTC timezone. The structure is documented below.
	BackupWindowStarts []GetMdbClickhouseClusterBackupWindowStart `pulumi:"backupWindowStarts"`
	// Configuration of the ClickHouse subcluster. The structure is documented below.
	Clickhouses  []GetMdbClickhouseClusterClickhouse  `pulumi:"clickhouses"`
	CloudStorage *GetMdbClickhouseClusterCloudStorage `pulumi:"cloudStorage"`
	ClusterId    string                               `pulumi:"clusterId"`
	// Creation timestamp of the key.
	CreatedAt string `pulumi:"createdAt"`
	// A database of the ClickHouse cluster. The structure is documented below.
	Databases          []GetMdbClickhouseClusterDatabase `pulumi:"databases"`
	DeletionProtection bool                              `pulumi:"deletionProtection"`
	// Description of the shard group.
	Description string `pulumi:"description"`
	// Deployment environment of the ClickHouse cluster.
	Environment string `pulumi:"environment"`
	FolderId    string `pulumi:"folderId"`
	// A set of protobuf or cap'n proto format schemas. The structure is documented below.
	FormatSchemas []GetMdbClickhouseClusterFormatSchema `pulumi:"formatSchemas"`
	// Aggregated health of the cluster.
	Health string `pulumi:"health"`
	// A host of the ClickHouse cluster. The structure is documented below.
	Hosts []GetMdbClickhouseClusterHost `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of key/value label pairs to assign to the ClickHouse cluster.
	Labels             map[string]string                          `pulumi:"labels"`
	MaintenanceWindows []GetMdbClickhouseClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// A group of machine learning models. The structure is documented below.
	MlModels []GetMdbClickhouseClusterMlModel `pulumi:"mlModels"`
	// Graphite rollup configuration name.
	Name string `pulumi:"name"`
	// ID of the network, to which the ClickHouse cluster belongs.
	NetworkId string `pulumi:"networkId"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	ServiceAccountId string   `pulumi:"serviceAccountId"`
	// A group of clickhouse shards. The structure is documented below.
	ShardGroups []GetMdbClickhouseClusterShardGroup `pulumi:"shardGroups"`
	// Grants `admin` user database management permission.
	SqlDatabaseManagement bool `pulumi:"sqlDatabaseManagement"`
	// Enables `admin` user with user management permission.
	SqlUserManagement bool `pulumi:"sqlUserManagement"`
	// Status of the cluster.
	Status string `pulumi:"status"`
	// A user of the ClickHouse cluster. The structure is documented below.
	Users   []GetMdbClickhouseClusterUser `pulumi:"users"`
	Version string                        `pulumi:"version"`
	// Configuration of the ZooKeeper subcluster. The structure is documented below.
	Zookeepers []GetMdbClickhouseClusterZookeeper `pulumi:"zookeepers"`
}

func LookupMdbClickhouseClusterOutput(ctx *pulumi.Context, args LookupMdbClickhouseClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMdbClickhouseClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMdbClickhouseClusterResult, error) {
			args := v.(LookupMdbClickhouseClusterArgs)
			r, err := LookupMdbClickhouseCluster(ctx, &args, opts...)
			var s LookupMdbClickhouseClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMdbClickhouseClusterResultOutput)
}

// A collection of arguments for invoking getMdbClickhouseCluster.
type LookupMdbClickhouseClusterOutputArgs struct {
	CloudStorage GetMdbClickhouseClusterCloudStoragePtrInput `pulumi:"cloudStorage"`
	// The ID of the ClickHouse cluster.
	ClusterId          pulumi.StringPtrInput `pulumi:"clusterId"`
	DeletionProtection pulumi.BoolPtrInput   `pulumi:"deletionProtection"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// The name of the ClickHouse cluster.
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ServiceAccountId pulumi.StringPtrInput `pulumi:"serviceAccountId"`
}

func (LookupMdbClickhouseClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbClickhouseClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbClickhouseCluster.
type LookupMdbClickhouseClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMdbClickhouseClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbClickhouseClusterResult)(nil)).Elem()
}

func (o LookupMdbClickhouseClusterResultOutput) ToLookupMdbClickhouseClusterResultOutput() LookupMdbClickhouseClusterResultOutput {
	return o
}

func (o LookupMdbClickhouseClusterResultOutput) ToLookupMdbClickhouseClusterResultOutputWithContext(ctx context.Context) LookupMdbClickhouseClusterResultOutput {
	return o
}

// Access policy to the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Accesses() GetMdbClickhouseClusterAccessArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterAccess { return v.Accesses }).(GetMdbClickhouseClusterAccessArrayOutput)
}

// Time to start the daily backup, in the UTC timezone. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) BackupWindowStarts() GetMdbClickhouseClusterBackupWindowStartArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterBackupWindowStart {
		return v.BackupWindowStarts
	}).(GetMdbClickhouseClusterBackupWindowStartArrayOutput)
}

// Configuration of the ClickHouse subcluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Clickhouses() GetMdbClickhouseClusterClickhouseArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterClickhouse { return v.Clickhouses }).(GetMdbClickhouseClusterClickhouseArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) CloudStorage() GetMdbClickhouseClusterCloudStoragePtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *GetMdbClickhouseClusterCloudStorage { return v.CloudStorage }).(GetMdbClickhouseClusterCloudStoragePtrOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation timestamp of the key.
func (o LookupMdbClickhouseClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A database of the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Databases() GetMdbClickhouseClusterDatabaseArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterDatabase { return v.Databases }).(GetMdbClickhouseClusterDatabaseArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the shard group.
func (o LookupMdbClickhouseClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

// Deployment environment of the ClickHouse cluster.
func (o LookupMdbClickhouseClusterResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// A set of protobuf or cap'n proto format schemas. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) FormatSchemas() GetMdbClickhouseClusterFormatSchemaArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterFormatSchema { return v.FormatSchemas }).(GetMdbClickhouseClusterFormatSchemaArrayOutput)
}

// Aggregated health of the cluster.
func (o LookupMdbClickhouseClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

// A host of the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Hosts() GetMdbClickhouseClusterHostArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterHost { return v.Hosts }).(GetMdbClickhouseClusterHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMdbClickhouseClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to the ClickHouse cluster.
func (o LookupMdbClickhouseClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) MaintenanceWindows() GetMdbClickhouseClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterMaintenanceWindow {
		return v.MaintenanceWindows
	}).(GetMdbClickhouseClusterMaintenanceWindowArrayOutput)
}

// A group of machine learning models. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) MlModels() GetMdbClickhouseClusterMlModelArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterMlModel { return v.MlModels }).(GetMdbClickhouseClusterMlModelArrayOutput)
}

// Graphite rollup configuration name.
func (o LookupMdbClickhouseClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the network, to which the ClickHouse cluster belongs.
func (o LookupMdbClickhouseClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// A set of ids of security groups assigned to hosts of the cluster.
func (o LookupMdbClickhouseClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

// A group of clickhouse shards. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) ShardGroups() GetMdbClickhouseClusterShardGroupArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterShardGroup { return v.ShardGroups }).(GetMdbClickhouseClusterShardGroupArrayOutput)
}

// Grants `admin` user database management permission.
func (o LookupMdbClickhouseClusterResultOutput) SqlDatabaseManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.SqlDatabaseManagement }).(pulumi.BoolOutput)
}

// Enables `admin` user with user management permission.
func (o LookupMdbClickhouseClusterResultOutput) SqlUserManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.SqlUserManagement }).(pulumi.BoolOutput)
}

// Status of the cluster.
func (o LookupMdbClickhouseClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// A user of the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Users() GetMdbClickhouseClusterUserArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterUser { return v.Users }).(GetMdbClickhouseClusterUserArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

// Configuration of the ZooKeeper subcluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Zookeepers() GetMdbClickhouseClusterZookeeperArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterZookeeper { return v.Zookeepers }).(GetMdbClickhouseClusterZookeeperArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMdbClickhouseClusterResultOutput{})
}
