// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbVirtualHostRouteHttpRouteHttpMatchArgs : Pulumi.ResourceArgs
    {
        [Input("httpMethods")]
        private InputList<object>? _httpMethods;

        /// <summary>
        /// List of methods(strings).
        /// </summary>
        public InputList<object> HttpMethods
        {
            get => _httpMethods ?? (_httpMethods = new InputList<object>());
            set => _httpMethods = value;
        }

        /// <summary>
        /// If not set, '/' is assumed. The structure is documented below.
        /// </summary>
        [Input("path")]
        public Input<Inputs.AlbVirtualHostRouteHttpRouteHttpMatchPathArgs>? Path { get; set; }

        public AlbVirtualHostRouteHttpRouteHttpMatchArgs()
        {
        }
    }
}