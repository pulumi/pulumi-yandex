// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbMysqlClusterAccessArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Allow access for [Yandex DataLens](https://cloud.yandex.com/services/datalens).
        /// </summary>
        [Input("dataLens", required: true)]
        public bool DataLens { get; set; }

        /// <summary>
        /// Allows access for [SQL queries in the management console](https://cloud.yandex.ru/docs/managed-mysql/operations/web-sql-query).
        /// </summary>
        [Input("webSql", required: true)]
        public bool WebSql { get; set; }

        public GetMdbMysqlClusterAccessArgs()
        {
        }
    }
}
