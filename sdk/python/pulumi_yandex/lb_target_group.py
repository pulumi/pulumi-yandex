# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LbTargetGroupArgs', 'LbTargetGroup']

@pulumi.input_type
class LbTargetGroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input['LbTargetGroupTargetArgs']]]] = None):
        """
        The set of arguments for constructing a LbTargetGroup resource.
        :param pulumi.Input[str] description: An optional description of the target group. Provide this property when
               you create the resource.
        :param pulumi.Input[str] folder_id: The ID of the folder to which the resource belongs.
               If omitted, the provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this target group. A list of key/value pairs.
        :param pulumi.Input[str] name: Name of the target group. Provided by the client when the target group is created.
        :param pulumi.Input[str] region_id: ID of the availability zone where the target group resides. 
               The default is 'ru-central1'.
        :param pulumi.Input[Sequence[pulumi.Input['LbTargetGroupTargetArgs']]] targets: A Target resource. The structure is documented below.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the target group. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the folder to which the resource belongs.
        If omitted, the provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to assign to this target group. A list of key/value pairs.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the target group. Provided by the client when the target group is created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the availability zone where the target group resides. 
        The default is 'ru-central1'.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LbTargetGroupTargetArgs']]]]:
        """
        A Target resource. The structure is documented below.
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LbTargetGroupTargetArgs']]]]):
        pulumi.set(self, "targets", value)


@pulumi.input_type
class _LbTargetGroupState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input['LbTargetGroupTargetArgs']]]] = None):
        """
        Input properties used for looking up and filtering LbTargetGroup resources.
        :param pulumi.Input[str] created_at: The target group creation timestamp.
        :param pulumi.Input[str] description: An optional description of the target group. Provide this property when
               you create the resource.
        :param pulumi.Input[str] folder_id: The ID of the folder to which the resource belongs.
               If omitted, the provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this target group. A list of key/value pairs.
        :param pulumi.Input[str] name: Name of the target group. Provided by the client when the target group is created.
        :param pulumi.Input[str] region_id: ID of the availability zone where the target group resides. 
               The default is 'ru-central1'.
        :param pulumi.Input[Sequence[pulumi.Input['LbTargetGroupTargetArgs']]] targets: A Target resource. The structure is documented below.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The target group creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the target group. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the folder to which the resource belongs.
        If omitted, the provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to assign to this target group. A list of key/value pairs.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the target group. Provided by the client when the target group is created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the availability zone where the target group resides. 
        The default is 'ru-central1'.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LbTargetGroupTargetArgs']]]]:
        """
        A Target resource. The structure is documented below.
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LbTargetGroupTargetArgs']]]]):
        pulumi.set(self, "targets", value)


class LbTargetGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbTargetGroupTargetArgs']]]]] = None,
                 __props__=None):
        """
        Creates a target group in the specified folder and adds the specified targets to it.
        For more information, see [the official documentation](https://cloud.yandex.com/docs/load-balancer/concepts/target-resources).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        foo = yandex.LbTargetGroup("foo",
            region_id="ru-central1",
            targets=[
                yandex.LbTargetGroupTargetArgs(
                    address=yandex_compute_instance["my-instance-1"]["network_interface"][0]["ip_address"],
                    subnet_id=yandex_vpc_subnet["my-subnet"]["id"],
                ),
                yandex.LbTargetGroupTargetArgs(
                    address=yandex_compute_instance["my-instance-2"]["network_interface"][0]["ip_address"],
                    subnet_id=yandex_vpc_subnet["my-subnet"]["id"],
                ),
            ])
        ```

        ## Import

        A target group can be imported using the `id` of the resource, e.g.

        ```sh
         $ pulumi import yandex:index/lbTargetGroup:LbTargetGroup default target_group_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of the target group. Provide this property when
               you create the resource.
        :param pulumi.Input[str] folder_id: The ID of the folder to which the resource belongs.
               If omitted, the provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this target group. A list of key/value pairs.
        :param pulumi.Input[str] name: Name of the target group. Provided by the client when the target group is created.
        :param pulumi.Input[str] region_id: ID of the availability zone where the target group resides. 
               The default is 'ru-central1'.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbTargetGroupTargetArgs']]]] targets: A Target resource. The structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LbTargetGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a target group in the specified folder and adds the specified targets to it.
        For more information, see [the official documentation](https://cloud.yandex.com/docs/load-balancer/concepts/target-resources).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        foo = yandex.LbTargetGroup("foo",
            region_id="ru-central1",
            targets=[
                yandex.LbTargetGroupTargetArgs(
                    address=yandex_compute_instance["my-instance-1"]["network_interface"][0]["ip_address"],
                    subnet_id=yandex_vpc_subnet["my-subnet"]["id"],
                ),
                yandex.LbTargetGroupTargetArgs(
                    address=yandex_compute_instance["my-instance-2"]["network_interface"][0]["ip_address"],
                    subnet_id=yandex_vpc_subnet["my-subnet"]["id"],
                ),
            ])
        ```

        ## Import

        A target group can be imported using the `id` of the resource, e.g.

        ```sh
         $ pulumi import yandex:index/lbTargetGroup:LbTargetGroup default target_group_id
        ```

        :param str resource_name: The name of the resource.
        :param LbTargetGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LbTargetGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbTargetGroupTargetArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LbTargetGroupArgs.__new__(LbTargetGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["region_id"] = region_id
            __props__.__dict__["targets"] = targets
            __props__.__dict__["created_at"] = None
        super(LbTargetGroup, __self__).__init__(
            'yandex:index/lbTargetGroup:LbTargetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region_id: Optional[pulumi.Input[str]] = None,
            targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbTargetGroupTargetArgs']]]]] = None) -> 'LbTargetGroup':
        """
        Get an existing LbTargetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The target group creation timestamp.
        :param pulumi.Input[str] description: An optional description of the target group. Provide this property when
               you create the resource.
        :param pulumi.Input[str] folder_id: The ID of the folder to which the resource belongs.
               If omitted, the provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this target group. A list of key/value pairs.
        :param pulumi.Input[str] name: Name of the target group. Provided by the client when the target group is created.
        :param pulumi.Input[str] region_id: ID of the availability zone where the target group resides. 
               The default is 'ru-central1'.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbTargetGroupTargetArgs']]]] targets: A Target resource. The structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LbTargetGroupState.__new__(_LbTargetGroupState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["region_id"] = region_id
        __props__.__dict__["targets"] = targets
        return LbTargetGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The target group creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of the target group. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        The ID of the folder to which the resource belongs.
        If omitted, the provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels to assign to this target group. A list of key/value pairs.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the target group. Provided by the client when the target group is created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the availability zone where the target group resides. 
        The default is 'ru-central1'.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Optional[Sequence['outputs.LbTargetGroupTarget']]]:
        """
        A Target resource. The structure is documented below.
        """
        return pulumi.get(self, "targets")

