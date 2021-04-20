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
    public sealed class MdbKafkaClusterConfigZookeeperResources
    {
        /// <summary>
        /// Volume of the storage available to a ZooKeeper host, in gigabytes.
        /// </summary>
        public readonly int DiskSize;
        /// <summary>
        /// Type of the storage of ZooKeeper hosts.
        /// For more information see [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts/storage).
        /// </summary>
        public readonly string DiskTypeId;
        public readonly string ResourcePresetId;

        [OutputConstructor]
        private MdbKafkaClusterConfigZookeeperResources(
            int diskSize,

            string diskTypeId,

            string resourcePresetId)
        {
            DiskSize = diskSize;
            DiskTypeId = diskTypeId;
            ResourcePresetId = resourcePresetId;
        }
    }
}
