// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("cast")]
        public Input<string>? Cast { get; set; }

        [Input("collation")]
        public Input<string>? Collation { get; set; }

        [Input("constraint")]
        public Input<string>? Constraint { get; set; }

        [Input("defaultValues")]
        public Input<string>? DefaultValues { get; set; }

        [Input("fkConstraint")]
        public Input<string>? FkConstraint { get; set; }

        [Input("function")]
        public Input<string>? Function { get; set; }

        [Input("index")]
        public Input<string>? Index { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        [Input("rule")]
        public Input<string>? Rule { get; set; }

        [Input("sequence")]
        public Input<string>? Sequence { get; set; }

        [Input("sequenceOwnedBy")]
        public Input<string>? SequenceOwnedBy { get; set; }

        [Input("table")]
        public Input<string>? Table { get; set; }

        [Input("trigger")]
        public Input<string>? Trigger { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        public DatatransferEndpointSettingsPostgresSourceObjectTransferSettingsArgs()
        {
        }
    }
}
