// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a single Secuirity Group Rule within the Yandex.Cloud. For more information, see the official documentation
// of [security groups](https://cloud.yandex.com/docs/vpc/concepts/security-groups)
// and [security group rules](https://cloud.yandex.com/docs/vpc/concepts/security-groups#rules).
//
// > **NOTE:** There is another way to manage security group rules by `ingress` and `egress` arguments in yandex_vpc_security_group. Both ways are equivalent but not compatible now. Using in-line rules of VpcSecurityGroup with Security Group Rule resource at the same time will cause a conflict of rules configuration.
type VpcSecurityGroupRule struct {
	pulumi.CustomResourceState

	// Description of the rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// direction of the rule. Can be `ingress` (inbound) or `egress` (outbound).
	Direction pulumi.StringOutput `pulumi:"direction"`
	// Minimum port number.
	FromPort pulumi.IntPtrOutput `pulumi:"fromPort"`
	// Labels to assign to this rule.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Port number (if applied to a single port).
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Special-purpose targets such as "selfSecurityGroup". [See docs](https://cloud.yandex.com/docs/vpc/concepts/security-groups) for possible options.
	PredefinedTarget pulumi.StringPtrOutput `pulumi:"predefinedTarget"`
	// One of `ANY`, `TCP`, `UDP`, `ICMP`, `IPV6_ICMP`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// ID of the security group this rule belongs to.
	SecurityGroupBinding pulumi.StringOutput `pulumi:"securityGroupBinding"`
	// Target security group ID for this rule.
	SecurityGroupId pulumi.StringPtrOutput `pulumi:"securityGroupId"`
	// Maximum port number.
	ToPort pulumi.IntPtrOutput `pulumi:"toPort"`
	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks pulumi.StringArrayOutput `pulumi:"v4CidrBlocks"`
	// The blocks of IPv6 addresses for this rule. `v6CidrBlocks` argument is currently not supported. It will be available in the future.
	V6CidrBlocks pulumi.StringArrayOutput `pulumi:"v6CidrBlocks"`
}

// NewVpcSecurityGroupRule registers a new resource with the given unique name, arguments, and options.
func NewVpcSecurityGroupRule(ctx *pulumi.Context,
	name string, args *VpcSecurityGroupRuleArgs, opts ...pulumi.ResourceOption) (*VpcSecurityGroupRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.SecurityGroupBinding == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupBinding'")
	}
	var resource VpcSecurityGroupRule
	err := ctx.RegisterResource("yandex:index/vpcSecurityGroupRule:VpcSecurityGroupRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcSecurityGroupRule gets an existing VpcSecurityGroupRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcSecurityGroupRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcSecurityGroupRuleState, opts ...pulumi.ResourceOption) (*VpcSecurityGroupRule, error) {
	var resource VpcSecurityGroupRule
	err := ctx.ReadResource("yandex:index/vpcSecurityGroupRule:VpcSecurityGroupRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcSecurityGroupRule resources.
type vpcSecurityGroupRuleState struct {
	// Description of the rule.
	Description *string `pulumi:"description"`
	// direction of the rule. Can be `ingress` (inbound) or `egress` (outbound).
	Direction *string `pulumi:"direction"`
	// Minimum port number.
	FromPort *int `pulumi:"fromPort"`
	// Labels to assign to this rule.
	Labels map[string]string `pulumi:"labels"`
	// Port number (if applied to a single port).
	Port *int `pulumi:"port"`
	// Special-purpose targets such as "selfSecurityGroup". [See docs](https://cloud.yandex.com/docs/vpc/concepts/security-groups) for possible options.
	PredefinedTarget *string `pulumi:"predefinedTarget"`
	// One of `ANY`, `TCP`, `UDP`, `ICMP`, `IPV6_ICMP`.
	Protocol *string `pulumi:"protocol"`
	// ID of the security group this rule belongs to.
	SecurityGroupBinding *string `pulumi:"securityGroupBinding"`
	// Target security group ID for this rule.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Maximum port number.
	ToPort *int `pulumi:"toPort"`
	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []string `pulumi:"v4CidrBlocks"`
	// The blocks of IPv6 addresses for this rule. `v6CidrBlocks` argument is currently not supported. It will be available in the future.
	V6CidrBlocks []string `pulumi:"v6CidrBlocks"`
}

type VpcSecurityGroupRuleState struct {
	// Description of the rule.
	Description pulumi.StringPtrInput
	// direction of the rule. Can be `ingress` (inbound) or `egress` (outbound).
	Direction pulumi.StringPtrInput
	// Minimum port number.
	FromPort pulumi.IntPtrInput
	// Labels to assign to this rule.
	Labels pulumi.StringMapInput
	// Port number (if applied to a single port).
	Port pulumi.IntPtrInput
	// Special-purpose targets such as "selfSecurityGroup". [See docs](https://cloud.yandex.com/docs/vpc/concepts/security-groups) for possible options.
	PredefinedTarget pulumi.StringPtrInput
	// One of `ANY`, `TCP`, `UDP`, `ICMP`, `IPV6_ICMP`.
	Protocol pulumi.StringPtrInput
	// ID of the security group this rule belongs to.
	SecurityGroupBinding pulumi.StringPtrInput
	// Target security group ID for this rule.
	SecurityGroupId pulumi.StringPtrInput
	// Maximum port number.
	ToPort pulumi.IntPtrInput
	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks pulumi.StringArrayInput
	// The blocks of IPv6 addresses for this rule. `v6CidrBlocks` argument is currently not supported. It will be available in the future.
	V6CidrBlocks pulumi.StringArrayInput
}

func (VpcSecurityGroupRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcSecurityGroupRuleState)(nil)).Elem()
}

type vpcSecurityGroupRuleArgs struct {
	// Description of the rule.
	Description *string `pulumi:"description"`
	// direction of the rule. Can be `ingress` (inbound) or `egress` (outbound).
	Direction string `pulumi:"direction"`
	// Minimum port number.
	FromPort *int `pulumi:"fromPort"`
	// Labels to assign to this rule.
	Labels map[string]string `pulumi:"labels"`
	// Port number (if applied to a single port).
	Port *int `pulumi:"port"`
	// Special-purpose targets such as "selfSecurityGroup". [See docs](https://cloud.yandex.com/docs/vpc/concepts/security-groups) for possible options.
	PredefinedTarget *string `pulumi:"predefinedTarget"`
	// One of `ANY`, `TCP`, `UDP`, `ICMP`, `IPV6_ICMP`.
	Protocol *string `pulumi:"protocol"`
	// ID of the security group this rule belongs to.
	SecurityGroupBinding string `pulumi:"securityGroupBinding"`
	// Target security group ID for this rule.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Maximum port number.
	ToPort *int `pulumi:"toPort"`
	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []string `pulumi:"v4CidrBlocks"`
	// The blocks of IPv6 addresses for this rule. `v6CidrBlocks` argument is currently not supported. It will be available in the future.
	V6CidrBlocks []string `pulumi:"v6CidrBlocks"`
}

// The set of arguments for constructing a VpcSecurityGroupRule resource.
type VpcSecurityGroupRuleArgs struct {
	// Description of the rule.
	Description pulumi.StringPtrInput
	// direction of the rule. Can be `ingress` (inbound) or `egress` (outbound).
	Direction pulumi.StringInput
	// Minimum port number.
	FromPort pulumi.IntPtrInput
	// Labels to assign to this rule.
	Labels pulumi.StringMapInput
	// Port number (if applied to a single port).
	Port pulumi.IntPtrInput
	// Special-purpose targets such as "selfSecurityGroup". [See docs](https://cloud.yandex.com/docs/vpc/concepts/security-groups) for possible options.
	PredefinedTarget pulumi.StringPtrInput
	// One of `ANY`, `TCP`, `UDP`, `ICMP`, `IPV6_ICMP`.
	Protocol pulumi.StringPtrInput
	// ID of the security group this rule belongs to.
	SecurityGroupBinding pulumi.StringInput
	// Target security group ID for this rule.
	SecurityGroupId pulumi.StringPtrInput
	// Maximum port number.
	ToPort pulumi.IntPtrInput
	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks pulumi.StringArrayInput
	// The blocks of IPv6 addresses for this rule. `v6CidrBlocks` argument is currently not supported. It will be available in the future.
	V6CidrBlocks pulumi.StringArrayInput
}

func (VpcSecurityGroupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcSecurityGroupRuleArgs)(nil)).Elem()
}

