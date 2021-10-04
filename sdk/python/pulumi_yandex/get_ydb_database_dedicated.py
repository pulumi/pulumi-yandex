# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetYdbDatabaseDedicatedResult',
    'AwaitableGetYdbDatabaseDedicatedResult',
    'get_ydb_database_dedicated',
]

@pulumi.output_type
class GetYdbDatabaseDedicatedResult:
    """
    A collection of values returned by getYdbDatabaseDedicated.
    """
    def __init__(__self__, assign_public_ips=None, created_at=None, database_id=None, database_path=None, description=None, folder_id=None, id=None, labels=None, location=None, location_id=None, name=None, network_id=None, resource_preset_id=None, scale_policy=None, status=None, storage_configs=None, subnet_ids=None, tls_enabled=None, ydb_api_endpoint=None, ydb_full_endpoint=None):
        if assign_public_ips and not isinstance(assign_public_ips, bool):
            raise TypeError("Expected argument 'assign_public_ips' to be a bool")
        pulumi.set(__self__, "assign_public_ips", assign_public_ips)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if database_path and not isinstance(database_path, str):
            raise TypeError("Expected argument 'database_path' to be a str")
        pulumi.set(__self__, "database_path", database_path)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if location and not isinstance(location, dict):
            raise TypeError("Expected argument 'location' to be a dict")
        pulumi.set(__self__, "location", location)
        if location_id and not isinstance(location_id, str):
            raise TypeError("Expected argument 'location_id' to be a str")
        pulumi.set(__self__, "location_id", location_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if resource_preset_id and not isinstance(resource_preset_id, str):
            raise TypeError("Expected argument 'resource_preset_id' to be a str")
        pulumi.set(__self__, "resource_preset_id", resource_preset_id)
        if scale_policy and not isinstance(scale_policy, dict):
            raise TypeError("Expected argument 'scale_policy' to be a dict")
        pulumi.set(__self__, "scale_policy", scale_policy)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if storage_configs and not isinstance(storage_configs, list):
            raise TypeError("Expected argument 'storage_configs' to be a list")
        pulumi.set(__self__, "storage_configs", storage_configs)
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tls_enabled and not isinstance(tls_enabled, bool):
            raise TypeError("Expected argument 'tls_enabled' to be a bool")
        pulumi.set(__self__, "tls_enabled", tls_enabled)
        if ydb_api_endpoint and not isinstance(ydb_api_endpoint, str):
            raise TypeError("Expected argument 'ydb_api_endpoint' to be a str")
        pulumi.set(__self__, "ydb_api_endpoint", ydb_api_endpoint)
        if ydb_full_endpoint and not isinstance(ydb_full_endpoint, str):
            raise TypeError("Expected argument 'ydb_full_endpoint' to be a str")
        pulumi.set(__self__, "ydb_full_endpoint", ydb_full_endpoint)

    @property
    @pulumi.getter(name="assignPublicIps")
    def assign_public_ips(self) -> bool:
        """
        Whether public IP addresses are assigned to the Yandex Database cluster.
        """
        return pulumi.get(self, "assign_public_ips")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The Yandex Database cluster creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> Optional[str]:
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter(name="databasePath")
    def database_path(self) -> str:
        """
        Full database path of the Yandex Database cluster.
        Useful for SDK configuration.
        """
        return pulumi.get(self, "database_path")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the Yandex Database cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[str]:
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
    def labels(self) -> Mapping[str, str]:
        """
        A set of key/value label pairs assigned to the Yandex Database cluster.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> 'outputs.GetYdbDatabaseDedicatedLocationResult':
        """
        Location of the Yandex Database cluster.
        The structure is documented below.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> str:
        """
        Location ID of the Yandex Database cluster.
        """
        return pulumi.get(self, "location_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        """
        ID of the network the Yandex Database cluster is attached to.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="resourcePresetId")
    def resource_preset_id(self) -> str:
        """
        The Yandex Database cluster preset.
        """
        return pulumi.get(self, "resource_preset_id")

    @property
    @pulumi.getter(name="scalePolicy")
    def scale_policy(self) -> 'outputs.GetYdbDatabaseDedicatedScalePolicyResult':
        """
        Scaling policy of the Yandex Database cluster.
        The structure is documented below.
        """
        return pulumi.get(self, "scale_policy")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the Yandex Database cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageConfigs")
    def storage_configs(self) -> Sequence['outputs.GetYdbDatabaseDedicatedStorageConfigResult']:
        """
        A list of storage configuration options of the Yandex Database cluster.
        The structure is documented below.
        """
        return pulumi.get(self, "storage_configs")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Sequence[str]:
        """
        List of subnet IDs the Yandex Database cluster is attached to.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> bool:
        """
        Whether TLS is enabled for the Yandex Database cluster.
        Useful for SDK configuration.
        """
        return pulumi.get(self, "tls_enabled")

    @property
    @pulumi.getter(name="ydbApiEndpoint")
    def ydb_api_endpoint(self) -> str:
        """
        API endpoint of the Yandex Database cluster.
        Useful for SDK configuration.
        """
        return pulumi.get(self, "ydb_api_endpoint")

    @property
    @pulumi.getter(name="ydbFullEndpoint")
    def ydb_full_endpoint(self) -> str:
        """
        Full endpoint of the Yandex Database cluster.
        """
        return pulumi.get(self, "ydb_full_endpoint")


class AwaitableGetYdbDatabaseDedicatedResult(GetYdbDatabaseDedicatedResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetYdbDatabaseDedicatedResult(
            assign_public_ips=self.assign_public_ips,
            created_at=self.created_at,
            database_id=self.database_id,
            database_path=self.database_path,
            description=self.description,
            folder_id=self.folder_id,
            id=self.id,
            labels=self.labels,
            location=self.location,
            location_id=self.location_id,
            name=self.name,
            network_id=self.network_id,
            resource_preset_id=self.resource_preset_id,
            scale_policy=self.scale_policy,
            status=self.status,
            storage_configs=self.storage_configs,
            subnet_ids=self.subnet_ids,
            tls_enabled=self.tls_enabled,
            ydb_api_endpoint=self.ydb_api_endpoint,
            ydb_full_endpoint=self.ydb_full_endpoint)


def get_ydb_database_dedicated(database_id: Optional[str] = None,
                               folder_id: Optional[str] = None,
                               name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetYdbDatabaseDedicatedResult:
    """
    Get information about a Yandex Database (dedicated) cluster.
    For more information, see [the official documentation](https://cloud.yandex.com/en/docs/ydb/concepts/serverless_and_dedicated).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_yandex as yandex

    my_database = yandex.get_ydb_database_dedicated(database_id="some_ydb_dedicated_database_id")
    pulumi.export("ydbApiEndpoint", my_database.ydb_api_endpoint)
    ```


    :param str database_id: ID of the Yandex Database cluster.
    :param str folder_id: ID of the folder that the Yandex Database cluster belongs to.
           It will be deduced from provider configuration if not set explicitly.
    :param str name: Name of the Yandex Database cluster.
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    __args__['folderId'] = folder_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('yandex:index/getYdbDatabaseDedicated:getYdbDatabaseDedicated', __args__, opts=opts, typ=GetYdbDatabaseDedicatedResult).value

    return AwaitableGetYdbDatabaseDedicatedResult(
        assign_public_ips=__ret__.assign_public_ips,
        created_at=__ret__.created_at,
        database_id=__ret__.database_id,
        database_path=__ret__.database_path,
        description=__ret__.description,
        folder_id=__ret__.folder_id,
        id=__ret__.id,
        labels=__ret__.labels,
        location=__ret__.location,
        location_id=__ret__.location_id,
        name=__ret__.name,
        network_id=__ret__.network_id,
        resource_preset_id=__ret__.resource_preset_id,
        scale_policy=__ret__.scale_policy,
        status=__ret__.status,
        storage_configs=__ret__.storage_configs,
        subnet_ids=__ret__.subnet_ids,
        tls_enabled=__ret__.tls_enabled,
        ydb_api_endpoint=__ret__.ydb_api_endpoint,
        ydb_full_endpoint=__ret__.ydb_full_endpoint)
