// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetMdbMongodbCluster
    {
        /// <summary>
        /// Get information about a Yandex Managed MongoDB cluster. For more information, see
        /// [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts).
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
        ///         var foo = Output.Create(Yandex.GetMdbMongodbCluster.InvokeAsync(new Yandex.GetMdbMongodbClusterArgs
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
        public static Task<GetMdbMongodbClusterResult> InvokeAsync(GetMdbMongodbClusterArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMdbMongodbClusterResult>("yandex:index/getMdbMongodbCluster:getMdbMongodbCluster", args ?? new GetMdbMongodbClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Managed MongoDB cluster. For more information, see
        /// [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts).
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
        ///         var foo = Output.Create(Yandex.GetMdbMongodbCluster.InvokeAsync(new Yandex.GetMdbMongodbClusterArgs
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
        public static Output<GetMdbMongodbClusterResult> Invoke(GetMdbMongodbClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMdbMongodbClusterResult>("yandex:index/getMdbMongodbCluster:getMdbMongodbCluster", args ?? new GetMdbMongodbClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMdbMongodbClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the MongoDB cluster.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("deletionProtection")]
        public bool? DeletionProtection { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// The name of the MongoDB cluster.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetMdbMongodbClusterArgs()
        {
        }
    }

    public sealed class GetMdbMongodbClusterInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the MongoDB cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// The name of the MongoDB cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetMdbMongodbClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMdbMongodbClusterResult
    {
        /// <summary>
        /// Configuration of the MongoDB cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMongodbClusterClusterConfigResult> ClusterConfigs;
        public readonly string ClusterId;
        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A database of the MongoDB cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMongodbClusterDatabaseResult> Databases;
        public readonly bool DeletionProtection;
        /// <summary>
        /// Description of the MongoDB cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Deployment environment of the MongoDB cluster.
        /// </summary>
        public readonly string Environment;
        public readonly string FolderId;
        /// <summary>
        /// The health of the host.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// A host of the MongoDB cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMongodbClusterHostResult> Hosts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of key/value label pairs to assign to the MongoDB cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<Outputs.GetMdbMongodbClusterMaintenanceWindowResult> MaintenanceWindows;
        /// <summary>
        /// The name of the database.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the network, to which the MongoDB cluster belongs.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMongodbClusterResourceResult> Resources;
        /// <summary>
        /// A set of ids of security groups assigned to hosts of the cluster.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// MongoDB Cluster mode enabled/disabled.
        /// </summary>
        public readonly bool Sharded;
        /// <summary>
        /// Status of the cluster.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A user of the MongoDB cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMongodbClusterUserResult> Users;

        [OutputConstructor]
        private GetMdbMongodbClusterResult(
            ImmutableArray<Outputs.GetMdbMongodbClusterClusterConfigResult> clusterConfigs,

            string clusterId,

            string createdAt,

            ImmutableArray<Outputs.GetMdbMongodbClusterDatabaseResult> databases,

            bool deletionProtection,

            string description,

            string environment,

            string folderId,

            string health,

            ImmutableArray<Outputs.GetMdbMongodbClusterHostResult> hosts,

            string id,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GetMdbMongodbClusterMaintenanceWindowResult> maintenanceWindows,

            string name,

            string networkId,

            ImmutableArray<Outputs.GetMdbMongodbClusterResourceResult> resources,

            ImmutableArray<string> securityGroupIds,

            bool sharded,

            string status,

            ImmutableArray<Outputs.GetMdbMongodbClusterUserResult> users)
        {
            ClusterConfigs = clusterConfigs;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Databases = databases;
            DeletionProtection = deletionProtection;
            Description = description;
            Environment = environment;
            FolderId = folderId;
            Health = health;
            Hosts = hosts;
            Id = id;
            Labels = labels;
            MaintenanceWindows = maintenanceWindows;
            Name = name;
            NetworkId = networkId;
            Resources = resources;
            SecurityGroupIds = securityGroupIds;
            Sharded = sharded;
            Status = status;
            Users = users;
        }
    }
}
