// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Application Load Balancer HTTP Router. For more information, see
 * [Yandex.Cloud Application Load Balancer](https://cloud.yandex.com/en/docs/application-load-balancer/quickstart).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const tf_router = pulumi.output(yandex.getAlbHttpRouter({
 *     httpRouterId: "my-http-router-id",
 * }));
 * ```
 *
 * This data source is used to define [Application Load Balancer HTTP Router] that can be used by other resources.
 */
export function getAlbHttpRouter(args?: GetAlbHttpRouterArgs, opts?: pulumi.InvokeOptions): Promise<GetAlbHttpRouterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getAlbHttpRouter:getAlbHttpRouter", {
        "description": args.description,
        "folderId": args.folderId,
        "httpRouterId": args.httpRouterId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlbHttpRouter.
 */
export interface GetAlbHttpRouterArgs {
    /**
     * Description of the HTTP Router.
     */
    description?: string;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: string;
    /**
     * HTTP Router ID.
     */
    httpRouterId?: string;
    /**
     * - Name of the HTTP Router.
     */
    name?: string;
}

/**
 * A collection of values returned by getAlbHttpRouter.
 */
export interface GetAlbHttpRouterResult {
    /**
     * Creation timestamp of this HTTP Router.
     */
    readonly createdAt: string;
    /**
     * Description of the HTTP Router.
     */
    readonly description: string;
    readonly folderId: string;
    readonly httpRouterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Labels to assign to this HTTP Router.
     */
    readonly labels: {[key: string]: string};
    readonly name: string;
}

export function getAlbHttpRouterOutput(args?: GetAlbHttpRouterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAlbHttpRouterResult> {
    return pulumi.output(args).apply(a => getAlbHttpRouter(a, opts))
}

/**
 * A collection of arguments for invoking getAlbHttpRouter.
 */
export interface GetAlbHttpRouterOutputArgs {
    /**
     * Description of the HTTP Router.
     */
    description?: pulumi.Input<string>;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * HTTP Router ID.
     */
    httpRouterId?: pulumi.Input<string>;
    /**
     * - Name of the HTTP Router.
     */
    name?: pulumi.Input<string>;
}
