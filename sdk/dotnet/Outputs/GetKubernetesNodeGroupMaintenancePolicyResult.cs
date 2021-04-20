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
    public sealed class GetKubernetesNodeGroupMaintenancePolicyResult
    {
        /// <summary>
        /// Boolean flag.
        /// </summary>
        public readonly bool AutoRepair;
        /// <summary>
        /// Boolean flag.
        /// </summary>
        public readonly bool AutoUpgrade;
        /// <summary>
        /// Set of day intervals, when maintenance is allowed for this node group.
        /// When omitted, it defaults to any time.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupMaintenancePolicyMaintenanceWindowResult> MaintenanceWindows;

        [OutputConstructor]
        private GetKubernetesNodeGroupMaintenancePolicyResult(
            bool autoRepair,

            bool autoUpgrade,

            ImmutableArray<Outputs.GetKubernetesNodeGroupMaintenancePolicyMaintenanceWindowResult> maintenanceWindows)
        {
            AutoRepair = autoRepair;
            AutoUpgrade = autoUpgrade;
            MaintenanceWindows = maintenanceWindows;
        }
    }
}
