// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbClickhouseClusterHostGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets whether the host should get a public IP address on creation. Can be either `true` or `false`.
        /// </summary>
        [Input("assignPublicIp")]
        public Input<bool>? AssignPublicIp { get; set; }

        /// <summary>
        /// The fully qualified domain name of the host.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The name of the shard to which the host belongs.
        /// </summary>
        [Input("shardName")]
        public Input<string>? ShardName { get; set; }

        /// <summary>
        /// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Type of the model.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The availability zone where the ClickHouse host will be created.
        /// For more information see [the official documentation](https://cloud.yandex.com/docs/overview/concepts/geo-scope).
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public MdbClickhouseClusterHostGetArgs()
        {
        }
    }
}
