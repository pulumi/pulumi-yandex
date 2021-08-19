// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generates an [IAM] role document that may be referenced by and applied to
// other Yandex.Cloud Platform resources, such as the `ResourcemanagerFolder` resource. For more information, see
// [the official documentation](https://cloud.yandex.ru/docs/iam/concepts/access-control/roles).
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
// 		_, err := yandex.GetIamRole(ctx, &yandex.GetIamRoleArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"members": []string{
// 						"userAccount:user_id_1",
// 					},
// 					"role": "admin",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// This data source is used to define [IAM] roles in order to apply them to other resources.
// Currently, defining a role through a data source and referencing that role
// from another resource is the only way to apply an IAM role to a resource.
func GetIamRole(ctx *pulumi.Context, args *GetIamRoleArgs, opts ...pulumi.InvokeOption) (*GetIamRoleResult, error) {
	var rv GetIamRoleResult
	err := ctx.Invoke("yandex:index/getIamRole:getIamRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamRole.
type GetIamRoleArgs struct {
	Description *string `pulumi:"description"`
	RoleId      *string `pulumi:"roleId"`
}

// A collection of values returned by getIamRole.
type GetIamRoleResult struct {
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	RoleId string `pulumi:"roleId"`
}
