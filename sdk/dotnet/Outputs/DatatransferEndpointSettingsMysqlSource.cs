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
    public sealed class DatatransferEndpointSettingsMysqlSource
    {
        /// <summary>
        /// Connection settings. The structure is documented below.
        /// </summary>
        public readonly Outputs.DatatransferEndpointSettingsMysqlSourceConnection? Connection;
        /// <summary>
        /// Name of the database to transfer.
        /// </summary>
        public readonly string? Database;
        /// <summary>
        /// Opposite of `include_table_regex`. The tables matching the specified regular expressions will not be transferred.
        /// </summary>
        public readonly ImmutableArray<string> ExcludeTablesRegexes;
        /// <summary>
        /// List of regular expressions of table names which should be transferred. A table name is formatted as schemaname.tablename. For example, a single regular expression may look like `^mydb.employees$`.
        /// </summary>
        public readonly ImmutableArray<string> IncludeTablesRegexes;
        /// <summary>
        /// Defines which database schema objects should be transferred, e.g. views, functions, etc.
        /// </summary>
        public readonly Outputs.DatatransferEndpointSettingsMysqlSourceObjectTransferSettings? ObjectTransferSettings;
        /// <summary>
        /// Password for the database access. This is a block with a single field named `raw` which should contain the password.
        /// </summary>
        public readonly Outputs.DatatransferEndpointSettingsMysqlSourcePassword? Password;
        /// <summary>
        /// Timezone to use for parsing timestamps for saving source timezones. Accepts values from IANA timezone database. Default: local timezone.
        /// </summary>
        public readonly string? Timezone;
        /// <summary>
        /// User for the database access.
        /// </summary>
        public readonly string? User;

        [OutputConstructor]
        private DatatransferEndpointSettingsMysqlSource(
            Outputs.DatatransferEndpointSettingsMysqlSourceConnection? connection,

            string? database,

            ImmutableArray<string> excludeTablesRegexes,

            ImmutableArray<string> includeTablesRegexes,

            Outputs.DatatransferEndpointSettingsMysqlSourceObjectTransferSettings? objectTransferSettings,

            Outputs.DatatransferEndpointSettingsMysqlSourcePassword? password,

            string? timezone,

            string? user)
        {
            Connection = connection;
            Database = database;
            ExcludeTablesRegexes = excludeTablesRegexes;
            IncludeTablesRegexes = includeTablesRegexes;
            ObjectTransferSettings = objectTransferSettings;
            Password = password;
            Timezone = timezone;
            User = user;
        }
    }
}
