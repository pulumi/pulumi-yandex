// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class LbNetworkLoadBalancerListenerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// External IP address specification. The structure is documented below.
        /// </summary>
        [Input("externalAddressSpec")]
        public Input<Inputs.LbNetworkLoadBalancerListenerExternalAddressSpecArgs>? ExternalAddressSpec { get; set; }

        /// <summary>
        /// Internal IP address specification. The structure is documented below.
        /// </summary>
        [Input("internalAddressSpec")]
        public Input<Inputs.LbNetworkLoadBalancerListenerInternalAddressSpecArgs>? InternalAddressSpec { get; set; }

        /// <summary>
        /// Name of the listener. The name must be unique for each listener on a single load balancer.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Port for incoming traffic.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Protocol for incoming traffic. TCP or UDP and the default is TCP.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Port of a target. The default is the same as listener's port.
        /// </summary>
        [Input("targetPort")]
        public Input<int>? TargetPort { get; set; }

        public LbNetworkLoadBalancerListenerArgs()
        {
        }
    }
}
