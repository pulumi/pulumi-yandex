// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceGroupInstanceTemplateSecondaryDiskInitializeParamsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the boot disk.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The disk image to initialize this disk from.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The number of instances in the instance group.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The snapshot to initialize this disk from.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// Network acceleration type. By default a network is in `STANDARD` mode.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ComputeInstanceGroupInstanceTemplateSecondaryDiskInitializeParamsGetArgs()
        {
        }
    }
}
