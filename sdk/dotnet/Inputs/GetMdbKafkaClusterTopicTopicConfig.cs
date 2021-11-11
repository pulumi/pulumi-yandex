// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbKafkaClusterTopicTopicConfigArgs : Pulumi.InvokeArgs
    {
        [Input("cleanupPolicy")]
        public string? CleanupPolicy { get; set; }

        [Input("compressionType")]
        public string? CompressionType { get; set; }

        [Input("deleteRetentionMs")]
        public string? DeleteRetentionMs { get; set; }

        [Input("fileDeleteDelayMs")]
        public string? FileDeleteDelayMs { get; set; }

        [Input("flushMessages")]
        public string? FlushMessages { get; set; }

        [Input("flushMs")]
        public string? FlushMs { get; set; }

        [Input("maxMessageBytes")]
        public string? MaxMessageBytes { get; set; }

        [Input("minCompactionLagMs")]
        public string? MinCompactionLagMs { get; set; }

        [Input("minInsyncReplicas")]
        public string? MinInsyncReplicas { get; set; }

        [Input("preallocate")]
        public bool? Preallocate { get; set; }

        [Input("retentionBytes")]
        public string? RetentionBytes { get; set; }

        [Input("retentionMs")]
        public string? RetentionMs { get; set; }

        [Input("segmentBytes")]
        public string? SegmentBytes { get; set; }

        public GetMdbKafkaClusterTopicTopicConfigArgs()
        {
        }
    }
}
