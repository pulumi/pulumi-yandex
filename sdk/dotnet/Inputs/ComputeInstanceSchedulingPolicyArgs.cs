// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceSchedulingPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies if the instance is preemptible. Defaults to false.
        /// </summary>
        [Input("preemptible")]
        public Input<bool>? Preemptible { get; set; }

        public ComputeInstanceSchedulingPolicyArgs()
        {
        }
    }
}
