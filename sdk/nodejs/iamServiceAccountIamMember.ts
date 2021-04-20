// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * When managing IAM roles, you can treat a service account either as a resource or as an identity.
 * This resource is used to add IAM policy bindings to a service account resource to configure permissions
 * that define who can edit the service account.
 *
 * There are three different resources that help you manage your IAM policy for a service account.
 * Each of these resources is used for a different use case:
 *
 * * yandex_iam_service_account_iam_policy: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
 * * yandex_iam_service_account_iam_binding: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
 * * yandex_iam_service_account_iam_member: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role of the service account are preserved.
 *
 * > **Note:** `yandex.IamServiceAccountIamPolicy` **cannot** be used in conjunction with `yandex.IamServiceAccountIamBinding` and `yandex.IamServiceAccountIamMember` or they will conflict over what your policy should be.
 *
 * > **Note:** `yandex.IamServiceAccountIamBinding` resources **can be** used in conjunction with `yandex.IamServiceAccountIamMember` resources **only if** they do not grant privileges to the same role.
 *
 * ## yandex\_service\_account\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const admin_account_iam = new yandex.IamServiceAccountIamMember("admin-account-iam", {
 *     member: "userAccount:bar_user_id",
 *     role: "admin",
 *     serviceAccountId: "your-service-account-id",
 * });
 * ```
 *
 * ## Import
 *
 * Service account IAM member resources can be imported using the service account ID, role and member.
 *
 * ```sh
 *  $ pulumi import yandex:index/iamServiceAccountIamMember:IamServiceAccountIamMember admin-account-iam "service_account_id roles/editor foo@example.com"
 * ```
 */
export class IamServiceAccountIamMember extends pulumi.CustomResource {
    /**
     * Get an existing IamServiceAccountIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamServiceAccountIamMemberState, opts?: pulumi.CustomResourceOptions): IamServiceAccountIamMember {
        return new IamServiceAccountIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/iamServiceAccountIamMember:IamServiceAccountIamMember';

    /**
     * Returns true if the given object is an instance of IamServiceAccountIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamServiceAccountIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamServiceAccountIamMember.__pulumiType;
    }

    /**
     * Identity that will be granted the privilege in `role`.
     * Entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     */
    public readonly member!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `yandex.IamServiceAccountIamBinding` can be used per role.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The service account ID to apply a policy to.
     */
    public readonly serviceAccountId!: pulumi.Output<string>;
    public readonly sleepAfter!: pulumi.Output<number | undefined>;

    /**
     * Create a IamServiceAccountIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamServiceAccountIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamServiceAccountIamMemberArgs | IamServiceAccountIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamServiceAccountIamMemberState | undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            inputs["sleepAfter"] = state ? state.sleepAfter : undefined;
        } else {
            const args = argsOrState as IamServiceAccountIamMemberArgs | undefined;
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.serviceAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            inputs["member"] = args ? args.member : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            inputs["sleepAfter"] = args ? args.sleepAfter : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IamServiceAccountIamMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamServiceAccountIamMember resources.
 */
export interface IamServiceAccountIamMemberState {
    /**
     * Identity that will be granted the privilege in `role`.
     * Entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     */
    readonly member?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `yandex.IamServiceAccountIamBinding` can be used per role.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The service account ID to apply a policy to.
     */
    readonly serviceAccountId?: pulumi.Input<string>;
    readonly sleepAfter?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a IamServiceAccountIamMember resource.
 */
export interface IamServiceAccountIamMemberArgs {
    /**
     * Identity that will be granted the privilege in `role`.
     * Entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     */
    readonly member: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `yandex.IamServiceAccountIamBinding` can be used per role.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The service account ID to apply a policy to.
     */
    readonly serviceAccountId: pulumi.Input<string>;
    readonly sleepAfter?: pulumi.Input<number>;
}
