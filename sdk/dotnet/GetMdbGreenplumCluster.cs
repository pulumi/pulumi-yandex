// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetMdbGreenplumCluster
    {
        public static Task<GetMdbGreenplumClusterResult> InvokeAsync(GetMdbGreenplumClusterArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMdbGreenplumClusterResult>("yandex:index/getMdbGreenplumCluster:getMdbGreenplumCluster", args ?? new GetMdbGreenplumClusterArgs(), options.WithDefaults());

        public static Output<GetMdbGreenplumClusterResult> Invoke(GetMdbGreenplumClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMdbGreenplumClusterResult>("yandex:index/getMdbGreenplumCluster:getMdbGreenplumCluster", args ?? new GetMdbGreenplumClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMdbGreenplumClusterArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetMdbGreenplumClusterArgs()
        {
        }
    }

    public sealed class GetMdbGreenplumClusterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetMdbGreenplumClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMdbGreenplumClusterResult
    {
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterAccessResult> Accesses;
        public readonly bool AssignPublicIp;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterBackupWindowStartResult> BackupWindowStarts;
        public readonly string ClusterId;
        public readonly string CreatedAt;
        public readonly bool DeletionProtection;
        public readonly string Description;
        public readonly string Environment;
        public readonly string FolderId;
        public readonly string Health;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly int MasterHostCount;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterMasterHostResult> MasterHosts;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterMasterSubclusterResult> MasterSubclusters;
        public readonly string Name;
        public readonly string NetworkId;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly int SegmentHostCount;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterSegmentHostResult> SegmentHosts;
        public readonly int SegmentInHost;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterSegmentSubclusterResult> SegmentSubclusters;
        public readonly string Status;
        public readonly string SubnetId;
        public readonly string UserName;
        public readonly string Version;
        public readonly string Zone;

        [OutputConstructor]
        private GetMdbGreenplumClusterResult(
            ImmutableArray<Outputs.GetMdbGreenplumClusterAccessResult> accesses,

            bool assignPublicIp,

            ImmutableArray<Outputs.GetMdbGreenplumClusterBackupWindowStartResult> backupWindowStarts,

            string clusterId,

            string createdAt,

            bool deletionProtection,

            string description,

            string environment,

            string folderId,

            string health,

            string id,

            ImmutableDictionary<string, string> labels,

            int masterHostCount,

            ImmutableArray<Outputs.GetMdbGreenplumClusterMasterHostResult> masterHosts,

            ImmutableArray<Outputs.GetMdbGreenplumClusterMasterSubclusterResult> masterSubclusters,

            string name,

            string networkId,

            ImmutableArray<string> securityGroupIds,

            int segmentHostCount,

            ImmutableArray<Outputs.GetMdbGreenplumClusterSegmentHostResult> segmentHosts,

            int segmentInHost,

            ImmutableArray<Outputs.GetMdbGreenplumClusterSegmentSubclusterResult> segmentSubclusters,

            string status,

            string subnetId,

            string userName,

            string version,

            string zone)
        {
            Accesses = accesses;
            AssignPublicIp = assignPublicIp;
            BackupWindowStarts = backupWindowStarts;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            DeletionProtection = deletionProtection;
            Description = description;
            Environment = environment;
            FolderId = folderId;
            Health = health;
            Id = id;
            Labels = labels;
            MasterHostCount = masterHostCount;
            MasterHosts = masterHosts;
            MasterSubclusters = masterSubclusters;
            Name = name;
            NetworkId = networkId;
            SecurityGroupIds = securityGroupIds;
            SegmentHostCount = segmentHostCount;
            SegmentHosts = segmentHosts;
            SegmentInHost = segmentInHost;
            SegmentSubclusters = segmentSubclusters;
            Status = status;
            SubnetId = subnetId;
            UserName = userName;
            Version = version;
            Zone = zone;
        }
    }
}
