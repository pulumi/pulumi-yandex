// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getComputeImage(args?: GetComputeImageArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeImageResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getComputeImage:getComputeImage", {
        "family": args.family,
        "folderId": args.folderId,
        "imageId": args.imageId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeImage.
 */
export interface GetComputeImageArgs {
    /**
     * The family name of an image. Used to search the latest image in a family.
     */
    readonly family?: string;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    readonly folderId?: string;
    /**
     * The ID of a specific image.
     */
    readonly imageId?: string;
    /**
     * The name of the image.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getComputeImage.
 */
export interface GetComputeImageResult {
    /**
     * Image creation timestamp.
     */
    readonly createdAt: string;
    /**
     * An optional description of this image.
     */
    readonly description: string;
    /**
     * The OS family name of the image.
     */
    readonly family: string;
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId: string;
    /**
     * A map of labels applied to this image.
     */
    readonly labels: {[key: string]: string};
    /**
     * Minimum size of the disk which is created from this image.
     */
    readonly minDiskSize: number;
    readonly name: string;
    /**
     * Operating system type that the image contains.
     */
    readonly osType: string;
    /**
     * License IDs that indicate which licenses are attached to this image.
     */
    readonly productIds: string[];
    /**
     * The size of the image, specified in Gb.
     */
    readonly size: number;
    /**
     * The status of the image.
     */
    readonly status: string;
}
