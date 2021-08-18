// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetIamPolicy
    {
        /// <summary>
        /// Generates an [IAM] policy document that may be referenced by and applied to
        /// other Yandex.Cloud Platform resources, such as the `yandex.ResourcemanagerFolder` resource. 
        /// 
        /// 
        /// 
        /// This data source is used to define [IAM] policies to apply to other resources.
        /// Currently, defining a policy through a data source and referencing that policy
        /// from another resource is the only way to apply an IAM policy to a resource.
        /// </summary>
        public static Task<GetIamPolicyResult> InvokeAsync(GetIamPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIamPolicyResult>("yandex:index/getIamPolicy:getIamPolicy", args ?? new GetIamPolicyArgs(), options.WithVersion());
    }


    public sealed class GetIamPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("bindings", required: true)]
        private List<Inputs.GetIamPolicyBindingArgs>? _bindings;

        /// <summary>
        /// A nested configuration block (described below)
        /// that defines a binding to be included in the policy document. Multiple
        /// `binding` arguments are supported.
        /// </summary>
        public List<Inputs.GetIamPolicyBindingArgs> Bindings
        {
            get => _bindings ?? (_bindings = new List<Inputs.GetIamPolicyBindingArgs>());
            set => _bindings = value;
        }

        public GetIamPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIamPolicyResult
    {
        public readonly ImmutableArray<Outputs.GetIamPolicyBindingResult> Bindings;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The above bindings serialized in a format suitable for
        /// referencing from a resource that supports IAM.
        /// </summary>
        public readonly string PolicyData;

        [OutputConstructor]
        private GetIamPolicyResult(
            ImmutableArray<Outputs.GetIamPolicyBindingResult> bindings,

            string id,

            string policyData)
        {
            Bindings = bindings;
            Id = id;
            PolicyData = policyData;
        }
    }
}
