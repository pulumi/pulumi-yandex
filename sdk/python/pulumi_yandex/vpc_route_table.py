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

__all__ = ['VpcRouteTableArgs', 'VpcRouteTable']

@pulumi.input_type
class VpcRouteTableArgs:
    def __init__(__self__, *,
                 network_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 static_routes: Optional[pulumi.Input[Sequence[pulumi.Input['VpcRouteTableStaticRouteArgs']]]] = None):
        """
        The set of arguments for constructing a VpcRouteTable resource.
        :param pulumi.Input[str] network_id: ID of the network this route table belongs to.
        :param pulumi.Input[str] description: An optional description of the route table. Provide this property when
               you create the resource.
        :param pulumi.Input[str] folder_id: The ID of the folder to which the resource belongs.
               If omitted, the provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this route table. A list of key/value pairs.
        :param pulumi.Input[str] name: Name of the route table. Provided by the client when the route table is created.
        :param pulumi.Input[Sequence[pulumi.Input['VpcRouteTableStaticRouteArgs']]] static_routes: A list of static route records for the route table. The structure is documented below.
        """
        pulumi.set(__self__, "network_id", network_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if static_routes is not None:
            pulumi.set(__self__, "static_routes", static_routes)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        """
        ID of the network this route table belongs to.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the route table. Provide this property when
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
        Labels to assign to this route table. A list of key/value pairs.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the route table. Provided by the client when the route table is created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="staticRoutes")
    def static_routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcRouteTableStaticRouteArgs']]]]:
        """
        A list of static route records for the route table. The structure is documented below.
        """
        return pulumi.get(self, "static_routes")

    @static_routes.setter
    def static_routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcRouteTableStaticRouteArgs']]]]):
        pulumi.set(self, "static_routes", value)


@pulumi.input_type
class _VpcRouteTableState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 static_routes: Optional[pulumi.Input[Sequence[pulumi.Input['VpcRouteTableStaticRouteArgs']]]] = None):
        """
        Input properties used for looking up and filtering VpcRouteTable resources.
        :param pulumi.Input[str] created_at: Creation timestamp of the route table.
        :param pulumi.Input[str] description: An optional description of the route table. Provide this property when
               you create the resource.
        :param pulumi.Input[str] folder_id: The ID of the folder to which the resource belongs.
               If omitted, the provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this route table. A list of key/value pairs.
        :param pulumi.Input[str] name: Name of the route table. Provided by the client when the route table is created.
        :param pulumi.Input[str] network_id: ID of the network this route table belongs to.
        :param pulumi.Input[Sequence[pulumi.Input['VpcRouteTableStaticRouteArgs']]] static_routes: A list of static route records for the route table. The structure is documented below.
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
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if static_routes is not None:
            pulumi.set(__self__, "static_routes", static_routes)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp of the route table.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the route table. Provide this property when
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
        Labels to assign to this route table. A list of key/value pairs.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the route table. Provided by the client when the route table is created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the network this route table belongs to.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="staticRoutes")
    def static_routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcRouteTableStaticRouteArgs']]]]:
        """
        A list of static route records for the route table. The structure is documented below.
        """
        return pulumi.get(self, "static_routes")

    @static_routes.setter
    def static_routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcRouteTableStaticRouteArgs']]]]):
        pulumi.set(self, "static_routes", value)


class VpcRouteTable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 static_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcRouteTableStaticRouteArgs']]]]] = None,
                 __props__=None):
        """
        Manages a route table within the Yandex.Cloud. For more information, see
        [the official documentation](https://cloud.yandex.com/docs/vpc/concepts).

        * How-to Guides
            * [Cloud Networking](https://cloud.yandex.com/docs/vpc/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        lab_net = yandex.VpcNetwork("lab-net")
        lab_rt_a = yandex.VpcRouteTable("lab-rt-a",
            network_id=lab_net.id,
            static_routes=[yandex.VpcRouteTableStaticRouteArgs(
                destination_prefix="10.2.0.0/16",
                next_hop_address="172.16.10.10",
            )])
        ```

        ## Import

        A route table can be imported using the `id` of the resource, e.g.

        ```sh
         $ pulumi import yandex:index/vpcRouteTable:VpcRouteTable default route_table_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of the route table. Provide this property when
               you create the resource.
        :param pulumi.Input[str] folder_id: The ID of the folder to which the resource belongs.
               If omitted, the provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this route table. A list of key/value pairs.
        :param pulumi.Input[str] name: Name of the route table. Provided by the client when the route table is created.
        :param pulumi.Input[str] network_id: ID of the network this route table belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcRouteTableStaticRouteArgs']]]] static_routes: A list of static route records for the route table. The structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcRouteTableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a route table within the Yandex.Cloud. For more information, see
        [the official documentation](https://cloud.yandex.com/docs/vpc/concepts).

        * How-to Guides
            * [Cloud Networking](https://cloud.yandex.com/docs/vpc/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        lab_net = yandex.VpcNetwork("lab-net")
        lab_rt_a = yandex.VpcRouteTable("lab-rt-a",
            network_id=lab_net.id,
            static_routes=[yandex.VpcRouteTableStaticRouteArgs(
                destination_prefix="10.2.0.0/16",
                next_hop_address="172.16.10.10",
            )])
        ```

        ## Import

        A route table can be imported using the `id` of the resource, e.g.

        ```sh
         $ pulumi import yandex:index/vpcRouteTable:VpcRouteTable default route_table_id
        ```

        :param str resource_name: The name of the resource.
        :param VpcRouteTableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcRouteTableArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 network_id: Optional[pulumi.Input[str]] = None,
                 static_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcRouteTableStaticRouteArgs']]]]] = None,
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
            __props__ = VpcRouteTableArgs.__new__(VpcRouteTableArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["static_routes"] = static_routes
            __props__.__dict__["created_at"] = None
        super(VpcRouteTable, __self__).__init__(
            'yandex:index/vpcRouteTable:VpcRouteTable',
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
            network_id: Optional[pulumi.Input[str]] = None,
            static_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcRouteTableStaticRouteArgs']]]]] = None) -> 'VpcRouteTable':
        """
        Get an existing VpcRouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: Creation timestamp of the route table.
        :param pulumi.Input[str] description: An optional description of the route table. Provide this property when
               you create the resource.
        :param pulumi.Input[str] folder_id: The ID of the folder to which the resource belongs.
               If omitted, the provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this route table. A list of key/value pairs.
        :param pulumi.Input[str] name: Name of the route table. Provided by the client when the route table is created.
        :param pulumi.Input[str] network_id: ID of the network this route table belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcRouteTableStaticRouteArgs']]]] static_routes: A list of static route records for the route table. The structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcRouteTableState.__new__(_VpcRouteTableState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["static_routes"] = static_routes
        return VpcRouteTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Creation timestamp of the route table.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of the route table. Provide this property when
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
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Labels to assign to this route table. A list of key/value pairs.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the route table. Provided by the client when the route table is created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        ID of the network this route table belongs to.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="staticRoutes")
    def static_routes(self) -> pulumi.Output[Optional[Sequence['outputs.VpcRouteTableStaticRoute']]]:
        """
        A list of static route records for the route table. The structure is documented below.
        """
        return pulumi.get(self, "static_routes")

