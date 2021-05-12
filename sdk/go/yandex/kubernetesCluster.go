// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Yandex Kubernetes Cluster.
//
// ## Import
//
// A Managed Kubernetes cluster can be imported using the `id` of the resource, e.g.
//
// ```sh
//  $ pulumi import yandex:index/kubernetesCluster:KubernetesCluster default cluster_id
// ```
type KubernetesCluster struct {
	pulumi.CustomResourceState

	// CIDR block. IP range for allocating pod addresses.
	// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
	// set up for this CIDR blocks in node subnets.
	ClusterIpv4Range pulumi.StringOutput `pulumi:"clusterIpv4Range"`
	// Identical to clusterIpv4Range but for IPv6 protocol.
	ClusterIpv6Range pulumi.StringOutput `pulumi:"clusterIpv6Range"`
	// (Computed) The Kubernetes cluster creation timestamp.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A description of the Kubernetes cluster.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ID of the folder that the Kubernetes cluster belongs to.
	// If it is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// (Computed) Health of the Kubernetes cluster.
	Health pulumi.StringOutput `pulumi:"health"`
	// cluster KMS provider parameters.
	KmsProvider KubernetesClusterKmsProviderPtrOutput `pulumi:"kmsProvider"`
	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Log group where cluster stores cluster system logs, like audit, events, or controlplane logs.
	LogGroupId pulumi.StringOutput `pulumi:"logGroupId"`
	// Kubernetes master configuration options. The structure is documented below.
	Master KubernetesClusterMasterOutput `pulumi:"master"`
	// Name of a specific Kubernetes cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the cluster network.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Network policy provider for the cluster. Possible values: `CALICO`.
	NetworkPolicyProvider pulumi.StringPtrOutput `pulumi:"networkPolicyProvider"`
	// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
	NodeIpv4CidrMaskSize pulumi.IntPtrOutput `pulumi:"nodeIpv4CidrMaskSize"`
	// Service account to be used by the worker nodes of the Kubernetes cluster
	// to access Container Registry or to push node logs and metrics.
	NodeServiceAccountId pulumi.StringOutput `pulumi:"nodeServiceAccountId"`
	// Cluster release channel.
	ReleaseChannel pulumi.StringOutput `pulumi:"releaseChannel"`
	// Service account to be used for provisioning Compute Cloud and VPC resources
	// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
	// cluster will be located and on the folder where selected network resides.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
	// CIDR block. IP range Kubernetes service Kubernetes cluster
	// IP addresses will be allocated from. It should not overlap with any subnet in the network
	// the Kubernetes cluster located in.
	ServiceIpv4Range pulumi.StringOutput `pulumi:"serviceIpv4Range"`
	// Identical to serviceIpv4Range but for IPv6 protocol.
	ServiceIpv6Range pulumi.StringOutput `pulumi:"serviceIpv6Range"`
	// (Computed)Status of the Kubernetes cluster.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewKubernetesCluster registers a new resource with the given unique name, arguments, and options.
func NewKubernetesCluster(ctx *pulumi.Context,
	name string, args *KubernetesClusterArgs, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Master == nil {
		return nil, errors.New("invalid value for required argument 'Master'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.NodeServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'NodeServiceAccountId'")
	}
	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	var resource KubernetesCluster
	err := ctx.RegisterResource("yandex:index/kubernetesCluster:KubernetesCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesCluster gets an existing KubernetesCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesClusterState, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	var resource KubernetesCluster
	err := ctx.ReadResource("yandex:index/kubernetesCluster:KubernetesCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesCluster resources.
type kubernetesClusterState struct {
	// CIDR block. IP range for allocating pod addresses.
	// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
	// set up for this CIDR blocks in node subnets.
	ClusterIpv4Range *string `pulumi:"clusterIpv4Range"`
	// Identical to clusterIpv4Range but for IPv6 protocol.
	ClusterIpv6Range *string `pulumi:"clusterIpv6Range"`
	// (Computed) The Kubernetes cluster creation timestamp.
	CreatedAt *string `pulumi:"createdAt"`
	// A description of the Kubernetes cluster.
	Description *string `pulumi:"description"`
	// The ID of the folder that the Kubernetes cluster belongs to.
	// If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// (Computed) Health of the Kubernetes cluster.
	Health *string `pulumi:"health"`
	// cluster KMS provider parameters.
	KmsProvider *KubernetesClusterKmsProvider `pulumi:"kmsProvider"`
	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels map[string]string `pulumi:"labels"`
	// Log group where cluster stores cluster system logs, like audit, events, or controlplane logs.
	LogGroupId *string `pulumi:"logGroupId"`
	// Kubernetes master configuration options. The structure is documented below.
	Master *KubernetesClusterMaster `pulumi:"master"`
	// Name of a specific Kubernetes cluster.
	Name *string `pulumi:"name"`
	// The ID of the cluster network.
	NetworkId *string `pulumi:"networkId"`
	// Network policy provider for the cluster. Possible values: `CALICO`.
	NetworkPolicyProvider *string `pulumi:"networkPolicyProvider"`
	// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
	NodeIpv4CidrMaskSize *int `pulumi:"nodeIpv4CidrMaskSize"`
	// Service account to be used by the worker nodes of the Kubernetes cluster
	// to access Container Registry or to push node logs and metrics.
	NodeServiceAccountId *string `pulumi:"nodeServiceAccountId"`
	// Cluster release channel.
	ReleaseChannel *string `pulumi:"releaseChannel"`
	// Service account to be used for provisioning Compute Cloud and VPC resources
	// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
	// cluster will be located and on the folder where selected network resides.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// CIDR block. IP range Kubernetes service Kubernetes cluster
	// IP addresses will be allocated from. It should not overlap with any subnet in the network
	// the Kubernetes cluster located in.
	ServiceIpv4Range *string `pulumi:"serviceIpv4Range"`
	// Identical to serviceIpv4Range but for IPv6 protocol.
	ServiceIpv6Range *string `pulumi:"serviceIpv6Range"`
	// (Computed)Status of the Kubernetes cluster.
	Status *string `pulumi:"status"`
}

type KubernetesClusterState struct {
	// CIDR block. IP range for allocating pod addresses.
	// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
	// set up for this CIDR blocks in node subnets.
	ClusterIpv4Range pulumi.StringPtrInput
	// Identical to clusterIpv4Range but for IPv6 protocol.
	ClusterIpv6Range pulumi.StringPtrInput
	// (Computed) The Kubernetes cluster creation timestamp.
	CreatedAt pulumi.StringPtrInput
	// A description of the Kubernetes cluster.
	Description pulumi.StringPtrInput
	// The ID of the folder that the Kubernetes cluster belongs to.
	// If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// (Computed) Health of the Kubernetes cluster.
	Health pulumi.StringPtrInput
	// cluster KMS provider parameters.
	KmsProvider KubernetesClusterKmsProviderPtrInput
	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels pulumi.StringMapInput
	// Log group where cluster stores cluster system logs, like audit, events, or controlplane logs.
	LogGroupId pulumi.StringPtrInput
	// Kubernetes master configuration options. The structure is documented below.
	Master KubernetesClusterMasterPtrInput
	// Name of a specific Kubernetes cluster.
	Name pulumi.StringPtrInput
	// The ID of the cluster network.
	NetworkId pulumi.StringPtrInput
	// Network policy provider for the cluster. Possible values: `CALICO`.
	NetworkPolicyProvider pulumi.StringPtrInput
	// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
	NodeIpv4CidrMaskSize pulumi.IntPtrInput
	// Service account to be used by the worker nodes of the Kubernetes cluster
	// to access Container Registry or to push node logs and metrics.
	NodeServiceAccountId pulumi.StringPtrInput
	// Cluster release channel.
	ReleaseChannel pulumi.StringPtrInput
	// Service account to be used for provisioning Compute Cloud and VPC resources
	// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
	// cluster will be located and on the folder where selected network resides.
	ServiceAccountId pulumi.StringPtrInput
	// CIDR block. IP range Kubernetes service Kubernetes cluster
	// IP addresses will be allocated from. It should not overlap with any subnet in the network
	// the Kubernetes cluster located in.
	ServiceIpv4Range pulumi.StringPtrInput
	// Identical to serviceIpv4Range but for IPv6 protocol.
	ServiceIpv6Range pulumi.StringPtrInput
	// (Computed)Status of the Kubernetes cluster.
	Status pulumi.StringPtrInput
}

func (KubernetesClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterState)(nil)).Elem()
}

type kubernetesClusterArgs struct {
	// CIDR block. IP range for allocating pod addresses.
	// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
	// set up for this CIDR blocks in node subnets.
	ClusterIpv4Range *string `pulumi:"clusterIpv4Range"`
	// Identical to clusterIpv4Range but for IPv6 protocol.
	ClusterIpv6Range *string `pulumi:"clusterIpv6Range"`
	// A description of the Kubernetes cluster.
	Description *string `pulumi:"description"`
	// The ID of the folder that the Kubernetes cluster belongs to.
	// If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// cluster KMS provider parameters.
	KmsProvider *KubernetesClusterKmsProvider `pulumi:"kmsProvider"`
	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels map[string]string `pulumi:"labels"`
	// Kubernetes master configuration options. The structure is documented below.
	Master KubernetesClusterMaster `pulumi:"master"`
	// Name of a specific Kubernetes cluster.
	Name *string `pulumi:"name"`
	// The ID of the cluster network.
	NetworkId string `pulumi:"networkId"`
	// Network policy provider for the cluster. Possible values: `CALICO`.
	NetworkPolicyProvider *string `pulumi:"networkPolicyProvider"`
	// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
	NodeIpv4CidrMaskSize *int `pulumi:"nodeIpv4CidrMaskSize"`
	// Service account to be used by the worker nodes of the Kubernetes cluster
	// to access Container Registry or to push node logs and metrics.
	NodeServiceAccountId string `pulumi:"nodeServiceAccountId"`
	// Cluster release channel.
	ReleaseChannel *string `pulumi:"releaseChannel"`
	// Service account to be used for provisioning Compute Cloud and VPC resources
	// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
	// cluster will be located and on the folder where selected network resides.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// CIDR block. IP range Kubernetes service Kubernetes cluster
	// IP addresses will be allocated from. It should not overlap with any subnet in the network
	// the Kubernetes cluster located in.
	ServiceIpv4Range *string `pulumi:"serviceIpv4Range"`
	// Identical to serviceIpv4Range but for IPv6 protocol.
	ServiceIpv6Range *string `pulumi:"serviceIpv6Range"`
}

// The set of arguments for constructing a KubernetesCluster resource.
type KubernetesClusterArgs struct {
	// CIDR block. IP range for allocating pod addresses.
	// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
	// set up for this CIDR blocks in node subnets.
	ClusterIpv4Range pulumi.StringPtrInput
	// Identical to clusterIpv4Range but for IPv6 protocol.
	ClusterIpv6Range pulumi.StringPtrInput
	// A description of the Kubernetes cluster.
	Description pulumi.StringPtrInput
	// The ID of the folder that the Kubernetes cluster belongs to.
	// If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// cluster KMS provider parameters.
	KmsProvider KubernetesClusterKmsProviderPtrInput
	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels pulumi.StringMapInput
	// Kubernetes master configuration options. The structure is documented below.
	Master KubernetesClusterMasterInput
	// Name of a specific Kubernetes cluster.
	Name pulumi.StringPtrInput
	// The ID of the cluster network.
	NetworkId pulumi.StringInput
	// Network policy provider for the cluster. Possible values: `CALICO`.
	NetworkPolicyProvider pulumi.StringPtrInput
	// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
	NodeIpv4CidrMaskSize pulumi.IntPtrInput
	// Service account to be used by the worker nodes of the Kubernetes cluster
	// to access Container Registry or to push node logs and metrics.
	NodeServiceAccountId pulumi.StringInput
	// Cluster release channel.
	ReleaseChannel pulumi.StringPtrInput
	// Service account to be used for provisioning Compute Cloud and VPC resources
	// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
	// cluster will be located and on the folder where selected network resides.
	ServiceAccountId pulumi.StringInput
	// CIDR block. IP range Kubernetes service Kubernetes cluster
	// IP addresses will be allocated from. It should not overlap with any subnet in the network
	// the Kubernetes cluster located in.
	ServiceIpv4Range pulumi.StringPtrInput
	// Identical to serviceIpv4Range but for IPv6 protocol.
	ServiceIpv6Range pulumi.StringPtrInput
}

func (KubernetesClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterArgs)(nil)).Elem()
}

