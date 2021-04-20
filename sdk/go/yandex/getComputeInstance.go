// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Compute instance. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/compute/concepts/vm).
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
// 		opt0 := "some_instance_id"
// 		myInstance, err := yandex.LookupComputeInstance(ctx, &yandex.LookupComputeInstanceArgs{
// 			InstanceId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("instanceExternalIp", myInstance.NetworkInterfaces[0].NatIpAddress)
// 		return nil
// 	})
// }
// ```
func LookupComputeInstance(ctx *pulumi.Context, args *LookupComputeInstanceArgs, opts ...pulumi.InvokeOption) (*LookupComputeInstanceResult, error) {
	var rv LookupComputeInstanceResult
	err := ctx.Invoke("yandex:index/getComputeInstance:getComputeInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeInstance.
type LookupComputeInstanceArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// The ID of a specific instance.
	InstanceId *string `pulumi:"instanceId"`
	// Name of the instance.
	Name            *string                            `pulumi:"name"`
	PlacementPolicy *GetComputeInstancePlacementPolicy `pulumi:"placementPolicy"`
	// Scheduling policy configuration. The structure is documented below.
	SchedulingPolicy *GetComputeInstanceSchedulingPolicy `pulumi:"schedulingPolicy"`
}

// A collection of values returned by getComputeInstance.
type LookupComputeInstanceResult struct {
	// The boot disk for the instance. Structure is documented below.
	BootDisk GetComputeInstanceBootDisk `pulumi:"bootDisk"`
	// Instance creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// Description of the boot disk.
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	// FQDN of the instance.
	Fqdn string `pulumi:"fqdn"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// A set of key/value label pairs assigned to the instance.
	Labels map[string]string `pulumi:"labels"`
	// Metadata key/value pairs to make available from
	// within the instance.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the boot disk.
	Name string `pulumi:"name"`
	// Type of network acceleration. The default is `standard`. Values: `standard`, `softwareAccelerated`
	NetworkAccelerationType string `pulumi:"networkAccelerationType"`
	// The networks attached to the instance. Structure is documented below.
	// * `network_interface.0.ip_address` - An internal IP address of the instance, either manually or dynamically assigned.
	// * `network_interface.0.nat_ip_address` - An assigned external IP address if the instance has NAT enabled.
	NetworkInterfaces []GetComputeInstanceNetworkInterface `pulumi:"networkInterfaces"`
	PlacementPolicy   *GetComputeInstancePlacementPolicy   `pulumi:"placementPolicy"`
	// Type of virtual machine to create. Default is 'standard-v1'.
	PlatformId string                      `pulumi:"platformId"`
	Resources  GetComputeInstanceResources `pulumi:"resources"`
	// Scheduling policy configuration. The structure is documented below.
	SchedulingPolicy GetComputeInstanceSchedulingPolicy `pulumi:"schedulingPolicy"`
	// List of secondary disks attached to the instance. Structure is documented below.
	SecondaryDisks []GetComputeInstanceSecondaryDisk `pulumi:"secondaryDisks"`
	// ID of the service account authorized for this instance.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// Status of the instance.
	// * `resources.0.memory` - Memory size allocated for the instance.
	// * `resources.0.cores` - Number of CPU cores allocated for the instance.
	// * `resources.0.core_fraction` - Baseline performance for a core, set as a percent.
	// * `resources.0.gpus` - Number of GPU cores allocated for the instance.
	Status string `pulumi:"status"`
	// Availability zone where the instance resides.
	Zone string `pulumi:"zone"`
}
