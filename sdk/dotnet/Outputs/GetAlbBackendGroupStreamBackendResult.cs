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
    public sealed class GetAlbBackendGroupStreamBackendResult
    {
        /// <summary>
        /// Healthcheck specification that will be used by this backend. Structure is documented below.
        /// </summary>
        public readonly Outputs.GetAlbBackendGroupStreamBackendHealthcheckResult Healthcheck;
        /// <summary>
        /// Load Balancing Config specification that will be used by this backend. Structure is documented below.
        /// </summary>
        public readonly Outputs.GetAlbBackendGroupStreamBackendLoadBalancingConfigResult LoadBalancingConfig;
        /// <summary>
        /// - Name of the Backend Group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Port for incoming traffic.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// References target groups for the backend.
        /// </summary>
        public readonly ImmutableArray<string> TargetGroupIds;
        /// <summary>
        /// Tls specification that will be used by this backend. Structure is documented below.
        /// </summary>
        public readonly Outputs.GetAlbBackendGroupStreamBackendTlsResult Tls;
        /// <summary>
        /// Weight of the backend. Traffic will be split between backends of the same BackendGroup according to their weights.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private GetAlbBackendGroupStreamBackendResult(
            Outputs.GetAlbBackendGroupStreamBackendHealthcheckResult healthcheck,

            Outputs.GetAlbBackendGroupStreamBackendLoadBalancingConfigResult loadBalancingConfig,

            string name,

            int port,

            ImmutableArray<string> targetGroupIds,

            Outputs.GetAlbBackendGroupStreamBackendTlsResult tls,

            int weight)
        {
            Healthcheck = healthcheck;
            LoadBalancingConfig = loadBalancingConfig;
            Name = name;
            Port = port;
            TargetGroupIds = targetGroupIds;
            Tls = tls;
            Weight = weight;
        }
    }
}