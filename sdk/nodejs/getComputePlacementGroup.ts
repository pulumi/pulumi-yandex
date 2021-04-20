// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Compute Placement group. For more information, see
 * [the official documentation](https://cloud.yandex.ru/docs/compute/concepts/placement-groups).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const myGroup = pulumi.output(yandex.getComputePlacementGroup({
 *     groupId: "some_group_id",
 * }, { async: true }));
 *
 * export const placementGroupName = myGroup.name!;
 * ```
 */
export function getComputePlacementGroup(args?: GetComputePlacementGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetComputePlacementGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getComputePlacementGroup:getComputePlacementGroup", {
        "description": args.description,
        "folderId": args.folderId,
        "groupId": args.groupId,
        "labels": args.labels,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputePlacementGroup.
 */
export interface GetComputePlacementGroupArgs {
    /**
     * Description of the group.
     */
    readonly description?: string;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    readonly folderId?: string;
    /**
     * The ID of a specific group.
     */
    readonly groupId?: string;
    /**
     * A set of key/value label pairs assigned to the group.
     */
    readonly labels?: {[key: string]: string};
    /**
     * Name of the group.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getComputePlacementGroup.
 */
export interface GetComputePlacementGroupResult {
    /**
     * Placement group creation timestamp.
     */
    readonly createdAt: string;
    /**
     * Description of the group.
     */
    readonly description?: string;
    readonly folderId: string;
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of key/value label pairs assigned to the group.
     */
    readonly labels?: {[key: string]: string};
    readonly name?: string;
}
