// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class CdnResourceSslCertificateGetArgs : Pulumi.ResourceArgs
    {
        [Input("certificateManagerId")]
        public Input<string>? CertificateManagerId { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CdnResourceSslCertificateGetArgs()
        {
        }
    }
}
