// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbVirtualHostRouteHttpRouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Direct response action resource. The structure is documented below.
        /// </summary>
        [Input("directResponseAction")]
        public Input<Inputs.AlbVirtualHostRouteHttpRouteDirectResponseActionArgs>? DirectResponseAction { get; set; }

        [Input("httpMatches")]
        private InputList<Inputs.AlbVirtualHostRouteHttpRouteHttpMatchArgs>? _httpMatches;

        /// <summary>
        /// Checks "/" prefix by default. The structure is documented below.
        /// </summary>
        public InputList<Inputs.AlbVirtualHostRouteHttpRouteHttpMatchArgs> HttpMatches
        {
            get => _httpMatches ?? (_httpMatches = new InputList<Inputs.AlbVirtualHostRouteHttpRouteHttpMatchArgs>());
            set => _httpMatches = value;
        }

        /// <summary>
        /// HTTP route action resource. The structure is documented below.
        /// </summary>
        [Input("httpRouteAction")]
        public Input<Inputs.AlbVirtualHostRouteHttpRouteHttpRouteActionArgs>? HttpRouteAction { get; set; }

        /// <summary>
        /// Redirect action resource. The structure is documented below.
        /// </summary>
        [Input("redirectAction")]
        public Input<Inputs.AlbVirtualHostRouteHttpRouteRedirectActionArgs>? RedirectAction { get; set; }

        public AlbVirtualHostRouteHttpRouteArgs()
        {
        }
    }
}
