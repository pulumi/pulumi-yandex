// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("yandex");

/**
 * ID of Yandex.Cloud tenant.
 */
export declare const cloudId: string | undefined;
Object.defineProperty(exports, "cloudId", {
    get() {
        return __config.get("cloudId");
    },
    enumerable: true,
});

/**
 * The API endpoint for Yandex.Cloud SDK client.
 */
export declare const endpoint: string | undefined;
Object.defineProperty(exports, "endpoint", {
    get() {
        return __config.get("endpoint");
    },
    enumerable: true,
});

/**
 * The default folder ID where resources will be placed.
 */
export declare const folderId: string | undefined;
Object.defineProperty(exports, "folderId", {
    get() {
        return __config.get("folderId");
    },
    enumerable: true,
});

/**
 * Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
 */
export declare const insecure: boolean | undefined;
Object.defineProperty(exports, "insecure", {
    get() {
        return __config.getObject<boolean>("insecure");
    },
    enumerable: true,
});

/**
 * The maximum number of times an API request is being executed. If the API request still fails, an error is thrown.
 */
export declare const maxRetries: number | undefined;
Object.defineProperty(exports, "maxRetries", {
    get() {
        return __config.getObject<number>("maxRetries");
    },
    enumerable: true,
});

export declare const organizationId: string | undefined;
Object.defineProperty(exports, "organizationId", {
    get() {
        return __config.get("organizationId");
    },
    enumerable: true,
});

/**
 * Disable use of TLS. Default value is `false`.
 */
export declare const plaintext: boolean | undefined;
Object.defineProperty(exports, "plaintext", {
    get() {
        return __config.getObject<boolean>("plaintext");
    },
    enumerable: true,
});

/**
 * Either the path to or the contents of a Service Account key file in JSON format.
 */
export declare const serviceAccountKeyFile: string | undefined;
Object.defineProperty(exports, "serviceAccountKeyFile", {
    get() {
        return __config.get("serviceAccountKeyFile");
    },
    enumerable: true,
});

/**
 * Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly
 * specified.
 */
export declare const storageAccessKey: string | undefined;
Object.defineProperty(exports, "storageAccessKey", {
    get() {
        return __config.get("storageAccessKey");
    },
    enumerable: true,
});

/**
 * Yandex.Cloud storage service endpoint. Default is storage.yandexcloud.net
 */
export declare const storageEndpoint: string | undefined;
Object.defineProperty(exports, "storageEndpoint", {
    get() {
        return __config.get("storageEndpoint");
    },
    enumerable: true,
});

/**
 * Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly
 * specified.
 */
export declare const storageSecretKey: string | undefined;
Object.defineProperty(exports, "storageSecretKey", {
    get() {
        return __config.get("storageSecretKey");
    },
    enumerable: true,
});

/**
 * The access token for API operations.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

/**
 * Yandex.Cloud Message Queue service access key. Used when a message queue resource doesn't have an access key explicitly
 * specified.
 */
export declare const ymqAccessKey: string | undefined;
Object.defineProperty(exports, "ymqAccessKey", {
    get() {
        return __config.get("ymqAccessKey");
    },
    enumerable: true,
});

/**
 * Yandex.Cloud Message Queue service endpoint. Default is message-queue.api.cloud.yandex.net
 */
export declare const ymqEndpoint: string | undefined;
Object.defineProperty(exports, "ymqEndpoint", {
    get() {
        return __config.get("ymqEndpoint");
    },
    enumerable: true,
});

/**
 * Yandex.Cloud Message Queue service secret key. Used when a message queue resource doesn't have a secret key explicitly
 * specified.
 */
export declare const ymqSecretKey: string | undefined;
Object.defineProperty(exports, "ymqSecretKey", {
    get() {
        return __config.get("ymqSecretKey");
    },
    enumerable: true,
});

/**
 * The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
 */
export declare const zone: string | undefined;
Object.defineProperty(exports, "zone", {
    get() {
        return __config.get("zone");
    },
    enumerable: true,
});

