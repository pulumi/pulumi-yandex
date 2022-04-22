// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Redis cluster within the Yandex.Cloud. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-redis/concepts).
//
// ## Example Usage
//
// Example of creating a Standalone Redis.
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
// 		fooVpcNetwork, err := yandex.NewVpcNetwork(ctx, "fooVpcNetwork", nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooVpcSubnet, err := yandex.NewVpcSubnet(ctx, "fooVpcSubnet", &yandex.VpcSubnetArgs{
// 			NetworkId: fooVpcNetwork.ID(),
// 			V4CidrBlocks: pulumi.StringArray{
// 				pulumi.String("10.5.0.0/24"),
// 			},
// 			Zone: pulumi.String("ru-central1-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewMdbRedisCluster(ctx, "fooMdbRedisCluster", &yandex.MdbRedisClusterArgs{
// 			Config: &MdbRedisClusterConfigArgs{
// 				Password: pulumi.String("your_password"),
// 				Version:  pulumi.String("6.0"),
// 			},
// 			Environment: pulumi.String("PRESTABLE"),
// 			Hosts: MdbRedisClusterHostArray{
// 				&MdbRedisClusterHostArgs{
// 					SubnetId: fooVpcSubnet.ID(),
// 					Zone:     pulumi.String("ru-central1-a"),
// 				},
// 			},
// 			MaintenanceWindow: &MdbRedisClusterMaintenanceWindowArgs{
// 				Type: pulumi.String("ANYTIME"),
// 			},
// 			NetworkId: fooVpcNetwork.ID(),
// 			Resources: &MdbRedisClusterResourcesArgs{
// 				DiskSize:         pulumi.Int(16),
// 				ResourcePresetId: pulumi.String("hm1.nano"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Example of creating a sharded Redis Cluster.
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
// 		fooVpcNetwork, err := yandex.NewVpcNetwork(ctx, "fooVpcNetwork", nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooVpcSubnet, err := yandex.NewVpcSubnet(ctx, "fooVpcSubnet", &yandex.VpcSubnetArgs{
// 			NetworkId: fooVpcNetwork.ID(),
// 			V4CidrBlocks: pulumi.StringArray{
// 				pulumi.String("10.1.0.0/24"),
// 			},
// 			Zone: pulumi.String("ru-central1-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := yandex.NewVpcSubnet(ctx, "bar", &yandex.VpcSubnetArgs{
// 			NetworkId: fooVpcNetwork.ID(),
// 			V4CidrBlocks: pulumi.StringArray{
// 				pulumi.String("10.2.0.0/24"),
// 			},
// 			Zone: pulumi.String("ru-central1-b"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		baz, err := yandex.NewVpcSubnet(ctx, "baz", &yandex.VpcSubnetArgs{
// 			NetworkId: fooVpcNetwork.ID(),
// 			V4CidrBlocks: pulumi.StringArray{
// 				pulumi.String("10.3.0.0/24"),
// 			},
// 			Zone: pulumi.String("ru-central1-c"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewMdbRedisCluster(ctx, "fooMdbRedisCluster", &yandex.MdbRedisClusterArgs{
// 			Config: &MdbRedisClusterConfigArgs{
// 				Password: pulumi.String("your_password"),
// 				Version:  pulumi.String("6.0"),
// 			},
// 			Environment: pulumi.String("PRESTABLE"),
// 			Hosts: MdbRedisClusterHostArray{
// 				&MdbRedisClusterHostArgs{
// 					ShardName: pulumi.String("first"),
// 					SubnetId:  fooVpcSubnet.ID(),
// 					Zone:      pulumi.String("ru-central1-a"),
// 				},
// 				&MdbRedisClusterHostArgs{
// 					ShardName: pulumi.String("second"),
// 					SubnetId:  bar.ID(),
// 					Zone:      pulumi.String("ru-central1-b"),
// 				},
// 				&MdbRedisClusterHostArgs{
// 					ShardName: pulumi.String("third"),
// 					SubnetId:  baz.ID(),
// 					Zone:      pulumi.String("ru-central1-c"),
// 				},
// 			},
// 			NetworkId: fooVpcNetwork.ID(),
// 			Resources: &MdbRedisClusterResourcesArgs{
// 				DiskSize:         pulumi.Int(16),
// 				ResourcePresetId: pulumi.String("hm1.nano"),
// 			},
// 			Sharded: pulumi.Bool(true),
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
// A cluster can be imported using the `id` of the resource, e.g.
//
// ```sh
//  $ pulumi import yandex:index/mdbRedisCluster:MdbRedisCluster foo cluster_id
// ```
type MdbRedisCluster struct {
	pulumi.CustomResourceState

	// Configuration of the Redis cluster. The structure is documented below.
	Config MdbRedisClusterConfigOutput `pulumi:"config"`
	// Creation timestamp of the key.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the Redis cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`.
	// For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
	Health pulumi.StringOutput `pulumi:"health"`
	// A host of the Redis cluster. The structure is documented below.
	Hosts MdbRedisClusterHostArrayOutput `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the Redis cluster.
	Labels            pulumi.StringMapOutput                 `pulumi:"labels"`
	MaintenanceWindow MdbRedisClusterMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// Name of the Redis cluster. Provided by the client when the cluster is created.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network, to which the Redis cluster belongs.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources MdbRedisClusterResourcesOutput `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Redis Cluster mode enabled/disabled.
	Sharded pulumi.BoolPtrOutput `pulumi:"sharded"`
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
	// For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
	Status pulumi.StringOutput `pulumi:"status"`
	// tls support mode enabled/disabled.
	TlsEnabled pulumi.BoolOutput `pulumi:"tlsEnabled"`
}

