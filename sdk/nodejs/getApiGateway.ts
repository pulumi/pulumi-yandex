// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getApiGateway(args?: GetApiGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetApiGatewayResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getApiGateway:getApiGateway", {
        "apiGatewayId": args.apiGatewayId,
        "folderId": args.folderId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiGateway.
 */
export interface GetApiGatewayArgs {
    readonly apiGatewayId?: string;
    readonly folderId?: string;
    readonly name?: string;
}

/**
 * A collection of values returned by getApiGateway.
 */
export interface GetApiGatewayResult {
    readonly apiGatewayId?: string;
    readonly createdAt: string;
    readonly description: string;
    readonly domain: string;
    readonly folderId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: {[key: string]: string};
    readonly logGroupId: string;
    readonly name?: string;
    readonly status: string;
}
