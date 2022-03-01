// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Load Balancer network load balancer. For more information, see
// [Yandex.Cloud Network Load Balancer](https://cloud.yandex.com/docs/load-balancer/concepts/).
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
// 		_, err := yandex.LookupLbNetworkLoadBalancer(ctx, &GetLbNetworkLoadBalancerArgs{
// 			NetworkLoadBalancerId: pulumi.StringRef("my-network-load-balancer"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// This data source is used to define [Load Balancer Network Load Balancers] that can be used by other resources.
func LookupLbNetworkLoadBalancer(ctx *pulumi.Context, args *LookupLbNetworkLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLbNetworkLoadBalancerResult, error) {
	var rv LookupLbNetworkLoadBalancerResult
	err := ctx.Invoke("yandex:index/getLbNetworkLoadBalancer:getLbNetworkLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbNetworkLoadBalancer.
type LookupLbNetworkLoadBalancerArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// - Name of the network load balancer.
	Name *string `pulumi:"name"`
	// Network load balancer ID.
	NetworkLoadBalancerId *string `pulumi:"networkLoadBalancerId"`
}

// A collection of values returned by getLbNetworkLoadBalancer.
type LookupLbNetworkLoadBalancerResult struct {
	// An attached target group is a group of targets that is attached to a load balancer. Structure is documented below.
	AttachedTargetGroups []GetLbNetworkLoadBalancerAttachedTargetGroup `pulumi:"attachedTargetGroups"`
	// Creation timestamp of this network load balancer.
	CreatedAt string `pulumi:"createdAt"`
	// Description of the network load balancer.
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Labels to assign to this network load balancer.
	Labels map[string]string `pulumi:"labels"`
	// Listener specification that will be used by a network load balancer. Structure is documented below.
	Listeners []GetLbNetworkLoadBalancerListener `pulumi:"listeners"`
	// Name of the listener.
	Name                  string `pulumi:"name"`
	NetworkLoadBalancerId string `pulumi:"networkLoadBalancerId"`
	// ID of the region where the network load balancer resides.
	RegionId string `pulumi:"regionId"`
	// Type of the network load balancer.
	Type string `pulumi:"type"`
}

func LookupLbNetworkLoadBalancerOutput(ctx *pulumi.Context, args LookupLbNetworkLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLbNetworkLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLbNetworkLoadBalancerResult, error) {
			args := v.(LookupLbNetworkLoadBalancerArgs)
			r, err := LookupLbNetworkLoadBalancer(ctx, &args, opts...)
			return *r, err
		}).(LookupLbNetworkLoadBalancerResultOutput)
}

// A collection of arguments for invoking getLbNetworkLoadBalancer.
type LookupLbNetworkLoadBalancerOutputArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// - Name of the network load balancer.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Network load balancer ID.
	NetworkLoadBalancerId pulumi.StringPtrInput `pulumi:"networkLoadBalancerId"`
}

func (LookupLbNetworkLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbNetworkLoadBalancerArgs)(nil)).Elem()
}

// A collection of values returned by getLbNetworkLoadBalancer.
type LookupLbNetworkLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLbNetworkLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbNetworkLoadBalancerResult)(nil)).Elem()
}

func (o LookupLbNetworkLoadBalancerResultOutput) ToLookupLbNetworkLoadBalancerResultOutput() LookupLbNetworkLoadBalancerResultOutput {
	return o
}

func (o LookupLbNetworkLoadBalancerResultOutput) ToLookupLbNetworkLoadBalancerResultOutputWithContext(ctx context.Context) LookupLbNetworkLoadBalancerResultOutput {
	return o
}

// An attached target group is a group of targets that is attached to a load balancer. Structure is documented below.
func (o LookupLbNetworkLoadBalancerResultOutput) AttachedTargetGroups() GetLbNetworkLoadBalancerAttachedTargetGroupArrayOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) []GetLbNetworkLoadBalancerAttachedTargetGroup {
		return v.AttachedTargetGroups
	}).(GetLbNetworkLoadBalancerAttachedTargetGroupArrayOutput)
}

// Creation timestamp of this network load balancer.
func (o LookupLbNetworkLoadBalancerResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the network load balancer.
func (o LookupLbNetworkLoadBalancerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupLbNetworkLoadBalancerResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbNetworkLoadBalancerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Labels to assign to this network load balancer.
func (o LookupLbNetworkLoadBalancerResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Listener specification that will be used by a network load balancer. Structure is documented below.
func (o LookupLbNetworkLoadBalancerResultOutput) Listeners() GetLbNetworkLoadBalancerListenerArrayOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) []GetLbNetworkLoadBalancerListener { return v.Listeners }).(GetLbNetworkLoadBalancerListenerArrayOutput)
}

// Name of the listener.
func (o LookupLbNetworkLoadBalancerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLbNetworkLoadBalancerResultOutput) NetworkLoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) string { return v.NetworkLoadBalancerId }).(pulumi.StringOutput)
}

// ID of the region where the network load balancer resides.
func (o LookupLbNetworkLoadBalancerResultOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) string { return v.RegionId }).(pulumi.StringOutput)
}

// Type of the network load balancer.
func (o LookupLbNetworkLoadBalancerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbNetworkLoadBalancerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbNetworkLoadBalancerResultOutput{})
}
