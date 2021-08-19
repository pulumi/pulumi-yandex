// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, will enable only HTTP1 protocol with HTTP1.0 support.
        /// </summary>
        [Input("allowHttp10")]
        public Input<bool>? AllowHttp10 { get; set; }

        /// <summary>
        /// If set, will enable HTTP2 protocol for the handler. The structure is documented below.
        /// </summary>
        [Input("http2Options")]
        public Input<Inputs.AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerHttp2OptionsArgs>? Http2Options { get; set; }

        /// <summary>
        /// HTTP router id.
        /// </summary>
        [Input("httpRouterId")]
        public Input<string>? HttpRouterId { get; set; }

        public AlbLoadBalancerListenerTlsSniHandlerHandlerHttpHandlerArgs()
        {
        }
    }
}
