// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsPostgresSourceGetArgs : Pulumi.ResourceArgs
    {
        [Input("collapseInheritTable")]
        public Input<bool>? CollapseInheritTable { get; set; }

        /// <summary>
        /// Connection settings. The structure is documented below.
        /// </summary>
        [Input("connection")]
        public Input<Inputs.DatatransferEndpointSettingsPostgresSourceConnectionGetArgs>? Connection { get; set; }

        /// <summary>
        /// Name of the database to transfer.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        [Input("excludeTables")]
        private InputList<string>? _excludeTables;

        /// <summary>
        /// List of tables which will not be transfered, formatted as `schemaname.tablename`.
        /// </summary>
        public InputList<string> ExcludeTables
        {
            get => _excludeTables ?? (_excludeTables = new InputList<string>());
            set => _excludeTables = value;
        }

        [Input("includeTables")]
        private InputList<string>? _includeTables;

        /// <summary>
        /// List of tables to transfer, formatted as `schemaname.tablename`. If omitted or an empty list is specified, all tables will be transferred.
        /// </summary>
        public InputList<string> IncludeTables
        {
            get => _includeTables ?? (_includeTables = new InputList<string>());
            set => _includeTables = value;
        }

        /// <summary>
        /// Defines which database schema objects should be transferred, e.g. views, functions, etc.
        /// </summary>
        [Input("objectTransferSettings")]
        public Input<Inputs.DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsGetArgs>? ObjectTransferSettings { get; set; }

        /// <summary>
        /// Password for the database access. This is a block with a single field named `raw` which should contain the password.
        /// </summary>
        [Input("password")]
        public Input<Inputs.DatatransferEndpointSettingsPostgresSourcePasswordGetArgs>? Password { get; set; }

        /// <summary>
        /// Name of the database schema in which auxiliary tables needed for the transfer will be created. Empty `service_schema` implies schema "public".
        /// </summary>
        [Input("serviceSchema")]
        public Input<string>? ServiceSchema { get; set; }

        /// <summary>
        /// Maximum WAL size held by the replication slot, in gigabytes. Exceeding this limit will result in a replication failure and deletion of the replication slot. Unlimited by default.
        /// </summary>
        [Input("slotGigabyteLagLimit")]
        public Input<int>? SlotGigabyteLagLimit { get; set; }

        /// <summary>
        /// User for the database access.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public DatatransferEndpointSettingsPostgresSourceGetArgs()
        {
        }
    }
}