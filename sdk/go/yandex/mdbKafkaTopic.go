// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a topic of a Kafka cluster within the Yandex.Cloud. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts).
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
// 		foo, err := yandex.NewMdbKafkaCluster(ctx, "foo", &yandex.MdbKafkaClusterArgs{
// 			NetworkId: pulumi.String("c64vs98keiqc7f24pvkd"),
// 			Config: &MdbKafkaClusterConfigArgs{
// 				Version: pulumi.String("2.8"),
// 				Zones: pulumi.StringArray{
// 					pulumi.String("ru-central1-a"),
// 				},
// 				UnmanagedTopics: pulumi.Bool(true),
// 				Kafka: &MdbKafkaClusterConfigKafkaArgs{
// 					Resources: &MdbKafkaClusterConfigKafkaResourcesArgs{
// 						ResourcePresetId: pulumi.String("s2.micro"),
// 						DiskTypeId:       pulumi.String("network-hdd"),
// 						DiskSize:         pulumi.Int(16),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewMdbKafkaTopic(ctx, "events", &yandex.MdbKafkaTopicArgs{
// 			ClusterId:         foo.ID(),
// 			Partitions:        pulumi.Int(4),
// 			ReplicationFactor: pulumi.Int(1),
// 			TopicConfig: &MdbKafkaTopicTopicConfigArgs{
// 				CleanupPolicy:      pulumi.String("CLEANUP_POLICY_COMPACT"),
// 				CompressionType:    pulumi.String("COMPRESSION_TYPE_LZ4"),
// 				DeleteRetentionMs:  pulumi.String("86400000"),
// 				FileDeleteDelayMs:  pulumi.String("60000"),
// 				FlushMessages:      pulumi.String("128"),
// 				FlushMs:            pulumi.String("1000"),
// 				MinCompactionLagMs: pulumi.String("0"),
// 				RetentionBytes:     pulumi.String("10737418240"),
// 				RetentionMs:        pulumi.String("604800000"),
// 				MaxMessageBytes:    pulumi.String("1048588"),
// 				MinInsyncReplicas:  pulumi.String("1"),
// 				SegmentBytes:       pulumi.String("268435456"),
// 				Preallocate:        pulumi.Bool(true),
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
// ## Import
//
// Kafka topic can be imported using following format
//
// ```sh
//  $ pulumi import yandex:index/mdbKafkaTopic:MdbKafkaTopic foo {{cluster_id}}:{{topic_name}}
// ```
type MdbKafkaTopic struct {
	pulumi.CustomResourceState

	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The name of the topic.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of the topic's partitions.
	Partitions pulumi.IntOutput `pulumi:"partitions"`
	// Amount of data copies (replicas) for the topic in the cluster.
	ReplicationFactor pulumi.IntOutput `pulumi:"replicationFactor"`
	// User-defined settings for the topic. The structure is documented below.
	TopicConfig MdbKafkaTopicTopicConfigPtrOutput `pulumi:"topicConfig"`
}

