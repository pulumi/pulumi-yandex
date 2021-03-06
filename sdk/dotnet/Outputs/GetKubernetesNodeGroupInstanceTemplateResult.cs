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
    public sealed class GetKubernetesNodeGroupInstanceTemplateResult
    {
        /// <summary>
        /// The specifications for boot disks that will be attached to the instance. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateBootDiskResult> BootDisks;
        /// <summary>
        /// Container runtime configuration. The structure is documented below.
        /// ---
        /// </summary>
        public readonly Outputs.GetKubernetesNodeGroupInstanceTemplateContainerRuntimeResult ContainerRuntime;
        /// <summary>
        /// The set of metadata `key:value` pairs assigned to this instance template. This includes custom metadata and predefined keys.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// A public address that can be used to access the internet over NAT.
        /// </summary>
        public readonly bool Nat;
        /// <summary>
        /// Type of network acceleration. Values: `standard`, `software_accelerated`.
        /// </summary>
        public readonly string NetworkAccelerationType;
        /// <summary>
        /// An array with the network interfaces that will be attached to the instance. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateNetworkInterfaceResult> NetworkInterfaces;
        /// <summary>
        /// (Optional) The placement policy configuration. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplatePlacementPolicyResult> PlacementPolicies;
        /// <summary>
        /// The ID of the hardware platform configuration for the instance.
        /// </summary>
        public readonly string PlatformId;
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateResourceResult> Resources;
        /// <summary>
        /// The scheduling policy for the instances in node group. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateSchedulingPolicyResult> SchedulingPolicies;

        [OutputConstructor]
        private GetKubernetesNodeGroupInstanceTemplateResult(
            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateBootDiskResult> bootDisks,

            Outputs.GetKubernetesNodeGroupInstanceTemplateContainerRuntimeResult containerRuntime,

            ImmutableDictionary<string, string> metadata,

            bool nat,

            string networkAccelerationType,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateNetworkInterfaceResult> networkInterfaces,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplatePlacementPolicyResult> placementPolicies,

            string platformId,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateResourceResult> resources,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateSchedulingPolicyResult> schedulingPolicies)
        {
            BootDisks = bootDisks;
            ContainerRuntime = containerRuntime;
            Metadata = metadata;
            Nat = nat;
            NetworkAccelerationType = networkAccelerationType;
            NetworkInterfaces = networkInterfaces;
            PlacementPolicies = placementPolicies;
            PlatformId = platformId;
            Resources = resources;
            SchedulingPolicies = schedulingPolicies;
        }
    }
}
