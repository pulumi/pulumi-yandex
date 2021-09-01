// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a Kafka cluster within the Yandex.Cloud. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts).
 *
 * ## Example Usage
 *
 * Example of creating a Single Node Kafka.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const fooVpcNetwork = new yandex.VpcNetwork("foo", {});
 * const fooVpcSubnet = new yandex.VpcSubnet("foo", {
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.5.0.0/24"],
 *     zone: "ru-central1-a",
 * });
 * const fooMdbKafkaCluster = new yandex.MdbKafkaCluster("foo", {
 *     config: {
 *         assignPublicIp: false,
 *         brokersCount: 1,
 *         kafka: {
 *             kafkaConfig: {
 *                 compressionType: "COMPRESSION_TYPE_ZSTD",
 *                 defaultReplicationFactor: 1,
 *                 logFlushIntervalMessages: 1024,
 *                 logFlushIntervalMs: 1000,
 *                 logFlushSchedulerIntervalMs: 1000,
 *                 logPreallocate: true,
 *                 logRetentionBytes: 1073741824,
 *                 logRetentionHours: 168,
 *                 logRetentionMinutes: 10080,
 *                 logRetentionMs: 86400000,
 *                 logSegmentBytes: 134217728,
 *                 numPartitions: 10,
 *             },
 *             resources: {
 *                 diskSize: 32,
 *                 diskTypeId: "network-ssd",
 *                 resourcePresetId: "s2.micro",
 *             },
 *         },
 *         unmanagedTopics: false,
 *         version: "2.6",
 *         zones: ["ru-central1-a"],
 *     },
 *     environment: "PRESTABLE",
 *     networkId: fooVpcNetwork.id,
 *     subnetIds: [fooVpcSubnet.id],
 *     topics: [
 *         {
 *             name: "input",
 *             partitions: 2,
 *             replicationFactor: 1,
 *             topicConfig: {
 *                 compressionType: "COMPRESSION_TYPE_LZ4",
 *                 deleteRetentionMs: 86400000,
 *                 fileDeleteDelayMs: 60000,
 *                 flushMessages: 128,
 *                 flushMs: 1000,
 *                 maxMessageBytes: 1048588,
 *                 minCompactionLagMs: 0,
 *                 minInsyncReplicas: 1,
 *                 preallocate: true,
 *                 retentionBytes: 10737418240,
 *                 retentionMs: 604800000,
 *                 segmentBytes: 268435456,
 *             },
 *         },
 *         {
 *             name: "output",
 *             partitions: 6,
 *             replicationFactor: 1,
 *             topicConfig: {
 *                 compressionType: "COMPRESSION_TYPE_GZIP",
 *                 maxMessageBytes: 1048588,
 *                 preallocate: false,
 *                 segmentBytes: 536870912,
 *             },
 *         },
 *     ],
 *     users: [
 *         {
 *             name: "producer-application",
 *             password: "password",
 *             permissions: [{
 *                 role: "ACCESS_ROLE_PRODUCER",
 *                 topicName: "input",
 *             }],
 *         },
 *         {
 *             name: "worker",
 *             password: "password",
 *             permissions: [
 *                 {
 *                     role: "ACCESS_ROLE_CONSUMER",
 *                     topicName: "input",
 *                 },
 *                 {
 *                     role: "ACCESS_ROLE_PRODUCER",
 *                     topicName: "output",
 *                 },
 *             ],
 *         },
 *     ],
 * });
 * ```
 *
 * Example of creating a HA Kafka Cluster with two brokers per AZ (6 brokers + 3 zk)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const fooVpcNetwork = new yandex.VpcNetwork("foo", {});
 * const fooVpcSubnet = new yandex.VpcSubnet("foo", {
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.1.0.0/24"],
 *     zone: "ru-central1-a",
 * });
 * const bar = new yandex.VpcSubnet("bar", {
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.2.0.0/24"],
 *     zone: "ru-central1-b",
 * });
 * const baz = new yandex.VpcSubnet("baz", {
 *     networkId: fooVpcNetwork.id,
 *     v4CidrBlocks: ["10.3.0.0/24"],
 *     zone: "ru-central1-c",
 * });
 * const fooMdbKafkaCluster = new yandex.MdbKafkaCluster("foo", {
 *     config: {
 *         assignPublicIp: true,
 *         brokersCount: 2,
 *         kafka: {
 *             kafkaConfig: {
 *                 compressionType: "COMPRESSION_TYPE_ZSTD",
 *                 defaultReplicationFactor: 6,
 *                 logFlushIntervalMessages: 1024,
 *                 logFlushIntervalMs: 1000,
 *                 logFlushSchedulerIntervalMs: 1000,
 *                 logPreallocate: true,
 *                 logRetentionBytes: 1073741824,
 *                 logRetentionHours: 168,
 *                 logRetentionMinutes: 10080,
 *                 logRetentionMs: 86400000,
 *                 logSegmentBytes: 134217728,
 *                 numPartitions: 10,
 *             },
 *             resources: {
 *                 diskSize: 128,
 *                 diskTypeId: "network-ssd",
 *                 resourcePresetId: "s2.medium",
 *             },
 *         },
 *         unmanagedTopics: false,
 *         version: "2.6",
 *         zones: [
 *             "ru-central1-a",
 *             "ru-central1-b",
 *             "ru-central1-c",
 *         ],
 *         zookeeper: {
 *             resources: {
 *                 diskSize: 20,
 *                 diskTypeId: "network-ssd",
 *                 resourcePresetId: "s2.micro",
 *             },
 *         },
 *     },
 *     environment: "PRESTABLE",
 *     networkId: fooVpcNetwork.id,
 *     subnetIds: [
 *         fooVpcSubnet.id,
 *         bar.id,
 *         baz.id,
 *     ],
 *     topics: [
 *         {
 *             name: "input",
 *             partitions: 2,
 *             replicationFactor: 1,
 *             topicConfig: {
 *                 compressionType: "COMPRESSION_TYPE_LZ4",
 *                 deleteRetentionMs: 86400000,
 *                 fileDeleteDelayMs: 60000,
 *                 flushMessages: 128,
 *                 flushMs: 1000,
 *                 maxMessageBytes: 1048588,
 *                 minCompactionLagMs: 0,
 *                 minInsyncReplicas: 1,
 *                 preallocate: true,
 *                 retentionBytes: 10737418240,
 *                 retentionMs: 604800000,
 *                 segmentBytes: 268435456,
 *             },
 *         },
 *         {
 *             name: "output",
 *             partitions: 6,
 *             replicationFactor: 1,
 *             topicConfig: {
 *                 compressionType: "COMPRESSION_TYPE_GZIP",
 *                 maxMessageBytes: 1048588,
 *                 preallocate: false,
 *                 segmentBytes: 536870912,
 *             },
 *         },
 *     ],
 *     users: [
 *         {
 *             name: "producer-application",
 *             password: "password",
 *             permissions: [{
 *                 role: "ACCESS_ROLE_PRODUCER",
 *                 topicName: "input",
 *             }],
 *         },
 *         {
 *             name: "worker",
 *             password: "password",
 *             permissions: [
 *                 {
 *                     role: "ACCESS_ROLE_CONSUMER",
 *                     topicName: "input",
 *                 },
 *                 {
 *                     role: "ACCESS_ROLE_PRODUCER",
 *                     topicName: "output",
 *                 },
 *             ],
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * A cluster can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/mdbKafkaCluster:MdbKafkaCluster foo cluster_id
 * ```
 */
