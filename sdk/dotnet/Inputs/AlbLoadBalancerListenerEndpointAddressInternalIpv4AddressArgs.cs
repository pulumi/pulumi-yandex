// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Provided by the client or computed automatically.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Provided by the client or computed automatically.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressArgs()
        {
        }
    }
}