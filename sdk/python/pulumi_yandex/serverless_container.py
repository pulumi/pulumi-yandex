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

__all__ = ['ServerlessContainerArgs', 'ServerlessContainer']

@pulumi.input_type
class ServerlessContainerArgs:
    def __init__(__self__, *,
                 image: pulumi.Input['ServerlessContainerImageArgs'],
                 memory: pulumi.Input[int],
                 concurrency: Optional[pulumi.Input[int]] = None,
                 core_fraction: Optional[pulumi.Input[int]] = None,
                 cores: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 execution_timeout: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServerlessContainer resource.
        :param pulumi.Input['ServerlessContainerImageArgs'] image: Revision deployment image for Yandex Cloud Serverless Container
               * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
               * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
               * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
               If presented, should be equal to digest that will be resolved at server side by URL.
               Container will be updated on digest change even if `image.0.url` stays the same.
               If field not specified then its value will be computed.
               * `image.0.command` - List of commands for Yandex Cloud Serverless Container
               * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
               * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
        :param pulumi.Input[int] memory: Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
        :param pulumi.Input[int] concurrency: Concurrency of Yandex Cloud Serverless Container
        :param pulumi.Input[int] core_fraction: Core fraction (**0...100**) of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] description: Description of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] execution_timeout: Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
        :param pulumi.Input[str] folder_id: Folder ID for the Yandex Cloud Serverless Container
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
        :param pulumi.Input[str] name: Yandex Cloud Serverless Container name
        :param pulumi.Input[str] service_account_id: Service account ID for Yandex Cloud Serverless Container
        """
        pulumi.set(__self__, "image", image)
        pulumi.set(__self__, "memory", memory)
        if concurrency is not None:
            pulumi.set(__self__, "concurrency", concurrency)
        if core_fraction is not None:
            pulumi.set(__self__, "core_fraction", core_fraction)
        if cores is not None:
            pulumi.set(__self__, "cores", cores)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if execution_timeout is not None:
            pulumi.set(__self__, "execution_timeout", execution_timeout)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)

    @property
    @pulumi.getter
    def image(self) -> pulumi.Input['ServerlessContainerImageArgs']:
        """
        Revision deployment image for Yandex Cloud Serverless Container
        * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
        * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
        * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
        If presented, should be equal to digest that will be resolved at server side by URL.
        Container will be updated on digest change even if `image.0.url` stays the same.
        If field not specified then its value will be computed.
        * `image.0.command` - List of commands for Yandex Cloud Serverless Container
        * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
        * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "image")

    @image.setter
    def image(self, value: pulumi.Input['ServerlessContainerImageArgs']):
        pulumi.set(self, "image", value)

    @property
    @pulumi.getter
    def memory(self) -> pulumi.Input[int]:
        """
        Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: pulumi.Input[int]):
        pulumi.set(self, "memory", value)

    @property
    @pulumi.getter
    def concurrency(self) -> Optional[pulumi.Input[int]]:
        """
        Concurrency of Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "concurrency")

    @concurrency.setter
    def concurrency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "concurrency", value)

    @property
    @pulumi.getter(name="coreFraction")
    def core_fraction(self) -> Optional[pulumi.Input[int]]:
        """
        Core fraction (**0...100**) of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "core_fraction")

    @core_fraction.setter
    def core_fraction(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "core_fraction", value)

    @property
    @pulumi.getter
    def cores(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "cores")

    @cores.setter
    def cores(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cores", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="executionTimeout")
    def execution_timeout(self) -> Optional[pulumi.Input[str]]:
        """
        Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "execution_timeout")

    @execution_timeout.setter
    def execution_timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "execution_timeout", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        Folder ID for the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Yandex Cloud Serverless Container name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Service account ID for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_id", value)


