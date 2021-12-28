// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Managed Redis cluster. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-redis/concepts).
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
// 		opt0 := "test"
// 		foo, err := yandex.LookupMdbRedisCluster(ctx, &GetMdbRedisClusterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("networkId", foo.NetworkId)
// 		return nil
// 	})
// }
// ```
func LookupMdbRedisCluster(ctx *pulumi.Context, args *LookupMdbRedisClusterArgs, opts ...pulumi.InvokeOption) (*LookupMdbRedisClusterResult, error) {
	var rv LookupMdbRedisClusterResult
	err := ctx.Invoke("yandex:index/getMdbRedisCluster:getMdbRedisCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbRedisCluster.
type LookupMdbRedisClusterArgs struct {
	// The ID of the Redis cluster.
	ClusterId          *string `pulumi:"clusterId"`
	DeletionProtection *bool   `pulumi:"deletionProtection"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// The name of the Redis cluster.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getMdbRedisCluster.
type LookupMdbRedisClusterResult struct {
	ClusterId string `pulumi:"clusterId"`
	// Configuration of the Redis cluster. The structure is documented below.
	Configs []GetMdbRedisClusterConfig `pulumi:"configs"`
	// Creation timestamp of the key.
	CreatedAt          string `pulumi:"createdAt"`
	DeletionProtection bool   `pulumi:"deletionProtection"`
	// Description of the Redis cluster.
	Description string `pulumi:"description"`
	// Deployment environment of the Redis cluster.
	Environment string `pulumi:"environment"`
	FolderId    string `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health string `pulumi:"health"`
	// A host of the Redis cluster. The structure is documented below.
	Hosts []GetMdbRedisClusterHost `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of key/value label pairs to assign to the Redis cluster.
	Labels             map[string]string                     `pulumi:"labels"`
	MaintenanceWindows []GetMdbRedisClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	Name               string                                `pulumi:"name"`
	// ID of the network, to which the Redis cluster belongs.
	NetworkId string `pulumi:"networkId"`
	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources []GetMdbRedisClusterResource `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Redis Cluster mode enabled/disabled.
	Sharded bool `pulumi:"sharded"`
	// Status of the cluster.
	Status string `pulumi:"status"`
	// tls support mode enabled/disabled.
	TlsEnabled bool `pulumi:"tlsEnabled"`
}

func LookupMdbRedisClusterOutput(ctx *pulumi.Context, args LookupMdbRedisClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMdbRedisClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMdbRedisClusterResult, error) {
			args := v.(LookupMdbRedisClusterArgs)
			r, err := LookupMdbRedisCluster(ctx, &args, opts...)
			return *r, err
		}).(LookupMdbRedisClusterResultOutput)
}

// A collection of arguments for invoking getMdbRedisCluster.
type LookupMdbRedisClusterOutputArgs struct {
	// The ID of the Redis cluster.
	ClusterId          pulumi.StringPtrInput `pulumi:"clusterId"`
	DeletionProtection pulumi.BoolPtrInput   `pulumi:"deletionProtection"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// The name of the Redis cluster.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupMdbRedisClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbRedisClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbRedisCluster.
type LookupMdbRedisClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMdbRedisClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbRedisClusterResult)(nil)).Elem()
}

func (o LookupMdbRedisClusterResultOutput) ToLookupMdbRedisClusterResultOutput() LookupMdbRedisClusterResultOutput {
	return o
}

func (o LookupMdbRedisClusterResultOutput) ToLookupMdbRedisClusterResultOutputWithContext(ctx context.Context) LookupMdbRedisClusterResultOutput {
	return o
}

func (o LookupMdbRedisClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Configuration of the Redis cluster. The structure is documented below.
func (o LookupMdbRedisClusterResultOutput) Configs() GetMdbRedisClusterConfigArrayOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) []GetMdbRedisClusterConfig { return v.Configs }).(GetMdbRedisClusterConfigArrayOutput)
}

// Creation timestamp of the key.
func (o LookupMdbRedisClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupMdbRedisClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the Redis cluster.
func (o LookupMdbRedisClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

// Deployment environment of the Redis cluster.
func (o LookupMdbRedisClusterResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o LookupMdbRedisClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// Aggregated health of the cluster.
func (o LookupMdbRedisClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

// A host of the Redis cluster. The structure is documented below.
func (o LookupMdbRedisClusterResultOutput) Hosts() GetMdbRedisClusterHostArrayOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) []GetMdbRedisClusterHost { return v.Hosts }).(GetMdbRedisClusterHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMdbRedisClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to the Redis cluster.
func (o LookupMdbRedisClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupMdbRedisClusterResultOutput) MaintenanceWindows() GetMdbRedisClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) []GetMdbRedisClusterMaintenanceWindow { return v.MaintenanceWindows }).(GetMdbRedisClusterMaintenanceWindowArrayOutput)
}

func (o LookupMdbRedisClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the network, to which the Redis cluster belongs.
func (o LookupMdbRedisClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// Resources allocated to hosts of the Redis cluster. The structure is documented below.
func (o LookupMdbRedisClusterResultOutput) Resources() GetMdbRedisClusterResourceArrayOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) []GetMdbRedisClusterResource { return v.Resources }).(GetMdbRedisClusterResourceArrayOutput)
}

// A set of ids of security groups assigned to hosts of the cluster.
func (o LookupMdbRedisClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Redis Cluster mode enabled/disabled.
func (o LookupMdbRedisClusterResultOutput) Sharded() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) bool { return v.Sharded }).(pulumi.BoolOutput)
}

// Status of the cluster.
func (o LookupMdbRedisClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// tls support mode enabled/disabled.
func (o LookupMdbRedisClusterResultOutput) TlsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbRedisClusterResult) bool { return v.TlsEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMdbRedisClusterResultOutput{})
}
