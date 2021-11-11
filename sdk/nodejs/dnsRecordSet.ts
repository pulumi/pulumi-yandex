// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a DNS Recordset.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = new yandex.VpcNetwork("foo", {});
 * const zone1 = new yandex.DnsZone("zone1", {
 *     description: "desc",
 *     labels: {
 *         label1: "label-1-value",
 *     },
 *     zone: "example.com.",
 *     "public": false,
 *     privateNetworks: [foo.id],
 * });
 * const rs1 = new yandex.DnsRecordSet("rs1", {
 *     zoneId: zone1.id,
 *     type: "A",
 *     ttl: 200,
 *     datas: ["10.1.0.1"],
 * });
 * const rs2 = new yandex.DnsRecordSet("rs2", {
 *     zoneId: zone1.id,
 *     type: "A",
 *     ttl: 200,
 *     datas: ["10.1.0.2"],
 * });
 * ```
 *
 * ## Import
 *
 * DNS recordset can be imported using this format
 *
 * ```sh
 *  $ pulumi import yandex:index/dnsRecordSet:DnsRecordSet rs1 {{zone_id}}/{{name}}/{{type}}
 * ```
 */
export class DnsRecordSet extends pulumi.CustomResource {
    /**
     * Get an existing DnsRecordSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsRecordSetState, opts?: pulumi.CustomResourceOptions): DnsRecordSet {
        return new DnsRecordSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/dnsRecordSet:DnsRecordSet';

    /**
     * Returns true if the given object is an instance of DnsRecordSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsRecordSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsRecordSet.__pulumiType;
    }

    /**
     * The string data for the records in this record set.
     */
    public readonly datas!: pulumi.Output<string[]>;
    /**
     * The DNS name this record set will apply to.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The time-to-live of this record set (seconds).
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * The DNS record set type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The id of the zone in which this record set will reside.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a DnsRecordSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsRecordSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsRecordSetArgs | DnsRecordSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsRecordSetState | undefined;
            inputs["datas"] = state ? state.datas : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as DnsRecordSetArgs | undefined;
            if ((!args || args.datas === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datas'");
            }
            if ((!args || args.ttl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttl'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            inputs["datas"] = args ? args.datas : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DnsRecordSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsRecordSet resources.
 */
export interface DnsRecordSetState {
    /**
     * The string data for the records in this record set.
     */
    datas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The DNS name this record set will apply to.
     */
    name?: pulumi.Input<string>;
    /**
     * The time-to-live of this record set (seconds).
     */
    ttl?: pulumi.Input<number>;
    /**
     * The DNS record set type.
     */
    type?: pulumi.Input<string>;
    /**
     * The id of the zone in which this record set will reside.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsRecordSet resource.
 */
export interface DnsRecordSetArgs {
    /**
     * The string data for the records in this record set.
     */
    datas: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The DNS name this record set will apply to.
     */
    name?: pulumi.Input<string>;
    /**
     * The time-to-live of this record set (seconds).
     */
    ttl: pulumi.Input<number>;
    /**
     * The DNS record set type.
     */
    type: pulumi.Input<string>;
    /**
     * The id of the zone in which this record set will reside.
     */
    zoneId: pulumi.Input<string>;
}
