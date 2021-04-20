// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbPostgresqlClusterRestoreArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backup ID. The cluster will be created from the specified backup. [How to get a list of PostgreSQL backups](https://cloud.yandex.com/docs/managed-postgresql/operations/cluster-backups).
        /// </summary>
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        /// <summary>
        /// Timestamp of the moment to which the PostgreSQL cluster should be restored. (Format: "2006-01-02T15:04:05" - UTC). When not set, current time is used.
        /// </summary>
        [Input("time")]
        public Input<string>? Time { get; set; }

        /// <summary>
        /// Flag that indicates whether a database should be restored to the first backup point available just after the timestamp specified in the [time] field instead of just before.  
        /// Possible values:
        /// - false (default) — the restore point refers to the first backup moment before [time].
        /// - true — the restore point refers to the first backup point after [time].
        /// </summary>
        [Input("timeInclusive")]
        public Input<bool>? TimeInclusive { get; set; }

        public MdbPostgresqlClusterRestoreArgs()
        {
        }
    }
}
