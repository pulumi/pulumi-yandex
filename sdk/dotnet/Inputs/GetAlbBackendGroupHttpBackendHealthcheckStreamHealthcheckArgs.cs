// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupHttpBackendHealthcheckStreamHealthcheckInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional text to search in reply.
        /// </summary>
        [Input("receive", required: true)]
        public Input<string> Receive { get; set; } = null!;

        /// <summary>
        /// Optional message to send. If empty, it's a connect-only health check.
        /// </summary>
        [Input("send", required: true)]
        public Input<string> Send { get; set; } = null!;

        public GetAlbBackendGroupHttpBackendHealthcheckStreamHealthcheckInputArgs()
        {
        }
    }
}
