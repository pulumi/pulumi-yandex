// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Application Load Balancer target group. For more information, see
// [Yandex.Cloud Application Load Balancer](https://cloud.yandex.com/en/docs/application-load-balancer/quickstart).
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
// 		opt0 := "my-target-group-id"
// 		_, err := yandex.LookupAlbTargetGroup(ctx, &GetAlbTargetGroupArgs{
// 			TargetGroupId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// This data source is used to define [Application Load Balancer Target Groups] that can be used by other resources.
func LookupAlbTargetGroup(ctx *pulumi.Context, args *LookupAlbTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupAlbTargetGroupResult, error) {
	var rv LookupAlbTargetGroupResult
	err := ctx.Invoke("yandex:index/getAlbTargetGroup:getAlbTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlbTargetGroup.
type LookupAlbTargetGroupArgs struct {
	// Description of the target group.
	Description *string `pulumi:"description"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// - Name of the Target Group.
	Name *string `pulumi:"name"`
	// Target Group ID.
	TargetGroupId *string `pulumi:"targetGroupId"`
}

// A collection of values returned by getAlbTargetGroup.
type LookupAlbTargetGroupResult struct {
	// Creation timestamp of this target group.
	CreatedAt string `pulumi:"createdAt"`
	// Description of the target group.
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Labels to assign to this target group.
	// * `target.0.ip_address` - IP address of the target.
	// * `target.0.subnet_id` - ID of the subnet that targets are connected to.
	Labels        map[string]string         `pulumi:"labels"`
	Name          string                    `pulumi:"name"`
	TargetGroupId string                    `pulumi:"targetGroupId"`
	Targets       []GetAlbTargetGroupTarget `pulumi:"targets"`
}

func LookupAlbTargetGroupOutput(ctx *pulumi.Context, args LookupAlbTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAlbTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlbTargetGroupResult, error) {
			args := v.(LookupAlbTargetGroupArgs)
			r, err := LookupAlbTargetGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupAlbTargetGroupResultOutput)
}

// A collection of arguments for invoking getAlbTargetGroup.
type LookupAlbTargetGroupOutputArgs struct {
	// Description of the target group.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// - Name of the Target Group.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Target Group ID.
	TargetGroupId pulumi.StringPtrInput `pulumi:"targetGroupId"`
}

func (LookupAlbTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlbTargetGroupArgs)(nil)).Elem()
}

// A collection of values returned by getAlbTargetGroup.
type LookupAlbTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAlbTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlbTargetGroupResult)(nil)).Elem()
}

func (o LookupAlbTargetGroupResultOutput) ToLookupAlbTargetGroupResultOutput() LookupAlbTargetGroupResultOutput {
	return o
}

func (o LookupAlbTargetGroupResultOutput) ToLookupAlbTargetGroupResultOutputWithContext(ctx context.Context) LookupAlbTargetGroupResultOutput {
	return o
}

// Creation timestamp of this target group.
func (o LookupAlbTargetGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbTargetGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the target group.
func (o LookupAlbTargetGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbTargetGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupAlbTargetGroupResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbTargetGroupResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAlbTargetGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbTargetGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Labels to assign to this target group.
// * `target.0.ip_address` - IP address of the target.
// * `target.0.subnet_id` - ID of the subnet that targets are connected to.
func (o LookupAlbTargetGroupResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAlbTargetGroupResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupAlbTargetGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbTargetGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAlbTargetGroupResultOutput) TargetGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbTargetGroupResult) string { return v.TargetGroupId }).(pulumi.StringOutput)
}

func (o LookupAlbTargetGroupResultOutput) Targets() GetAlbTargetGroupTargetArrayOutput {
	return o.ApplyT(func(v LookupAlbTargetGroupResult) []GetAlbTargetGroupTarget { return v.Targets }).(GetAlbTargetGroupTargetArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlbTargetGroupResultOutput{})
}
