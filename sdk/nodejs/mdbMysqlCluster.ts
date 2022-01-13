// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a MySQL cluster within the Yandex.Cloud. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
 *
 * ## Example Usage
 *
 * Example of creating a Single Node MySQL.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const fooVpcNetwork = new yandex.VpcNetwork("fooVpcNetwork", {});
 * const fooVpcSubnet = new yandex.VpcSubnet("fooVpcSubnet", {
 *     zone: "ru-central1-a",
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.5.0.0/24"],
 * });
 * const fooMdbMysqlCluster = new yandex.MdbMysqlCluster("fooMdbMysqlCluster", {
 *     environment: "PRESTABLE",
 *     networkId: fooVpcNetwork.id,
 *     version: "8.0",
 *     resources: {
 *         resourcePresetId: "s2.micro",
 *         diskTypeId: "network-ssd",
 *         diskSize: 16,
 *     },
 *     mysqlConfig: {
 *         sql_mode: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION",
 *         max_connections: 100,
 *         default_authentication_plugin: "MYSQL_NATIVE_PASSWORD",
 *         innodb_print_all_deadlocks: true,
 *     },
 *     access: {
 *         webSql: true,
 *     },
 *     databases: [{
 *         name: "db_name",
 *     }],
 *     users: [{
 *         name: "user_name",
 *         password: "your_password",
 *         permissions: [{
 *             databaseName: "db_name",
 *             roles: ["ALL"],
 *         }],
 *     }],
 *     hosts: [{
 *         zone: "ru-central1-a",
 *         subnetId: fooVpcSubnet.id,
 *     }],
 * });
 * ```
 *
 * Example of creating a High-Availability(HA) MySQL Cluster.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const fooVpcNetwork = new yandex.VpcNetwork("fooVpcNetwork", {});
 * const fooVpcSubnet = new yandex.VpcSubnet("fooVpcSubnet", {
 *     zone: "ru-central1-a",
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.1.0.0/24"],
 * });
 * const bar = new yandex.VpcSubnet("bar", {
 *     zone: "ru-central1-b",
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.2.0.0/24"],
 * });
 * const fooMdbMysqlCluster = new yandex.MdbMysqlCluster("fooMdbMysqlCluster", {
 *     environment: "PRESTABLE",
 *     networkId: fooVpcNetwork.id,
 *     version: "8.0",
 *     resources: {
 *         resourcePresetId: "s2.micro",
 *         diskTypeId: "network-ssd",
 *         diskSize: 16,
 *     },
 *     databases: [{
 *         name: "db_name",
 *     }],
 *     maintenanceWindow: {
 *         type: "WEEKLY",
 *         day: "SAT",
 *         hour: 12,
 *     },
 *     users: [{
 *         name: "user_name",
 *         password: "your_password",
 *         permissions: [{
 *             databaseName: "db_name",
 *             roles: ["ALL"],
 *         }],
 *     }],
 *     hosts: [
 *         {
 *             zone: "ru-central1-a",
 *             subnetId: fooVpcSubnet.id,
 *         },
 *         {
 *             zone: "ru-central1-b",
 *             subnetId: bar.id,
 *         },
 *     ],
 * });
 * ```
 *
 * Example of creating a MySQL Cluster with cascade replicas: HA-group consist of 'na-1' and 'na-2', cascade replicas form a chain 'na-1' > 'nb-1' > 'nb-2'
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const fooVpcNetwork = new yandex.VpcNetwork("fooVpcNetwork", {});
 * const fooVpcSubnet = new yandex.VpcSubnet("fooVpcSubnet", {
 *     zone: "ru-central1-a",
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.1.0.0/24"],
 * });
 * const bar = new yandex.VpcSubnet("bar", {
 *     zone: "ru-central1-b",
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.2.0.0/24"],
 * });
 * const fooMdbMysqlCluster = new yandex.MdbMysqlCluster("fooMdbMysqlCluster", {
 *     environment: "PRESTABLE",
 *     networkId: fooVpcNetwork.id,
 *     version: "8.0",
 *     resources: {
 *         resourcePresetId: "s2.micro",
 *         diskTypeId: "network-ssd",
 *         diskSize: 16,
 *     },
 *     databases: [{
 *         name: "db_name",
 *     }],
 *     maintenanceWindow: {
 *         type: "WEEKLY",
 *         day: "SAT",
 *         hour: 12,
 *     },
 *     users: [{
 *         name: "user_name",
 *         password: "your_password",
 *         permissions: [{
 *             databaseName: "db_name",
 *             roles: ["ALL"],
 *         }],
 *     }],
 *     hosts: [
 *         {
 *             zone: "ru-central1-a",
 *             name: "na-1",
 *             subnetId: fooVpcSubnet.id,
 *         },
 *         {
 *             zone: "ru-central1-a",
 *             name: "na-2",
 *             subnetId: fooVpcSubnet.id,
 *         },
 *         {
 *             zone: "ru-central1-b",
 *             name: "nb-1",
 *             replicationSourceName: "na-1",
 *             subnetId: bar.id,
 *         },
 *         {
 *             zone: "ru-central1-b",
 *             name: "nb-2",
 *             replicationSourceName: "nb-1",
 *             subnetId: bar.id,
 *         },
 *     ],
 * });
 * ```
 *
 * Example of creating a Single Node MySQL with user params.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const fooVpcNetwork = new yandex.VpcNetwork("fooVpcNetwork", {});
 * const fooVpcSubnet = new yandex.VpcSubnet("fooVpcSubnet", {
 *     zone: "ru-central1-a",
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.5.0.0/24"],
 * });
 * const fooMdbMysqlCluster = new yandex.MdbMysqlCluster("fooMdbMysqlCluster", {
 *     environment: "PRESTABLE",
 *     networkId: fooVpcNetwork.id,
 *     version: "8.0",
 *     resources: {
 *         resourcePresetId: "s2.micro",
 *         diskTypeId: "network-ssd",
 *         diskSize: 16,
 *     },
 *     databases: [{
 *         name: "db_name",
 *     }],
 *     maintenanceWindow: {
 *         type: "ANYTIME",
 *     },
 *     users: [{
 *         name: "user_name",
 *         password: "your_password",
 *         permissions: [{
 *             databaseName: "db_name",
 *             roles: ["ALL"],
 *         }],
 *         connectionLimits: {
 *             maxQuestionsPerHour: 10,
 *         },
 *         globalPermissions: [
 *             "REPLICATION_SLAVE",
 *             "PROCESS",
 *         ],
 *         authenticationPlugin: "CACHING_SHA2_PASSWORD",
 *     }],
 *     hosts: [{
 *         zone: "ru-central1-a",
 *         subnetId: fooVpcSubnet.id,
 *     }],
 * });
 * ```
 *
 * Example of restoring MySQL cluster.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const fooVpcNetwork = new yandex.VpcNetwork("fooVpcNetwork", {});
 * const fooVpcSubnet = new yandex.VpcSubnet("fooVpcSubnet", {
 *     zone: "ru-central1-a",
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.5.0.0/24"],
 * });
 * const fooMdbMysqlCluster = new yandex.MdbMysqlCluster("fooMdbMysqlCluster", {
 *     environment: "PRESTABLE",
 *     networkId: fooVpcNetwork.id,
 *     version: "8.0",
 *     restore: {
 *         backupId: "c9qj2tns23432471d9qha:stream_20210122T141717Z",
 *         time: "2021-01-23T15:04:05",
 *     },
 *     resources: {
 *         resourcePresetId: "s2.micro",
 *         diskTypeId: "network-ssd",
 *         diskSize: 16,
 *     },
 *     databases: [{
 *         name: "db_name",
 *     }],
 *     users: [{
 *         name: "user_name",
 *         password: "your_password",
 *         permissions: [{
 *             databaseName: "db_name",
 *             roles: ["ALL"],
 *         }],
 *     }],
 *     hosts: [{
 *         zone: "ru-central1-a",
 *         subnetId: fooVpcSubnet.id,
 *     }],
 * });
 * ```
 * ## MySQL config
 *
 * If not specified `mysqlConfig` then does not make any changes.
 *
 * * `sqlMode` default value: `ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION`
 *
 * some of:\
 * 	-	1: "ALLOW_INVALID_DATES"
 * 	-	2: "ANSI_QUOTES"
 * 	-	3: "ERROR_FOR_DIVISION_BY_ZERO"
 * 	-	4: "HIGH_NOT_PRECEDENCE"
 * 	-	5: "IGNORE_SPACE"
 * 	-	6: "NO_AUTO_VALUE_ON_ZERO"
 * 	-	7: "NO_BACKSLASH_ESCAPES"
 * 	-	8: "NO_ENGINE_SUBSTITUTION"
 * 	-	9: "NO_UNSIGNED_SUBTRACTION"
 * 	-	10: "NO_ZERO_DATE"
 * 	-	11: "NO_ZERO_IN_DATE"
 * 	-	15: "ONLY_FULL_GROUP_BY"
 * 	-	16: "PAD_CHAR_TO_FULL_LENGTH"
 * 	-	17: "PIPES_AS_CONCAT"
 * 	-	18: "REAL_AS_FLOAT"
 * 	-	19: "STRICT_ALL_TABLES"
 * 	-	20: "STRICT_TRANS_TABLES"
 * 	-	21: "TIME_TRUNCATE_FRACTIONAL"
 * 	-	22: "ANSI"
 * 	-	23: "TRADITIONAL"
 * 	-	24: "NO_DIR_IN_CREATE"
 * or:
 *   - 0: "SQLMODE_UNSPECIFIED"
 *
 * ### MysqlConfig 8.0
 * * `auditLog` boolean
 *
 * * `autoIncrementIncrement` integer
 *
 * * `autoIncrementOffset` integer
 *
 * * `binlogCacheSize` integer
 *
 * * `binlogGroupCommitSyncDelay` integer
 *
 * * `binlogRowImage` one of:
 *   - 0: "BINLOG_ROW_IMAGE_UNSPECIFIED"
 *   - 1: "FULL"
 *   - 2: "MINIMAL"
 *   - 3: "NOBLOB"
 *
 * * `binlogRowsQueryLogEvents` boolean
 *
 * * `characterSetServer` text
 *
 * * `collationServer` text
 *
 * * `defaultAuthenticationPlugin` one of:
 *   - 0: "AUTH_PLUGIN_UNSPECIFIED"
 *   - 1: "MYSQL_NATIVE_PASSWORD"
 *   - 2: "CACHING_SHA2_PASSWORD"
 *   - 3: "SHA256_PASSWORD"
 *
 * * `defaultTimeZone` text
 *
 * * `explicitDefaultsForTimestamp` boolean
 *
 * * `generalLog` boolean
 *
 * * `groupConcatMaxLen` integer
 *
 * * `innodbAdaptiveHashIndex` boolean
 *
 * * `innodbBufferPoolSize` integer
 *
 * * `innodbFlushLogAtTrxCommit` integer
 *
 * * `innodbIoCapacity` integer
 *
 * * `innodbIoCapacityMax` integer
 *
 * * `innodbLockWaitTimeout` integer
 *
 * * `innodbLogBufferSize` integer
 *
 * * `innodbLogFileSize` integer
 *
 * * `innodbNumaInterleave` boolean
 *
 * * `innodbPrintAllDeadlocks` boolean
 *
 * * `innodbPurgeThreads` integer
 *
 * * `innodbReadIoThreads` integer
 *
 * * `innodbTempDataFileMaxSize` integer
 *
 * * `innodbThreadConcurrency` integer
 *
 * * `innodbWriteIoThreads` integer
 *
 * * `joinBufferSize` integer
 *
 * * `longQueryTime` float
 *
 * * `maxAllowedPacket` integer
 *
 * * `maxConnections` integer
 *
 * * `maxHeapTableSize` integer
 *
 * * `netReadTimeout` integer
 *
 * * `netWriteTimeout` integer
 *
 * * `regexpTimeLimit` integer
 *
 * * `rplSemiSyncMasterWaitForSlaveCount` integer
 *
 * * `slaveParallelType` one of:
 *   - 0: "SLAVE_PARALLEL_TYPE_UNSPECIFIED"
 *   - 1: "DATABASE"
 *   - 2: "LOGICAL_CLOCK"
 *
 * * `slaveParallelWorkers` integer
 *
 * * `sortBufferSize` integer
 *
 * * `syncBinlog` integer
 *
 * * `tableDefinitionCache` integer
 *
 * * `tableOpenCache` integer
 *
 * * `tableOpenCacheInstances` integer
 *
 * * `threadCacheSize` integer
 *
 * * `threadStack` integer
 *
 * * `tmpTableSize` integer
 *
 * * `transactionIsolation` one of:
 *   - 0: "TRANSACTION_ISOLATION_UNSPECIFIED"
 *   - 1: "READ_COMMITTED"
 *   - 2: "REPEATABLE_READ"
 *   - 3: "SERIALIZABLE"
 *
 * ### MysqlConfig 5.7
 * * `auditLog` boolean
 *
 * * `autoIncrementIncrement` integer
 *
 * * `autoIncrementOffset` integer
 *
 * * `binlogCacheSize` integer
 *
 * * `binlogGroupCommitSyncDelay` integer
 *
 * * `binlogRowImage` one of:
 *   - 0: "BINLOG_ROW_IMAGE_UNSPECIFIED"
 *   - 1: "FULL"
 *   - 2: "MINIMAL"
 *   - 3: "NOBLOB"
 *
 * * `binlogRowsQueryLogEvents` boolean
 *
 * * `characterSetServer` text
 *
 * * `collationServer` text
 *
 * * `defaultAuthenticationPlugin` one of:
 *   - 0: "AUTH_PLUGIN_UNSPECIFIED"
 *   - 1: "MYSQL_NATIVE_PASSWORD"
 *   - 2: "CACHING_SHA2_PASSWORD"
 *   - 3: "SHA256_PASSWORD"
 *
 * * `defaultTimeZone` text
 *
 * * `explicitDefaultsForTimestamp` boolean
 *
 * * `generalLog` boolean
 *
 * * `groupConcatMaxLen` integer
 *
 * * `innodbAdaptiveHashIndex` boolean
 *
 * * `innodbBufferPoolSize` integer
 *
 * * `innodbFlushLogAtTrxCommit` integer
 *
 * * `innodbIoCapacity` integer
 *
 * * `innodbIoCapacityMax` integer
 *
 * * `innodbLockWaitTimeout` integer
 *
 * * `innodbLogBufferSize` integer
 *
 * * `innodbLogFileSize` integer
 *
 * * `innodbNumaInterleave` boolean
 *
 * * `innodbPrintAllDeadlocks` boolean
 *
 * * `innodbPurgeThreads` integer
 *
 * * `innodbReadIoThreads` integer
 *
 * * `innodbTempDataFileMaxSize` integer
 *
 * * `innodbThreadConcurrency` integer
 *
 * * `innodbWriteIoThreads` integer
 *
 * * `joinBufferSize` integer
 *
 * * `longQueryTime` float
 *
 * * `maxAllowedPacket` integer
 *
 * * `maxConnections` integer
 *
 * * `maxHeapTableSize` integer
 *
 * * `netReadTimeout` integer
 *
 * * `netWriteTimeout` integer
 *
 * * `rplSemiSyncMasterWaitForSlaveCount` integer
 *
 * * `slaveParallelType` one of:
 *   - 0: "SLAVE_PARALLEL_TYPE_UNSPECIFIED"
 *   - 1: "DATABASE"
 *   - 2: "LOGICAL_CLOCK"
 *
 * * `slaveParallelWorkers` integer
 *
 * * `sortBufferSize` integer
 *
 * * `syncBinlog` integer
 *
 * * `tableDefinitionCache` integer
 *
 * * `tableOpenCache` integer
 *
 * * `tableOpenCacheInstances` integer
 *
 * * `threadCacheSize` integer
 *
 * * `threadStack` integer
 *
 * * `tmpTableSize` integer
 *
 * * `transactionIsolation` one of:
 *   - 0: "TRANSACTION_ISOLATION_UNSPECIFIED"
 *   - 1: "READ_COMMITTED"
 *   - 2: "REPEATABLE_READ"
 *   - 3: "SERIALIZABLE"
 *
 * ## Import
 *
 * A cluster can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/mdbMysqlCluster:MdbMysqlCluster foo cluster_id
 * ```
 */
