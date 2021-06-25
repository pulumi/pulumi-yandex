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
    public sealed class AlbVirtualHostRouteHttpRouteDirectResponseAction
    {
        /// <summary>
        /// Response body text.
        /// </summary>
        public readonly string? Body;
        /// <summary>
        /// The status of the response. Supported values are: ok, invalid_argumet, not_found, 
        /// permission_denied, unauthenticated, unimplemented, internal, unavailable.
        /// </summary>
        public readonly int? Status;

        [OutputConstructor]
        private AlbVirtualHostRouteHttpRouteDirectResponseAction(
            string? body,

            int? status)
        {
            Body = body;
            Status = status;
        }
    }
}
