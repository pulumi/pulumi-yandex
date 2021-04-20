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
    /// Manages a route table within the Yandex.Cloud. For more information, see
    /// [the official documentation](https://cloud.yandex.com/docs/vpc/concepts).
    /// 
    /// * How-to Guides
    ///     * [Cloud Networking](https://cloud.yandex.com/docs/vpc/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var lab_net = new Yandex.VpcNetwork("lab-net", new Yandex.VpcNetworkArgs
    ///         {
    ///         });
    ///         var lab_rt_a = new Yandex.VpcRouteTable("lab-rt-a", new Yandex.VpcRouteTableArgs
    ///         {
    ///             NetworkId = lab_net.Id,
    ///             StaticRoutes = 
    ///             {
    ///                 new Yandex.Inputs.VpcRouteTableStaticRouteArgs
    ///                 {
    ///                     DestinationPrefix = "10.2.0.0/16",
    ///                     NextHopAddress = "172.16.10.10",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// A route table can be imported using the `id` of the resource, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import yandex:index/vpcRouteTable:VpcRouteTable default route_table_id
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/vpcRouteTable:VpcRouteTable")]
    public partial class VpcRouteTable : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp of the route table.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// An optional description of the route table. Provide this property when
        /// you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the folder to which the resource belongs.
        /// If omitted, the provider folder is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Labels to assign to this route table. A list of key/value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the route table. Provided by the client when the route table is created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the network this route table belongs to.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// A list of static route records for the route table. The structure is documented below.
        /// </summary>
        [Output("staticRoutes")]
        public Output<ImmutableArray<Outputs.VpcRouteTableStaticRoute>> StaticRoutes { get; private set; } = null!;


        /// <summary>
        /// Create a VpcRouteTable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcRouteTable(string name, VpcRouteTableArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/vpcRouteTable:VpcRouteTable", name, args ?? new VpcRouteTableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcRouteTable(string name, Input<string> id, VpcRouteTableState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/vpcRouteTable:VpcRouteTable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcRouteTable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcRouteTable Get(string name, Input<string> id, VpcRouteTableState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcRouteTable(name, id, state, options);
        }
    }

    public sealed class VpcRouteTableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of the route table. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the folder to which the resource belongs.
        /// If omitted, the provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to assign to this route table. A list of key/value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the route table. Provided by the client when the route table is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network this route table belongs to.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        [Input("staticRoutes")]
        private InputList<Inputs.VpcRouteTableStaticRouteArgs>? _staticRoutes;

        /// <summary>
        /// A list of static route records for the route table. The structure is documented below.
        /// </summary>
        public InputList<Inputs.VpcRouteTableStaticRouteArgs> StaticRoutes
        {
            get => _staticRoutes ?? (_staticRoutes = new InputList<Inputs.VpcRouteTableStaticRouteArgs>());
            set => _staticRoutes = value;
        }

        public VpcRouteTableArgs()
        {
        }
    }

    public sealed class VpcRouteTableState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp of the route table.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// An optional description of the route table. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the folder to which the resource belongs.
        /// If omitted, the provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to assign to this route table. A list of key/value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the route table. Provided by the client when the route table is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network this route table belongs to.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("staticRoutes")]
        private InputList<Inputs.VpcRouteTableStaticRouteGetArgs>? _staticRoutes;

        /// <summary>
        /// A list of static route records for the route table. The structure is documented below.
        /// </summary>
        public InputList<Inputs.VpcRouteTableStaticRouteGetArgs> StaticRoutes
        {
            get => _staticRoutes ?? (_staticRoutes = new InputList<Inputs.VpcRouteTableStaticRouteGetArgs>());
            set => _staticRoutes = value;
        }

        public VpcRouteTableState()
        {
        }
    }
}
