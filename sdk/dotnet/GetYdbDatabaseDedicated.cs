// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Yandex
{
    public static class GetYdbDatabaseDedicated
    {
        /// <summary>
        /// Get information about a Yandex Database (dedicated) cluster.
        /// For more information, see [the official documentation](https://cloud.yandex.com/en/docs/ydb/concepts/serverless_and_dedicated).
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
        ///         var myDatabase = Output.Create(Yandex.GetYdbDatabaseDedicated.InvokeAsync(new Yandex.GetYdbDatabaseDedicatedArgs
        ///         {
        ///             DatabaseId = "some_ydb_dedicated_database_id",
        ///         }));
        ///         this.YdbApiEndpoint = myDatabase.Apply(myDatabase =&gt; myDatabase.YdbApiEndpoint);
        ///     }
        /// 
        ///     [Output("ydbApiEndpoint")]
        ///     public Output&lt;string&gt; YdbApiEndpoint { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetYdbDatabaseDedicatedResult> InvokeAsync(GetYdbDatabaseDedicatedArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetYdbDatabaseDedicatedResult>("yandex:index/getYdbDatabaseDedicated:getYdbDatabaseDedicated", args ?? new GetYdbDatabaseDedicatedArgs(), options.WithVersion());

        /// <summary>
        /// Get information about a Yandex Database (dedicated) cluster.
        /// For more information, see [the official documentation](https://cloud.yandex.com/en/docs/ydb/concepts/serverless_and_dedicated).
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
        ///         var myDatabase = Output.Create(Yandex.GetYdbDatabaseDedicated.InvokeAsync(new Yandex.GetYdbDatabaseDedicatedArgs
        ///         {
        ///             DatabaseId = "some_ydb_dedicated_database_id",
        ///         }));
        ///         this.YdbApiEndpoint = myDatabase.Apply(myDatabase =&gt; myDatabase.YdbApiEndpoint);
        ///     }
        /// 
        ///     [Output("ydbApiEndpoint")]
        ///     public Output&lt;string&gt; YdbApiEndpoint { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetYdbDatabaseDedicatedResult> Invoke(GetYdbDatabaseDedicatedInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetYdbDatabaseDedicatedResult>("yandex:index/getYdbDatabaseDedicated:getYdbDatabaseDedicated", args ?? new GetYdbDatabaseDedicatedInvokeArgs(), options.WithVersion());
    }


    public sealed class GetYdbDatabaseDedicatedArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the Yandex Database cluster.
        /// </summary>
        [Input("databaseId")]
        public string? DatabaseId { get; set; }

        /// <summary>
        /// ID of the folder that the Yandex Database cluster belongs to.
        /// It will be deduced from provider configuration if not set explicitly.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// Name of the Yandex Database cluster.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetYdbDatabaseDedicatedArgs()
        {
        }
    }

    public sealed class GetYdbDatabaseDedicatedInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the Yandex Database cluster.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// ID of the folder that the Yandex Database cluster belongs to.
        /// It will be deduced from provider configuration if not set explicitly.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Name of the Yandex Database cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetYdbDatabaseDedicatedInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetYdbDatabaseDedicatedResult
    {
        /// <summary>
        /// Whether public IP addresses are assigned to the Yandex Database cluster.
        /// </summary>
        public readonly bool AssignPublicIps;
        /// <summary>
        /// The Yandex Database cluster creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        public readonly string? DatabaseId;
        /// <summary>
        /// Full database path of the Yandex Database cluster.
        /// Useful for SDK configuration.
        /// </summary>
        public readonly string DatabasePath;
        /// <summary>
        /// A description of the Yandex Database cluster.
        /// </summary>
        public readonly string Description;
        public readonly string? FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of key/value label pairs assigned to the Yandex Database cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Location of the Yandex Database cluster.
        /// The structure is documented below.
        /// </summary>
        public readonly Outputs.GetYdbDatabaseDedicatedLocationResult Location;
        /// <summary>
        /// Location ID of the Yandex Database cluster.
        /// </summary>
        public readonly string LocationId;
        public readonly string? Name;
        /// <summary>
        /// ID of the network the Yandex Database cluster is attached to.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// The Yandex Database cluster preset.
        /// </summary>
        public readonly string ResourcePresetId;
        /// <summary>
        /// Scaling policy of the Yandex Database cluster.
        /// The structure is documented below.
        /// </summary>
        public readonly Outputs.GetYdbDatabaseDedicatedScalePolicyResult ScalePolicy;
        /// <summary>
        /// Status of the Yandex Database cluster.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A list of storage configuration options of the Yandex Database cluster.
        /// The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetYdbDatabaseDedicatedStorageConfigResult> StorageConfigs;
        /// <summary>
        /// List of subnet IDs the Yandex Database cluster is attached to.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// Whether TLS is enabled for the Yandex Database cluster.
        /// Useful for SDK configuration.
        /// </summary>
        public readonly bool TlsEnabled;
        /// <summary>
        /// API endpoint of the Yandex Database cluster.
        /// Useful for SDK configuration.
        /// </summary>
        public readonly string YdbApiEndpoint;
        /// <summary>
        /// Full endpoint of the Yandex Database cluster.
        /// </summary>
        public readonly string YdbFullEndpoint;

        [OutputConstructor]
        private GetYdbDatabaseDedicatedResult(
            bool assignPublicIps,

            string createdAt,

            string? databaseId,

            string databasePath,

            string description,

            string? folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            Outputs.GetYdbDatabaseDedicatedLocationResult location,

            string locationId,

            string? name,

            string networkId,

            string resourcePresetId,

            Outputs.GetYdbDatabaseDedicatedScalePolicyResult scalePolicy,

            string status,

            ImmutableArray<Outputs.GetYdbDatabaseDedicatedStorageConfigResult> storageConfigs,

            ImmutableArray<string> subnetIds,

            bool tlsEnabled,

            string ydbApiEndpoint,

            string ydbFullEndpoint)
        {
            AssignPublicIps = assignPublicIps;
            CreatedAt = createdAt;
            DatabaseId = databaseId;
            DatabasePath = databasePath;
            Description = description;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Location = location;
            LocationId = locationId;
            Name = name;
            NetworkId = networkId;
            ResourcePresetId = resourcePresetId;
            ScalePolicy = scalePolicy;
            Status = status;
            StorageConfigs = storageConfigs;
            SubnetIds = subnetIds;
            TlsEnabled = tlsEnabled;
            YdbApiEndpoint = ydbApiEndpoint;
            YdbFullEndpoint = ydbFullEndpoint;
        }
    }
}
