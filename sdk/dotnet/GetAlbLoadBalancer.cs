// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetAlbLoadBalancer
    {
        public static Task<GetAlbLoadBalancerResult> InvokeAsync(GetAlbLoadBalancerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlbLoadBalancerResult>("yandex:index/getAlbLoadBalancer:getAlbLoadBalancer", args ?? new GetAlbLoadBalancerArgs(), options.WithVersion());
    }


    public sealed class GetAlbLoadBalancerArgs : Pulumi.InvokeArgs
    {
        [Input("loadBalancerId")]
        public string? LoadBalancerId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetAlbLoadBalancerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlbLoadBalancerResult
    {
        public readonly Outputs.GetAlbLoadBalancerAllocationPolicyResult AllocationPolicy;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<Outputs.GetAlbLoadBalancerListenerResult> Listeners;
        public readonly string LoadBalancerId;
        public readonly string LogGroupId;
        public readonly string Name;
        public readonly string NetworkId;
        public readonly string RegionId;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly string Status;

        [OutputConstructor]
        private GetAlbLoadBalancerResult(
            Outputs.GetAlbLoadBalancerAllocationPolicyResult allocationPolicy,

            string createdAt,

            string description,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GetAlbLoadBalancerListenerResult> listeners,

            string loadBalancerId,

            string logGroupId,

            string name,

            string networkId,

            string regionId,

            ImmutableArray<string> securityGroupIds,

            string status)
        {
            AllocationPolicy = allocationPolicy;
            CreatedAt = createdAt;
            Description = description;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Listeners = listeners;
            LoadBalancerId = loadBalancerId;
            LogGroupId = logGroupId;
            Name = name;
            NetworkId = networkId;
            RegionId = regionId;
            SecurityGroupIds = securityGroupIds;
            Status = status;
        }
    }
}