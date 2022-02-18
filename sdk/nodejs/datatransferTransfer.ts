// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Data Transfer transfer. For more information, see [the official documentation](https://cloud.yandex.com/docs/data-transfer/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const pgSource = new yandex.DatatransferEndpoint("pgSource", {settings: {
 *     postgresSource: {
 *         connection: {
 *             onPremise: {
 *                 hosts: ["example.org"],
 *                 port: 5432,
 *             },
 *         },
 *         slotGigabyteLagLimit: 100,
 *         database: "db1",
 *         user: "user1",
 *         password: {
 *             raw: "123",
 *         },
 *     },
 * }});
 * const pgTarget = new yandex.DatatransferEndpoint("pgTarget", {
 *     folderId: "some_folder_id",
 *     settings: {
 *         postgresTarget: {
 *             connection: {
 *                 mdbClusterId: "some_cluster_id",
 *             },
 *             database: "db2",
 *             user: "user2",
 *             password: {
 *                 raw: "321",
 *             },
 *         },
 *     },
 * });
 * const pgpgTransfer = new yandex.DatatransferTransfer("pgpgTransfer", {
 *     folderId: "some_folder_id",
 *     sourceId: pgSource.id,
 *     targetId: pgTarget.id,
 *     type: "SNAPSHOT_AND_INCREMENT",
 * });
 * ```
 *
 * ## Import
 *
 * A transfer can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/datatransferTransfer:DatatransferTransfer foo transfer_id
 * ```
 */
export class DatatransferTransfer extends pulumi.CustomResource {
    /**
     * Get an existing DatatransferTransfer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatatransferTransferState, opts?: pulumi.CustomResourceOptions): DatatransferTransfer {
        return new DatatransferTransfer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/datatransferTransfer:DatatransferTransfer';

    /**
     * Returns true if the given object is an instance of DatatransferTransfer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatatransferTransfer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatatransferTransfer.__pulumiType;
    }

    /**
     * Arbitrary description text for the transfer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ID of the folder to create the transfer in. If it is not provided, the default provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the Data Transfer transfer.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the transfer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the source endpoint for the transfer.
     */
    public readonly sourceId!: pulumi.Output<string | undefined>;
    /**
     * ID of the target endpoint for the transfer.
     */
    public readonly targetId!: pulumi.Output<string | undefined>;
    /**
     * Type of the transfer. One of "SNAPSHOT_ONLY", "INCREMENT_ONLY", "SNAPSHOT_AND_INCREMENT".
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * (Computed) Error description if transfer has any errors.
     */
    public /*out*/ readonly warning!: pulumi.Output<string>;

    /**
     * Create a DatatransferTransfer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatatransferTransferArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatatransferTransferArgs | DatatransferTransferState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatatransferTransferState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sourceId"] = state ? state.sourceId : undefined;
            resourceInputs["targetId"] = state ? state.targetId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["warning"] = state ? state.warning : undefined;
        } else {
            const args = argsOrState as DatatransferTransferArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sourceId"] = args ? args.sourceId : undefined;
            resourceInputs["targetId"] = args ? args.targetId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["warning"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatatransferTransfer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatatransferTransfer resources.
 */
export interface DatatransferTransferState {
    /**
     * Arbitrary description text for the transfer.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the folder to create the transfer in. If it is not provided, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Data Transfer transfer.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the transfer.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the source endpoint for the transfer.
     */
    sourceId?: pulumi.Input<string>;
    /**
     * ID of the target endpoint for the transfer.
     */
    targetId?: pulumi.Input<string>;
    /**
     * Type of the transfer. One of "SNAPSHOT_ONLY", "INCREMENT_ONLY", "SNAPSHOT_AND_INCREMENT".
     */
    type?: pulumi.Input<string>;
    /**
     * (Computed) Error description if transfer has any errors.
     */
    warning?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatatransferTransfer resource.
 */
export interface DatatransferTransferArgs {
    /**
     * Arbitrary description text for the transfer.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the folder to create the transfer in. If it is not provided, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Data Transfer transfer.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the transfer.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the source endpoint for the transfer.
     */
    sourceId?: pulumi.Input<string>;
    /**
     * ID of the target endpoint for the transfer.
     */
    targetId?: pulumi.Input<string>;
    /**
     * Type of the transfer. One of "SNAPSHOT_ONLY", "INCREMENT_ONLY", "SNAPSHOT_AND_INCREMENT".
     */
    type?: pulumi.Input<string>;
}