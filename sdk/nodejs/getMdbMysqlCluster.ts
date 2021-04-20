// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Managed MySQL cluster. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = pulumi.output(yandex.getMdbMysqlCluster({
 *     name: "test",
 * }, { async: true }));
 *
 * export const networkId = foo.networkId;
 * ```
 */
export function getMdbMysqlCluster(args?: GetMdbMysqlClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetMdbMysqlClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getMdbMysqlCluster:getMdbMysqlCluster", {
        "access": args.access,
        "clusterId": args.clusterId,
        "description": args.description,
        "folderId": args.folderId,
        "labels": args.labels,
        "mysqlConfig": args.mysqlConfig,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getMdbMysqlCluster.
 */
export interface GetMdbMysqlClusterArgs {
    /**
     * Access policy to the MySQL cluster. The structure is documented below.
     */
    readonly access?: inputs.GetMdbMysqlClusterAccess;
    /**
     * The ID of the MySQL cluster.
     */
    readonly clusterId?: string;
    /**
     * Description of the MySQL cluster.
     */
    readonly description?: string;
    /**
     * The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
     */
    readonly folderId?: string;
    /**
     * A set of key/value label pairs to assign to the MySQL cluster.
     */
    readonly labels?: {[key: string]: string};
    /**
     * MySQL cluster config.
     */
    readonly mysqlConfig?: {[key: string]: string};
    /**
     * The name of the MySQL cluster.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getMdbMysqlCluster.
 */
export interface GetMdbMysqlClusterResult {
    /**
     * Access policy to the MySQL cluster. The structure is documented below.
     */
    readonly access: outputs.GetMdbMysqlClusterAccess;
    readonly backupWindowStart: outputs.GetMdbMysqlClusterBackupWindowStart;
    readonly clusterId: string;
    /**
     * Creation timestamp of the key.
     */
    readonly createdAt: string;
    /**
     * A database of the MySQL cluster. The structure is documented below.
     */
    readonly databases: outputs.GetMdbMysqlClusterDatabase[];
    /**
     * Description of the MySQL cluster.
     */
    readonly description?: string;
    /**
     * Deployment environment of the MySQL cluster.
     */
    readonly environment: string;
    readonly folderId: string;
    /**
     * Aggregated health of the cluster.
     */
    readonly health: string;
    /**
     * A host of the MySQL cluster. The structure is documented below.
     */
    readonly hosts: outputs.GetMdbMysqlClusterHost[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of key/value label pairs to assign to the MySQL cluster.
     */
    readonly labels?: {[key: string]: string};
    /**
     * MySQL cluster config.
     */
    readonly mysqlConfig: {[key: string]: string};
    /**
     * The name of the database.
     */
    readonly name: string;
    /**
     * ID of the network, to which the MySQL cluster belongs.
     */
    readonly networkId: string;
    /**
     * Resources allocated to hosts of the MySQL cluster. The structure is documented below.
     */
    readonly resources: outputs.GetMdbMysqlClusterResources;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    readonly securityGroupIds: string[];
    /**
     * Status of the cluster.
     */
    readonly status: string;
    /**
     * A user of the MySQL cluster. The structure is documented below.
     */
    readonly users: outputs.GetMdbMysqlClusterUser[];
    /**
     * Version of the MySQL cluster.
     */
    readonly version: string;
}
