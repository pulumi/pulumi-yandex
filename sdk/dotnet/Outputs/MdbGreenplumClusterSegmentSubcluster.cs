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
    public sealed class MdbGreenplumClusterSegmentSubcluster
    {
        /// <summary>
        /// Resources allocated to hosts for segment subcluster of the Greenplum cluster. The structure is documented below.
        /// </summary>
        public readonly Outputs.MdbGreenplumClusterSegmentSubclusterResources Resources;

        [OutputConstructor]
        private MdbGreenplumClusterSegmentSubcluster(Outputs.MdbGreenplumClusterSegmentSubclusterResources resources)
        {
            Resources = resources;
        }
    }
}
