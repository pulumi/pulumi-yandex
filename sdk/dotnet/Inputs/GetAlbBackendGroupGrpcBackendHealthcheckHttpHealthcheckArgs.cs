// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheckArgs : Pulumi.InvokeArgs
    {
        [Input("host", required: true)]
        public string Host { get; set; } = null!;

        [Input("http2", required: true)]
        public bool Http2 { get; set; }

        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheckArgs()
        {
        }
    }
}
