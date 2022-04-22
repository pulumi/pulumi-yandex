// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A VM instance resource. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/compute/concepts/vm).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooVpcNetwork, err := yandex.NewVpcNetwork(ctx, "fooVpcNetwork", nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooVpcSubnet, err := yandex.NewVpcSubnet(ctx, "fooVpcSubnet", &yandex.VpcSubnetArgs{
// 			NetworkId: fooVpcNetwork.ID(),
// 			Zone:      pulumi.String("ru-central1-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewComputeInstance(ctx, "default", &yandex.ComputeInstanceArgs{
// 			BootDisk: &ComputeInstanceBootDiskArgs{
// 				InitializeParams: &ComputeInstanceBootDiskInitializeParamsArgs{
// 					ImageId: pulumi.String("image_id"),
// 				},
// 			},
// 			Metadata: pulumi.StringMap{
// 				"foo":      pulumi.String("bar"),
// 				"ssh-keys": pulumi.String(fmt.Sprintf("%v%v", "ubuntu:", readFileOrPanic("~/.ssh/id_rsa.pub"))),
// 			},
// 			NetworkInterfaces: ComputeInstanceNetworkInterfaceArray{
// 				&ComputeInstanceNetworkInterfaceArgs{
// 					SubnetId: fooVpcSubnet.ID(),
// 				},
// 			},
// 			PlatformId: pulumi.String("standard-v1"),
// 			Resources: &ComputeInstanceResourcesArgs{
// 				Cores:  pulumi.Int(2),
// 				Memory: pulumi.Float64(4),
// 			},
// 			Zone: pulumi.String("ru-central1-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Instances can be imported using the `ID` of an instance, e.g.
//
// ```sh
//  $ pulumi import yandex:index/computeInstance:ComputeInstance default instance_id
// ```
type ComputeInstance struct {
	pulumi.CustomResourceState

	AllowStoppingForUpdate pulumi.BoolPtrOutput `pulumi:"allowStoppingForUpdate"`
	// The boot disk for the instance. The structure is documented below.
	BootDisk ComputeInstanceBootDiskOutput `pulumi:"bootDisk"`
	// Creation timestamp of the instance.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the boot disk.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// DNS record FQDN (must have a dot at the end).
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Host name for the instance. This field is used to generate the instance `fqdn` value.
	// The host name must be unique within the network and region. If not specified, the host name will be equal
	// to `id` of the instance and `fqdn` will be `<id>.auto.internal`.
	// Otherwise FQDN will be `<hostname>.<region_id>.internal`.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// A set of key/value label pairs to assign to the instance.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Metadata key/value pairs to make available from
	// within the instance.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Name of the boot disk.
	Name pulumi.StringOutput `pulumi:"name"`
	// Type of network acceleration. The default is `standard`. Values: `standard`, `softwareAccelerated`
	NetworkAccelerationType pulumi.StringPtrOutput `pulumi:"networkAccelerationType"`
	// Networks to attach to the instance. This can
	// be specified multiple times. The structure is documented below.
	NetworkInterfaces ComputeInstanceNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// The placement policy configuration. The structure is documented below.
	PlacementPolicy ComputeInstancePlacementPolicyOutput `pulumi:"placementPolicy"`
	// The type of virtual machine to create. The default is 'standard-v1'.
	PlatformId pulumi.StringPtrOutput `pulumi:"platformId"`
	// Compute resources that are allocated for the instance. The structure is documented below.
	Resources ComputeInstanceResourcesOutput `pulumi:"resources"`
	// Scheduling policy configuration. The structure is documented below.
	SchedulingPolicy ComputeInstanceSchedulingPolicyOutput `pulumi:"schedulingPolicy"`
	// A list of disks to attach to the instance. The structure is documented below.
	// **Note**: The `allowStoppingForUpdate` property must be set to true in order to update this structure.
	SecondaryDisks ComputeInstanceSecondaryDiskArrayOutput `pulumi:"secondaryDisks"`
	// ID of the service account authorized for this instance.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
	// The status of this instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// The availability zone where the virtual machine will be created. If it is not provided,
	// the default provider folder is used.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewComputeInstance registers a new resource with the given unique name, arguments, and options.
func NewComputeInstance(ctx *pulumi.Context,
	name string, args *ComputeInstanceArgs, opts ...pulumi.ResourceOption) (*ComputeInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BootDisk == nil {
		return nil, errors.New("invalid value for required argument 'BootDisk'")
	}
	if args.NetworkInterfaces == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaces'")
	}
	if args.Resources == nil {
		return nil, errors.New("invalid value for required argument 'Resources'")
	}
	var resource ComputeInstance
	err := ctx.RegisterResource("yandex:index/computeInstance:ComputeInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeInstance gets an existing ComputeInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeInstanceState, opts ...pulumi.ResourceOption) (*ComputeInstance, error) {
	var resource ComputeInstance
	err := ctx.ReadResource("yandex:index/computeInstance:ComputeInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeInstance resources.
type computeInstanceState struct {
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	// The boot disk for the instance. The structure is documented below.
	BootDisk *ComputeInstanceBootDisk `pulumi:"bootDisk"`
	// Creation timestamp of the instance.
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the boot disk.
	Description *string `pulumi:"description"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// DNS record FQDN (must have a dot at the end).
	Fqdn *string `pulumi:"fqdn"`
	// Host name for the instance. This field is used to generate the instance `fqdn` value.
	// The host name must be unique within the network and region. If not specified, the host name will be equal
	// to `id` of the instance and `fqdn` will be `<id>.auto.internal`.
	// Otherwise FQDN will be `<hostname>.<region_id>.internal`.
	Hostname *string `pulumi:"hostname"`
	// A set of key/value label pairs to assign to the instance.
	Labels map[string]string `pulumi:"labels"`
	// Metadata key/value pairs to make available from
	// within the instance.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the boot disk.
	Name *string `pulumi:"name"`
	// Type of network acceleration. The default is `standard`. Values: `standard`, `softwareAccelerated`
	NetworkAccelerationType *string `pulumi:"networkAccelerationType"`
	// Networks to attach to the instance. This can
	// be specified multiple times. The structure is documented below.
	NetworkInterfaces []ComputeInstanceNetworkInterface `pulumi:"networkInterfaces"`
	// The placement policy configuration. The structure is documented below.
	PlacementPolicy *ComputeInstancePlacementPolicy `pulumi:"placementPolicy"`
	// The type of virtual machine to create. The default is 'standard-v1'.
	PlatformId *string `pulumi:"platformId"`
	// Compute resources that are allocated for the instance. The structure is documented below.
	Resources *ComputeInstanceResources `pulumi:"resources"`
	// Scheduling policy configuration. The structure is documented below.
	SchedulingPolicy *ComputeInstanceSchedulingPolicy `pulumi:"schedulingPolicy"`
	// A list of disks to attach to the instance. The structure is documented below.
	// **Note**: The `allowStoppingForUpdate` property must be set to true in order to update this structure.
	SecondaryDisks []ComputeInstanceSecondaryDisk `pulumi:"secondaryDisks"`
	// ID of the service account authorized for this instance.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// The status of this instance.
	Status *string `pulumi:"status"`
	// The availability zone where the virtual machine will be created. If it is not provided,
	// the default provider folder is used.
	Zone *string `pulumi:"zone"`
}

type ComputeInstanceState struct {
	AllowStoppingForUpdate pulumi.BoolPtrInput
	// The boot disk for the instance. The structure is documented below.
	BootDisk ComputeInstanceBootDiskPtrInput
	// Creation timestamp of the instance.
	CreatedAt pulumi.StringPtrInput
	// Description of the boot disk.
	Description pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// DNS record FQDN (must have a dot at the end).
	Fqdn pulumi.StringPtrInput
	// Host name for the instance. This field is used to generate the instance `fqdn` value.
	// The host name must be unique within the network and region. If not specified, the host name will be equal
	// to `id` of the instance and `fqdn` will be `<id>.auto.internal`.
	// Otherwise FQDN will be `<hostname>.<region_id>.internal`.
	Hostname pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the instance.
	Labels pulumi.StringMapInput
	// Metadata key/value pairs to make available from
	// within the instance.
	Metadata pulumi.StringMapInput
	// Name of the boot disk.
	Name pulumi.StringPtrInput
	// Type of network acceleration. The default is `standard`. Values: `standard`, `softwareAccelerated`
	NetworkAccelerationType pulumi.StringPtrInput
	// Networks to attach to the instance. This can
	// be specified multiple times. The structure is documented below.
	NetworkInterfaces ComputeInstanceNetworkInterfaceArrayInput
	// The placement policy configuration. The structure is documented below.
	PlacementPolicy ComputeInstancePlacementPolicyPtrInput
	// The type of virtual machine to create. The default is 'standard-v1'.
	PlatformId pulumi.StringPtrInput
	// Compute resources that are allocated for the instance. The structure is documented below.
	Resources ComputeInstanceResourcesPtrInput
	// Scheduling policy configuration. The structure is documented below.
	SchedulingPolicy ComputeInstanceSchedulingPolicyPtrInput
	// A list of disks to attach to the instance. The structure is documented below.
	// **Note**: The `allowStoppingForUpdate` property must be set to true in order to update this structure.
	SecondaryDisks ComputeInstanceSecondaryDiskArrayInput
	// ID of the service account authorized for this instance.
	ServiceAccountId pulumi.StringPtrInput
	// The status of this instance.
	Status pulumi.StringPtrInput
	// The availability zone where the virtual machine will be created. If it is not provided,
	// the default provider folder is used.
	Zone pulumi.StringPtrInput
}

func (ComputeInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeInstanceState)(nil)).Elem()
}

type computeInstanceArgs struct {
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	// The boot disk for the instance. The structure is documented below.
	BootDisk ComputeInstanceBootDisk `pulumi:"bootDisk"`
	// Description of the boot disk.
	Description *string `pulumi:"description"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Host name for the instance. This field is used to generate the instance `fqdn` value.
	// The host name must be unique within the network and region. If not specified, the host name will be equal
	// to `id` of the instance and `fqdn` will be `<id>.auto.internal`.
	// Otherwise FQDN will be `<hostname>.<region_id>.internal`.
	Hostname *string `pulumi:"hostname"`
	// A set of key/value label pairs to assign to the instance.
	Labels map[string]string `pulumi:"labels"`
	// Metadata key/value pairs to make available from
	// within the instance.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the boot disk.
	Name *string `pulumi:"name"`
	// Type of network acceleration. The default is `standard`. Values: `standard`, `softwareAccelerated`
	NetworkAccelerationType *string `pulumi:"networkAccelerationType"`
	// Networks to attach to the instance. This can
	// be specified multiple times. The structure is documented below.
	NetworkInterfaces []ComputeInstanceNetworkInterface `pulumi:"networkInterfaces"`
	// The placement policy configuration. The structure is documented below.
	PlacementPolicy *ComputeInstancePlacementPolicy `pulumi:"placementPolicy"`
	// The type of virtual machine to create. The default is 'standard-v1'.
	PlatformId *string `pulumi:"platformId"`
	// Compute resources that are allocated for the instance. The structure is documented below.
	Resources ComputeInstanceResources `pulumi:"resources"`
	// Scheduling policy configuration. The structure is documented below.
	SchedulingPolicy *ComputeInstanceSchedulingPolicy `pulumi:"schedulingPolicy"`
	// A list of disks to attach to the instance. The structure is documented below.
	// **Note**: The `allowStoppingForUpdate` property must be set to true in order to update this structure.
	SecondaryDisks []ComputeInstanceSecondaryDisk `pulumi:"secondaryDisks"`
	// ID of the service account authorized for this instance.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// The availability zone where the virtual machine will be created. If it is not provided,
	// the default provider folder is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ComputeInstance resource.
type ComputeInstanceArgs struct {
	AllowStoppingForUpdate pulumi.BoolPtrInput
	// The boot disk for the instance. The structure is documented below.
	BootDisk ComputeInstanceBootDiskInput
	// Description of the boot disk.
	Description pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Host name for the instance. This field is used to generate the instance `fqdn` value.
	// The host name must be unique within the network and region. If not specified, the host name will be equal
	// to `id` of the instance and `fqdn` will be `<id>.auto.internal`.
	// Otherwise FQDN will be `<hostname>.<region_id>.internal`.
	Hostname pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the instance.
	Labels pulumi.StringMapInput
	// Metadata key/value pairs to make available from
	// within the instance.
	Metadata pulumi.StringMapInput
	// Name of the boot disk.
	Name pulumi.StringPtrInput
	// Type of network acceleration. The default is `standard`. Values: `standard`, `softwareAccelerated`
	NetworkAccelerationType pulumi.StringPtrInput
	// Networks to attach to the instance. This can
	// be specified multiple times. The structure is documented below.
	NetworkInterfaces ComputeInstanceNetworkInterfaceArrayInput
	// The placement policy configuration. The structure is documented below.
	PlacementPolicy ComputeInstancePlacementPolicyPtrInput
	// The type of virtual machine to create. The default is 'standard-v1'.
	PlatformId pulumi.StringPtrInput
	// Compute resources that are allocated for the instance. The structure is documented below.
	Resources ComputeInstanceResourcesInput
	// Scheduling policy configuration. The structure is documented below.
	SchedulingPolicy ComputeInstanceSchedulingPolicyPtrInput
	// A list of disks to attach to the instance. The structure is documented below.
	// **Note**: The `allowStoppingForUpdate` property must be set to true in order to update this structure.
	SecondaryDisks ComputeInstanceSecondaryDiskArrayInput
	// ID of the service account authorized for this instance.
	ServiceAccountId pulumi.StringPtrInput
	// The availability zone where the virtual machine will be created. If it is not provided,
	// the default provider folder is used.
	Zone pulumi.StringPtrInput
}

func (ComputeInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeInstanceArgs)(nil)).Elem()
}

type ComputeInstanceInput interface {
	pulumi.Input

	ToComputeInstanceOutput() ComputeInstanceOutput
	ToComputeInstanceOutputWithContext(ctx context.Context) ComputeInstanceOutput
}

func (*ComputeInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstance)(nil)).Elem()
}

