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
    public sealed class GetComputeInstanceNetworkInterfaceResult
    {
        /// <summary>
        /// List of configurations for creating ipv4 DNS records. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceNetworkInterfaceDnsRecordResult> DnsRecords;
        /// <summary>
        /// The index of the network interface, generated by the server.
        /// </summary>
        public readonly int Index;
        /// <summary>
        /// The assignd private IP address to the network interface.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// Show if IPv4 address is assigned to the network interface.
        /// </summary>
        public readonly bool Ipv4;
        public readonly bool Ipv6;
        public readonly string Ipv6Address;
        /// <summary>
        /// List of configurations for creating ipv6 DNS records. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceNetworkInterfaceIpv6DnsRecordResult> Ipv6DnsRecords;
        /// <summary>
        /// MAC address that is assigned to the network interface.
        /// </summary>
        public readonly string MacAddress;
        /// <summary>
        /// Assigned for the instance's public address that is used to access the internet over NAT.
        /// </summary>
        public readonly bool Nat;
        /// <summary>
        /// List of configurations for creating ipv4 NAT DNS records. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstanceNetworkInterfaceNatDnsRecordResult> NatDnsRecords;
        /// <summary>
        /// Public IP address of the instance.
        /// </summary>
        public readonly string NatIpAddress;
        /// <summary>
        /// IP version for the public address.
        /// </summary>
        public readonly string NatIpVersion;
        /// <summary>
        /// Security group ids for network interface.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private GetComputeInstanceNetworkInterfaceResult(
            ImmutableArray<Outputs.GetComputeInstanceNetworkInterfaceDnsRecordResult> dnsRecords,

            int index,

            string ipAddress,

            bool ipv4,

            bool ipv6,

            string ipv6Address,

            ImmutableArray<Outputs.GetComputeInstanceNetworkInterfaceIpv6DnsRecordResult> ipv6DnsRecords,

            string macAddress,

            bool nat,

            ImmutableArray<Outputs.GetComputeInstanceNetworkInterfaceNatDnsRecordResult> natDnsRecords,

            string natIpAddress,

            string natIpVersion,

            ImmutableArray<string> securityGroupIds,

            string subnetId)
        {
            DnsRecords = dnsRecords;
            Index = index;
            IpAddress = ipAddress;
            Ipv4 = ipv4;
            Ipv6 = ipv6;
            Ipv6Address = ipv6Address;
            Ipv6DnsRecords = ipv6DnsRecords;
            MacAddress = macAddress;
            Nat = nat;
            NatDnsRecords = natDnsRecords;
            NatIpAddress = natIpAddress;
            NatIpVersion = natIpVersion;
            SecurityGroupIds = securityGroupIds;
            SubnetId = subnetId;
        }
    }
}
