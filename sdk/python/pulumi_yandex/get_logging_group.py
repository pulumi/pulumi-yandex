# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetLoggingGroupResult',
    'AwaitableGetLoggingGroupResult',
    'get_logging_group',
    'get_logging_group_output',
]

@pulumi.output_type
class GetLoggingGroupResult:
    """
    A collection of values returned by getLoggingGroup.
    """
    def __init__(__self__, cloud_id=None, created_at=None, description=None, folder_id=None, group_id=None, id=None, labels=None, name=None, retention_period=None, status=None):
        if cloud_id and not isinstance(cloud_id, str):
            raise TypeError("Expected argument 'cloud_id' to be a str")
        pulumi.set(__self__, "cloud_id", cloud_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if retention_period and not isinstance(retention_period, str):
            raise TypeError("Expected argument 'retention_period' to be a str")
        pulumi.set(__self__, "retention_period", retention_period)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="cloudId")
    def cloud_id(self) -> str:
        return pulumi.get(self, "cloud_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> str:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> str:
        return pulumi.get(self, "retention_period")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


class AwaitableGetLoggingGroupResult(GetLoggingGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoggingGroupResult(
            cloud_id=self.cloud_id,
            created_at=self.created_at,
            description=self.description,
            folder_id=self.folder_id,
            group_id=self.group_id,
            id=self.id,
            labels=self.labels,
            name=self.name,
            retention_period=self.retention_period,
            status=self.status)


def get_logging_group(folder_id: Optional[str] = None,
                      group_id: Optional[str] = None,
                      name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoggingGroupResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['groupId'] = group_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('yandex:index/getLoggingGroup:getLoggingGroup', __args__, opts=opts, typ=GetLoggingGroupResult).value

    return AwaitableGetLoggingGroupResult(
        cloud_id=__ret__.cloud_id,
        created_at=__ret__.created_at,
        description=__ret__.description,
        folder_id=__ret__.folder_id,
        group_id=__ret__.group_id,
        id=__ret__.id,
        labels=__ret__.labels,
        name=__ret__.name,
        retention_period=__ret__.retention_period,
        status=__ret__.status)


@_utilities.lift_output_func(get_logging_group)
def get_logging_group_output(folder_id: Optional[pulumi.Input[Optional[str]]] = None,
                             group_id: Optional[pulumi.Input[Optional[str]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoggingGroupResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
