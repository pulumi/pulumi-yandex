// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbVirtualHostRouteGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// GRPC route resource. The structure is documented below.
        /// </summary>
        [Input("grpcRoute")]
        public Input<Inputs.AlbVirtualHostRouteGrpcRouteGetArgs>? GrpcRoute { get; set; }

        /// <summary>
        /// HTTP route resource. The structure is documented below.
        /// </summary>
        [Input("httpRoute")]
        public Input<Inputs.AlbVirtualHostRouteHttpRouteGetArgs>? HttpRoute { get; set; }

        /// <summary>
        /// name of the route.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AlbVirtualHostRouteGetArgs()
        {
        }
    }
}