type KubernetesClusterInput interface {
	pulumi.Input

	ToKubernetesClusterOutput() KubernetesClusterOutput
	ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput
}

func (*KubernetesCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesCluster)(nil))
}

func (i *KubernetesCluster) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return i.ToKubernetesClusterOutputWithContext(context.Background())
}

func (i *KubernetesCluster) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterOutput)
}

func (i *KubernetesCluster) ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput {
	return i.ToKubernetesClusterPtrOutputWithContext(context.Background())
}

func (i *KubernetesCluster) ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPtrOutput)
}

type KubernetesClusterPtrInput interface {
	pulumi.Input

	ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput
	ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput
}

type kubernetesClusterPtrType KubernetesClusterArgs

func (*kubernetesClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesCluster)(nil))
}

func (i *kubernetesClusterPtrType) ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput {
	return i.ToKubernetesClusterPtrOutputWithContext(context.Background())
}

func (i *kubernetesClusterPtrType) ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPtrOutput)
}

// KubernetesClusterArrayInput is an input type that accepts KubernetesClusterArray and KubernetesClusterArrayOutput values.
// You can construct a concrete instance of `KubernetesClusterArrayInput` via:
//
//          KubernetesClusterArray{ KubernetesClusterArgs{...} }
type KubernetesClusterArrayInput interface {
	pulumi.Input

	ToKubernetesClusterArrayOutput() KubernetesClusterArrayOutput
	ToKubernetesClusterArrayOutputWithContext(context.Context) KubernetesClusterArrayOutput
}

