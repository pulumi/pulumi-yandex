// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class VpcRouteTableStaticRouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Route prefix in CIDR notation.
        /// </summary>
        [Input("destinationPrefix")]
        public Input<string>? DestinationPrefix { get; set; }

        /// <summary>
        /// Address of the next hop.
        /// </summary>
        [Input("nextHopAddress")]
        public Input<string>? NextHopAddress { get; set; }

        public VpcRouteTableStaticRouteArgs()
        {
        }
    }
}
