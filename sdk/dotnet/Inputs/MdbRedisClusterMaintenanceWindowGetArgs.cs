// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbRedisClusterMaintenanceWindowGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Day of week for maintenance window if window type is weekly. Possible values: `MON`, `TUE`, `WED`, `THU`, `FRI`, `SAT`, `SUN`.
        /// </summary>
        [Input("day")]
        public Input<string>? Day { get; set; }

        /// <summary>
        /// Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
        /// </summary>
        [Input("hour")]
        public Input<int>? Hour { get; set; }

        /// <summary>
        /// Type of maintenance window. Can be either `ANYTIME` or `WEEKLY`. A day and hour of window need to be specified with weekly window.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MdbRedisClusterMaintenanceWindowGetArgs()
        {
        }
    }
}
