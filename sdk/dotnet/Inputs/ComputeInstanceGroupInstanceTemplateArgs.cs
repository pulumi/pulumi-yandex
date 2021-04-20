// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceGroupInstanceTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boot disk specifications for the instance. The structure is documented below.
        /// </summary>
        [Input("bootDisk", required: true)]
        public Input<Inputs.ComputeInstanceGroupInstanceTemplateBootDiskArgs> BootDisk { get; set; } = null!;

        /// <summary>
        /// A description of the boot disk.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Hostname template for the instance.   
        /// This field is used to generate the FQDN value of instance.
        /// The hostname must be unique within the network and region.
        /// If not specified, the hostname will be equal to id of the instance
        /// and FQDN will be `&lt;id&gt;.auto.internal`. Otherwise FQDN will be `&lt;hostname&gt;.&lt;region_id&gt;.internal`.
        /// In order to be unique it must contain at least on of instance unique placeholders:
        /// {instance.short_id}
        /// {instance.index}
        /// combination of {instance.zone_id} and {instance.index_in_zone}
        /// Example: my-instance-{instance.index}
        /// If not set, `name` value will be used
        /// It may also contain another placeholders, see metadata doc for full list.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A map of labels of metric.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A set of metadata key/value pairs to make available from within the instance.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Name template of the instance.  
        /// In order to be unique it must contain at least one of instance unique placeholders:
        /// {instance.short_id}
        /// {instance.index}
        /// combination of {instance.zone_id} and {instance.index_in_zone}
        /// Example: my-instance-{instance.index}
        /// If not set, default is used: {instance_group.id}-{instance.short_id}
        /// It may also contain another placeholders, see metadata doc for full list.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces", required: true)]
        private InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// Network specifications for the instance. This can be used multiple times for adding multiple interfaces. The structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("networkSettings")]
        private InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkSettingArgs>? _networkSettings;

        /// <summary>
        /// Network acceleration type for instance. The structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkSettingArgs> NetworkSettings
        {
            get => _networkSettings ?? (_networkSettings = new InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkSettingArgs>());
            set => _networkSettings = value;
        }

        /// <summary>
        /// The placement policy configuration. The structure is documented below.
        /// </summary>
        [Input("placementPolicy")]
        public Input<Inputs.ComputeInstanceGroupInstanceTemplatePlacementPolicyArgs>? PlacementPolicy { get; set; }

        /// <summary>
        /// The ID of the hardware platform configuration for the instance. The default is 'standard-v1'.
        /// </summary>
        [Input("platformId")]
        public Input<string>? PlatformId { get; set; }

        /// <summary>
        /// Compute resource specifications for the instance. The structure is documented below.
        /// </summary>
        [Input("resources", required: true)]
        public Input<Inputs.ComputeInstanceGroupInstanceTemplateResourcesArgs> Resources { get; set; } = null!;

        /// <summary>
        /// The scheduling policy configuration. The structure is documented below.
        /// </summary>
        [Input("schedulingPolicy")]
        public Input<Inputs.ComputeInstanceGroupInstanceTemplateSchedulingPolicyArgs>? SchedulingPolicy { get; set; }

        [Input("secondaryDisks")]
        private InputList<Inputs.ComputeInstanceGroupInstanceTemplateSecondaryDiskArgs>? _secondaryDisks;

        /// <summary>
        /// A list of disks to attach to the instance. The structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceGroupInstanceTemplateSecondaryDiskArgs> SecondaryDisks
        {
            get => _secondaryDisks ?? (_secondaryDisks = new InputList<Inputs.ComputeInstanceGroupInstanceTemplateSecondaryDiskArgs>());
            set => _secondaryDisks = value;
        }

        /// <summary>
        /// The ID of the service account authorized for this instance.
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        public ComputeInstanceGroupInstanceTemplateArgs()
        {
        }
    }
}
