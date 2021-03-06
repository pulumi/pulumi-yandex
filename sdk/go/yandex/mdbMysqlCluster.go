// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a MySQL cluster within the Yandex.Cloud. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
//
// ## Example Usage
//
// Example of creating a Single Node MySQL.
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
// 		_, err = yandex.NewMdbMysqlCluster(ctx, "fooMdbMysqlCluster", &yandex.MdbMysqlClusterArgs{
// 			Environment: pulumi.String("PRESTABLE"),
// 			NetworkId:   fooVpcNetwork.ID(),
// 			Version:     pulumi.String("8.0"),
// 			Resources: &MdbMysqlClusterResourcesArgs{
// 				ResourcePresetId: pulumi.String("s2.micro"),
// 				DiskTypeId:       pulumi.String("network-ssd"),
// 				DiskSize:         pulumi.Int(16),
// 			},
// 			MysqlConfig: pulumi.StringMap{
// 				"sql_mode":                      pulumi.String("ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION"),
// 				"max_connections":               pulumi.String("100"),
// 				"default_authentication_plugin": pulumi.String("MYSQL_NATIVE_PASSWORD"),
// 				"innodb_print_all_deadlocks":    pulumi.String("true"),
// 			},
// 			Access: &MdbMysqlClusterAccessArgs{
// 				WebSql: pulumi.Bool(true),
// 			},
// 			Databases: MdbMysqlClusterDatabaseArray{
// 				&MdbMysqlClusterDatabaseArgs{
// 					Name: pulumi.String("db_name"),
// 				},
// 			},
// 			Users: MdbMysqlClusterUserArray{
// 				&MdbMysqlClusterUserArgs{
// 					Name:     pulumi.String("user_name"),
// 					Password: pulumi.String("your_password"),
// 					Permissions: MdbMysqlClusterUserPermissionArray{
// 						&MdbMysqlClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("ALL"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Hosts: MdbMysqlClusterHostArray{
// 				&MdbMysqlClusterHostArgs{
// 					Zone:     pulumi.String("ru-central1-a"),
// 					SubnetId: fooVpcSubnet.ID(),
// 				},
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
// Example of creating a High-Availability(HA) MySQL Cluster.
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
// 				pulumi.String("10.1.0.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := yandex.NewVpcSubnet(ctx, "bar", &yandex.VpcSubnetArgs{
// 			Zone:      pulumi.String("ru-central1-b"),
// 			NetworkId: fooVpcNetwork.ID(),
// 			V4CidrBlocks: pulumi.StringArray{
// 				pulumi.String("10.2.0.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewMdbMysqlCluster(ctx, "fooMdbMysqlCluster", &yandex.MdbMysqlClusterArgs{
// 			Environment: pulumi.String("PRESTABLE"),
// 			NetworkId:   fooVpcNetwork.ID(),
// 			Version:     pulumi.String("8.0"),
// 			Resources: &MdbMysqlClusterResourcesArgs{
// 				ResourcePresetId: pulumi.String("s2.micro"),
// 				DiskTypeId:       pulumi.String("network-ssd"),
// 				DiskSize:         pulumi.Int(16),
// 			},
// 			Databases: MdbMysqlClusterDatabaseArray{
// 				&MdbMysqlClusterDatabaseArgs{
// 					Name: pulumi.String("db_name"),
// 				},
// 			},
// 			MaintenanceWindow: &MdbMysqlClusterMaintenanceWindowArgs{
// 				Type: pulumi.String("WEEKLY"),
// 				Day:  pulumi.String("SAT"),
// 				Hour: pulumi.Int(12),
// 			},
// 			Users: MdbMysqlClusterUserArray{
// 				&MdbMysqlClusterUserArgs{
// 					Name:     pulumi.String("user_name"),
// 					Password: pulumi.String("your_password"),
// 					Permissions: MdbMysqlClusterUserPermissionArray{
// 						&MdbMysqlClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("ALL"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Hosts: MdbMysqlClusterHostArray{
// 				&MdbMysqlClusterHostArgs{
// 					Zone:     pulumi.String("ru-central1-a"),
// 					SubnetId: fooVpcSubnet.ID(),
// 				},
// 				&MdbMysqlClusterHostArgs{
// 					Zone:     pulumi.String("ru-central1-b"),
// 					SubnetId: bar.ID(),
// 				},
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
// Example of creating a MySQL Cluster with cascade replicas: HA-group consist of 'na-1' and 'na-2', cascade replicas form a chain 'na-1' > 'nb-1' > 'nb-2'
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
// 				pulumi.String("10.1.0.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := yandex.NewVpcSubnet(ctx, "bar", &yandex.VpcSubnetArgs{
// 			Zone:      pulumi.String("ru-central1-b"),
// 			NetworkId: fooVpcNetwork.ID(),
// 			V4CidrBlocks: pulumi.StringArray{
// 				pulumi.String("10.2.0.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewMdbMysqlCluster(ctx, "fooMdbMysqlCluster", &yandex.MdbMysqlClusterArgs{
// 			Environment: pulumi.String("PRESTABLE"),
// 			NetworkId:   fooVpcNetwork.ID(),
// 			Version:     pulumi.String("8.0"),
// 			Resources: &MdbMysqlClusterResourcesArgs{
// 				ResourcePresetId: pulumi.String("s2.micro"),
// 				DiskTypeId:       pulumi.String("network-ssd"),
// 				DiskSize:         pulumi.Int(16),
// 			},
// 			Databases: MdbMysqlClusterDatabaseArray{
// 				&MdbMysqlClusterDatabaseArgs{
// 					Name: pulumi.String("db_name"),
// 				},
// 			},
// 			MaintenanceWindow: &MdbMysqlClusterMaintenanceWindowArgs{
// 				Type: pulumi.String("WEEKLY"),
// 				Day:  pulumi.String("SAT"),
// 				Hour: pulumi.Int(12),
// 			},
// 			Users: MdbMysqlClusterUserArray{
// 				&MdbMysqlClusterUserArgs{
// 					Name:     pulumi.String("user_name"),
// 					Password: pulumi.String("your_password"),
// 					Permissions: MdbMysqlClusterUserPermissionArray{
// 						&MdbMysqlClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("ALL"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Hosts: MdbMysqlClusterHostArray{
// 				&MdbMysqlClusterHostArgs{
// 					Zone:     pulumi.String("ru-central1-a"),
// 					Name:     pulumi.String("na-1"),
// 					SubnetId: fooVpcSubnet.ID(),
// 				},
// 				&MdbMysqlClusterHostArgs{
// 					Zone:     pulumi.String("ru-central1-a"),
// 					Name:     pulumi.String("na-2"),
// 					SubnetId: fooVpcSubnet.ID(),
// 				},
// 				&MdbMysqlClusterHostArgs{
// 					Zone:                  pulumi.String("ru-central1-b"),
// 					Name:                  pulumi.String("nb-1"),
// 					ReplicationSourceName: pulumi.String("na-1"),
// 					SubnetId:              bar.ID(),
// 				},
// 				&MdbMysqlClusterHostArgs{
// 					Zone:                  pulumi.String("ru-central1-b"),
// 					Name:                  pulumi.String("nb-2"),
// 					ReplicationSourceName: pulumi.String("nb-1"),
// 					SubnetId:              bar.ID(),
// 				},
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
// Example of creating a Single Node MySQL with user params.
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
// 		_, err = yandex.NewMdbMysqlCluster(ctx, "fooMdbMysqlCluster", &yandex.MdbMysqlClusterArgs{
// 			Environment: pulumi.String("PRESTABLE"),
// 			NetworkId:   fooVpcNetwork.ID(),
// 			Version:     pulumi.String("8.0"),
// 			Resources: &MdbMysqlClusterResourcesArgs{
// 				ResourcePresetId: pulumi.String("s2.micro"),
// 				DiskTypeId:       pulumi.String("network-ssd"),
// 				DiskSize:         pulumi.Int(16),
// 			},
// 			Databases: MdbMysqlClusterDatabaseArray{
// 				&MdbMysqlClusterDatabaseArgs{
// 					Name: pulumi.String("db_name"),
// 				},
// 			},
// 			MaintenanceWindow: &MdbMysqlClusterMaintenanceWindowArgs{
// 				Type: pulumi.String("ANYTIME"),
// 			},
// 			Users: MdbMysqlClusterUserArray{
// 				&MdbMysqlClusterUserArgs{
// 					Name:     pulumi.String("user_name"),
// 					Password: pulumi.String("your_password"),
// 					Permissions: MdbMysqlClusterUserPermissionArray{
// 						&MdbMysqlClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("ALL"),
// 							},
// 						},
// 					},
// 					ConnectionLimits: &MdbMysqlClusterUserConnectionLimitsArgs{
// 						MaxQuestionsPerHour: pulumi.Int(10),
// 					},
// 					GlobalPermissions: pulumi.StringArray{
// 						pulumi.String("REPLICATION_SLAVE"),
// 						pulumi.String("PROCESS"),
// 					},
// 					AuthenticationPlugin: pulumi.String("CACHING_SHA2_PASSWORD"),
// 				},
// 			},
// 			Hosts: MdbMysqlClusterHostArray{
// 				&MdbMysqlClusterHostArgs{
// 					Zone:     pulumi.String("ru-central1-a"),
// 					SubnetId: fooVpcSubnet.ID(),
// 				},
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
// Example of restoring MySQL cluster.
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
// 		_, err = yandex.NewMdbMysqlCluster(ctx, "fooMdbMysqlCluster", &yandex.MdbMysqlClusterArgs{
// 			Environment: pulumi.String("PRESTABLE"),
// 			NetworkId:   fooVpcNetwork.ID(),
// 			Version:     pulumi.String("8.0"),
// 			Restore: &MdbMysqlClusterRestoreArgs{
// 				BackupId: pulumi.String("c9qj2tns23432471d9qha:stream_20210122T141717Z"),
// 				Time:     pulumi.String("2021-01-23T15:04:05"),
// 			},
// 			Resources: &MdbMysqlClusterResourcesArgs{
// 				ResourcePresetId: pulumi.String("s2.micro"),
// 				DiskTypeId:       pulumi.String("network-ssd"),
// 				DiskSize:         pulumi.Int(16),
// 			},
// 			Databases: MdbMysqlClusterDatabaseArray{
// 				&MdbMysqlClusterDatabaseArgs{
// 					Name: pulumi.String("db_name"),
// 				},
// 			},
// 			Users: MdbMysqlClusterUserArray{
// 				&MdbMysqlClusterUserArgs{
// 					Name:     pulumi.String("user_name"),
// 					Password: pulumi.String("your_password"),
// 					Permissions: MdbMysqlClusterUserPermissionArray{
// 						&MdbMysqlClusterUserPermissionArgs{
// 							DatabaseName: pulumi.String("db_name"),
// 							Roles: pulumi.StringArray{
// 								pulumi.String("ALL"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Hosts: MdbMysqlClusterHostArray{
// 				&MdbMysqlClusterHostArgs{
// 					Zone:     pulumi.String("ru-central1-a"),
// 					SubnetId: fooVpcSubnet.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## MySQL config
//
// If not specified `mysqlConfig` then does not make any changes.
//
// * `sqlMode` default value: `ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION`
//
// some of:\
// 	-	1: "ALLOW_INVALID_DATES"
// 	-	2: "ANSI_QUOTES"
// 	-	3: "ERROR_FOR_DIVISION_BY_ZERO"
// 	-	4: "HIGH_NOT_PRECEDENCE"
// 	-	5: "IGNORE_SPACE"
// 	-	6: "NO_AUTO_VALUE_ON_ZERO"
// 	-	7: "NO_BACKSLASH_ESCAPES"
// 	-	8: "NO_ENGINE_SUBSTITUTION"
// 	-	9: "NO_UNSIGNED_SUBTRACTION"
// 	-	10: "NO_ZERO_DATE"
// 	-	11: "NO_ZERO_IN_DATE"
// 	-	15: "ONLY_FULL_GROUP_BY"
// 	-	16: "PAD_CHAR_TO_FULL_LENGTH"
// 	-	17: "PIPES_AS_CONCAT"
// 	-	18: "REAL_AS_FLOAT"
// 	-	19: "STRICT_ALL_TABLES"
// 	-	20: "STRICT_TRANS_TABLES"
// 	-	21: "TIME_TRUNCATE_FRACTIONAL"
// 	-	22: "ANSI"
// 	-	23: "TRADITIONAL"
// 	-	24: "NO_DIR_IN_CREATE"
// or:
//   - 0: "SQLMODE_UNSPECIFIED"
//
// ### MysqlConfig 8.0
// * `auditLog` boolean
//
// * `autoIncrementIncrement` integer
//
// * `autoIncrementOffset` integer
//
// * `binlogCacheSize` integer
//
// * `binlogGroupCommitSyncDelay` integer
//
// * `binlogRowImage` one of:
//   - 0: "BINLOG_ROW_IMAGE_UNSPECIFIED"
//   - 1: "FULL"
//   - 2: "MINIMAL"
//   - 3: "NOBLOB"
//
// * `binlogRowsQueryLogEvents` boolean
//
// * `characterSetServer` text
//
// * `collationServer` text
//
// * `defaultAuthenticationPlugin` one of:
//   - 0: "AUTH_PLUGIN_UNSPECIFIED"
//   - 1: "MYSQL_NATIVE_PASSWORD"
//   - 2: "CACHING_SHA2_PASSWORD"
//   - 3: "SHA256_PASSWORD"
//
// * `defaultTimeZone` text
//
// * `explicitDefaultsForTimestamp` boolean
//
// * `generalLog` boolean
//
// * `groupConcatMaxLen` integer
//
// * `innodbAdaptiveHashIndex` boolean
//
// * `innodbBufferPoolSize` integer
//
// * `innodbFlushLogAtTrxCommit` integer
//
// * `innodbIoCapacity` integer
//
// * `innodbIoCapacityMax` integer
//
// * `innodbLockWaitTimeout` integer
//
// * `innodbLogBufferSize` integer
//
// * `innodbLogFileSize` integer
//
// * `innodbNumaInterleave` boolean
//
// * `innodbPrintAllDeadlocks` boolean
//
// * `innodbPurgeThreads` integer
//
// * `innodbReadIoThreads` integer
//
// * `innodbTempDataFileMaxSize` integer
//
// * `innodbThreadConcurrency` integer
//
// * `innodbWriteIoThreads` integer
//
// * `joinBufferSize` integer
//
// * `longQueryTime` float
//
// * `maxAllowedPacket` integer
//
// * `maxConnections` integer
//
// * `maxHeapTableSize` integer
//
// * `netReadTimeout` integer
//
// * `netWriteTimeout` integer
//
// * `regexpTimeLimit` integer
//
// * `rplSemiSyncMasterWaitForSlaveCount` integer
//
// * `slaveParallelType` one of:
//   - 0: "SLAVE_PARALLEL_TYPE_UNSPECIFIED"
//   - 1: "DATABASE"
//   - 2: "LOGICAL_CLOCK"
//
// * `slaveParallelWorkers` integer
//
// * `sortBufferSize` integer
//
// * `syncBinlog` integer
//
// * `tableDefinitionCache` integer
//
// * `tableOpenCache` integer
//
// * `tableOpenCacheInstances` integer
//
// * `threadCacheSize` integer
//
// * `threadStack` integer
//
// * `tmpTableSize` integer
//
// * `transactionIsolation` one of:
//   - 0: "TRANSACTION_ISOLATION_UNSPECIFIED"
//   - 1: "READ_COMMITTED"
//   - 2: "REPEATABLE_READ"
//   - 3: "SERIALIZABLE"
//
// ### MysqlConfig 5.7
// * `auditLog` boolean
//
// * `autoIncrementIncrement` integer
//
// * `autoIncrementOffset` integer
//
// * `binlogCacheSize` integer
//
// * `binlogGroupCommitSyncDelay` integer
//
// * `binlogRowImage` one of:
//   - 0: "BINLOG_ROW_IMAGE_UNSPECIFIED"
//   - 1: "FULL"
//   - 2: "MINIMAL"
//   - 3: "NOBLOB"
//
// * `binlogRowsQueryLogEvents` boolean
//
// * `characterSetServer` text
//
// * `collationServer` text
//
// * `defaultAuthenticationPlugin` one of:
//   - 0: "AUTH_PLUGIN_UNSPECIFIED"
//   - 1: "MYSQL_NATIVE_PASSWORD"
//   - 2: "CACHING_SHA2_PASSWORD"
//   - 3: "SHA256_PASSWORD"
//
// * `defaultTimeZone` text
//
// * `explicitDefaultsForTimestamp` boolean
//
// * `generalLog` boolean
//
// * `groupConcatMaxLen` integer
//
// * `innodbAdaptiveHashIndex` boolean
//
// * `innodbBufferPoolSize` integer
//
// * `innodbFlushLogAtTrxCommit` integer
//
// * `innodbIoCapacity` integer
//
// * `innodbIoCapacityMax` integer
//
// * `innodbLockWaitTimeout` integer
//
// * `innodbLogBufferSize` integer
//
// * `innodbLogFileSize` integer
//
// * `innodbNumaInterleave` boolean
//
// * `innodbPrintAllDeadlocks` boolean
//
// * `innodbPurgeThreads` integer
//
// * `innodbReadIoThreads` integer
//
// * `innodbTempDataFileMaxSize` integer
//
// * `innodbThreadConcurrency` integer
//
// * `innodbWriteIoThreads` integer
//
// * `joinBufferSize` integer
//
// * `longQueryTime` float
//
// * `maxAllowedPacket` integer
//
// * `maxConnections` integer
//
// * `maxHeapTableSize` integer
//
// * `netReadTimeout` integer
//
// * `netWriteTimeout` integer
//
// * `rplSemiSyncMasterWaitForSlaveCount` integer
//
// * `slaveParallelType` one of:
//   - 0: "SLAVE_PARALLEL_TYPE_UNSPECIFIED"
//   - 1: "DATABASE"
//   - 2: "LOGICAL_CLOCK"
//
// * `slaveParallelWorkers` integer
//
// * `sortBufferSize` integer
//
// * `syncBinlog` integer
//
// * `tableDefinitionCache` integer
//
// * `tableOpenCache` integer
//
// * `tableOpenCacheInstances` integer
//
// * `threadCacheSize` integer
//
// * `threadStack` integer
//
// * `tmpTableSize` integer
//
// * `transactionIsolation` one of:
//   - 0: "TRANSACTION_ISOLATION_UNSPECIFIED"
//   - 1: "READ_COMMITTED"
//   - 2: "REPEATABLE_READ"
//   - 3: "SERIALIZABLE"
//
// ## Import
//
// A cluster can be imported using the `id` of the resource, e.g.
//
// ```sh
//  $ pulumi import yandex:index/mdbMysqlCluster:MdbMysqlCluster foo cluster_id
// ```
type MdbMysqlCluster struct {
	pulumi.CustomResourceState

	// Access policy to the MySQL cluster. The structure is documented below.
	Access MdbMysqlClusterAccessOutput `pulumi:"access"`
	// Deprecated: You can safely remove this option. There is no need to recreate host if assign_public_ip is changed.
	AllowRegenerationHost pulumi.BoolPtrOutput `pulumi:"allowRegenerationHost"`
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbMysqlClusterBackupWindowStartOutput `pulumi:"backupWindowStart"`
	// Creation timestamp of the cluster.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A database of the MySQL cluster. The structure is documented below.
	Databases MdbMysqlClusterDatabaseArrayOutput `pulumi:"databases"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the MySQL cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Deployment environment of the MySQL cluster.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health pulumi.StringOutput `pulumi:"health"`
	// A host of the MySQL cluster. The structure is documented below.
	Hosts MdbMysqlClusterHostArrayOutput `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the MySQL cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Maintenance policy of the MySQL cluster. The structure is documented below.
	MaintenanceWindow MdbMysqlClusterMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// MySQL cluster config. Detail info in "MySQL config" section (documented below).
	MysqlConfig pulumi.StringMapOutput `pulumi:"mysqlConfig"`
	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replicationSourceName` parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network, to which the MySQL cluster uses.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	Resources MdbMysqlClusterResourcesOutput `pulumi:"resources"`
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore MdbMysqlClusterRestorePtrOutput `pulumi:"restore"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Status of the cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// A user of the MySQL cluster. The structure is documented below.
	Users MdbMysqlClusterUserArrayOutput `pulumi:"users"`
	// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewMdbMysqlCluster registers a new resource with the given unique name, arguments, and options.
func NewMdbMysqlCluster(ctx *pulumi.Context,
	name string, args *MdbMysqlClusterArgs, opts ...pulumi.ResourceOption) (*MdbMysqlCluster, error) {
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
	var resource MdbMysqlCluster
	err := ctx.RegisterResource("yandex:index/mdbMysqlCluster:MdbMysqlCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbMysqlCluster gets an existing MdbMysqlCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbMysqlCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbMysqlClusterState, opts ...pulumi.ResourceOption) (*MdbMysqlCluster, error) {
	var resource MdbMysqlCluster
	err := ctx.ReadResource("yandex:index/mdbMysqlCluster:MdbMysqlCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbMysqlCluster resources.
type mdbMysqlClusterState struct {
	// Access policy to the MySQL cluster. The structure is documented below.
	Access *MdbMysqlClusterAccess `pulumi:"access"`
	// Deprecated: You can safely remove this option. There is no need to recreate host if assign_public_ip is changed.
	AllowRegenerationHost *bool `pulumi:"allowRegenerationHost"`
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart *MdbMysqlClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// Creation timestamp of the cluster.
	CreatedAt *string `pulumi:"createdAt"`
	// A database of the MySQL cluster. The structure is documented below.
	Databases []MdbMysqlClusterDatabase `pulumi:"databases"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the MySQL cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the MySQL cluster.
	Environment *string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health *string `pulumi:"health"`
	// A host of the MySQL cluster. The structure is documented below.
	Hosts []MdbMysqlClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the MySQL cluster.
	Labels map[string]string `pulumi:"labels"`
	// Maintenance policy of the MySQL cluster. The structure is documented below.
	MaintenanceWindow *MdbMysqlClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// MySQL cluster config. Detail info in "MySQL config" section (documented below).
	MysqlConfig map[string]string `pulumi:"mysqlConfig"`
	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replicationSourceName` parameter.
	Name *string `pulumi:"name"`
	// ID of the network, to which the MySQL cluster uses.
	NetworkId *string `pulumi:"networkId"`
	// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	Resources *MdbMysqlClusterResources `pulumi:"resources"`
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore *MdbMysqlClusterRestore `pulumi:"restore"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Status of the cluster.
	Status *string `pulumi:"status"`
	// A user of the MySQL cluster. The structure is documented below.
	Users []MdbMysqlClusterUser `pulumi:"users"`
	// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	Version *string `pulumi:"version"`
}

type MdbMysqlClusterState struct {
	// Access policy to the MySQL cluster. The structure is documented below.
	Access MdbMysqlClusterAccessPtrInput
	// Deprecated: You can safely remove this option. There is no need to recreate host if assign_public_ip is changed.
	AllowRegenerationHost pulumi.BoolPtrInput
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbMysqlClusterBackupWindowStartPtrInput
	// Creation timestamp of the cluster.
	CreatedAt pulumi.StringPtrInput
	// A database of the MySQL cluster. The structure is documented below.
	Databases MdbMysqlClusterDatabaseArrayInput
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the MySQL cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the MySQL cluster.
	Environment pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Aggregated health of the cluster.
	Health pulumi.StringPtrInput
	// A host of the MySQL cluster. The structure is documented below.
	Hosts MdbMysqlClusterHostArrayInput
	// A set of key/value label pairs to assign to the MySQL cluster.
	Labels pulumi.StringMapInput
	// Maintenance policy of the MySQL cluster. The structure is documented below.
	MaintenanceWindow MdbMysqlClusterMaintenanceWindowPtrInput
	// MySQL cluster config. Detail info in "MySQL config" section (documented below).
	MysqlConfig pulumi.StringMapInput
	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replicationSourceName` parameter.
	Name pulumi.StringPtrInput
	// ID of the network, to which the MySQL cluster uses.
	NetworkId pulumi.StringPtrInput
	// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	Resources MdbMysqlClusterResourcesPtrInput
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore MdbMysqlClusterRestorePtrInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// Status of the cluster.
	Status pulumi.StringPtrInput
	// A user of the MySQL cluster. The structure is documented below.
	Users MdbMysqlClusterUserArrayInput
	// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	Version pulumi.StringPtrInput
}

func (MdbMysqlClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbMysqlClusterState)(nil)).Elem()
}

type mdbMysqlClusterArgs struct {
	// Access policy to the MySQL cluster. The structure is documented below.
	Access *MdbMysqlClusterAccess `pulumi:"access"`
	// Deprecated: You can safely remove this option. There is no need to recreate host if assign_public_ip is changed.
	AllowRegenerationHost *bool `pulumi:"allowRegenerationHost"`
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart *MdbMysqlClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// A database of the MySQL cluster. The structure is documented below.
	Databases []MdbMysqlClusterDatabase `pulumi:"databases"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the MySQL cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the MySQL cluster.
	Environment string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A host of the MySQL cluster. The structure is documented below.
	Hosts []MdbMysqlClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the MySQL cluster.
	Labels map[string]string `pulumi:"labels"`
	// Maintenance policy of the MySQL cluster. The structure is documented below.
	MaintenanceWindow *MdbMysqlClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// MySQL cluster config. Detail info in "MySQL config" section (documented below).
	MysqlConfig map[string]string `pulumi:"mysqlConfig"`
	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replicationSourceName` parameter.
	Name *string `pulumi:"name"`
	// ID of the network, to which the MySQL cluster uses.
	NetworkId string `pulumi:"networkId"`
	// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	Resources MdbMysqlClusterResources `pulumi:"resources"`
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore *MdbMysqlClusterRestore `pulumi:"restore"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A user of the MySQL cluster. The structure is documented below.
	Users []MdbMysqlClusterUser `pulumi:"users"`
	// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a MdbMysqlCluster resource.
type MdbMysqlClusterArgs struct {
	// Access policy to the MySQL cluster. The structure is documented below.
	Access MdbMysqlClusterAccessPtrInput
	// Deprecated: You can safely remove this option. There is no need to recreate host if assign_public_ip is changed.
	AllowRegenerationHost pulumi.BoolPtrInput
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbMysqlClusterBackupWindowStartPtrInput
	// A database of the MySQL cluster. The structure is documented below.
	Databases MdbMysqlClusterDatabaseArrayInput
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the MySQL cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the MySQL cluster.
	Environment pulumi.StringInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A host of the MySQL cluster. The structure is documented below.
	Hosts MdbMysqlClusterHostArrayInput
	// A set of key/value label pairs to assign to the MySQL cluster.
	Labels pulumi.StringMapInput
	// Maintenance policy of the MySQL cluster. The structure is documented below.
	MaintenanceWindow MdbMysqlClusterMaintenanceWindowPtrInput
	// MySQL cluster config. Detail info in "MySQL config" section (documented below).
	MysqlConfig pulumi.StringMapInput
	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replicationSourceName` parameter.
	Name pulumi.StringPtrInput
	// ID of the network, to which the MySQL cluster uses.
	NetworkId pulumi.StringInput
	// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	Resources MdbMysqlClusterResourcesInput
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore MdbMysqlClusterRestorePtrInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// A user of the MySQL cluster. The structure is documented below.
	Users MdbMysqlClusterUserArrayInput
	// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	Version pulumi.StringInput
}

func (MdbMysqlClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbMysqlClusterArgs)(nil)).Elem()
}