// NewMdbKafkaTopic registers a new resource with the given unique name, arguments, and options.
func NewMdbKafkaTopic(ctx *pulumi.Context,
	name string, args *MdbKafkaTopicArgs, opts ...pulumi.ResourceOption) (*MdbKafkaTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Partitions == nil {
		return nil, errors.New("invalid value for required argument 'Partitions'")
	}
	if args.ReplicationFactor == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationFactor'")
	}
	var resource MdbKafkaTopic
	err := ctx.RegisterResource("yandex:index/mdbKafkaTopic:MdbKafkaTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbKafkaTopic gets an existing MdbKafkaTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbKafkaTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbKafkaTopicState, opts ...pulumi.ResourceOption) (*MdbKafkaTopic, error) {
	var resource MdbKafkaTopic
	err := ctx.ReadResource("yandex:index/mdbKafkaTopic:MdbKafkaTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbKafkaTopic resources.
type mdbKafkaTopicState struct {
	ClusterId *string `pulumi:"clusterId"`
	// The name of the topic.
	Name *string `pulumi:"name"`
	// The number of the topic's partitions.
	Partitions *int `pulumi:"partitions"`
	// Amount of data copies (replicas) for the topic in the cluster.
	ReplicationFactor *int `pulumi:"replicationFactor"`
	// User-defined settings for the topic. The structure is documented below.
	TopicConfig *MdbKafkaTopicTopicConfig `pulumi:"topicConfig"`
}

type MdbKafkaTopicState struct {
	ClusterId pulumi.StringPtrInput
	// The name of the topic.
	Name pulumi.StringPtrInput
	// The number of the topic's partitions.
	Partitions pulumi.IntPtrInput
	// Amount of data copies (replicas) for the topic in the cluster.
	ReplicationFactor pulumi.IntPtrInput
	// User-defined settings for the topic. The structure is documented below.
	TopicConfig MdbKafkaTopicTopicConfigPtrInput
}

func (MdbKafkaTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbKafkaTopicState)(nil)).Elem()
}

type mdbKafkaTopicArgs struct {
	ClusterId string `pulumi:"clusterId"`
	// The name of the topic.
	Name *string `pulumi:"name"`
	// The number of the topic's partitions.
	Partitions int `pulumi:"partitions"`
	// Amount of data copies (replicas) for the topic in the cluster.
	ReplicationFactor int `pulumi:"replicationFactor"`
	// User-defined settings for the topic. The structure is documented below.
	TopicConfig *MdbKafkaTopicTopicConfig `pulumi:"topicConfig"`
}

// The set of arguments for constructing a MdbKafkaTopic resource.
type MdbKafkaTopicArgs struct {
	ClusterId pulumi.StringInput
	// The name of the topic.
	Name pulumi.StringPtrInput
	// The number of the topic's partitions.
	Partitions pulumi.IntInput
	// Amount of data copies (replicas) for the topic in the cluster.
	ReplicationFactor pulumi.IntInput
	// User-defined settings for the topic. The structure is documented below.
	TopicConfig MdbKafkaTopicTopicConfigPtrInput
}

func (MdbKafkaTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbKafkaTopicArgs)(nil)).Elem()
}

type MdbKafkaTopicInput interface {
	pulumi.Input

	ToMdbKafkaTopicOutput() MdbKafkaTopicOutput
	ToMdbKafkaTopicOutputWithContext(ctx context.Context) MdbKafkaTopicOutput
}

func (*MdbKafkaTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbKafkaTopic)(nil)).Elem()
}

func (i *MdbKafkaTopic) ToMdbKafkaTopicOutput() MdbKafkaTopicOutput {
	return i.ToMdbKafkaTopicOutputWithContext(context.Background())
}

func (i *MdbKafkaTopic) ToMdbKafkaTopicOutputWithContext(ctx context.Context) MdbKafkaTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbKafkaTopicOutput)
}

// MdbKafkaTopicArrayInput is an input type that accepts MdbKafkaTopicArray and MdbKafkaTopicArrayOutput values.
// You can construct a concrete instance of `MdbKafkaTopicArrayInput` via:
//
//          MdbKafkaTopicArray{ MdbKafkaTopicArgs{...} }
type MdbKafkaTopicArrayInput interface {
	pulumi.Input

	ToMdbKafkaTopicArrayOutput() MdbKafkaTopicArrayOutput
	ToMdbKafkaTopicArrayOutputWithContext(context.Context) MdbKafkaTopicArrayOutput
}

type MdbKafkaTopicArray []MdbKafkaTopicInput

func (MdbKafkaTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbKafkaTopic)(nil)).Elem()
}

func (i MdbKafkaTopicArray) ToMdbKafkaTopicArrayOutput() MdbKafkaTopicArrayOutput {
	return i.ToMdbKafkaTopicArrayOutputWithContext(context.Background())
}

