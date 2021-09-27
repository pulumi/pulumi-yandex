// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Greenplum cluster within the Yandex.Cloud. For more information, see
// [the official documentation](https://cloud.yandex.ru/docs/managed-greenplum/).
//
// Please read [Pricing for Managed Service for Greenplum](https://cloud.yandex.ru/docs/managed-greenplum/) before using Greenplum cluster.
//
// Yandex Managed Service for Greenplum® is now in preview
//
// ## Import
//
// A cluster can be imported using the `id` of the resource, e.g.
//
// ```sh
//  $ pulumi import yandex:index/mdbGreenplumCluster:MdbGreenplumCluster foo cluster_id
// ```
type MdbGreenplumCluster struct {
	pulumi.CustomResourceState

	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment.
	AssignPublicIp pulumi.BoolOutput `pulumi:"assignPublicIp"`
	// Creation timestamp of the cluster.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the Greenplum cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health pulumi.StringOutput `pulumi:"health"`
	// A set of key/value label pairs to assign to the Greenplum cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount pulumi.IntOutput `pulumi:"masterHostCount"`
	// (Computed) Info about hosts in master subcluster. The structure is documented below.
	MasterHosts MdbGreenplumClusterMasterHostArrayOutput `pulumi:"masterHosts"`
	// Settings for master subcluster. The structure is documented below.
	MasterSubcluster MdbGreenplumClusterMasterSubclusterOutput `pulumi:"masterSubcluster"`
	// Name of the Greenplum cluster. Provided by the client when the cluster is created.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network, to which the Greenplum cluster uses.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount pulumi.IntOutput `pulumi:"segmentHostCount"`
	// (Computed) Info about hosts in segment subcluster. The structure is documented below.
	SegmentHosts MdbGreenplumClusterSegmentHostArrayOutput `pulumi:"segmentHosts"`
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost pulumi.IntOutput `pulumi:"segmentInHost"`
	// Settings for segment subcluster. The structure is documented below.
	SegmentSubcluster MdbGreenplumClusterSegmentSubclusterOutput `pulumi:"segmentSubcluster"`
	// Status of the cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Greenplum cluster admin user name.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Greenplum cluster admin password name.
	UserPassword pulumi.StringOutput `pulumi:"userPassword"`
	// Version of the Greenplum cluster. (6.17)
	Version pulumi.StringOutput `pulumi:"version"`
	// The availability zone where the Greenplum hosts will be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewMdbGreenplumCluster registers a new resource with the given unique name, arguments, and options.
func NewMdbGreenplumCluster(ctx *pulumi.Context,
	name string, args *MdbGreenplumClusterArgs, opts ...pulumi.ResourceOption) (*MdbGreenplumCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssignPublicIp == nil {
		return nil, errors.New("invalid value for required argument 'AssignPublicIp'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.MasterHostCount == nil {
		return nil, errors.New("invalid value for required argument 'MasterHostCount'")
	}
	if args.MasterSubcluster == nil {
		return nil, errors.New("invalid value for required argument 'MasterSubcluster'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.SegmentHostCount == nil {
		return nil, errors.New("invalid value for required argument 'SegmentHostCount'")
	}
	if args.SegmentInHost == nil {
		return nil, errors.New("invalid value for required argument 'SegmentInHost'")
	}
	if args.SegmentSubcluster == nil {
		return nil, errors.New("invalid value for required argument 'SegmentSubcluster'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	if args.UserPassword == nil {
		return nil, errors.New("invalid value for required argument 'UserPassword'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource MdbGreenplumCluster
	err := ctx.RegisterResource("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbGreenplumCluster gets an existing MdbGreenplumCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbGreenplumCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbGreenplumClusterState, opts ...pulumi.ResourceOption) (*MdbGreenplumCluster, error) {
	var resource MdbGreenplumCluster
	err := ctx.ReadResource("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbGreenplumCluster resources.
type mdbGreenplumClusterState struct {
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment.
	AssignPublicIp *bool `pulumi:"assignPublicIp"`
	// Creation timestamp of the cluster.
	CreatedAt *string `pulumi:"createdAt"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the Greenplum cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment *string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health *string `pulumi:"health"`
	// A set of key/value label pairs to assign to the Greenplum cluster.
	Labels map[string]string `pulumi:"labels"`
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount *int `pulumi:"masterHostCount"`
	// (Computed) Info about hosts in master subcluster. The structure is documented below.
	MasterHosts []MdbGreenplumClusterMasterHost `pulumi:"masterHosts"`
	// Settings for master subcluster. The structure is documented below.
	MasterSubcluster *MdbGreenplumClusterMasterSubcluster `pulumi:"masterSubcluster"`
	// Name of the Greenplum cluster. Provided by the client when the cluster is created.
	Name *string `pulumi:"name"`
	// ID of the network, to which the Greenplum cluster uses.
	NetworkId *string `pulumi:"networkId"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount *int `pulumi:"segmentHostCount"`
	// (Computed) Info about hosts in segment subcluster. The structure is documented below.
	SegmentHosts []MdbGreenplumClusterSegmentHost `pulumi:"segmentHosts"`
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost *int `pulumi:"segmentInHost"`
	// Settings for segment subcluster. The structure is documented below.
	SegmentSubcluster *MdbGreenplumClusterSegmentSubcluster `pulumi:"segmentSubcluster"`
	// Status of the cluster.
	Status *string `pulumi:"status"`
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId *string `pulumi:"subnetId"`
	// Greenplum cluster admin user name.
	UserName *string `pulumi:"userName"`
	// Greenplum cluster admin password name.
	UserPassword *string `pulumi:"userPassword"`
	// Version of the Greenplum cluster. (6.17)
	Version *string `pulumi:"version"`
	// The availability zone where the Greenplum hosts will be created.
	Zone *string `pulumi:"zone"`
}

type MdbGreenplumClusterState struct {
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment.
	AssignPublicIp pulumi.BoolPtrInput
	// Creation timestamp of the cluster.
	CreatedAt pulumi.StringPtrInput
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the Greenplum cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Aggregated health of the cluster.
	Health pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Greenplum cluster.
	Labels pulumi.StringMapInput
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount pulumi.IntPtrInput
	// (Computed) Info about hosts in master subcluster. The structure is documented below.
	MasterHosts MdbGreenplumClusterMasterHostArrayInput
	// Settings for master subcluster. The structure is documented below.
	MasterSubcluster MdbGreenplumClusterMasterSubclusterPtrInput
	// Name of the Greenplum cluster. Provided by the client when the cluster is created.
	Name pulumi.StringPtrInput
	// ID of the network, to which the Greenplum cluster uses.
	NetworkId pulumi.StringPtrInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount pulumi.IntPtrInput
	// (Computed) Info about hosts in segment subcluster. The structure is documented below.
	SegmentHosts MdbGreenplumClusterSegmentHostArrayInput
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost pulumi.IntPtrInput
	// Settings for segment subcluster. The structure is documented below.
	SegmentSubcluster MdbGreenplumClusterSegmentSubclusterPtrInput
	// Status of the cluster.
	Status pulumi.StringPtrInput
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId pulumi.StringPtrInput
	// Greenplum cluster admin user name.
	UserName pulumi.StringPtrInput
	// Greenplum cluster admin password name.
	UserPassword pulumi.StringPtrInput
	// Version of the Greenplum cluster. (6.17)
	Version pulumi.StringPtrInput
	// The availability zone where the Greenplum hosts will be created.
	Zone pulumi.StringPtrInput
}

func (MdbGreenplumClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbGreenplumClusterState)(nil)).Elem()
}

type mdbGreenplumClusterArgs struct {
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment.
	AssignPublicIp bool `pulumi:"assignPublicIp"`
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the Greenplum cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Greenplum cluster.
	Labels map[string]string `pulumi:"labels"`
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount int `pulumi:"masterHostCount"`
	// Settings for master subcluster. The structure is documented below.
	MasterSubcluster MdbGreenplumClusterMasterSubcluster `pulumi:"masterSubcluster"`
	// Name of the Greenplum cluster. Provided by the client when the cluster is created.
	Name *string `pulumi:"name"`
	// ID of the network, to which the Greenplum cluster uses.
	NetworkId string `pulumi:"networkId"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount int `pulumi:"segmentHostCount"`
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost int `pulumi:"segmentInHost"`
	// Settings for segment subcluster. The structure is documented below.
	SegmentSubcluster MdbGreenplumClusterSegmentSubcluster `pulumi:"segmentSubcluster"`
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId string `pulumi:"subnetId"`
	// Greenplum cluster admin user name.
	UserName string `pulumi:"userName"`
	// Greenplum cluster admin password name.
	UserPassword string `pulumi:"userPassword"`
	// Version of the Greenplum cluster. (6.17)
	Version string `pulumi:"version"`
	// The availability zone where the Greenplum hosts will be created.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a MdbGreenplumCluster resource.
type MdbGreenplumClusterArgs struct {
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment.
	AssignPublicIp pulumi.BoolInput
	// Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the Greenplum cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringInput
	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Greenplum cluster.
	Labels pulumi.StringMapInput
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount pulumi.IntInput
	// Settings for master subcluster. The structure is documented below.
	MasterSubcluster MdbGreenplumClusterMasterSubclusterInput
	// Name of the Greenplum cluster. Provided by the client when the cluster is created.
	Name pulumi.StringPtrInput
	// ID of the network, to which the Greenplum cluster uses.
	NetworkId pulumi.StringInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount pulumi.IntInput
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost pulumi.IntInput
	// Settings for segment subcluster. The structure is documented below.
	SegmentSubcluster MdbGreenplumClusterSegmentSubclusterInput
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId pulumi.StringInput
	// Greenplum cluster admin user name.
	UserName pulumi.StringInput
	// Greenplum cluster admin password name.
	UserPassword pulumi.StringInput
	// Version of the Greenplum cluster. (6.17)
	Version pulumi.StringInput
	// The availability zone where the Greenplum hosts will be created.
	Zone pulumi.StringInput
}

func (MdbGreenplumClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbGreenplumClusterArgs)(nil)).Elem()
}

type MdbGreenplumClusterInput interface {
	pulumi.Input

	ToMdbGreenplumClusterOutput() MdbGreenplumClusterOutput
	ToMdbGreenplumClusterOutputWithContext(ctx context.Context) MdbGreenplumClusterOutput
}

func (*MdbGreenplumCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*MdbGreenplumCluster)(nil))
}

func (i *MdbGreenplumCluster) ToMdbGreenplumClusterOutput() MdbGreenplumClusterOutput {
	return i.ToMdbGreenplumClusterOutputWithContext(context.Background())
}

func (i *MdbGreenplumCluster) ToMdbGreenplumClusterOutputWithContext(ctx context.Context) MdbGreenplumClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbGreenplumClusterOutput)
}

func (i *MdbGreenplumCluster) ToMdbGreenplumClusterPtrOutput() MdbGreenplumClusterPtrOutput {
	return i.ToMdbGreenplumClusterPtrOutputWithContext(context.Background())
}

func (i *MdbGreenplumCluster) ToMdbGreenplumClusterPtrOutputWithContext(ctx context.Context) MdbGreenplumClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbGreenplumClusterPtrOutput)
}

type MdbGreenplumClusterPtrInput interface {
	pulumi.Input

	ToMdbGreenplumClusterPtrOutput() MdbGreenplumClusterPtrOutput
	ToMdbGreenplumClusterPtrOutputWithContext(ctx context.Context) MdbGreenplumClusterPtrOutput
}

type mdbGreenplumClusterPtrType MdbGreenplumClusterArgs

func (*mdbGreenplumClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbGreenplumCluster)(nil))
}

func (i *mdbGreenplumClusterPtrType) ToMdbGreenplumClusterPtrOutput() MdbGreenplumClusterPtrOutput {
	return i.ToMdbGreenplumClusterPtrOutputWithContext(context.Background())
}

func (i *mdbGreenplumClusterPtrType) ToMdbGreenplumClusterPtrOutputWithContext(ctx context.Context) MdbGreenplumClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbGreenplumClusterPtrOutput)
}

// MdbGreenplumClusterArrayInput is an input type that accepts MdbGreenplumClusterArray and MdbGreenplumClusterArrayOutput values.
// You can construct a concrete instance of `MdbGreenplumClusterArrayInput` via:
//
//          MdbGreenplumClusterArray{ MdbGreenplumClusterArgs{...} }
type MdbGreenplumClusterArrayInput interface {
	pulumi.Input

	ToMdbGreenplumClusterArrayOutput() MdbGreenplumClusterArrayOutput
	ToMdbGreenplumClusterArrayOutputWithContext(context.Context) MdbGreenplumClusterArrayOutput
}

type MdbGreenplumClusterArray []MdbGreenplumClusterInput

func (MdbGreenplumClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MdbGreenplumCluster)(nil))
}

func (i MdbGreenplumClusterArray) ToMdbGreenplumClusterArrayOutput() MdbGreenplumClusterArrayOutput {
	return i.ToMdbGreenplumClusterArrayOutputWithContext(context.Background())
}

func (i MdbGreenplumClusterArray) ToMdbGreenplumClusterArrayOutputWithContext(ctx context.Context) MdbGreenplumClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbGreenplumClusterArrayOutput)
}

// MdbGreenplumClusterMapInput is an input type that accepts MdbGreenplumClusterMap and MdbGreenplumClusterMapOutput values.
// You can construct a concrete instance of `MdbGreenplumClusterMapInput` via:
//
//          MdbGreenplumClusterMap{ "key": MdbGreenplumClusterArgs{...} }
type MdbGreenplumClusterMapInput interface {
	pulumi.Input

	ToMdbGreenplumClusterMapOutput() MdbGreenplumClusterMapOutput
	ToMdbGreenplumClusterMapOutputWithContext(context.Context) MdbGreenplumClusterMapOutput
}

type MdbGreenplumClusterMap map[string]MdbGreenplumClusterInput

func (MdbGreenplumClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MdbGreenplumCluster)(nil))
}

func (i MdbGreenplumClusterMap) ToMdbGreenplumClusterMapOutput() MdbGreenplumClusterMapOutput {
	return i.ToMdbGreenplumClusterMapOutputWithContext(context.Background())
}

func (i MdbGreenplumClusterMap) ToMdbGreenplumClusterMapOutputWithContext(ctx context.Context) MdbGreenplumClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbGreenplumClusterMapOutput)
}

type MdbGreenplumClusterOutput struct {
	*pulumi.OutputState
}

func (MdbGreenplumClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MdbGreenplumCluster)(nil))
}

func (o MdbGreenplumClusterOutput) ToMdbGreenplumClusterOutput() MdbGreenplumClusterOutput {
	return o
}

func (o MdbGreenplumClusterOutput) ToMdbGreenplumClusterOutputWithContext(ctx context.Context) MdbGreenplumClusterOutput {
	return o
}

func (o MdbGreenplumClusterOutput) ToMdbGreenplumClusterPtrOutput() MdbGreenplumClusterPtrOutput {
	return o.ToMdbGreenplumClusterPtrOutputWithContext(context.Background())
}

func (o MdbGreenplumClusterOutput) ToMdbGreenplumClusterPtrOutputWithContext(ctx context.Context) MdbGreenplumClusterPtrOutput {
	return o.ApplyT(func(v MdbGreenplumCluster) *MdbGreenplumCluster {
		return &v
	}).(MdbGreenplumClusterPtrOutput)
}

type MdbGreenplumClusterPtrOutput struct {
	*pulumi.OutputState
}

func (MdbGreenplumClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbGreenplumCluster)(nil))
}

func (o MdbGreenplumClusterPtrOutput) ToMdbGreenplumClusterPtrOutput() MdbGreenplumClusterPtrOutput {
	return o
}

func (o MdbGreenplumClusterPtrOutput) ToMdbGreenplumClusterPtrOutputWithContext(ctx context.Context) MdbGreenplumClusterPtrOutput {
	return o
}

type MdbGreenplumClusterArrayOutput struct{ *pulumi.OutputState }

func (MdbGreenplumClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MdbGreenplumCluster)(nil))
}

func (o MdbGreenplumClusterArrayOutput) ToMdbGreenplumClusterArrayOutput() MdbGreenplumClusterArrayOutput {
	return o
}

func (o MdbGreenplumClusterArrayOutput) ToMdbGreenplumClusterArrayOutputWithContext(ctx context.Context) MdbGreenplumClusterArrayOutput {
	return o
}

func (o MdbGreenplumClusterArrayOutput) Index(i pulumi.IntInput) MdbGreenplumClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MdbGreenplumCluster {
		return vs[0].([]MdbGreenplumCluster)[vs[1].(int)]
	}).(MdbGreenplumClusterOutput)
}

type MdbGreenplumClusterMapOutput struct{ *pulumi.OutputState }

func (MdbGreenplumClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MdbGreenplumCluster)(nil))
}

func (o MdbGreenplumClusterMapOutput) ToMdbGreenplumClusterMapOutput() MdbGreenplumClusterMapOutput {
	return o
}

func (o MdbGreenplumClusterMapOutput) ToMdbGreenplumClusterMapOutputWithContext(ctx context.Context) MdbGreenplumClusterMapOutput {
	return o
}

func (o MdbGreenplumClusterMapOutput) MapIndex(k pulumi.StringInput) MdbGreenplumClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MdbGreenplumCluster {
		return vs[0].(map[string]MdbGreenplumCluster)[vs[1].(string)]
	}).(MdbGreenplumClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(MdbGreenplumClusterOutput{})
	pulumi.RegisterOutputType(MdbGreenplumClusterPtrOutput{})
	pulumi.RegisterOutputType(MdbGreenplumClusterArrayOutput{})
	pulumi.RegisterOutputType(MdbGreenplumClusterMapOutput{})
}