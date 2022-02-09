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

__all__ = ['OrganizationmanagerSamlFederationArgs', 'OrganizationmanagerSamlFederation']

@pulumi.input_type
class OrganizationmanagerSamlFederationArgs:
    def __init__(__self__, *,
                 issuer: pulumi.Input[str],
                 organization_id: pulumi.Input[str],
                 sso_binding: pulumi.Input[str],
                 sso_url: pulumi.Input[str],
                 auto_create_account_on_login: Optional[pulumi.Input[bool]] = None,
                 case_insensitive_name_ids: Optional[pulumi.Input[bool]] = None,
                 cookie_max_age: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_settings: Optional[pulumi.Input['OrganizationmanagerSamlFederationSecuritySettingsArgs']] = None):
        """
        The set of arguments for constructing a OrganizationmanagerSamlFederation resource.
        :param pulumi.Input[str] issuer: The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
        :param pulumi.Input[str] organization_id: The organization to attach this SAML Federation to.
        :param pulumi.Input[str] sso_binding: Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
        :param pulumi.Input[str] sso_url: Single sign-on endpoint URL. Specify the link to the IdP login page here.
        :param pulumi.Input[bool] auto_create_account_on_login: Add new users automatically on successful authentication. The user will get the `resource-manager.clouds.member` role automatically, but you need to grant other roles to them. If the value is `false`, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
        :param pulumi.Input[bool] case_insensitive_name_ids: Use case-insensitive name ids.
        :param pulumi.Input[str] cookie_max_age: The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is `8h`.
        :param pulumi.Input[str] description: The description of the SAML Federation.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs assigned to the SAML Federation.
        :param pulumi.Input[str] name: The name of the SAML Federation.
        :param pulumi.Input['OrganizationmanagerSamlFederationSecuritySettingsArgs'] security_settings: Federation security settings, structure is documented below.
        """
        pulumi.set(__self__, "issuer", issuer)
        pulumi.set(__self__, "organization_id", organization_id)
        pulumi.set(__self__, "sso_binding", sso_binding)
        pulumi.set(__self__, "sso_url", sso_url)
        if auto_create_account_on_login is not None:
            pulumi.set(__self__, "auto_create_account_on_login", auto_create_account_on_login)
        if case_insensitive_name_ids is not None:
            pulumi.set(__self__, "case_insensitive_name_ids", case_insensitive_name_ids)
        if cookie_max_age is not None:
            pulumi.set(__self__, "cookie_max_age", cookie_max_age)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if security_settings is not None:
            pulumi.set(__self__, "security_settings", security_settings)

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Input[str]:
        """
        The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: pulumi.Input[str]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        """
        The organization to attach this SAML Federation to.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="ssoBinding")
    def sso_binding(self) -> pulumi.Input[str]:
        """
        Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
        """
        return pulumi.get(self, "sso_binding")

    @sso_binding.setter
    def sso_binding(self, value: pulumi.Input[str]):
        pulumi.set(self, "sso_binding", value)

    @property
    @pulumi.getter(name="ssoUrl")
    def sso_url(self) -> pulumi.Input[str]:
        """
        Single sign-on endpoint URL. Specify the link to the IdP login page here.
        """
        return pulumi.get(self, "sso_url")

    @sso_url.setter
    def sso_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "sso_url", value)

    @property
    @pulumi.getter(name="autoCreateAccountOnLogin")
    def auto_create_account_on_login(self) -> Optional[pulumi.Input[bool]]:
        """
        Add new users automatically on successful authentication. The user will get the `resource-manager.clouds.member` role automatically, but you need to grant other roles to them. If the value is `false`, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
        """
        return pulumi.get(self, "auto_create_account_on_login")

    @auto_create_account_on_login.setter
    def auto_create_account_on_login(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_account_on_login", value)

    @property
    @pulumi.getter(name="caseInsensitiveNameIds")
    def case_insensitive_name_ids(self) -> Optional[pulumi.Input[bool]]:
        """
        Use case-insensitive name ids.
        """
        return pulumi.get(self, "case_insensitive_name_ids")

    @case_insensitive_name_ids.setter
    def case_insensitive_name_ids(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "case_insensitive_name_ids", value)

    @property
    @pulumi.getter(name="cookieMaxAge")
    def cookie_max_age(self) -> Optional[pulumi.Input[str]]:
        """
        The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is `8h`.
        """
        return pulumi.get(self, "cookie_max_age")

    @cookie_max_age.setter
    def cookie_max_age(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cookie_max_age", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the SAML Federation.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs assigned to the SAML Federation.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SAML Federation.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="securitySettings")
    def security_settings(self) -> Optional[pulumi.Input['OrganizationmanagerSamlFederationSecuritySettingsArgs']]:
        """
        Federation security settings, structure is documented below.
        """
        return pulumi.get(self, "security_settings")

    @security_settings.setter
    def security_settings(self, value: Optional[pulumi.Input['OrganizationmanagerSamlFederationSecuritySettingsArgs']]):
        pulumi.set(self, "security_settings", value)


@pulumi.input_type
class _OrganizationmanagerSamlFederationState:
    def __init__(__self__, *,
                 auto_create_account_on_login: Optional[pulumi.Input[bool]] = None,
                 case_insensitive_name_ids: Optional[pulumi.Input[bool]] = None,
                 cookie_max_age: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 security_settings: Optional[pulumi.Input['OrganizationmanagerSamlFederationSecuritySettingsArgs']] = None,
                 sso_binding: Optional[pulumi.Input[str]] = None,
                 sso_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrganizationmanagerSamlFederation resources.
        :param pulumi.Input[bool] auto_create_account_on_login: Add new users automatically on successful authentication. The user will get the `resource-manager.clouds.member` role automatically, but you need to grant other roles to them. If the value is `false`, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
        :param pulumi.Input[bool] case_insensitive_name_ids: Use case-insensitive name ids.
        :param pulumi.Input[str] cookie_max_age: The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is `8h`.
        :param pulumi.Input[str] created_at: (Computed) The SAML Federation creation timestamp.
        :param pulumi.Input[str] description: The description of the SAML Federation.
        :param pulumi.Input[str] issuer: The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs assigned to the SAML Federation.
        :param pulumi.Input[str] name: The name of the SAML Federation.
        :param pulumi.Input[str] organization_id: The organization to attach this SAML Federation to.
        :param pulumi.Input['OrganizationmanagerSamlFederationSecuritySettingsArgs'] security_settings: Federation security settings, structure is documented below.
        :param pulumi.Input[str] sso_binding: Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
        :param pulumi.Input[str] sso_url: Single sign-on endpoint URL. Specify the link to the IdP login page here.
        """
        if auto_create_account_on_login is not None:
            pulumi.set(__self__, "auto_create_account_on_login", auto_create_account_on_login)
        if case_insensitive_name_ids is not None:
            pulumi.set(__self__, "case_insensitive_name_ids", case_insensitive_name_ids)
        if cookie_max_age is not None:
            pulumi.set(__self__, "cookie_max_age", cookie_max_age)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if security_settings is not None:
            pulumi.set(__self__, "security_settings", security_settings)
        if sso_binding is not None:
            pulumi.set(__self__, "sso_binding", sso_binding)
        if sso_url is not None:
            pulumi.set(__self__, "sso_url", sso_url)

    @property
    @pulumi.getter(name="autoCreateAccountOnLogin")
    def auto_create_account_on_login(self) -> Optional[pulumi.Input[bool]]:
        """
        Add new users automatically on successful authentication. The user will get the `resource-manager.clouds.member` role automatically, but you need to grant other roles to them. If the value is `false`, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
        """
        return pulumi.get(self, "auto_create_account_on_login")

    @auto_create_account_on_login.setter
    def auto_create_account_on_login(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_account_on_login", value)

    @property
    @pulumi.getter(name="caseInsensitiveNameIds")
    def case_insensitive_name_ids(self) -> Optional[pulumi.Input[bool]]:
        """
        Use case-insensitive name ids.
        """
        return pulumi.get(self, "case_insensitive_name_ids")

    @case_insensitive_name_ids.setter
    def case_insensitive_name_ids(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "case_insensitive_name_ids", value)

    @property
    @pulumi.getter(name="cookieMaxAge")
    def cookie_max_age(self) -> Optional[pulumi.Input[str]]:
        """
        The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is `8h`.
        """
        return pulumi.get(self, "cookie_max_age")

    @cookie_max_age.setter
    def cookie_max_age(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cookie_max_age", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The SAML Federation creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the SAML Federation.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs assigned to the SAML Federation.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SAML Federation.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization to attach this SAML Federation to.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="securitySettings")
    def security_settings(self) -> Optional[pulumi.Input['OrganizationmanagerSamlFederationSecuritySettingsArgs']]:
        """
        Federation security settings, structure is documented below.
        """
        return pulumi.get(self, "security_settings")

    @security_settings.setter
    def security_settings(self, value: Optional[pulumi.Input['OrganizationmanagerSamlFederationSecuritySettingsArgs']]):
        pulumi.set(self, "security_settings", value)

    @property
    @pulumi.getter(name="ssoBinding")
    def sso_binding(self) -> Optional[pulumi.Input[str]]:
        """
        Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
        """
        return pulumi.get(self, "sso_binding")

    @sso_binding.setter
    def sso_binding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sso_binding", value)

    @property
    @pulumi.getter(name="ssoUrl")
    def sso_url(self) -> Optional[pulumi.Input[str]]:
        """
        Single sign-on endpoint URL. Specify the link to the IdP login page here.
        """
        return pulumi.get(self, "sso_url")

    @sso_url.setter
    def sso_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sso_url", value)


class OrganizationmanagerSamlFederation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create_account_on_login: Optional[pulumi.Input[bool]] = None,
                 case_insensitive_name_ids: Optional[pulumi.Input[bool]] = None,
                 cookie_max_age: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 security_settings: Optional[pulumi.Input[pulumi.InputType['OrganizationmanagerSamlFederationSecuritySettingsArgs']]] = None,
                 sso_binding: Optional[pulumi.Input[str]] = None,
                 sso_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows management of a single SAML Federation within an existing Yandex.Cloud Organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        federation = yandex.OrganizationmanagerSamlFederation("federation",
            description="My new SAML federation",
            issuer="my-issuer",
            organization_id="sdf4*********3fr",
            sso_binding="POST",
            sso_url="https://my-sso.url")
        ```

        ## Import

        A Yandex SAML Federation can be imported using the `id` of the resource, e.g.

        ```sh
         $ pulumi import yandex:index/organizationmanagerSamlFederation:OrganizationmanagerSamlFederation federation "federation_id"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_account_on_login: Add new users automatically on successful authentication. The user will get the `resource-manager.clouds.member` role automatically, but you need to grant other roles to them. If the value is `false`, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
        :param pulumi.Input[bool] case_insensitive_name_ids: Use case-insensitive name ids.
        :param pulumi.Input[str] cookie_max_age: The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is `8h`.
        :param pulumi.Input[str] description: The description of the SAML Federation.
        :param pulumi.Input[str] issuer: The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs assigned to the SAML Federation.
        :param pulumi.Input[str] name: The name of the SAML Federation.
        :param pulumi.Input[str] organization_id: The organization to attach this SAML Federation to.
        :param pulumi.Input[pulumi.InputType['OrganizationmanagerSamlFederationSecuritySettingsArgs']] security_settings: Federation security settings, structure is documented below.
        :param pulumi.Input[str] sso_binding: Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
        :param pulumi.Input[str] sso_url: Single sign-on endpoint URL. Specify the link to the IdP login page here.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationmanagerSamlFederationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows management of a single SAML Federation within an existing Yandex.Cloud Organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_yandex as yandex

        federation = yandex.OrganizationmanagerSamlFederation("federation",
            description="My new SAML federation",
            issuer="my-issuer",
            organization_id="sdf4*********3fr",
            sso_binding="POST",
            sso_url="https://my-sso.url")
        ```

        ## Import

        A Yandex SAML Federation can be imported using the `id` of the resource, e.g.

        ```sh
         $ pulumi import yandex:index/organizationmanagerSamlFederation:OrganizationmanagerSamlFederation federation "federation_id"
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationmanagerSamlFederationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationmanagerSamlFederationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create_account_on_login: Optional[pulumi.Input[bool]] = None,
                 case_insensitive_name_ids: Optional[pulumi.Input[bool]] = None,
                 cookie_max_age: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 security_settings: Optional[pulumi.Input[pulumi.InputType['OrganizationmanagerSamlFederationSecuritySettingsArgs']]] = None,
                 sso_binding: Optional[pulumi.Input[str]] = None,
                 sso_url: Optional[pulumi.Input[str]] = None,
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
            __props__ = OrganizationmanagerSamlFederationArgs.__new__(OrganizationmanagerSamlFederationArgs)

            __props__.__dict__["auto_create_account_on_login"] = auto_create_account_on_login
            __props__.__dict__["case_insensitive_name_ids"] = case_insensitive_name_ids
            __props__.__dict__["cookie_max_age"] = cookie_max_age
            __props__.__dict__["description"] = description
            if issuer is None and not opts.urn:
                raise TypeError("Missing required property 'issuer'")
            __props__.__dict__["issuer"] = issuer
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["security_settings"] = security_settings
            if sso_binding is None and not opts.urn:
                raise TypeError("Missing required property 'sso_binding'")
            __props__.__dict__["sso_binding"] = sso_binding
            if sso_url is None and not opts.urn:
                raise TypeError("Missing required property 'sso_url'")
            __props__.__dict__["sso_url"] = sso_url
            __props__.__dict__["created_at"] = None
        super(OrganizationmanagerSamlFederation, __self__).__init__(
            'yandex:index/organizationmanagerSamlFederation:OrganizationmanagerSamlFederation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_create_account_on_login: Optional[pulumi.Input[bool]] = None,
            case_insensitive_name_ids: Optional[pulumi.Input[bool]] = None,
            cookie_max_age: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            issuer: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            security_settings: Optional[pulumi.Input[pulumi.InputType['OrganizationmanagerSamlFederationSecuritySettingsArgs']]] = None,
            sso_binding: Optional[pulumi.Input[str]] = None,
            sso_url: Optional[pulumi.Input[str]] = None) -> 'OrganizationmanagerSamlFederation':
        """
        Get an existing OrganizationmanagerSamlFederation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_account_on_login: Add new users automatically on successful authentication. The user will get the `resource-manager.clouds.member` role automatically, but you need to grant other roles to them. If the value is `false`, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
        :param pulumi.Input[bool] case_insensitive_name_ids: Use case-insensitive name ids.
        :param pulumi.Input[str] cookie_max_age: The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is `8h`.
        :param pulumi.Input[str] created_at: (Computed) The SAML Federation creation timestamp.
        :param pulumi.Input[str] description: The description of the SAML Federation.
        :param pulumi.Input[str] issuer: The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs assigned to the SAML Federation.
        :param pulumi.Input[str] name: The name of the SAML Federation.
        :param pulumi.Input[str] organization_id: The organization to attach this SAML Federation to.
        :param pulumi.Input[pulumi.InputType['OrganizationmanagerSamlFederationSecuritySettingsArgs']] security_settings: Federation security settings, structure is documented below.
        :param pulumi.Input[str] sso_binding: Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
        :param pulumi.Input[str] sso_url: Single sign-on endpoint URL. Specify the link to the IdP login page here.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationmanagerSamlFederationState.__new__(_OrganizationmanagerSamlFederationState)

        __props__.__dict__["auto_create_account_on_login"] = auto_create_account_on_login
        __props__.__dict__["case_insensitive_name_ids"] = case_insensitive_name_ids
        __props__.__dict__["cookie_max_age"] = cookie_max_age
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["issuer"] = issuer
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["security_settings"] = security_settings
        __props__.__dict__["sso_binding"] = sso_binding
        __props__.__dict__["sso_url"] = sso_url
        return OrganizationmanagerSamlFederation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoCreateAccountOnLogin")
    def auto_create_account_on_login(self) -> pulumi.Output[bool]:
        """
        Add new users automatically on successful authentication. The user will get the `resource-manager.clouds.member` role automatically, but you need to grant other roles to them. If the value is `false`, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
        """
        return pulumi.get(self, "auto_create_account_on_login")

    @property
    @pulumi.getter(name="caseInsensitiveNameIds")
    def case_insensitive_name_ids(self) -> pulumi.Output[bool]:
        """
        Use case-insensitive name ids.
        """
        return pulumi.get(self, "case_insensitive_name_ids")

    @property
    @pulumi.getter(name="cookieMaxAge")
    def cookie_max_age(self) -> pulumi.Output[str]:
        """
        The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is `8h`.
        """
        return pulumi.get(self, "cookie_max_age")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        (Computed) The SAML Federation creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the SAML Federation.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Output[str]:
        """
        The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs assigned to the SAML Federation.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the SAML Federation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization to attach this SAML Federation to.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="securitySettings")
    def security_settings(self) -> pulumi.Output['outputs.OrganizationmanagerSamlFederationSecuritySettings']:
        """
        Federation security settings, structure is documented below.
        """
        return pulumi.get(self, "security_settings")

    @property
    @pulumi.getter(name="ssoBinding")
    def sso_binding(self) -> pulumi.Output[str]:
        """
        Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
        """
        return pulumi.get(self, "sso_binding")

    @property
    @pulumi.getter(name="ssoUrl")
    def sso_url(self) -> pulumi.Output[str]:
        """
        Single sign-on endpoint URL. Specify the link to the IdP login page here.
        """
        return pulumi.get(self, "sso_url")

