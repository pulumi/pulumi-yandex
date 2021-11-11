// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getMdbGreenplumCluster(args?: GetMdbGreenplumClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetMdbGreenplumClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getMdbGreenplumCluster:getMdbGreenplumCluster", {
        "clusterId": args.clusterId,
        "folderId": args.folderId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getMdbGreenplumCluster.
 */
export interface GetMdbGreenplumClusterArgs {
    clusterId?: string;
    folderId?: string;
    name?: string;
}

/**
 * A collection of values returned by getMdbGreenplumCluster.
 */
export interface GetMdbGreenplumClusterResult {
    readonly assignPublicIp: boolean;
    readonly clusterId: string;
    readonly createdAt: string;
    readonly deletionProtection: boolean;
    readonly description: string;
    readonly environment: string;
    readonly folderId: string;
    readonly health: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: {[key: string]: string};
    readonly masterHostCount: number;
    readonly masterHosts: outputs.GetMdbGreenplumClusterMasterHost[];
    readonly masterSubcluster: outputs.GetMdbGreenplumClusterMasterSubcluster;
    readonly name: string;
    readonly networkId: string;
    readonly securityGroupIds: string[];
    readonly segmentHostCount: number;
    readonly segmentHosts: outputs.GetMdbGreenplumClusterSegmentHost[];
    readonly segmentInHost: number;
    readonly segmentSubcluster: outputs.GetMdbGreenplumClusterSegmentSubcluster;
    readonly status: string;
    readonly subnetId: string;
    readonly userName: string;
    readonly version: string;
    readonly zone: string;
}

export function getMdbGreenplumClusterOutput(args?: GetMdbGreenplumClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMdbGreenplumClusterResult> {
    return pulumi.output(args).apply(a => getMdbGreenplumCluster(a, opts))
}

/**
 * A collection of arguments for invoking getMdbGreenplumCluster.
 */
export interface GetMdbGreenplumClusterOutputArgs {
    clusterId?: pulumi.Input<string>;
    folderId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
