// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a subnet within the Yandex.Cloud. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/vpc/concepts/network#subnet).
 *
 * * How-to Guides
 *     * [Cloud Networking](https://cloud.yandex.com/docs/vpc/)
 *     * [VPC Addressing](https://cloud.yandex.com/docs/vpc/concepts/address)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const lab_net = new yandex.VpcNetwork("lab-net", {});
 * const lab_subnet_a = new yandex.VpcSubnet("lab-subnet-a", {
 *     networkId: lab_net.id,
 *     v4CidrBlocks: ["10.2.0.0/16"],
 *     zone: "ru-central1-a",
 * });
 * ```
 *
 * ## Import
 *
 * A subnet can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/vpcSubnet:VpcSubnet default subnet_id
 * ```
 */
export class VpcSubnet extends pulumi.CustomResource {
    /**
     * Get an existing VpcSubnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcSubnetState, opts?: pulumi.CustomResourceOptions): VpcSubnet {
        return new VpcSubnet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/vpcSubnet:VpcSubnet';

    /**
     * Returns true if the given object is an instance of VpcSubnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcSubnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcSubnet.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * An optional description of the subnet. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Options for DHCP client. The structure is documented below.
     */
    public readonly dhcpOptions!: pulumi.Output<outputs.VpcSubnetDhcpOptions | undefined>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Labels to assign to this subnet. A list of key/value pairs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of the subnet. Provided by the client when the subnet is created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network this subnet belongs to.
     * Only networks that are in the distributed mode can have subnets.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * The ID of the route table to assign to this subnet. Assigned route table should
     * belong to the same network as this subnet.
     */
    public readonly routeTableId!: pulumi.Output<string | undefined>;
    /**
     * A list of blocks of internal IPv4 addresses that are owned by this subnet.
     * Provide this property when you create the subnet. For example, 10.0.0.0/22 or 192.168.0.0/16.
     * Blocks of addresses must be unique and non-overlapping within a network.
     * Minimum subnet size is /28, and maximum subnet size is /16. Only IPv4 is supported.
     */
    public readonly v4CidrBlocks!: pulumi.Output<string[]>;
    /**
     * An optional list of blocks of IPv6 addresses that are owned by this subnet.
     */
    public /*out*/ readonly v6CidrBlocks!: pulumi.Output<string[]>;
    /**
     * Name of the Yandex.Cloud zone for this subnet.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a VpcSubnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcSubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcSubnetArgs | VpcSubnetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcSubnetState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dhcpOptions"] = state ? state.dhcpOptions : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
            resourceInputs["v4CidrBlocks"] = state ? state.v4CidrBlocks : undefined;
            resourceInputs["v6CidrBlocks"] = state ? state.v6CidrBlocks : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as VpcSubnetArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.v4CidrBlocks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'v4CidrBlocks'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dhcpOptions"] = args ? args.dhcpOptions : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["routeTableId"] = args ? args.routeTableId : undefined;
            resourceInputs["v4CidrBlocks"] = args ? args.v4CidrBlocks : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["v6CidrBlocks"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VpcSubnet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcSubnet resources.
 */
export interface VpcSubnetState {
    createdAt?: pulumi.Input<string>;
    /**
     * An optional description of the subnet. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Options for DHCP client. The structure is documented below.
     */
    dhcpOptions?: pulumi.Input<inputs.VpcSubnetDhcpOptions>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this subnet. A list of key/value pairs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the subnet. Provided by the client when the subnet is created.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network this subnet belongs to.
     * Only networks that are in the distributed mode can have subnets.
     */
    networkId?: pulumi.Input<string>;
    /**
     * The ID of the route table to assign to this subnet. Assigned route table should
     * belong to the same network as this subnet.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * A list of blocks of internal IPv4 addresses that are owned by this subnet.
     * Provide this property when you create the subnet. For example, 10.0.0.0/22 or 192.168.0.0/16.
     * Blocks of addresses must be unique and non-overlapping within a network.
     * Minimum subnet size is /28, and maximum subnet size is /16. Only IPv4 is supported.
     */
    v4CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional list of blocks of IPv6 addresses that are owned by this subnet.
     */
    v6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the Yandex.Cloud zone for this subnet.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcSubnet resource.
 */
export interface VpcSubnetArgs {
    /**
     * An optional description of the subnet. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Options for DHCP client. The structure is documented below.
     */
    dhcpOptions?: pulumi.Input<inputs.VpcSubnetDhcpOptions>;
    /**
     * The ID of the folder to which the resource belongs.
     * If omitted, the provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this subnet. A list of key/value pairs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the subnet. Provided by the client when the subnet is created.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network this subnet belongs to.
     * Only networks that are in the distributed mode can have subnets.
     */
    networkId: pulumi.Input<string>;
    /**
     * The ID of the route table to assign to this subnet. Assigned route table should
     * belong to the same network as this subnet.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * A list of blocks of internal IPv4 addresses that are owned by this subnet.
     * Provide this property when you create the subnet. For example, 10.0.0.0/22 or 192.168.0.0/16.
     * Blocks of addresses must be unique and non-overlapping within a network.
     * Minimum subnet size is /28, and maximum subnet size is /16. Only IPv4 is supported.
     */
    v4CidrBlocks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the Yandex.Cloud zone for this subnet.
     */
    zone?: pulumi.Input<string>;
}
