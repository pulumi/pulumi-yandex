// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetVpcSubnet
    {
        /// <summary>
        /// Get information about a Yandex VPC subnet. For more information, see
        /// [Yandex.Cloud VPC](https://cloud.yandex.com/docs/vpc/concepts/index).
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Yandex = Pulumi.Yandex;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var admin = Output.Create(Yandex.GetVpcSubnet.InvokeAsync(new Yandex.GetVpcSubnetArgs
        ///         {
        ///             SubnetId = "my-subnet-id",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// This data source is used to define [VPC Subnets] that can be used by other resources.
        /// </summary>
        public static Task<GetVpcSubnetResult> InvokeAsync(GetVpcSubnetArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcSubnetResult>("yandex:index/getVpcSubnet:getVpcSubnet", args ?? new GetVpcSubnetArgs(), options.WithVersion());
    }


    public sealed class GetVpcSubnetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// - Name of the subnet.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Subnet ID.
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        public GetVpcSubnetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcSubnetResult
    {
        /// <summary>
        /// Creation timestamp of this subnet.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Description of the subnet.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Options for DHCP client. The structure is documented below.
        /// </summary>
        public readonly Outputs.GetVpcSubnetDhcpOptionsResult DhcpOptions;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Labels to assign to this subnet.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        /// <summary>
        /// ID of the network this subnet belongs to.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// ID of the route table to assign to this subnet.
        /// </summary>
        public readonly string RouteTableId;
        public readonly string SubnetId;
        /// <summary>
        /// The blocks of internal IPv4 addresses owned by this subnet.
        /// </summary>
        public readonly ImmutableArray<string> V4CidrBlocks;
        /// <summary>
        /// The blocks of internal IPv6 addresses owned by this subnet.
        /// </summary>
        public readonly ImmutableArray<string> V6CidrBlocks;
        /// <summary>
        /// Name of the availability zone for this subnet.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetVpcSubnetResult(
            string createdAt,

            string description,

            Outputs.GetVpcSubnetDhcpOptionsResult dhcpOptions,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            string networkId,

            string routeTableId,

            string subnetId,

            ImmutableArray<string> v4CidrBlocks,

            ImmutableArray<string> v6CidrBlocks,

            string zone)
        {
            CreatedAt = createdAt;
            Description = description;
            DhcpOptions = dhcpOptions;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Name = name;
            NetworkId = networkId;
            RouteTableId = routeTableId;
            SubnetId = subnetId;
            V4CidrBlocks = v4CidrBlocks;
            V6CidrBlocks = v6CidrBlocks;
            Zone = zone;
        }
    }
}
