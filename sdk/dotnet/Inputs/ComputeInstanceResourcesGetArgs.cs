// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceResourcesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If provided, specifies baseline performance for a core as a percent.
        /// </summary>
        [Input("coreFraction")]
        public Input<int>? CoreFraction { get; set; }

        /// <summary>
        /// CPU cores for the instance.
        /// </summary>
        [Input("cores", required: true)]
        public Input<int> Cores { get; set; } = null!;

        [Input("gpus")]
        public Input<int>? Gpus { get; set; }

        /// <summary>
        /// Memory size in GB.
        /// </summary>
        [Input("memory", required: true)]
        public Input<double> Memory { get; set; } = null!;

        public ComputeInstanceResourcesGetArgs()
        {
        }
    }
}
