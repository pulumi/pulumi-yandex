// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbKafkaClusterHostGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets whether the host should get a public IP address on creation. Can be either `true` or `false`.
        /// </summary>
        [Input("assignPublicIp")]
        public Input<bool>? AssignPublicIp { get; set; }

        /// <summary>
        /// Health of the host.
        /// </summary>
        [Input("health")]
        public Input<string>? Health { get; set; }

        /// <summary>
        /// The name of the topic.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The role type to grant to the topic.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The ID of the subnet, to which the host belongs.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The availability zone where the Kafka host was created.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public MdbKafkaClusterHostGetArgs()
        {
        }
    }
}
