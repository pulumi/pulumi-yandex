// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a backend group in the specified folder and adds the specified backends to it.
 * For more information, see [the official documentation](https://cloud.yandex.com/en/docs/application-load-balancer/concepts/backend-group).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const test_backend_group = new yandex.AlbBackendGroup("test-backend-group", {
 *     httpBackends: [{
 *         healthcheck: {
 *             httpHealthcheck: {
 *                 path: "/",
 *             },
 *             interval: "1s",
 *             timeout: "1s",
 *         },
 *         http2: true,
 *         loadBalancingConfig: {
 *             panicThreshold: 50,
 *         },
 *         name: "test-http-backend",
 *         port: 8080,
 *         targetGroupIds: [yandex_alb_target_group_test_target_group.id],
 *         tls: {
 *             sni: "backend-domain.internal",
 *         },
 *         weight: 1,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * A backend group can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/albBackendGroup:AlbBackendGroup default backend_group_id
 * ```
 */
export class AlbBackendGroup extends pulumi.CustomResource {
    /**
     * Get an existing AlbBackendGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlbBackendGroupState, opts?: pulumi.CustomResourceOptions): AlbBackendGroup {
        return new AlbBackendGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/albBackendGroup:AlbBackendGroup';

    /**
     * Returns true if the given object is an instance of AlbBackendGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlbBackendGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlbBackendGroup.__pulumiType;
    }

    /**
     * The backend group creation timestamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Description of the backend group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    public readonly grpcBackends!: pulumi.Output<outputs.AlbBackendGroupGrpcBackend[] | undefined>;
    /**
     * Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    public readonly httpBackends!: pulumi.Output<outputs.AlbBackendGroupHttpBackend[] | undefined>;
    /**
     * Labels to assign to this backend group.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the backend.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a AlbBackendGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AlbBackendGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlbBackendGroupArgs | AlbBackendGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlbBackendGroupState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["grpcBackends"] = state ? state.grpcBackends : undefined;
            inputs["httpBackends"] = state ? state.httpBackends : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as AlbBackendGroupArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["grpcBackends"] = args ? args.grpcBackends : undefined;
            inputs["httpBackends"] = args ? args.httpBackends : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["createdAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlbBackendGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlbBackendGroup resources.
 */
export interface AlbBackendGroupState {
    /**
     * The backend group creation timestamp.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Description of the backend group.
     */
    description?: pulumi.Input<string>;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    grpcBackends?: pulumi.Input<pulumi.Input<inputs.AlbBackendGroupGrpcBackend>[]>;
    /**
     * Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    httpBackends?: pulumi.Input<pulumi.Input<inputs.AlbBackendGroupHttpBackend>[]>;
    /**
     * Labels to assign to this backend group.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the backend.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlbBackendGroup resource.
 */
export interface AlbBackendGroupArgs {
    /**
     * Description of the backend group.
     */
    description?: pulumi.Input<string>;
    /**
     * Folder that the resource belongs to. If value is omitted, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    grpcBackends?: pulumi.Input<pulumi.Input<inputs.AlbBackendGroupGrpcBackend>[]>;
    /**
     * Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
     */
    httpBackends?: pulumi.Input<pulumi.Input<inputs.AlbBackendGroupHttpBackend>[]>;
    /**
     * Labels to assign to this backend group.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the backend.
     */
    name?: pulumi.Input<string>;
}
