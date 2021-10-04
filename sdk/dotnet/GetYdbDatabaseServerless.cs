// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetYdbDatabaseServerless
    {
        /// <summary>
        /// Get information about a Yandex Database serverless cluster.
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
        ///         var myDatabase = Output.Create(Yandex.GetYdbDatabaseServerless.InvokeAsync(new Yandex.GetYdbDatabaseServerlessArgs
        ///         {
        ///             DatabaseId = "some_ydb_serverless_database_id",
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
        public static Task<GetYdbDatabaseServerlessResult> InvokeAsync(GetYdbDatabaseServerlessArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetYdbDatabaseServerlessResult>("yandex:index/getYdbDatabaseServerless:getYdbDatabaseServerless", args ?? new GetYdbDatabaseServerlessArgs(), options.WithVersion());
    }


    public sealed class GetYdbDatabaseServerlessArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the Yandex Database serverless cluster.
        /// </summary>
        [Input("databaseId")]
        public string? DatabaseId { get; set; }

        /// <summary>
        /// ID of the folder that the Yandex Database serverless cluster belongs to.
        /// It will be deduced from provider configuration if not set explicitly.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// Name of the Yandex Database serverless cluster.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetYdbDatabaseServerlessArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetYdbDatabaseServerlessResult
    {
        /// <summary>
        /// The Yandex Database serverless cluster creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        public readonly string? DatabaseId;
        /// <summary>
        /// Full database path of the Yandex Database serverless cluster.
        /// Useful for SDK configuration.
        /// </summary>
        public readonly string DatabasePath;
        /// <summary>
        /// A description of the Yandex Database serverless cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Document API endpoint of the Yandex Database serverless cluster.
        /// </summary>
        public readonly string DocumentApiEndpoint;
        public readonly string? FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of key/value label pairs assigned to the Yandex Database serverless cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Location ID of the Yandex Database serverless cluster.
        /// </summary>
        public readonly string LocationId;
        public readonly string? Name;
        /// <summary>
        /// Status of the Yandex Database serverless cluster.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Whether TLS is enabled for the Yandex Database serverless cluster.
        /// Useful for SDK configuration.
        /// </summary>
        public readonly bool TlsEnabled;
        /// <summary>
        /// API endpoint of the Yandex Database serverless cluster.
        /// Useful for SDK configuration.
        /// </summary>
        public readonly string YdbApiEndpoint;
        /// <summary>
        /// Full endpoint of the Yandex Database serverless cluster.
        /// </summary>
        public readonly string YdbFullEndpoint;

        [OutputConstructor]
        private GetYdbDatabaseServerlessResult(
            string createdAt,

            string? databaseId,

            string databasePath,

            string description,

            string documentApiEndpoint,

            string? folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string locationId,

            string? name,

            string status,

            bool tlsEnabled,

            string ydbApiEndpoint,

            string ydbFullEndpoint)
        {
            CreatedAt = createdAt;
            DatabaseId = databaseId;
            DatabasePath = databasePath;
            Description = description;
            DocumentApiEndpoint = documentApiEndpoint;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            LocationId = locationId;
            Name = name;
            Status = status;
            TlsEnabled = tlsEnabled;
            YdbApiEndpoint = ydbApiEndpoint;
            YdbFullEndpoint = ydbFullEndpoint;
        }
    }
}
