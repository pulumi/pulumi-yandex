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
    public sealed class GetMdbPostgresqlClusterConfigPerformanceDiagnosticResult
    {
        /// <summary>
        /// Flag, when true, performance diagnostics is enabled
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Interval (in seconds) for pg_stat_activity sampling Acceptable values are 1 to 86400, inclusive.
        /// </summary>
        public readonly int SessionsSamplingInterval;
        /// <summary>
        /// Interval (in seconds) for pg_stat_statements sampling Acceptable values are 1 to 86400, inclusive.
        /// </summary>
        public readonly int StatementsSamplingInterval;

        [OutputConstructor]
        private GetMdbPostgresqlClusterConfigPerformanceDiagnosticResult(
            bool enabled,

            int sessionsSamplingInterval,

            int statementsSamplingInterval)
        {
            Enabled = enabled;
            SessionsSamplingInterval = sessionsSamplingInterval;
            StatementsSamplingInterval = statementsSamplingInterval;
        }
    }
}