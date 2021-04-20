// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Outputs
{

    [OutputType]
    public sealed class GetMdbRedisClusterHostResult
    {
        /// <summary>
        /// The fully qualified domain name of the host.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The name of the shard to which the host belongs.
        /// </summary>
        public readonly string ShardName;
        /// <summary>
        /// The ID of the subnet, to which the host belongs. The subnet must
        /// be a part of the network to which the cluster belongs.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The availability zone where the Redis host will be created.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetMdbRedisClusterHostResult(
            string fqdn,

            string shardName,

            string subnetId,

            string zone)
        {
            Fqdn = fqdn;
            ShardName = shardName;
            SubnetId = subnetId;
            Zone = zone;
        }
    }
}
