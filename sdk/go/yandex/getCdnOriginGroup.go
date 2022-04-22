// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex CDN Origin Group. For more information, see
// [the official documentation](https://cloud.yandex.ru/docs/cdn/concepts/origins).
//
// > **_NOTE:_**  CDN provider must be activated prior usage of CDN resources, either via UI console or via yc cli command: ```yc cdn provider activate --folder-id <folder-id> --type gcore```
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
// 		myGroup, err := yandex.LookupCdnOriginGroup(ctx, &GetCdnOriginGroupArgs{
// 			OriginGroupId: pulumi.IntRef("some_instance_id"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("originGroupName", myGroup.Name)
// 		return nil
// 	})
// }
// ```
func LookupCdnOriginGroup(ctx *pulumi.Context, args *LookupCdnOriginGroupArgs, opts ...pulumi.InvokeOption) (*LookupCdnOriginGroupResult, error) {
	var rv LookupCdnOriginGroupResult
	err := ctx.Invoke("yandex:index/getCdnOriginGroup:getCdnOriginGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCdnOriginGroup.
type LookupCdnOriginGroupArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Name of the origin group.
	Name *string `pulumi:"name"`
	// The ID of a specific origin group.
	OriginGroupId *int `pulumi:"originGroupId"`
}

// A collection of values returned by getCdnOriginGroup.
type LookupCdnOriginGroupResult struct {
	FolderId string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id            string                    `pulumi:"id"`
	Name          string                    `pulumi:"name"`
	OriginGroupId int                       `pulumi:"originGroupId"`
	Origins       []GetCdnOriginGroupOrigin `pulumi:"origins"`
	UseNext       bool                      `pulumi:"useNext"`
}

func LookupCdnOriginGroupOutput(ctx *pulumi.Context, args LookupCdnOriginGroupOutputArgs, opts ...pulumi.InvokeOption) LookupCdnOriginGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCdnOriginGroupResult, error) {
			args := v.(LookupCdnOriginGroupArgs)
			r, err := LookupCdnOriginGroup(ctx, &args, opts...)
			var s LookupCdnOriginGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCdnOriginGroupResultOutput)
}

// A collection of arguments for invoking getCdnOriginGroup.
type LookupCdnOriginGroupOutputArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// Name of the origin group.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of a specific origin group.
	OriginGroupId pulumi.IntPtrInput `pulumi:"originGroupId"`
}

func (LookupCdnOriginGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCdnOriginGroupArgs)(nil)).Elem()
}

// A collection of values returned by getCdnOriginGroup.
type LookupCdnOriginGroupResultOutput struct{ *pulumi.OutputState }

func (LookupCdnOriginGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCdnOriginGroupResult)(nil)).Elem()
}

func (o LookupCdnOriginGroupResultOutput) ToLookupCdnOriginGroupResultOutput() LookupCdnOriginGroupResultOutput {
	return o
}

func (o LookupCdnOriginGroupResultOutput) ToLookupCdnOriginGroupResultOutputWithContext(ctx context.Context) LookupCdnOriginGroupResultOutput {
	return o
}

func (o LookupCdnOriginGroupResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnOriginGroupResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCdnOriginGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnOriginGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCdnOriginGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnOriginGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCdnOriginGroupResultOutput) OriginGroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCdnOriginGroupResult) int { return v.OriginGroupId }).(pulumi.IntOutput)
}

func (o LookupCdnOriginGroupResultOutput) Origins() GetCdnOriginGroupOriginArrayOutput {
	return o.ApplyT(func(v LookupCdnOriginGroupResult) []GetCdnOriginGroupOrigin { return v.Origins }).(GetCdnOriginGroupOriginArrayOutput)
}

func (o LookupCdnOriginGroupResultOutput) UseNext() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCdnOriginGroupResult) bool { return v.UseNext }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCdnOriginGroupResultOutput{})
}
