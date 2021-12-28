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
    public sealed class ServerlessContainerImage
    {
        public readonly ImmutableArray<string> Args;
        public readonly ImmutableArray<string> Commands;
        public readonly string? Digest;
        public readonly ImmutableDictionary<string, string>? Environment;
        /// <summary>
        /// Invoke URL for the Yandex Cloud Serverless Container
        /// </summary>
        public readonly string Url;
        public readonly string? WorkDir;

        [OutputConstructor]
        private ServerlessContainerImage(
            ImmutableArray<string> args,

            ImmutableArray<string> commands,

            string? digest,

            ImmutableDictionary<string, string>? environment,

            string url,

            string? workDir)
        {
            Args = args;
            Commands = commands;
            Digest = digest;
            Environment = environment;
            Url = url;
            WorkDir = workDir;
        }
    }
}
