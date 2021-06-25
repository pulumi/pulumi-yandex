// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetMdbElasticSearchCluster
    {
        /// <summary>
        /// Get information about a Yandex Managed Elasticsearch cluster. For more information, see
        /// [the official documentation](https://cloud.yandex.com/docs/managed-elasticsearch/concepts).
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
        ///         var foo = Output.Create(Yandex.GetMdbElasticSearchCluster.InvokeAsync(new Yandex.GetMdbElasticSearchClusterArgs
        ///         {
        ///             Name = "test",
        ///         }));
        ///         this.NetworkId = foo.Apply(foo =&gt; foo.NetworkId);
        ///     }
        /// 
        ///     [Output("networkId")]
        ///     public Output&lt;string&gt; NetworkId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMdbElasticSearchClusterResult> InvokeAsync(GetMdbElasticSearchClusterArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMdbElasticSearchClusterResult>("yandex:index/getMdbElasticSearchCluster:getMdbElasticSearchCluster", args ?? new GetMdbElasticSearchClusterArgs(), options.WithVersion());
    }


    public sealed class GetMdbElasticSearchClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Elasticsearch cluster.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        /// <summary>
        /// Description of the Elasticsearch cluster.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("labels")]
        private Dictionary<string, string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Elasticsearch cluster.
        /// </summary>
        public Dictionary<string, string> Labels
        {
            get => _labels ?? (_labels = new Dictionary<string, string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the Elasticsearch cluster.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("securityGroupIds")]
        private List<string>? _securityGroupIds;

        /// <summary>
        /// A set of ids of security groups assigned to hosts of the cluster.
        /// </summary>
        public List<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new List<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// ID of the service account authorized for this cluster.
        /// </summary>
        [Input("serviceAccountId")]
        public string? ServiceAccountId { get; set; }

        public GetMdbElasticSearchClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMdbElasticSearchClusterResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// Configuration of the Elasticsearch cluster. The structure is documented below.
        /// </summary>
        public readonly Outputs.GetMdbElasticSearchClusterConfigResult Config;
        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Description of the Elasticsearch cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Deployment environment of the Elasticsearch cluster.
        /// </summary>
        public readonly string Environment;
        public readonly string FolderId;
        /// <summary>
        /// Aggregated health of the cluster.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// A host of the Elasticsearch cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbElasticSearchClusterHostResult> Hosts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of key/value label pairs to assign to the Elasticsearch cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        /// <summary>
        /// ID of the network, to which the Elasticsearch cluster belongs.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// A set of ids of security groups assigned to hosts of the cluster.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// ID of the service account authorized for this cluster.
        /// </summary>
        public readonly string ServiceAccountId;
        /// <summary>
        /// Status of the cluster.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetMdbElasticSearchClusterResult(
            string clusterId,

            Outputs.GetMdbElasticSearchClusterConfigResult config,

            string createdAt,

            string description,

            string environment,

            string folderId,

            string health,

            ImmutableArray<Outputs.GetMdbElasticSearchClusterHostResult> hosts,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            string networkId,

            ImmutableArray<string> securityGroupIds,

            string serviceAccountId,

            string status)
        {
            ClusterId = clusterId;
            Config = config;
            CreatedAt = createdAt;
            Description = description;
            Environment = environment;
            FolderId = folderId;
            Health = health;
            Hosts = hosts;
            Id = id;
            Labels = labels;
            Name = name;
            NetworkId = networkId;
            SecurityGroupIds = securityGroupIds;
            ServiceAccountId = serviceAccountId;
            Status = status;
        }
    }
}