// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Compute snapshot. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/compute/concepts/snapshot).
 */
export function getComputeSnapshot(args?: GetComputeSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeSnapshotResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("yandex:index/getComputeSnapshot:getComputeSnapshot", {
        "folderId": args.folderId,
        "name": args.name,
        "snapshotId": args.snapshotId,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeSnapshot.
 */
export interface GetComputeSnapshotArgs {
    /**
     * ID of the folder that the snapshot belongs to.
     */
    folderId?: string;
    /**
     * The name of the snapshot.
     */
    name?: string;
    /**
     * The ID of a specific snapshot.
     */
    snapshotId?: string;
}

/**
 * A collection of values returned by getComputeSnapshot.
 */
export interface GetComputeSnapshotResult {
    /**
     * Snapshot creation timestamp.
     */
    readonly createdAt: string;
    /**
     * An optional description of this snapshot.
     */
    readonly description: string;
    /**
     * Minimum required size of the disk which is created from this snapshot.
     */
    readonly diskSize: number;
    /**
     * ID of the folder that the snapshot belongs to.
     */
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A map of labels applied to this snapshot.
     */
    readonly labels: {[key: string]: string};
    readonly name: string;
    /**
     * License IDs that indicate which licenses are attached to this snapshot.
     */
    readonly productIds: string[];
    readonly snapshotId: string;
    /**
     * ID of the source disk.
     */
    readonly sourceDiskId: string;
    /**
     * The status of the snapshot.
     */
    readonly status: string;
    /**
     * The size of the snapshot, specified in Gb.
     */
    readonly storageSize: number;
}

export function getComputeSnapshotOutput(args?: GetComputeSnapshotOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeSnapshotResult> {
    return pulumi.output(args).apply(a => getComputeSnapshot(a, opts))
}

/**
 * A collection of arguments for invoking getComputeSnapshot.
 */
export interface GetComputeSnapshotOutputArgs {
    /**
     * ID of the folder that the snapshot belongs to.
     */
    folderId?: pulumi.Input<string>;
    /**
     * The name of the snapshot.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of a specific snapshot.
     */
    snapshotId?: pulumi.Input<string>;
}
