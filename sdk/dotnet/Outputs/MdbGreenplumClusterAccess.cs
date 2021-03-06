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
    public sealed class MdbGreenplumClusterAccess
    {
        /// <summary>
        /// Allow access for [Yandex DataLens](https://cloud.yandex.com/services/datalens).
        /// </summary>
        public readonly bool? DataLens;
        /// <summary>
        /// Allows access for SQL queries in the management console
        /// </summary>
        public readonly bool? WebSql;

        [OutputConstructor]
        private MdbGreenplumClusterAccess(
            bool? dataLens,

            bool? webSql)
        {
            DataLens = dataLens;
            WebSql = webSql;
        }
    }
}
