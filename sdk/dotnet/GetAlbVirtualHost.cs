// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Yandex
{
    public static class GetAlbVirtualHost
    {
        /// <summary>
        /// Get information about a Yandex ALB Virtual Host. For more information, see
        /// [Yandex.Cloud Application Load Balancer](https://cloud.yandex.com/en/docs/application-load-balancer/quickstart).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Yandex = Pulumi.Yandex;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_vh_data = Output.Create(Yandex.GetAlbVirtualHost.InvokeAsync(new Yandex.GetAlbVirtualHostArgs
        ///         {
        ///             Name = yandex_alb_virtual_host.My_vh.Name,
        ///             HttpRouterId = yandex_alb_virtual_host.My_router.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// This data source is used to define [Application Load Balancer Virtual Host] that can be used by other resources.
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAlbVirtualHostResult> InvokeAsync(GetAlbVirtualHostArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlbVirtualHostResult>("yandex:index/getAlbVirtualHost:getAlbVirtualHost", args ?? new GetAlbVirtualHostArgs(), options.WithVersion());

        /// <summary>
        /// Get information about a Yandex ALB Virtual Host. For more information, see
        /// [Yandex.Cloud Application Load Balancer](https://cloud.yandex.com/en/docs/application-load-balancer/quickstart).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Yandex = Pulumi.Yandex;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_vh_data = Output.Create(Yandex.GetAlbVirtualHost.InvokeAsync(new Yandex.GetAlbVirtualHostArgs
        ///         {
        ///             Name = yandex_alb_virtual_host.My_vh.Name,
        ///             HttpRouterId = yandex_alb_virtual_host.My_router.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// This data source is used to define [Application Load Balancer Virtual Host] that can be used by other resources.
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAlbVirtualHostResult> Invoke(GetAlbVirtualHostInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAlbVirtualHostResult>("yandex:index/getAlbVirtualHost:getAlbVirtualHost", args ?? new GetAlbVirtualHostInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAlbVirtualHostArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// HTTP Router that the resource belongs to.
        /// </summary>
        [Input("httpRouterId")]
        public string? HttpRouterId { get; set; }

        /// <summary>
        /// Name of the Virtual Host.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of a specific Virtual Host. Virtual Host ID is concatenation of HTTP Router ID
        /// and Virtual Host name with `/` symbol between them. For Example, "http_router_id/vhost_name".
        /// </summary>
        [Input("virtualHostId")]
        public string? VirtualHostId { get; set; }

        public GetAlbVirtualHostArgs()
        {
        }
    }

    public sealed class GetAlbVirtualHostInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// HTTP Router that the resource belongs to.
        /// </summary>
        [Input("httpRouterId")]
        public Input<string>? HttpRouterId { get; set; }

        /// <summary>
        /// Name of the Virtual Host.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of a specific Virtual Host. Virtual Host ID is concatenation of HTTP Router ID
        /// and Virtual Host name with `/` symbol between them. For Example, "http_router_id/vhost_name".
        /// </summary>
        [Input("virtualHostId")]
        public Input<string>? VirtualHostId { get; set; }

        public GetAlbVirtualHostInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlbVirtualHostResult
    {
        /// <summary>
        /// A list of domains (host/authority header) that will be matched to this virtual host. Wildcard hosts are
        /// supported in the form of '*.foo.com' or '*-bar.foo.com'. If not specified, all domains will be matched.
        /// </summary>
        public readonly ImmutableArray<string> Authorities;
        public readonly string HttpRouterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Apply the following modifications to the request headers. The structure is documented
        /// below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlbVirtualHostModifyRequestHeaderResult> ModifyRequestHeaders;
        /// <summary>
        /// Apply the following modifications to the response headers. The structure is documented
        /// below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlbVirtualHostModifyResponseHeaderResult> ModifyResponseHeaders;
        /// <summary>
        /// name of the route.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A Route resource. Routes are matched *in-order*. Be careful when adding them to the end. For instance,
        /// having http '/' match first makes all other routes unused. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlbVirtualHostRouteResult> Routes;
        public readonly string VirtualHostId;

        [OutputConstructor]
        private GetAlbVirtualHostResult(
            ImmutableArray<string> authorities,

            string httpRouterId,

            string id,

            ImmutableArray<Outputs.GetAlbVirtualHostModifyRequestHeaderResult> modifyRequestHeaders,

            ImmutableArray<Outputs.GetAlbVirtualHostModifyResponseHeaderResult> modifyResponseHeaders,

            string name,

            ImmutableArray<Outputs.GetAlbVirtualHostRouteResult> routes,

            string virtualHostId)
        {
            Authorities = authorities;
            HttpRouterId = httpRouterId;
            Id = id;
            ModifyRequestHeaders = modifyRequestHeaders;
            ModifyResponseHeaders = modifyResponseHeaders;
            Name = name;
            Routes = routes;
            VirtualHostId = virtualHostId;
        }
    }
}
