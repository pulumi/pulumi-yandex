// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbKafkaClusterConfigZookeeperInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Resources allocated to hosts of the ZooKeeper subcluster. The structure is documented below.
        /// </summary>
        [Input("resources", required: true)]
        public Input<Inputs.GetMdbKafkaClusterConfigZookeeperResourcesInputArgs> Resources { get; set; } = null!;

        public GetMdbKafkaClusterConfigZookeeperInputArgs()
        {
        }
    }
}
