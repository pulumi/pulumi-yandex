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
    'GetMdbKafkaTopicResult',
    'AwaitableGetMdbKafkaTopicResult',
    'get_mdb_kafka_topic',
    'get_mdb_kafka_topic_output',
]

@pulumi.output_type
class GetMdbKafkaTopicResult:
    """
    A collection of values returned by getMdbKafkaTopic.
    """
    def __init__(__self__, cluster_id=None, id=None, name=None, partitions=None, replication_factor=None, topic_config=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partitions and not isinstance(partitions, int):
            raise TypeError("Expected argument 'partitions' to be a int")
        pulumi.set(__self__, "partitions", partitions)
        if replication_factor and not isinstance(replication_factor, int):
            raise TypeError("Expected argument 'replication_factor' to be a int")
        pulumi.set(__self__, "replication_factor", replication_factor)
        if topic_config and not isinstance(topic_config, dict):
            raise TypeError("Expected argument 'topic_config' to be a dict")
        pulumi.set(__self__, "topic_config", topic_config)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

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
    @pulumi.getter
    def partitions(self) -> int:
        return pulumi.get(self, "partitions")

    @property
    @pulumi.getter(name="replicationFactor")
    def replication_factor(self) -> int:
        return pulumi.get(self, "replication_factor")

    @property
    @pulumi.getter(name="topicConfig")
    def topic_config(self) -> 'outputs.GetMdbKafkaTopicTopicConfigResult':
        return pulumi.get(self, "topic_config")


class AwaitableGetMdbKafkaTopicResult(GetMdbKafkaTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMdbKafkaTopicResult(
            cluster_id=self.cluster_id,
            id=self.id,
            name=self.name,
            partitions=self.partitions,
            replication_factor=self.replication_factor,
            topic_config=self.topic_config)


def get_mdb_kafka_topic(cluster_id: Optional[str] = None,
                        name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMdbKafkaTopicResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('yandex:index/getMdbKafkaTopic:getMdbKafkaTopic', __args__, opts=opts, typ=GetMdbKafkaTopicResult).value

    return AwaitableGetMdbKafkaTopicResult(
        cluster_id=__ret__.cluster_id,
        id=__ret__.id,
        name=__ret__.name,
        partitions=__ret__.partitions,
        replication_factor=__ret__.replication_factor,
        topic_config=__ret__.topic_config)


@_utilities.lift_output_func(get_mdb_kafka_topic)
def get_mdb_kafka_topic_output(cluster_id: Optional[pulumi.Input[str]] = None,
                               name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMdbKafkaTopicResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
