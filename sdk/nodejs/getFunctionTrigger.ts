// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Cloud Function Trigger. For more information about Yandex Cloud Functions, see
 * [Yandex Cloud Functions](https://cloud.yandex.com/docs/functions/).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const myTrigger = pulumi.output(yandex.getFunctionTrigger({
 *     triggerId: "are1sampletrigger11",
 * }));
 * ```
 *
 * This data source is used to define [Yandex Cloud Functions Trigger](https://cloud.yandex.com/docs/functions/concepts/trigger) that can be used by other resources.
 */
export function getFunctionTrigger(args?: GetFunctionTriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionTriggerResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("yandex:index/getFunctionTrigger:getFunctionTrigger", {
        "folderId": args.folderId,
        "name": args.name,
        "triggerId": args.triggerId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctionTrigger.
 */
export interface GetFunctionTriggerArgs {
    /**
     * Folder ID for the Yandex Cloud Functions Trigger
     */
    folderId?: string;
    /**
     * Yandex Cloud Functions Trigger name used to define trigger
     */
    name?: string;
    /**
     * Yandex Cloud Functions Trigger id used to define trigger
     */
    triggerId?: string;
}

/**
 * A collection of values returned by getFunctionTrigger.
 */
export interface GetFunctionTriggerResult {
    /**
     * Creation timestamp of the Yandex Cloud Functions Trigger
     */
    readonly createdAt: string;
    /**
     * Description of the Yandex Cloud Functions Trigger
     */
    readonly description: string;
    /**
     * Dead Letter Queue settings definition for Yandex Cloud Functions Trigger
     * * `dlq.0.queue_id` - ID of Dead Letter Queue for Trigger (Queue ARN)
     * * `dlq.0.service_account_id` - Service Account ID for Dead Letter Queue for Yandex Cloud Functions Trigger
     */
    readonly dlqs: outputs.GetFunctionTriggerDlq[];
    /**
     * Folder ID for the Yandex Cloud Functions Trigger
     */
    readonly folderId?: string;
    /**
     * [Yandex.Cloud Function](https://cloud.yandex.com/docs/functions/concepts/function) settings definition for Yandex Cloud Functions Trigger
     * * `function.0.id` - Yandex.Cloud Function ID for Yandex Cloud Functions Trigger
     * * `function.0.service_account_id` - Service account ID for Yandex.Cloud Function for Yandex Cloud Functions Trigger
     * * `function.0.tag` - Tag for Yandex.Cloud Function for Yandex Cloud Functions Trigger
     * * `function.0.retry_attempts` - Retry attempts for Yandex.Cloud Function for Yandex Cloud Functions Trigger
     * * `function.0.retry_interval` - Retry interval in seconds for Yandex.Cloud Function for Yandex Cloud Functions Trigger
     */
    readonly functions: outputs.GetFunctionTriggerFunction[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * [IoT](https://cloud.yandex.com/docs/functions/concepts/trigger/iot-core-trigger) settings definition for Yandex Cloud Functions Trigger, if present
     * * `iot.0.registry_id` - IoT Registry ID for Yandex Cloud Functions Trigger
     * * `iot.0.device_id` - IoT Device ID for Yandex Cloud Functions Trigger
     * * `iot.0.topic` - IoT Topic for Yandex Cloud Functions Trigger
     */
    readonly iots: outputs.GetFunctionTriggerIot[];
    /**
     * A set of key/value label pairs to assign to the Yandex Cloud Functions Trigger
     */
    readonly labels: {[key: string]: string};
    readonly logGroups: outputs.GetFunctionTriggerLogGroup[];
    readonly loggings: outputs.GetFunctionTriggerLogging[];
    /**
     * [Message Queue](https://cloud.yandex.com/docs/functions/concepts/trigger/ymq-trigger) settings definition for Yandex Cloud Functions Trigger, if present
     * * `message_queue.0.queue_id` - Message Queue ID for Yandex Cloud Functions Trigger
     * * `message_queue.0.service_account_id` - Message Queue Service Account ID for Yandex Cloud Functions Trigger
     * * `message_queue.0.batch_cutoff` - Batch Duration in seconds for Yandex Cloud Functions Trigger
     * * `message_queue.0.batch_size` - Batch Size for Yandex Cloud Functions Trigger
     * * `message_queue.0.visibility_timeout` - Visibility timeout for Yandex Cloud Functions Trigger
     */
    readonly messageQueues: outputs.GetFunctionTriggerMessageQueue[];
    readonly name?: string;
    /**
     * [Object Storage](https://cloud.yandex.com/docs/functions/concepts/trigger/os-trigger) settings definition for Yandex Cloud Functions Trigger, if present
     * * `object_storage.0.bucket_id` - Object Storage Bucket ID for Yandex Cloud Functions Trigger
     * * `object_storage.0.prefix` - Prefix for Object Storage for Yandex Cloud Functions Trigger
     * * `object_storage.0.suffix` - Suffix for Object Storage for Yandex Cloud Functions Trigger
     * * `object_storage.0.create` - Boolean flag for setting create event for Yandex Cloud Functions Trigger
     * * `object_storage.0.update` - Boolean flag for setting update event for Yandex Cloud Functions Trigger
     * * `object_storage.0.delete` - Boolean flag for setting delete event for Yandex Cloud Functions Trigger
     */
    readonly objectStorages: outputs.GetFunctionTriggerObjectStorage[];
    /**
     * [Timer](https://cloud.yandex.com/docs/functions/concepts/trigger/timer) settings definition for Yandex Cloud Functions Trigger, if present
     * * `timer.0.cron_expression` - Cron expression for timer for Yandex Cloud Functions Trigger
     */
    readonly timers: outputs.GetFunctionTriggerTimer[];
    readonly triggerId?: string;
}

export function getFunctionTriggerOutput(args?: GetFunctionTriggerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionTriggerResult> {
    return pulumi.output(args).apply(a => getFunctionTrigger(a, opts))
}

/**
 * A collection of arguments for invoking getFunctionTrigger.
 */
export interface GetFunctionTriggerOutputArgs {
    /**
     * Folder ID for the Yandex Cloud Functions Trigger
     */
    folderId?: pulumi.Input<string>;
    /**
     * Yandex Cloud Functions Trigger name used to define trigger
     */
    name?: pulumi.Input<string>;
    /**
     * Yandex Cloud Functions Trigger id used to define trigger
     */
    triggerId?: pulumi.Input<string>;
}
