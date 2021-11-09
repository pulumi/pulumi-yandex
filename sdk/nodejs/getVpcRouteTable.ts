// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex VPC route table. For more information, see
 * [Yandex.Cloud VPC](https://cloud.yandex.com/docs/vpc/concepts/index).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const thisVpcRouteTable = pulumi.output(yandex.getVpcRouteTable({
 *     routeTableId: "my-rt-id",
 * }));
 * ```
 *
 * This data source is used to define [VPC Route Table] that can be used by other resources.
 */
export function getVpcRouteTable(args?: GetVpcRouteTableArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcRouteTableResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getVpcRouteTable:getVpcRouteTable", {
        "folderId": args.folderId,
        "name": args.name,
        "routeTableId": args.routeTableId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcRouteTable.
 */
export interface GetVpcRouteTableArgs {
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: string;
    /**
     * - Name of the route table.
     */
    name?: string;
    /**
     * Route table ID.
     */
    routeTableId?: string;
}

/**
 * A collection of values returned by getVpcRouteTable.
 */
export interface GetVpcRouteTableResult {
    /**
     * Creation timestamp of this route table.
     */
    readonly createdAt: string;
    /**
     * Description of the route table.
     */
    readonly description: string;
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Labels to assign to this route table.
     */
    readonly labels: {[key: string]: string};
    readonly name: string;
    /**
     * ID of the network this route table belongs to.
     */
    readonly networkId: string;
    readonly routeTableId: string;
    /**
     * List of static route records of the route table. Structure is documented below.
     */
    readonly staticRoutes: outputs.GetVpcRouteTableStaticRoute[];
}

export function getVpcRouteTableOutput(args?: GetVpcRouteTableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcRouteTableResult> {
    return pulumi.output(args).apply(a => getVpcRouteTable(a, opts))
}

/**
 * A collection of arguments for invoking getVpcRouteTable.
 */
export interface GetVpcRouteTableOutputArgs {
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * - Name of the route table.
     */
    name?: pulumi.Input<string>;
    /**
     * Route table ID.
     */
    routeTableId?: pulumi.Input<string>;
}
