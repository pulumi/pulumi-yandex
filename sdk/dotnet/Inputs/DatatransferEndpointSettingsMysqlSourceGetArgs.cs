// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsMysqlSourceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection settings. The structure is documented below.
        /// </summary>
        [Input("connection")]
        public Input<Inputs.DatatransferEndpointSettingsMysqlSourceConnectionGetArgs>? Connection { get; set; }

        /// <summary>
        /// Name of the database to transfer.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        [Input("excludeTablesRegexes")]
        private InputList<string>? _excludeTablesRegexes;

        /// <summary>
        /// Opposite of `include_table_regex`. The tables matching the specified regular expressions will not be transferred.
        /// </summary>
        public InputList<string> ExcludeTablesRegexes
        {
            get => _excludeTablesRegexes ?? (_excludeTablesRegexes = new InputList<string>());
            set => _excludeTablesRegexes = value;
        }

        [Input("includeTablesRegexes")]
        private InputList<string>? _includeTablesRegexes;

        /// <summary>
        /// List of regular expressions of table names which should be transferred. A table name is formatted as schemaname.tablename. For example, a single regular expression may look like `^mydb.employees$`.
        /// </summary>
        public InputList<string> IncludeTablesRegexes
        {
            get => _includeTablesRegexes ?? (_includeTablesRegexes = new InputList<string>());
            set => _includeTablesRegexes = value;
        }

        /// <summary>
        /// Defines which database schema objects should be transferred, e.g. views, functions, etc.
        /// </summary>
        [Input("objectTransferSettings")]
        public Input<Inputs.DatatransferEndpointSettingsMysqlSourceObjectTransferSettingsGetArgs>? ObjectTransferSettings { get; set; }

        /// <summary>
        /// Password for the database access. This is a block with a single field named `raw` which should contain the password.
        /// </summary>
        [Input("password")]
        public Input<Inputs.DatatransferEndpointSettingsMysqlSourcePasswordGetArgs>? Password { get; set; }

        /// <summary>
        /// Timezone to use for parsing timestamps for saving source timezones. Accepts values from IANA timezone database. Default: local timezone.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// User for the database access.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public DatatransferEndpointSettingsMysqlSourceGetArgs()
        {
        }
    }
}