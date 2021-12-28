// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbBackendGroupStreamBackendHealthcheckGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Grpc Healthcheck specification that will be used by this healthcheck. Structure is documented below.
        /// </summary>
        [Input("grpcHealthcheck")]
        public Input<Inputs.AlbBackendGroupStreamBackendHealthcheckGrpcHealthcheckGetArgs>? GrpcHealthcheck { get; set; }

        /// <summary>
        /// Optional alternative port for health checking.
        /// </summary>
        [Input("healthcheckPort")]
        public Input<int>? HealthcheckPort { get; set; }

        /// <summary>
        /// Number of consecutive successful health checks required to promote endpoint into the healthy state. 0 means 1. Note that during startup, only a single successful health check is required to mark a host healthy.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// Http Healthcheck specification that will be used by this healthcheck. Structure is documented below.
        /// </summary>
        [Input("httpHealthcheck")]
        public Input<Inputs.AlbBackendGroupStreamBackendHealthcheckHttpHealthcheckGetArgs>? HttpHealthcheck { get; set; }

        /// <summary>
        /// Interval between health checks.
        /// </summary>
        [Input("interval", required: true)]
        public Input<string> Interval { get; set; } = null!;

        /// <summary>
        /// An optional jitter amount as a percentage of interval. If specified, during every interval value of (interval_ms * interval_jitter_percent / 100) will be added to the wait time.
        /// </summary>
        [Input("intervalJitterPercent")]
        public Input<double>? IntervalJitterPercent { get; set; }

        /// <summary>
        /// Stream Healthcheck specification that will be used by this healthcheck. Structure is documented below.
        /// </summary>
        [Input("streamHealthcheck")]
        public Input<Inputs.AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckGetArgs>? StreamHealthcheck { get; set; }

        /// <summary>
        /// Time to wait for a health check response.
        /// </summary>
        [Input("timeout", required: true)]
        public Input<string> Timeout { get; set; } = null!;

        /// <summary>
        /// Number of consecutive failed health checks required to demote endpoint into the unhealthy state. 0 means 1. Note that for HTTP health checks, a single 503 immediately makes endpoint unhealthy.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public AlbBackendGroupStreamBackendHealthcheckGetArgs()
        {
        }
    }
}