export class MdbKafkaCluster extends pulumi.CustomResource {
    /**
     * Get an existing MdbKafkaCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MdbKafkaClusterState, opts?: pulumi.CustomResourceOptions): MdbKafkaCluster {
        return new MdbKafkaCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/mdbKafkaCluster:MdbKafkaCluster';

    /**
     * Returns true if the given object is an instance of MdbKafkaCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MdbKafkaCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MdbKafkaCluster.__pulumiType;
    }

    /**
     * Configuration of the Kafka cluster. The structure is documented below.
     */
    public readonly config!: pulumi.Output<outputs.MdbKafkaClusterConfig>;
    /**
     * Timestamp of cluster creation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * Description of the Kafka cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Deployment environment of the Kafka cluster. Can be either `PRESTABLE` or `PRODUCTION`.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Health of the host.
     */
    public /*out*/ readonly health!: pulumi.Output<string>;
    /**
     * A list of IDs of the host groups to place VMs of the cluster on.
     */
    public readonly hostGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * A host of the Kafka cluster. The structure is documented below.
     */
    public /*out*/ readonly hosts!: pulumi.Output<outputs.MdbKafkaClusterHost[]>;
    /**
     * A set of key/value label pairs to assign to the Kafka cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the topic.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network, to which the Kafka cluster belongs.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * Security group ids, to which the Kafka cluster belongs.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
     * For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-kafka/api-ref/Cluster/).
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * IDs of the subnets, to which the Kafka cluster belongs.
     */
    public readonly subnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * A topic of the Kafka cluster. The structure is documented below.
     */
    public readonly topics!: pulumi.Output<outputs.MdbKafkaClusterTopic[] | undefined>;
    /**
     * A user of the Kafka cluster. The structure is documented below.
     */
    public readonly users!: pulumi.Output<outputs.MdbKafkaClusterUser[] | undefined>;