type MdbMysqlClusterInput interface {
	pulumi.Input

	ToMdbMysqlClusterOutput() MdbMysqlClusterOutput
	ToMdbMysqlClusterOutputWithContext(ctx context.Context) MdbMysqlClusterOutput
}

func (*MdbMysqlCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbMysqlCluster)(nil)).Elem()
}

func (i *MdbMysqlCluster) ToMdbMysqlClusterOutput() MdbMysqlClusterOutput {
	return i.ToMdbMysqlClusterOutputWithContext(context.Background())
}

func (i *MdbMysqlCluster) ToMdbMysqlClusterOutputWithContext(ctx context.Context) MdbMysqlClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMysqlClusterOutput)
}

// MdbMysqlClusterArrayInput is an input type that accepts MdbMysqlClusterArray and MdbMysqlClusterArrayOutput values.
// You can construct a concrete instance of `MdbMysqlClusterArrayInput` via:
//
//          MdbMysqlClusterArray{ MdbMysqlClusterArgs{...} }
type MdbMysqlClusterArrayInput interface {
	pulumi.Input

	ToMdbMysqlClusterArrayOutput() MdbMysqlClusterArrayOutput
	ToMdbMysqlClusterArrayOutputWithContext(context.Context) MdbMysqlClusterArrayOutput
}

