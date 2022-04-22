// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlbLoadBalancer(ctx *pulumi.Context, args *LookupAlbLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupAlbLoadBalancerResult, error) {
	var rv LookupAlbLoadBalancerResult
	err := ctx.Invoke("yandex:index/getAlbLoadBalancer:getAlbLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlbLoadBalancer.
type LookupAlbLoadBalancerArgs struct {
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	Name           *string `pulumi:"name"`
}

// A collection of values returned by getAlbLoadBalancer.
type LookupAlbLoadBalancerResult struct {
	AllocationPolicies []GetAlbLoadBalancerAllocationPolicy `pulumi:"allocationPolicies"`
	CreatedAt          string                               `pulumi:"createdAt"`
	Description        string                               `pulumi:"description"`
	FolderId           string                               `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                       `pulumi:"id"`
	Labels           map[string]string            `pulumi:"labels"`
	Listeners        []GetAlbLoadBalancerListener `pulumi:"listeners"`
	LoadBalancerId   string                       `pulumi:"loadBalancerId"`
	LogGroupId       string                       `pulumi:"logGroupId"`
	Name             string                       `pulumi:"name"`
	NetworkId        string                       `pulumi:"networkId"`
	RegionId         string                       `pulumi:"regionId"`
	SecurityGroupIds []string                     `pulumi:"securityGroupIds"`
	Status           string                       `pulumi:"status"`
}

func LookupAlbLoadBalancerOutput(ctx *pulumi.Context, args LookupAlbLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupAlbLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlbLoadBalancerResult, error) {
			args := v.(LookupAlbLoadBalancerArgs)
			r, err := LookupAlbLoadBalancer(ctx, &args, opts...)
			var s LookupAlbLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlbLoadBalancerResultOutput)
}

// A collection of arguments for invoking getAlbLoadBalancer.
type LookupAlbLoadBalancerOutputArgs struct {
	LoadBalancerId pulumi.StringPtrInput `pulumi:"loadBalancerId"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAlbLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlbLoadBalancerArgs)(nil)).Elem()
}

// A collection of values returned by getAlbLoadBalancer.
type LookupAlbLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupAlbLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlbLoadBalancerResult)(nil)).Elem()
}

func (o LookupAlbLoadBalancerResultOutput) ToLookupAlbLoadBalancerResultOutput() LookupAlbLoadBalancerResultOutput {
	return o
}

func (o LookupAlbLoadBalancerResultOutput) ToLookupAlbLoadBalancerResultOutputWithContext(ctx context.Context) LookupAlbLoadBalancerResultOutput {
	return o
}

func (o LookupAlbLoadBalancerResultOutput) AllocationPolicies() GetAlbLoadBalancerAllocationPolicyArrayOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) []GetAlbLoadBalancerAllocationPolicy { return v.AllocationPolicies }).(GetAlbLoadBalancerAllocationPolicyArrayOutput)
}

func (o LookupAlbLoadBalancerResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupAlbLoadBalancerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupAlbLoadBalancerResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAlbLoadBalancerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlbLoadBalancerResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupAlbLoadBalancerResultOutput) Listeners() GetAlbLoadBalancerListenerArrayOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) []GetAlbLoadBalancerListener { return v.Listeners }).(GetAlbLoadBalancerListenerArrayOutput)
}

func (o LookupAlbLoadBalancerResultOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.LoadBalancerId }).(pulumi.StringOutput)
}

func (o LookupAlbLoadBalancerResultOutput) LogGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.LogGroupId }).(pulumi.StringOutput)
}

func (o LookupAlbLoadBalancerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAlbLoadBalancerResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

func (o LookupAlbLoadBalancerResultOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.RegionId }).(pulumi.StringOutput)
}

func (o LookupAlbLoadBalancerResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupAlbLoadBalancerResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbLoadBalancerResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlbLoadBalancerResultOutput{})
}
