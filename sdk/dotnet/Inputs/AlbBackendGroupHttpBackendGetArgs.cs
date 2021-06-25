// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbBackendGroupHttpBackendGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Healthcheck specification that will be used by this backend. Structure is documented below.
        /// </summary>
        [Input("healthcheck")]
        public Input<Inputs.AlbBackendGroupHttpBackendHealthcheckGetArgs>? Healthcheck { get; set; }

        /// <summary>
        /// If set, health checks will use HTTP2.
        /// </summary>
        [Input("http2")]
        public Input<bool>? Http2 { get; set; }

        /// <summary>
        /// Load Balancing Config specification that will be used by this backend. Structure is documented below.
        /// </summary>
        [Input("loadBalancingConfig")]
        public Input<Inputs.AlbBackendGroupHttpBackendLoadBalancingConfigGetArgs>? LoadBalancingConfig { get; set; }

        /// <summary>
        /// Name of the backend.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Port for incoming traffic.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("targetGroupIds", required: true)]
        private InputList<string>? _targetGroupIds;

        /// <summary>
        /// References target groups for the backend.
        /// </summary>
        public InputList<string> TargetGroupIds
        {
            get => _targetGroupIds ?? (_targetGroupIds = new InputList<string>());
            set => _targetGroupIds = value;
        }

        /// <summary>
        /// Tls specification that will be used by this backend. Structure is documented below.
        /// </summary>
        [Input("tls")]
        public Input<Inputs.AlbBackendGroupHttpBackendTlsGetArgs>? Tls { get; set; }

        /// <summary>
        /// Weight of the backend. Traffic will be split between backends of the same BackendGroup according to their weights.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public AlbBackendGroupHttpBackendGetArgs()
        {
        }
    }
}