// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Compute instance group.
func LookupComputeInstanceGroup(ctx *pulumi.Context, args *LookupComputeInstanceGroupArgs, opts ...pulumi.InvokeOption) (*LookupComputeInstanceGroupResult, error) {
	var rv LookupComputeInstanceGroupResult
	err := ctx.Invoke("yandex:index/getComputeInstanceGroup:getComputeInstanceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeInstanceGroup.
type LookupComputeInstanceGroupArgs struct {
	// The ID of a specific instance group.
	InstanceGroupId string `pulumi:"instanceGroupId"`
}

// A collection of values returned by getComputeInstanceGroup.
type LookupComputeInstanceGroupResult struct {
	// The allocation policy of the instance group by zone and region. The structure is documented below.
	AllocationPolicy GetComputeInstanceGroupAllocationPolicy `pulumi:"allocationPolicy"`
	// The instance group creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// Flag that protects the instance group from accidental deletion.
	DeletionProtection bool `pulumi:"deletionProtection"`
	// The deployment policy of the instance group. The structure is documented below.
	DeployPolicy GetComputeInstanceGroupDeployPolicy `pulumi:"deployPolicy"`
	// A description of the boot disk.
	Description string `pulumi:"description"`
	// The ID of the folder that the instance group belongs to.
	FolderId string `pulumi:"folderId"`
	// Health check specification. The structure is documented below.
	HealthChecks []GetComputeInstanceGroupHealthCheck `pulumi:"healthChecks"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	InstanceGroupId string `pulumi:"instanceGroupId"`
	// The instance template that the instance group belongs to. The structure is documented below.
	InstanceTemplate GetComputeInstanceGroupInstanceTemplate `pulumi:"instanceTemplate"`
	// A list of instances in the specified instance group. The structure is documented below.
	Instances []GetComputeInstanceGroupInstance `pulumi:"instances"`
	// A map of labels applied to this instance.
	// * `resources.0.memory` - The memory size allocated to the instance.
	// * `resources.0.cores` - Number of CPU cores allocated to the instance.
	// * `resources.0.core_fraction` - Baseline core performance as a percent.
	// * `resources.0.gpus` - Number of GPU cores allocated to the instance.
	Labels map[string]string `pulumi:"labels"`
	// Load balancing specification. The structure is documented below.
	LoadBalancer GetComputeInstanceGroupLoadBalancer `pulumi:"loadBalancer"`
	// Information about which entities can be attached to this load balancer. The structure is documented below.
	LoadBalancerState GetComputeInstanceGroupLoadBalancerState `pulumi:"loadBalancerState"`
	// The name of the managed instance.
	Name string `pulumi:"name"`
	// The scaling policy of the instance group. The structure is documented below.
	ScalePolicy GetComputeInstanceGroupScalePolicy `pulumi:"scalePolicy"`
	// The service account ID for the instance.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// The status of the instance.
	Status string `pulumi:"status"`
	// A set of key/value  variables pairs to assign to the instance group.
	Variables map[string]string `pulumi:"variables"`
}
