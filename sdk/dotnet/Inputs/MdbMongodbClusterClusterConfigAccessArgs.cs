// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbMongodbClusterClusterConfigAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow access for DataLens.
        /// </summary>
        [Input("dataLens")]
        public Input<bool>? DataLens { get; set; }

        public MdbMongodbClusterClusterConfigAccessArgs()
        {
        }
    }
}
