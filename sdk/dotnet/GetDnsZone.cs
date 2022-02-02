// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetDnsZone
    {
        /// <summary>
        /// Get information about a DNS Zone.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Yandex = Pulumi.Yandex;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(Yandex.GetDnsZone.InvokeAsync(new Yandex.GetDnsZoneArgs
        ///         {
        ///             DnsZoneId = yandex_dns_zone.Zone1.Id,
        ///         }));
        ///         this.Zone = foo.Apply(foo =&gt; foo.Zone);
        ///     }
        /// 
        ///     [Output("zone")]
        ///     public Output&lt;string&gt; Zone { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDnsZoneResult> InvokeAsync(GetDnsZoneArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDnsZoneResult>("yandex:index/getDnsZone:getDnsZone", args ?? new GetDnsZoneArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a DNS Zone.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Yandex = Pulumi.Yandex;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(Yandex.GetDnsZone.InvokeAsync(new Yandex.GetDnsZoneArgs
        ///         {
        ///             DnsZoneId = yandex_dns_zone.Zone1.Id,
        ///         }));
        ///         this.Zone = foo.Apply(foo =&gt; foo.Zone);
        ///     }
        /// 
        ///     [Output("zone")]
        ///     public Output&lt;string&gt; Zone { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDnsZoneResult> Invoke(GetDnsZoneInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDnsZoneResult>("yandex:index/getDnsZone:getDnsZone", args ?? new GetDnsZoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDnsZoneArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the DNS Zone.
        /// </summary>
        [Input("dnsZoneId")]
        public string? DnsZoneId { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// - Name of the DNS Zone.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetDnsZoneArgs()
        {
        }
    }

    public sealed class GetDnsZoneInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the DNS Zone.
        /// </summary>
        [Input("dnsZoneId")]
        public Input<string>? DnsZoneId { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// - Name of the DNS Zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetDnsZoneInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDnsZoneResult
    {
        /// <summary>
        /// (Computed) The DNS zone creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// (Computed) Description of the DNS zone.
        /// </summary>
        public readonly string Description;
        public readonly string DnsZoneId;
        /// <summary>
        /// (Computed) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) A set of key/value label pairs to assign to the DNS zone.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// (Computed) User assigned name of a specific resource. Must be unique within the folder.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (Computed) For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        /// </summary>
        public readonly ImmutableArray<string> PrivateNetworks;
        /// <summary>
        /// (Computed) The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
        /// </summary>
        public readonly bool Public;
        /// <summary>
        /// (Computed) The DNS name of this zone, e.g. "example.com.". Must ends with dot.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetDnsZoneResult(
            string createdAt,

            string description,

            string dnsZoneId,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<string> privateNetworks,

            bool @public,

            string zone)
        {
            CreatedAt = createdAt;
            Description = description;
            DnsZoneId = dnsZoneId;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Name = name;
            PrivateNetworks = privateNetworks;
            Public = @public;
            Zone = zone;
        }
    }
}
