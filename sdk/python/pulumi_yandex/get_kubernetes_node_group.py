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
    'GetKubernetesNodeGroupResult',
    'AwaitableGetKubernetesNodeGroupResult',
    'get_kubernetes_node_group',
    'get_kubernetes_node_group_output',
]

@pulumi.output_type
class GetKubernetesNodeGroupResult:
    """
    A collection of values returned by getKubernetesNodeGroup.
    """
    def __init__(__self__, allocation_policy=None, allowed_unsafe_sysctls=None, cluster_id=None, created_at=None, deploy_policy=None, description=None, folder_id=None, id=None, instance_group_id=None, instance_template=None, labels=None, maintenance_policy=None, name=None, node_group_id=None, node_labels=None, node_taints=None, scale_policy=None, status=None, version_info=None):
        if allocation_policy and not isinstance(allocation_policy, dict):
            raise TypeError("Expected argument 'allocation_policy' to be a dict")
        pulumi.set(__self__, "allocation_policy", allocation_policy)
        if allowed_unsafe_sysctls and not isinstance(allowed_unsafe_sysctls, list):
            raise TypeError("Expected argument 'allowed_unsafe_sysctls' to be a list")
        pulumi.set(__self__, "allowed_unsafe_sysctls", allowed_unsafe_sysctls)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if deploy_policy and not isinstance(deploy_policy, dict):
            raise TypeError("Expected argument 'deploy_policy' to be a dict")
        pulumi.set(__self__, "deploy_policy", deploy_policy)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_group_id and not isinstance(instance_group_id, str):
            raise TypeError("Expected argument 'instance_group_id' to be a str")
        pulumi.set(__self__, "instance_group_id", instance_group_id)
        if instance_template and not isinstance(instance_template, dict):
            raise TypeError("Expected argument 'instance_template' to be a dict")
        pulumi.set(__self__, "instance_template", instance_template)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if maintenance_policy and not isinstance(maintenance_policy, dict):
            raise TypeError("Expected argument 'maintenance_policy' to be a dict")
        pulumi.set(__self__, "maintenance_policy", maintenance_policy)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_group_id and not isinstance(node_group_id, str):
            raise TypeError("Expected argument 'node_group_id' to be a str")
        pulumi.set(__self__, "node_group_id", node_group_id)
        if node_labels and not isinstance(node_labels, dict):
            raise TypeError("Expected argument 'node_labels' to be a dict")
        pulumi.set(__self__, "node_labels", node_labels)
        if node_taints and not isinstance(node_taints, list):
            raise TypeError("Expected argument 'node_taints' to be a list")
        pulumi.set(__self__, "node_taints", node_taints)
        if scale_policy and not isinstance(scale_policy, dict):
            raise TypeError("Expected argument 'scale_policy' to be a dict")
        pulumi.set(__self__, "scale_policy", scale_policy)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if version_info and not isinstance(version_info, dict):
            raise TypeError("Expected argument 'version_info' to be a dict")
        pulumi.set(__self__, "version_info", version_info)

    @property
    @pulumi.getter(name="allocationPolicy")
    def allocation_policy(self) -> 'outputs.GetKubernetesNodeGroupAllocationPolicyResult':
        """
        This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.
        """
        return pulumi.get(self, "allocation_policy")

    @property
    @pulumi.getter(name="allowedUnsafeSysctls")
    def allowed_unsafe_sysctls(self) -> Sequence[str]:
        """
        A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).
        """
        return pulumi.get(self, "allowed_unsafe_sysctls")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The ID of the Kubernetes cluster that this node group belongs to.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The Kubernetes node group creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deployPolicy")
    def deploy_policy(self) -> 'outputs.GetKubernetesNodeGroupDeployPolicyResult':
        """
        Deploy policy of the node group. The structure is documented below.
        """
        return pulumi.get(self, "deploy_policy")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the Kubernetes node group.
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
    @pulumi.getter(name="instanceGroupId")
    def instance_group_id(self) -> str:
        """
        ID of instance group that is used to manage this Kubernetes node group.
        """
        return pulumi.get(self, "instance_group_id")

    @property
    @pulumi.getter(name="instanceTemplate")
    def instance_template(self) -> 'outputs.GetKubernetesNodeGroupInstanceTemplateResult':
        """
        Template used to create compute instances in this Kubernetes node group. The structure is documented below.
        """
        return pulumi.get(self, "instance_template")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        A map of labels applied to this instance.
        * `resources.0.memory` - The memory size allocated to the instance.
        * `resources.0.cores` - Number of CPU cores allocated to the instance.
        * `resources.0.core_fraction` - Baseline core performance as a percent.
        * `resources.0.gpus` - Number of GPU cores allocated to the instance.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> 'outputs.GetKubernetesNodeGroupMaintenancePolicyResult':
        """
        Information about maintenance policy for this Kubernetes node group. The structure is documented below.
        """
        return pulumi.get(self, "maintenance_policy")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeGroupId")
    def node_group_id(self) -> str:
        return pulumi.get(self, "node_group_id")

    @property
    @pulumi.getter(name="nodeLabels")
    def node_labels(self) -> Mapping[str, str]:
        """
        A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.
        """
        return pulumi.get(self, "node_labels")

    @property
    @pulumi.getter(name="nodeTaints")
    def node_taints(self) -> Sequence[str]:
        """
        A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.
        """
        return pulumi.get(self, "node_taints")

    @property
    @pulumi.getter(name="scalePolicy")
    def scale_policy(self) -> 'outputs.GetKubernetesNodeGroupScalePolicyResult':
        """
        Scale policy of the node group. The structure is documented below.
        """
        return pulumi.get(self, "scale_policy")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the Kubernetes node group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="versionInfo")
    def version_info(self) -> 'outputs.GetKubernetesNodeGroupVersionInfoResult':
        """
        Information about Kubernetes node group version. The structure is documented below.
        """
        return pulumi.get(self, "version_info")


