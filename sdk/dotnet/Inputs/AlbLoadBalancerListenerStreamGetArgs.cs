// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerStreamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP handler that sets plaintext HTTP router. The structure is documented below.
        /// </summary>
        [Input("handler")]
        public Input<Inputs.AlbLoadBalancerListenerStreamHandlerGetArgs>? Handler { get; set; }

        public AlbLoadBalancerListenerStreamGetArgs()
        {
        }
    }
}