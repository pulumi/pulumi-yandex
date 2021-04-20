// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a new container repository. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/container-registry/concepts/repository)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const my_registry = new yandex.ContainerRegistry("my-registry", {});
 * const my_repository = new yandex.ContainerRepository("my-repository", {});
 * ```
 *
 * ## Import
 *
 * A repository can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/containerRepository:ContainerRepository my-repository repository_id
 * ```
 */
export class ContainerRepository extends pulumi.CustomResource {
    /**
     * Get an existing ContainerRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerRepositoryState, opts?: pulumi.CustomResourceOptions): ContainerRepository {
        return new ContainerRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/containerRepository:ContainerRepository';

    /**
     * Returns true if the given object is an instance of ContainerRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerRepository.__pulumiType;
    }

    /**
     * A name of the repository. The name of the repository should start with id of a container registry and match the name of the images that will be pushed in the repository.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ContainerRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerRepositoryArgs | ContainerRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerRepositoryState | undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ContainerRepositoryArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ContainerRepository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerRepository resources.
 */
export interface ContainerRepositoryState {
    /**
     * A name of the repository. The name of the repository should start with id of a container registry and match the name of the images that will be pushed in the repository.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerRepository resource.
 */
export interface ContainerRepositoryArgs {
    /**
     * A name of the repository. The name of the repository should start with id of a container registry and match the name of the images that will be pushed in the repository.
     */
    readonly name?: pulumi.Input<string>;
}
