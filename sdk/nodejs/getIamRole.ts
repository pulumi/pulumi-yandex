// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Generates an [IAM] role document that may be referenced by and applied to
 * other Yandex.Cloud Platform resources, such as the `yandex.ResourcemanagerFolder` resource. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/roles).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const admin = pulumi.output(yandex.getIamRole({
 *     binding: [{
 *         members: ["userAccount:user_id_1"],
 *         role: "admin",
 *     }],
 * }));
 * ```
 *
 * This data source is used to define [IAM] roles in order to apply them to other resources.
 * Currently, defining a role through a data source and referencing that role
 * from another resource is the only way to apply an IAM role to a resource.
 */
export function getIamRole(args?: GetIamRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetIamRoleResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("yandex:index/getIamRole:getIamRole", {
        "description": args.description,
        "roleId": args.roleId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIamRole.
 */
export interface GetIamRoleArgs {
    description?: string;
    roleId?: string;
}

/**
 * A collection of values returned by getIamRole.
 */
export interface GetIamRoleResult {
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly roleId: string;
}

export function getIamRoleOutput(args?: GetIamRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIamRoleResult> {
    return pulumi.output(args).apply(a => getIamRole(a, opts))
}

/**
 * A collection of arguments for invoking getIamRole.
 */
export interface GetIamRoleOutputArgs {
    description?: pulumi.Input<string>;
    roleId?: pulumi.Input<string>;
}
