// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbVirtualHostRouteHttpRouteHttpMatchPathGetArgs : Pulumi.ResourceArgs
    {
        [Input("exact")]
        public Input<string>? Exact { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public AlbVirtualHostRouteHttpRouteHttpMatchPathGetArgs()
        {
        }
    }
}
