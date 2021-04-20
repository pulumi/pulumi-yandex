// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a new snapshot of a disk. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/compute/concepts/snapshot).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const defaultComputeSnapshot = new yandex.ComputeSnapshot("default", {
 *     labels: {
 *         "my-label": "my-label-value",
 *     },
 *     sourceDiskId: "test_disk_id",
 * });
 * ```
 *
 * ## Import
 *
 * A snapshot can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/computeSnapshot:ComputeSnapshot disk-snapshot shapshot_id
 * ```
 */
export class ComputeSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing ComputeSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeSnapshotState, opts?: pulumi.CustomResourceOptions): ComputeSnapshot {
        return new ComputeSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/computeSnapshot:ComputeSnapshot';

    /**
     * Returns true if the given object is an instance of ComputeSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeSnapshot.__pulumiType;
    }

    /**
     * Creation timestamp of the snapshot.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Size of the disk when the snapshot was created, specified in GB.
     */
    public /*out*/ readonly diskSize!: pulumi.Output<number>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the snapshot.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A name for the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the disk to create a snapshot from.
     */
    public readonly sourceDiskId!: pulumi.Output<string>;
    /**
     * Size of the snapshot, specified in GB.
     */
    public /*out*/ readonly storageSize!: pulumi.Output<number>;

    /**
     * Create a ComputeSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeSnapshotArgs | ComputeSnapshotState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeSnapshotState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["diskSize"] = state ? state.diskSize : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["sourceDiskId"] = state ? state.sourceDiskId : undefined;
            inputs["storageSize"] = state ? state.storageSize : undefined;
        } else {
            const args = argsOrState as ComputeSnapshotArgs | undefined;
            if ((!args || args.sourceDiskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceDiskId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sourceDiskId"] = args ? args.sourceDiskId : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["diskSize"] = undefined /*out*/;
            inputs["storageSize"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ComputeSnapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeSnapshot resources.
 */
export interface ComputeSnapshotState {
    /**
     * Creation timestamp of the snapshot.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * Description of the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Size of the disk when the snapshot was created, specified in GB.
     */
    readonly diskSize?: pulumi.Input<number>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the snapshot.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A name for the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the disk to create a snapshot from.
     */
    readonly sourceDiskId?: pulumi.Input<string>;
    /**
     * Size of the snapshot, specified in GB.
     */
    readonly storageSize?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ComputeSnapshot resource.
 */
export interface ComputeSnapshotArgs {
    /**
     * Description of the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the snapshot.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A name for the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the disk to create a snapshot from.
     */
    readonly sourceDiskId: pulumi.Input<string>;
}
