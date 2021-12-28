# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

cloudId: Optional[str]
"""
ID of Yandex.Cloud tenant.
"""

endpoint: Optional[str]
"""
The API endpoint for Yandex.Cloud SDK client.
"""

folderId: Optional[str]
"""
The default folder ID where resources will be placed.
"""

insecure: Optional[bool]
"""
Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
"""

maxRetries: Optional[int]
"""
The maximum number of times an API request is being executed. If the API request still fails, an error is thrown.
"""

organizationId: Optional[str]

plaintext: Optional[bool]
"""
Disable use of TLS. Default value is `false`.
"""

serviceAccountKeyFile: Optional[str]
"""
Either the path to or the contents of a Service Account key file in JSON format.
"""

storageAccessKey: Optional[str]
"""
Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly
specified.
"""

storageEndpoint: Optional[str]
"""
Yandex.Cloud storage service endpoint. Default is storage.yandexcloud.net
"""

storageSecretKey: Optional[str]
"""
Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly
specified.
"""

token: Optional[str]
"""
The access token for API operations.
"""

ymqAccessKey: Optional[str]
"""
Yandex.Cloud Message Queue service access key. Used when a message queue resource doesn't have an access key explicitly
specified.
"""

ymqEndpoint: Optional[str]
"""
Yandex.Cloud Message Queue service endpoint. Default is message-queue.api.cloud.yandex.net
"""

ymqSecretKey: Optional[str]
"""
Yandex.Cloud Message Queue service secret key. Used when a message queue resource doesn't have a secret key explicitly
specified.
"""

zone: Optional[str]
"""
The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
"""

