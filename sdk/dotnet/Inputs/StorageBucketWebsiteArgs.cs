// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class StorageBucketWebsiteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An absolute path to the document to return in case of a 4XX error.
        /// </summary>
        [Input("errorDocument")]
        public Input<string>? ErrorDocument { get; set; }

        /// <summary>
        /// Storage returns this index document when requests are made to the root domain or any of the subfolders.
        /// </summary>
        [Input("indexDocument")]
        public Input<string>? IndexDocument { get; set; }

        /// <summary>
        /// A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (`http://` or `https://`) to use when redirecting requests. The default is the protocol that is used in the original request.
        /// </summary>
        [Input("redirectAllRequestsTo")]
        public Input<string>? RedirectAllRequestsTo { get; set; }

        /// <summary>
        /// A json array containing [routing rules](https://cloud.yandex.ru/docs/storage/s3/api-ref/hosting/upload#request-scheme) describing redirect behavior and when redirects are applied.
        /// </summary>
        [Input("routingRules")]
        public Input<string>? RoutingRules { get; set; }

        public StorageBucketWebsiteArgs()
        {
        }
    }
}
