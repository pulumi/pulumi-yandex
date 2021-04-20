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
    public sealed class GetDataprocClusterClusterConfigResult
    {
        /// <summary>
        /// Data Proc specific options. The structure is documented below.
        /// </summary>
        public readonly Outputs.GetDataprocClusterClusterConfigHadoopResult Hadoop;
        /// <summary>
        /// Configuration of the Data Proc subcluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDataprocClusterClusterConfigSubclusterSpecResult> SubclusterSpecs;
        /// <summary>
        /// Version of Data Proc image.
        /// </summary>
        public readonly string VersionId;

        [OutputConstructor]
        private GetDataprocClusterClusterConfigResult(
            Outputs.GetDataprocClusterClusterConfigHadoopResult hadoop,

            ImmutableArray<Outputs.GetDataprocClusterClusterConfigSubclusterSpecResult> subclusterSpecs,

            string versionId)
        {
            Hadoop = hadoop;
            SubclusterSpecs = subclusterSpecs;
            VersionId = versionId;
        }
    }
}
