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
    public sealed class GetYdbDatabaseDedicatedStorageConfigResult
    {
        /// <summary>
        /// Amount of storage groups of selected type in the Yandex Database cluster.
        /// </summary>
        public readonly int GroupCount;
        /// <summary>
        /// Storage type ID of the Yandex Database cluster.
        /// </summary>
        public readonly string StorageTypeId;

        [OutputConstructor]
        private GetYdbDatabaseDedicatedStorageConfigResult(
            int groupCount,

            string storageTypeId)
        {
            GroupCount = groupCount;
            StorageTypeId = storageTypeId;
        }
    }
}
