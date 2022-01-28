// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class KubernetesNodeGroupInstanceTemplateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The specifications for boot disks that will be attached to the instance. The structure is documented below.
        /// </summary>
        [Input("bootDisk")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplateBootDiskGetArgs>? BootDisk { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// The set of metadata `key:value` pairs assigned to this instance template. This includes custom metadata and predefined keys.
        /// * `resources.0.memory` - The memory size allocated to the instance.
        /// * `resources.0.cores` - Number of CPU cores allocated to the instance.
        /// * `resources.0.core_fraction` - Baseline core performance as a percent.
        /// * `resources.0.gpus` - Number of GPU cores allocated to the instance.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// A public address that can be used to access the internet over NAT.
        /// </summary>
        [Input("nat")]
        public Input<bool>? Nat { get; set; }

        /// <summary>
        /// Type of network acceleration. Values: `standard`, `software_accelerated`.
        /// </summary>
        [Input("networkAccelerationType")]
        public Input<string>? NetworkAccelerationType { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceGetArgs>? _networkInterfaces;

        /// <summary>
        /// An array with the network interfaces that will be attached to the instance. The structure is documented below.
        /// </summary>
        public InputList<Inputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// The placement policy configuration. The structure is documented below.
        /// </summary>
        [Input("placementPolicy")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplatePlacementPolicyGetArgs>? PlacementPolicy { get; set; }

        /// <summary>
        /// The ID of the hardware platform configuration for the node group compute instances.
        /// </summary>
        [Input("platformId")]
        public Input<string>? PlatformId { get; set; }

        [Input("resources")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplateResourcesGetArgs>? Resources { get; set; }

        /// <summary>
        /// The scheduling policy for the instances in node group. The structure is documented below.
        /// </summary>
        [Input("schedulingPolicy")]
        public Input<Inputs.KubernetesNodeGroupInstanceTemplateSchedulingPolicyGetArgs>? SchedulingPolicy { get; set; }

        public KubernetesNodeGroupInstanceTemplateGetArgs()
        {
        }
    }
}
