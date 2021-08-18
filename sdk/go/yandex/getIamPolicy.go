// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generates an [IAM] policy document that may be referenced by and applied to
// other Yandex.Cloud Platform resources, such as the `ResourcemanagerFolder` resource.
//
// This data source is used to define [IAM] policies to apply to other resources.
// Currently, defining a policy through a data source and referencing that policy
// from another resource is the only way to apply an IAM policy to a resource.
func GetIamPolicy(ctx *pulumi.Context, args *GetIamPolicyArgs, opts ...pulumi.InvokeOption) (*GetIamPolicyResult, error) {
	var rv GetIamPolicyResult
	err := ctx.Invoke("yandex:index/getIamPolicy:getIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamPolicy.
type GetIamPolicyArgs struct {
	// A nested configuration block (described below)
	// that defines a binding to be included in the policy document. Multiple
	// `binding` arguments are supported.
	Bindings []GetIamPolicyBinding `pulumi:"bindings"`
}

// A collection of values returned by getIamPolicy.
type GetIamPolicyResult struct {
	Bindings []GetIamPolicyBinding `pulumi:"bindings"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The above bindings serialized in a format suitable for
	// referencing from a resource that supports IAM.
	PolicyData string `pulumi:"policyData"`
}
