// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetComputeImage
    {
        public static Task<GetComputeImageResult> InvokeAsync(GetComputeImageArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetComputeImageResult>("yandex:index/getComputeImage:getComputeImage", args ?? new GetComputeImageArgs(), options.WithVersion());
    }


    public sealed class GetComputeImageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The family name of an image. Used to search the latest image in a family.
        /// </summary>
        [Input("family")]
        public string? Family { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// The ID of a specific image.
        /// </summary>
        [Input("imageId")]
        public string? ImageId { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetComputeImageArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetComputeImageResult
    {
        /// <summary>
        /// Image creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// An optional description of this image.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The OS family name of the image.
        /// </summary>
        public readonly string Family;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ImageId;
        /// <summary>
        /// A map of labels applied to this image.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Minimum size of the disk which is created from this image.
        /// </summary>
        public readonly int MinDiskSize;
        public readonly string Name;
        /// <summary>
        /// Operating system type that the image contains.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// License IDs that indicate which licenses are attached to this image.
        /// </summary>
        public readonly ImmutableArray<string> ProductIds;
        /// <summary>
        /// The size of the image, specified in Gb.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The status of the image.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetComputeImageResult(
            string createdAt,

            string description,

            string family,

            string folderId,

            string id,

            string imageId,

            ImmutableDictionary<string, string> labels,

            int minDiskSize,

            string name,

            string osType,

            ImmutableArray<string> productIds,

            int size,

            string status)
        {
            CreatedAt = createdAt;
            Description = description;
            Family = family;
            FolderId = folderId;
            Id = id;
            ImageId = imageId;
            Labels = labels;
            MinDiskSize = minDiskSize;
            Name = name;
            OsType = osType;
            ProductIds = productIds;
            Size = size;
            Status = status;
        }
    }
}
