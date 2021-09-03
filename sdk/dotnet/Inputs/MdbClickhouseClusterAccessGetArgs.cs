// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbClickhouseClusterAccessGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow access for DataLens. Can be either `true` or `false`.
        /// </summary>
        [Input("dataLens")]
        public Input<bool>? DataLens { get; set; }

        /// <summary>
        /// Allow access for Yandex.Metrika. Can be either `true` or `false`.
        /// </summary>
        [Input("metrika")]
        public Input<bool>? Metrika { get; set; }

        /// <summary>
        /// Allow access for Serverless. Can be either `true` or `false`.
        /// </summary>
        [Input("serverless")]
        public Input<bool>? Serverless { get; set; }

        /// <summary>
        /// Allow access for Web SQL. Can be either `true` or `false`.
        /// </summary>
        [Input("webSql")]
        public Input<bool>? WebSql { get; set; }

        public MdbClickhouseClusterAccessGetArgs()
        {
        }
    }
}
