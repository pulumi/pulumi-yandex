// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a Greenplum cluster within the Yandex.Cloud. For more information, see
 * [the official documentation](https://cloud.yandex.ru/docs/managed-greenplum/).
 *
 * Please read [Pricing for Managed Service for Greenplum](https://cloud.yandex.ru/docs/managed-greenplum/) before using Greenplum cluster.
 *
 * Yandex Managed Service for Greenplum® is now in preview
 *
 * ## Example Usage
 *
 * Example of creating a Single Node Greenplum.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const fooVpcNetwork = new yandex.VpcNetwork("fooVpcNetwork", {});
 * const fooVpcSubnet = new yandex.VpcSubnet("fooVpcSubnet", {
 *     zone: "ru-central1-a",
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.5.0.0/24"],
 * });
 * const test_sg_x = new yandex.VpcSecurityGroup("test-sg-x", {
 *     networkId: fooVpcNetwork.id,
 *     ingresses: [{
 *         protocol: "ANY",
 *         description: "Allow incoming traffic from members of the same security group",
 *         fromPort: 0,
 *         toPort: 65535,
 *         v4CidrBlocks: ["0.0.0.0/0"],
 *     }],
 *     egresses: [{
 *         protocol: "ANY",
 *         description: "Allow outgoing traffic to members of the same security group",
 *         fromPort: 0,
 *         toPort: 65535,
 *         v4CidrBlocks: ["0.0.0.0/0"],
 *     }],
 * });
 * const fooMdbGreenplumCluster = new yandex.MdbGreenplumCluster("fooMdbGreenplumCluster", {
 *     description: "test greenplum cluster",
 *     environment: "PRESTABLE",
 *     networkId: fooVpcNetwork.id,
 *     zoneId: "ru-central1-a",
 *     subnetId: fooVpcSubnet.id,
 *     assignPublicIp: true,
 *     version: "6.17",
 *     masterHostCount: 2,
 *     segmentHostCount: 5,
 *     segmentInHost: 1,
 *     masterSubcluster: {
 *         resources: {
 *             resourcePresetId: "s2.micro",
 *             diskSize: 24,
 *             diskTypeId: "network-ssd",
 *         },
 *     },
 *     segmentSubcluster: {
 *         resources: {
 *             resourcePresetId: "s2.micro",
 *             diskSize: 24,
 *             diskTypeId: "network-ssd",
 *         },
 *     },
 *     userName: "admin_user",
 *     userPassword: "your_super_secret_password",
 *     securityGroupIds: [test_sg_x.id],
 * });
 * ```
 *
 * ## Import
 *
 * A cluster can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/mdbGreenplumCluster:MdbGreenplumCluster foo cluster_id
 * ```
 */
export class MdbGreenplumCluster extends pulumi.CustomResource {
    /**
     * Get an existing MdbGreenplumCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MdbGreenplumClusterState, opts?: pulumi.CustomResourceOptions): MdbGreenplumCluster {
        return new MdbGreenplumCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/mdbGreenplumCluster:MdbGreenplumCluster';

    /**
     * Returns true if the given object is an instance of MdbGreenplumCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MdbGreenplumCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MdbGreenplumCluster.__pulumiType;
    }

    /**
     * Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment.
     */
    public readonly assignPublicIp!: pulumi.Output<boolean>;
    /**
     * Creation timestamp of the cluster.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * Description of the Greenplum cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Aggregated health of the cluster.
     */
    public /*out*/ readonly health!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the Greenplum cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Number of hosts in master subcluster (1 or 2).
     */
    public readonly masterHostCount!: pulumi.Output<number>;
    /**
     * (Computed) Info about hosts in master subcluster. The structure is documented below.
     */
    public /*out*/ readonly masterHosts!: pulumi.Output<outputs.MdbGreenplumClusterMasterHost[]>;
    /**
     * Settings for master subcluster. The structure is documented below.
     */
    public readonly masterSubcluster!: pulumi.Output<outputs.MdbGreenplumClusterMasterSubcluster>;
    /**
     * Name of the Greenplum cluster. Provided by the client when the cluster is created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network, to which the Greenplum cluster uses.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Number of hosts in segment subcluster (from 1 to 32).
     */
    public readonly segmentHostCount!: pulumi.Output<number>;
    /**
     * (Computed) Info about hosts in segment subcluster. The structure is documented below.
     */
    public /*out*/ readonly segmentHosts!: pulumi.Output<outputs.MdbGreenplumClusterSegmentHost[]>;
    /**
     * Number of segments on segment host (not more then 1 + RAM/8).
     */
    public readonly segmentInHost!: pulumi.Output<number>;
    /**
     * Settings for segment subcluster. The structure is documented below.
     */
    public readonly segmentSubcluster!: pulumi.Output<outputs.MdbGreenplumClusterSegmentSubcluster>;
    /**
     * Status of the cluster.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Greenplum cluster admin user name.
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * Greenplum cluster admin password name.
     */
    public readonly userPassword!: pulumi.Output<string>;
    /**
     * Version of the Greenplum cluster. (6.17)
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The availability zone where the Greenplum hosts will be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a MdbGreenplumCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MdbGreenplumClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MdbGreenplumClusterArgs | MdbGreenplumClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MdbGreenplumClusterState | undefined;
            inputs["assignPublicIp"] = state ? state.assignPublicIp : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["health"] = state ? state.health : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["masterHostCount"] = state ? state.masterHostCount : undefined;
            inputs["masterHosts"] = state ? state.masterHosts : undefined;
            inputs["masterSubcluster"] = state ? state.masterSubcluster : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkId"] = state ? state.networkId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["segmentHostCount"] = state ? state.segmentHostCount : undefined;
            inputs["segmentHosts"] = state ? state.segmentHosts : undefined;
            inputs["segmentInHost"] = state ? state.segmentInHost : undefined;
            inputs["segmentSubcluster"] = state ? state.segmentSubcluster : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["userName"] = state ? state.userName : undefined;
            inputs["userPassword"] = state ? state.userPassword : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as MdbGreenplumClusterArgs | undefined;
            if ((!args || args.assignPublicIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assignPublicIp'");
            }
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.masterHostCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'masterHostCount'");
            }
            if ((!args || args.masterSubcluster === undefined) && !opts.urn) {
                throw new Error("Missing required property 'masterSubcluster'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.segmentHostCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'segmentHostCount'");
            }
            if ((!args || args.segmentInHost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'segmentInHost'");
            }
            if ((!args || args.segmentSubcluster === undefined) && !opts.urn) {
                throw new Error("Missing required property 'segmentSubcluster'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            if ((!args || args.userPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPassword'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["assignPublicIp"] = args ? args.assignPublicIp : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["masterHostCount"] = args ? args.masterHostCount : undefined;
            inputs["masterSubcluster"] = args ? args.masterSubcluster : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkId"] = args ? args.networkId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["segmentHostCount"] = args ? args.segmentHostCount : undefined;
            inputs["segmentInHost"] = args ? args.segmentInHost : undefined;
            inputs["segmentSubcluster"] = args ? args.segmentSubcluster : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["userName"] = args ? args.userName : undefined;
            inputs["userPassword"] = args ? args.userPassword : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["health"] = undefined /*out*/;
            inputs["masterHosts"] = undefined /*out*/;
            inputs["segmentHosts"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MdbGreenplumCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MdbGreenplumCluster resources.
 */
export interface MdbGreenplumClusterState {
    /**
     * Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment.
     */
    readonly assignPublicIp?: pulumi.Input<boolean>;
    /**
     * Creation timestamp of the cluster.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * Description of the Greenplum cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * Aggregated health of the cluster.
     */
    readonly health?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Greenplum cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Number of hosts in master subcluster (1 or 2).
     */
    readonly masterHostCount?: pulumi.Input<number>;
    /**
     * (Computed) Info about hosts in master subcluster. The structure is documented below.
     */
    readonly masterHosts?: pulumi.Input<pulumi.Input<inputs.MdbGreenplumClusterMasterHost>[]>;
    /**
     * Settings for master subcluster. The structure is documented below.
     */
    readonly masterSubcluster?: pulumi.Input<inputs.MdbGreenplumClusterMasterSubcluster>;
    /**
     * Name of the Greenplum cluster. Provided by the client when the cluster is created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the network, to which the Greenplum cluster uses.
     */
    readonly networkId?: pulumi.Input<string>;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of hosts in segment subcluster (from 1 to 32).
     */
    readonly segmentHostCount?: pulumi.Input<number>;
    /**
     * (Computed) Info about hosts in segment subcluster. The structure is documented below.
     */
    readonly segmentHosts?: pulumi.Input<pulumi.Input<inputs.MdbGreenplumClusterSegmentHost>[]>;
    /**
     * Number of segments on segment host (not more then 1 + RAM/8).
     */
    readonly segmentInHost?: pulumi.Input<number>;
    /**
     * Settings for segment subcluster. The structure is documented below.
     */
    readonly segmentSubcluster?: pulumi.Input<inputs.MdbGreenplumClusterSegmentSubcluster>;
    /**
     * Status of the cluster.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * Greenplum cluster admin user name.
     */
    readonly userName?: pulumi.Input<string>;
    /**
     * Greenplum cluster admin password name.
     */
    readonly userPassword?: pulumi.Input<string>;
    /**
     * Version of the Greenplum cluster. (6.17)
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The availability zone where the Greenplum hosts will be created.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MdbGreenplumCluster resource.
 */
export interface MdbGreenplumClusterArgs {
    /**
     * Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment.
     */
    readonly assignPublicIp: pulumi.Input<boolean>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * Description of the Greenplum cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
     */
    readonly environment: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Greenplum cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Number of hosts in master subcluster (1 or 2).
     */
    readonly masterHostCount: pulumi.Input<number>;
    /**
     * Settings for master subcluster. The structure is documented below.
     */
    readonly masterSubcluster: pulumi.Input<inputs.MdbGreenplumClusterMasterSubcluster>;
    /**
     * Name of the Greenplum cluster. Provided by the client when the cluster is created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the network, to which the Greenplum cluster uses.
     */
    readonly networkId: pulumi.Input<string>;
    /**
     * A set of ids of security groups assigned to hosts of the cluster.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of hosts in segment subcluster (from 1 to 32).
     */
    readonly segmentHostCount: pulumi.Input<number>;
    /**
     * Number of segments on segment host (not more then 1 + RAM/8).
     */
    readonly segmentInHost: pulumi.Input<number>;
    /**
     * Settings for segment subcluster. The structure is documented below.
     */
    readonly segmentSubcluster: pulumi.Input<inputs.MdbGreenplumClusterSegmentSubcluster>;
    /**
     * The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * Greenplum cluster admin user name.
     */
    readonly userName: pulumi.Input<string>;
    /**
     * Greenplum cluster admin password name.
     */
    readonly userPassword: pulumi.Input<string>;
    /**
     * Version of the Greenplum cluster. (6.17)
     */
    readonly version: pulumi.Input<string>;
    /**
     * The availability zone where the Greenplum hosts will be created.
     */
    readonly zone: pulumi.Input<string>;
}
