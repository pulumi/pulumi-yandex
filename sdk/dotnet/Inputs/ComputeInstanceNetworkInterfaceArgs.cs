// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceNetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// The private IP address to assign to the instance. If
        /// empty, the address will be automatically assigned from the specified subnet.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Allocate an IPv4 address for the interface. The default value is `true`.
        /// </summary>
        [Input("ipv4")]
        public Input<bool>? Ipv4 { get; set; }

        /// <summary>
        /// If true, allocate an IPv6 address for the interface.
        /// The address will be automatically assigned from the specified subnet.
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        /// <summary>
        /// The private IPv6 address to assign to the instance.
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// Provide a public address, for instance, to access the internet over NAT.
        /// </summary>
        [Input("nat")]
        public Input<bool>? Nat { get; set; }

        /// <summary>
        /// Provide a public address, for instance, to access the internet over NAT. Address should be already reserved in web UI.
        /// </summary>
        [Input("natIpAddress")]
        public Input<string>? NatIpAddress { get; set; }

        [Input("natIpVersion")]
        public Input<string>? NatIpVersion { get; set; }

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

        /// <summary>
        /// ID of the subnet to attach this
        /// interface to. The subnet must exist in the same zone where this instance will be
        /// created.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public ComputeInstanceNetworkInterfaceArgs()
        {
        }
    }
}
