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
    public sealed class AlbBackendGroupGrpcBackendHealthcheck
    {
        /// <summary>
        /// Grpc Healthcheck specification that will be used by this healthcheck. Structure is documented below.
        /// </summary>
        public readonly Outputs.AlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheck? GrpcHealthcheck;
        /// <summary>
        /// Optional alternative port for health checking.
        /// </summary>
        public readonly int? HealthcheckPort;
        /// <summary>
        /// Number of consecutive successful health checks required to promote endpoint into the healthy state. 0 means 1. Note that during startup, only a single successful health check is required to mark a host healthy.
        /// </summary>
        public readonly int? HealthyThreshold;
        /// <summary>
        /// Http Healthcheck specification that will be used by this healthcheck. Structure is documented below.
        /// </summary>
        public readonly Outputs.AlbBackendGroupGrpcBackendHealthcheckHttpHealthcheck? HttpHealthcheck;
        /// <summary>
        /// Interval between health checks.
        /// </summary>
        public readonly string Interval;
        /// <summary>
        /// An optional jitter amount as a percentage of interval. If specified, during every interval value of (interval_ms * interval_jitter_percent / 100) will be added to the wait time.
        /// </summary>
        public readonly double? IntervalJitterPercent;
        /// <summary>
        /// Stream Healthcheck specification that will be used by this healthcheck. Structure is documented below.
        /// </summary>
        public readonly Outputs.AlbBackendGroupGrpcBackendHealthcheckStreamHealthcheck? StreamHealthcheck;
        /// <summary>
        /// Time to wait for a health check response.
        /// </summary>
        public readonly string Timeout;
        /// <summary>
        /// Number of consecutive failed health checks required to demote endpoint into the unhealthy state. 0 means 1. Note that for HTTP health checks, a single 503 immediately makes endpoint unhealthy.
        /// </summary>
        public readonly int? UnhealthyThreshold;

        [OutputConstructor]
        private AlbBackendGroupGrpcBackendHealthcheck(
            Outputs.AlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheck? grpcHealthcheck,

            int? healthcheckPort,

            int? healthyThreshold,

            Outputs.AlbBackendGroupGrpcBackendHealthcheckHttpHealthcheck? httpHealthcheck,

            string interval,

            double? intervalJitterPercent,

            Outputs.AlbBackendGroupGrpcBackendHealthcheckStreamHealthcheck? streamHealthcheck,

            string timeout,

            int? unhealthyThreshold)
        {
            GrpcHealthcheck = grpcHealthcheck;
            HealthcheckPort = healthcheckPort;
            HealthyThreshold = healthyThreshold;
            HttpHealthcheck = httpHealthcheck;
            Interval = interval;
            IntervalJitterPercent = intervalJitterPercent;
            StreamHealthcheck = streamHealthcheck;
            Timeout = timeout;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}
