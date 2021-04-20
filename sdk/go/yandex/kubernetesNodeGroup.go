// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Yandex Kubernetes Node Group.
//
// ## Import
//
// A Yandex Kubernetes Node Group can be imported using the `id` of the resource, e.g.
//
// ```sh
//  $ pulumi import yandex:index/kubernetesNodeGroup:KubernetesNodeGroup default node_group_id
// ```
type KubernetesNodeGroup struct {
	pulumi.CustomResourceState

	// This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.
	AllocationPolicy KubernetesNodeGroupAllocationPolicyOutput `pulumi:"allocationPolicy"`
	// A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).
	AllowedUnsafeSysctls pulumi.StringArrayOutput `pulumi:"allowedUnsafeSysctls"`
	// The ID of the Kubernetes cluster that this node group belongs to.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// (Computed) The Kubernetes node group creation timestamp.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Deploy policy of the node group. The structure is documented below.
	DeployPolicy KubernetesNodeGroupDeployPolicyOutput `pulumi:"deployPolicy"`
	// A description of the Kubernetes node group.
	Description pulumi.StringOutput `pulumi:"description"`
	// ID of instance group that is used to manage this Kubernetes node group.
	InstanceGroupId pulumi.StringOutput `pulumi:"instanceGroupId"`
	// Template used to create compute instances in this Kubernetes node group. The structure is documented below.
	InstanceTemplate KubernetesNodeGroupInstanceTemplateOutput `pulumi:"instanceTemplate"`
	// A set of key/value label pairs assigned to the Kubernetes node group.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// (Computed) Maintenance policy for this Kubernetes node group.
	// If policy is omitted, automatic revision upgrades are enabled and could happen at any time.
	// Revision upgrades are performed only within the same minor version, e.g. 1.13.
	// Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.
	MaintenancePolicy KubernetesNodeGroupMaintenancePolicyOutput `pulumi:"maintenancePolicy"`
	// Name of a specific Kubernetes node group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.
	NodeLabels pulumi.StringMapOutput `pulumi:"nodeLabels"`
	// A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.
	NodeTaints pulumi.StringArrayOutput `pulumi:"nodeTaints"`
	// Scale policy of the node group. The structure is documented below.
	ScalePolicy KubernetesNodeGroupScalePolicyOutput `pulumi:"scalePolicy"`
	// (Computed) Status of the Kubernetes node group.
	Status pulumi.StringOutput `pulumi:"status"`
	// Version of Kubernetes that will be used for Kubernetes node group.
	Version pulumi.StringOutput `pulumi:"version"`
	// Information about Kubernetes node group version. The structure is documented below.
	VersionInfo KubernetesNodeGroupVersionInfoOutput `pulumi:"versionInfo"`
}

