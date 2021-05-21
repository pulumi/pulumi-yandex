// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	Description   *string `pulumi:"description"`
	FolderId      *string `pulumi:"folderId"`
	Name          *string `pulumi:"name"`
	TargetGroupId *string `pulumi:"targetGroupId"`
}

// A collection of values returned by getAlbTargetGroup.
type LookupAlbTargetGroupResult struct {
	CreatedAt   string `pulumi:"createdAt"`
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id            string                    `pulumi:"id"`
	Labels        map[string]string         `pulumi:"labels"`
	Name          string                    `pulumi:"name"`
	TargetGroupId string                    `pulumi:"targetGroupId"`
	Targets       []GetAlbTargetGroupTarget `pulumi:"targets"`
}