type VpcSecurityGroupRuleInput interface {
	pulumi.Input

	ToVpcSecurityGroupRuleOutput() VpcSecurityGroupRuleOutput
	ToVpcSecurityGroupRuleOutputWithContext(ctx context.Context) VpcSecurityGroupRuleOutput
}

func (*VpcSecurityGroupRule) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcSecurityGroupRule)(nil))
}

func (i *VpcSecurityGroupRule) ToVpcSecurityGroupRuleOutput() VpcSecurityGroupRuleOutput {
	return i.ToVpcSecurityGroupRuleOutputWithContext(context.Background())
}

func (i *VpcSecurityGroupRule) ToVpcSecurityGroupRuleOutputWithContext(ctx context.Context) VpcSecurityGroupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSecurityGroupRuleOutput)
}

func (i *VpcSecurityGroupRule) ToVpcSecurityGroupRulePtrOutput() VpcSecurityGroupRulePtrOutput {
	return i.ToVpcSecurityGroupRulePtrOutputWithContext(context.Background())
}

func (i *VpcSecurityGroupRule) ToVpcSecurityGroupRulePtrOutputWithContext(ctx context.Context) VpcSecurityGroupRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSecurityGroupRulePtrOutput)
}

type VpcSecurityGroupRulePtrInput interface {
	pulumi.Input

	ToVpcSecurityGroupRulePtrOutput() VpcSecurityGroupRulePtrOutput
	ToVpcSecurityGroupRulePtrOutputWithContext(ctx context.Context) VpcSecurityGroupRulePtrOutput
}

