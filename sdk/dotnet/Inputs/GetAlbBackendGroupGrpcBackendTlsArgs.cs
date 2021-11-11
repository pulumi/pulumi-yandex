// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupGrpcBackendTlsInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [SNI](https://en.wikipedia.org/wiki/Server_Name_Indication) string for TLS connections.
        /// * `validation_context.0.trusted_ca_id` - Trusted CA certificate ID in the Certificate Manager.
        /// * `validation_context.0.trusted_ca_bytes` - PEM-encoded trusted CA certificate chain.
        /// </summary>
        [Input("sni", required: true)]
        public Input<string> Sni { get; set; } = null!;

        [Input("validationContext", required: true)]
        public Input<Inputs.GetAlbBackendGroupGrpcBackendTlsValidationContextInputArgs> ValidationContext { get; set; } = null!;

        public GetAlbBackendGroupGrpcBackendTlsInputArgs()
        {
        }
    }
}
