// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Outputs
{

    [OutputType]
    public sealed class GetComputeInstanceGroupInstanceTemplateResult
    {
        /// <summary>
        /// The specifications for boot disk that will be attached to the instance. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateBootDiskResult> BootDisks;
        /// <summary>
        /// A description of the boot disk.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Hostname temaplate for the instance.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// A map of labels applied to this instance.
        /// * `resources.0.memory` - The memory size allocated to the instance.
        /// * `resources.0.cores` - Number of CPU cores allocated to the instance.
        /// * `resources.0.core_fraction` - Baseline core performance as a percent.
        /// * `resources.0.gpus` - Number of GPU cores allocated to the instance.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The set of metadata `key:value` pairs assigned to this instance template. This includes custom metadata and predefined keys.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// The name of the managed instance.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An array with the network interfaces attached to the managed instance. The structure is documented below.
        /// * `status_changed_at` -The timestamp when the status of the managed instance was last changed.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateNetworkInterfaceResult> NetworkInterfaces;
        /// <summary>
        /// Network acceleration settings. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateNetworkSettingResult> NetworkSettings;
        public readonly Outputs.GetComputeInstanceGroupInstanceTemplatePlacementPolicyResult? PlacementPolicy;
        /// <summary>
        /// The ID of the hardware platform configuration for the instance.
        /// </summary>
        public readonly string PlatformId;
        public readonly ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateResourceResult> Resources;
        /// <summary>
        /// The scheduling policy for the instance. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateSchedulingPolicyResult> SchedulingPolicies;
        /// <summary>
        /// An array with the secondary disks that will be attached to the instance. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateSecondaryDiskResult> SecondaryDisks;
        /// <summary>
        /// The service account ID for the instance.
        /// </summary>
        public readonly string ServiceAccountId;

        [OutputConstructor]
        private GetComputeInstanceGroupInstanceTemplateResult(
            ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateBootDiskResult> bootDisks,

            string description,

            string hostname,

            ImmutableDictionary<string, string> labels,

            ImmutableDictionary<string, string> metadata,

            string name,

            ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateNetworkInterfaceResult> networkInterfaces,

            ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateNetworkSettingResult> networkSettings,

            Outputs.GetComputeInstanceGroupInstanceTemplatePlacementPolicyResult? placementPolicy,

            string platformId,

            ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateResourceResult> resources,

            ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateSchedulingPolicyResult> schedulingPolicies,

            ImmutableArray<Outputs.GetComputeInstanceGroupInstanceTemplateSecondaryDiskResult> secondaryDisks,

            string serviceAccountId)
        {
            BootDisks = bootDisks;
            Description = description;
            Hostname = hostname;
            Labels = labels;
            Metadata = metadata;
            Name = name;
            NetworkInterfaces = networkInterfaces;
            NetworkSettings = networkSettings;
            PlacementPolicy = placementPolicy;
            PlatformId = platformId;
            Resources = resources;
            SchedulingPolicies = schedulingPolicies;
            SecondaryDisks = secondaryDisks;
            ServiceAccountId = serviceAccountId;
        }
    }
}
