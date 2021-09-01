// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
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
// 		opt0 := "test"
// 		foo, err := yandex.GetMdbPostgresqlCluster(ctx, &yandex.GetMdbPostgresqlClusterArgs{
// 			Name: &opt0,
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
	Config GetMdbPostgresqlClusterConfig `pulumi:"config"`
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
	MaintenanceWindow GetMdbPostgresqlClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
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
