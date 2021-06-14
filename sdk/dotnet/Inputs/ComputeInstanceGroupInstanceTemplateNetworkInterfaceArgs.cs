// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceGroupInstanceTemplateNetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        [Input("dnsRecords")]
        private InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordArgs>? _dnsRecords;

        /// <summary>
        /// List of dns records.  The structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordArgs> DnsRecords
        {
            get => _dnsRecords ?? (_dnsRecords = new InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceDnsRecordArgs>());
            set => _dnsRecords = value;
        }

        /// <summary>
        /// Manual set static IP address.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// True if IPv4 address allocated for the network interface.
        /// </summary>
        [Input("ipv4")]
        public Input<bool>? Ipv4 { get; set; }

        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        /// <summary>
        /// Manual set static IPv6 address.
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("ipv6DnsRecords")]
        private InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordArgs>? _ipv6DnsRecords;

        /// <summary>
        /// List of ipv6 dns records.  The structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordArgs> Ipv6DnsRecords
        {
            get => _ipv6DnsRecords ?? (_ipv6DnsRecords = new InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceIpv6DnsRecordArgs>());
            set => _ipv6DnsRecords = value;
        }

        /// <summary>
        /// A public address that can be used to access the internet over NAT.
        /// </summary>
        [Input("nat")]
        public Input<bool>? Nat { get; set; }

        [Input("natDnsRecords")]
        private InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordArgs>? _natDnsRecords;

        /// <summary>
        /// List of nat dns records.  The structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordArgs> NatDnsRecords
        {
            get => _natDnsRecords ?? (_natDnsRecords = new InputList<Inputs.ComputeInstanceGroupInstanceTemplateNetworkInterfaceNatDnsRecordArgs>());
            set => _natDnsRecords = value;
        }

        /// <summary>
        /// The ID of the network.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group ids for network interface.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The ID of the subnets to attach this interface to.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public ComputeInstanceGroupInstanceTemplateNetworkInterfaceArgs()
        {
        }
    }
}
