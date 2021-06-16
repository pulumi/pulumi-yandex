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

__all__ = ['AlbBackendGroupArgs', 'AlbBackendGroup']

@pulumi.input_type
class AlbBackendGroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 grpc_backends: Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupGrpcBackendArgs']]]] = None,
                 http_backends: Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupHttpBackendArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AlbBackendGroup resource.
        :param pulumi.Input[str] description: Description of the backend group.
        :param pulumi.Input[str] folder_id: Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        :param pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupGrpcBackendArgs']]] grpc_backends: Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupHttpBackendArgs']]] http_backends: Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this backend group.
        :param pulumi.Input[str] name: Name of the backend.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if grpc_backends is not None:
            pulumi.set(__self__, "grpc_backends", grpc_backends)
        if http_backends is not None:
            pulumi.set(__self__, "http_backends", http_backends)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the backend group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter(name="grpcBackends")
    def grpc_backends(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupGrpcBackendArgs']]]]:
        """
        Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
        """
        return pulumi.get(self, "grpc_backends")

    @grpc_backends.setter
    def grpc_backends(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupGrpcBackendArgs']]]]):
        pulumi.set(self, "grpc_backends", value)

    @property
    @pulumi.getter(name="httpBackends")
    def http_backends(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupHttpBackendArgs']]]]:
        """
        Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
        """
        return pulumi.get(self, "http_backends")

    @http_backends.setter
    def http_backends(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupHttpBackendArgs']]]]):
        pulumi.set(self, "http_backends", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to assign to this backend group.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the backend.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AlbBackendGroupState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 grpc_backends: Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupGrpcBackendArgs']]]] = None,
                 http_backends: Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupHttpBackendArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AlbBackendGroup resources.
        :param pulumi.Input[str] created_at: The backend group creation timestamp.
        :param pulumi.Input[str] description: Description of the backend group.
        :param pulumi.Input[str] folder_id: Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        :param pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupGrpcBackendArgs']]] grpc_backends: Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupHttpBackendArgs']]] http_backends: Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this backend group.
        :param pulumi.Input[str] name: Name of the backend.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if grpc_backends is not None:
            pulumi.set(__self__, "grpc_backends", grpc_backends)
        if http_backends is not None:
            pulumi.set(__self__, "http_backends", http_backends)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The backend group creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the backend group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter(name="grpcBackends")
    def grpc_backends(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupGrpcBackendArgs']]]]:
        """
        Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
        """
        return pulumi.get(self, "grpc_backends")

    @grpc_backends.setter
    def grpc_backends(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupGrpcBackendArgs']]]]):
        pulumi.set(self, "grpc_backends", value)

    @property
    @pulumi.getter(name="httpBackends")
    def http_backends(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupHttpBackendArgs']]]]:
        """
        Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
        """
        return pulumi.get(self, "http_backends")

    @http_backends.setter
    def http_backends(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlbBackendGroupHttpBackendArgs']]]]):
        pulumi.set(self, "http_backends", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to assign to this backend group.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the backend.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class AlbBackendGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 grpc_backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupGrpcBackendArgs']]]]] = None,
                 http_backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupHttpBackendArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a backend group in the specified folder and adds the specified backends to it.
        For more information, see [the official documentation](https://cloud.yandex.com/en/docs/application-load-balancer/concepts/backend-group).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        test_backend_group = yandex.AlbBackendGroup("test-backend-group", http_backends=[yandex.AlbBackendGroupHttpBackendArgs(
            healthcheck=yandex.AlbBackendGroupHttpBackendHealthcheckArgs(
                http_healthcheck=yandex.AlbBackendGroupHttpBackendHealthcheckHttpHealthcheckArgs(
                    path="/",
                ),
                interval="1s",
                timeout="1s",
            ),
            http2=True,
            load_balancing_config=yandex.AlbBackendGroupHttpBackendLoadBalancingConfigArgs(
                panic_threshold=50,
            ),
            name="test-http-backend",
            port=8080,
            target_group_ids=[yandex_alb_target_group["test-target-group"]["id"]],
            tls=yandex.AlbBackendGroupHttpBackendTlsArgs(
                sni="backend-domain.internal",
            ),
            weight=1,
        )])
        ```

        ## Import

        A backend group can be imported using the `id` of the resource, e.g.

        ```sh
         $ pulumi import yandex:index/albBackendGroup:AlbBackendGroup default backend_group_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the backend group.
        :param pulumi.Input[str] folder_id: Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupGrpcBackendArgs']]]] grpc_backends: Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupHttpBackendArgs']]]] http_backends: Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this backend group.
        :param pulumi.Input[str] name: Name of the backend.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AlbBackendGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a backend group in the specified folder and adds the specified backends to it.
        For more information, see [the official documentation](https://cloud.yandex.com/en/docs/application-load-balancer/concepts/backend-group).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        test_backend_group = yandex.AlbBackendGroup("test-backend-group", http_backends=[yandex.AlbBackendGroupHttpBackendArgs(
            healthcheck=yandex.AlbBackendGroupHttpBackendHealthcheckArgs(
                http_healthcheck=yandex.AlbBackendGroupHttpBackendHealthcheckHttpHealthcheckArgs(
                    path="/",
                ),
                interval="1s",
                timeout="1s",
            ),
            http2=True,
            load_balancing_config=yandex.AlbBackendGroupHttpBackendLoadBalancingConfigArgs(
                panic_threshold=50,
            ),
            name="test-http-backend",
            port=8080,
            target_group_ids=[yandex_alb_target_group["test-target-group"]["id"]],
            tls=yandex.AlbBackendGroupHttpBackendTlsArgs(
                sni="backend-domain.internal",
            ),
            weight=1,
        )])
        ```

        ## Import

        A backend group can be imported using the `id` of the resource, e.g.

        ```sh
         $ pulumi import yandex:index/albBackendGroup:AlbBackendGroup default backend_group_id
        ```

        :param str resource_name: The name of the resource.
        :param AlbBackendGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlbBackendGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 grpc_backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupGrpcBackendArgs']]]]] = None,
                 http_backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupHttpBackendArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = AlbBackendGroupArgs.__new__(AlbBackendGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["grpc_backends"] = grpc_backends
            __props__.__dict__["http_backends"] = http_backends
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["created_at"] = None
        super(AlbBackendGroup, __self__).__init__(
            'yandex:index/albBackendGroup:AlbBackendGroup',
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
            grpc_backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupGrpcBackendArgs']]]]] = None,
            http_backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupHttpBackendArgs']]]]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'AlbBackendGroup':
        """
        Get an existing AlbBackendGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The backend group creation timestamp.
        :param pulumi.Input[str] description: Description of the backend group.
        :param pulumi.Input[str] folder_id: Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupGrpcBackendArgs']]]] grpc_backends: Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlbBackendGroupHttpBackendArgs']]]] http_backends: Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to assign to this backend group.
        :param pulumi.Input[str] name: Name of the backend.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlbBackendGroupState.__new__(_AlbBackendGroupState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["grpc_backends"] = grpc_backends
        __props__.__dict__["http_backends"] = http_backends
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        return AlbBackendGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The backend group creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the backend group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="grpcBackends")
    def grpc_backends(self) -> pulumi.Output[Optional[Sequence['outputs.AlbBackendGroupGrpcBackend']]]:
        """
        Grpc backend specification that will be used by the ALB Backend Group. Structure is documented below.
        """
        return pulumi.get(self, "grpc_backends")

    @property
    @pulumi.getter(name="httpBackends")
    def http_backends(self) -> pulumi.Output[Optional[Sequence['outputs.AlbBackendGroupHttpBackend']]]:
        """
        Http backend specification that will be used by the ALB Backend Group. Structure is documented below.
        """
        return pulumi.get(self, "http_backends")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels to assign to this backend group.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the backend.
        """
        return pulumi.get(self, "name")