// NewMdbRedisCluster registers a new resource with the given unique name, arguments, and options.
func NewMdbRedisCluster(ctx *pulumi.Context,
	name string, args *MdbRedisClusterArgs, opts ...pulumi.ResourceOption) (*MdbRedisCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Hosts == nil {
		return nil, errors.New("invalid value for required argument 'Hosts'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.Resources == nil {
		return nil, errors.New("invalid value for required argument 'Resources'")
	}
	var resource MdbRedisCluster
	err := ctx.RegisterResource("yandex:index/mdbRedisCluster:MdbRedisCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbRedisCluster gets an existing MdbRedisCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbRedisCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbRedisClusterState, opts ...pulumi.ResourceOption) (*MdbRedisCluster, error) {
	var resource MdbRedisCluster
	err := ctx.ReadResource("yandex:index/mdbRedisCluster:MdbRedisCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbRedisCluster resources.
type mdbRedisClusterState struct {
	// Configuration of the Redis cluster. The structure is documented below.
	Config *MdbRedisClusterConfig `pulumi:"config"`
	// Creation timestamp of the key.
	CreatedAt *string `pulumi:"createdAt"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the Redis cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment *string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`.
	// For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
	Health *string `pulumi:"health"`
	// A host of the Redis cluster. The structure is documented below.
	Hosts []MdbRedisClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the Redis cluster.
	Labels            map[string]string                 `pulumi:"labels"`
	MaintenanceWindow *MdbRedisClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// Name of the Redis cluster. Provided by the client when the cluster is created.
	Name *string `pulumi:"name"`
	// ID of the network, to which the Redis cluster belongs.
	NetworkId *string `pulumi:"networkId"`
	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources *MdbRedisClusterResources `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Redis Cluster mode enabled/disabled.
	Sharded *bool `pulumi:"sharded"`
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
	// For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
	Status *string `pulumi:"status"`
	// tls support mode enabled/disabled.
	TlsEnabled *bool `pulumi:"tlsEnabled"`
}

type MdbRedisClusterState struct {
	// Configuration of the Redis cluster. The structure is documented below.
	Config MdbRedisClusterConfigPtrInput
	// Creation timestamp of the key.
	CreatedAt pulumi.StringPtrInput
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the Redis cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`.
	// For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
	Health pulumi.StringPtrInput
	// A host of the Redis cluster. The structure is documented below.
	Hosts MdbRedisClusterHostArrayInput
	// A set of key/value label pairs to assign to the Redis cluster.
	Labels            pulumi.StringMapInput
	MaintenanceWindow MdbRedisClusterMaintenanceWindowPtrInput
	// Name of the Redis cluster. Provided by the client when the cluster is created.
	Name pulumi.StringPtrInput
	// ID of the network, to which the Redis cluster belongs.
	NetworkId pulumi.StringPtrInput
	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources MdbRedisClusterResourcesPtrInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// Redis Cluster mode enabled/disabled.
	Sharded pulumi.BoolPtrInput
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
	// For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
	Status pulumi.StringPtrInput
	// tls support mode enabled/disabled.
	TlsEnabled pulumi.BoolPtrInput
}

func (MdbRedisClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbRedisClusterState)(nil)).Elem()
}

type mdbRedisClusterArgs struct {
	// Configuration of the Redis cluster. The structure is documented below.
	Config MdbRedisClusterConfig `pulumi:"config"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the Redis cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A host of the Redis cluster. The structure is documented below.
	Hosts []MdbRedisClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the Redis cluster.
	Labels            map[string]string                 `pulumi:"labels"`
	MaintenanceWindow *MdbRedisClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// Name of the Redis cluster. Provided by the client when the cluster is created.
	Name *string `pulumi:"name"`
	// ID of the network, to which the Redis cluster belongs.
	NetworkId string `pulumi:"networkId"`
	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources MdbRedisClusterResources `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Redis Cluster mode enabled/disabled.
	Sharded *bool `pulumi:"sharded"`
	// tls support mode enabled/disabled.
	TlsEnabled *bool `pulumi:"tlsEnabled"`
}

// The set of arguments for constructing a MdbRedisCluster resource.
type MdbRedisClusterArgs struct {
	// Configuration of the Redis cluster. The structure is documented below.
	Config MdbRedisClusterConfigInput
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the Redis cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A host of the Redis cluster. The structure is documented below.
	Hosts MdbRedisClusterHostArrayInput
	// A set of key/value label pairs to assign to the Redis cluster.
	Labels            pulumi.StringMapInput
	MaintenanceWindow MdbRedisClusterMaintenanceWindowPtrInput
	// Name of the Redis cluster. Provided by the client when the cluster is created.
	Name pulumi.StringPtrInput
	// ID of the network, to which the Redis cluster belongs.
	NetworkId pulumi.StringInput
	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources MdbRedisClusterResourcesInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// Redis Cluster mode enabled/disabled.
	Sharded pulumi.BoolPtrInput
	// tls support mode enabled/disabled.
	TlsEnabled pulumi.BoolPtrInput
}

func (MdbRedisClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbRedisClusterArgs)(nil)).Elem()
}