type MdbMysqlClusterArray []MdbMysqlClusterInput

func (MdbMysqlClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbMysqlCluster)(nil)).Elem()
}

func (i MdbMysqlClusterArray) ToMdbMysqlClusterArrayOutput() MdbMysqlClusterArrayOutput {
	return i.ToMdbMysqlClusterArrayOutputWithContext(context.Background())
}

func (i MdbMysqlClusterArray) ToMdbMysqlClusterArrayOutputWithContext(ctx context.Context) MdbMysqlClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMysqlClusterArrayOutput)
}

// MdbMysqlClusterMapInput is an input type that accepts MdbMysqlClusterMap and MdbMysqlClusterMapOutput values.
// You can construct a concrete instance of `MdbMysqlClusterMapInput` via:
//
//          MdbMysqlClusterMap{ "key": MdbMysqlClusterArgs{...} }
type MdbMysqlClusterMapInput interface {
	pulumi.Input

	ToMdbMysqlClusterMapOutput() MdbMysqlClusterMapOutput
	ToMdbMysqlClusterMapOutputWithContext(context.Context) MdbMysqlClusterMapOutput
}

type MdbMysqlClusterMap map[string]MdbMysqlClusterInput

func (MdbMysqlClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbMysqlCluster)(nil)).Elem()
}

