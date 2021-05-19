// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetAlbTargetGroup
    {
        public static Task<GetAlbTargetGroupResult> InvokeAsync(GetAlbTargetGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlbTargetGroupResult>("yandex:index/getAlbTargetGroup:getAlbTargetGroup", args ?? new GetAlbTargetGroupArgs(), options.WithVersion());
    }


    public sealed class GetAlbTargetGroupArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("targetGroupId")]
        public string? TargetGroupId { get; set; }

        public GetAlbTargetGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlbTargetGroupResult
    {
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly string TargetGroupId;
        public readonly ImmutableArray<Outputs.GetAlbTargetGroupTargetResult> Targets;

        [OutputConstructor]
        private GetAlbTargetGroupResult(
            string createdAt,

            string description,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            string targetGroupId,

            ImmutableArray<Outputs.GetAlbTargetGroupTargetResult> targets)
        {
            CreatedAt = createdAt;
            Description = description;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Name = name;
            TargetGroupId = targetGroupId;
            Targets = targets;
        }
    }
}
