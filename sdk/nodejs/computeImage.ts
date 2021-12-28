// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a virtual machine image resource for the Yandex Compute Cloud service from an existing
 * tarball. For more information, see [the official documentation](https://cloud.yandex.com/docs/compute/concepts/image).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const foo_image = new yandex.ComputeImage("foo-image", {
 *     sourceUrl: "https://storage.yandexcloud.net/lucky-images/kube-it.img",
 * });
 * const vm = new yandex.ComputeInstance("vm", {
 *     bootDisk: {
 *         initializeParams: {
 *             imageId: foo_image.id,
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * A VM image can be imported using the `id` of the resource, e.g.
 *
 * ```sh
 *  $ pulumi import yandex:index/computeImage:ComputeImage web-image image_id
 * ```
 */
export class ComputeImage extends pulumi.CustomResource {
    /**
     * Get an existing ComputeImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeImageState, opts?: pulumi.CustomResourceOptions): ComputeImage {
        return new ComputeImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/computeImage:ComputeImage';

    /**
     * Returns true if the given object is an instance of ComputeImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeImage.__pulumiType;
    }

    /**
     * Creation timestamp of the image.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * An optional description of the image. Provide this property when
     * you create a resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the image family to which this image belongs.
     */
    public readonly family!: pulumi.Output<string | undefined>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the image.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Minimum size in GB of the disk that will be created from this image.
     */
    public readonly minDiskSize!: pulumi.Output<number>;
    /**
     * Name of the disk.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Operating system type that is contained in the image. Possible values: "LINUX", "WINDOWS".
     */
    public readonly osType!: pulumi.Output<string>;
    /**
     * Optimize the image to create a disk.
     */
    public readonly pooled!: pulumi.Output<boolean>;
    /**
     * License IDs that indicate which licenses are
     * attached to this image.
     */
    public readonly productIds!: pulumi.Output<string[]>;
    /**
     * The size of the image, specified in GB.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * The ID of a disk to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    public readonly sourceDisk!: pulumi.Output<string>;
    /**
     * The name of the family to use as the source of the new image.
     * The ID of the latest image is taken from the "standard-images" folder. Changing the family forces
     * a new resource to be created.
     */
    public readonly sourceFamily!: pulumi.Output<string>;
    /**
     * The ID of an existing image to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    public readonly sourceImage!: pulumi.Output<string>;
    /**
     * The ID of a snapshot to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    public readonly sourceSnapshot!: pulumi.Output<string>;
    /**
     * The URL to use as the source of the
     * image. Changing this URL forces a new resource to be created.
     */
    public readonly sourceUrl!: pulumi.Output<string>;
    /**
     * The status of the image.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ComputeImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ComputeImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeImageArgs | ComputeImageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeImageState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["family"] = state ? state.family : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["minDiskSize"] = state ? state.minDiskSize : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["pooled"] = state ? state.pooled : undefined;
            inputs["productIds"] = state ? state.productIds : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["sourceDisk"] = state ? state.sourceDisk : undefined;
            inputs["sourceFamily"] = state ? state.sourceFamily : undefined;
            inputs["sourceImage"] = state ? state.sourceImage : undefined;
            inputs["sourceSnapshot"] = state ? state.sourceSnapshot : undefined;
            inputs["sourceUrl"] = state ? state.sourceUrl : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ComputeImageArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["family"] = args ? args.family : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["minDiskSize"] = args ? args.minDiskSize : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["pooled"] = args ? args.pooled : undefined;
            inputs["productIds"] = args ? args.productIds : undefined;
            inputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            inputs["sourceFamily"] = args ? args.sourceFamily : undefined;
            inputs["sourceImage"] = args ? args.sourceImage : undefined;
            inputs["sourceSnapshot"] = args ? args.sourceSnapshot : undefined;
            inputs["sourceUrl"] = args ? args.sourceUrl : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ComputeImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeImage resources.
 */
export interface ComputeImageState {
    /**
     * Creation timestamp of the image.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * An optional description of the image. Provide this property when
     * you create a resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the image family to which this image belongs.
     */
    family?: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the image.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Minimum size in GB of the disk that will be created from this image.
     */
    minDiskSize?: pulumi.Input<number>;
    /**
     * Name of the disk.
     */
    name?: pulumi.Input<string>;
    /**
     * Operating system type that is contained in the image. Possible values: "LINUX", "WINDOWS".
     */
    osType?: pulumi.Input<string>;
    /**
     * Optimize the image to create a disk.
     */
    pooled?: pulumi.Input<boolean>;
    /**
     * License IDs that indicate which licenses are
     * attached to this image.
     */
    productIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The size of the image, specified in GB.
     */
    size?: pulumi.Input<number>;
    /**
     * The ID of a disk to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * The name of the family to use as the source of the new image.
     * The ID of the latest image is taken from the "standard-images" folder. Changing the family forces
     * a new resource to be created.
     */
    sourceFamily?: pulumi.Input<string>;
    /**
     * The ID of an existing image to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * The ID of a snapshot to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    sourceSnapshot?: pulumi.Input<string>;
    /**
     * The URL to use as the source of the
     * image. Changing this URL forces a new resource to be created.
     */
    sourceUrl?: pulumi.Input<string>;
    /**
     * The status of the image.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeImage resource.
 */
export interface ComputeImageArgs {
    /**
     * An optional description of the image. Provide this property when
     * you create a resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the image family to which this image belongs.
     */
    family?: pulumi.Input<string>;
    /**
     * The ID of the folder that the resource belongs to. If it
     * is not provided, the default provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the image.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Minimum size in GB of the disk that will be created from this image.
     */
    minDiskSize?: pulumi.Input<number>;
    /**
     * Name of the disk.
     */
    name?: pulumi.Input<string>;
    /**
     * Operating system type that is contained in the image. Possible values: "LINUX", "WINDOWS".
     */
    osType?: pulumi.Input<string>;
    /**
     * Optimize the image to create a disk.
     */
    pooled?: pulumi.Input<boolean>;
    /**
     * License IDs that indicate which licenses are
     * attached to this image.
     */
    productIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of a disk to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * The name of the family to use as the source of the new image.
     * The ID of the latest image is taken from the "standard-images" folder. Changing the family forces
     * a new resource to be created.
     */
    sourceFamily?: pulumi.Input<string>;
    /**
     * The ID of an existing image to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * The ID of a snapshot to use as the source of the
     * image. Changing this ID forces a new resource to be created.
     */
    sourceSnapshot?: pulumi.Input<string>;
    /**
     * The URL to use as the source of the
     * image. Changing this URL forces a new resource to be created.
     */
    sourceUrl?: pulumi.Input<string>;
}
