// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a route table within the Yandex.Cloud. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/vpc/concepts).
 *
 * * How-to Guides
 *     * [Cloud Networking](https://cloud.yandex.com/docs/vpc/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const lab_net = new yandex.VpcNetwork("lab-net", {});
 * const lab_rt_a = new yandex.VpcRouteTable("lab-rt-a", {
 *     networkId: lab_net.id,
 *     staticRoutes: [{
 *         destinationPrefix: "10.2.0.0/16",
 *         nextHopAddress: "172.16.10.10",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * A route table can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/vpcRouteTable:VpcRouteTable default route_table_id
 * ```
 */
export class VpcRouteTable extends pulumi.CustomResource {
    /**
     * Get an existing VpcRouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcRouteTableState, opts?: pulumi.CustomResourceOptions): VpcRouteTable {
        return new VpcRouteTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/vpcRouteTable:VpcRouteTable';

    /**
     * Returns true if the given object is an instance of VpcRouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcRouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcRouteTable.__pulumiType;
    }

    /**
     * Creation timestamp of the route table.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * An optional description of the route table. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Labels to assign to this route table. A list of key/value pairs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of the route table. Provided by the client when the route table is created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network this route table belongs to.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * A list of static route records for the route table. The structure is documented below.
     */
    public readonly staticRoutes!: pulumi.Output<outputs.VpcRouteTableStaticRoute[] | undefined>;

    /**
     * Create a VpcRouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcRouteTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcRouteTableArgs | VpcRouteTableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcRouteTableState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkId"] = state ? state.networkId : undefined;
            inputs["staticRoutes"] = state ? state.staticRoutes : undefined;
        } else {
            const args = argsOrState as VpcRouteTableArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkId"] = args ? args.networkId : undefined;
            inputs["staticRoutes"] = args ? args.staticRoutes : undefined;
            inputs["createdAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VpcRouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcRouteTable resources.
 */
export interface VpcRouteTableState {
    /**
     * Creation timestamp of the route table.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * An optional description of the route table. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this route table. A list of key/value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the route table. Provided by the client when the route table is created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the network this route table belongs to.
     */
    readonly networkId?: pulumi.Input<string>;
    /**
     * A list of static route records for the route table. The structure is documented below.
     */
    readonly staticRoutes?: pulumi.Input<pulumi.Input<inputs.VpcRouteTableStaticRoute>[]>;
}

/**
 * The set of arguments for constructing a VpcRouteTable resource.
 */
export interface VpcRouteTableArgs {
    /**
     * An optional description of the route table. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this route table. A list of key/value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the route table. Provided by the client when the route table is created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the network this route table belongs to.
     */
    readonly networkId: pulumi.Input<string>;
    /**
     * A list of static route records for the route table. The structure is documented below.
     */
    readonly staticRoutes?: pulumi.Input<pulumi.Input<inputs.VpcRouteTableStaticRoute>[]>;
}
