// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows creation and management of a single member for a single binding within
 * the IAM policy for an existing Yandex Organization Manager organization.
 *
 * > **Note:** Roles controlled by `yandex.OrganizationManagerOrganizationIamBinding`
 *    should not be assigned using `yandex.OrganizationManagerOrganizationIamMember`.
 *
 * > **Note:** When you delete `yandex.OrganizationManagerOrganizationIamBinding` resource,
 *    the roles can be deleted from other users within the organization as well. Be careful!
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const editor = new yandex.OrganizationManagerOrganizationIamMember("editor", {
 *     member: "userAccount:user_id",
 *     organizationId: "some_organization_id",
 *     role: "editor",
 * });
 * ```
 *
 * ## Import
 *
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the `organization id`, role, and account, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/organizationManagerOrganizationIamMember:OrganizationManagerOrganizationIamMember my_project "organization_id viewer foo@example.com"
 * ```
 */
export class OrganizationManagerOrganizationIamMember extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationManagerOrganizationIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationManagerOrganizationIamMemberState, opts?: pulumi.CustomResourceOptions): OrganizationManagerOrganizationIamMember {
        return new OrganizationManagerOrganizationIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/organizationManagerOrganizationIamMember:OrganizationManagerOrganizationIamMember';

    /**
     * Returns true if the given object is an instance of OrganizationManagerOrganizationIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationManagerOrganizationIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationManagerOrganizationIamMember.__pulumiType;
    }

    /**
     * The identity that will be granted the privilege that is specified in the `role` field.
     * This field can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **federatedUser:{federated_user_id}**: A unique federated user ID.
     */
    public readonly member!: pulumi.Output<string>;
    /**
     * ID of the organization to attach a policy to.
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The role that should be assigned.
     */
    public readonly role!: pulumi.Output<string>;
    public readonly sleepAfter!: pulumi.Output<number | undefined>;

    /**
     * Create a OrganizationManagerOrganizationIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationManagerOrganizationIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationManagerOrganizationIamMemberArgs | OrganizationManagerOrganizationIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationManagerOrganizationIamMemberState | undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["organizationId"] = state ? state.organizationId : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["sleepAfter"] = state ? state.sleepAfter : undefined;
        } else {
            const args = argsOrState as OrganizationManagerOrganizationIamMemberArgs | undefined;
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["member"] = args ? args.member : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["sleepAfter"] = args ? args.sleepAfter : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationManagerOrganizationIamMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationManagerOrganizationIamMember resources.
 */
export interface OrganizationManagerOrganizationIamMemberState {
    /**
     * The identity that will be granted the privilege that is specified in the `role` field.
     * This field can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **federatedUser:{federated_user_id}**: A unique federated user ID.
     */
    member?: pulumi.Input<string>;
    /**
     * ID of the organization to attach a policy to.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The role that should be assigned.
     */
    role?: pulumi.Input<string>;
    sleepAfter?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OrganizationManagerOrganizationIamMember resource.
 */
export interface OrganizationManagerOrganizationIamMemberArgs {
    /**
     * The identity that will be granted the privilege that is specified in the `role` field.
     * This field can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **federatedUser:{federated_user_id}**: A unique federated user ID.
     */
    member: pulumi.Input<string>;
    /**
     * ID of the organization to attach a policy to.
     */
    organizationId: pulumi.Input<string>;
    /**
     * The role that should be assigned.
     */
    role: pulumi.Input<string>;
    sleepAfter?: pulumi.Input<number>;
}
