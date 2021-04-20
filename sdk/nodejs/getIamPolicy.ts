// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Generates an [IAM] policy document that may be referenced by and applied to
 * other Yandex.Cloud Platform resources, such as the `yandex.getResourcemanagerFolder` resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const admin = pulumi.output(yandex.getIamPolicy({
 *     bindings: [
 *         {
 *             members: ["userAccount:user_id_1"],
 *             role: "admin",
 *         },
 *         {
 *             members: ["userAccount:user_id_2"],
 *             role: "viewer",
 *         },
 *     ],
 * }, { async: true }));
 * ```
 *
 * This data source is used to define [IAM] policies to apply to other resources.
 * Currently, defining a policy through a data source and referencing that policy
 * from another resource is the only way to apply an IAM policy to a resource.
 */
export function getIamPolicy(args: GetIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetIamPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getIamPolicy:getIamPolicy", {
        "bindings": args.bindings,
    }, opts);
}

/**
 * A collection of arguments for invoking getIamPolicy.
 */
export interface GetIamPolicyArgs {
    /**
     * A nested configuration block (described below)
     * that defines a binding to be included in the policy document. Multiple
     * `binding` arguments are supported.
     */
    readonly bindings: inputs.GetIamPolicyBinding[];
}

/**
 * A collection of values returned by getIamPolicy.
 */
export interface GetIamPolicyResult {
    readonly bindings: outputs.GetIamPolicyBinding[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The above bindings serialized in a format suitable for
     * referencing from a resource that supports IAM.
     */
    readonly policyData: string;
}