type KubernetesClusterArray []KubernetesClusterInput

func (KubernetesClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*KubernetesCluster)(nil))
}

func (i KubernetesClusterArray) ToKubernetesClusterArrayOutput() KubernetesClusterArrayOutput {
	return i.ToKubernetesClusterArrayOutputWithContext(context.Background())
}

func (i KubernetesClusterArray) ToKubernetesClusterArrayOutputWithContext(ctx context.Context) KubernetesClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterArrayOutput)
}

// KubernetesClusterMapInput is an input type that accepts KubernetesClusterMap and KubernetesClusterMapOutput values.
// You can construct a concrete instance of `KubernetesClusterMapInput` via:
//
//          KubernetesClusterMap{ "key": KubernetesClusterArgs{...} }
type KubernetesClusterMapInput interface {
	pulumi.Input

	ToKubernetesClusterMapOutput() KubernetesClusterMapOutput
	ToKubernetesClusterMapOutputWithContext(context.Context) KubernetesClusterMapOutput
}

type KubernetesClusterMap map[string]KubernetesClusterInput

func (KubernetesClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*KubernetesCluster)(nil))
}

func (i KubernetesClusterMap) ToKubernetesClusterMapOutput() KubernetesClusterMapOutput {
	return i.ToKubernetesClusterMapOutputWithContext(context.Background())
}

