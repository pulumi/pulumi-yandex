// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbKafkaClusterConfigArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The flag that defines whether a public IP address is assigned to the node.
        /// </summary>
        [Input("assignPublicIp")]
        public bool? AssignPublicIp { get; set; }

        /// <summary>
        /// (Optional) Count of brokers per availability zone.
        /// </summary>
        [Input("brokersCount")]
        public int? BrokersCount { get; set; }

        /// <summary>
        /// (Optional) Configuration of the Kafka subcluster. The structure is documented below.
        /// </summary>
        [Input("kafka", required: true)]
        public Inputs.GetMdbKafkaClusterConfigKafkaArgs Kafka { get; set; } = null!;

        /// <summary>
        /// (Optional) Enables managed schema registry on cluster. Can be either `true` or `false`.
        /// </summary>
        [Input("schemaRegistry")]
        public bool? SchemaRegistry { get; set; }

        /// <summary>
        /// (Optional) Allows to use Kafka AdminAPI to manage topics. Can be either `true` or `false`.
        /// </summary>
        [Input("unmanagedTopics")]
        public bool? UnmanagedTopics { get; set; }

        /// <summary>
        /// (Required) Version of the Kafka server software.
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        [Input("zones", required: true)]
        private List<string>? _zones;

        /// <summary>
        /// (Optional) List of availability zones.
        /// </summary>
        public List<string> Zones
        {
            get => _zones ?? (_zones = new List<string>());
            set => _zones = value;
        }

        /// <summary>
        /// (Optional) Configuration of the ZooKeeper subcluster. The structure is documented below.
        /// </summary>
        [Input("zookeeper", required: true)]
        public Inputs.GetMdbKafkaClusterConfigZookeeperArgs Zookeeper { get; set; } = null!;

        public GetMdbKafkaClusterConfigArgs()
        {
        }
    }
}