// NewKubernetesNodeGroup registers a new resource with the given unique name, arguments, and options.
func NewKubernetesNodeGroup(ctx *pulumi.Context,
	name string, args *KubernetesNodeGroupArgs, opts ...pulumi.ResourceOption) (*KubernetesNodeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceTemplate == nil {
		return nil, errors.New("invalid value for required argument 'InstanceTemplate'")
	}
	if args.ScalePolicy == nil {
		return nil, errors.New("invalid value for required argument 'ScalePolicy'")
	}
	var resource KubernetesNodeGroup
	err := ctx.RegisterResource("yandex:index/kubernetesNodeGroup:KubernetesNodeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesNodeGroup gets an existing KubernetesNodeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesNodeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesNodeGroupState, opts ...pulumi.ResourceOption) (*KubernetesNodeGroup, error) {
	var resource KubernetesNodeGroup
	err := ctx.ReadResource("yandex:index/kubernetesNodeGroup:KubernetesNodeGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesNodeGroup resources.
type kubernetesNodeGroupState struct {
	// This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.
	AllocationPolicy *KubernetesNodeGroupAllocationPolicy `pulumi:"allocationPolicy"`
	// A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).
	AllowedUnsafeSysctls []string `pulumi:"allowedUnsafeSysctls"`
	// The ID of the Kubernetes cluster that this node group belongs to.
	ClusterId *string `pulumi:"clusterId"`
	// (Computed) The Kubernetes node group creation timestamp.
	CreatedAt *string `pulumi:"createdAt"`
	// Deploy policy of the node group. The structure is documented below.
	DeployPolicy *KubernetesNodeGroupDeployPolicy `pulumi:"deployPolicy"`
	// A description of the Kubernetes node group.
	Description *string `pulumi:"description"`
	// ID of instance group that is used to manage this Kubernetes node group.
	InstanceGroupId *string `pulumi:"instanceGroupId"`
	// Template used to create compute instances in this Kubernetes node group. The structure is documented below.
	InstanceTemplate *KubernetesNodeGroupInstanceTemplate `pulumi:"instanceTemplate"`
	// A set of key/value label pairs assigned to the Kubernetes node group.
	Labels map[string]string `pulumi:"labels"`
	// (Computed) Maintenance policy for this Kubernetes node group.
	// If policy is omitted, automatic revision upgrades are enabled and could happen at any time.
	// Revision upgrades are performed only within the same minor version, e.g. 1.13.
	// Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.
	MaintenancePolicy *KubernetesNodeGroupMaintenancePolicy `pulumi:"maintenancePolicy"`
	// Name of a specific Kubernetes node group.
	Name *string `pulumi:"name"`
	// A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.
	NodeTaints []string `pulumi:"nodeTaints"`
	// Scale policy of the node group. The structure is documented below.
	ScalePolicy *KubernetesNodeGroupScalePolicy `pulumi:"scalePolicy"`
	// (Computed) Status of the Kubernetes node group.
	Status *string `pulumi:"status"`
	// Version of Kubernetes that will be used for Kubernetes node group.
	Version *string `pulumi:"version"`
	// Information about Kubernetes node group version. The structure is documented below.
	VersionInfo *KubernetesNodeGroupVersionInfo `pulumi:"versionInfo"`
}

type KubernetesNodeGroupState struct {
	// This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.
	AllocationPolicy KubernetesNodeGroupAllocationPolicyPtrInput
	// A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).
	AllowedUnsafeSysctls pulumi.StringArrayInput
	// The ID of the Kubernetes cluster that this node group belongs to.
	ClusterId pulumi.StringPtrInput
	// (Computed) The Kubernetes node group creation timestamp.
	CreatedAt pulumi.StringPtrInput
	// Deploy policy of the node group. The structure is documented below.
	DeployPolicy KubernetesNodeGroupDeployPolicyPtrInput
	// A description of the Kubernetes node group.
	Description pulumi.StringPtrInput
	// ID of instance group that is used to manage this Kubernetes node group.
	InstanceGroupId pulumi.StringPtrInput
	// Template used to create compute instances in this Kubernetes node group. The structure is documented below.
	InstanceTemplate KubernetesNodeGroupInstanceTemplatePtrInput
	// A set of key/value label pairs assigned to the Kubernetes node group.
	Labels pulumi.StringMapInput
	// (Computed) Maintenance policy for this Kubernetes node group.
	// If policy is omitted, automatic revision upgrades are enabled and could happen at any time.
	// Revision upgrades are performed only within the same minor version, e.g. 1.13.
	// Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.
	MaintenancePolicy KubernetesNodeGroupMaintenancePolicyPtrInput
	// Name of a specific Kubernetes node group.
	Name pulumi.StringPtrInput
	// A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.
	NodeLabels pulumi.StringMapInput
	// A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.
	NodeTaints pulumi.StringArrayInput
	// Scale policy of the node group. The structure is documented below.
	ScalePolicy KubernetesNodeGroupScalePolicyPtrInput
	// (Computed) Status of the Kubernetes node group.
	Status pulumi.StringPtrInput
	// Version of Kubernetes that will be used for Kubernetes node group.
	Version pulumi.StringPtrInput
	// Information about Kubernetes node group version. The structure is documented below.
	VersionInfo KubernetesNodeGroupVersionInfoPtrInput
}

func (KubernetesNodeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesNodeGroupState)(nil)).Elem()
}

type kubernetesNodeGroupArgs struct {
	// This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.
	AllocationPolicy *KubernetesNodeGroupAllocationPolicy `pulumi:"allocationPolicy"`
	// A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).
	AllowedUnsafeSysctls []string `pulumi:"allowedUnsafeSysctls"`
	// The ID of the Kubernetes cluster that this node group belongs to.
	ClusterId string `pulumi:"clusterId"`
	// Deploy policy of the node group. The structure is documented below.
	DeployPolicy *KubernetesNodeGroupDeployPolicy `pulumi:"deployPolicy"`
	// A description of the Kubernetes node group.
	Description *string `pulumi:"description"`
	// Template used to create compute instances in this Kubernetes node group. The structure is documented below.
	InstanceTemplate KubernetesNodeGroupInstanceTemplate `pulumi:"instanceTemplate"`
	// A set of key/value label pairs assigned to the Kubernetes node group.
	Labels map[string]string `pulumi:"labels"`
	// (Computed) Maintenance policy for this Kubernetes node group.
	// If policy is omitted, automatic revision upgrades are enabled and could happen at any time.
	// Revision upgrades are performed only within the same minor version, e.g. 1.13.
	// Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.
	MaintenancePolicy *KubernetesNodeGroupMaintenancePolicy `pulumi:"maintenancePolicy"`
	// Name of a specific Kubernetes node group.
	Name *string `pulumi:"name"`
	// A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.
	NodeTaints []string `pulumi:"nodeTaints"`
	// Scale policy of the node group. The structure is documented below.
	ScalePolicy KubernetesNodeGroupScalePolicy `pulumi:"scalePolicy"`
	// Version of Kubernetes that will be used for Kubernetes node group.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a KubernetesNodeGroup resource.
type KubernetesNodeGroupArgs struct {
	// This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.
	AllocationPolicy KubernetesNodeGroupAllocationPolicyPtrInput
	// A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).
	AllowedUnsafeSysctls pulumi.StringArrayInput
	// The ID of the Kubernetes cluster that this node group belongs to.
	ClusterId pulumi.StringInput
	// Deploy policy of the node group. The structure is documented below.
	DeployPolicy KubernetesNodeGroupDeployPolicyPtrInput
	// A description of the Kubernetes node group.
	Description pulumi.StringPtrInput
	// Template used to create compute instances in this Kubernetes node group. The structure is documented below.
	InstanceTemplate KubernetesNodeGroupInstanceTemplateInput
	// A set of key/value label pairs assigned to the Kubernetes node group.
	Labels pulumi.StringMapInput
	// (Computed) Maintenance policy for this Kubernetes node group.
	// If policy is omitted, automatic revision upgrades are enabled and could happen at any time.
	// Revision upgrades are performed only within the same minor version, e.g. 1.13.
	// Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.
	MaintenancePolicy KubernetesNodeGroupMaintenancePolicyPtrInput
	// Name of a specific Kubernetes node group.
	Name pulumi.StringPtrInput
	// A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.
	NodeLabels pulumi.StringMapInput
	// A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.
	NodeTaints pulumi.StringArrayInput
	// Scale policy of the node group. The structure is documented below.
	ScalePolicy KubernetesNodeGroupScalePolicyInput
	// Version of Kubernetes that will be used for Kubernetes node group.
	Version pulumi.StringPtrInput
}

func (KubernetesNodeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesNodeGroupArgs)(nil)).Elem()
}

