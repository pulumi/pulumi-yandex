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

__all__ = [
    'GetIamPolicyResult',
    'AwaitableGetIamPolicyResult',
    'get_iam_policy',
]

@pulumi.output_type
class GetIamPolicyResult:
    """
    A collection of values returned by getIamPolicy.
    """
    def __init__(__self__, bindings=None, id=None, policy_data=None):
        if bindings and not isinstance(bindings, list):
            raise TypeError("Expected argument 'bindings' to be a list")
        pulumi.set(__self__, "bindings", bindings)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter
    def bindings(self) -> Sequence['outputs.GetIamPolicyBindingResult']:
        return pulumi.get(self, "bindings")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> str:
        """
        The above bindings serialized in a format suitable for
        referencing from a resource that supports IAM.
        """
        return pulumi.get(self, "policy_data")


class AwaitableGetIamPolicyResult(GetIamPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIamPolicyResult(
            bindings=self.bindings,
            id=self.id,
            policy_data=self.policy_data)


def get_iam_policy(bindings: Optional[Sequence[pulumi.InputType['GetIamPolicyBindingArgs']]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIamPolicyResult:
    """
    Generates an [IAM] policy document that may be referenced by and applied to
    other Yandex.Cloud Platform resources, such as the `ResourcemanagerFolder` resource.

    This data source is used to define [IAM] policies to apply to other resources.
    Currently, defining a policy through a data source and referencing that policy
    from another resource is the only way to apply an IAM policy to a resource.


    :param Sequence[pulumi.InputType['GetIamPolicyBindingArgs']] bindings: A nested configuration block (described below)
           that defines a binding to be included in the policy document. Multiple
           `binding` arguments are supported.
    """
    __args__ = dict()
    __args__['bindings'] = bindings
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('yandex:index/getIamPolicy:getIamPolicy', __args__, opts=opts, typ=GetIamPolicyResult).value

    return AwaitableGetIamPolicyResult(
        bindings=__ret__.bindings,
        id=__ret__.id,
        policy_data=__ret__.policy_data)
