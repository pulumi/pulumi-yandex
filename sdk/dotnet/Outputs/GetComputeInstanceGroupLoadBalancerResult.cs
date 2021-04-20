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
    public sealed class GetComputeInstanceGroupLoadBalancerResult
    {
        /// <summary>
        /// The status message of the target group.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// A description of the target group.
        /// </summary>
        public readonly string TargetGroupDescription;
        /// <summary>
        /// The ID of the target group.
        /// </summary>
        public readonly string TargetGroupId;
        /// <summary>
        /// A set of key/value label pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> TargetGroupLabels;
        /// <summary>
        /// The name of the target group.
        /// </summary>
        public readonly string TargetGroupName;

        [OutputConstructor]
        private GetComputeInstanceGroupLoadBalancerResult(
            string statusMessage,

            string targetGroupDescription,

            string targetGroupId,

            ImmutableDictionary<string, string> targetGroupLabels,

            string targetGroupName)
        {
            StatusMessage = statusMessage;
            TargetGroupDescription = targetGroupDescription;
            TargetGroupId = targetGroupId;
            TargetGroupLabels = targetGroupLabels;
            TargetGroupName = targetGroupName;
        }
    }
}