class AwaitableGetKubernetesNodeGroupResult(GetKubernetesNodeGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesNodeGroupResult(
            allocation_policy=self.allocation_policy,
            allowed_unsafe_sysctls=self.allowed_unsafe_sysctls,
            cluster_id=self.cluster_id,
            created_at=self.created_at,
            deploy_policy=self.deploy_policy,
            description=self.description,
            folder_id=self.folder_id,
            id=self.id,
            instance_group_id=self.instance_group_id,
            instance_template=self.instance_template,
            labels=self.labels,
            maintenance_policy=self.maintenance_policy,
            name=self.name,
            node_group_id=self.node_group_id,
            node_labels=self.node_labels,
            node_taints=self.node_taints,
            scale_policy=self.scale_policy,
            status=self.status,
            version_info=self.version_info)


def get_kubernetes_node_group(folder_id: Optional[str] = None,
                              name: Optional[str] = None,
                              node_group_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubernetesNodeGroupResult:
    """
    Get information about a Yandex Kubernetes Node Group. For more information, see
    [the official documentation](https://cloud.yandex.com/docs/managed-kubernetes/concepts/#node-group).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_yandex as yandex

    my_node_group = yandex.get_kubernetes_node_group(node_group_id="some_k8s_node_group_id")
    pulumi.export("myNodeGroup.status", my_node_group.status)
    ```


    :param str folder_id: Folder that the resource belongs to. If value is omitted, the default provider folder is used.
    :param str name: Name of a specific Kubernetes node group.
    :param str node_group_id: ID of a specific Kubernetes node group.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['nodeGroupId'] = node_group_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('yandex:index/getKubernetesNodeGroup:getKubernetesNodeGroup', __args__, opts=opts, typ=GetKubernetesNodeGroupResult).value

    return AwaitableGetKubernetesNodeGroupResult(
        allocation_policy=__ret__.allocation_policy,
        allowed_unsafe_sysctls=__ret__.allowed_unsafe_sysctls,
        cluster_id=__ret__.cluster_id,
        created_at=__ret__.created_at,
        deploy_policy=__ret__.deploy_policy,
        description=__ret__.description,
        folder_id=__ret__.folder_id,
        id=__ret__.id,
        instance_group_id=__ret__.instance_group_id,
        instance_template=__ret__.instance_template,
        labels=__ret__.labels,
        maintenance_policy=__ret__.maintenance_policy,
        name=__ret__.name,
        node_group_id=__ret__.node_group_id,
        node_labels=__ret__.node_labels,
        node_taints=__ret__.node_taints,
        scale_policy=__ret__.scale_policy,
        status=__ret__.status,
        version_info=__ret__.version_info)


@_utilities.lift_output_func(get_kubernetes_node_group)
def get_kubernetes_node_group_output(folder_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     name: Optional[pulumi.Input[Optional[str]]] = None,
                                     node_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKubernetesNodeGroupResult]:
    """
    Get information about a Yandex Kubernetes Node Group. For more information, see
    [the official documentation](https://cloud.yandex.com/docs/managed-kubernetes/concepts/#node-group).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_yandex as yandex

    my_node_group = yandex.get_kubernetes_node_group(node_group_id="some_k8s_node_group_id")
    pulumi.export("myNodeGroup.status", my_node_group.status)
    ```


    :param str folder_id: Folder that the resource belongs to. If value is omitted, the default provider folder is used.
    :param str name: Name of a specific Kubernetes node group.
    :param str node_group_id: ID of a specific Kubernetes node group.
    """
    ...