type MdbRedisClusterInput interface {
	pulumi.Input

	ToMdbRedisClusterOutput() MdbRedisClusterOutput
	ToMdbRedisClusterOutputWithContext(ctx context.Context) MdbRedisClusterOutput
}

func (*MdbRedisCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbRedisCluster)(nil)).Elem()
}

func (i *MdbRedisCluster) ToMdbRedisClusterOutput() MdbRedisClusterOutput {
	return i.ToMdbRedisClusterOutputWithContext(context.Background())
}

func (i *MdbRedisCluster) ToMdbRedisClusterOutputWithContext(ctx context.Context) MdbRedisClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbRedisClusterOutput)
}

// MdbRedisClusterArrayInput is an input type that accepts MdbRedisClusterArray and MdbRedisClusterArrayOutput values.
// You can construct a concrete instance of `MdbRedisClusterArrayInput` via:
//
//          MdbRedisClusterArray{ MdbRedisClusterArgs{...} }
type MdbRedisClusterArrayInput interface {
	pulumi.Input

	ToMdbRedisClusterArrayOutput() MdbRedisClusterArrayOutput
	ToMdbRedisClusterArrayOutputWithContext(context.Context) MdbRedisClusterArrayOutput
}

type MdbRedisClusterArray []MdbRedisClusterInput

