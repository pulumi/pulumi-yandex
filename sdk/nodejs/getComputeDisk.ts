// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Compute disk. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/compute/concepts/disk).
 */
export function getComputeDisk(args?: GetComputeDiskArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeDiskResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getComputeDisk:getComputeDisk", {
        "diskId": args.diskId,
        "diskPlacementPolicy": args.diskPlacementPolicy,
        "folderId": args.folderId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeDisk.
 */
export interface GetComputeDiskArgs {
    /**
     * The ID of a specific disk.
     */
    diskId?: string;
    diskPlacementPolicy?: inputs.GetComputeDiskDiskPlacementPolicy;
    /**
     * ID of the folder that the disk belongs to.
     */
    folderId?: string;
    /**
     * Name of the disk.
     */
    name?: string;
}

/**
 * A collection of values returned by getComputeDisk.
 */
export interface GetComputeDiskResult {
    /**
     * Disk creation timestamp.
     */
    readonly createdAt: string;
    /**
     * Optional description of this disk.
     */
    readonly description: string;
    readonly diskId: string;
    readonly diskPlacementPolicy?: outputs.GetComputeDiskDiskPlacementPolicy;
    /**
     * ID of the folder that the disk belongs to.
     */
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * ID of the source image that was used to create this disk.
     */
    readonly imageId: string;
    /**
     * IDs of instances to which this disk is attached.
     */
    readonly instanceIds: string[];
    /**
     * Map of labels applied to this disk.
     */
    readonly labels: {[key: string]: string};
    readonly name: string;
    /**
     * License IDs that indicate which licenses are attached to this disk.
     */
    readonly productIds: string[];
    /**
     * Size of the disk, specified in Gb.
     */
    readonly size: number;
    /**
     * Source snapshot that was used to create this disk.
     */
    readonly snapshotId: string;
    /**
     * Status of the disk.
     */
    readonly status: string;
    /**
     * Type of the disk.
     */
    readonly type: string;
    /**
     * ID of the zone where the disk resides.
     */
    readonly zone: string;
}

export function getComputeDiskOutput(args?: GetComputeDiskOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeDiskResult> {
    return pulumi.output(args).apply(a => getComputeDisk(a, opts))
}

/**
 * A collection of arguments for invoking getComputeDisk.
 */
export interface GetComputeDiskOutputArgs {
    /**
     * The ID of a specific disk.
     */
    diskId?: pulumi.Input<string>;
    diskPlacementPolicy?: pulumi.Input<inputs.GetComputeDiskDiskPlacementPolicyArgs>;
    /**
     * ID of the folder that the disk belongs to.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Name of the disk.
     */
    name?: pulumi.Input<string>;
}