func (i MdbMysqlClusterMap) ToMdbMysqlClusterMapOutput() MdbMysqlClusterMapOutput {
	return i.ToMdbMysqlClusterMapOutputWithContext(context.Background())
}

func (i MdbMysqlClusterMap) ToMdbMysqlClusterMapOutputWithContext(ctx context.Context) MdbMysqlClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMysqlClusterMapOutput)
}

type MdbMysqlClusterOutput struct{ *pulumi.OutputState }

func (MdbMysqlClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbMysqlCluster)(nil)).Elem()
}

func (o MdbMysqlClusterOutput) ToMdbMysqlClusterOutput() MdbMysqlClusterOutput {
	return o
}

func (o MdbMysqlClusterOutput) ToMdbMysqlClusterOutputWithContext(ctx context.Context) MdbMysqlClusterOutput {
	return o
}

type MdbMysqlClusterArrayOutput struct{ *pulumi.OutputState }

func (MdbMysqlClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbMysqlCluster)(nil)).Elem()
}

func (o MdbMysqlClusterArrayOutput) ToMdbMysqlClusterArrayOutput() MdbMysqlClusterArrayOutput {
	return o
}

func (o MdbMysqlClusterArrayOutput) ToMdbMysqlClusterArrayOutputWithContext(ctx context.Context) MdbMysqlClusterArrayOutput {
	return o
}

