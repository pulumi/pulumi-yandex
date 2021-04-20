// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetApiGateway
    {
        public static Task<GetApiGatewayResult> InvokeAsync(GetApiGatewayArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiGatewayResult>("yandex:index/getApiGateway:getApiGateway", args ?? new GetApiGatewayArgs(), options.WithVersion());
    }


    public sealed class GetApiGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("apiGatewayId")]
        public string? ApiGatewayId { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetApiGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiGatewayResult
    {
        public readonly string? ApiGatewayId;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string Domain;
        public readonly string? FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string LogGroupId;
        public readonly string? Name;
        public readonly string Status;

        [OutputConstructor]
        private GetApiGatewayResult(
            string? apiGatewayId,

            string createdAt,

            string description,

            string domain,

            string? folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string logGroupId,

            string? name,

            string status)
        {
            ApiGatewayId = apiGatewayId;
            CreatedAt = createdAt;
            Description = description;
            Domain = domain;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            LogGroupId = logGroupId;
            Name = name;
            Status = status;
        }
    }
}
