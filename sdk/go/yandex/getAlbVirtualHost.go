// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex ALB Virtual Host. For more information, see
// [Yandex.Cloud Application Load Balancer](https://cloud.yandex.com/en/docs/application-load-balancer/quickstart).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := yandex_alb_virtual_host.My - vh.Name
// 		opt1 := yandex_alb_virtual_host.My - router.Id
// 		_, err := yandex.LookupAlbVirtualHost(ctx, &yandex.LookupAlbVirtualHostArgs{
// 			Name:         &opt0,
// 			HttpRouterId: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// This data source is used to define [Application Load Balancer Virtual Host] that can be used by other resources.
func LookupAlbVirtualHost(ctx *pulumi.Context, args *LookupAlbVirtualHostArgs, opts ...pulumi.InvokeOption) (*LookupAlbVirtualHostResult, error) {
	var rv LookupAlbVirtualHostResult
	err := ctx.Invoke("yandex:index/getAlbVirtualHost:getAlbVirtualHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlbVirtualHost.
type LookupAlbVirtualHostArgs struct {
	// HTTP Router that the resource belongs to.
	HttpRouterId *string `pulumi:"httpRouterId"`
	// Name of the Virtual Host.
	Name *string `pulumi:"name"`
	// The ID of a specific Virtual Host. Virtual Host ID is concatenation of HTTP Router ID
	// and Virtual Host name with `/` symbol between them. For Example, "http_router_id/vhost_name".
	VirtualHostId *string `pulumi:"virtualHostId"`
}

// A collection of values returned by getAlbVirtualHost.
type LookupAlbVirtualHostResult struct {
	// A list of domains (host/authority header) that will be matched to this virtual host. Wildcard hosts are
	// supported in the form of '*.foo.com' or '*-bar.foo.com'. If not specified, all domains will be matched.
	Authorities  []string `pulumi:"authorities"`
	HttpRouterId string   `pulumi:"httpRouterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Apply the following modifications to the request headers. The structure is documented
	// below.
	ModifyRequestHeaders []GetAlbVirtualHostModifyRequestHeader `pulumi:"modifyRequestHeaders"`
	// Apply the following modifications to the response headers. The structure is documented
	// below.
	ModifyResponseHeaders []GetAlbVirtualHostModifyResponseHeader `pulumi:"modifyResponseHeaders"`
	// name of the route.
	Name string `pulumi:"name"`
	// A Route resource. Routes are matched *in-order*. Be careful when adding them to the end. For instance,
	// having http '/' match first makes all other routes unused. The structure is documented below.
	Routes        []GetAlbVirtualHostRoute `pulumi:"routes"`
	VirtualHostId string                   `pulumi:"virtualHostId"`
}
