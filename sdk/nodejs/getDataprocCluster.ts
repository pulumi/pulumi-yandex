// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Data Proc cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/data-proc/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = pulumi.output(yandex.getDataprocCluster({
 *     name: "test",
 * }));
 *
 * export const serviceAccountId = foo.serviceAccountId;
 * ```
 */
export function getDataprocCluster(args?: GetDataprocClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetDataprocClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getDataprocCluster:getDataprocCluster", {
        "clusterId": args.clusterId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDataprocCluster.
 */
export interface GetDataprocClusterArgs {
    /**
     * The ID of the Data Proc cluster.
     */
    clusterId?: string;
    /**
     * The name of the Data Proc cluster.
     */
    name?: string;
}

/**
 * A collection of values returned by getDataprocCluster.
 */
export interface GetDataprocClusterResult {
    /**
     * Name of the Object Storage bucket used for Data Proc jobs.
     */
    readonly bucket: string;
    /**
     * Configuration and resources of the cluster. The structure is documented below.
     */
    readonly clusterConfigs: outputs.GetDataprocClusterClusterConfig[];
    readonly clusterId: string;
    /**
     * The Data Proc cluster creation timestamp.
     */
    readonly createdAt: string;
    readonly deletionProtection: boolean;
    /**
     * Description of the Data Proc cluster.
     */
    readonly description: string;
    readonly folderId: string;
    /**
     * A list of IDs of the host groups hosting VMs of the cluster.
     */
    readonly hostGroupIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of key/value label pairs assigned to the Data Proc cluster.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the Data Proc subcluster.
     */
    readonly name: string;
    readonly securityGroupIds: string[];
    /**
     * Service account used by the Data Proc agent to access resources of Yandex.Cloud.
     */
    readonly serviceAccountId: string;
    /**
     * Whether UI proxy feature is enabled.
     */
    readonly uiProxy: boolean;
    /**
     * ID of the availability zone where the cluster resides.
     */
    readonly zoneId: string;
}

export function getDataprocClusterOutput(args?: GetDataprocClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataprocClusterResult> {
    return pulumi.output(args).apply(a => getDataprocCluster(a, opts))
}

/**
 * A collection of arguments for invoking getDataprocCluster.
 */
export interface GetDataprocClusterOutputArgs {
    /**
     * The ID of the Data Proc cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The name of the Data Proc cluster.
     */
    name?: pulumi.Input<string>;
}
