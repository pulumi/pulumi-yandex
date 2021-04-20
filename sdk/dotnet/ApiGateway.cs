// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/apiGateway:ApiGateway")]
    public partial class ApiGateway : Pulumi.CustomResource
    {
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        [Output("logGroupId")]
        public Output<string> LogGroupId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("spec")]
        public Output<string> Spec { get; private set; } = null!;

        [Output("specContentHash")]
        public Output<int?> SpecContentHash { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ApiGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiGateway(string name, ApiGatewayArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/apiGateway:ApiGateway", name, args ?? new ApiGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiGateway(string name, Input<string> id, ApiGatewayState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/apiGateway:ApiGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiGateway Get(string name, Input<string> id, ApiGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiGateway(name, id, state, options);
        }
    }

    public sealed class ApiGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("spec", required: true)]
        public Input<string> Spec { get; set; } = null!;

        [Input("specContentHash")]
        public Input<int>? SpecContentHash { get; set; }

        public ApiGatewayArgs()
        {
        }
    }

    public sealed class ApiGatewayState : Pulumi.ResourceArgs
    {
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("logGroupId")]
        public Input<string>? LogGroupId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("spec")]
        public Input<string>? Spec { get; set; }

        [Input("specContentHash")]
        public Input<int>? SpecContentHash { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public ApiGatewayState()
        {
        }
    }
}
