// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional service name for grpc.health.v1.HealthCheckRequest message.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetAlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckInputArgs()
        {
        }
    }
}
