// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbKafkaClusterTopicInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Kafka cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// (Required) The number of the topic's partitions.
        /// </summary>
        [Input("partitions", required: true)]
        public Input<int> Partitions { get; set; } = null!;

        /// <summary>
        /// (Required) Amount of data copies (replicas) for the topic in the cluster.
        /// </summary>
        [Input("replicationFactor", required: true)]
        public Input<int> ReplicationFactor { get; set; } = null!;

        /// <summary>
        /// (Required) User-defined settings for the topic. The structure is documented below.
        /// </summary>
        [Input("topicConfig")]
        public Input<Inputs.GetMdbKafkaClusterTopicTopicConfigInputArgs>? TopicConfig { get; set; }

        public GetMdbKafkaClusterTopicInputArgs()
        {
        }
    }
}
