// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsMysqlSourceObjectTransferSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("routine")]
        public Input<string>? Routine { get; set; }

        [Input("trigger")]
        public Input<string>? Trigger { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        public DatatransferEndpointSettingsMysqlSourceObjectTransferSettingsArgs()
        {
        }
    }
}