func (i MdbKafkaTopicArray) ToMdbKafkaTopicArrayOutputWithContext(ctx context.Context) MdbKafkaTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbKafkaTopicArrayOutput)
}

// MdbKafkaTopicMapInput is an input type that accepts MdbKafkaTopicMap and MdbKafkaTopicMapOutput values.
// You can construct a concrete instance of `MdbKafkaTopicMapInput` via:
//
//          MdbKafkaTopicMap{ "key": MdbKafkaTopicArgs{...} }
type MdbKafkaTopicMapInput interface {
	pulumi.Input

	ToMdbKafkaTopicMapOutput() MdbKafkaTopicMapOutput
	ToMdbKafkaTopicMapOutputWithContext(context.Context) MdbKafkaTopicMapOutput
}

type MdbKafkaTopicMap map[string]MdbKafkaTopicInput

func (MdbKafkaTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbKafkaTopic)(nil)).Elem()
}

func (i MdbKafkaTopicMap) ToMdbKafkaTopicMapOutput() MdbKafkaTopicMapOutput {
	return i.ToMdbKafkaTopicMapOutputWithContext(context.Background())
}

func (i MdbKafkaTopicMap) ToMdbKafkaTopicMapOutputWithContext(ctx context.Context) MdbKafkaTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbKafkaTopicMapOutput)
}

type MdbKafkaTopicOutput struct{ *pulumi.OutputState }

func (MdbKafkaTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbKafkaTopic)(nil)).Elem()
}

func (o MdbKafkaTopicOutput) ToMdbKafkaTopicOutput() MdbKafkaTopicOutput {
	return o
}

func (o MdbKafkaTopicOutput) ToMdbKafkaTopicOutputWithContext(ctx context.Context) MdbKafkaTopicOutput {
	return o
}

type MdbKafkaTopicArrayOutput struct{ *pulumi.OutputState }

func (MdbKafkaTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbKafkaTopic)(nil)).Elem()
}

func (o MdbKafkaTopicArrayOutput) ToMdbKafkaTopicArrayOutput() MdbKafkaTopicArrayOutput {
	return o
}

func (o MdbKafkaTopicArrayOutput) ToMdbKafkaTopicArrayOutputWithContext(ctx context.Context) MdbKafkaTopicArrayOutput {
	return o
}

func (o MdbKafkaTopicArrayOutput) Index(i pulumi.IntInput) MdbKafkaTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MdbKafkaTopic {
		return vs[0].([]*MdbKafkaTopic)[vs[1].(int)]
	}).(MdbKafkaTopicOutput)
}

type MdbKafkaTopicMapOutput struct{ *pulumi.OutputState }

func (MdbKafkaTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbKafkaTopic)(nil)).Elem()
}

func (o MdbKafkaTopicMapOutput) ToMdbKafkaTopicMapOutput() MdbKafkaTopicMapOutput {
	return o
}

func (o MdbKafkaTopicMapOutput) ToMdbKafkaTopicMapOutputWithContext(ctx context.Context) MdbKafkaTopicMapOutput {
	return o
}

func (o MdbKafkaTopicMapOutput) MapIndex(k pulumi.StringInput) MdbKafkaTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MdbKafkaTopic {
		return vs[0].(map[string]*MdbKafkaTopic)[vs[1].(string)]
	}).(MdbKafkaTopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MdbKafkaTopicInput)(nil)).Elem(), &MdbKafkaTopic{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbKafkaTopicArrayInput)(nil)).Elem(), MdbKafkaTopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbKafkaTopicMapInput)(nil)).Elem(), MdbKafkaTopicMap{})
	pulumi.RegisterOutputType(MdbKafkaTopicOutput{})
	pulumi.RegisterOutputType(MdbKafkaTopicArrayOutput{})
	pulumi.RegisterOutputType(MdbKafkaTopicMapOutput{})
}
