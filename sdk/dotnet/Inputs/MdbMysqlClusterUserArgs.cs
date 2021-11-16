// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbMysqlClusterUserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD` (for version 5.7 `MYSQL_NATIVE_PASSWORD`, `SHA256_PASSWORD`)
        /// </summary>
        [Input("authenticationPlugin")]
        public Input<string>? AuthenticationPlugin { get; set; }

        /// <summary>
        /// User's connection limits. The structure is documented below.
        /// If the attribute is not specified there will be no changes.
        /// </summary>
        [Input("connectionLimits")]
        public Input<Inputs.MdbMysqlClusterUserConnectionLimitsArgs>? ConnectionLimits { get; set; }

        [Input("globalPermissions")]
        private InputList<string>? _globalPermissions;

        /// <summary>
        /// List user's global permissions     
        /// Allowed permissions:  `REPLICATION_CLIENT`, `REPLICATION_SLAVE`, `PROCESS` for clear list use empty list.
        /// If the attribute is not specified there will be no changes.
        /// </summary>
        public InputList<string> GlobalPermissions
        {
            get => _globalPermissions ?? (_globalPermissions = new InputList<string>());
            set => _globalPermissions = value;
        }

        /// <summary>
        /// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replication_source_name` parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The password of the user.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("permissions")]
        private InputList<Inputs.MdbMysqlClusterUserPermissionArgs>? _permissions;

        /// <summary>
        /// Set of permissions granted to the user. The structure is documented below.
        /// </summary>
        public InputList<Inputs.MdbMysqlClusterUserPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.MdbMysqlClusterUserPermissionArgs>());
            set => _permissions = value;
        }

        public MdbMysqlClusterUserArgs()
        {
        }
    }
}
