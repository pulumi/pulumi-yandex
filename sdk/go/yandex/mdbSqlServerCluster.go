// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a SQLServer cluster within the Yandex.Cloud. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-sqlserver/).
//
// Please read [Pricing for Managed Service for SQL Server](https://cloud.yandex.com/docs/managed-sqlserver/pricing#prices) before using SQLServer cluster.
//
// ## Example Usage
//
// Example of creating a Single Node SQLServer.
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
// 		fooVpcNetwork, err := yandex.NewVpcNetwork(ctx, "fooVpcNetwork", nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooVpcSubnet, err := yandex.NewVpcSubnet(ctx, "fooVpcSubnet", &yandex.VpcSubnetArgs{
// 			Zone:      pulumi.String("ru-central1-a"),
// 			NetworkId: fooVpcNetwork.ID(),
// 			V4CidrBlocks: pulumi.StringArray{
// 				pulumi.String("10.5.0.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewVpcSecurityGroup(ctx, "test_sg_x", &yandex.VpcSecurityGroupArgs{
// 			NetworkId: fooVpcNetwork.ID(),
// 			Ingresses: yandex.VpcSecurityGroupIngressArray{
// 				&yandex.VpcSecurityGroupIngressArgs{
// 					Protocol:    pulumi.String("ANY"),
// 					Description: pulumi.String("Allow incoming traffic from members of the same security group"),
// 					FromPort:    pulumi.Int(0),
// 					ToPort:      pulumi.Int(65535),
// 					V4CidrBlocks: pulumi.StringArray{
// 						pulumi.String("0.0.0.0/0"),
// 					},
// 				},
// 			},
// 			Egresses: yandex.VpcSecurityGroupEgressArray{
// 				&yandex.VpcSecurityGroupEgressArgs{
// 					Protocol:    pulumi.String("ANY"),
// 					Description: pulumi.String("Allow outgoing traffic to members of the same security group"),
// 					FromPort:    pulumi.Int(0),
// 					ToPort:      pulumi.Int(65535),
// 					V4CidrBlocks: pulumi.StringArray{
// 						pulumi.String("0.0.0.0/0"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewMdbSqlServerCluster(ctx, "fooMdbSqlServerCluster", &yandex.MdbSqlServerClusterArgs{
// 			Environment: pulumi.String("PRESTABLE"),
// 			NetworkId:   fooVpcNetwork.ID(),
// 			Version:     pulumi.String("2016sp2std"),
// 			Resources: &yandex.MdbSqlServerClusterResourcesArgs{
// 				ResourcePresetId: pulumi.String("s2.small"),
// 				DiskTypeId:       pulumi.String("network-ssd"),
// 				DiskSize:         pulumi.Int(20),
// 			},
// 			Labels: pulumi.StringMap{
// 				"test_key": pulumi.String("test_value"),
// 			},
// 			BackupWindowStart: &yandex.MdbSqlServerClusterBackupWindowStartArgs{
// 				Hours:   pulumi.Int(20),
// 				Minutes: pulumi.Int(30),
// 			},
// 			SqlserverConfig: pulumi.StringMap{
// 				"fill_factor_percent":           pulumi.String("49"),
// 				"optimize_for_ad_hoc_workloads": pulumi.String("true"),
// 			},
// 			Databases: yandex.MdbSqlServerClusterDatabaseArray{
// 				&yandex.MdbSqlServerClusterDatabaseArgs{
// 					Name: pulumi.String("db_name_a"),
// 				},
// 				&yandex.MdbSqlServerClusterDatabaseArgs{
// 					Name: pulumi.String("db_name"),
// 				},
// 				&yandex.MdbSqlServerClusterDatabaseArgs{
// 					Name: pulumi.String("db_name_b"),
// 				},
// 			},
// 			Users: yandex.MdbSqlServerClusterUserArray{
// 				&yandex.MdbSqlServerClusterUserArgs{
// 					Name:     pulumi.String("bob"),
// 					Password: pulumi.String("mysecurepassword"),
// 				},
// 				&yandex.MdbSqlServerClusterUserArgs{
// 					Name:     pulumi.String("alice"),
// 					Password: pulumi.String("mysecurepassword"),
// 					Permissions: yandex.MdbSqlServerClusterUserPermissionArray{
// 						&yandex.MdbSqlServerClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("DDLADMIN"),
// 							},
// 						},
// 					},
// 				},
// 				&yandex.MdbSqlServerClusterUserArgs{
// 					Name:     pulumi.String("chuck"),
// 					Password: pulumi.String("mysecurepassword"),
// 					Permissions: yandex.MdbSqlServerClusterUserPermissionArray{
// 						&yandex.MdbSqlServerClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name_a"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("OWNER"),
// 							},
// 						},
// 						&yandex.MdbSqlServerClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("OWNER"),
// 								pulumi.String("DDLADMIN"),
// 							},
// 						},
// 						&yandex.MdbSqlServerClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name_b"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("OWNER"),
// 								pulumi.String("DDLADMIN"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Hosts: yandex.MdbSqlServerClusterHostArray{
// 				&yandex.MdbSqlServerClusterHostArgs{
// 					Zone:     pulumi.String("ru-central1-a"),
// 					SubnetId: fooVpcSubnet.ID(),
// 				},
// 			},
// 			SecurityGroupIds: pulumi.StringArray{
// 				test_sg_x.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## SQLServer config
//
// If not specified `sqlserverConfig` then does not make any changes.
//
// * maxDegreeOfParallelism - Limits the number of processors to use in parallel plan execution per task. See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-max-degree-of-parallelism-server-configuration-option?view=sql-server-2016).
//
// * costThresholdForParallelism - Specifies the threshold at which SQL Server creates and runs parallel plans for queries. SQL Server creates and runs a parallel plan for a query only when the estimated cost to run a serial plan for the same query is higher than the value of the option. See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-cost-threshold-for-parallelism-server-configuration-option?view=sql-server-2016).
//
// * auditLevel - Describes how to configure login auditing to monitor SQL Server Database Engine login activity. Possible values:
//   - 0 — do not log login attempts,˚√
//   - 1 — log only failed login attempts,
//   - 2 — log only successful login attempts (not recommended),
//   - 3 — log all login attempts (not recommended).
//      See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/ssms/configure-login-auditing-sql-server-management-studio?view=sql-server-2016).
//
// * fillFactorPercent - Manages the fill factor server configuration option. When an index is created or rebuilt the fill factor determines the percentage of space on each index leaf-level page to be filled with data, reserving the rest as free space for future growth. Values 0 and 100 mean full page usage (no space reserved). See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-fill-factor-server-configuration-option?view=sql-server-2016).
// * optimizeForAdHocWorkloads - Determines whether plans should be cached only after second execution. Allows to avoid SQL cache bloat because of single-use plans. See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/optimize-for-ad-hoc-workloads-server-configuration-option?view=sql-server-2016).
//
// ## Import
//
// A cluster can be imported using the `id` of the resource, e.g.
//
// ```sh
//  $ pulumi import yandex:index/mdbSqlServerCluster:MdbSqlServerCluster foo cluster_id
// ```
type MdbSqlServerCluster struct {
	pulumi.CustomResourceState

	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbSqlServerClusterBackupWindowStartOutput `pulumi:"backupWindowStart"`
	// Creation timestamp of the cluster.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A database of the SQLServer cluster. The structure is documented below.
	Databases MdbSqlServerClusterDatabaseArrayOutput `pulumi:"databases"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the SQLServer cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health pulumi.StringOutput `pulumi:"health"`
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts MdbSqlServerClusterHostArrayOutput `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the database.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources MdbSqlServerClusterResourcesOutput `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig pulumi.StringMapOutput `pulumi:"sqlserverConfig"`
	// Status of the cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// A user of the SQLServer cluster. The structure is documented below.
	Users MdbSqlServerClusterUserArrayOutput `pulumi:"users"`
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewMdbSqlServerCluster registers a new resource with the given unique name, arguments, and options.
func NewMdbSqlServerCluster(ctx *pulumi.Context,
	name string, args *MdbSqlServerClusterArgs, opts ...pulumi.ResourceOption) (*MdbSqlServerCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Databases == nil {
		return nil, errors.New("invalid value for required argument 'Databases'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Hosts == nil {
		return nil, errors.New("invalid value for required argument 'Hosts'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.Resources == nil {
		return nil, errors.New("invalid value for required argument 'Resources'")
	}
	if args.Users == nil {
		return nil, errors.New("invalid value for required argument 'Users'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource MdbSqlServerCluster
	err := ctx.RegisterResource("yandex:index/mdbSqlServerCluster:MdbSqlServerCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbSqlServerCluster gets an existing MdbSqlServerCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbSqlServerCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbSqlServerClusterState, opts ...pulumi.ResourceOption) (*MdbSqlServerCluster, error) {
	var resource MdbSqlServerCluster
	err := ctx.ReadResource("yandex:index/mdbSqlServerCluster:MdbSqlServerCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbSqlServerCluster resources.
type mdbSqlServerClusterState struct {
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart *MdbSqlServerClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// Creation timestamp of the cluster.
	CreatedAt *string `pulumi:"createdAt"`
	// A database of the SQLServer cluster. The structure is documented below.
	Databases []MdbSqlServerClusterDatabase `pulumi:"databases"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the SQLServer cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment *string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health *string `pulumi:"health"`
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts []MdbSqlServerClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels map[string]string `pulumi:"labels"`
	// The name of the database.
	Name *string `pulumi:"name"`
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId *string `pulumi:"networkId"`
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources *MdbSqlServerClusterResources `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig map[string]string `pulumi:"sqlserverConfig"`
	// Status of the cluster.
	Status *string `pulumi:"status"`
	// A user of the SQLServer cluster. The structure is documented below.
	Users []MdbSqlServerClusterUser `pulumi:"users"`
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version *string `pulumi:"version"`
}

type MdbSqlServerClusterState struct {
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbSqlServerClusterBackupWindowStartPtrInput
	// Creation timestamp of the cluster.
	CreatedAt pulumi.StringPtrInput
	// A database of the SQLServer cluster. The structure is documented below.
	Databases MdbSqlServerClusterDatabaseArrayInput
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the SQLServer cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Aggregated health of the cluster.
	Health pulumi.StringPtrInput
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts MdbSqlServerClusterHostArrayInput
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels pulumi.StringMapInput
	// The name of the database.
	Name pulumi.StringPtrInput
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId pulumi.StringPtrInput
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources MdbSqlServerClusterResourcesPtrInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig pulumi.StringMapInput
	// Status of the cluster.
	Status pulumi.StringPtrInput
	// A user of the SQLServer cluster. The structure is documented below.
	Users MdbSqlServerClusterUserArrayInput
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version pulumi.StringPtrInput
}

func (MdbSqlServerClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbSqlServerClusterState)(nil)).Elem()
}

type mdbSqlServerClusterArgs struct {
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart *MdbSqlServerClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// A database of the SQLServer cluster. The structure is documented below.
	Databases []MdbSqlServerClusterDatabase `pulumi:"databases"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the SQLServer cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts []MdbSqlServerClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels map[string]string `pulumi:"labels"`
	// The name of the database.
	Name *string `pulumi:"name"`
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId string `pulumi:"networkId"`
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources MdbSqlServerClusterResources `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig map[string]string `pulumi:"sqlserverConfig"`
	// A user of the SQLServer cluster. The structure is documented below.
	Users []MdbSqlServerClusterUser `pulumi:"users"`
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a MdbSqlServerCluster resource.
type MdbSqlServerClusterArgs struct {
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbSqlServerClusterBackupWindowStartPtrInput
	// A database of the SQLServer cluster. The structure is documented below.
	Databases MdbSqlServerClusterDatabaseArrayInput
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the SQLServer cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts MdbSqlServerClusterHostArrayInput
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels pulumi.StringMapInput
	// The name of the database.
	Name pulumi.StringPtrInput
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId pulumi.StringInput
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources MdbSqlServerClusterResourcesInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig pulumi.StringMapInput
	// A user of the SQLServer cluster. The structure is documented below.
	Users MdbSqlServerClusterUserArrayInput
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version pulumi.StringInput
}

func (MdbSqlServerClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbSqlServerClusterArgs)(nil)).Elem()
}

type MdbSqlServerClusterInput interface {
	pulumi.Input

	ToMdbSqlServerClusterOutput() MdbSqlServerClusterOutput
	ToMdbSqlServerClusterOutputWithContext(ctx context.Context) MdbSqlServerClusterOutput
}

func (*MdbSqlServerCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*MdbSqlServerCluster)(nil))
}

func (i *MdbSqlServerCluster) ToMdbSqlServerClusterOutput() MdbSqlServerClusterOutput {
	return i.ToMdbSqlServerClusterOutputWithContext(context.Background())
}

func (i *MdbSqlServerCluster) ToMdbSqlServerClusterOutputWithContext(ctx context.Context) MdbSqlServerClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbSqlServerClusterOutput)
}

func (i *MdbSqlServerCluster) ToMdbSqlServerClusterPtrOutput() MdbSqlServerClusterPtrOutput {
	return i.ToMdbSqlServerClusterPtrOutputWithContext(context.Background())
}

func (i *MdbSqlServerCluster) ToMdbSqlServerClusterPtrOutputWithContext(ctx context.Context) MdbSqlServerClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbSqlServerClusterPtrOutput)
}

type MdbSqlServerClusterPtrInput interface {
	pulumi.Input

	ToMdbSqlServerClusterPtrOutput() MdbSqlServerClusterPtrOutput
	ToMdbSqlServerClusterPtrOutputWithContext(ctx context.Context) MdbSqlServerClusterPtrOutput
}

type mdbSqlServerClusterPtrType MdbSqlServerClusterArgs

func (*mdbSqlServerClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbSqlServerCluster)(nil))
}

func (i *mdbSqlServerClusterPtrType) ToMdbSqlServerClusterPtrOutput() MdbSqlServerClusterPtrOutput {
	return i.ToMdbSqlServerClusterPtrOutputWithContext(context.Background())
}

func (i *mdbSqlServerClusterPtrType) ToMdbSqlServerClusterPtrOutputWithContext(ctx context.Context) MdbSqlServerClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbSqlServerClusterPtrOutput)
}

// MdbSqlServerClusterArrayInput is an input type that accepts MdbSqlServerClusterArray and MdbSqlServerClusterArrayOutput values.
// You can construct a concrete instance of `MdbSqlServerClusterArrayInput` via:
//
//          MdbSqlServerClusterArray{ MdbSqlServerClusterArgs{...} }
type MdbSqlServerClusterArrayInput interface {
	pulumi.Input

	ToMdbSqlServerClusterArrayOutput() MdbSqlServerClusterArrayOutput
	ToMdbSqlServerClusterArrayOutputWithContext(context.Context) MdbSqlServerClusterArrayOutput
}

type MdbSqlServerClusterArray []MdbSqlServerClusterInput

func (MdbSqlServerClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MdbSqlServerCluster)(nil))
}

func (i MdbSqlServerClusterArray) ToMdbSqlServerClusterArrayOutput() MdbSqlServerClusterArrayOutput {
	return i.ToMdbSqlServerClusterArrayOutputWithContext(context.Background())
}

func (i MdbSqlServerClusterArray) ToMdbSqlServerClusterArrayOutputWithContext(ctx context.Context) MdbSqlServerClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbSqlServerClusterArrayOutput)
}

// MdbSqlServerClusterMapInput is an input type that accepts MdbSqlServerClusterMap and MdbSqlServerClusterMapOutput values.
// You can construct a concrete instance of `MdbSqlServerClusterMapInput` via:
//
//          MdbSqlServerClusterMap{ "key": MdbSqlServerClusterArgs{...} }
type MdbSqlServerClusterMapInput interface {
	pulumi.Input

	ToMdbSqlServerClusterMapOutput() MdbSqlServerClusterMapOutput
	ToMdbSqlServerClusterMapOutputWithContext(context.Context) MdbSqlServerClusterMapOutput
}

type MdbSqlServerClusterMap map[string]MdbSqlServerClusterInput

func (MdbSqlServerClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MdbSqlServerCluster)(nil))
}

func (i MdbSqlServerClusterMap) ToMdbSqlServerClusterMapOutput() MdbSqlServerClusterMapOutput {
	return i.ToMdbSqlServerClusterMapOutputWithContext(context.Background())
}

func (i MdbSqlServerClusterMap) ToMdbSqlServerClusterMapOutputWithContext(ctx context.Context) MdbSqlServerClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbSqlServerClusterMapOutput)
}

type MdbSqlServerClusterOutput struct {
	*pulumi.OutputState
}

func (MdbSqlServerClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MdbSqlServerCluster)(nil))
}

func (o MdbSqlServerClusterOutput) ToMdbSqlServerClusterOutput() MdbSqlServerClusterOutput {
	return o
}

func (o MdbSqlServerClusterOutput) ToMdbSqlServerClusterOutputWithContext(ctx context.Context) MdbSqlServerClusterOutput {
	return o
}

func (o MdbSqlServerClusterOutput) ToMdbSqlServerClusterPtrOutput() MdbSqlServerClusterPtrOutput {
	return o.ToMdbSqlServerClusterPtrOutputWithContext(context.Background())
}

func (o MdbSqlServerClusterOutput) ToMdbSqlServerClusterPtrOutputWithContext(ctx context.Context) MdbSqlServerClusterPtrOutput {
	return o.ApplyT(func(v MdbSqlServerCluster) *MdbSqlServerCluster {
		return &v
	}).(MdbSqlServerClusterPtrOutput)
}

type MdbSqlServerClusterPtrOutput struct {
	*pulumi.OutputState
}

func (MdbSqlServerClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbSqlServerCluster)(nil))
}

func (o MdbSqlServerClusterPtrOutput) ToMdbSqlServerClusterPtrOutput() MdbSqlServerClusterPtrOutput {
	return o
}

func (o MdbSqlServerClusterPtrOutput) ToMdbSqlServerClusterPtrOutputWithContext(ctx context.Context) MdbSqlServerClusterPtrOutput {
	return o
}

type MdbSqlServerClusterArrayOutput struct{ *pulumi.OutputState }

func (MdbSqlServerClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MdbSqlServerCluster)(nil))
}

func (o MdbSqlServerClusterArrayOutput) ToMdbSqlServerClusterArrayOutput() MdbSqlServerClusterArrayOutput {
	return o
}

func (o MdbSqlServerClusterArrayOutput) ToMdbSqlServerClusterArrayOutputWithContext(ctx context.Context) MdbSqlServerClusterArrayOutput {
	return o
}

func (o MdbSqlServerClusterArrayOutput) Index(i pulumi.IntInput) MdbSqlServerClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MdbSqlServerCluster {
		return vs[0].([]MdbSqlServerCluster)[vs[1].(int)]
	}).(MdbSqlServerClusterOutput)
}

type MdbSqlServerClusterMapOutput struct{ *pulumi.OutputState }

func (MdbSqlServerClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MdbSqlServerCluster)(nil))
}

func (o MdbSqlServerClusterMapOutput) ToMdbSqlServerClusterMapOutput() MdbSqlServerClusterMapOutput {
	return o
}

func (o MdbSqlServerClusterMapOutput) ToMdbSqlServerClusterMapOutputWithContext(ctx context.Context) MdbSqlServerClusterMapOutput {
	return o
}

func (o MdbSqlServerClusterMapOutput) MapIndex(k pulumi.StringInput) MdbSqlServerClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MdbSqlServerCluster {
		return vs[0].(map[string]MdbSqlServerCluster)[vs[1].(string)]
	}).(MdbSqlServerClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(MdbSqlServerClusterOutput{})
	pulumi.RegisterOutputType(MdbSqlServerClusterPtrOutput{})
	pulumi.RegisterOutputType(MdbSqlServerClusterArrayOutput{})
	pulumi.RegisterOutputType(MdbSqlServerClusterMapOutput{})
}
