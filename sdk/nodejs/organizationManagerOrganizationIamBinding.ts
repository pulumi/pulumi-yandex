// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows creation and management of a single binding within IAM policy for
 * an existing Yandex.Cloud Organization Manager organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const editor = new yandex.OrganizationManagerOrganizationIamBinding("editor", {
 *     members: ["userAccount:some_user_id"],
 *     organizationId: "some_organization_id",
 *     role: "editor",
 * });
 * ```
 *
 * ## Import
 *
 * IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the `organization_id` and role, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/organizationManagerOrganizationIamBinding:OrganizationManagerOrganizationIamBinding viewer "organization_id viewer"
 * ```
 */
export class OrganizationManagerOrganizationIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationManagerOrganizationIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationManagerOrganizationIamBindingState, opts?: pulumi.CustomResourceOptions): OrganizationManagerOrganizationIamBinding {
        return new OrganizationManagerOrganizationIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/organizationManagerOrganizationIamBinding:OrganizationManagerOrganizationIamBinding';

    /**
     * Returns true if the given object is an instance of OrganizationManagerOrganizationIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationManagerOrganizationIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationManagerOrganizationIamBinding.__pulumiType;
    }

    /**
     * An array of identities that will be granted the privilege in the `role`.
     * Each entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **federatedUser:{federated_user_id}**: A unique federated user ID.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * ID of the organization to attach the policy to.
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The role that should be assigned. Only one
     * `yandex.OrganizationManagerOrganizationIamBinding` can be used per role.
     */
    public readonly role!: pulumi.Output<string>;
    public readonly sleepAfter!: pulumi.Output<number | undefined>;

    /**
     * Create a OrganizationManagerOrganizationIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationManagerOrganizationIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationManagerOrganizationIamBindingArgs | OrganizationManagerOrganizationIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationManagerOrganizationIamBindingState | undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["organizationId"] = state ? state.organizationId : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["sleepAfter"] = state ? state.sleepAfter : undefined;
        } else {
            const args = argsOrState as OrganizationManagerOrganizationIamBindingArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["members"] = args ? args.members : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["sleepAfter"] = args ? args.sleepAfter : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationManagerOrganizationIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationManagerOrganizationIamBinding resources.
 */
export interface OrganizationManagerOrganizationIamBindingState {
    /**
     * An array of identities that will be granted the privilege in the `role`.
     * Each entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **federatedUser:{federated_user_id}**: A unique federated user ID.
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the organization to attach the policy to.
     */
    readonly organizationId?: pulumi.Input<string>;
    /**
     * The role that should be assigned. Only one
     * `yandex.OrganizationManagerOrganizationIamBinding` can be used per role.
     */
    readonly role?: pulumi.Input<string>;
    readonly sleepAfter?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OrganizationManagerOrganizationIamBinding resource.
 */
export interface OrganizationManagerOrganizationIamBindingArgs {
    /**
     * An array of identities that will be granted the privilege in the `role`.
     * Each entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **federatedUser:{federated_user_id}**: A unique federated user ID.
     */
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the organization to attach the policy to.
     */
    readonly organizationId: pulumi.Input<string>;
    /**
     * The role that should be assigned. Only one
     * `yandex.OrganizationManagerOrganizationIamBinding` can be used per role.
     */
    readonly role: pulumi.Input<string>;
    readonly sleepAfter?: pulumi.Input<number>;
}