type KubernetesNodeGroupInput interface {
	pulumi.Input

	ToKubernetesNodeGroupOutput() KubernetesNodeGroupOutput
	ToKubernetesNodeGroupOutputWithContext(ctx context.Context) KubernetesNodeGroupOutput
}

func (*KubernetesNodeGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesNodeGroup)(nil))
}

func (i *KubernetesNodeGroup) ToKubernetesNodeGroupOutput() KubernetesNodeGroupOutput {
	return i.ToKubernetesNodeGroupOutputWithContext(context.Background())
}

func (i *KubernetesNodeGroup) ToKubernetesNodeGroupOutputWithContext(ctx context.Context) KubernetesNodeGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesNodeGroupOutput)
}

func (i *KubernetesNodeGroup) ToKubernetesNodeGroupPtrOutput() KubernetesNodeGroupPtrOutput {
	return i.ToKubernetesNodeGroupPtrOutputWithContext(context.Background())
}

func (i *KubernetesNodeGroup) ToKubernetesNodeGroupPtrOutputWithContext(ctx context.Context) KubernetesNodeGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesNodeGroupPtrOutput)
}

type KubernetesNodeGroupPtrInput interface {
	pulumi.Input

	ToKubernetesNodeGroupPtrOutput() KubernetesNodeGroupPtrOutput
	ToKubernetesNodeGroupPtrOutputWithContext(ctx context.Context) KubernetesNodeGroupPtrOutput
}

