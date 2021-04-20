// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbRedisClusterConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Redis key eviction policy for a dataset that reaches maximum memory.
        /// Can be any of the listed in [the official RedisDB documentation](https://docs.redislabs.com/latest/rs/administering/database-operations/eviction-policy/).
        /// </summary>
        [Input("maxmemoryPolicy")]
        public Input<string>? MaxmemoryPolicy { get; set; }

        /// <summary>
        /// Password for the Redis cluster.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Close the connection after a client is idle for N seconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Version of Redis (either 5.0 or 6.0).
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public MdbRedisClusterConfigGetArgs()
        {
        }
    }
}