type vpcSecurityGroupRulePtrType VpcSecurityGroupRuleArgs

func (*vpcSecurityGroupRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcSecurityGroupRule)(nil))
}

func (i *vpcSecurityGroupRulePtrType) ToVpcSecurityGroupRulePtrOutput() VpcSecurityGroupRulePtrOutput {
	return i.ToVpcSecurityGroupRulePtrOutputWithContext(context.Background())
}

func (i *vpcSecurityGroupRulePtrType) ToVpcSecurityGroupRulePtrOutputWithContext(ctx context.Context) VpcSecurityGroupRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSecurityGroupRulePtrOutput)
}

// VpcSecurityGroupRuleArrayInput is an input type that accepts VpcSecurityGroupRuleArray and VpcSecurityGroupRuleArrayOutput values.
// You can construct a concrete instance of `VpcSecurityGroupRuleArrayInput` via:
//
//          VpcSecurityGroupRuleArray{ VpcSecurityGroupRuleArgs{...} }
type VpcSecurityGroupRuleArrayInput interface {
	pulumi.Input

	ToVpcSecurityGroupRuleArrayOutput() VpcSecurityGroupRuleArrayOutput
	ToVpcSecurityGroupRuleArrayOutputWithContext(context.Context) VpcSecurityGroupRuleArrayOutput
}

type VpcSecurityGroupRuleArray []VpcSecurityGroupRuleInput

func (VpcSecurityGroupRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcSecurityGroupRule)(nil)).Elem()
}

func (i VpcSecurityGroupRuleArray) ToVpcSecurityGroupRuleArrayOutput() VpcSecurityGroupRuleArrayOutput {
	return i.ToVpcSecurityGroupRuleArrayOutputWithContext(context.Background())
}

func (i VpcSecurityGroupRuleArray) ToVpcSecurityGroupRuleArrayOutputWithContext(ctx context.Context) VpcSecurityGroupRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSecurityGroupRuleArrayOutput)
}

// VpcSecurityGroupRuleMapInput is an input type that accepts VpcSecurityGroupRuleMap and VpcSecurityGroupRuleMapOutput values.
// You can construct a concrete instance of `VpcSecurityGroupRuleMapInput` via:
//
//          VpcSecurityGroupRuleMap{ "key": VpcSecurityGroupRuleArgs{...} }
type VpcSecurityGroupRuleMapInput interface {
	pulumi.Input

	ToVpcSecurityGroupRuleMapOutput() VpcSecurityGroupRuleMapOutput
	ToVpcSecurityGroupRuleMapOutputWithContext(context.Context) VpcSecurityGroupRuleMapOutput
}

type VpcSecurityGroupRuleMap map[string]VpcSecurityGroupRuleInput

func (VpcSecurityGroupRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcSecurityGroupRule)(nil)).Elem()
}

func (i VpcSecurityGroupRuleMap) ToVpcSecurityGroupRuleMapOutput() VpcSecurityGroupRuleMapOutput {
	return i.ToVpcSecurityGroupRuleMapOutputWithContext(context.Background())
}

func (i VpcSecurityGroupRuleMap) ToVpcSecurityGroupRuleMapOutputWithContext(ctx context.Context) VpcSecurityGroupRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSecurityGroupRuleMapOutput)
}

