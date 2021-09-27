// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * A Disk Placement Group resource. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/compute/concepts/disk#nr-disks).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const group1 = new yandex.ComputeDiskPlacementGroup("group1", {
 *     description: "my description",
 *     folderId: "abc*********123",
 * });
 * ```
 */
export class ComputeDiskPlacementGroup extends pulumi.CustomResource {
    /**
     * Get an existing ComputeDiskPlacementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeDiskPlacementGroupState, opts?: pulumi.CustomResourceOptions): ComputeDiskPlacementGroup {
        return new ComputeDiskPlacementGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/computeDiskPlacementGroup:ComputeDiskPlacementGroup';

    /**
     * Returns true if the given object is an instance of ComputeDiskPlacementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeDiskPlacementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeDiskPlacementGroup.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A description of the Disk Placement Group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the Disk Placement Group.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the Disk Placement Group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Status of the Disk Placement Group.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * ID of the zone where the Disk Placement Group resides.
     */
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a ComputeDiskPlacementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ComputeDiskPlacementGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeDiskPlacementGroupArgs | ComputeDiskPlacementGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeDiskPlacementGroupState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ComputeDiskPlacementGroupArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ComputeDiskPlacementGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeDiskPlacementGroup resources.
 */
export interface ComputeDiskPlacementGroupState {
    readonly createdAt?: pulumi.Input<string>;
    /**
     * A description of the Disk Placement Group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Disk Placement Group.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Disk Placement Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Status of the Disk Placement Group.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * ID of the zone where the Disk Placement Group resides.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeDiskPlacementGroup resource.
 */
export interface ComputeDiskPlacementGroupArgs {
    /**
     * A description of the Disk Placement Group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Disk Placement Group.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Disk Placement Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the zone where the Disk Placement Group resides.
     */
    readonly zone?: pulumi.Input<string>;
}
