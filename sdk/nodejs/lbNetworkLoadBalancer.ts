// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a network load balancer in the specified folder using the data specified in the config.
 * For more information, see [the official documentation](https://cloud.yandex.com/docs/load-balancer/concepts).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = new yandex.LbNetworkLoadBalancer("foo", {
 *     attachedTargetGroups: [{
 *         healthchecks: [{
 *             httpOptions: {
 *                 path: "/ping",
 *                 port: 8080,
 *             },
 *             name: "http",
 *         }],
 *         targetGroupId: yandex_lb_target_group_my_target_group.id,
 *     }],
 *     listeners: [{
 *         externalAddressSpec: {
 *             ipVersion: "ipv4",
 *         },
 *         name: "my-listener",
 *         port: 8080,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * A network load balancer can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer default network_load_balancer_id
 * ```
 */
export class LbNetworkLoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LbNetworkLoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbNetworkLoadBalancerState, opts?: pulumi.CustomResourceOptions): LbNetworkLoadBalancer {
        return new LbNetworkLoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer';

    /**
     * Returns true if the given object is an instance of LbNetworkLoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbNetworkLoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbNetworkLoadBalancer.__pulumiType;
    }

    /**
     * An AttachedTargetGroup resource. The structure is documented below.
     */
    public readonly attachedTargetGroups!: pulumi.Output<outputs.LbNetworkLoadBalancerAttachedTargetGroup[] | undefined>;
    /**
     * The network load balancer creation timestamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * An optional description of the network load balancer. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Labels to assign to this network load balancer. A list of key/value pairs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Listener specification that will be used by a network load balancer. The structure is documented below.
     */
    public readonly listeners!: pulumi.Output<outputs.LbNetworkLoadBalancerListener[] | undefined>;
    /**
     * Name of the listener. The name must be unique for each listener on a single load balancer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the availability zone where the network load balancer resides.
     * The default is 'ru-central1'.
     */
    public readonly regionId!: pulumi.Output<string | undefined>;
    /**
     * Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a LbNetworkLoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LbNetworkLoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbNetworkLoadBalancerArgs | LbNetworkLoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbNetworkLoadBalancerState | undefined;
            resourceInputs["attachedTargetGroups"] = state ? state.attachedTargetGroups : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["listeners"] = state ? state.listeners : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["regionId"] = state ? state.regionId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as LbNetworkLoadBalancerArgs | undefined;
            resourceInputs["attachedTargetGroups"] = args ? args.attachedTargetGroups : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["listeners"] = args ? args.listeners : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regionId"] = args ? args.regionId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbNetworkLoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbNetworkLoadBalancer resources.
 */
export interface LbNetworkLoadBalancerState {
    /**
     * An AttachedTargetGroup resource. The structure is documented below.
     */
    attachedTargetGroups?: pulumi.Input<pulumi.Input<inputs.LbNetworkLoadBalancerAttachedTargetGroup>[]>;
    /**
     * The network load balancer creation timestamp.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * An optional description of the network load balancer. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this network load balancer. A list of key/value pairs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Listener specification that will be used by a network load balancer. The structure is documented below.
     */
    listeners?: pulumi.Input<pulumi.Input<inputs.LbNetworkLoadBalancerListener>[]>;
    /**
     * Name of the listener. The name must be unique for each listener on a single load balancer.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the availability zone where the network load balancer resides.
     * The default is 'ru-central1'.
     */
    regionId?: pulumi.Input<string>;
    /**
     * Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbNetworkLoadBalancer resource.
 */
export interface LbNetworkLoadBalancerArgs {
    /**
     * An AttachedTargetGroup resource. The structure is documented below.
     */
    attachedTargetGroups?: pulumi.Input<pulumi.Input<inputs.LbNetworkLoadBalancerAttachedTargetGroup>[]>;
    /**
     * An optional description of the network load balancer. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this network load balancer. A list of key/value pairs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Listener specification that will be used by a network load balancer. The structure is documented below.
     */
    listeners?: pulumi.Input<pulumi.Input<inputs.LbNetworkLoadBalancerListener>[]>;
    /**
     * Name of the listener. The name must be unique for each listener on a single load balancer.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the availability zone where the network load balancer resides.
     * The default is 'ru-central1'.
     */
    regionId?: pulumi.Input<string>;
    /**
     * Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
     */
    type?: pulumi.Input<string>;
}
