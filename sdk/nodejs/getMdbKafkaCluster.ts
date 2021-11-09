// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Managed Kafka cluster. For more information, see
 * [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo = pulumi.output(yandex.getMdbKafkaCluster({
 *     name: "test",
 * }));
 *
 * export const networkId = foo.networkId;
 * ```
 */
export function getMdbKafkaCluster(args?: GetMdbKafkaClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetMdbKafkaClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getMdbKafkaCluster:getMdbKafkaCluster", {
        "clusterId": args.clusterId,
        "config": args.config,
        "deletionProtection": args.deletionProtection,
        "folderId": args.folderId,
        "name": args.name,
        "subnetIds": args.subnetIds,
        "topics": args.topics,
        "users": args.users,
    }, opts);
}

/**
 * A collection of arguments for invoking getMdbKafkaCluster.
 */
export interface GetMdbKafkaClusterArgs {
    /**
     * The ID of the Kafka cluster.
     */
    clusterId?: string;
    /**
     * Configuration of the Kafka cluster. The structure is documented below.
     */
    config?: inputs.GetMdbKafkaClusterConfig;
    deletionProtection?: boolean;
    /**
     * The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
     */
    folderId?: string;
    /**
     * The name of the Kafka cluster.
     */
    name?: string;
    subnetIds?: string[];
    /**
     * A topic of the Kafka cluster. The structure is documented below.
     */
    topics?: inputs.GetMdbKafkaClusterTopic[];
    /**
     * A user of the Kafka cluster. The structure is documented below.
     */
    users?: inputs.GetMdbKafkaClusterUser[];
}

/**
 * A collection of values returned by getMdbKafkaCluster.
 */
export interface GetMdbKafkaClusterResult {
    readonly clusterId: string;
    /**
     * Configuration of the Kafka cluster. The structure is documented below.
     */
    readonly config?: outputs.GetMdbKafkaClusterConfig;
    /**
     * Creation timestamp of the key.
     */
    readonly createdAt: string;
    readonly deletionProtection: boolean;
    /**
     * Description of the Kafka cluster.
     */
    readonly description: string;
    /**
     * Deployment environment of the Kafka cluster.
     */
    readonly environment: string;
    readonly folderId: string;
    /**
     * Health of the host.
     */
    readonly health: string;
    /**
     * A list of IDs of the host groups hosting VMs of the cluster.
     */
    readonly hostGroupIds: string[];
    /**
     * A host of the Kafka cluster. The structure is documented below.
     */
    readonly hosts: outputs.GetMdbKafkaClusterHost[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of key/value label pairs to assign to the Kafka cluster.
     */
    readonly labels: {[key: string]: string};
    /**
     * The fully qualified domain name of the host.
     */
    readonly name: string;
    /**
     * ID of the network, to which the Kafka cluster belongs.
     */
    readonly networkId: string;
    /**
     * A list of security groups IDs of the Kafka cluster.
     */
    readonly securityGroupIds: string[];
    /**
     * Status of the cluster.
     */
    readonly status: string;
    readonly subnetIds?: string[];
    /**
     * A topic of the Kafka cluster. The structure is documented below.
     */
    readonly topics?: outputs.GetMdbKafkaClusterTopic[];
    /**
     * A user of the Kafka cluster. The structure is documented below.
     */
    readonly users?: outputs.GetMdbKafkaClusterUser[];
}

export function getMdbKafkaClusterOutput(args?: GetMdbKafkaClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMdbKafkaClusterResult> {
    return pulumi.output(args).apply(a => getMdbKafkaCluster(a, opts))
}

/**
 * A collection of arguments for invoking getMdbKafkaCluster.
 */
export interface GetMdbKafkaClusterOutputArgs {
    /**
     * The ID of the Kafka cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Configuration of the Kafka cluster. The structure is documented below.
     */
    config?: pulumi.Input<inputs.GetMdbKafkaClusterConfigArgs>;
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * The name of the Kafka cluster.
     */
    name?: pulumi.Input<string>;
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A topic of the Kafka cluster. The structure is documented below.
     */
    topics?: pulumi.Input<pulumi.Input<inputs.GetMdbKafkaClusterTopicArgs>[]>;
    /**
     * A user of the Kafka cluster. The structure is documented below.
     */
    users?: pulumi.Input<pulumi.Input<inputs.GetMdbKafkaClusterUserArgs>[]>;
}