type kubernetesNodeGroupPtrType KubernetesNodeGroupArgs

func (*kubernetesNodeGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesNodeGroup)(nil))
}

func (i *kubernetesNodeGroupPtrType) ToKubernetesNodeGroupPtrOutput() KubernetesNodeGroupPtrOutput {
	return i.ToKubernetesNodeGroupPtrOutputWithContext(context.Background())
}

func (i *kubernetesNodeGroupPtrType) ToKubernetesNodeGroupPtrOutputWithContext(ctx context.Context) KubernetesNodeGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesNodeGroupPtrOutput)
}

// KubernetesNodeGroupArrayInput is an input type that accepts KubernetesNodeGroupArray and KubernetesNodeGroupArrayOutput values.
// You can construct a concrete instance of `KubernetesNodeGroupArrayInput` via:
//
//          KubernetesNodeGroupArray{ KubernetesNodeGroupArgs{...} }
type KubernetesNodeGroupArrayInput interface {
	pulumi.Input

	ToKubernetesNodeGroupArrayOutput() KubernetesNodeGroupArrayOutput
	ToKubernetesNodeGroupArrayOutputWithContext(context.Context) KubernetesNodeGroupArrayOutput
}

type KubernetesNodeGroupArray []KubernetesNodeGroupInput

func (KubernetesNodeGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*KubernetesNodeGroup)(nil))
}

func (i KubernetesNodeGroupArray) ToKubernetesNodeGroupArrayOutput() KubernetesNodeGroupArrayOutput {
	return i.ToKubernetesNodeGroupArrayOutputWithContext(context.Background())
}

func (i KubernetesNodeGroupArray) ToKubernetesNodeGroupArrayOutputWithContext(ctx context.Context) KubernetesNodeGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesNodeGroupArrayOutput)
}

// KubernetesNodeGroupMapInput is an input type that accepts KubernetesNodeGroupMap and KubernetesNodeGroupMapOutput values.
// You can construct a concrete instance of `KubernetesNodeGroupMapInput` via:
//
//          KubernetesNodeGroupMap{ "key": KubernetesNodeGroupArgs{...} }
type KubernetesNodeGroupMapInput interface {
	pulumi.Input

	ToKubernetesNodeGroupMapOutput() KubernetesNodeGroupMapOutput
	ToKubernetesNodeGroupMapOutputWithContext(context.Context) KubernetesNodeGroupMapOutput
}

type KubernetesNodeGroupMap map[string]KubernetesNodeGroupInput

func (KubernetesNodeGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*KubernetesNodeGroup)(nil))
}

