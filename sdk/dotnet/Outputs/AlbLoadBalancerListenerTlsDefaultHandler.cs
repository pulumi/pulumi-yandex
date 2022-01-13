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
    public sealed class AlbLoadBalancerListenerTlsDefaultHandler
    {
        /// <summary>
        /// Certificate IDs in the Certificate Manager. Multiple TLS certificates can be associated
        /// with the same context to allow both RSA and ECDSA certificates. Only the first certificate of each type will be used.
        /// </summary>
        public readonly ImmutableArray<string> CertificateIds;
        /// <summary>
        /// HTTP handler resource. The structure is documented below.
        /// </summary>
        public readonly Outputs.AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler? HttpHandler;
        /// <summary>
        /// Stream handler resource. The structure is documented below.
        /// </summary>
        public readonly Outputs.AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler? StreamHandler;

        [OutputConstructor]
        private AlbLoadBalancerListenerTlsDefaultHandler(
            ImmutableArray<string> certificateIds,

            Outputs.AlbLoadBalancerListenerTlsDefaultHandlerHttpHandler? httpHandler,

            Outputs.AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler? streamHandler)
        {
            CertificateIds = certificateIds;
            HttpHandler = httpHandler;
            StreamHandler = streamHandler;
        }
    }
}
