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
    def __init__(__self__, allocation_policies=None, allowed_unsafe_sysctls=None, cluster_id=None, created_at=None, deploy_policies=None, description=None, folder_id=None, id=None, instance_group_id=None, instance_templates=None, labels=None, maintenance_policies=None, name=None, node_group_id=None, node_labels=None, node_taints=None, scale_policies=None, status=None, version_infos=None):
        if allocation_policies and not isinstance(allocation_policies, list):
            raise TypeError("Expected argument 'allocation_policies' to be a list")
        pulumi.set(__self__, "allocation_policies", allocation_policies)
        if allowed_unsafe_sysctls and not isinstance(allowed_unsafe_sysctls, list):
            raise TypeError("Expected argument 'allowed_unsafe_sysctls' to be a list")
        pulumi.set(__self__, "allowed_unsafe_sysctls", allowed_unsafe_sysctls)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if deploy_policies and not isinstance(deploy_policies, list):
            raise TypeError("Expected argument 'deploy_policies' to be a list")
        pulumi.set(__self__, "deploy_policies", deploy_policies)
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
        if instance_templates and not isinstance(instance_templates, list):
            raise TypeError("Expected argument 'instance_templates' to be a list")
        pulumi.set(__self__, "instance_templates", instance_templates)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if maintenance_policies and not isinstance(maintenance_policies, list):
            raise TypeError("Expected argument 'maintenance_policies' to be a list")
        pulumi.set(__self__, "maintenance_policies", maintenance_policies)
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
        if scale_policies and not isinstance(scale_policies, list):
            raise TypeError("Expected argument 'scale_policies' to be a list")
        pulumi.set(__self__, "scale_policies", scale_policies)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if version_infos and not isinstance(version_infos, list):
            raise TypeError("Expected argument 'version_infos' to be a list")
        pulumi.set(__self__, "version_infos", version_infos)

    @property
    @pulumi.getter(name="allocationPolicies")
    def allocation_policies(self) -> Sequence['outputs.GetKubernetesNodeGroupAllocationPolicyResult']:
        """
        This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.
        """
        return pulumi.get(self, "allocation_policies")

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
    @pulumi.getter(name="deployPolicies")
    def deploy_policies(self) -> Sequence['outputs.GetKubernetesNodeGroupDeployPolicyResult']:
        """
        Deploy policy of the node group. The structure is documented below.
        """
        return pulumi.get(self, "deploy_policies")

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
    @pulumi.getter(name="instanceTemplates")
    def instance_templates(self) -> Sequence['outputs.GetKubernetesNodeGroupInstanceTemplateResult']:
        """
        Template used to create compute instances in this Kubernetes node group. The structure is documented below.
        """
        return pulumi.get(self, "instance_templates")

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
    @pulumi.getter(name="maintenancePolicies")
    def maintenance_policies(self) -> Sequence['outputs.GetKubernetesNodeGroupMaintenancePolicyResult']:
        """
        Information about maintenance policy for this Kubernetes node group. The structure is documented below.
        """
        return pulumi.get(self, "maintenance_policies")

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
    @pulumi.getter(name="scalePolicies")
    def scale_policies(self) -> Sequence['outputs.GetKubernetesNodeGroupScalePolicyResult']:
        """
        Scale policy of the node group. The structure is documented below.
        """
        return pulumi.get(self, "scale_policies")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the Kubernetes node group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="versionInfos")
    def version_infos(self) -> Sequence['outputs.GetKubernetesNodeGroupVersionInfoResult']:
        """
        Information about Kubernetes node group version. The structure is documented below.
        """
        return pulumi.get(self, "version_infos")


class AwaitableGetKubernetesNodeGroupResult(GetKubernetesNodeGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesNodeGroupResult(
            allocation_policies=self.allocation_policies,
            allowed_unsafe_sysctls=self.allowed_unsafe_sysctls,
            cluster_id=self.cluster_id,
            created_at=self.created_at,
            deploy_policies=self.deploy_policies,
            description=self.description,
            folder_id=self.folder_id,
            id=self.id,
            instance_group_id=self.instance_group_id,
            instance_templates=self.instance_templates,
            labels=self.labels,
            maintenance_policies=self.maintenance_policies,
            name=self.name,
            node_group_id=self.node_group_id,
            node_labels=self.node_labels,
            node_taints=self.node_taints,
            scale_policies=self.scale_policies,
            status=self.status,
            version_infos=self.version_infos)


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
        allocation_policies=__ret__.allocation_policies,
        allowed_unsafe_sysctls=__ret__.allowed_unsafe_sysctls,
        cluster_id=__ret__.cluster_id,
        created_at=__ret__.created_at,
        deploy_policies=__ret__.deploy_policies,
        description=__ret__.description,
        folder_id=__ret__.folder_id,
        id=__ret__.id,
        instance_group_id=__ret__.instance_group_id,
        instance_templates=__ret__.instance_templates,
        labels=__ret__.labels,
        maintenance_policies=__ret__.maintenance_policies,
        name=__ret__.name,
        node_group_id=__ret__.node_group_id,
        node_labels=__ret__.node_labels,
        node_taints=__ret__.node_taints,
        scale_policies=__ret__.scale_policies,
        status=__ret__.status,
        version_infos=__ret__.version_infos)


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
