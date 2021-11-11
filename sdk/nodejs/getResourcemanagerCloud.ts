// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get cloud details.
 * For more information, see [Cloud](https://cloud.yandex.com/docs/resource-manager/concepts/resources-hierarchy#cloud).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = pulumi.output(yandex.getResourcemanagerCloud({
 *     name: "foo-cloud",
 * }));
 *
 * export const cloudCreateTimestamp = foo.createdAt;
 * ```
 */
export function getResourcemanagerCloud(args?: GetResourcemanagerCloudArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcemanagerCloudResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getResourcemanagerCloud:getResourcemanagerCloud", {
        "cloudId": args.cloudId,
        "description": args.description,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourcemanagerCloud.
 */
export interface GetResourcemanagerCloudArgs {
    /**
     * ID of the cloud.
     */
    cloudId?: string;
    /**
     * Description of the cloud.
     */
    description?: string;
    /**
     * Name of the cloud.
     */
    name?: string;
}

/**
 * A collection of values returned by getResourcemanagerCloud.
 */
export interface GetResourcemanagerCloudResult {
    readonly cloudId: string;
    /**
     * Cloud creation timestamp.
     */
    readonly createdAt: string;
    /**
     * Description of the cloud.
     */
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the cloud.
     */
    readonly name: string;
}

export function getResourcemanagerCloudOutput(args?: GetResourcemanagerCloudOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourcemanagerCloudResult> {
    return pulumi.output(args).apply(a => getResourcemanagerCloud(a, opts))
}

/**
 * A collection of arguments for invoking getResourcemanagerCloud.
 */
export interface GetResourcemanagerCloudOutputArgs {
    /**
     * ID of the cloud.
     */
    cloudId?: pulumi.Input<string>;
    /**
     * Description of the cloud.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the cloud.
     */
    name?: pulumi.Input<string>;
}
