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
    public sealed class GetCdnResourceSslCertificateResult
    {
        public readonly string? CertificateManagerId;
        public readonly string Status;
        public readonly string Type;

        [OutputConstructor]
        private GetCdnResourceSslCertificateResult(
            string? certificateManagerId,

            string status,

            string type)
        {
            CertificateManagerId = certificateManagerId;
            Status = status;
            Type = type;
        }
    }
}
