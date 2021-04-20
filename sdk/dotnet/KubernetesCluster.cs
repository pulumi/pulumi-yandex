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
    /// Creates a Yandex Kubernetes Cluster.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var zonalClusterResourceName = new Yandex.KubernetesCluster("zonalClusterResourceName", new Yandex.KubernetesClusterArgs
    ///         {
    ///             Description = "description",
    ///             KmsProvider = new Yandex.Inputs.KubernetesClusterKmsProviderArgs
    ///             {
    ///                 KeyId = yandex_kms_symmetric_key.Kms_key_resource_name.Id,
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "my_key", "my_value" },
    ///                 { "my_other_key", "my_other_value" },
    ///             },
    ///             Master = new Yandex.Inputs.KubernetesClusterMasterArgs
    ///             {
    ///                 MaintenancePolicy = new Yandex.Inputs.KubernetesClusterMasterMaintenancePolicyArgs
    ///                 {
    ///                     AutoUpgrade = true,
    ///                     MaintenanceWindow = 
    ///                     {
    ///                         
    ///                         {
    ///                             { "duration", "3h" },
    ///                             { "startTime", "15:00" },
    ///                         },
    ///                     },
    ///                 },
    ///                 PublicIp = true,
    ///                 SecurityGroupIds = 
    ///                 {
    ///                     yandex_vpc_security_group.Security_group_name.Id,
    ///                 },
    ///                 Version = "1.15",
    ///                 Zonal = new Yandex.Inputs.KubernetesClusterMasterZonalArgs
    ///                 {
    ///                     SubnetId = yandex_vpc_subnet.Subnet_resource_name.Id,
    ///                     Zone = yandex_vpc_subnet.Subnet_resource_name.Zone,
    ///                 },
    ///             },
    ///             NetworkId = yandex_vpc_network.Network_resource_name.Id,
    ///             NetworkPolicyProvider = "CALICO",
    ///             NodeServiceAccountId = yandex_iam_service_account.Node_service_account_resource_name.Id,
    ///             ReleaseChannel = "RAPID",
    ///             ServiceAccountId = yandex_iam_service_account.Service_account_resource_name.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var regionalClusterResourceName = new Yandex.KubernetesCluster("regionalClusterResourceName", new Yandex.KubernetesClusterArgs
    ///         {
    ///             Description = "description",
    ///             Labels = 
    ///             {
    ///                 { "my_key", "my_value" },
    ///                 { "my_other_key", "my_other_value" },
    ///             },
    ///             Master = new Yandex.Inputs.KubernetesClusterMasterArgs
    ///             {
    ///                 MaintenancePolicy = new Yandex.Inputs.KubernetesClusterMasterMaintenancePolicyArgs
    ///                 {
    ///                     AutoUpgrade = true,
    ///                     MaintenanceWindow = 
    ///                     {
    ///                         
    ///                         {
    ///                             { "day", "monday" },
    ///                             { "duration", "3h" },
    ///                             { "startTime", "15:00" },
    ///                         },
    ///                         
    ///                         {
    ///                             { "day", "friday" },
    ///                             { "duration", "4h30m" },
    ///                             { "startTime", "10:00" },
    ///                         },
    ///                     },
    ///                 },
    ///                 PublicIp = true,
    ///                 Regional = new Yandex.Inputs.KubernetesClusterMasterRegionalArgs
    ///                 {
    ///                     Location = 
    ///                     {
    ///                         
    ///                         {
    ///                             { "subnetId", yandex_vpc_subnet.Subnet_a_resource_name.Id },
    ///                             { "zone", yandex_vpc_subnet.Subnet_a_resource_name.Zone },
    ///                         },
    ///                         
    ///                         {
    ///                             { "subnetId", yandex_vpc_subnet.Subnet_b_resource_name.Id },
    ///                             { "zone", yandex_vpc_subnet.Subnet_b_resource_name.Zone },
    ///                         },
    ///                         
    ///                         {
    ///                             { "subnetId", yandex_vpc_subnet.Subnet_c_resource_name.Id },
    ///                             { "zone", yandex_vpc_subnet.Subnet_c_resource_name.Zone },
    ///                         },
    ///                     },
    ///                     Region = "ru-central1",
    ///                 },
    ///                 Version = "1.14",
    ///             },
    ///             NetworkId = yandex_vpc_network.Network_resource_name.Id,
    ///             NodeServiceAccountId = yandex_iam_service_account.Node_service_account_resource_name.Id,
    ///             ReleaseChannel = "STABLE",
    ///             ServiceAccountId = yandex_iam_service_account.Service_account_resource_name.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// A Managed Kubernetes cluster can be imported using the `id` of the resource, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import yandex:index/kubernetesCluster:KubernetesCluster default cluster_id
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/kubernetesCluster:KubernetesCluster")]
    public partial class KubernetesCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// CIDR block. IP range for allocating pod addresses.
        /// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
        /// set up for this CIDR blocks in node subnets.
        /// </summary>
        [Output("clusterIpv4Range")]
        public Output<string> ClusterIpv4Range { get; private set; } = null!;

        /// <summary>
        /// (Computed) The Kubernetes cluster creation timestamp.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A description of the Kubernetes cluster.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the folder that the Kubernetes cluster belongs to.
        /// If it is not provided, the default provider folder is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// (Computed) Health of the Kubernetes cluster.
        /// </summary>
        [Output("health")]
        public Output<string> Health { get; private set; } = null!;

        /// <summary>
        /// cluster KMS provider parameters.
        /// </summary>
        [Output("kmsProvider")]
        public Output<Outputs.KubernetesClusterKmsProvider?> KmsProvider { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the Kubernetes cluster.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Kubernetes master configuration options. The structure is documented below.
        /// </summary>
        [Output("master")]
        public Output<Outputs.KubernetesClusterMaster> Master { get; private set; } = null!;

        /// <summary>
        /// Name of a specific Kubernetes cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the cluster network.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Network policy provider for the cluster. Possible values: `CALICO`.
        /// </summary>
        [Output("networkPolicyProvider")]
        public Output<string?> NetworkPolicyProvider { get; private set; } = null!;

        /// <summary>
        /// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
        /// </summary>
        [Output("nodeIpv4CidrMaskSize")]
        public Output<int?> NodeIpv4CidrMaskSize { get; private set; } = null!;

        /// <summary>
        /// Service account to be used by the worker nodes of the Kubernetes cluster
        /// to access Container Registry or to push node logs and metrics.
        /// </summary>
        [Output("nodeServiceAccountId")]
        public Output<string> NodeServiceAccountId { get; private set; } = null!;

        /// <summary>
        /// Cluster release channel.
        /// </summary>
        [Output("releaseChannel")]
        public Output<string> ReleaseChannel { get; private set; } = null!;

        /// <summary>
        /// Service account to be used for provisioning Compute Cloud and VPC resources
        /// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
        /// cluster will be located and on the folder where selected network resides.
        /// </summary>
        [Output("serviceAccountId")]
        public Output<string> ServiceAccountId { get; private set; } = null!;

        /// <summary>
        /// CIDR block. IP range Kubernetes service Kubernetes cluster
        /// IP addresses will be allocated from. It should not overlap with any subnet in the network
        /// the Kubernetes cluster located in.
        /// </summary>
        [Output("serviceIpv4Range")]
        public Output<string> ServiceIpv4Range { get; private set; } = null!;

        /// <summary>
        /// (Computed)Status of the Kubernetes cluster.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesCluster(string name, KubernetesClusterArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/kubernetesCluster:KubernetesCluster", name, args ?? new KubernetesClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesCluster(string name, Input<string> id, KubernetesClusterState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/kubernetesCluster:KubernetesCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KubernetesCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesCluster Get(string name, Input<string> id, KubernetesClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new KubernetesCluster(name, id, state, options);
        }
    }

    public sealed class KubernetesClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CIDR block. IP range for allocating pod addresses.
        /// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
        /// set up for this CIDR blocks in node subnets.
        /// </summary>
        [Input("clusterIpv4Range")]
        public Input<string>? ClusterIpv4Range { get; set; }

        /// <summary>
        /// A description of the Kubernetes cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the folder that the Kubernetes cluster belongs to.
        /// If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// cluster KMS provider parameters.
        /// </summary>
        [Input("kmsProvider")]
        public Input<Inputs.KubernetesClusterKmsProviderArgs>? KmsProvider { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Kubernetes cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Kubernetes master configuration options. The structure is documented below.
        /// </summary>
        [Input("master", required: true)]
        public Input<Inputs.KubernetesClusterMasterArgs> Master { get; set; } = null!;

        /// <summary>
        /// Name of a specific Kubernetes cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the cluster network.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Network policy provider for the cluster. Possible values: `CALICO`.
        /// </summary>
        [Input("networkPolicyProvider")]
        public Input<string>? NetworkPolicyProvider { get; set; }

        /// <summary>
        /// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
        /// </summary>
        [Input("nodeIpv4CidrMaskSize")]
        public Input<int>? NodeIpv4CidrMaskSize { get; set; }

        /// <summary>
        /// Service account to be used by the worker nodes of the Kubernetes cluster
        /// to access Container Registry or to push node logs and metrics.
        /// </summary>
        [Input("nodeServiceAccountId", required: true)]
        public Input<string> NodeServiceAccountId { get; set; } = null!;

        /// <summary>
        /// Cluster release channel.
        /// </summary>
        [Input("releaseChannel")]
        public Input<string>? ReleaseChannel { get; set; }

        /// <summary>
        /// Service account to be used for provisioning Compute Cloud and VPC resources
        /// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
        /// cluster will be located and on the folder where selected network resides.
        /// </summary>
        [Input("serviceAccountId", required: true)]
        public Input<string> ServiceAccountId { get; set; } = null!;

        /// <summary>
        /// CIDR block. IP range Kubernetes service Kubernetes cluster
        /// IP addresses will be allocated from. It should not overlap with any subnet in the network
        /// the Kubernetes cluster located in.
        /// </summary>
        [Input("serviceIpv4Range")]
        public Input<string>? ServiceIpv4Range { get; set; }

        public KubernetesClusterArgs()
        {
        }
    }

    public sealed class KubernetesClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CIDR block. IP range for allocating pod addresses.
        /// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
        /// set up for this CIDR blocks in node subnets.
        /// </summary>
        [Input("clusterIpv4Range")]
        public Input<string>? ClusterIpv4Range { get; set; }

        /// <summary>
        /// (Computed) The Kubernetes cluster creation timestamp.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A description of the Kubernetes cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the folder that the Kubernetes cluster belongs to.
        /// If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// (Computed) Health of the Kubernetes cluster.
        /// </summary>
        [Input("health")]
        public Input<string>? Health { get; set; }

        /// <summary>
        /// cluster KMS provider parameters.
        /// </summary>
        [Input("kmsProvider")]
        public Input<Inputs.KubernetesClusterKmsProviderGetArgs>? KmsProvider { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Kubernetes cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Kubernetes master configuration options. The structure is documented below.
        /// </summary>
        [Input("master")]
        public Input<Inputs.KubernetesClusterMasterGetArgs>? Master { get; set; }

        /// <summary>
        /// Name of a specific Kubernetes cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the cluster network.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Network policy provider for the cluster. Possible values: `CALICO`.
        /// </summary>
        [Input("networkPolicyProvider")]
        public Input<string>? NetworkPolicyProvider { get; set; }

        /// <summary>
        /// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
        /// </summary>
        [Input("nodeIpv4CidrMaskSize")]
        public Input<int>? NodeIpv4CidrMaskSize { get; set; }

        /// <summary>
        /// Service account to be used by the worker nodes of the Kubernetes cluster
        /// to access Container Registry or to push node logs and metrics.
        /// </summary>
        [Input("nodeServiceAccountId")]
        public Input<string>? NodeServiceAccountId { get; set; }

        /// <summary>
        /// Cluster release channel.
        /// </summary>
        [Input("releaseChannel")]
        public Input<string>? ReleaseChannel { get; set; }

        /// <summary>
        /// Service account to be used for provisioning Compute Cloud and VPC resources
        /// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
        /// cluster will be located and on the folder where selected network resides.
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        /// <summary>
        /// CIDR block. IP range Kubernetes service Kubernetes cluster
        /// IP addresses will be allocated from. It should not overlap with any subnet in the network
        /// the Kubernetes cluster located in.
        /// </summary>
        [Input("serviceIpv4Range")]
        public Input<string>? ServiceIpv4Range { get; set; }

        /// <summary>
        /// (Computed)Status of the Kubernetes cluster.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public KubernetesClusterState()
        {
        }
    }
}
