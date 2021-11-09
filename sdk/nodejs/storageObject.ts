// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows management of [Yandex.Cloud Storage Object](https://cloud.yandex.com/docs/storage/concepts/object).
 *
 * ## Example Usage
 *
 * Example creating an object in an existing `cat-pictures` bucket.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const cute_cat_picture = new yandex.StorageObject("cute-cat-picture", {
 *     bucket: "cat-pictures",
 *     key: "cute-cat",
 *     source: "/images/cats/cute-cat.jpg",
 * });
 * ```
 */
export class StorageObject extends pulumi.CustomResource {
    /**
     * Get an existing StorageObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StorageObjectState, opts?: pulumi.CustomResourceOptions): StorageObject {
        return new StorageObject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/storageObject:StorageObject';

    /**
     * Returns true if the given object is an instance of StorageObject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageObject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageObject.__pulumiType;
    }

    /**
     * The access key to use when applying changes. If omitted, `storageAccessKey` specified in config is used.
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`.
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * The name of the containing bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
     */
    public readonly contentBase64!: pulumi.Output<string | undefined>;
    /**
     * A standard MIME type describing the format of the object data, e.g. `application/octet-stream`. All Valid MIME Types are valid for this input.
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * The name of the object once it is in the bucket.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The secret key to use when applying changes. If omitted, `storageSecretKey` specified in config is used.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * The path to a file that will be read and uploaded as raw bytes for the object content.
     */
    public readonly source!: pulumi.Output<string | undefined>;

    /**
     * Create a StorageObject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageObjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageObjectArgs | StorageObjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StorageObjectState | undefined;
            inputs["accessKey"] = state ? state.accessKey : undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["contentBase64"] = state ? state.contentBase64 : undefined;
            inputs["contentType"] = state ? state.contentType : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["secretKey"] = state ? state.secretKey : undefined;
            inputs["source"] = state ? state.source : undefined;
        } else {
            const args = argsOrState as StorageObjectArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            inputs["accessKey"] = args ? args.accessKey : undefined;
            inputs["acl"] = args ? args.acl : undefined;
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["content"] = args ? args.content : undefined;
            inputs["contentBase64"] = args ? args.contentBase64 : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["secretKey"] = args ? args.secretKey : undefined;
            inputs["source"] = args ? args.source : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StorageObject.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StorageObject resources.
 */
export interface StorageObjectState {
    /**
     * The access key to use when applying changes. If omitted, `storageAccessKey` specified in config is used.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`.
     */
    acl?: pulumi.Input<string>;
    /**
     * The name of the containing bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
     */
    content?: pulumi.Input<string>;
    /**
     * Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
     */
    contentBase64?: pulumi.Input<string>;
    /**
     * A standard MIME type describing the format of the object data, e.g. `application/octet-stream`. All Valid MIME Types are valid for this input.
     */
    contentType?: pulumi.Input<string>;
    /**
     * The name of the object once it is in the bucket.
     */
    key?: pulumi.Input<string>;
    /**
     * The secret key to use when applying changes. If omitted, `storageSecretKey` specified in config is used.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The path to a file that will be read and uploaded as raw bytes for the object content.
     */
    source?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StorageObject resource.
 */
export interface StorageObjectArgs {
    /**
     * The access key to use when applying changes. If omitted, `storageAccessKey` specified in config is used.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`.
     */
    acl?: pulumi.Input<string>;
    /**
     * The name of the containing bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
     */
    content?: pulumi.Input<string>;
    /**
     * Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
     */
    contentBase64?: pulumi.Input<string>;
    /**
     * A standard MIME type describing the format of the object data, e.g. `application/octet-stream`. All Valid MIME Types are valid for this input.
     */
    contentType?: pulumi.Input<string>;
    /**
     * The name of the object once it is in the bucket.
     */
    key: pulumi.Input<string>;
    /**
     * The secret key to use when applying changes. If omitted, `storageSecretKey` specified in config is used.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The path to a file that will be read and uploaded as raw bytes for the object content.
     */
    source?: pulumi.Input<string>;
}
