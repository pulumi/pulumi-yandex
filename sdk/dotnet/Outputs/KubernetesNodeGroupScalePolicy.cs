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
    public sealed class KubernetesNodeGroupScalePolicy
    {
        /// <summary>
        /// Scale policy for an autoscaled node group. The structure is documented below.
        /// </summary>
        public readonly Outputs.KubernetesNodeGroupScalePolicyAutoScale? AutoScale;
        /// <summary>
        /// Scale policy for a fixed scale node group. The structure is documented below.
        /// </summary>
        public readonly Outputs.KubernetesNodeGroupScalePolicyFixedScale? FixedScale;

        [OutputConstructor]
        private KubernetesNodeGroupScalePolicy(
            Outputs.KubernetesNodeGroupScalePolicyAutoScale? autoScale,

            Outputs.KubernetesNodeGroupScalePolicyFixedScale? fixedScale)
        {
            AutoScale = autoScale;
            FixedScale = fixedScale;
        }
    }
}
