// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbKafkaClusterUserGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the topic.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The password of the user.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("permissions")]
        private InputList<Inputs.MdbKafkaClusterUserPermissionGetArgs>? _permissions;

        /// <summary>
        /// Set of permissions granted to the user. The structure is documented below.
        /// </summary>
        public InputList<Inputs.MdbKafkaClusterUserPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.MdbKafkaClusterUserPermissionGetArgs>());
            set => _permissions = value;
        }

        public MdbKafkaClusterUserGetArgs()
        {
        }
    }
}
