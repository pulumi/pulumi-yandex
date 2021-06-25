// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a target group in the specified folder and adds the specified targets to it.
 * For more information, see [the official documentation](https://cloud.yandex.com/en/docs/application-load-balancer/concepts/target-group).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = new yandex.AlbTargetGroup("foo", {
 *     targets: [
 *         {
 *             address: yandex_compute_instance_my_instance_1.networkInterface.0.ipAddress,
 *             subnetId: yandex_vpc_subnet_my_subnet.id,
 *         },
 *         {
 *             address: yandex_compute_instance_my_instance_2.networkInterface.0.ipAddress,
 *             subnetId: yandex_vpc_subnet_my_subnet.id,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * A target group can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/albTargetGroup:AlbTargetGroup default target_group_id
 * ```
 */
export class AlbTargetGroup extends pulumi.CustomResource {
    /**
     * Get an existing AlbTargetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlbTargetGroupState, opts?: pulumi.CustomResourceOptions): AlbTargetGroup {
        return new AlbTargetGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/albTargetGroup:AlbTargetGroup';

    /**
     * Returns true if the given object is an instance of AlbTargetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlbTargetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlbTargetGroup.__pulumiType;
    }

    /**
     * The target group creation timestamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * An optional description of the target group. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Labels to assign to this target group. A list of key/value pairs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the target group. Provided by the client when the target group is created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A Target resource. The structure is documented below.
     */
    public readonly targets!: pulumi.Output<outputs.AlbTargetGroupTarget[] | undefined>;

    /**
     * Create a AlbTargetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AlbTargetGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlbTargetGroupArgs | AlbTargetGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlbTargetGroupState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["targets"] = state ? state.targets : undefined;
        } else {
            const args = argsOrState as AlbTargetGroupArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["createdAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlbTargetGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlbTargetGroup resources.
 */
export interface AlbTargetGroupState {
    /**
     * The target group creation timestamp.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * An optional description of the target group. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this target group. A list of key/value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the target group. Provided by the client when the target group is created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A Target resource. The structure is documented below.
     */
    readonly targets?: pulumi.Input<pulumi.Input<inputs.AlbTargetGroupTarget>[]>;
}

/**
 * The set of arguments for constructing a AlbTargetGroup resource.
 */
export interface AlbTargetGroupArgs {
    /**
     * An optional description of the target group. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this target group. A list of key/value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the target group. Provided by the client when the target group is created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A Target resource. The structure is documented below.
     */
    readonly targets?: pulumi.Input<pulumi.Input<inputs.AlbTargetGroupTarget>[]>;
}