// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class KubernetesNodeGroupMaintenancePolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean flag that specifies if node group can be repaired automatically. When omitted, default value is TRUE.
        /// </summary>
        [Input("autoRepair", required: true)]
        public Input<bool> AutoRepair { get; set; } = null!;

        /// <summary>
        /// Boolean flag that specifies if node group can be upgraded automatically. When omitted, default value is TRUE.
        /// </summary>
        [Input("autoUpgrade", required: true)]
        public Input<bool> AutoUpgrade { get; set; } = null!;

        [Input("maintenanceWindows")]
        private InputList<Inputs.KubernetesNodeGroupMaintenancePolicyMaintenanceWindowGetArgs>? _maintenanceWindows;

        /// <summary>
        /// (Computed) Set of day intervals, when maintenance is allowed for this node group. When omitted, it defaults to any time.
        /// </summary>
        public InputList<Inputs.KubernetesNodeGroupMaintenancePolicyMaintenanceWindowGetArgs> MaintenanceWindows
        {
            get => _maintenanceWindows ?? (_maintenanceWindows = new InputList<Inputs.KubernetesNodeGroupMaintenancePolicyMaintenanceWindowGetArgs>());
            set => _maintenanceWindows = value;
        }

        public KubernetesNodeGroupMaintenancePolicyGetArgs()
        {
        }
    }
}
