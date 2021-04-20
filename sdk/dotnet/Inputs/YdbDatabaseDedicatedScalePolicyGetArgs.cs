// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class YdbDatabaseDedicatedScalePolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fixed scaling policy for the Yandex Database cluster.
        /// The structure is documented below.
        /// </summary>
        [Input("fixedScale", required: true)]
        public Input<Inputs.YdbDatabaseDedicatedScalePolicyFixedScaleGetArgs> FixedScale { get; set; } = null!;

        public YdbDatabaseDedicatedScalePolicyGetArgs()
        {
        }
    }
}
