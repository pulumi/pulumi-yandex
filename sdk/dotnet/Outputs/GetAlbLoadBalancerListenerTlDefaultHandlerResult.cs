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
    public sealed class GetAlbLoadBalancerListenerTlDefaultHandlerResult
    {
        public readonly ImmutableArray<string> CertificateIds;
        public readonly ImmutableArray<Outputs.GetAlbLoadBalancerListenerTlDefaultHandlerHttpHandlerResult> HttpHandlers;
        public readonly ImmutableArray<Outputs.GetAlbLoadBalancerListenerTlDefaultHandlerStreamHandlerResult> StreamHandlers;

        [OutputConstructor]
        private GetAlbLoadBalancerListenerTlDefaultHandlerResult(
            ImmutableArray<string> certificateIds,

            ImmutableArray<Outputs.GetAlbLoadBalancerListenerTlDefaultHandlerHttpHandlerResult> httpHandlers,

            ImmutableArray<Outputs.GetAlbLoadBalancerListenerTlDefaultHandlerStreamHandlerResult> streamHandlers)
        {
            CertificateIds = certificateIds;
            HttpHandlers = httpHandlers;
            StreamHandlers = streamHandlers;
        }
    }
}
