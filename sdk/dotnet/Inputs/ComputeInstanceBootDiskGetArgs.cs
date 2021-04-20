// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceBootDiskGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the disk is auto-deleted when the instance
        /// is deleted. The default value is false.
        /// </summary>
        [Input("autoDelete")]
        public Input<bool>? AutoDelete { get; set; }

        /// <summary>
        /// Name that can be used to access an attached disk
        /// under `/dev/disk/by-id/`.
        /// </summary>
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        /// <summary>
        /// ID of the disk that is attached to the instance.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// Parameters for a new disk that will be created
        /// alongside the new instance. Either `initialize_params` or `disk_id` must be set. The structure is documented below.
        /// </summary>
        [Input("initializeParams")]
        public Input<Inputs.ComputeInstanceBootDiskInitializeParamsGetArgs>? InitializeParams { get; set; }

        /// <summary>
        /// Type of access to the disk resource. By default, a disk is attached in `READ_WRITE` mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public ComputeInstanceBootDiskGetArgs()
        {
        }
    }
}
