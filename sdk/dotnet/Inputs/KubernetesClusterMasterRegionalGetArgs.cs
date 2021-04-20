// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class KubernetesClusterMasterRegionalGetArgs : Pulumi.ResourceArgs
    {
        [Input("locations")]
        private InputList<Inputs.KubernetesClusterMasterRegionalLocationGetArgs>? _locations;

        /// <summary>
        /// Array of locations, where master instances will be allocated. The structure is documented below.
        /// </summary>
        public InputList<Inputs.KubernetesClusterMasterRegionalLocationGetArgs> Locations
        {
            get => _locations ?? (_locations = new InputList<Inputs.KubernetesClusterMasterRegionalLocationGetArgs>());
            set => _locations = value;
        }

        /// <summary>
        /// (Required) Name of availability region (e.g. "ru-central1"), where master instances will be allocated.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public KubernetesClusterMasterRegionalGetArgs()
        {
        }
    }
}
