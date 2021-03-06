// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information about a Yandex Resource Manager Folder. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/resource-manager/concepts/resources-hierarchy#folder).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const myFolder1 = pulumi.output(yandex.getResourcemanagerFolder({
 *     folderId: "folder_id_number_1",
 * }));
 * const myFolder2 = pulumi.output(yandex.getResourcemanagerFolder({
 *     cloudId: "some_cloud_id",
 *     name: "folder_name",
 * }));
 *
 * export const myFolder1Name = myFolder1.name!;
 * export const myFolder2CloudId = myFolder2.cloudId!;
 * ```
 */
export function getResourcemanagerFolder(args?: GetResourcemanagerFolderArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcemanagerFolderResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("yandex:index/getResourcemanagerFolder:getResourcemanagerFolder", {
        "cloudId": args.cloudId,
        "folderId": args.folderId,
        "labels": args.labels,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourcemanagerFolder.
 */
export interface GetResourcemanagerFolderArgs {
    /**
     * Cloud that the resource belongs to. If value is omitted, the default provider cloud is used.
     */
    cloudId?: string;
    /**
     * ID of the folder.
     */
    folderId?: string;
    /**
     * A map of labels applied to this folder.
     */
    labels?: {[key: string]: string};
    /**
     * Name of the folder.
     */
    name?: string;
}

/**
 * A collection of values returned by getResourcemanagerFolder.
 */
export interface GetResourcemanagerFolderResult {
    /**
     * ID of the cloud that contains the folder.
     */
    readonly cloudId: string;
    /**
     * Folder creation timestamp.
     */
    readonly createdAt: string;
    /**
     * Description of the folder.
     */
    readonly description: string;
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A map of labels applied to this folder.
     */
    readonly labels?: {[key: string]: string};
    readonly name: string;
    /**
     * Current status of the folder.
     */
    readonly status: string;
}

export function getResourcemanagerFolderOutput(args?: GetResourcemanagerFolderOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourcemanagerFolderResult> {
    return pulumi.output(args).apply(a => getResourcemanagerFolder(a, opts))
}

/**
 * A collection of arguments for invoking getResourcemanagerFolder.
 */
export interface GetResourcemanagerFolderOutputArgs {
    /**
     * Cloud that the resource belongs to. If value is omitted, the default provider cloud is used.
     */
    cloudId?: pulumi.Input<string>;
    /**
     * ID of the folder.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A map of labels applied to this folder.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the folder.
     */
    name?: pulumi.Input<string>;
}
