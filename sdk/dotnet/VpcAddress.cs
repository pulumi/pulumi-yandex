// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    /// <summary>
    /// Manages a address within the Yandex.Cloud. For more information, see
    /// [the official documentation](https://cloud.yandex.com/docs/vpc/concepts/address).
    /// 
    /// * How-to Guides
    ///     * [Cloud Networking](https://cloud.yandex.com/docs/vpc/)
    ///     * [VPC Addressing](https://cloud.yandex.com/docs/vpc/concepts/address)
    /// 
    /// ## Example Usage
    /// ### External ipv4 address
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var addr = new Yandex.VpcAddress("addr", new Yandex.VpcAddressArgs
    ///         {
    ///             ExternalIpv4Address = new Yandex.Inputs.VpcAddressExternalIpv4AddressArgs
    ///             {
    ///                 ZoneId = "ru-central1-a",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Address with DDoS protection
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var vpnaddr = new Yandex.VpcAddress("vpnaddr", new Yandex.VpcAddressArgs
    ///         {
    ///             ExternalIpv4Address = new Yandex.Inputs.VpcAddressExternalIpv4AddressArgs
    ///             {
    ///                 DdosProtectionProvider = "qrator",
    ///                 ZoneId = "ru-central1-a",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// A address can be imported using the `id` of the resource, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import yandex:index/vpcAddress:VpcAddress addr address_id
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/vpcAddress:VpcAddress")]
    public partial class VpcAddress : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// spec of IP v4 address
        /// ---
        /// </summary>
        [Output("externalIpv4Address")]
        public Output<Outputs.VpcAddressExternalIpv4Address?> ExternalIpv4Address { get; private set; } = null!;

        /// <summary>
        /// ID of the folder that the resource belongs to. If it
        /// is not provided, the default provider folder is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this resource. A list of key/value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the address. Provided by the client when the address is created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `false` means that address is ephemeral.
        /// </summary>
        [Output("reserved")]
        public Output<bool> Reserved { get; private set; } = null!;

        /// <summary>
        /// `true` if address is used.
        /// </summary>
        [Output("used")]
        public Output<bool> Used { get; private set; } = null!;


        /// <summary>
        /// Create a VpcAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcAddress(string name, VpcAddressArgs? args = null, CustomResourceOptions? options = null)
            : base("yandex:index/vpcAddress:VpcAddress", name, args ?? new VpcAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcAddress(string name, Input<string> id, VpcAddressState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/vpcAddress:VpcAddress", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcAddress Get(string name, Input<string> id, VpcAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcAddress(name, id, state, options);
        }
    }

    public sealed class VpcAddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// spec of IP v4 address
        /// ---
        /// </summary>
        [Input("externalIpv4Address")]
        public Input<Inputs.VpcAddressExternalIpv4AddressArgs>? ExternalIpv4Address { get; set; }

        /// <summary>
        /// ID of the folder that the resource belongs to. If it
        /// is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this resource. A list of key/value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the address. Provided by the client when the address is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VpcAddressArgs()
        {
        }
    }

    public sealed class VpcAddressState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// spec of IP v4 address
        /// ---
        /// </summary>
        [Input("externalIpv4Address")]
        public Input<Inputs.VpcAddressExternalIpv4AddressGetArgs>? ExternalIpv4Address { get; set; }

        /// <summary>
        /// ID of the folder that the resource belongs to. If it
        /// is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this resource. A list of key/value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the address. Provided by the client when the address is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `false` means that address is ephemeral.
        /// </summary>
        [Input("reserved")]
        public Input<bool>? Reserved { get; set; }

        /// <summary>
        /// `true` if address is used.
        /// </summary>
        [Input("used")]
        public Input<bool>? Used { get; set; }

        public VpcAddressState()
        {
        }
    }
}
