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
    public sealed class GetOrganizationmanagerSamlFederationSecuritySettingResult
    {
        /// <summary>
        /// Indicates whether encrypted assertions are enabled.
        /// </summary>
        public readonly bool EncryptedAssertions;

        [OutputConstructor]
        private GetOrganizationmanagerSamlFederationSecuritySettingResult(bool encryptedAssertions)
        {
            EncryptedAssertions = encryptedAssertions;
        }
    }
}