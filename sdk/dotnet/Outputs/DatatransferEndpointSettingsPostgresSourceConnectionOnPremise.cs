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
    public sealed class DatatransferEndpointSettingsPostgresSourceConnectionOnPremise
    {
        /// <summary>
        /// X.509 certificate of the certificate authority which issued the server's certificate, in PEM format. If empty, the server's certificate must be signed by a well-known CA.
        /// </summary>
        public readonly string? CaCertificate;
        /// <summary>
        /// List of host names of the PostgreSQL server. Exactly one host is expected currently.
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// Port for the database connection.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Identifier of the Yandex Cloud VPC subnetwork to user for accessing the database. If omitted, the server has to be accessible via Internet.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// TLS settings for the server connection. Empty implies plaintext connection. The structure is documented below.
        /// </summary>
        public readonly Outputs.DatatransferEndpointSettingsPostgresSourceConnectionOnPremiseTlsMode? TlsMode;

        [OutputConstructor]
        private DatatransferEndpointSettingsPostgresSourceConnectionOnPremise(
            string? caCertificate,

            ImmutableArray<string> hosts,

            int? port,

            string? subnetId,

            Outputs.DatatransferEndpointSettingsPostgresSourceConnectionOnPremiseTlsMode? tlsMode)
        {
            CaCertificate = caCertificate;
            Hosts = hosts;
            Port = port;
            SubnetId = subnetId;
            TlsMode = tlsMode;
        }
    }
}
