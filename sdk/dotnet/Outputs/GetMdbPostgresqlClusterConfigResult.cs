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
    public sealed class GetMdbPostgresqlClusterConfigResult
    {
        /// <summary>
        /// Access policy to the PostgreSQL cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigAccessResult> Accesses;
        /// <summary>
        /// Configuration setting which enables/disables autofailover in cluster.
        /// </summary>
        public readonly bool Autofailover;
        /// <summary>
        /// The period in days during which backups are stored.
        /// </summary>
        public readonly int BackupRetainPeriodDays;
        /// <summary>
        /// Time to start the daily backup, in the UTC timezone. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigBackupWindowStartResult> BackupWindowStarts;
        /// <summary>
        /// Cluster performance diagnostics settings. The structure is documented below. [YC Documentation](https://cloud.yandex.com/docs/managed-postgresql/grpc/cluster_service#PerformanceDiagnostics)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigPerformanceDiagnosticResult> PerformanceDiagnostics;
        /// <summary>
        /// Configuration of the connection pooler. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigPoolerConfigResult> PoolerConfigs;
        /// <summary>
        /// PostgreSQL cluster config.
        /// </summary>
        public readonly ImmutableDictionary<string, string> PostgresqlConfig;
        /// <summary>
        /// Resources allocated to hosts of the PostgreSQL cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigResourceResult> Resources;
        /// <summary>
        /// Version of the extension.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetMdbPostgresqlClusterConfigResult(
            ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigAccessResult> accesses,

            bool autofailover,

            int backupRetainPeriodDays,

            ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigBackupWindowStartResult> backupWindowStarts,

            ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigPerformanceDiagnosticResult> performanceDiagnostics,

            ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigPoolerConfigResult> poolerConfigs,

            ImmutableDictionary<string, string> postgresqlConfig,

            ImmutableArray<Outputs.GetMdbPostgresqlClusterConfigResourceResult> resources,

            string version)
        {
            Accesses = accesses;
            Autofailover = autofailover;
            BackupRetainPeriodDays = backupRetainPeriodDays;
            BackupWindowStarts = backupWindowStarts;
            PerformanceDiagnostics = performanceDiagnostics;
            PoolerConfigs = poolerConfigs;
            PostgresqlConfig = postgresqlConfig;
            Resources = resources;
            Version = version;
        }
    }
}
