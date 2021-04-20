// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceGroupInstanceTemplateSecondaryDiskGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This value can be used to reference the device under `/dev/disk/by-id/`.
        /// </summary>
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        /// <summary>
        /// Parameters for creating a disk alongside the instance. The structure is documented below.
        /// </summary>
        [Input("initializeParams", required: true)]
        public Input<Inputs.ComputeInstanceGroupInstanceTemplateSecondaryDiskInitializeParamsGetArgs> InitializeParams { get; set; } = null!;

        /// <summary>
        /// The access mode to the disk resource. By default a disk is attached in `READ_WRITE` mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public ComputeInstanceGroupInstanceTemplateSecondaryDiskGetArgs()
        {
        }
    }
}