func (o MdbMysqlClusterArrayOutput) Index(i pulumi.IntInput) MdbMysqlClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MdbMysqlCluster {
		return vs[0].([]*MdbMysqlCluster)[vs[1].(int)]
	}).(MdbMysqlClusterOutput)
}

type MdbMysqlClusterMapOutput struct{ *pulumi.OutputState }

func (MdbMysqlClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbMysqlCluster)(nil)).Elem()
}

func (o MdbMysqlClusterMapOutput) ToMdbMysqlClusterMapOutput() MdbMysqlClusterMapOutput {
	return o
}

func (o MdbMysqlClusterMapOutput) ToMdbMysqlClusterMapOutputWithContext(ctx context.Context) MdbMysqlClusterMapOutput {
	return o
}

func (o MdbMysqlClusterMapOutput) MapIndex(k pulumi.StringInput) MdbMysqlClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MdbMysqlCluster {
		return vs[0].(map[string]*MdbMysqlCluster)[vs[1].(string)]
	}).(MdbMysqlClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMysqlClusterInput)(nil)).Elem(), &MdbMysqlCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMysqlClusterArrayInput)(nil)).Elem(), MdbMysqlClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMysqlClusterMapInput)(nil)).Elem(), MdbMysqlClusterMap{})
	pulumi.RegisterOutputType(MdbMysqlClusterOutput{})
	pulumi.RegisterOutputType(MdbMysqlClusterArrayOutput{})
	pulumi.RegisterOutputType(MdbMysqlClusterMapOutput{})
}
