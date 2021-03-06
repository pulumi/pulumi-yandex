// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Managed PostgreSQL cluster. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/).
// [How to connect to the DB](https://cloud.yandex.com/en-ru/docs/managed-postgresql/quickstart#connect). To connect, use port 6432. The port number is not configurable.
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
// 		foo, err := yandex.GetMdbPostgresqlCluster(ctx, &GetMdbPostgresqlClusterArgs{
// 			Name: pulumi.StringRef("test"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("fqdn", foo.Hosts[0].Fqdn)
// 		return nil
// 	})
// }
// ```
func GetMdbPostgresqlCluster(ctx *pulumi.Context, args *GetMdbPostgresqlClusterArgs, opts ...pulumi.InvokeOption) (*GetMdbPostgresqlClusterResult, error) {
	var rv GetMdbPostgresqlClusterResult
	err := ctx.Invoke("yandex:index/getMdbPostgresqlCluster:getMdbPostgresqlCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbPostgresqlCluster.
type GetMdbPostgresqlClusterArgs struct {
	// The ID of the PostgreSQL cluster.
	ClusterId          *string `pulumi:"clusterId"`
	DeletionProtection *bool   `pulumi:"deletionProtection"`
	// Description of the PostgreSQL cluster.
	Description *string `pulumi:"description"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// The name of the PostgreSQL cluster.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getMdbPostgresqlCluster.
type GetMdbPostgresqlClusterResult struct {
	ClusterId string `pulumi:"clusterId"`
	// Configuration of the PostgreSQL cluster. The structure is documented below.
	Configs []GetMdbPostgresqlClusterConfig `pulumi:"configs"`
	// Timestamp of cluster creation.
	CreatedAt string `pulumi:"createdAt"`
	// A database of the PostgreSQL cluster. The structure is documented below.
	Databases          []GetMdbPostgresqlClusterDatabase `pulumi:"databases"`
	DeletionProtection bool                              `pulumi:"deletionProtection"`
	// Description of the PostgreSQL cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the PostgreSQL cluster.
	Environment string `pulumi:"environment"`
	FolderId    string `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health string `pulumi:"health"`
	// A host of the PostgreSQL cluster. The structure is documented below.
	Hosts []GetMdbPostgresqlClusterHost `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of key/value label pairs to assign to the PostgreSQL cluster.
	Labels map[string]string `pulumi:"labels"`
	// Maintenance window settings of the PostgreSQL cluster. The structure is documented below.
	MaintenanceWindows []GetMdbPostgresqlClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// Name of the database extension. For more information on available extensions see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/operations/cluster-extensions).
	Name string `pulumi:"name"`
	// ID of the network, to which the PostgreSQL cluster belongs.
	NetworkId string `pulumi:"networkId"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Status of the cluster.
	Status string `pulumi:"status"`
	// A user of the PostgreSQL cluster. The structure is documented below.
	Users []GetMdbPostgresqlClusterUser `pulumi:"users"`
}

func GetMdbPostgresqlClusterOutput(ctx *pulumi.Context, args GetMdbPostgresqlClusterOutputArgs, opts ...pulumi.InvokeOption) GetMdbPostgresqlClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMdbPostgresqlClusterResult, error) {
			args := v.(GetMdbPostgresqlClusterArgs)
			r, err := GetMdbPostgresqlCluster(ctx, &args, opts...)
			var s GetMdbPostgresqlClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMdbPostgresqlClusterResultOutput)
}

// A collection of arguments for invoking getMdbPostgresqlCluster.
type GetMdbPostgresqlClusterOutputArgs struct {
	// The ID of the PostgreSQL cluster.
	ClusterId          pulumi.StringPtrInput `pulumi:"clusterId"`
	DeletionProtection pulumi.BoolPtrInput   `pulumi:"deletionProtection"`
	// Description of the PostgreSQL cluster.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// The name of the PostgreSQL cluster.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetMdbPostgresqlClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMdbPostgresqlClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbPostgresqlCluster.
type GetMdbPostgresqlClusterResultOutput struct{ *pulumi.OutputState }

func (GetMdbPostgresqlClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMdbPostgresqlClusterResult)(nil)).Elem()
}

func (o GetMdbPostgresqlClusterResultOutput) ToGetMdbPostgresqlClusterResultOutput() GetMdbPostgresqlClusterResultOutput {
	return o
}

func (o GetMdbPostgresqlClusterResultOutput) ToGetMdbPostgresqlClusterResultOutputWithContext(ctx context.Context) GetMdbPostgresqlClusterResultOutput {
	return o
}

func (o GetMdbPostgresqlClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Configuration of the PostgreSQL cluster. The structure is documented below.
func (o GetMdbPostgresqlClusterResultOutput) Configs() GetMdbPostgresqlClusterConfigArrayOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) []GetMdbPostgresqlClusterConfig { return v.Configs }).(GetMdbPostgresqlClusterConfigArrayOutput)
}

// Timestamp of cluster creation.
func (o GetMdbPostgresqlClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A database of the PostgreSQL cluster. The structure is documented below.
func (o GetMdbPostgresqlClusterResultOutput) Databases() GetMdbPostgresqlClusterDatabaseArrayOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) []GetMdbPostgresqlClusterDatabase { return v.Databases }).(GetMdbPostgresqlClusterDatabaseArrayOutput)
}

func (o GetMdbPostgresqlClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the PostgreSQL cluster.
func (o GetMdbPostgresqlClusterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Deployment environment of the PostgreSQL cluster.
func (o GetMdbPostgresqlClusterResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o GetMdbPostgresqlClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// Aggregated health of the cluster.
func (o GetMdbPostgresqlClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

// A host of the PostgreSQL cluster. The structure is documented below.
func (o GetMdbPostgresqlClusterResultOutput) Hosts() GetMdbPostgresqlClusterHostArrayOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) []GetMdbPostgresqlClusterHost { return v.Hosts }).(GetMdbPostgresqlClusterHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMdbPostgresqlClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to the PostgreSQL cluster.
func (o GetMdbPostgresqlClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Maintenance window settings of the PostgreSQL cluster. The structure is documented below.
func (o GetMdbPostgresqlClusterResultOutput) MaintenanceWindows() GetMdbPostgresqlClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) []GetMdbPostgresqlClusterMaintenanceWindow {
		return v.MaintenanceWindows
	}).(GetMdbPostgresqlClusterMaintenanceWindowArrayOutput)
}

// Name of the database extension. For more information on available extensions see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/operations/cluster-extensions).
func (o GetMdbPostgresqlClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the network, to which the PostgreSQL cluster belongs.
func (o GetMdbPostgresqlClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// A set of ids of security groups assigned to hosts of the cluster.
func (o GetMdbPostgresqlClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Status of the cluster.
func (o GetMdbPostgresqlClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// A user of the PostgreSQL cluster. The structure is documented below.
func (o GetMdbPostgresqlClusterResultOutput) Users() GetMdbPostgresqlClusterUserArrayOutput {
	return o.ApplyT(func(v GetMdbPostgresqlClusterResult) []GetMdbPostgresqlClusterUser { return v.Users }).(GetMdbPostgresqlClusterUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMdbPostgresqlClusterResultOutput{})
}