func (i KubernetesNodeGroupMap) ToKubernetesNodeGroupMapOutput() KubernetesNodeGroupMapOutput {
	return i.ToKubernetesNodeGroupMapOutputWithContext(context.Background())
}

func (i KubernetesNodeGroupMap) ToKubernetesNodeGroupMapOutputWithContext(ctx context.Context) KubernetesNodeGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesNodeGroupMapOutput)
}

type KubernetesNodeGroupOutput struct {
	*pulumi.OutputState
}

func (KubernetesNodeGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesNodeGroup)(nil))
}

func (o KubernetesNodeGroupOutput) ToKubernetesNodeGroupOutput() KubernetesNodeGroupOutput {
	return o
}

func (o KubernetesNodeGroupOutput) ToKubernetesNodeGroupOutputWithContext(ctx context.Context) KubernetesNodeGroupOutput {
	return o
}

func (o KubernetesNodeGroupOutput) ToKubernetesNodeGroupPtrOutput() KubernetesNodeGroupPtrOutput {
	return o.ToKubernetesNodeGroupPtrOutputWithContext(context.Background())
}

func (o KubernetesNodeGroupOutput) ToKubernetesNodeGroupPtrOutputWithContext(ctx context.Context) KubernetesNodeGroupPtrOutput {
	return o.ApplyT(func(v KubernetesNodeGroup) *KubernetesNodeGroup {
		return &v
	}).(KubernetesNodeGroupPtrOutput)
}

type KubernetesNodeGroupPtrOutput struct {
	*pulumi.OutputState
}

func (KubernetesNodeGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesNodeGroup)(nil))
}

func (o KubernetesNodeGroupPtrOutput) ToKubernetesNodeGroupPtrOutput() KubernetesNodeGroupPtrOutput {
	return o
}

func (o KubernetesNodeGroupPtrOutput) ToKubernetesNodeGroupPtrOutputWithContext(ctx context.Context) KubernetesNodeGroupPtrOutput {
	return o
}

type KubernetesNodeGroupArrayOutput struct{ *pulumi.OutputState }

func (KubernetesNodeGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesNodeGroup)(nil))
}

func (o KubernetesNodeGroupArrayOutput) ToKubernetesNodeGroupArrayOutput() KubernetesNodeGroupArrayOutput {
	return o
}

func (o KubernetesNodeGroupArrayOutput) ToKubernetesNodeGroupArrayOutputWithContext(ctx context.Context) KubernetesNodeGroupArrayOutput {
	return o
}

func (o KubernetesNodeGroupArrayOutput) Index(i pulumi.IntInput) KubernetesNodeGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesNodeGroup {
		return vs[0].([]KubernetesNodeGroup)[vs[1].(int)]
	}).(KubernetesNodeGroupOutput)
}

type KubernetesNodeGroupMapOutput struct{ *pulumi.OutputState }

func (KubernetesNodeGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KubernetesNodeGroup)(nil))
}

func (o KubernetesNodeGroupMapOutput) ToKubernetesNodeGroupMapOutput() KubernetesNodeGroupMapOutput {
	return o
}

func (o KubernetesNodeGroupMapOutput) ToKubernetesNodeGroupMapOutputWithContext(ctx context.Context) KubernetesNodeGroupMapOutput {
	return o
}

func (o KubernetesNodeGroupMapOutput) MapIndex(k pulumi.StringInput) KubernetesNodeGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KubernetesNodeGroup {
		return vs[0].(map[string]KubernetesNodeGroup)[vs[1].(string)]
	}).(KubernetesNodeGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(KubernetesNodeGroupOutput{})
	pulumi.RegisterOutputType(KubernetesNodeGroupPtrOutput{})
	pulumi.RegisterOutputType(KubernetesNodeGroupArrayOutput{})
	pulumi.RegisterOutputType(KubernetesNodeGroupMapOutput{})
}
