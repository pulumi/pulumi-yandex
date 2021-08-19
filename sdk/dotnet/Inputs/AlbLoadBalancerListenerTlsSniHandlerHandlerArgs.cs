// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerTlsSniHandlerHandlerArgs : Pulumi.ResourceArgs
    {
        [Input("certificateIds", required: true)]
        private InputList<string>? _certificateIds;

        /// <summary>
        /// Certificate IDs in the Certificate Manager. Multiple TLS certificates can be associated
        /// with the same context to allow both RSA and ECDSA certificates. Only the first certificate of each type will be used.
        /// </summary>
        public InputList<string> CertificateIds
        {
            get => _certificateIds ?? (_certificateIds = new InputList<string>());
            set => _certificateIds = value;
        }

        /// <summary>
        /// HTTP handler resource. The structure is documented below.
        /// </summary>
        [Input("httpHandler")]
        public Input<Inputs.AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerArgs>? HttpHandler { get; set; }

        public AlbLoadBalancerListenerTlsSniHandlerHandlerArgs()
        {
        }
    }
}
