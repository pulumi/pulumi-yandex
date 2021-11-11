// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex VPC network. For more information, see
 * [Yandex.Cloud VPC](https://cloud.yandex.com/docs/vpc/concepts/index).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const admin = pulumi.output(yandex.getVpcNetwork({
 *     networkId: "my-network-id",
 * }));
 * ```
 *
 * This data source is used to define [VPC Networks] that can be used by other resources.
 */
export function getVpcNetwork(args?: GetVpcNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcNetworkResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getVpcNetwork:getVpcNetwork", {
        "folderId": args.folderId,
        "name": args.name,
        "networkId": args.networkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcNetwork.
 */
export interface GetVpcNetworkArgs {
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: string;
    /**
     * Name of the network.
     */
    name?: string;
    /**
     * ID of the network.
     */
    networkId?: string;
}

/**
 * A collection of values returned by getVpcNetwork.
 */
export interface GetVpcNetworkResult {
    /**
     * Creation timestamp of this network.
     */
    readonly createdAt: string;
    /**
     * ID of default Security Group of this network.
     */
    readonly defaultSecurityGroupId: string;
    /**
     * Description of the network.
     */
    readonly description: string;
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Labels assigned to this network.
     */
    readonly labels: {[key: string]: string};
    readonly name: string;
    readonly networkId: string;
    readonly subnetIds: string[];
}

export function getVpcNetworkOutput(args?: GetVpcNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcNetworkResult> {
    return pulumi.output(args).apply(a => getVpcNetwork(a, opts))
}

/**
 * A collection of arguments for invoking getVpcNetwork.
 */
export interface GetVpcNetworkOutputArgs {
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Name of the network.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network.
     */
    networkId?: pulumi.Input<string>;
}
