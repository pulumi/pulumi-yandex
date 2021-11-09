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
    public static class GetComputeSnapshot
    {
        /// <summary>
        /// Get information about a Yandex Compute snapshot. For more information, see
        /// [the official documentation](https://cloud.yandex.com/docs/compute/concepts/snapshot).
        /// </summary>
        public static Task<GetComputeSnapshotResult> InvokeAsync(GetComputeSnapshotArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetComputeSnapshotResult>("yandex:index/getComputeSnapshot:getComputeSnapshot", args ?? new GetComputeSnapshotArgs(), options.WithVersion());

        /// <summary>
        /// Get information about a Yandex Compute snapshot. For more information, see
        /// [the official documentation](https://cloud.yandex.com/docs/compute/concepts/snapshot).
        /// </summary>
        public static Output<GetComputeSnapshotResult> Invoke(GetComputeSnapshotInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetComputeSnapshotResult>("yandex:index/getComputeSnapshot:getComputeSnapshot", args ?? new GetComputeSnapshotInvokeArgs(), options.WithVersion());
    }


    public sealed class GetComputeSnapshotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the folder that the snapshot belongs to.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of a specific snapshot.
        /// </summary>
        [Input("snapshotId")]
        public string? SnapshotId { get; set; }

        public GetComputeSnapshotArgs()
        {
        }
    }

    public sealed class GetComputeSnapshotInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the folder that the snapshot belongs to.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of a specific snapshot.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public GetComputeSnapshotInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetComputeSnapshotResult
    {
        /// <summary>
        /// Snapshot creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// An optional description of this snapshot.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Minimum required size of the disk which is created from this snapshot.
        /// </summary>
        public readonly int DiskSize;
        /// <summary>
        /// ID of the folder that the snapshot belongs to.
        /// </summary>
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A map of labels applied to this snapshot.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        /// <summary>
        /// License IDs that indicate which licenses are attached to this snapshot.
        /// </summary>
        public readonly ImmutableArray<string> ProductIds;
        public readonly string SnapshotId;
        /// <summary>
        /// ID of the source disk.
        /// </summary>
        public readonly string SourceDiskId;
        /// <summary>
        /// The status of the snapshot.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The size of the snapshot, specified in Gb.
        /// </summary>
        public readonly int StorageSize;

        [OutputConstructor]
        private GetComputeSnapshotResult(
            string createdAt,

            string description,

            int diskSize,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<string> productIds,

            string snapshotId,

            string sourceDiskId,

            string status,

            int storageSize)
        {
            CreatedAt = createdAt;
            Description = description;
            DiskSize = diskSize;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Name = name;
            ProductIds = productIds;
            SnapshotId = snapshotId;
            SourceDiskId = sourceDiskId;
            Status = status;
            StorageSize = storageSize;
        }
    }
}
