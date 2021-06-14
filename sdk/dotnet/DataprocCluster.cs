// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    /// <summary>
    /// Manages a Data Proc cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/data-proc/).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fooVpcNetwork = new Yandex.VpcNetwork("fooVpcNetwork", new Yandex.VpcNetworkArgs
    ///         {
    ///         });
    ///         var fooVpcSubnet = new Yandex.VpcSubnet("fooVpcSubnet", new Yandex.VpcSubnetArgs
    ///         {
    ///             Zone = "ru-central1-b",
    ///             NetworkId = fooVpcNetwork.Id,
    ///             V4CidrBlocks = 
    ///             {
    ///                 "10.1.0.0/24",
    ///             },
    ///         });
    ///         var dataprocIamServiceAccount = new Yandex.IamServiceAccount("dataprocIamServiceAccount", new Yandex.IamServiceAccountArgs
    ///         {
    ///             Description = "service account to manage Dataproc Cluster",
    ///         });
    ///         var fooResourcemanagerFolder = Output.Create(Yandex.GetResourcemanagerFolder.InvokeAsync(new Yandex.GetResourcemanagerFolderArgs
    ///         {
    ///             FolderId = "some_folder_id",
    ///         }));
    ///         var dataprocResourcemanagerFolderIamBinding = new Yandex.ResourcemanagerFolderIamBinding("dataprocResourcemanagerFolderIamBinding", new Yandex.ResourcemanagerFolderIamBindingArgs
    ///         {
    ///             FolderId = fooResourcemanagerFolder.Apply(fooResourcemanagerFolder =&gt; fooResourcemanagerFolder.Id),
    ///             Role = "mdb.dataproc.agent",
    ///             Members = 
    ///             {
    ///                 dataprocIamServiceAccount.Id.Apply(id =&gt; $"serviceAccount:{id}"),
    ///             },
    ///         });
    ///         // required in order to create bucket
    ///         var bucket_creator = new Yandex.ResourcemanagerFolderIamBinding("bucket-creator", new Yandex.ResourcemanagerFolderIamBindingArgs
    ///         {
    ///             FolderId = fooResourcemanagerFolder.Apply(fooResourcemanagerFolder =&gt; fooResourcemanagerFolder.Id),
    ///             Role = "editor",
    ///             Members = 
    ///             {
    ///                 dataprocIamServiceAccount.Id.Apply(id =&gt; $"serviceAccount:{id}"),
    ///             },
    ///         });
    ///         var fooIamServiceAccountStaticAccessKey = new Yandex.IamServiceAccountStaticAccessKey("fooIamServiceAccountStaticAccessKey", new Yandex.IamServiceAccountStaticAccessKeyArgs
    ///         {
    ///             ServiceAccountId = dataprocIamServiceAccount.Id,
    ///         });
    ///         var fooStorageBucket = new Yandex.StorageBucket("fooStorageBucket", new Yandex.StorageBucketArgs
    ///         {
    ///             Bucket = "foo",
    ///             AccessKey = fooIamServiceAccountStaticAccessKey.AccessKey,
    ///             SecretKey = fooIamServiceAccountStaticAccessKey.SecretKey,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 bucket_creator,
    ///             },
    ///         });
    ///         var fooDataprocCluster = new Yandex.DataprocCluster("fooDataprocCluster", new Yandex.DataprocClusterArgs
    ///         {
    ///             Bucket = fooStorageBucket.Bucket,
    ///             Description = "Dataproc Cluster created by Terraform",
    ///             Labels = 
    ///             {
    ///                 { "created_by", "terraform" },
    ///             },
    ///             ServiceAccountId = dataprocIamServiceAccount.Id,
    ///             ZoneId = "ru-central1-b",
    ///             ClusterConfig = new Yandex.Inputs.DataprocClusterClusterConfigArgs
    ///             {
    ///                 Hadoop = new Yandex.Inputs.DataprocClusterClusterConfigHadoopArgs
    ///                 {
    ///                     Services = 
    ///                     {
    ///                         "HDFS",
    ///                         "YARN",
    ///                         "SPARK",
    ///                         "TEZ",
    ///                         "MAPREDUCE",
    ///                         "HIVE",
    ///                     },
    ///                     Properties = 
    ///                     {
    ///                         { "yarn:yarn.resourcemanager.am.max-attempts", "5" },
    ///                     },
    ///                     SshPublicKeys = 
    ///                     {
    ///                         File.ReadAllText("~/.ssh/id_rsa.pub"),
    ///                     },
    ///                 },
    ///                 SubclusterSpecs = 
    ///                 {
    ///                     new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecArgs
    ///                     {
    ///                         Name = "main",
    ///                         Role = "MASTERNODE",
    ///                         Resources = new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecResourcesArgs
    ///                         {
    ///                             ResourcePresetId = "s2.small",
    ///                             DiskTypeId = "network-hdd",
    ///                             DiskSize = 20,
    ///                         },
    ///                         SubnetId = fooVpcSubnet.Id,
    ///                         HostsCount = 1,
    ///                     },
    ///                     new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecArgs
    ///                     {
    ///                         Name = "data",
    ///                         Role = "DATANODE",
    ///                         Resources = new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecResourcesArgs
    ///                         {
    ///                             ResourcePresetId = "s2.small",
    ///                             DiskTypeId = "network-hdd",
    ///                             DiskSize = 20,
    ///                         },
    ///                         SubnetId = fooVpcSubnet.Id,
    ///                         HostsCount = 2,
    ///                     },
    ///                     new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecArgs
    ///                     {
    ///                         Name = "compute",
    ///                         Role = "COMPUTENODE",
    ///                         Resources = new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecResourcesArgs
    ///                         {
    ///                             ResourcePresetId = "s2.small",
    ///                             DiskTypeId = "network-hdd",
    ///                             DiskSize = 20,
    ///                         },
    ///                         SubnetId = fooVpcSubnet.Id,
    ///                         HostsCount = 2,
    ///                     },
    ///                     new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecArgs
    ///                     {
    ///                         Name = "compute_autoscaling",
    ///                         Role = "COMPUTENODE",
    ///                         Resources = new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecResourcesArgs
    ///                         {
    ///                             ResourcePresetId = "s2.small",
    ///                             DiskTypeId = "network-hdd",
    ///                             DiskSize = 20,
    ///                         },
    ///                         SubnetId = fooVpcSubnet.Id,
    ///                         HostsCount = 2,
    ///                         AutoscalingConfig = new Yandex.Inputs.DataprocClusterClusterConfigSubclusterSpecAutoscalingConfigArgs
    ///                         {
    ///                             MaxHostsCount = 10,
    ///                             MeasurementDuration = 60,
    ///                             WarmupDuration = 60,
    ///                             StabilizationDuration = 120,
    ///                             Preemptible = false,
    ///                             DecommissionTimeout = 60,
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 dataprocResourcemanagerFolderIamBinding,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// A cluster can be imported using the `id` of the resource, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import yandex:index/dataprocCluster:DataprocCluster foo cluster_id
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/dataprocCluster:DataprocCluster")]
    public partial class DataprocCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Object Storage bucket to use for Data Proc jobs. Data Proc Agent saves output of job driver's process to specified bucket. In order for this to work service account (specified by the `service_account_id` argument) should be given permission to create objects within this bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string?> Bucket { get; private set; } = null!;

        /// <summary>
        /// Configuration and resources for hosts that should be created with the cluster. The structure is documented below.
        /// </summary>
        [Output("clusterConfig")]
        public Output<Outputs.DataprocClusterClusterConfig> ClusterConfig { get; private set; } = null!;

        /// <summary>
        /// (Computed) The Data Proc cluster creation timestamp.
        /// * `cluster_config.0.subcluster_spec.X.id` - (Computed) ID of the subcluster.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Description of the Data Proc cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ID of the folder to create a cluster in. If it is not provided, the default provider folder is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// A list of host group IDs to place VMs of the cluster on.
        /// </summary>
        [Output("hostGroupIds")]
        public Output<ImmutableArray<string>> HostGroupIds { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the Data Proc cluster.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the Data Proc subcluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs that the cluster belongs to.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Service account to be used by the Data Proc agent to access resources of Yandex.Cloud. Selected service account should have `mdb.dataproc.agent` role on the folder where the Data Proc cluster will be located.
        /// </summary>
        [Output("serviceAccountId")]
        public Output<string> ServiceAccountId { get; private set; } = null!;

        /// <summary>
        /// Whether to enable UI Proxy feature.
        /// </summary>
        [Output("uiProxy")]
        public Output<bool?> UiProxy { get; private set; } = null!;

        /// <summary>
        /// ID of the availability zone to create cluster in. If it is not provided, the default provider zone is used.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a DataprocCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataprocCluster(string name, DataprocClusterArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/dataprocCluster:DataprocCluster", name, args ?? new DataprocClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataprocCluster(string name, Input<string> id, DataprocClusterState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/dataprocCluster:DataprocCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataprocCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataprocCluster Get(string name, Input<string> id, DataprocClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new DataprocCluster(name, id, state, options);
        }
    }

    public sealed class DataprocClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Object Storage bucket to use for Data Proc jobs. Data Proc Agent saves output of job driver's process to specified bucket. In order for this to work service account (specified by the `service_account_id` argument) should be given permission to create objects within this bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Configuration and resources for hosts that should be created with the cluster. The structure is documented below.
        /// </summary>
        [Input("clusterConfig", required: true)]
        public Input<Inputs.DataprocClusterClusterConfigArgs> ClusterConfig { get; set; } = null!;

        /// <summary>
        /// Description of the Data Proc cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the folder to create a cluster in. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("hostGroupIds")]
        private InputList<string>? _hostGroupIds;

        /// <summary>
        /// A list of host group IDs to place VMs of the cluster on.
        /// </summary>
        public InputList<string> HostGroupIds
        {
            get => _hostGroupIds ?? (_hostGroupIds = new InputList<string>());
            set => _hostGroupIds = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Data Proc cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the Data Proc subcluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs that the cluster belongs to.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Service account to be used by the Data Proc agent to access resources of Yandex.Cloud. Selected service account should have `mdb.dataproc.agent` role on the folder where the Data Proc cluster will be located.
        /// </summary>
        [Input("serviceAccountId", required: true)]
        public Input<string> ServiceAccountId { get; set; } = null!;

        /// <summary>
        /// Whether to enable UI Proxy feature.
        /// </summary>
        [Input("uiProxy")]
        public Input<bool>? UiProxy { get; set; }

        /// <summary>
        /// ID of the availability zone to create cluster in. If it is not provided, the default provider zone is used.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DataprocClusterArgs()
        {
        }
    }

    public sealed class DataprocClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Object Storage bucket to use for Data Proc jobs. Data Proc Agent saves output of job driver's process to specified bucket. In order for this to work service account (specified by the `service_account_id` argument) should be given permission to create objects within this bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Configuration and resources for hosts that should be created with the cluster. The structure is documented below.
        /// </summary>
        [Input("clusterConfig")]
        public Input<Inputs.DataprocClusterClusterConfigGetArgs>? ClusterConfig { get; set; }

        /// <summary>
        /// (Computed) The Data Proc cluster creation timestamp.
        /// * `cluster_config.0.subcluster_spec.X.id` - (Computed) ID of the subcluster.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the Data Proc cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the folder to create a cluster in. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("hostGroupIds")]
        private InputList<string>? _hostGroupIds;

        /// <summary>
        /// A list of host group IDs to place VMs of the cluster on.
        /// </summary>
        public InputList<string> HostGroupIds
        {
            get => _hostGroupIds ?? (_hostGroupIds = new InputList<string>());
            set => _hostGroupIds = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Data Proc cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the Data Proc subcluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs that the cluster belongs to.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Service account to be used by the Data Proc agent to access resources of Yandex.Cloud. Selected service account should have `mdb.dataproc.agent` role on the folder where the Data Proc cluster will be located.
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        /// <summary>
        /// Whether to enable UI Proxy feature.
        /// </summary>
        [Input("uiProxy")]
        public Input<bool>? UiProxy { get; set; }

        /// <summary>
        /// ID of the availability zone to create cluster in. If it is not provided, the default provider zone is used.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DataprocClusterState()
        {
        }
    }
}