func (i KubernetesClusterMap) ToKubernetesClusterMapOutputWithContext(ctx context.Context) KubernetesClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterMapOutput)
}

type KubernetesClusterOutput struct {
	*pulumi.OutputState
}

func (KubernetesClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesCluster)(nil))
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return o
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return o
}

func (o KubernetesClusterOutput) ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput {
	return o.ToKubernetesClusterPtrOutputWithContext(context.Background())
}

func (o KubernetesClusterOutput) ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput {
	return o.ApplyT(func(v KubernetesCluster) *KubernetesCluster {
		return &v
	}).(KubernetesClusterPtrOutput)
}

type KubernetesClusterPtrOutput struct {
	*pulumi.OutputState
}

func (KubernetesClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesCluster)(nil))
}

func (o KubernetesClusterPtrOutput) ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput {
	return o
}

func (o KubernetesClusterPtrOutput) ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput {
	return o
}

type KubernetesClusterArrayOutput struct{ *pulumi.OutputState }

func (KubernetesClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesCluster)(nil))
}

func (o KubernetesClusterArrayOutput) ToKubernetesClusterArrayOutput() KubernetesClusterArrayOutput {
	return o
}

func (o KubernetesClusterArrayOutput) ToKubernetesClusterArrayOutputWithContext(ctx context.Context) KubernetesClusterArrayOutput {
	return o
}

func (o KubernetesClusterArrayOutput) Index(i pulumi.IntInput) KubernetesClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesCluster {
		return vs[0].([]KubernetesCluster)[vs[1].(int)]
	}).(KubernetesClusterOutput)
}

type KubernetesClusterMapOutput struct{ *pulumi.OutputState }

func (KubernetesClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KubernetesCluster)(nil))
}

func (o KubernetesClusterMapOutput) ToKubernetesClusterMapOutput() KubernetesClusterMapOutput {
	return o
}

func (o KubernetesClusterMapOutput) ToKubernetesClusterMapOutputWithContext(ctx context.Context) KubernetesClusterMapOutput {
	return o
}

func (o KubernetesClusterMapOutput) MapIndex(k pulumi.StringInput) KubernetesClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KubernetesCluster {
		return vs[0].(map[string]KubernetesCluster)[vs[1].(string)]
	}).(KubernetesClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(KubernetesClusterOutput{})
	pulumi.RegisterOutputType(KubernetesClusterPtrOutput{})
	pulumi.RegisterOutputType(KubernetesClusterArrayOutput{})
	pulumi.RegisterOutputType(KubernetesClusterMapOutput{})
}