@pulumi.input_type
class _ServerlessContainerState:
    def __init__(__self__, *,
                 concurrency: Optional[pulumi.Input[int]] = None,
                 core_fraction: Optional[pulumi.Input[int]] = None,
                 cores: Optional[pulumi.Input[int]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 execution_timeout: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input['ServerlessContainerImageArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 memory: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 revision_id: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServerlessContainer resources.
        :param pulumi.Input[int] concurrency: Concurrency of Yandex Cloud Serverless Container
        :param pulumi.Input[int] core_fraction: Core fraction (**0...100**) of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] created_at: Creation timestamp of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] description: Description of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] execution_timeout: Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
        :param pulumi.Input[str] folder_id: Folder ID for the Yandex Cloud Serverless Container
        :param pulumi.Input['ServerlessContainerImageArgs'] image: Revision deployment image for Yandex Cloud Serverless Container
               * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
               * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
               * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
               If presented, should be equal to digest that will be resolved at server side by URL.
               Container will be updated on digest change even if `image.0.url` stays the same.
               If field not specified then its value will be computed.
               * `image.0.command` - List of commands for Yandex Cloud Serverless Container
               * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
               * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
        :param pulumi.Input[int] memory: Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
        :param pulumi.Input[str] name: Yandex Cloud Serverless Container name
        :param pulumi.Input[str] revision_id: Last revision ID of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] service_account_id: Service account ID for Yandex Cloud Serverless Container
        :param pulumi.Input[str] url: Invoke URL for the Yandex Cloud Serverless Container
        """
        if concurrency is not None:
            pulumi.set(__self__, "concurrency", concurrency)
        if core_fraction is not None:
            pulumi.set(__self__, "core_fraction", core_fraction)
        if cores is not None:
            pulumi.set(__self__, "cores", cores)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if execution_timeout is not None:
            pulumi.set(__self__, "execution_timeout", execution_timeout)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if image is not None:
            pulumi.set(__self__, "image", image)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if memory is not None:
            pulumi.set(__self__, "memory", memory)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if revision_id is not None:
            pulumi.set(__self__, "revision_id", revision_id)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def concurrency(self) -> Optional[pulumi.Input[int]]:
        """
        Concurrency of Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "concurrency")

    @concurrency.setter
    def concurrency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "concurrency", value)

    @property
    @pulumi.getter(name="coreFraction")
    def core_fraction(self) -> Optional[pulumi.Input[int]]:
        """
        Core fraction (**0...100**) of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "core_fraction")

    @core_fraction.setter
    def core_fraction(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "core_fraction", value)

    @property
    @pulumi.getter
    def cores(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "cores")

    @cores.setter
    def cores(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cores", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="executionTimeout")
    def execution_timeout(self) -> Optional[pulumi.Input[str]]:
        """
        Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "execution_timeout")

    @execution_timeout.setter
    def execution_timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "execution_timeout", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        Folder ID for the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def image(self) -> Optional[pulumi.Input['ServerlessContainerImageArgs']]:
        """
        Revision deployment image for Yandex Cloud Serverless Container
        * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
        * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
        * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
        If presented, should be equal to digest that will be resolved at server side by URL.
        Container will be updated on digest change even if `image.0.url` stays the same.
        If field not specified then its value will be computed.
        * `image.0.command` - List of commands for Yandex Cloud Serverless Container
        * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
        * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "image")

    @image.setter
    def image(self, value: Optional[pulumi.Input['ServerlessContainerImageArgs']]):
        pulumi.set(self, "image", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def memory(self) -> Optional[pulumi.Input[int]]:
        """
        Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Yandex Cloud Serverless Container name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="revisionId")
    def revision_id(self) -> Optional[pulumi.Input[str]]:
        """
        Last revision ID of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "revision_id")

    @revision_id.setter
    def revision_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "revision_id", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Service account ID for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_id", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Invoke URL for the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ServerlessContainer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 concurrency: Optional[pulumi.Input[int]] = None,
                 core_fraction: Optional[pulumi.Input[int]] = None,
                 cores: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 execution_timeout: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input[pulumi.InputType['ServerlessContainerImageArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 memory: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows management of Yandex Cloud Serverless Containers

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        test_container = yandex.ServerlessContainer("test-container",
            core_fraction=100,
            cores=1,
            description="any description",
            execution_timeout="15s",
            image=yandex.ServerlessContainerImageArgs(
                url="cr.yandex/yc/test-image:v1",
            ),
            memory=256,
            service_account_id="are1service2account3id")
        ```
        ```python
        import pulumi
        import pulumi_yandex as yandex

        test_container_with_digest = yandex.ServerlessContainer("test-container-with-digest",
            image=yandex.ServerlessContainerImageArgs(
                digest="sha256:e1d772fa8795adac847a2420c87d0d2e3d38fb02f168cab8c0b5fe2fb95c47f4",
                url="cr.yandex/yc/test-image:v1",
            ),
            memory=128)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] concurrency: Concurrency of Yandex Cloud Serverless Container
        :param pulumi.Input[int] core_fraction: Core fraction (**0...100**) of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] description: Description of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] execution_timeout: Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
        :param pulumi.Input[str] folder_id: Folder ID for the Yandex Cloud Serverless Container
        :param pulumi.Input[pulumi.InputType['ServerlessContainerImageArgs']] image: Revision deployment image for Yandex Cloud Serverless Container
               * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
               * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
               * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
               If presented, should be equal to digest that will be resolved at server side by URL.
               Container will be updated on digest change even if `image.0.url` stays the same.
               If field not specified then its value will be computed.
               * `image.0.command` - List of commands for Yandex Cloud Serverless Container
               * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
               * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
        :param pulumi.Input[int] memory: Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
        :param pulumi.Input[str] name: Yandex Cloud Serverless Container name
        :param pulumi.Input[str] service_account_id: Service account ID for Yandex Cloud Serverless Container
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerlessContainerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows management of Yandex Cloud Serverless Containers

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        test_container = yandex.ServerlessContainer("test-container",
            core_fraction=100,
            cores=1,
            description="any description",
            execution_timeout="15s",
            image=yandex.ServerlessContainerImageArgs(
                url="cr.yandex/yc/test-image:v1",
            ),
            memory=256,
            service_account_id="are1service2account3id")
        ```
        ```python
        import pulumi
        import pulumi_yandex as yandex

        test_container_with_digest = yandex.ServerlessContainer("test-container-with-digest",
            image=yandex.ServerlessContainerImageArgs(
                digest="sha256:e1d772fa8795adac847a2420c87d0d2e3d38fb02f168cab8c0b5fe2fb95c47f4",
                url="cr.yandex/yc/test-image:v1",
            ),
            memory=128)
        ```

        :param str resource_name: The name of the resource.
        :param ServerlessContainerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerlessContainerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 concurrency: Optional[pulumi.Input[int]] = None,
                 core_fraction: Optional[pulumi.Input[int]] = None,
                 cores: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 execution_timeout: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input[pulumi.InputType['ServerlessContainerImageArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 memory: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServerlessContainerArgs.__new__(ServerlessContainerArgs)

            __props__.__dict__["concurrency"] = concurrency
            __props__.__dict__["core_fraction"] = core_fraction
            __props__.__dict__["cores"] = cores
            __props__.__dict__["description"] = description
            __props__.__dict__["execution_timeout"] = execution_timeout
            __props__.__dict__["folder_id"] = folder_id
            if image is None and not opts.urn:
                raise TypeError("Missing required property 'image'")
            __props__.__dict__["image"] = image
            __props__.__dict__["labels"] = labels
            if memory is None and not opts.urn:
                raise TypeError("Missing required property 'memory'")
            __props__.__dict__["memory"] = memory
            __props__.__dict__["name"] = name
            __props__.__dict__["service_account_id"] = service_account_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["revision_id"] = None
            __props__.__dict__["url"] = None
        super(ServerlessContainer, __self__).__init__(
            'yandex:index/serverlessContainer:ServerlessContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            concurrency: Optional[pulumi.Input[int]] = None,
            core_fraction: Optional[pulumi.Input[int]] = None,
            cores: Optional[pulumi.Input[int]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            execution_timeout: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            image: Optional[pulumi.Input[pulumi.InputType['ServerlessContainerImageArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            memory: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            revision_id: Optional[pulumi.Input[str]] = None,
            service_account_id: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServerlessContainer':
        """
        Get an existing ServerlessContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] concurrency: Concurrency of Yandex Cloud Serverless Container
        :param pulumi.Input[int] core_fraction: Core fraction (**0...100**) of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] created_at: Creation timestamp of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] description: Description of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] execution_timeout: Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
        :param pulumi.Input[str] folder_id: Folder ID for the Yandex Cloud Serverless Container
        :param pulumi.Input[pulumi.InputType['ServerlessContainerImageArgs']] image: Revision deployment image for Yandex Cloud Serverless Container
               * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
               * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
               * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
               If presented, should be equal to digest that will be resolved at server side by URL.
               Container will be updated on digest change even if `image.0.url` stays the same.
               If field not specified then its value will be computed.
               * `image.0.command` - List of commands for Yandex Cloud Serverless Container
               * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
               * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
        :param pulumi.Input[int] memory: Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
        :param pulumi.Input[str] name: Yandex Cloud Serverless Container name
        :param pulumi.Input[str] revision_id: Last revision ID of the Yandex Cloud Serverless Container
        :param pulumi.Input[str] service_account_id: Service account ID for Yandex Cloud Serverless Container
        :param pulumi.Input[str] url: Invoke URL for the Yandex Cloud Serverless Container
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerlessContainerState.__new__(_ServerlessContainerState)

        __props__.__dict__["concurrency"] = concurrency
        __props__.__dict__["core_fraction"] = core_fraction
        __props__.__dict__["cores"] = cores
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["execution_timeout"] = execution_timeout
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["image"] = image
        __props__.__dict__["labels"] = labels
        __props__.__dict__["memory"] = memory
        __props__.__dict__["name"] = name
        __props__.__dict__["revision_id"] = revision_id
        __props__.__dict__["service_account_id"] = service_account_id
        __props__.__dict__["url"] = url
        return ServerlessContainer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def concurrency(self) -> pulumi.Output[Optional[int]]:
        """
        Concurrency of Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "concurrency")

    @property
    @pulumi.getter(name="coreFraction")
    def core_fraction(self) -> pulumi.Output[int]:
        """
        Core fraction (**0...100**) of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "core_fraction")

    @property
    @pulumi.getter
    def cores(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "cores")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Creation timestamp of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="executionTimeout")
    def execution_timeout(self) -> pulumi.Output[str]:
        """
        Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "execution_timeout")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        Folder ID for the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def image(self) -> pulumi.Output['outputs.ServerlessContainerImage']:
        """
        Revision deployment image for Yandex Cloud Serverless Container
        * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
        * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
        * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
        If presented, should be equal to digest that will be resolved at server side by URL.
        Container will be updated on digest change even if `image.0.url` stays the same.
        If field not specified then its value will be computed.
        * `image.0.command` - List of commands for Yandex Cloud Serverless Container
        * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
        * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def memory(self) -> pulumi.Output[int]:
        """
        Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Yandex Cloud Serverless Container name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="revisionId")
    def revision_id(self) -> pulumi.Output[str]:
        """
        Last revision ID of the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "revision_id")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        Service account ID for Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "service_account_id")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        Invoke URL for the Yandex Cloud Serverless Container
        """
        return pulumi.get(self, "url")

