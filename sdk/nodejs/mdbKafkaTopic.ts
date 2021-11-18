// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a topic of a Kafka cluster within the Yandex.Cloud. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = new yandex.MdbKafkaCluster("foo", {
 *     networkId: "c64vs98keiqc7f24pvkd",
 *     config: {
 *         version: "2.8",
 *         zones: ["ru-central1-a"],
 *         unmanagedTopics: true,
 *         kafka: {
 *             resources: {
 *                 resourcePresetId: "s2.micro",
 *                 diskTypeId: "network-hdd",
 *                 diskSize: 16,
 *             },
 *         },
 *     },
 * });
 * const events = new yandex.MdbKafkaTopic("events", {
 *     clusterId: foo.id,
 *     partitions: 4,
 *     replicationFactor: 1,
 *     topicConfig: {
 *         cleanupPolicy: "CLEANUP_POLICY_COMPACT",
 *         compressionType: "COMPRESSION_TYPE_LZ4",
 *         deleteRetentionMs: 86400000,
 *         fileDeleteDelayMs: 60000,
 *         flushMessages: 128,
 *         flushMs: 1000,
 *         minCompactionLagMs: 0,
 *         retentionBytes: 10737418240,
 *         retentionMs: 604800000,
 *         maxMessageBytes: 1048588,
 *         minInsyncReplicas: 1,
 *         segmentBytes: 268435456,
 *         preallocate: true,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Kafka topic can be imported using following format
 *
 * ```sh
 *  $ pulumi import yandex:index/mdbKafkaTopic:MdbKafkaTopic foo {{cluster_id}}:{{topic_name}}
 * ```
 */
export class MdbKafkaTopic extends pulumi.CustomResource {
    /**
     * Get an existing MdbKafkaTopic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MdbKafkaTopicState, opts?: pulumi.CustomResourceOptions): MdbKafkaTopic {
        return new MdbKafkaTopic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/mdbKafkaTopic:MdbKafkaTopic';

    /**
     * Returns true if the given object is an instance of MdbKafkaTopic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MdbKafkaTopic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MdbKafkaTopic.__pulumiType;
    }

    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The name of the topic.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of the topic's partitions.
     */
    public readonly partitions!: pulumi.Output<number>;
    /**
     * Amount of data copies (replicas) for the topic in the cluster.
     */
    public readonly replicationFactor!: pulumi.Output<number>;
    /**
     * User-defined settings for the topic. The structure is documented below.
     */
    public readonly topicConfig!: pulumi.Output<outputs.MdbKafkaTopicTopicConfig | undefined>;

    /**
     * Create a MdbKafkaTopic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MdbKafkaTopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MdbKafkaTopicArgs | MdbKafkaTopicState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MdbKafkaTopicState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partitions"] = state ? state.partitions : undefined;
            inputs["replicationFactor"] = state ? state.replicationFactor : undefined;
            inputs["topicConfig"] = state ? state.topicConfig : undefined;
        } else {
            const args = argsOrState as MdbKafkaTopicArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.partitions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitions'");
            }
            if ((!args || args.replicationFactor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicationFactor'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partitions"] = args ? args.partitions : undefined;
            inputs["replicationFactor"] = args ? args.replicationFactor : undefined;
            inputs["topicConfig"] = args ? args.topicConfig : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MdbKafkaTopic.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MdbKafkaTopic resources.
 */
export interface MdbKafkaTopicState {
    clusterId?: pulumi.Input<string>;
    /**
     * The name of the topic.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of the topic's partitions.
     */
    partitions?: pulumi.Input<number>;
    /**
     * Amount of data copies (replicas) for the topic in the cluster.
     */
    replicationFactor?: pulumi.Input<number>;
    /**
     * User-defined settings for the topic. The structure is documented below.
     */
    topicConfig?: pulumi.Input<inputs.MdbKafkaTopicTopicConfig>;
}

/**
 * The set of arguments for constructing a MdbKafkaTopic resource.
 */
export interface MdbKafkaTopicArgs {
    clusterId: pulumi.Input<string>;
    /**
     * The name of the topic.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of the topic's partitions.
     */
    partitions: pulumi.Input<number>;
    /**
     * Amount of data copies (replicas) for the topic in the cluster.
     */
    replicationFactor: pulumi.Input<number>;
    /**
     * User-defined settings for the topic. The structure is documented below.
     */
    topicConfig?: pulumi.Input<inputs.MdbKafkaTopicTopicConfig>;
}