func (i *ComputeInstance) ToComputeInstanceOutput() ComputeInstanceOutput {
	return i.ToComputeInstanceOutputWithContext(context.Background())
}

func (i *ComputeInstance) ToComputeInstanceOutputWithContext(ctx context.Context) ComputeInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceOutput)
}

// ComputeInstanceArrayInput is an input type that accepts ComputeInstanceArray and ComputeInstanceArrayOutput values.
// You can construct a concrete instance of `ComputeInstanceArrayInput` via:
//
//          ComputeInstanceArray{ ComputeInstanceArgs{...} }
type ComputeInstanceArrayInput interface {
	pulumi.Input

	ToComputeInstanceArrayOutput() ComputeInstanceArrayOutput
	ToComputeInstanceArrayOutputWithContext(context.Context) ComputeInstanceArrayOutput
}

type ComputeInstanceArray []ComputeInstanceInput

func (ComputeInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeInstance)(nil)).Elem()
}

func (i ComputeInstanceArray) ToComputeInstanceArrayOutput() ComputeInstanceArrayOutput {
	return i.ToComputeInstanceArrayOutputWithContext(context.Background())
}

func (i ComputeInstanceArray) ToComputeInstanceArrayOutputWithContext(ctx context.Context) ComputeInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceArrayOutput)
}

