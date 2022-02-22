// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Managed SQLServer cluster. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-sqlserver/).
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
// 		opt0 := "test"
// 		foo, err := yandex.GetMdbSqlserverCluster(ctx, &GetMdbSqlserverClusterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("networkId", foo.NetworkId)
// 		return nil
// 	})
// }
// ```
func GetMdbSqlserverCluster(ctx *pulumi.Context, args *GetMdbSqlserverClusterArgs, opts ...pulumi.InvokeOption) (*GetMdbSqlserverClusterResult, error) {
	var rv GetMdbSqlserverClusterResult
	err := ctx.Invoke("yandex:index/getMdbSqlserverCluster:getMdbSqlserverCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbSqlserverCluster.
type GetMdbSqlserverClusterArgs struct {
	// The ID of the SQLServer cluster.
	ClusterId          *string `pulumi:"clusterId"`
	DeletionProtection *bool   `pulumi:"deletionProtection"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// The name of the SQLServer cluster.
	Name *string `pulumi:"name"`
	// SQLServer cluster config.
	SqlserverConfig map[string]string `pulumi:"sqlserverConfig"`
}

// A collection of values returned by getMdbSqlserverCluster.
type GetMdbSqlserverClusterResult struct {
	BackupWindowStarts []GetMdbSqlserverClusterBackupWindowStart `pulumi:"backupWindowStarts"`
	ClusterId          string                                    `pulumi:"clusterId"`
	// Creation timestamp of the key.
	CreatedAt string `pulumi:"createdAt"`
	// A database of the SQLServer cluster. The structure is documented below.
	Databases          []GetMdbSqlserverClusterDatabase `pulumi:"databases"`
	DeletionProtection bool                             `pulumi:"deletionProtection"`
	// Description of the SQLServer cluster.
	Description string `pulumi:"description"`
	// Deployment environment of the SQLServer cluster.
	Environment string `pulumi:"environment"`
	FolderId    string `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health string `pulumi:"health"`
	// A list of IDs of the host groups hosting VMs of the cluster.
	HostGroupIds []string `pulumi:"hostGroupIds"`
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts []GetMdbSqlserverClusterHost `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels map[string]string `pulumi:"labels"`
	// The name of the database.
	Name string `pulumi:"name"`
	// ID of the network, to which the SQLServer cluster belongs.
	NetworkId string `pulumi:"networkId"`
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources []GetMdbSqlserverClusterResource `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// SQLServer cluster config.
	SqlserverConfig map[string]string `pulumi:"sqlserverConfig"`
	// Status of the cluster.
	Status string `pulumi:"status"`
	// A user of the SQLServer cluster. The structure is documented below.
	Users []GetMdbSqlserverClusterUser `pulumi:"users"`
	// Version of the SQLServer cluster.
	Version string `pulumi:"version"`
}

func GetMdbSqlserverClusterOutput(ctx *pulumi.Context, args GetMdbSqlserverClusterOutputArgs, opts ...pulumi.InvokeOption) GetMdbSqlserverClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMdbSqlserverClusterResult, error) {
			args := v.(GetMdbSqlserverClusterArgs)
			r, err := GetMdbSqlserverCluster(ctx, &args, opts...)
			return *r, err
		}).(GetMdbSqlserverClusterResultOutput)
}

// A collection of arguments for invoking getMdbSqlserverCluster.
type GetMdbSqlserverClusterOutputArgs struct {
	// The ID of the SQLServer cluster.
	ClusterId          pulumi.StringPtrInput `pulumi:"clusterId"`
	DeletionProtection pulumi.BoolPtrInput   `pulumi:"deletionProtection"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// The name of the SQLServer cluster.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// SQLServer cluster config.
	SqlserverConfig pulumi.StringMapInput `pulumi:"sqlserverConfig"`
}

func (GetMdbSqlserverClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMdbSqlserverClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbSqlserverCluster.
type GetMdbSqlserverClusterResultOutput struct{ *pulumi.OutputState }

func (GetMdbSqlserverClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMdbSqlserverClusterResult)(nil)).Elem()
}

func (o GetMdbSqlserverClusterResultOutput) ToGetMdbSqlserverClusterResultOutput() GetMdbSqlserverClusterResultOutput {
	return o
}

func (o GetMdbSqlserverClusterResultOutput) ToGetMdbSqlserverClusterResultOutputWithContext(ctx context.Context) GetMdbSqlserverClusterResultOutput {
	return o
}

func (o GetMdbSqlserverClusterResultOutput) BackupWindowStarts() GetMdbSqlserverClusterBackupWindowStartArrayOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) []GetMdbSqlserverClusterBackupWindowStart {
		return v.BackupWindowStarts
	}).(GetMdbSqlserverClusterBackupWindowStartArrayOutput)
}

func (o GetMdbSqlserverClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation timestamp of the key.
func (o GetMdbSqlserverClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A database of the SQLServer cluster. The structure is documented below.
func (o GetMdbSqlserverClusterResultOutput) Databases() GetMdbSqlserverClusterDatabaseArrayOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) []GetMdbSqlserverClusterDatabase { return v.Databases }).(GetMdbSqlserverClusterDatabaseArrayOutput)
}

func (o GetMdbSqlserverClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the SQLServer cluster.
func (o GetMdbSqlserverClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

// Deployment environment of the SQLServer cluster.
func (o GetMdbSqlserverClusterResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o GetMdbSqlserverClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// Aggregated health of the cluster.
func (o GetMdbSqlserverClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

// A list of IDs of the host groups hosting VMs of the cluster.
func (o GetMdbSqlserverClusterResultOutput) HostGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) []string { return v.HostGroupIds }).(pulumi.StringArrayOutput)
}

// A host of the SQLServer cluster. The structure is documented below.
func (o GetMdbSqlserverClusterResultOutput) Hosts() GetMdbSqlserverClusterHostArrayOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) []GetMdbSqlserverClusterHost { return v.Hosts }).(GetMdbSqlserverClusterHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMdbSqlserverClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to the SQLServer cluster.
func (o GetMdbSqlserverClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the database.
func (o GetMdbSqlserverClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the network, to which the SQLServer cluster belongs.
func (o GetMdbSqlserverClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
func (o GetMdbSqlserverClusterResultOutput) Resources() GetMdbSqlserverClusterResourceArrayOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) []GetMdbSqlserverClusterResource { return v.Resources }).(GetMdbSqlserverClusterResourceArrayOutput)
}

// A set of ids of security groups assigned to hosts of the cluster.
func (o GetMdbSqlserverClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// SQLServer cluster config.
func (o GetMdbSqlserverClusterResultOutput) SqlserverConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) map[string]string { return v.SqlserverConfig }).(pulumi.StringMapOutput)
}

// Status of the cluster.
func (o GetMdbSqlserverClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// A user of the SQLServer cluster. The structure is documented below.
func (o GetMdbSqlserverClusterResultOutput) Users() GetMdbSqlserverClusterUserArrayOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) []GetMdbSqlserverClusterUser { return v.Users }).(GetMdbSqlserverClusterUserArrayOutput)
}

// Version of the SQLServer cluster.
func (o GetMdbSqlserverClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbSqlserverClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMdbSqlserverClusterResultOutput{})
}
