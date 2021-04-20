// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class KubernetesClusterMasterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) PEM-encoded public certificate that is the root of trust for the Kubernetes cluster.
        /// </summary>
        [Input("clusterCaCertificate")]
        public Input<string>? ClusterCaCertificate { get; set; }

        /// <summary>
        /// (Computed) An IPv4 external network address that is assigned to the master.
        /// </summary>
        [Input("externalV4Address")]
        public Input<string>? ExternalV4Address { get; set; }

        /// <summary>
        /// (Computed) External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).
        /// </summary>
        [Input("externalV4Endpoint")]
        public Input<string>? ExternalV4Endpoint { get; set; }

        /// <summary>
        /// (Computed) An IPv4 internal network address that is assigned to the master.
        /// </summary>
        [Input("internalV4Address")]
        public Input<string>? InternalV4Address { get; set; }

        /// <summary>
        /// (Computed) Internal endpoint that can be used to connect to the master from cloud networks.
        /// </summary>
        [Input("internalV4Endpoint")]
        public Input<string>? InternalV4Endpoint { get; set; }

        /// <summary>
        /// (Optional) (Computed) Maintenance policy for Kubernetes master.
        /// If policy is omitted, automatic revision upgrades of the kubernetes master are enabled and could happen at any time.
        /// Revision upgrades are performed only within the same minor version, e.g. 1.13.
        /// Minor version upgrades (e.g. 1.13-&gt;1.14) should be performed manually. The structure is documented below.
        /// </summary>
        [Input("maintenancePolicy")]
        public Input<Inputs.KubernetesClusterMasterMaintenancePolicyGetArgs>? MaintenancePolicy { get; set; }

        /// <summary>
        /// (Optional) (Computed) Boolean flag. When `true`, Kubernetes master will have visible ipv4 address.
        /// </summary>
        [Input("publicIp")]
        public Input<bool>? PublicIp { get; set; }

        /// <summary>
        /// (Optional) Initialize parameters for Regional Master (highly available master). The structure is documented below.
        /// </summary>
        [Input("regional")]
        public Input<Inputs.KubernetesClusterMasterRegionalGetArgs>? Regional { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// (Optional) List of security group IDs to which the Kubernetes cluster belongs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// (Optional) (Computed) Version of Kubernetes that will be used for master.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// (Computed) Information about cluster version. The structure is documented below.
        /// </summary>
        [Input("versionInfo")]
        public Input<Inputs.KubernetesClusterMasterVersionInfoGetArgs>? VersionInfo { get; set; }

        /// <summary>
        /// (Optional) Initialize parameters for Zonal Master (single node master). The structure is documented below.
        /// </summary>
        [Input("zonal")]
        public Input<Inputs.KubernetesClusterMasterZonalGetArgs>? Zonal { get; set; }

        public KubernetesClusterMasterGetArgs()
        {
        }
    }
}
