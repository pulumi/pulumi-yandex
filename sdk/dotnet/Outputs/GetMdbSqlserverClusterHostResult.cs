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
    public sealed class GetMdbSqlserverClusterHostResult
    {
        /// <summary>
        /// Sets whether the host should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment
        /// </summary>
        public readonly bool AssignPublicIp;
        /// <summary>
        /// The fully qualified domain name of the host.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The availability zone where the SQLServer host will be created.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetMdbSqlserverClusterHostResult(
            bool assignPublicIp,

            string fqdn,

            string subnetId,

            string zone)
        {
            AssignPublicIp = assignPublicIp;
            Fqdn = fqdn;
            SubnetId = subnetId;
            Zone = zone;
        }
    }
}
