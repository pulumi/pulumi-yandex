// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// X.509 certificate of the certificate authority which issued the server's certificate, in PEM format. If empty, the server's certificate must be signed by a well-known CA.
        /// </summary>
        [Input("caCertificate")]
        public Input<string>? CaCertificate { get; set; }

        [Input("hosts")]
        private InputList<string>? _hosts;

        /// <summary>
        /// List of host names of the PostgreSQL server. Exactly one host is expected currently.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// Port for the database connection.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Identifier of the Yandex Cloud VPC subnetwork to user for accessing the database. If omitted, the server has to be accessible via Internet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// TLS settings for the server connection. Empty implies plaintext connection. The structure is documented below.
        /// </summary>
        [Input("tlsMode")]
        public Input<Inputs.DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseTlsModeArgs>? TlsMode { get; set; }

        public DatatransferEndpointSettingsMysqlTargetConnectionOnPremiseArgs()
        {
        }
    }
}
