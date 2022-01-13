// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a target group in the specified folder and adds the specified targets to it.
 * For more information, see [the official documentation](https://cloud.yandex.com/docs/load-balancer/concepts/target-resources).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = new yandex.LbTargetGroup("foo", {
 *     regionId: "ru-central1",
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
 *  $ pulumi import yandex:index/lbTargetGroup:LbTargetGroup default target_group_id
 * ```
 */
export class LbTargetGroup extends pulumi.CustomResource {
    /**
     * Get an existing LbTargetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbTargetGroupState, opts?: pulumi.CustomResourceOptions): LbTargetGroup {
        return new LbTargetGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/lbTargetGroup:LbTargetGroup';

    /**
     * Returns true if the given object is an instance of LbTargetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbTargetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbTargetGroup.__pulumiType;
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
     * ID of the availability zone where the target group resides. 
     * The default is 'ru-central1'.
     */
    public readonly regionId!: pulumi.Output<string | undefined>;
    /**
     * A Target resource. The structure is documented below.
     */
    public readonly targets!: pulumi.Output<outputs.LbTargetGroupTarget[] | undefined>;

    /**
     * Create a LbTargetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LbTargetGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbTargetGroupArgs | LbTargetGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbTargetGroupState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["regionId"] = state ? state.regionId : undefined;
            resourceInputs["targets"] = state ? state.targets : undefined;
        } else {
            const args = argsOrState as LbTargetGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regionId"] = args ? args.regionId : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbTargetGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbTargetGroup resources.
 */
export interface LbTargetGroupState {
    /**
     * The target group creation timestamp.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * An optional description of the target group. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this target group. A list of key/value pairs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the target group. Provided by the client when the target group is created.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the availability zone where the target group resides. 
     * The default is 'ru-central1'.
     */
    regionId?: pulumi.Input<string>;
    /**
     * A Target resource. The structure is documented below.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.LbTargetGroupTarget>[]>;
}

/**
 * The set of arguments for constructing a LbTargetGroup resource.
 */
export interface LbTargetGroupArgs {
    /**
     * An optional description of the target group. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this target group. A list of key/value pairs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the target group. Provided by the client when the target group is created.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the availability zone where the target group resides. 
     * The default is 'ru-central1'.
     */
    regionId?: pulumi.Input<string>;
    /**
     * A Target resource. The structure is documented below.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.LbTargetGroupTarget>[]>;
}
