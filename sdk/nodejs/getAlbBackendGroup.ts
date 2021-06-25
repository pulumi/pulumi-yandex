// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Application Load Balancer Backend Group. For more information, see
 * [Yandex.Cloud Application Load Balancer](https://cloud.yandex.com/en/docs/application-load-balancer/quickstart).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = pulumi.output(yandex.getAlbBackendGroup({
 *     backendGroupId: "my-backend-group-id",
 * }, { async: true }));
 * ```
 *
 * This data source is used to define [Application Load Balancer Backend Groups] that can be used by other resources.
 */
export function getAlbBackendGroup(args?: GetAlbBackendGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetAlbBackendGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getAlbBackendGroup:getAlbBackendGroup", {
        "backendGroupId": args.backendGroupId,
        "description": args.description,
        "folderId": args.folderId,
        "grpcBackends": args.grpcBackends,
        "httpBackends": args.httpBackends,
        "labels": args.labels,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlbBackendGroup.
 */
export interface GetAlbBackendGroupArgs {
    /**
     * Backend Group ID.
     */
    readonly backendGroupId?: string;
    /**
     * Description of the backend group.
     */
    readonly description?: string;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    readonly folderId?: string;
    /**
     * Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    readonly grpcBackends?: inputs.GetAlbBackendGroupGrpcBackend[];
    /**
     * Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    readonly httpBackends?: inputs.GetAlbBackendGroupHttpBackend[];
    /**
     * Labels to assign to this backend group.
     */
    readonly labels?: {[key: string]: string};
    /**
     * - Name of the Backend Group.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getAlbBackendGroup.
 */
export interface GetAlbBackendGroupResult {
    readonly backendGroupId: string;
    /**
     * Creation timestamp of this backend group.
     */
    readonly createdAt: string;
    /**
     * Description of the backend group.
     */
    readonly description: string;
    readonly folderId: string;
    /**
     * Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    readonly grpcBackends: outputs.GetAlbBackendGroupGrpcBackend[];
    /**
     * Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    readonly httpBackends: outputs.GetAlbBackendGroupHttpBackend[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Labels to assign to this backend group.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the backend.
     */
    readonly name: string;
}