export class MdbMysqlCluster extends pulumi.CustomResource {
    /**
     * Get an existing MdbMysqlCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MdbMysqlClusterState, opts?: pulumi.CustomResourceOptions): MdbMysqlCluster {
        return new MdbMysqlCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/mdbMysqlCluster:MdbMysqlCluster';

    /**
     * Returns true if the given object is an instance of MdbMysqlCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MdbMysqlCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MdbMysqlCluster.__pulumiType;
    }

    /**
     * Access policy to the MySQL cluster. The structure is documented below.
     */
    public readonly access!: pulumi.Output<outputs.MdbMysqlClusterAccess>;
    /**
     * @deprecated You can safely remove this option. There is no need to recreate host if assign_public_ip is changed.
     */
    public readonly allowRegenerationHost!: pulumi.Output<boolean | undefined>;
    /**
     * Time to start the daily backup, in the UTC. The structure is documented below.
     */
    public readonly backupWindowStart!: pulumi.Output<outputs.MdbMysqlClusterBackupWindowStart>;
    /**
     * Creation timestamp of the cluster.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A database of the MySQL cluster. The structure is documented below.
     */
    public readonly databases!: pulumi.Output<outputs.MdbMysqlClusterDatabase[]>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * Description of the MySQL cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Deployment environment of the MySQL cluster.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Aggregated health of the cluster.
     */
    public /*out*/ readonly health!: pulumi.Output<string>;
    /**
     * A host of the MySQL cluster. The structure is documented below.
     */
    public readonly hosts!: pulumi.Output<outputs.MdbMysqlClusterHost[]>;
    /**
     * A set of key/value label pairs to assign to the MySQL cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Maintenance policy of the MySQL cluster. The structure is documented below.
     */
    public readonly maintenanceWindow!: pulumi.Output<outputs.MdbMysqlClusterMaintenanceWindow>;
    /**
     * MySQL cluster config. Detail info in "MySQL config" section (documented below).
     */
    public readonly mysqlConfig!: pulumi.Output<{[key: string]: string}>;
    /**
     * Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replicationSourceName` parameter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network, to which the MySQL cluster uses.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * Resources allocated to hosts of the MySQL cluster. The structure is documented below.
     */
    public readonly resources!: pulumi.Output<outputs.MdbMysqlClusterResources>;
    /**
     * The cluster will be created from the specified backup. The structure is documented below.
     */
    public readonly restore!: pulumi.Output<outputs.MdbMysqlClusterRestore | undefined>;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Status of the cluster.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A user of the MySQL cluster. The structure is documented below.
     */
    public readonly users!: pulumi.Output<outputs.MdbMysqlClusterUser[]>;
    /**
     * Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a MdbMysqlCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MdbMysqlClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MdbMysqlClusterArgs | MdbMysqlClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MdbMysqlClusterState | undefined;
            resourceInputs["access"] = state ? state.access : undefined;
            resourceInputs["allowRegenerationHost"] = state ? state.allowRegenerationHost : undefined;
            resourceInputs["backupWindowStart"] = state ? state.backupWindowStart : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["databases"] = state ? state.databases : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["health"] = state ? state.health : undefined;
            resourceInputs["hosts"] = state ? state.hosts : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maintenanceWindow"] = state ? state.maintenanceWindow : undefined;
            resourceInputs["mysqlConfig"] = state ? state.mysqlConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
            resourceInputs["restore"] = state ? state.restore : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as MdbMysqlClusterArgs | undefined;
            if ((!args || args.databases === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databases'");
            }
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.hosts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hosts'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.resources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resources'");
            }
            if ((!args || args.users === undefined) && !opts.urn) {
                throw new Error("Missing required property 'users'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["access"] = args ? args.access : undefined;
            resourceInputs["allowRegenerationHost"] = args ? args.allowRegenerationHost : undefined;
            resourceInputs["backupWindowStart"] = args ? args.backupWindowStart : undefined;
            resourceInputs["databases"] = args ? args.databases : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["hosts"] = args ? args.hosts : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            resourceInputs["mysqlConfig"] = args ? args.mysqlConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["restore"] = args ? args.restore : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["health"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MdbMysqlCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MdbMysqlCluster resources.
 */
export interface MdbMysqlClusterState {
    /**
     * Access policy to the MySQL cluster. The structure is documented below.
     */
    access?: pulumi.Input<inputs.MdbMysqlClusterAccess>;
    /**
     * @deprecated You can safely remove this option. There is no need to recreate host if assign_public_ip is changed.
     */
    allowRegenerationHost?: pulumi.Input<boolean>;
    /**
     * Time to start the daily backup, in the UTC. The structure is documented below.
     */
    backupWindowStart?: pulumi.Input<inputs.MdbMysqlClusterBackupWindowStart>;
    /**
     * Creation timestamp of the cluster.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A database of the MySQL cluster. The structure is documented below.
     */
    databases?: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterDatabase>[]>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Description of the MySQL cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * Deployment environment of the MySQL cluster.
     */
    environment?: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Aggregated health of the cluster.
     */
    health?: pulumi.Input<string>;
    /**
     * A host of the MySQL cluster. The structure is documented below.
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterHost>[]>;
    /**
     * A set of key/value label pairs to assign to the MySQL cluster.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Maintenance policy of the MySQL cluster. The structure is documented below.
     */
    maintenanceWindow?: pulumi.Input<inputs.MdbMysqlClusterMaintenanceWindow>;
    /**
     * MySQL cluster config. Detail info in "MySQL config" section (documented below).
     */
    mysqlConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replicationSourceName` parameter.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network, to which the MySQL cluster uses.
     */
    networkId?: pulumi.Input<string>;
    /**
     * Resources allocated to hosts of the MySQL cluster. The structure is documented below.
     */
    resources?: pulumi.Input<inputs.MdbMysqlClusterResources>;
    /**
     * The cluster will be created from the specified backup. The structure is documented below.
     */
    restore?: pulumi.Input<inputs.MdbMysqlClusterRestore>;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the cluster.
     */
    status?: pulumi.Input<string>;
    /**
     * A user of the MySQL cluster. The structure is documented below.
     */
    users?: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterUser>[]>;
    /**
     * Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MdbMysqlCluster resource.
 */
export interface MdbMysqlClusterArgs {
    /**
     * Access policy to the MySQL cluster. The structure is documented below.
     */
    access?: pulumi.Input<inputs.MdbMysqlClusterAccess>;
    /**
     * @deprecated You can safely remove this option. There is no need to recreate host if assign_public_ip is changed.
     */
    allowRegenerationHost?: pulumi.Input<boolean>;
    /**
     * Time to start the daily backup, in the UTC. The structure is documented below.
     */
    backupWindowStart?: pulumi.Input<inputs.MdbMysqlClusterBackupWindowStart>;
    /**
     * A database of the MySQL cluster. The structure is documented below.
     */
    databases: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterDatabase>[]>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Description of the MySQL cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * Deployment environment of the MySQL cluster.
     */
    environment: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A host of the MySQL cluster. The structure is documented below.
     */
    hosts: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterHost>[]>;
    /**
     * A set of key/value label pairs to assign to the MySQL cluster.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Maintenance policy of the MySQL cluster. The structure is documented below.
     */
    maintenanceWindow?: pulumi.Input<inputs.MdbMysqlClusterMaintenanceWindow>;
    /**
     * MySQL cluster config. Detail info in "MySQL config" section (documented below).
     */
    mysqlConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replicationSourceName` parameter.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network, to which the MySQL cluster uses.
     */
    networkId: pulumi.Input<string>;
    /**
     * Resources allocated to hosts of the MySQL cluster. The structure is documented below.
     */
    resources: pulumi.Input<inputs.MdbMysqlClusterResources>;
    /**
     * The cluster will be created from the specified backup. The structure is documented below.
     */
    restore?: pulumi.Input<inputs.MdbMysqlClusterRestore>;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A user of the MySQL cluster. The structure is documented below.
     */
    users: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterUser>[]>;
    /**
     * Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
     */
    version: pulumi.Input<string>;
}
