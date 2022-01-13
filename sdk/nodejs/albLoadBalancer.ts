// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates an Application Load Balancer in the specified folder. For more information, see
 * [the official documentation](https://cloud.yandex.com/en/docs/application-load-balancer/concepts/application-load-balancer)
 * .
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const test_balancer = new yandex.AlbLoadBalancer("test-balancer", {
 *     networkId: yandex_vpc_network["test-network"].id,
 *     allocationPolicy: {
 *         locations: [{
 *             zoneId: "ru-central1-a",
 *             subnetId: yandex_vpc_subnet["test-subnet"].id,
 *         }],
 *     },
 *     listeners: [{
 *         name: "my-listener",
 *         endpoints: [{
 *             addresses: [{
 *                 externalIpv4Address: {},
 *             }],
 *             ports: [8080],
 *         }],
 *         http: {
 *             handler: {
 *                 httpRouterId: yandex_alb_http_router["test-router"].id,
 *             },
 *         },
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * An Application Load Balancer can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/albLoadBalancer:AlbLoadBalancer default load_balancer_id
 * ```
 */
export class AlbLoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing AlbLoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlbLoadBalancerState, opts?: pulumi.CustomResourceOptions): AlbLoadBalancer {
        return new AlbLoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/albLoadBalancer:AlbLoadBalancer';

    /**
     * Returns true if the given object is an instance of AlbLoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlbLoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlbLoadBalancer.__pulumiType;
    }

    /**
     * Allocation zones for the Load Balancer instance. The structure is documented below.
     */
    public readonly allocationPolicy!: pulumi.Output<outputs.AlbLoadBalancerAllocationPolicy>;
    /**
     * The Load Balancer creation timestamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * An optional description of the Load Balancer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the folder to which the resource belongs. If omitted, the provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Labels to assign to this Load Balancer. A list of key/value pairs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of listeners for the Load Balancer. The structure is documented below.
     */
    public readonly listeners!: pulumi.Output<outputs.AlbLoadBalancerListener[] | undefined>;
    /**
     * Cloud log group used by the Load Balancer to store access logs.
     */
    public /*out*/ readonly logGroupId!: pulumi.Output<string>;
    /**
     * name of SNI match.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network that the Load Balancer is located at.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * ID of the region that the Load Balancer is located at.
     */
    public readonly regionId!: pulumi.Output<string | undefined>;
    /**
     * A list of ID's of security groups attached to the Load Balancer.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Status of the Load Balancer.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a AlbLoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlbLoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlbLoadBalancerArgs | AlbLoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlbLoadBalancerState | undefined;
            resourceInputs["allocationPolicy"] = state ? state.allocationPolicy : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["listeners"] = state ? state.listeners : undefined;
            resourceInputs["logGroupId"] = state ? state.logGroupId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["regionId"] = state ? state.regionId : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AlbLoadBalancerArgs | undefined;
            if ((!args || args.allocationPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allocationPolicy'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            resourceInputs["allocationPolicy"] = args ? args.allocationPolicy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["listeners"] = args ? args.listeners : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["regionId"] = args ? args.regionId : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["logGroupId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AlbLoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlbLoadBalancer resources.
 */
export interface AlbLoadBalancerState {
    /**
     * Allocation zones for the Load Balancer instance. The structure is documented below.
     */
    allocationPolicy?: pulumi.Input<inputs.AlbLoadBalancerAllocationPolicy>;
    /**
     * The Load Balancer creation timestamp.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * An optional description of the Load Balancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs. If omitted, the provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this Load Balancer. A list of key/value pairs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of listeners for the Load Balancer. The structure is documented below.
     */
    listeners?: pulumi.Input<pulumi.Input<inputs.AlbLoadBalancerListener>[]>;
    /**
     * Cloud log group used by the Load Balancer to store access logs.
     */
    logGroupId?: pulumi.Input<string>;
    /**
     * name of SNI match.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network that the Load Balancer is located at.
     */
    networkId?: pulumi.Input<string>;
    /**
     * ID of the region that the Load Balancer is located at.
     */
    regionId?: pulumi.Input<string>;
    /**
     * A list of ID's of security groups attached to the Load Balancer.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the Load Balancer.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlbLoadBalancer resource.
 */
export interface AlbLoadBalancerArgs {
    /**
     * Allocation zones for the Load Balancer instance. The structure is documented below.
     */
    allocationPolicy: pulumi.Input<inputs.AlbLoadBalancerAllocationPolicy>;
    /**
     * An optional description of the Load Balancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs. If omitted, the provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this Load Balancer. A list of key/value pairs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of listeners for the Load Balancer. The structure is documented below.
     */
    listeners?: pulumi.Input<pulumi.Input<inputs.AlbLoadBalancerListener>[]>;
    /**
     * name of SNI match.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network that the Load Balancer is located at.
     */
    networkId: pulumi.Input<string>;
    /**
     * ID of the region that the Load Balancer is located at.
     */
    regionId?: pulumi.Input<string>;
    /**
     * A list of ID's of security groups attached to the Load Balancer.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