// ComputeInstanceMapInput is an input type that accepts ComputeInstanceMap and ComputeInstanceMapOutput values.
// You can construct a concrete instance of `ComputeInstanceMapInput` via:
//
//          ComputeInstanceMap{ "key": ComputeInstanceArgs{...} }
type ComputeInstanceMapInput interface {
	pulumi.Input

	ToComputeInstanceMapOutput() ComputeInstanceMapOutput
	ToComputeInstanceMapOutputWithContext(context.Context) ComputeInstanceMapOutput
}

type ComputeInstanceMap map[string]ComputeInstanceInput

func (ComputeInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeInstance)(nil)).Elem()
}

func (i ComputeInstanceMap) ToComputeInstanceMapOutput() ComputeInstanceMapOutput {
	return i.ToComputeInstanceMapOutputWithContext(context.Background())
}

func (i ComputeInstanceMap) ToComputeInstanceMapOutputWithContext(ctx context.Context) ComputeInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceMapOutput)
}

type ComputeInstanceOutput struct{ *pulumi.OutputState }

func (ComputeInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstance)(nil)).Elem()
}

func (o ComputeInstanceOutput) ToComputeInstanceOutput() ComputeInstanceOutput {
	return o
}

func (o ComputeInstanceOutput) ToComputeInstanceOutputWithContext(ctx context.Context) ComputeInstanceOutput {
	return o
}

