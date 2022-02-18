// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings specific to the MySQL source endpoint.
        /// </summary>
        [Input("mysqlSource")]
        public Input<Inputs.DatatransferEndpointSettingsMysqlSourceArgs>? MysqlSource { get; set; }

        /// <summary>
        /// Settings specific to the MySQL target endpoint.
        /// </summary>
        [Input("mysqlTarget")]
        public Input<Inputs.DatatransferEndpointSettingsMysqlTargetArgs>? MysqlTarget { get; set; }

        /// <summary>
        /// Settings specific to the PostgreSQL source endpoint.
        /// </summary>
        [Input("postgresSource")]
        public Input<Inputs.DatatransferEndpointSettingsPostgresSourceArgs>? PostgresSource { get; set; }

        /// <summary>
        /// Settings specific to the PostgreSQL target endpoint.
        /// </summary>
        [Input("postgresTarget")]
        public Input<Inputs.DatatransferEndpointSettingsPostgresTargetArgs>? PostgresTarget { get; set; }

        public DatatransferEndpointSettingsArgs()
        {
        }
    }
}