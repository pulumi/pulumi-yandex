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
    public sealed class ComputeInstanceGroupScalePolicyAutoScaleCustomRule
    {
        /// <summary>
        /// Folder ID of custom metric in Yandex Monitoring that should be used for scaling.
        /// </summary>
        public readonly string? FolderId;
        /// <summary>
        /// A map of labels of metric.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// The name of metric.
        /// </summary>
        public readonly string MetricName;
        /// <summary>
        /// Metric type, `GAUGE` or `COUNTER`.
        /// </summary>
        public readonly string MetricType;
        /// <summary>
        /// Rule type: `UTILIZATION` - This type means that the metric applies to one instance.
        /// First, Instance Groups calculates the average metric value for each instance,
        /// then averages the values for instances in one availability zone.
        /// This type of metric must have the `instance_id` label. `WORKLOAD` - This type means that the metric applies to instances in one availability zone.
        /// This type of metric must have the `zone_id` label.
        /// </summary>
        public readonly string RuleType;
        /// <summary>
        /// Service of custom metric in Yandex Monitoring that should be used for scaling.
        /// </summary>
        public readonly string? Service;
        /// <summary>
        /// Target metric value level.
        /// </summary>
        public readonly double Target;

        [OutputConstructor]
        private ComputeInstanceGroupScalePolicyAutoScaleCustomRule(
            string? folderId,

            ImmutableDictionary<string, string>? labels,

            string metricName,

            string metricType,

            string ruleType,

            string? service,

            double target)
        {
            FolderId = folderId;
            Labels = labels;
            MetricName = metricName;
            MetricType = metricType;
            RuleType = ruleType;
            Service = service;
            Target = target;
        }
    }
}