type VpcSecurityGroupRuleOutput struct{ *pulumi.OutputState }

func (VpcSecurityGroupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcSecurityGroupRule)(nil))
}

func (o VpcSecurityGroupRuleOutput) ToVpcSecurityGroupRuleOutput() VpcSecurityGroupRuleOutput {
	return o
}

func (o VpcSecurityGroupRuleOutput) ToVpcSecurityGroupRuleOutputWithContext(ctx context.Context) VpcSecurityGroupRuleOutput {
	return o
}

func (o VpcSecurityGroupRuleOutput) ToVpcSecurityGroupRulePtrOutput() VpcSecurityGroupRulePtrOutput {
	return o.ToVpcSecurityGroupRulePtrOutputWithContext(context.Background())
}

func (o VpcSecurityGroupRuleOutput) ToVpcSecurityGroupRulePtrOutputWithContext(ctx context.Context) VpcSecurityGroupRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpcSecurityGroupRule) *VpcSecurityGroupRule {
		return &v
	}).(VpcSecurityGroupRulePtrOutput)
}

type VpcSecurityGroupRulePtrOutput struct{ *pulumi.OutputState }

func (VpcSecurityGroupRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcSecurityGroupRule)(nil))
}

func (o VpcSecurityGroupRulePtrOutput) ToVpcSecurityGroupRulePtrOutput() VpcSecurityGroupRulePtrOutput {
	return o
}

func (o VpcSecurityGroupRulePtrOutput) ToVpcSecurityGroupRulePtrOutputWithContext(ctx context.Context) VpcSecurityGroupRulePtrOutput {
	return o
}

func (o VpcSecurityGroupRulePtrOutput) Elem() VpcSecurityGroupRuleOutput {
	return o.ApplyT(func(v *VpcSecurityGroupRule) VpcSecurityGroupRule {
		if v != nil {
			return *v
		}
		var ret VpcSecurityGroupRule
		return ret
	}).(VpcSecurityGroupRuleOutput)
}

type VpcSecurityGroupRuleArrayOutput struct{ *pulumi.OutputState }

func (VpcSecurityGroupRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcSecurityGroupRule)(nil))
}

func (o VpcSecurityGroupRuleArrayOutput) ToVpcSecurityGroupRuleArrayOutput() VpcSecurityGroupRuleArrayOutput {
	return o
}

func (o VpcSecurityGroupRuleArrayOutput) ToVpcSecurityGroupRuleArrayOutputWithContext(ctx context.Context) VpcSecurityGroupRuleArrayOutput {
	return o
}

func (o VpcSecurityGroupRuleArrayOutput) Index(i pulumi.IntInput) VpcSecurityGroupRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcSecurityGroupRule {
		return vs[0].([]VpcSecurityGroupRule)[vs[1].(int)]
	}).(VpcSecurityGroupRuleOutput)
}

type VpcSecurityGroupRuleMapOutput struct{ *pulumi.OutputState }

func (VpcSecurityGroupRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcSecurityGroupRule)(nil))
}

func (o VpcSecurityGroupRuleMapOutput) ToVpcSecurityGroupRuleMapOutput() VpcSecurityGroupRuleMapOutput {
	return o
}

func (o VpcSecurityGroupRuleMapOutput) ToVpcSecurityGroupRuleMapOutputWithContext(ctx context.Context) VpcSecurityGroupRuleMapOutput {
	return o
}

func (o VpcSecurityGroupRuleMapOutput) MapIndex(k pulumi.StringInput) VpcSecurityGroupRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcSecurityGroupRule {
		return vs[0].(map[string]VpcSecurityGroupRule)[vs[1].(string)]
	}).(VpcSecurityGroupRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSecurityGroupRuleInput)(nil)).Elem(), &VpcSecurityGroupRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSecurityGroupRulePtrInput)(nil)).Elem(), &VpcSecurityGroupRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSecurityGroupRuleArrayInput)(nil)).Elem(), VpcSecurityGroupRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSecurityGroupRuleMapInput)(nil)).Elem(), VpcSecurityGroupRuleMap{})
	pulumi.RegisterOutputType(VpcSecurityGroupRuleOutput{})
	pulumi.RegisterOutputType(VpcSecurityGroupRulePtrOutput{})
	pulumi.RegisterOutputType(VpcSecurityGroupRuleArrayOutput{})
	pulumi.RegisterOutputType(VpcSecurityGroupRuleMapOutput{})
}
