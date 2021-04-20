// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Managed ClickHouse cluster. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = pulumi.output(yandex.getMdbClickhouseCluster({
 *     name: "test",
 * }, { async: true }));
 *
 * export const networkId = foo.networkId;
 * ```
 */
export function getMdbClickhouseCluster(args?: GetMdbClickhouseClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetMdbClickhouseClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getMdbClickhouseCluster:getMdbClickhouseCluster", {
        "cloudStorage": args.cloudStorage,
        "clusterId": args.clusterId,
        "folderId": args.folderId,
        "name": args.name,
        "serviceAccountId": args.serviceAccountId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMdbClickhouseCluster.
 */
export interface GetMdbClickhouseClusterArgs {
    readonly cloudStorage?: inputs.GetMdbClickhouseClusterCloudStorage;
    /**
     * The ID of the ClickHouse cluster.
     */
    readonly clusterId?: string;
    /**
     * The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
     */
    readonly folderId?: string;
    /**
     * The name of the ClickHouse cluster.
     */
    readonly name?: string;
    readonly serviceAccountId?: string;
}

/**
 * A collection of values returned by getMdbClickhouseCluster.
 */
export interface GetMdbClickhouseClusterResult {
    /**
     * Access policy to the ClickHouse cluster. The structure is documented below.
     */
    readonly access: outputs.GetMdbClickhouseClusterAccess;
    /**
     * Time to start the daily backup, in the UTC timezone. The structure is documented below.
     */
    readonly backupWindowStart: outputs.GetMdbClickhouseClusterBackupWindowStart;
    /**
     * Configuration of the ClickHouse subcluster. The structure is documented below.
     */
    readonly clickhouse: outputs.GetMdbClickhouseClusterClickhouse;
    readonly cloudStorage?: outputs.GetMdbClickhouseClusterCloudStorage;
    readonly clusterId: string;
    /**
     * Creation timestamp of the key.
     */
    readonly createdAt: string;
    /**
     * A database of the ClickHouse cluster. The structure is documented below.
     */
    readonly databases: outputs.GetMdbClickhouseClusterDatabase[];
    /**
     * Description of the shard group.
     */
    readonly description: string;
    /**
     * Deployment environment of the ClickHouse cluster.
     */
    readonly environment: string;
    readonly folderId: string;
    /**
     * A set of protobuf or cap'n proto format schemas. The structure is documented below.
     */
    readonly formatSchemas: outputs.GetMdbClickhouseClusterFormatSchema[];
    /**
     * Aggregated health of the cluster.
     */
    readonly health: string;
    /**
     * A host of the ClickHouse cluster. The structure is documented below.
     */
    readonly hosts: outputs.GetMdbClickhouseClusterHost[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of key/value label pairs to assign to the ClickHouse cluster.
     */
    readonly labels: {[key: string]: string};
    /**
     * A group of machine learning models. The structure is documented below.
     */
    readonly mlModels: outputs.GetMdbClickhouseClusterMlModel[];
    /**
     * Graphite rollup configuration name.
     */
    readonly name: string;
    /**
     * ID of the network, to which the ClickHouse cluster belongs.
     */
    readonly networkId: string;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    readonly securityGroupIds: string[];
    readonly serviceAccountId: string;
    /**
     * A group of clickhouse shards. The structure is documented below.
     */
    readonly shardGroups: outputs.GetMdbClickhouseClusterShardGroup[];
    /**
     * Grants `admin` user database management permission.
     */
    readonly sqlDatabaseManagement: boolean;
    /**
     * Enables `admin` user with user management permission.
     */
    readonly sqlUserManagement: boolean;
    /**
     * Status of the cluster.
     */
    readonly status: string;
    /**
     * A user of the ClickHouse cluster. The structure is documented below.
     */
    readonly users: outputs.GetMdbClickhouseClusterUser[];
    readonly version: string;
    /**
     * Configuration of the ZooKeeper subcluster. The structure is documented below.
     */
    readonly zookeeper: outputs.GetMdbClickhouseClusterZookeeper;
}
