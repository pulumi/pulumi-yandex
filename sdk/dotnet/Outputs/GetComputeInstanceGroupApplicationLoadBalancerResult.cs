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
    public sealed class GetComputeInstanceGroupApplicationLoadBalancerResult
    {
        /// <summary>
        /// Timeout for waiting for the VM to be checked by the load balancer. If the timeout is exceeded, the VM will be turned off based on the deployment policy. Specified in seconds.
        /// </summary>
        public readonly int MaxOpeningTrafficDuration;
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
        private GetComputeInstanceGroupApplicationLoadBalancerResult(
            int maxOpeningTrafficDuration,

            string statusMessage,

            string targetGroupDescription,

            string targetGroupId,

            ImmutableDictionary<string, string> targetGroupLabels,

            string targetGroupName)
        {
            MaxOpeningTrafficDuration = maxOpeningTrafficDuration;
            StatusMessage = statusMessage;
            TargetGroupDescription = targetGroupDescription;
            TargetGroupId = targetGroupId;
            TargetGroupLabels = targetGroupLabels;
            TargetGroupName = targetGroupName;
        }
    }
}