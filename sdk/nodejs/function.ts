// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Allows management of [Yandex Cloud Function](https://cloud.yandex.com/docs/functions/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const test_function = new yandex.Function("test-function", {
 *     content: {
 *         zipFilename: "function.zip",
 *     },
 *     description: "any description",
 *     entrypoint: "main",
 *     executionTimeout: "10",
 *     memory: 128,
 *     runtime: "python37",
 *     serviceAccountId: "are1service2account3id",
 *     tags: ["my_tag"],
 *     userHash: "any_user_defined_string",
 * });
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * Version deployment content for Yandex Cloud Function code. Can be only one `package` or `content` section.
     * * `content.0.zip_filename` - Filename to zip archive for the version.
     */
    public readonly content!: pulumi.Output<outputs.FunctionContent | undefined>;
    /**
     * Creation timestamp of the Yandex Cloud Function.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Description of the Yandex Cloud Function
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Entrypoint for Yandex Cloud Function
     */
    public readonly entrypoint!: pulumi.Output<string>;
    /**
     * A set of key/value environment variables for Yandex Cloud Function
     */
    public readonly environment!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Execution timeout in seconds for Yandex Cloud Function
     */
    public readonly executionTimeout!: pulumi.Output<string | undefined>;
    /**
     * Folder ID for the Yandex Cloud Function
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Image size for Yandex Cloud Function.
     */
    public /*out*/ readonly imageSize!: pulumi.Output<number>;
    /**
     * A set of key/value label pairs to assign to the Yandex Cloud Function
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Loggroup ID size for Yandex Cloud Function.
     */
    public /*out*/ readonly loggroupId!: pulumi.Output<string>;
    /**
     * Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Function
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * Yandex Cloud Function name used to define trigger
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Version deployment package for Yandex Cloud Function code. Can be only one `package` or `content` section.
     * * `package.0.sha_256` - SHA256 hash of the version deployment package.
     * * `package.0.bucket_name` - Name of the bucket that stores the code for the version.
     * * `package.0.object_name` - Name of the object in the bucket that stores the code for the version.
     */
    public readonly package!: pulumi.Output<outputs.FunctionPackage | undefined>;
    /**
     * Runtime for Yandex Cloud Function
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * Service account ID for Yandex Cloud Function
     */
    public readonly serviceAccountId!: pulumi.Output<string | undefined>;
    /**
     * Tags for Yandex Cloud Function. Tag "$latest" isn't returned.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * User-defined string for current function version. User must change this string any times when function changed. Function will be updated when hash is changed.
     */
    public readonly userHash!: pulumi.Output<string>;
    /**
     * Version for Yandex Cloud Function.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["entrypoint"] = state ? state.entrypoint : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["executionTimeout"] = state ? state.executionTimeout : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["imageSize"] = state ? state.imageSize : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["loggroupId"] = state ? state.loggroupId : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["package"] = state ? state.package : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["userHash"] = state ? state.userHash : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.entrypoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entrypoint'");
            }
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            if ((!args || args.userHash === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userHash'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["entrypoint"] = args ? args.entrypoint : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["executionTimeout"] = args ? args.executionTimeout : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["package"] = args ? args.package : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userHash"] = args ? args.userHash : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["imageSize"] = undefined /*out*/;
            resourceInputs["loggroupId"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Function.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * Version deployment content for Yandex Cloud Function code. Can be only one `package` or `content` section.
     * * `content.0.zip_filename` - Filename to zip archive for the version.
     */
    content?: pulumi.Input<inputs.FunctionContent>;
    /**
     * Creation timestamp of the Yandex Cloud Function.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Description of the Yandex Cloud Function
     */
    description?: pulumi.Input<string>;
    /**
     * Entrypoint for Yandex Cloud Function
     */
    entrypoint?: pulumi.Input<string>;
    /**
     * A set of key/value environment variables for Yandex Cloud Function
     */
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Execution timeout in seconds for Yandex Cloud Function
     */
    executionTimeout?: pulumi.Input<string>;
    /**
     * Folder ID for the Yandex Cloud Function
     */
    folderId?: pulumi.Input<string>;
    /**
     * Image size for Yandex Cloud Function.
     */
    imageSize?: pulumi.Input<number>;
    /**
     * A set of key/value label pairs to assign to the Yandex Cloud Function
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Loggroup ID size for Yandex Cloud Function.
     */
    loggroupId?: pulumi.Input<string>;
    /**
     * Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Function
     */
    memory?: pulumi.Input<number>;
    /**
     * Yandex Cloud Function name used to define trigger
     */
    name?: pulumi.Input<string>;
    /**
     * Version deployment package for Yandex Cloud Function code. Can be only one `package` or `content` section.
     * * `package.0.sha_256` - SHA256 hash of the version deployment package.
     * * `package.0.bucket_name` - Name of the bucket that stores the code for the version.
     * * `package.0.object_name` - Name of the object in the bucket that stores the code for the version.
     */
    package?: pulumi.Input<inputs.FunctionPackage>;
    /**
     * Runtime for Yandex Cloud Function
     */
    runtime?: pulumi.Input<string>;
    /**
     * Service account ID for Yandex Cloud Function
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * Tags for Yandex Cloud Function. Tag "$latest" isn't returned.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User-defined string for current function version. User must change this string any times when function changed. Function will be updated when hash is changed.
     */
    userHash?: pulumi.Input<string>;
    /**
     * Version for Yandex Cloud Function.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * Version deployment content for Yandex Cloud Function code. Can be only one `package` or `content` section.
     * * `content.0.zip_filename` - Filename to zip archive for the version.
     */
    content?: pulumi.Input<inputs.FunctionContent>;
    /**
     * Description of the Yandex Cloud Function
     */
    description?: pulumi.Input<string>;
    /**
     * Entrypoint for Yandex Cloud Function
     */
    entrypoint: pulumi.Input<string>;
    /**
     * A set of key/value environment variables for Yandex Cloud Function
     */
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Execution timeout in seconds for Yandex Cloud Function
     */
    executionTimeout?: pulumi.Input<string>;
    /**
     * Folder ID for the Yandex Cloud Function
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Yandex Cloud Function
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Function
     */
    memory: pulumi.Input<number>;
    /**
     * Yandex Cloud Function name used to define trigger
     */
    name?: pulumi.Input<string>;
    /**
     * Version deployment package for Yandex Cloud Function code. Can be only one `package` or `content` section.
     * * `package.0.sha_256` - SHA256 hash of the version deployment package.
     * * `package.0.bucket_name` - Name of the bucket that stores the code for the version.
     * * `package.0.object_name` - Name of the object in the bucket that stores the code for the version.
     */
    package?: pulumi.Input<inputs.FunctionPackage>;
    /**
     * Runtime for Yandex Cloud Function
     */
    runtime: pulumi.Input<string>;
    /**
     * Service account ID for Yandex Cloud Function
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * Tags for Yandex Cloud Function. Tag "$latest" isn't returned.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User-defined string for current function version. User must change this string any times when function changed. Function will be updated when hash is changed.
     */
    userHash: pulumi.Input<string>;
}
