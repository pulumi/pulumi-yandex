// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbClickhouseClusterUserQuotaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of queries that threw exception.
        /// </summary>
        [Input("errors")]
        public Input<int>? Errors { get; set; }

        /// <summary>
        /// The total query execution time, in milliseconds (wall time).
        /// </summary>
        [Input("executionTime")]
        public Input<int>? ExecutionTime { get; set; }

        /// <summary>
        /// Duration of interval for quota in milliseconds.
        /// </summary>
        [Input("intervalDuration", required: true)]
        public Input<int> IntervalDuration { get; set; } = null!;

        /// <summary>
        /// The total number of queries.
        /// </summary>
        [Input("queries")]
        public Input<int>? Queries { get; set; }

        /// <summary>
        /// The total number of source rows read from tables for running the query, on all remote servers.
        /// </summary>
        [Input("readRows")]
        public Input<int>? ReadRows { get; set; }

        /// <summary>
        /// The total number of rows given as the result.
        /// </summary>
        [Input("resultRows")]
        public Input<int>? ResultRows { get; set; }

        public MdbClickhouseClusterUserQuotaArgs()
        {
        }
    }
}
