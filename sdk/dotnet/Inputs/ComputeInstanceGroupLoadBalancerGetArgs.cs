// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceGroupLoadBalancerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The status message of the instance.
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        /// <summary>
        /// A description of the target group.
        /// </summary>
        [Input("targetGroupDescription")]
        public Input<string>? TargetGroupDescription { get; set; }

        [Input("targetGroupId")]
        public Input<string>? TargetGroupId { get; set; }

        [Input("targetGroupLabels")]
        private InputMap<string>? _targetGroupLabels;

        /// <summary>
        /// A set of key/value label pairs.
        /// </summary>
        public InputMap<string> TargetGroupLabels
        {
            get => _targetGroupLabels ?? (_targetGroupLabels = new InputMap<string>());
            set => _targetGroupLabels = value;
        }

        /// <summary>
        /// The name of the target group.
        /// </summary>
        [Input("targetGroupName")]
        public Input<string>? TargetGroupName { get; set; }

        public ComputeInstanceGroupLoadBalancerGetArgs()
        {
        }
    }
}
