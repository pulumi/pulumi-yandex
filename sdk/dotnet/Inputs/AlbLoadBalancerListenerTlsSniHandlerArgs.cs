// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerTlsSniHandlerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP handler that sets plaintext HTTP router. The structure is documented below.
        /// </summary>
        [Input("handler", required: true)]
        public Input<Inputs.AlbLoadBalancerListenerTlsSniHandlerHandlerArgs> Handler { get; set; } = null!;

        /// <summary>
        /// name of SNI match.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("serviceNames", required: true)]
        private InputList<string>? _serviceNames;
        public InputList<string> ServiceNames
        {
            get => _serviceNames ?? (_serviceNames = new InputList<string>());
            set => _serviceNames = value;
        }

        public AlbLoadBalancerListenerTlsSniHandlerArgs()
        {
        }
    }
}
