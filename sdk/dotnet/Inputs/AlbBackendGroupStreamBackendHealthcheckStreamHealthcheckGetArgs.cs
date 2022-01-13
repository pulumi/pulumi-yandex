// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text to search in reply.
        /// </summary>
        [Input("receive")]
        public Input<string>? Receive { get; set; }

        /// <summary>
        /// Message to send. If empty, it's a connect-only health check.
        /// </summary>
        [Input("send")]
        public Input<string>? Send { get; set; }

        public AlbBackendGroupStreamBackendHealthcheckStreamHealthcheckGetArgs()
        {
        }
    }
}
