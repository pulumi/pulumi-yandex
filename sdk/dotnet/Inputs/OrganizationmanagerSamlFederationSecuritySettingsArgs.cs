// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class OrganizationmanagerSamlFederationSecuritySettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable encrypted assertions.
        /// </summary>
        [Input("encryptedAssertions", required: true)]
        public Input<bool> EncryptedAssertions { get; set; } = null!;

        public OrganizationmanagerSamlFederationSecuritySettingsArgs()
        {
        }
    }
}