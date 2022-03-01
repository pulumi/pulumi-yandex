// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex SAML Federation. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/organization/add-federation).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		federation, err := yandex.LookupOrganizationmanagerSamlFederation(ctx, &GetOrganizationmanagerSamlFederationArgs{
// 			FederationId:   pulumi.StringRef("some_federation_id"),
// 			OrganizationId: pulumi.StringRef("some_organization_id"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("myFederation.name", federation.Name)
// 		return nil
// 	})
// }
// ```
func LookupOrganizationmanagerSamlFederation(ctx *pulumi.Context, args *LookupOrganizationmanagerSamlFederationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationmanagerSamlFederationResult, error) {
	var rv LookupOrganizationmanagerSamlFederationResult
	err := ctx.Invoke("yandex:index/getOrganizationmanagerSamlFederation:getOrganizationmanagerSamlFederation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationmanagerSamlFederation.
type LookupOrganizationmanagerSamlFederationArgs struct {
	// ID of a SAML Federation.
	FederationId *string `pulumi:"federationId"`
	// A set of key/value label pairs assigned to the SAML Federation.
	Labels map[string]string `pulumi:"labels"`
	// Name of a SAML Federation.
	Name *string `pulumi:"name"`
	// Organization that the federation belongs to. If value is omitted, the default provider organization is used.
	OrganizationId *string `pulumi:"organizationId"`
}

// A collection of values returned by getOrganizationmanagerSamlFederation.
type LookupOrganizationmanagerSamlFederationResult struct {
	// Indicates whether new users get added automatically on successful authentication.
	AutoCreateAccountOnLogin bool `pulumi:"autoCreateAccountOnLogin"`
	// Indicates whether case-insensitive name ids are in use.
	CaseInsensitiveNameIds bool `pulumi:"caseInsensitiveNameIds"`
	// The lifetime of a Browser cookie in seconds.
	CookieMaxAge string `pulumi:"cookieMaxAge"`
	// The SAML Federation creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// The description of the SAML Federation.
	Description  string `pulumi:"description"`
	FederationId string `pulumi:"federationId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the IdP server used for authentication.
	Issuer string `pulumi:"issuer"`
	// A set of key/value label pairs assigned to the SAML Federation.
	Labels         map[string]string `pulumi:"labels"`
	Name           string            `pulumi:"name"`
	OrganizationId *string           `pulumi:"organizationId"`
	// Federation security settings, structure is documented below.
	SecuritySettings []GetOrganizationmanagerSamlFederationSecuritySetting `pulumi:"securitySettings"`
	// Single sign-on endpoint binding type.
	SsoBinding string `pulumi:"ssoBinding"`
	// Single sign-on endpoint URL.
	SsoUrl string `pulumi:"ssoUrl"`
}

func LookupOrganizationmanagerSamlFederationOutput(ctx *pulumi.Context, args LookupOrganizationmanagerSamlFederationOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationmanagerSamlFederationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationmanagerSamlFederationResult, error) {
			args := v.(LookupOrganizationmanagerSamlFederationArgs)
			r, err := LookupOrganizationmanagerSamlFederation(ctx, &args, opts...)
			return *r, err
		}).(LookupOrganizationmanagerSamlFederationResultOutput)
}

// A collection of arguments for invoking getOrganizationmanagerSamlFederation.
type LookupOrganizationmanagerSamlFederationOutputArgs struct {
	// ID of a SAML Federation.
	FederationId pulumi.StringPtrInput `pulumi:"federationId"`
	// A set of key/value label pairs assigned to the SAML Federation.
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// Name of a SAML Federation.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Organization that the federation belongs to. If value is omitted, the default provider organization is used.
	OrganizationId pulumi.StringPtrInput `pulumi:"organizationId"`
}

func (LookupOrganizationmanagerSamlFederationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationmanagerSamlFederationArgs)(nil)).Elem()
}

// A collection of values returned by getOrganizationmanagerSamlFederation.
type LookupOrganizationmanagerSamlFederationResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationmanagerSamlFederationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationmanagerSamlFederationResult)(nil)).Elem()
}

func (o LookupOrganizationmanagerSamlFederationResultOutput) ToLookupOrganizationmanagerSamlFederationResultOutput() LookupOrganizationmanagerSamlFederationResultOutput {
	return o
}

func (o LookupOrganizationmanagerSamlFederationResultOutput) ToLookupOrganizationmanagerSamlFederationResultOutputWithContext(ctx context.Context) LookupOrganizationmanagerSamlFederationResultOutput {
	return o
}

// Indicates whether new users get added automatically on successful authentication.
func (o LookupOrganizationmanagerSamlFederationResultOutput) AutoCreateAccountOnLogin() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) bool { return v.AutoCreateAccountOnLogin }).(pulumi.BoolOutput)
}

// Indicates whether case-insensitive name ids are in use.
func (o LookupOrganizationmanagerSamlFederationResultOutput) CaseInsensitiveNameIds() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) bool { return v.CaseInsensitiveNameIds }).(pulumi.BoolOutput)
}

// The lifetime of a Browser cookie in seconds.
func (o LookupOrganizationmanagerSamlFederationResultOutput) CookieMaxAge() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.CookieMaxAge }).(pulumi.StringOutput)
}

// The SAML Federation creation timestamp.
func (o LookupOrganizationmanagerSamlFederationResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the SAML Federation.
func (o LookupOrganizationmanagerSamlFederationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupOrganizationmanagerSamlFederationResultOutput) FederationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.FederationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrganizationmanagerSamlFederationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the IdP server used for authentication.
func (o LookupOrganizationmanagerSamlFederationResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// A set of key/value label pairs assigned to the SAML Federation.
func (o LookupOrganizationmanagerSamlFederationResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupOrganizationmanagerSamlFederationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrganizationmanagerSamlFederationResultOutput) OrganizationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) *string { return v.OrganizationId }).(pulumi.StringPtrOutput)
}

// Federation security settings, structure is documented below.
func (o LookupOrganizationmanagerSamlFederationResultOutput) SecuritySettings() GetOrganizationmanagerSamlFederationSecuritySettingArrayOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) []GetOrganizationmanagerSamlFederationSecuritySetting {
		return v.SecuritySettings
	}).(GetOrganizationmanagerSamlFederationSecuritySettingArrayOutput)
}

// Single sign-on endpoint binding type.
func (o LookupOrganizationmanagerSamlFederationResultOutput) SsoBinding() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.SsoBinding }).(pulumi.StringOutput)
}

// Single sign-on endpoint URL.
func (o LookupOrganizationmanagerSamlFederationResultOutput) SsoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationmanagerSamlFederationResult) string { return v.SsoUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationmanagerSamlFederationResultOutput{})
}
