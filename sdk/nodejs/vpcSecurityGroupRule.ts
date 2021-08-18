// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a single Secuirity Group Rule within the Yandex.Cloud. For more information, see the official documentation
 * of [security groups](https://cloud.yandex.ru/docs/vpc/concepts/security-groups)
 * and [security group rules](https://cloud.yandex.ru/docs/vpc/concepts/security-groups#rules).
 *
 * > **NOTE:** There is another way to manage security group rules by `ingress` and `egress` arguments in yandex_vpc_security_group. Both ways are equivalent but not compatible now. Using in-line rules of yandex.VpcSecurityGroup with Security Group Rule resource at the same time will cause a conflict of rules configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const lab_net = new yandex.VpcNetwork("lab-net", {});
 * const group1 = new yandex.VpcSecurityGroup("group1", {
 *     description: "description for my security group",
 *     networkId: lab_net.id,
 *     labels: {
 *         "my-label": "my-label-value",
 *     },
 * });
 * const rule1 = new yandex.VpcSecurityGroupRule("rule1", {
 *     securityGroupBinding: group1.id,
 *     direction: "ingress",
 *     description: "rule1 description",
 *     v4CidrBlokcs: [
 *         "10.0.1.0/24",
 *         "10.0.2.0/24",
 *     ],
 *     port: 8080,
 *     protocol: "TCP",
 * });
 * const rule2 = new yandex.VpcSecurityGroupRule("rule2", {
 *     securityGroupBinding: group1.id,
 *     direction: "egress",
 *     description: "rule2 description",
 *     v4CidrBlokcs: ["10.0.1.0/24"],
 *     fromPort: 8090,
 *     toPort: 8099,
 *     protocol: "UDP",
 * });
 * ```
 */
export class VpcSecurityGroupRule extends pulumi.CustomResource {
    /**
     * Get an existing VpcSecurityGroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcSecurityGroupRuleState, opts?: pulumi.CustomResourceOptions): VpcSecurityGroupRule {
        return new VpcSecurityGroupRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/vpcSecurityGroupRule:VpcSecurityGroupRule';

    /**
     * Returns true if the given object is an instance of VpcSecurityGroupRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcSecurityGroupRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcSecurityGroupRule.__pulumiType;
    }

    /**
     * Description of the rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * direction of the rule. Can be `ingress` (inbound) or `egress` (outbound).
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * Minimum port number.
     */
    public readonly fromPort!: pulumi.Output<number | undefined>;
    /**
     * Labels to assign to this rule.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Port number (if applied to a single port).
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * Special-purpose targets such as "selfSecurityGroup". [See docs](https://cloud.yandex.ru/docs/vpc/concepts/security-groups) for possible options.
     */
    public readonly predefinedTarget!: pulumi.Output<string | undefined>;
    /**
     * One of `ANY`, `TCP`, `UDP`, `ICMP`, `IPV6_ICMP`.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * ID of the security group this rule belongs to.
     */
    public readonly securityGroupBinding!: pulumi.Output<string>;
    /**
     * Target security group ID for this rule.
     */
    public readonly securityGroupId!: pulumi.Output<string | undefined>;
    /**
     * Maximum port number.
     */
    public readonly toPort!: pulumi.Output<number | undefined>;
    /**
     * The blocks of IPv4 addresses for this rule.
     */
    public readonly v4CidrBlocks!: pulumi.Output<string[] | undefined>;
    /**
     * The blocks of IPv6 addresses for this rule. `v6CidrBlocks` argument is currently not supported. It will be available in the future.
     */
    public readonly v6CidrBlocks!: pulumi.Output<string[] | undefined>;

    /**
     * Create a VpcSecurityGroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcSecurityGroupRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcSecurityGroupRuleArgs | VpcSecurityGroupRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcSecurityGroupRuleState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["direction"] = state ? state.direction : undefined;
            inputs["fromPort"] = state ? state.fromPort : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["predefinedTarget"] = state ? state.predefinedTarget : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["securityGroupBinding"] = state ? state.securityGroupBinding : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["toPort"] = state ? state.toPort : undefined;
            inputs["v4CidrBlocks"] = state ? state.v4CidrBlocks : undefined;
            inputs["v6CidrBlocks"] = state ? state.v6CidrBlocks : undefined;
        } else {
            const args = argsOrState as VpcSecurityGroupRuleArgs | undefined;
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.securityGroupBinding === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupBinding'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["fromPort"] = args ? args.fromPort : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["predefinedTarget"] = args ? args.predefinedTarget : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["securityGroupBinding"] = args ? args.securityGroupBinding : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["toPort"] = args ? args.toPort : undefined;
            inputs["v4CidrBlocks"] = args ? args.v4CidrBlocks : undefined;
            inputs["v6CidrBlocks"] = args ? args.v6CidrBlocks : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VpcSecurityGroupRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcSecurityGroupRule resources.
 */
export interface VpcSecurityGroupRuleState {
    /**
     * Description of the rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * direction of the rule. Can be `ingress` (inbound) or `egress` (outbound).
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * Minimum port number.
     */
    readonly fromPort?: pulumi.Input<number>;
    /**
     * Labels to assign to this rule.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Port number (if applied to a single port).
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Special-purpose targets such as "selfSecurityGroup". [See docs](https://cloud.yandex.ru/docs/vpc/concepts/security-groups) for possible options.
     */
    readonly predefinedTarget?: pulumi.Input<string>;
    /**
     * One of `ANY`, `TCP`, `UDP`, `ICMP`, `IPV6_ICMP`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * ID of the security group this rule belongs to.
     */
    readonly securityGroupBinding?: pulumi.Input<string>;
    /**
     * Target security group ID for this rule.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * Maximum port number.
     */
    readonly toPort?: pulumi.Input<number>;
    /**
     * The blocks of IPv4 addresses for this rule.
     */
    readonly v4CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The blocks of IPv6 addresses for this rule. `v6CidrBlocks` argument is currently not supported. It will be available in the future.
     */
    readonly v6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a VpcSecurityGroupRule resource.
 */
export interface VpcSecurityGroupRuleArgs {
    /**
     * Description of the rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * direction of the rule. Can be `ingress` (inbound) or `egress` (outbound).
     */
    readonly direction: pulumi.Input<string>;
    /**
     * Minimum port number.
     */
    readonly fromPort?: pulumi.Input<number>;
    /**
     * Labels to assign to this rule.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Port number (if applied to a single port).
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Special-purpose targets such as "selfSecurityGroup". [See docs](https://cloud.yandex.ru/docs/vpc/concepts/security-groups) for possible options.
     */
    readonly predefinedTarget?: pulumi.Input<string>;
    /**
     * One of `ANY`, `TCP`, `UDP`, `ICMP`, `IPV6_ICMP`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * ID of the security group this rule belongs to.
     */
    readonly securityGroupBinding: pulumi.Input<string>;
    /**
     * Target security group ID for this rule.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * Maximum port number.
     */
    readonly toPort?: pulumi.Input<number>;
    /**
     * The blocks of IPv4 addresses for this rule.
     */
    readonly v4CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The blocks of IPv6 addresses for this rule. `v6CidrBlocks` argument is currently not supported. It will be available in the future.
     */
    readonly v6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
}
