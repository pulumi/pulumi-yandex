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
    public sealed class MdbMysqlClusterMaintenanceWindow
    {
        /// <summary>
        /// Day of the week (in `DDD` format). Allowed values: "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"
        /// </summary>
        public readonly string? Day;
        /// <summary>
        /// Hour of the day in UTC (in `HH` format). Allowed value is between 0 and 23.
        /// </summary>
        public readonly int? Hour;
        /// <summary>
        /// Type of maintenance window. Can be either `ANYTIME` or `WEEKLY`. A day and hour of window need to be specified with weekly window.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private MdbMysqlClusterMaintenanceWindow(
            string? day,

            int? hour,

            string type)
        {
            Day = day;
            Hour = hour;
            Type = type;
        }
    }
}
