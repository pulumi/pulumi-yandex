// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupHttpBackendHealthcheckHttpHealthcheckArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Optional "Host" HTTP header value.
        /// </summary>
        [Input("host", required: true)]
        public string Host { get; set; } = null!;

        /// <summary>
        /// If set, health checks will use HTTP2.
        /// </summary>
        [Input("http2", required: true)]
        public bool Http2 { get; set; }

        /// <summary>
        /// HTTP path.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetAlbBackendGroupHttpBackendHealthcheckHttpHealthcheckArgs()
        {
        }
    }
}
