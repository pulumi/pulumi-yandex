// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class KubernetesClusterKmsProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// KMS key ID.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        public KubernetesClusterKmsProviderArgs()
        {
        }
    }
}
