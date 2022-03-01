// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a DNS Zone.
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
// 		foo, err := yandex.LookupDnsZone(ctx, &GetDnsZoneArgs{
// 			DnsZoneId: pulumi.StringRef(yandex_dns_zone.Zone1.Id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("zone", foo.Zone)
// 		return nil
// 	})
// }
// ```
func LookupDnsZone(ctx *pulumi.Context, args *LookupDnsZoneArgs, opts ...pulumi.InvokeOption) (*LookupDnsZoneResult, error) {
	var rv LookupDnsZoneResult
	err := ctx.Invoke("yandex:index/getDnsZone:getDnsZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsZone.
type LookupDnsZoneArgs struct {
	// The ID of the DNS Zone.
	DnsZoneId *string `pulumi:"dnsZoneId"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// - Name of the DNS Zone.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDnsZone.
type LookupDnsZoneResult struct {
	// (Computed) The DNS zone creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// (Computed) Description of the DNS zone.
	Description string `pulumi:"description"`
	DnsZoneId   string `pulumi:"dnsZoneId"`
	// (Computed) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) A set of key/value label pairs to assign to the DNS zone.
	Labels map[string]string `pulumi:"labels"`
	// (Computed) User assigned name of a specific resource. Must be unique within the folder.
	Name string `pulumi:"name"`
	// (Computed) For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
	PrivateNetworks []string `pulumi:"privateNetworks"`
	// (Computed) The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
	Public bool `pulumi:"public"`
	// (Computed) The DNS name of this zone, e.g. "example.com.". Must ends with dot.
	Zone string `pulumi:"zone"`
}

func LookupDnsZoneOutput(ctx *pulumi.Context, args LookupDnsZoneOutputArgs, opts ...pulumi.InvokeOption) LookupDnsZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsZoneResult, error) {
			args := v.(LookupDnsZoneArgs)
			r, err := LookupDnsZone(ctx, &args, opts...)
			return *r, err
		}).(LookupDnsZoneResultOutput)
}

// A collection of arguments for invoking getDnsZone.
type LookupDnsZoneOutputArgs struct {
	// The ID of the DNS Zone.
	DnsZoneId pulumi.StringPtrInput `pulumi:"dnsZoneId"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// - Name of the DNS Zone.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupDnsZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsZoneArgs)(nil)).Elem()
}

// A collection of values returned by getDnsZone.
type LookupDnsZoneResultOutput struct{ *pulumi.OutputState }

func (LookupDnsZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsZoneResult)(nil)).Elem()
}

func (o LookupDnsZoneResultOutput) ToLookupDnsZoneResultOutput() LookupDnsZoneResultOutput {
	return o
}

func (o LookupDnsZoneResultOutput) ToLookupDnsZoneResultOutputWithContext(ctx context.Context) LookupDnsZoneResultOutput {
	return o
}

// (Computed) The DNS zone creation timestamp.
func (o LookupDnsZoneResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// (Computed) Description of the DNS zone.
func (o LookupDnsZoneResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupDnsZoneResultOutput) DnsZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.DnsZoneId }).(pulumi.StringOutput)
}

// (Computed) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
func (o LookupDnsZoneResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDnsZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Computed) A set of key/value label pairs to assign to the DNS zone.
func (o LookupDnsZoneResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// (Computed) User assigned name of a specific resource. Must be unique within the folder.
func (o LookupDnsZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// (Computed) For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
func (o LookupDnsZoneResultOutput) PrivateNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) []string { return v.PrivateNetworks }).(pulumi.StringArrayOutput)
}

// (Computed) The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
func (o LookupDnsZoneResultOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) bool { return v.Public }).(pulumi.BoolOutput)
}

// (Computed) The DNS name of this zone, e.g. "example.com.". Must ends with dot.
func (o LookupDnsZoneResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsZoneResultOutput{})
}