type ComputeInstanceArrayOutput struct{ *pulumi.OutputState }

func (ComputeInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeInstance)(nil)).Elem()
}

func (o ComputeInstanceArrayOutput) ToComputeInstanceArrayOutput() ComputeInstanceArrayOutput {
	return o
}

func (o ComputeInstanceArrayOutput) ToComputeInstanceArrayOutputWithContext(ctx context.Context) ComputeInstanceArrayOutput {
	return o
}

func (o ComputeInstanceArrayOutput) Index(i pulumi.IntInput) ComputeInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ComputeInstance {
		return vs[0].([]*ComputeInstance)[vs[1].(int)]
	}).(ComputeInstanceOutput)
}

type ComputeInstanceMapOutput struct{ *pulumi.OutputState }

func (ComputeInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeInstance)(nil)).Elem()
}

func (o ComputeInstanceMapOutput) ToComputeInstanceMapOutput() ComputeInstanceMapOutput {
	return o
}

func (o ComputeInstanceMapOutput) ToComputeInstanceMapOutputWithContext(ctx context.Context) ComputeInstanceMapOutput {
	return o
}

func (o ComputeInstanceMapOutput) MapIndex(k pulumi.StringInput) ComputeInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ComputeInstance {
		return vs[0].(map[string]*ComputeInstance)[vs[1].(string)]
	}).(ComputeInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceInput)(nil)).Elem(), &ComputeInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceArrayInput)(nil)).Elem(), ComputeInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceMapInput)(nil)).Elem(), ComputeInstanceMap{})
	pulumi.RegisterOutputType(ComputeInstanceOutput{})
	pulumi.RegisterOutputType(ComputeInstanceArrayOutput{})
	pulumi.RegisterOutputType(ComputeInstanceMapOutput{})
}
