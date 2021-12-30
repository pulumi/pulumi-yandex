// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Yandex Database (dedicated) resource.
 * For more information, see [the official documentation](https://cloud.yandex.com/en/docs/ydb/concepts/serverless_and_dedicated).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const database1 = new yandex.YdbDatabaseDedicated("database1", {
 *     folderId: yandex_resourcemanager_folder_test_folder.id,
 *     location: {
 *         region: {
 *             id: "ru-central1",
 *         },
 *     },
 *     networkId: yandex_vpc_network_my_inst_group_network.id,
 *     resourcePresetId: "medium",
 *     scalePolicy: {
 *         fixedScale: {
 *             size: 1,
 *         },
 *     },
 *     storageConfig: {
 *         groupCount: 1,
 *         storageTypeId: "ssd",
 *     },
 *     subnetIds: [yandex_vpc_subnet_my_inst_group_subnet.id],
 * });
 * ```
 */
export class YdbDatabaseDedicated extends pulumi.CustomResource {
    /**
     * Get an existing YdbDatabaseDedicated resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: YdbDatabaseDedicatedState, opts?: pulumi.CustomResourceOptions): YdbDatabaseDedicated {
        return new YdbDatabaseDedicated(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/ydbDatabaseDedicated:YdbDatabaseDedicated';

    /**
     * Returns true if the given object is an instance of YdbDatabaseDedicated.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is YdbDatabaseDedicated {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === YdbDatabaseDedicated.__pulumiType;
    }

    /**
     * Whether public IP addresses should be assigned to the Yandex Database cluster.
     */
    public readonly assignPublicIps!: pulumi.Output<boolean | undefined>;
    /**
     * The Yandex Database cluster creation timestamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Full database path of the Yandex Database cluster.
     * Useful for SDK configuration.
     */
    public /*out*/ readonly databasePath!: pulumi.Output<string>;
    /**
     * A description for the Yandex Database cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ID of the folder that the Yandex Database cluster belongs to.
     * It will be deduced from provider configuration if not set explicitly.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the Yandex Database cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Location for the Yandex Database cluster.
     * The structure is documented below.
     */
    public readonly location!: pulumi.Output<outputs.YdbDatabaseDedicatedLocation | undefined>;
    /**
     * Location ID for the Yandex Database cluster.
     */
    public readonly locationId!: pulumi.Output<string>;
    /**
     * Name of the Yandex Database cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network to attach the Yandex Database cluster to.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * The Yandex Database cluster preset.
     * Available presets can be obtained via `yc ydb resource-preset list` command.
     */
    public readonly resourcePresetId!: pulumi.Output<string>;
    /**
     * Scaling policy for the Yandex Database cluster.
     * The structure is documented below.
     */
    public readonly scalePolicy!: pulumi.Output<outputs.YdbDatabaseDedicatedScalePolicy>;
    /**
     * Status of the Yandex Database cluster.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A list of storage configuration options for the Yandex Database cluster.
     * The structure is documented below.
     */
    public readonly storageConfig!: pulumi.Output<outputs.YdbDatabaseDedicatedStorageConfig>;
    /**
     * List of subnet IDs to attach the Yandex Database cluster to.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * Whether TLS is enabled for the Yandex Database cluster.
     * Useful for SDK configuration.
     */
    public /*out*/ readonly tlsEnabled!: pulumi.Output<boolean>;
    /**
     * API endpoint of the Yandex Database cluster.
     * Useful for SDK configuration.
     */
    public /*out*/ readonly ydbApiEndpoint!: pulumi.Output<string>;
    /**
     * Full endpoint of the Yandex Database cluster.
     */
    public /*out*/ readonly ydbFullEndpoint!: pulumi.Output<string>;

    /**
     * Create a YdbDatabaseDedicated resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: YdbDatabaseDedicatedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: YdbDatabaseDedicatedArgs | YdbDatabaseDedicatedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as YdbDatabaseDedicatedState | undefined;
            resourceInputs["assignPublicIps"] = state ? state.assignPublicIps : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["databasePath"] = state ? state.databasePath : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["locationId"] = state ? state.locationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["resourcePresetId"] = state ? state.resourcePresetId : undefined;
            resourceInputs["scalePolicy"] = state ? state.scalePolicy : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storageConfig"] = state ? state.storageConfig : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["tlsEnabled"] = state ? state.tlsEnabled : undefined;
            resourceInputs["ydbApiEndpoint"] = state ? state.ydbApiEndpoint : undefined;
            resourceInputs["ydbFullEndpoint"] = state ? state.ydbFullEndpoint : undefined;
        } else {
            const args = argsOrState as YdbDatabaseDedicatedArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.resourcePresetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourcePresetId'");
            }
            if ((!args || args.scalePolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalePolicy'");
            }
            if ((!args || args.storageConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageConfig'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            resourceInputs["assignPublicIps"] = args ? args.assignPublicIps : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["locationId"] = args ? args.locationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["resourcePresetId"] = args ? args.resourcePresetId : undefined;
            resourceInputs["scalePolicy"] = args ? args.scalePolicy : undefined;
            resourceInputs["storageConfig"] = args ? args.storageConfig : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["databasePath"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tlsEnabled"] = undefined /*out*/;
            resourceInputs["ydbApiEndpoint"] = undefined /*out*/;
            resourceInputs["ydbFullEndpoint"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(YdbDatabaseDedicated.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering YdbDatabaseDedicated resources.
 */
export interface YdbDatabaseDedicatedState {
    /**
     * Whether public IP addresses should be assigned to the Yandex Database cluster.
     */
    assignPublicIps?: pulumi.Input<boolean>;
    /**
     * The Yandex Database cluster creation timestamp.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Full database path of the Yandex Database cluster.
     * Useful for SDK configuration.
     */
    databasePath?: pulumi.Input<string>;
    /**
     * A description for the Yandex Database cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the folder that the Yandex Database cluster belongs to.
     * It will be deduced from provider configuration if not set explicitly.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Yandex Database cluster.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location for the Yandex Database cluster.
     * The structure is documented below.
     */
    location?: pulumi.Input<inputs.YdbDatabaseDedicatedLocation>;
    /**
     * Location ID for the Yandex Database cluster.
     */
    locationId?: pulumi.Input<string>;
    /**
     * Name of the Yandex Database cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network to attach the Yandex Database cluster to.
     */
    networkId?: pulumi.Input<string>;
    /**
     * The Yandex Database cluster preset.
     * Available presets can be obtained via `yc ydb resource-preset list` command.
     */
    resourcePresetId?: pulumi.Input<string>;
    /**
     * Scaling policy for the Yandex Database cluster.
     * The structure is documented below.
     */
    scalePolicy?: pulumi.Input<inputs.YdbDatabaseDedicatedScalePolicy>;
    /**
     * Status of the Yandex Database cluster.
     */
    status?: pulumi.Input<string>;
    /**
     * A list of storage configuration options for the Yandex Database cluster.
     * The structure is documented below.
     */
    storageConfig?: pulumi.Input<inputs.YdbDatabaseDedicatedStorageConfig>;
    /**
     * List of subnet IDs to attach the Yandex Database cluster to.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether TLS is enabled for the Yandex Database cluster.
     * Useful for SDK configuration.
     */
    tlsEnabled?: pulumi.Input<boolean>;
    /**
     * API endpoint of the Yandex Database cluster.
     * Useful for SDK configuration.
     */
    ydbApiEndpoint?: pulumi.Input<string>;
    /**
     * Full endpoint of the Yandex Database cluster.
     */
    ydbFullEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a YdbDatabaseDedicated resource.
 */
export interface YdbDatabaseDedicatedArgs {
    /**
     * Whether public IP addresses should be assigned to the Yandex Database cluster.
     */
    assignPublicIps?: pulumi.Input<boolean>;
    /**
     * A description for the Yandex Database cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the folder that the Yandex Database cluster belongs to.
     * It will be deduced from provider configuration if not set explicitly.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Yandex Database cluster.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location for the Yandex Database cluster.
     * The structure is documented below.
     */
    location?: pulumi.Input<inputs.YdbDatabaseDedicatedLocation>;
    /**
     * Location ID for the Yandex Database cluster.
     */
    locationId?: pulumi.Input<string>;
    /**
     * Name of the Yandex Database cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network to attach the Yandex Database cluster to.
     */
    networkId: pulumi.Input<string>;
    /**
     * The Yandex Database cluster preset.
     * Available presets can be obtained via `yc ydb resource-preset list` command.
     */
    resourcePresetId: pulumi.Input<string>;
    /**
     * Scaling policy for the Yandex Database cluster.
     * The structure is documented below.
     */
    scalePolicy: pulumi.Input<inputs.YdbDatabaseDedicatedScalePolicy>;
    /**
     * A list of storage configuration options for the Yandex Database cluster.
     * The structure is documented below.
     */
    storageConfig: pulumi.Input<inputs.YdbDatabaseDedicatedStorageConfig>;
    /**
     * List of subnet IDs to attach the Yandex Database cluster to.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
}
