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
    public sealed class GetComputeInstanceGroupScalePolicyFixedScaleResult
    {
        /// <summary>
        /// The size of the disk in GB.
        /// </summary>
        public readonly int Size;

        [OutputConstructor]
        private GetComputeInstanceGroupScalePolicyFixedScaleResult(int size)
        {
            Size = size;
        }
    }
}
