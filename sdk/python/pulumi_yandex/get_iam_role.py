# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetIamRoleResult',
    'AwaitableGetIamRoleResult',
    'get_iam_role',
]

@pulumi.output_type
class GetIamRoleResult:
    """
    A collection of values returned by getIamRole.
    """
    def __init__(__self__, description=None, id=None, role_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if role_id and not isinstance(role_id, str):
            raise TypeError("Expected argument 'role_id' to be a str")
        pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> str:
        return pulumi.get(self, "role_id")


class AwaitableGetIamRoleResult(GetIamRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIamRoleResult(
            description=self.description,
            id=self.id,
            role_id=self.role_id)


def get_iam_role(description: Optional[str] = None,
                 role_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIamRoleResult:
    """
    Generates an [IAM] role document that may be referenced by and applied to
    other Yandex.Cloud Platform resources, such as the `ResourcemanagerFolder` resource. For more information, see
    [the official documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/roles).

    ```python
    import pulumi
    import pulumi_yandex as yandex

    admin = yandex.get_iam_role(binding=[{
        "members": ["userAccount:user_id_1"],
        "role": "admin",
    }])
    ```

    This data source is used to define [IAM] roles in order to apply them to other resources.
    Currently, defining a role through a data source and referencing that role
    from another resource is the only way to apply an IAM role to a resource.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['roleId'] = role_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('yandex:index/getIamRole:getIamRole', __args__, opts=opts, typ=GetIamRoleResult).value

    return AwaitableGetIamRoleResult(
        description=__ret__.description,
        id=__ret__.id,
        role_id=__ret__.role_id)
