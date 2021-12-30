// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## yandex\_container\_repository\_iam\_binding
 *
 * Allows creation and management of a single binding within IAM policy for
 * an existing Yandex Container Repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const your_registry = new yandex.ContainerRegistry("your-registry", {folderId: "your-folder-id"});
 * const repo_1 = new yandex.ContainerRepository("repo-1", {});
 * const puller = new yandex.ContainerRepositoryIamBinding("puller", {
 *     repositoryId: repo_1.id,
 *     role: "container-registry.images.puller",
 *     members: ["system:allUsers"],
 * });
 * const repo-2 = yandex.getContainerRepository({
 *     name: "some_repository_name",
 * });
 * const pusher = new yandex.ContainerRepositoryIamBinding("pusher", {
 *     repositoryId: repo_2.then(repo_2 => repo_2.id),
 *     role: "container-registry.images.pusher",
 *     members: ["serviceAccount:your-service-account-id"],
 * });
 * ```
 *
 * ## Import
 *
 * IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the `repository_id` and role, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/containerRepositoryIamBinding:ContainerRepositoryIamBinding puller "repository_id container-registry.images.puller"
 * ```
 */
export class ContainerRepositoryIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing ContainerRepositoryIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerRepositoryIamBindingState, opts?: pulumi.CustomResourceOptions): ContainerRepositoryIamBinding {
        return new ContainerRepositoryIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/containerRepositoryIamBinding:ContainerRepositoryIamBinding';

    /**
     * Returns true if the given object is an instance of ContainerRepositoryIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerRepositoryIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerRepositoryIamBinding.__pulumiType;
    }

    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The [Yandex Container Repository](https://cloud.yandex.com/docs/container-registry/concepts/repository) ID to apply a binding to.
     */
    public readonly repositoryId!: pulumi.Output<string>;
    /**
     * The role that should be applied. See [roles](https://cloud.yandex.com/docs/container-registry/security/).
     */
    public readonly role!: pulumi.Output<string>;
    public readonly sleepAfter!: pulumi.Output<number | undefined>;

    /**
     * Create a ContainerRepositoryIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerRepositoryIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerRepositoryIamBindingArgs | ContainerRepositoryIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerRepositoryIamBindingState | undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["repositoryId"] = state ? state.repositoryId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["sleepAfter"] = state ? state.sleepAfter : undefined;
        } else {
            const args = argsOrState as ContainerRepositoryIamBindingArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.repositoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryId'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["sleepAfter"] = args ? args.sleepAfter : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ContainerRepositoryIamBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerRepositoryIamBinding resources.
 */
export interface ContainerRepositoryIamBindingState {
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The [Yandex Container Repository](https://cloud.yandex.com/docs/container-registry/concepts/repository) ID to apply a binding to.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The role that should be applied. See [roles](https://cloud.yandex.com/docs/container-registry/security/).
     */
    role?: pulumi.Input<string>;
    sleepAfter?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ContainerRepositoryIamBinding resource.
 */
export interface ContainerRepositoryIamBindingArgs {
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
     * * **serviceAccount:{service_account_id}**: A unique service account ID.
     * * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
     */
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The [Yandex Container Repository](https://cloud.yandex.com/docs/container-registry/concepts/repository) ID to apply a binding to.
     */
    repositoryId: pulumi.Input<string>;
    /**
     * The role that should be applied. See [roles](https://cloud.yandex.com/docs/container-registry/security/).
     */
    role: pulumi.Input<string>;
    sleepAfter?: pulumi.Input<number>;
}
