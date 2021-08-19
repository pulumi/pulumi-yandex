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
    public sealed class AlbVirtualHostRouteHttpRouteHttpMatch
    {
        /// <summary>
        /// List of methods(strings).
        /// </summary>
        public readonly ImmutableArray<object> HttpMethods;
        /// <summary>
        /// If not set, '/' is assumed. The structure is documented below.
        /// </summary>
        public readonly Outputs.AlbVirtualHostRouteHttpRouteHttpMatchPath? Path;

        [OutputConstructor]
        private AlbVirtualHostRouteHttpRouteHttpMatch(
            ImmutableArray<object> httpMethods,

            Outputs.AlbVirtualHostRouteHttpRouteHttpMatchPath? path)
        {
            HttpMethods = httpMethods;
            Path = path;
        }
    }
}
