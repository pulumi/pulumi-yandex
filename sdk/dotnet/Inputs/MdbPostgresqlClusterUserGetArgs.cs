// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbPostgresqlClusterUserGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of connections per user. (Default 50)
        /// </summary>
        [Input("connLimit")]
        public Input<int>? ConnLimit { get; set; }

        [Input("grants")]
        private InputList<string>? _grants;

        /// <summary>
        /// List of the user's grants.
        /// </summary>
        public InputList<string> Grants
        {
            get => _grants ?? (_grants = new InputList<string>());
            set => _grants = value;
        }

        /// <summary>
        /// User's ability to login.
        /// </summary>
        [Input("login")]
        public Input<bool>? Login { get; set; }

        /// <summary>
        /// Host state name. Is should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please see `replication_source_name` parameter.
        /// Also, this field is used to select which host will be selected as a master host. Please see `host_master_name` parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The password of the user.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("permissions")]
        private InputList<Inputs.MdbPostgresqlClusterUserPermissionGetArgs>? _permissions;

        /// <summary>
        /// Set of permissions granted to the user. The structure is documented below.
        /// </summary>
        public InputList<Inputs.MdbPostgresqlClusterUserPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.MdbPostgresqlClusterUserPermissionGetArgs>());
            set => _permissions = value;
        }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of user settings. List of settings is documented below.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        public MdbPostgresqlClusterUserGetArgs()
        {
        }
    }
}
