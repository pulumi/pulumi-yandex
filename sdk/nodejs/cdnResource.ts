// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Allows management of [Yandex.Cloud CDN Resource](https://cloud.yandex.ru/docs/cdn/concepts/resource).
 *
 * > **_NOTE:_**  CDN provider must be activated prior usage of CDN resources, either via UI console or via yc cli command: ```yc cdn provider activate --folder-id <folder-id> --type gcore```
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const myResource = new yandex.CdnResource("myResource", {
 *     cname: "cdn1.yandex-example.ru",
 *     active: false,
 *     originProtocol: "https",
 *     secondaryHostnames: [
 *         "cdn-example-1.yandex.ru",
 *         "cdn-example-2.yandex.ru",
 *     ],
 *     originGroupId: yandex_cdn_origin_group.foo_cdn_group_by_id.id,
 *     options: {
 *         edgeCacheSettings: 345600,
 *         ignoreCookie: true,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * A origin group can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import yandex:index/cdnResource:CdnResource default origin_group_id
 * ```
 */
export class CdnResource extends pulumi.CustomResource {
    /**
     * Get an existing CdnResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CdnResourceState, opts?: pulumi.CustomResourceOptions): CdnResource {
        return new CdnResource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/cdnResource:CdnResource';

    /**
     * Returns true if the given object is an instance of CdnResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CdnResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CdnResource.__pulumiType;
    }

    /**
     * Flag to create Resource either in active or disabled state. True - the content from CDN is available to clients.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * CDN endpoint CNAME, must be unique among resources.
     */
    public readonly cname!: pulumi.Output<string>;
    /**
     * Creation timestamp of the IoT Core Device
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public /*out*/ readonly folderId!: pulumi.Output<string>;
    /**
     * CDN Resource settings and options to tune CDN edge behavior.
     */
    public readonly options!: pulumi.Output<outputs.CdnResourceOptions>;
    public readonly originGroupId!: pulumi.Output<number | undefined>;
    public readonly originGroupName!: pulumi.Output<string | undefined>;
    public readonly originProtocol!: pulumi.Output<string | undefined>;
    /**
     * list of secondary hostname strings.
     */
    public readonly secondaryHostnames!: pulumi.Output<string[] | undefined>;
    public readonly sslCertificate!: pulumi.Output<outputs.CdnResourceSslCertificate | undefined>;
    public readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a CdnResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CdnResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CdnResourceArgs | CdnResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CdnResourceState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["cname"] = state ? state.cname : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["originGroupId"] = state ? state.originGroupId : undefined;
            resourceInputs["originGroupName"] = state ? state.originGroupName : undefined;
            resourceInputs["originProtocol"] = state ? state.originProtocol : undefined;
            resourceInputs["secondaryHostnames"] = state ? state.secondaryHostnames : undefined;
            resourceInputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as CdnResourceArgs | undefined;
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["cname"] = args ? args.cname : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["originGroupId"] = args ? args.originGroupId : undefined;
            resourceInputs["originGroupName"] = args ? args.originGroupName : undefined;
            resourceInputs["originProtocol"] = args ? args.originProtocol : undefined;
            resourceInputs["secondaryHostnames"] = args ? args.secondaryHostnames : undefined;
            resourceInputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            resourceInputs["updatedAt"] = args ? args.updatedAt : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["folderId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CdnResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CdnResource resources.
 */
export interface CdnResourceState {
    /**
     * Flag to create Resource either in active or disabled state. True - the content from CDN is available to clients.
     */
    active?: pulumi.Input<boolean>;
    /**
     * CDN endpoint CNAME, must be unique among resources.
     */
    cname?: pulumi.Input<string>;
    /**
     * Creation timestamp of the IoT Core Device
     */
    createdAt?: pulumi.Input<string>;
    folderId?: pulumi.Input<string>;
    /**
     * CDN Resource settings and options to tune CDN edge behavior.
     */
    options?: pulumi.Input<inputs.CdnResourceOptions>;
    originGroupId?: pulumi.Input<number>;
    originGroupName?: pulumi.Input<string>;
    originProtocol?: pulumi.Input<string>;
    /**
     * list of secondary hostname strings.
     */
    secondaryHostnames?: pulumi.Input<pulumi.Input<string>[]>;
    sslCertificate?: pulumi.Input<inputs.CdnResourceSslCertificate>;
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CdnResource resource.
 */
export interface CdnResourceArgs {
    /**
     * Flag to create Resource either in active or disabled state. True - the content from CDN is available to clients.
     */
    active?: pulumi.Input<boolean>;
    /**
     * CDN endpoint CNAME, must be unique among resources.
     */
    cname?: pulumi.Input<string>;
    /**
     * CDN Resource settings and options to tune CDN edge behavior.
     */
    options?: pulumi.Input<inputs.CdnResourceOptions>;
    originGroupId?: pulumi.Input<number>;
    originGroupName?: pulumi.Input<string>;
    originProtocol?: pulumi.Input<string>;
    /**
     * list of secondary hostname strings.
     */
    secondaryHostnames?: pulumi.Input<pulumi.Input<string>[]>;
    sslCertificate?: pulumi.Input<inputs.CdnResourceSslCertificate>;
    updatedAt?: pulumi.Input<string>;
}
