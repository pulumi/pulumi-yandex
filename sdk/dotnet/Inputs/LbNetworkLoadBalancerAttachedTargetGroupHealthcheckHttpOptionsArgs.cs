// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL path to set for health checking requests for every target in the target group. For example `/ping`. The default path is `/`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Port for incoming traffic.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptionsArgs()
        {
        }
    }
}
