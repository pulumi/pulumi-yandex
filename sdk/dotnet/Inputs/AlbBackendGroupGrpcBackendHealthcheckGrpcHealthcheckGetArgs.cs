// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheckGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Service name for grpc.health.v1.HealthCheckRequest message.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public AlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheckGetArgs()
        {
        }
    }
}