    /**
     * Create a MdbKafkaCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MdbKafkaClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MdbKafkaClusterArgs | MdbKafkaClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MdbKafkaClusterState | undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["health"] = state ? state.health : undefined;
            inputs["hostGroupIds"] = state ? state.hostGroupIds : undefined;
            inputs["hosts"] = state ? state.hosts : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkId"] = state ? state.networkId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["topics"] = state ? state.topics : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as MdbKafkaClusterArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            inputs["config"] = args ? args.config : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["hostGroupIds"] = args ? args.hostGroupIds : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkId"] = args ? args.networkId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["topics"] = args ? args.topics : undefined;
            inputs["users"] = args ? args.users : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["health"] = undefined /*out*/;
            inputs["hosts"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MdbKafkaCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MdbKafkaCluster resources.
 */
export interface MdbKafkaClusterState {
    /**
     * Configuration of the Kafka cluster. The structure is documented below.
     */
    readonly config?: pulumi.Input<inputs.MdbKafkaClusterConfig>;
    /**
     * Timestamp of cluster creation.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * Description of the Kafka cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Deployment environment of the Kafka cluster. Can be either `PRESTABLE` or `PRODUCTION`.
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * Health of the host.
     */
    readonly health?: pulumi.Input<string>;
    /**
     * A list of IDs of the host groups to place VMs of the cluster on.
     */
    readonly hostGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A host of the Kafka cluster. The structure is documented below.
     */
    readonly hosts?: pulumi.Input<pulumi.Input<inputs.MdbKafkaClusterHost>[]>;
    /**
     * A set of key/value label pairs to assign to the Kafka cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the topic.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the network, to which the Kafka cluster belongs.
     */
    readonly networkId?: pulumi.Input<string>;
    /**
     * Security group ids, to which the Kafka cluster belongs.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
     * For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-kafka/api-ref/Cluster/).
     */
    readonly status?: pulumi.Input<string>;
    /**
     * IDs of the subnets, to which the Kafka cluster belongs.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A topic of the Kafka cluster. The structure is documented below.
     */
    readonly topics?: pulumi.Input<pulumi.Input<inputs.MdbKafkaClusterTopic>[]>;
    /**
     * A user of the Kafka cluster. The structure is documented below.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.MdbKafkaClusterUser>[]>;
}

/**
 * The set of arguments for constructing a MdbKafkaCluster resource.
 */
export interface MdbKafkaClusterArgs {
    /**
     * Configuration of the Kafka cluster. The structure is documented below.
     */
    readonly config: pulumi.Input<inputs.MdbKafkaClusterConfig>;
    /**
     * Inhibits deletion of the cluster.  Can be either `true` or `false`.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * Description of the Kafka cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Deployment environment of the Kafka cluster. Can be either `PRESTABLE` or `PRODUCTION`.
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * A list of IDs of the host groups to place VMs of the cluster on.
     */
    readonly hostGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of key/value label pairs to assign to the Kafka cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the topic.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the network, to which the Kafka cluster belongs.
     */
    readonly networkId: pulumi.Input<string>;
    /**
     * Security group ids, to which the Kafka cluster belongs.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IDs of the subnets, to which the Kafka cluster belongs.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A topic of the Kafka cluster. The structure is documented below.
     */
    readonly topics?: pulumi.Input<pulumi.Input<inputs.MdbKafkaClusterTopic>[]>;
    /**
     * A user of the Kafka cluster. The structure is documented below.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.MdbKafkaClusterUser>[]>;
}
