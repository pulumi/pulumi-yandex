// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbMysqlClusterUserConnectionLimitsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Max connections per hour.
        /// </summary>
        [Input("maxConnectionsPerHour")]
        public Input<int>? MaxConnectionsPerHour { get; set; }

        /// <summary>
        /// Max questions per hour.
        /// </summary>
        [Input("maxQuestionsPerHour")]
        public Input<int>? MaxQuestionsPerHour { get; set; }

        /// <summary>
        /// Max updates per hour.
        /// </summary>
        [Input("maxUpdatesPerHour")]
        public Input<int>? MaxUpdatesPerHour { get; set; }

        /// <summary>
        /// Max user connections.
        /// </summary>
        [Input("maxUserConnections")]
        public Input<int>? MaxUserConnections { get; set; }

        public MdbMysqlClusterUserConnectionLimitsArgs()
        {
        }
    }
}