func (MdbRedisClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbRedisCluster)(nil)).Elem()
}

func (i MdbRedisClusterArray) ToMdbRedisClusterArrayOutput() MdbRedisClusterArrayOutput {
	return i.ToMdbRedisClusterArrayOutputWithContext(context.Background())
}

func (i MdbRedisClusterArray) ToMdbRedisClusterArrayOutputWithContext(ctx context.Context) MdbRedisClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbRedisClusterArrayOutput)
}

// MdbRedisClusterMapInput is an input type that accepts MdbRedisClusterMap and MdbRedisClusterMapOutput values.
// You can construct a concrete instance of `MdbRedisClusterMapInput` via:
//
//          MdbRedisClusterMap{ "key": MdbRedisClusterArgs{...} }
type MdbRedisClusterMapInput interface {
	pulumi.Input

	ToMdbRedisClusterMapOutput() MdbRedisClusterMapOutput
	ToMdbRedisClusterMapOutputWithContext(context.Context) MdbRedisClusterMapOutput
}

type MdbRedisClusterMap map[string]MdbRedisClusterInput

func (MdbRedisClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbRedisCluster)(nil)).Elem()
}

func (i MdbRedisClusterMap) ToMdbRedisClusterMapOutput() MdbRedisClusterMapOutput {
	return i.ToMdbRedisClusterMapOutputWithContext(context.Background())
}

func (i MdbRedisClusterMap) ToMdbRedisClusterMapOutputWithContext(ctx context.Context) MdbRedisClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbRedisClusterMapOutput)
}

type MdbRedisClusterOutput struct{ *pulumi.OutputState }

func (MdbRedisClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbRedisCluster)(nil)).Elem()
}

func (o MdbRedisClusterOutput) ToMdbRedisClusterOutput() MdbRedisClusterOutput {
	return o
}

func (o MdbRedisClusterOutput) ToMdbRedisClusterOutputWithContext(ctx context.Context) MdbRedisClusterOutput {
	return o
}

type MdbRedisClusterArrayOutput struct{ *pulumi.OutputState }

func (MdbRedisClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbRedisCluster)(nil)).Elem()
}

func (o MdbRedisClusterArrayOutput) ToMdbRedisClusterArrayOutput() MdbRedisClusterArrayOutput {
	return o
}

func (o MdbRedisClusterArrayOutput) ToMdbRedisClusterArrayOutputWithContext(ctx context.Context) MdbRedisClusterArrayOutput {
	return o
}

func (o MdbRedisClusterArrayOutput) Index(i pulumi.IntInput) MdbRedisClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MdbRedisCluster {
		return vs[0].([]*MdbRedisCluster)[vs[1].(int)]
	}).(MdbRedisClusterOutput)
}

type MdbRedisClusterMapOutput struct{ *pulumi.OutputState }

func (MdbRedisClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbRedisCluster)(nil)).Elem()
}

func (o MdbRedisClusterMapOutput) ToMdbRedisClusterMapOutput() MdbRedisClusterMapOutput {
	return o
}

func (o MdbRedisClusterMapOutput) ToMdbRedisClusterMapOutputWithContext(ctx context.Context) MdbRedisClusterMapOutput {
	return o
}

func (o MdbRedisClusterMapOutput) MapIndex(k pulumi.StringInput) MdbRedisClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MdbRedisCluster {
		return vs[0].(map[string]*MdbRedisCluster)[vs[1].(string)]
	}).(MdbRedisClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MdbRedisClusterInput)(nil)).Elem(), &MdbRedisCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbRedisClusterArrayInput)(nil)).Elem(), MdbRedisClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbRedisClusterMapInput)(nil)).Elem(), MdbRedisClusterMap{})
	pulumi.RegisterOutputType(MdbRedisClusterOutput{})
	pulumi.RegisterOutputType(MdbRedisClusterArrayOutput{})
	pulumi.RegisterOutputType(MdbRedisClusterMapOutput{})
}
