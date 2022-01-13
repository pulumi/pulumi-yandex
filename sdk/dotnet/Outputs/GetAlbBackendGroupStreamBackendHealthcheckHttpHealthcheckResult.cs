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
    public sealed class GetAlbBackendGroupStreamBackendHealthcheckHttpHealthcheckResult
    {
        /// <summary>
        /// Optional "Host" HTTP header value.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// If set, health checks will use HTTP2.
        /// </summary>
        public readonly bool Http2;
        /// <summary>
        /// HTTP path.
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private GetAlbBackendGroupStreamBackendHealthcheckHttpHealthcheckResult(
            string host,

            bool http2,

            string path)
        {
            Host = host;
            Http2 = http2;
            Path = path;
        }
    }
}
