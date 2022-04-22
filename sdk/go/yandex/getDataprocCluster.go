// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Data Proc cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/data-proc/).
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
// 		foo, err := yandex.LookupDataprocCluster(ctx, &GetDataprocClusterArgs{
// 			Name: pulumi.StringRef("test"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("serviceAccountId", foo.ServiceAccountId)
// 		return nil
// 	})
// }
// ```
func LookupDataprocCluster(ctx *pulumi.Context, args *LookupDataprocClusterArgs, opts ...pulumi.InvokeOption) (*LookupDataprocClusterResult, error) {
	var rv LookupDataprocClusterResult
	err := ctx.Invoke("yandex:index/getDataprocCluster:getDataprocCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataprocCluster.
type LookupDataprocClusterArgs struct {
	// The ID of the Data Proc cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The name of the Data Proc cluster.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDataprocCluster.
type LookupDataprocClusterResult struct {
	// Name of the Object Storage bucket used for Data Proc jobs.
	Bucket string `pulumi:"bucket"`
	// Configuration and resources of the cluster. The structure is documented below.
	ClusterConfigs []GetDataprocClusterClusterConfig `pulumi:"clusterConfigs"`
	ClusterId      string                            `pulumi:"clusterId"`
	// The Data Proc cluster creation timestamp.
	CreatedAt          string `pulumi:"createdAt"`
	DeletionProtection bool   `pulumi:"deletionProtection"`
	// Description of the Data Proc cluster.
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	// A list of IDs of the host groups hosting VMs of the cluster.
	HostGroupIds []string `pulumi:"hostGroupIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of key/value label pairs assigned to the Data Proc cluster.
	Labels map[string]string `pulumi:"labels"`
	// Name of the Data Proc subcluster.
	Name             string   `pulumi:"name"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Service account used by the Data Proc agent to access resources of Yandex.Cloud.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// Whether UI proxy feature is enabled.
	UiProxy bool `pulumi:"uiProxy"`
	// ID of the availability zone where the cluster resides.
	ZoneId string `pulumi:"zoneId"`
}

func LookupDataprocClusterOutput(ctx *pulumi.Context, args LookupDataprocClusterOutputArgs, opts ...pulumi.InvokeOption) LookupDataprocClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataprocClusterResult, error) {
			args := v.(LookupDataprocClusterArgs)
			r, err := LookupDataprocCluster(ctx, &args, opts...)
			var s LookupDataprocClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataprocClusterResultOutput)
}

// A collection of arguments for invoking getDataprocCluster.
type LookupDataprocClusterOutputArgs struct {
	// The ID of the Data Proc cluster.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// The name of the Data Proc cluster.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupDataprocClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataprocClusterArgs)(nil)).Elem()
}

// A collection of values returned by getDataprocCluster.
type LookupDataprocClusterResultOutput struct{ *pulumi.OutputState }

func (LookupDataprocClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataprocClusterResult)(nil)).Elem()
}

func (o LookupDataprocClusterResultOutput) ToLookupDataprocClusterResultOutput() LookupDataprocClusterResultOutput {
	return o
}

func (o LookupDataprocClusterResultOutput) ToLookupDataprocClusterResultOutputWithContext(ctx context.Context) LookupDataprocClusterResultOutput {
	return o
}

// Name of the Object Storage bucket used for Data Proc jobs.
func (o LookupDataprocClusterResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// Configuration and resources of the cluster. The structure is documented below.
func (o LookupDataprocClusterResultOutput) ClusterConfigs() GetDataprocClusterClusterConfigArrayOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) []GetDataprocClusterClusterConfig { return v.ClusterConfigs }).(GetDataprocClusterClusterConfigArrayOutput)
}

func (o LookupDataprocClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The Data Proc cluster creation timestamp.
func (o LookupDataprocClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupDataprocClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the Data Proc cluster.
func (o LookupDataprocClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupDataprocClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// A list of IDs of the host groups hosting VMs of the cluster.
func (o LookupDataprocClusterResultOutput) HostGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) []string { return v.HostGroupIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDataprocClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of key/value label pairs assigned to the Data Proc cluster.
func (o LookupDataprocClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the Data Proc subcluster.
func (o LookupDataprocClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataprocClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Service account used by the Data Proc agent to access resources of Yandex.Cloud.
func (o LookupDataprocClusterResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

// Whether UI proxy feature is enabled.
func (o LookupDataprocClusterResultOutput) UiProxy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) bool { return v.UiProxy }).(pulumi.BoolOutput)
}

// ID of the availability zone where the cluster resides.
func (o LookupDataprocClusterResultOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocClusterResult) string { return v.ZoneId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataprocClusterResultOutput{})
}
