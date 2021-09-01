// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Managed MongoDB cluster. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = pulumi.output(yandex.getMdbMongodbCluster({
 *     name: "test",
 * }, { async: true }));
 *
 * export const networkId = foo.networkId;
 * ```
 */
export function getMdbMongodbCluster(args?: GetMdbMongodbClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetMdbMongodbClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getMdbMongodbCluster:getMdbMongodbCluster", {
        "clusterId": args.clusterId,
        "deletionProtection": args.deletionProtection,
        "folderId": args.folderId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getMdbMongodbCluster.
 */
export interface GetMdbMongodbClusterArgs {
    /**
     * The ID of the MongoDB cluster.
     */
    readonly clusterId?: string;
    readonly deletionProtection?: boolean;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    readonly folderId?: string;
    /**
     * The name of the MongoDB cluster.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getMdbMongodbCluster.
 */
export interface GetMdbMongodbClusterResult {
    /**
     * Configuration of the MongoDB cluster. The structure is documented below.
     */
    readonly clusterConfig: outputs.GetMdbMongodbClusterClusterConfig;
    readonly clusterId: string;
    /**
     * Creation timestamp of the key.
     */
    readonly createdAt: string;
    /**
     * A database of the MongoDB cluster. The structure is documented below.
     */
    readonly databases: outputs.GetMdbMongodbClusterDatabase[];
    readonly deletionProtection: boolean;
    /**
     * Description of the MongoDB cluster.
     */
    readonly description: string;
    /**
     * Deployment environment of the MongoDB cluster.
     */
    readonly environment: string;
    readonly folderId: string;
    /**
     * The health of the host.
     */
    readonly health: string;
    /**
     * A host of the MongoDB cluster. The structure is documented below.
     */
    readonly hosts: outputs.GetMdbMongodbClusterHost[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of key/value label pairs to assign to the MongoDB cluster.
     */
    readonly labels: {[key: string]: string};
    readonly maintenanceWindow: outputs.GetMdbMongodbClusterMaintenanceWindow;
    /**
     * The name of the database.
     */
    readonly name: string;
    /**
     * ID of the network, to which the MongoDB cluster belongs.
     */
    readonly networkId: string;
    /**
     * Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
     */
    readonly resources: outputs.GetMdbMongodbClusterResources;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    readonly securityGroupIds: string[];
    /**
     * MongoDB Cluster mode enabled/disabled.
     */
    readonly sharded: boolean;
    /**
     * Status of the cluster.
     */
    readonly status: string;
    /**
     * A user of the MongoDB cluster. The structure is documented below.
     */
    readonly users: outputs.GetMdbMongodbClusterUser[];
}
