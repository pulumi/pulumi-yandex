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
    public sealed class MdbClickhouseClusterFormatSchema
    {
        /// <summary>
        /// Graphite rollup configuration name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of the model.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Model file URL. You can only use models stored in Yandex Object Storage.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private MdbClickhouseClusterFormatSchema(
            string name,

            string type,

            string uri)
        {
            Name = name;
            Type = type;
            Uri = uri;
        }
    }
}
