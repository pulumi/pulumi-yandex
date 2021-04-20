// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceGroupScalePolicyTestAutoScaleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target CPU load level.
        /// </summary>
        [Input("cpuUtilizationTarget")]
        public Input<double>? CpuUtilizationTarget { get; set; }

        [Input("customRules")]
        private InputList<Inputs.ComputeInstanceGroupScalePolicyTestAutoScaleCustomRuleArgs>? _customRules;

        /// <summary>
        /// A list of custom rules. The structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceGroupScalePolicyTestAutoScaleCustomRuleArgs> CustomRules
        {
            get => _customRules ?? (_customRules = new InputList<Inputs.ComputeInstanceGroupScalePolicyTestAutoScaleCustomRuleArgs>());
            set => _customRules = value;
        }

        /// <summary>
        /// The initial number of instances in the instance group.
        /// </summary>
        [Input("initialSize", required: true)]
        public Input<int> InitialSize { get; set; } = null!;

        /// <summary>
        /// The maximum number of virtual machines in the group.
        /// </summary>
        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        /// <summary>
        /// The amount of time, in seconds, that metrics are averaged for.
        /// If the average value at the end of the interval is higher than the `cpu_utilization_target`,
        /// the instance group will increase the number of virtual machines in the group.
        /// </summary>
        [Input("measurementDuration", required: true)]
        public Input<int> MeasurementDuration { get; set; } = null!;

        /// <summary>
        /// The minimum number of virtual machines in a single availability zone.
        /// </summary>
        [Input("minZoneSize")]
        public Input<int>? MinZoneSize { get; set; }

        /// <summary>
        /// The minimum time interval, in seconds, to monitor the load before
        /// an instance group can reduce the number of virtual machines in the group. During this time, the group
        /// will not decrease even if the average load falls below the value of `cpu_utilization_target`.
        /// </summary>
        [Input("stabilizationDuration")]
        public Input<int>? StabilizationDuration { get; set; }

        /// <summary>
        /// The warm-up time of the virtual machine, in seconds. During this time,
        /// traffic is fed to the virtual machine, but load metrics are not taken into account.
        /// </summary>
        [Input("warmupDuration")]
        public Input<int>? WarmupDuration { get; set; }

        public ComputeInstanceGroupScalePolicyTestAutoScaleArgs()
        {
        }
    }
}
