# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetIamServiceAccountResult',
    'AwaitableGetIamServiceAccountResult',
    'get_iam_service_account',
]

@pulumi.output_type
class GetIamServiceAccountResult:
    """
    A collection of values returned by getIamServiceAccount.
    """
    def __init__(__self__, created_at=None, description=None, folder_id=None, id=None, name=None, service_account_id=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_account_id and not isinstance(service_account_id, str):
            raise TypeError("Expected argument 'service_account_id' to be a str")
        pulumi.set(__self__, "service_account_id", service_account_id)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the service account.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> str:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> str:
        return pulumi.get(self, "service_account_id")


class AwaitableGetIamServiceAccountResult(GetIamServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIamServiceAccountResult(
            created_at=self.created_at,
            description=self.description,
            folder_id=self.folder_id,
            id=self.id,
            name=self.name,
            service_account_id=self.service_account_id)


def get_iam_service_account(folder_id: Optional[str] = None,
                            name: Optional[str] = None,
                            service_account_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIamServiceAccountResult:
    """
    Get information about a Yandex IAM service account. For more information about accounts, see
    [Yandex.Cloud IAM accounts](https://cloud.yandex.com/docs/iam/concepts/#accounts).

    ```python
    import pulumi
    import pulumi_yandex as yandex

    builder = yandex.get_iam_service_account(service_account_id="sa_id")
    deployer = yandex.get_iam_service_account(name="sa_name")
    ```

    ## Argument reference

    * `service_account_id` - (Optional) ID of a specific service account.

    * `name` - (Optional) Name of a specific service account.

    > **NOTE:** One of `service_account_id` or `name` should be specified.

    * `folder_id` - (Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['serviceAccountId'] = service_account_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('yandex:index/getIamServiceAccount:getIamServiceAccount', __args__, opts=opts, typ=GetIamServiceAccountResult).value

    return AwaitableGetIamServiceAccountResult(
        created_at=__ret__.created_at,
        description=__ret__.description,
        folder_id=__ret__.folder_id,
        id=__ret__.id,
        name=__ret__.name,
        service_account_id=__ret__.service_account_id)
