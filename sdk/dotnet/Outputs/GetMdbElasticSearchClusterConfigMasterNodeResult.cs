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
    public sealed class GetMdbElasticSearchClusterConfigMasterNodeResult
    {
        /// <summary>
        /// Resources allocated to hosts of the Elasticsearch master nodes subcluster. The structure is documented below.
        /// </summary>
        public readonly Outputs.GetMdbElasticSearchClusterConfigMasterNodeResourcesResult Resources;

        [OutputConstructor]
        private GetMdbElasticSearchClusterConfigMasterNodeResult(Outputs.GetMdbElasticSearchClusterConfigMasterNodeResourcesResult resources)
        {
            Resources = resources;
        }
    }
}