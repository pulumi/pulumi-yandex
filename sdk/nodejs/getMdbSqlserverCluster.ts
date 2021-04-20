// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Managed SQLServer cluster. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/managed-sqlserver/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = pulumi.output(yandex.getMdbSqlserverCluster({
 *     name: "test",
 * }, { async: true }));
 *
 * export const networkId = foo.networkId;
 * ```
 */
export function getMdbSqlserverCluster(args?: GetMdbSqlserverClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetMdbSqlserverClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getMdbSqlserverCluster:getMdbSqlserverCluster", {
        "clusterId": args.clusterId,
        "folderId": args.folderId,
        "name": args.name,
        "sqlserverConfig": args.sqlserverConfig,
    }, opts);
}

/**
 * A collection of arguments for invoking getMdbSqlserverCluster.
 */
export interface GetMdbSqlserverClusterArgs {
    /**
     * The ID of the SQLServer cluster.
     */
    readonly clusterId?: string;
    /**
     * The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
     */
    readonly folderId?: string;
    /**
     * The name of the SQLServer cluster.
     */
    readonly name?: string;
    /**
     * SQLServer cluster config.
     */
    readonly sqlserverConfig?: {[key: string]: string};
}

/**
 * A collection of values returned by getMdbSqlserverCluster.
 */
export interface GetMdbSqlserverClusterResult {
    readonly backupWindowStarts: outputs.GetMdbSqlserverClusterBackupWindowStart[];
    readonly clusterId: string;
    /**
     * Creation timestamp of the key.
     */
    readonly createdAt: string;
    /**
     * A database of the SQLServer cluster. The structure is documented below.
     */
    readonly databases: outputs.GetMdbSqlserverClusterDatabase[];
    /**
     * Description of the SQLServer cluster.
     */
    readonly description: string;
    /**
     * Deployment environment of the SQLServer cluster.
     */
    readonly environment: string;
    readonly folderId: string;
    /**
     * Aggregated health of the cluster.
     */
    readonly health: string;
    /**
     * A host of the SQLServer cluster. The structure is documented below.
     */
    readonly hosts: outputs.GetMdbSqlserverClusterHost[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of key/value label pairs to assign to the SQLServer cluster.
     */
    readonly labels: {[key: string]: string};
    /**
     * The name of the database.
     */
    readonly name: string;
    /**
     * ID of the network, to which the SQLServer cluster belongs.
     */
    readonly networkId: string;
    /**
     * Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
     */
    readonly resources: outputs.GetMdbSqlserverClusterResource[];
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    readonly securityGroupIds: string[];
    /**
     * SQLServer cluster config.
     */
    readonly sqlserverConfig: {[key: string]: string};
    /**
     * Status of the cluster.
     */
    readonly status: string;
    /**
     * A user of the SQLServer cluster. The structure is documented below.
     */
    readonly users: outputs.GetMdbSqlserverClusterUser[];
    /**
     * Version of the SQLServer cluster.
     */
    readonly version: string;
}
