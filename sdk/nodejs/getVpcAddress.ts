// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex VPC address. For more information, see
 * Yandex.Cloud VPC.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const addr = pulumi.output(yandex.getVpcAddress({
 *     addressId: "my-address-id",
 * }));
 * ```
 *
 * This data source is used to define [VPC Address] that can be used by other resources.
 */
export function getVpcAddress(args?: GetVpcAddressArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcAddressResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getVpcAddress:getVpcAddress", {
        "addressId": args.addressId,
        "folderId": args.folderId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcAddress.
 */
export interface GetVpcAddressArgs {
    /**
     * ID of the address.
     */
    addressId?: string;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: string;
    /**
     * Name of the address.
     */
    name?: string;
}

/**
 * A collection of values returned by getVpcAddress.
 */
export interface GetVpcAddressResult {
    readonly addressId: string;
    /**
     * Creation timestamp of this address.
     */
    readonly createdAt: string;
    /**
     * Description of the address.
     */
    readonly description: string;
    /**
     * spec of IP v4 address.
     */
    readonly externalIpv4Addresses: outputs.GetVpcAddressExternalIpv4Address[];
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Labels assigned to this address.
     */
    readonly labels: {[key: string]: string};
    readonly name: string;
    /**
     * `false` means that address is ephemeral.
     */
    readonly reserved: boolean;
    /**
     * `true` if address is used.
     */
    readonly used: boolean;
}

export function getVpcAddressOutput(args?: GetVpcAddressOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcAddressResult> {
    return pulumi.output(args).apply(a => getVpcAddress(a, opts))
}

/**
 * A collection of arguments for invoking getVpcAddress.
 */
export interface GetVpcAddressOutputArgs {
    /**
     * ID of the address.
     */
    addressId?: pulumi.Input<string>;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Name of the address.
     */
    name?: pulumi.Input<string>;
}
