// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DNS Recordset.
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
// 		foo, err := yandex.NewVpcNetwork(ctx, "foo", nil)
// 		if err != nil {
// 			return err
// 		}
// 		zone1, err := yandex.NewDnsZone(ctx, "zone1", &yandex.DnsZoneArgs{
// 			Description: pulumi.String("desc"),
// 			Labels: pulumi.StringMap{
// 				"label1": pulumi.String("label-1-value"),
// 			},
// 			Zone:   pulumi.String("example.com."),
// 			Public: pulumi.Bool(false),
// 			PrivateNetworks: pulumi.StringArray{
// 				foo.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewDnsRecordSet(ctx, "rs1", &yandex.DnsRecordSetArgs{
// 			ZoneId: zone1.ID(),
// 			Type:   pulumi.String("A"),
// 			Ttl:    pulumi.Int(200),
// 			Datas: pulumi.StringArray{
// 				pulumi.String("10.1.0.1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewDnsRecordSet(ctx, "rs2", &yandex.DnsRecordSetArgs{
// 			ZoneId: zone1.ID(),
// 			Type:   pulumi.String("A"),
// 			Ttl:    pulumi.Int(200),
// 			Datas: pulumi.StringArray{
// 				pulumi.String("10.1.0.2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// DNS recordset can be imported using this format
//
// ```sh
//  $ pulumi import yandex:index/dnsRecordSet:DnsRecordSet rs1 {{zone_id}}/{{name}}/{{type}}
// ```
type DnsRecordSet struct {
	pulumi.CustomResourceState

	// The string data for the records in this record set.
	Datas pulumi.StringArrayOutput `pulumi:"datas"`
	// The DNS name this record set will apply to.
	Name pulumi.StringOutput `pulumi:"name"`
	// The time-to-live of this record set (seconds).
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The DNS record set type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The id of the zone in which this record set will reside.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewDnsRecordSet registers a new resource with the given unique name, arguments, and options.
func NewDnsRecordSet(ctx *pulumi.Context,
	name string, args *DnsRecordSetArgs, opts ...pulumi.ResourceOption) (*DnsRecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Datas == nil {
		return nil, errors.New("invalid value for required argument 'Datas'")
	}
	if args.Ttl == nil {
		return nil, errors.New("invalid value for required argument 'Ttl'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource DnsRecordSet
	err := ctx.RegisterResource("yandex:index/dnsRecordSet:DnsRecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsRecordSet gets an existing DnsRecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsRecordSetState, opts ...pulumi.ResourceOption) (*DnsRecordSet, error) {
	var resource DnsRecordSet
	err := ctx.ReadResource("yandex:index/dnsRecordSet:DnsRecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsRecordSet resources.
type dnsRecordSetState struct {
	// The string data for the records in this record set.
	Datas []string `pulumi:"datas"`
	// The DNS name this record set will apply to.
	Name *string `pulumi:"name"`
	// The time-to-live of this record set (seconds).
	Ttl *int `pulumi:"ttl"`
	// The DNS record set type.
	Type *string `pulumi:"type"`
	// The id of the zone in which this record set will reside.
	ZoneId *string `pulumi:"zoneId"`
}

type DnsRecordSetState struct {
	// The string data for the records in this record set.
	Datas pulumi.StringArrayInput
	// The DNS name this record set will apply to.
	Name pulumi.StringPtrInput
	// The time-to-live of this record set (seconds).
	Ttl pulumi.IntPtrInput
	// The DNS record set type.
	Type pulumi.StringPtrInput
	// The id of the zone in which this record set will reside.
	ZoneId pulumi.StringPtrInput
}

func (DnsRecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordSetState)(nil)).Elem()
}

type dnsRecordSetArgs struct {
	// The string data for the records in this record set.
	Datas []string `pulumi:"datas"`
	// The DNS name this record set will apply to.
	Name *string `pulumi:"name"`
	// The time-to-live of this record set (seconds).
	Ttl int `pulumi:"ttl"`
	// The DNS record set type.
	Type string `pulumi:"type"`
	// The id of the zone in which this record set will reside.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a DnsRecordSet resource.
type DnsRecordSetArgs struct {
	// The string data for the records in this record set.
	Datas pulumi.StringArrayInput
	// The DNS name this record set will apply to.
	Name pulumi.StringPtrInput
	// The time-to-live of this record set (seconds).
	Ttl pulumi.IntInput
	// The DNS record set type.
	Type pulumi.StringInput
	// The id of the zone in which this record set will reside.
	ZoneId pulumi.StringInput
}

func (DnsRecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordSetArgs)(nil)).Elem()
}

type DnsRecordSetInput interface {
	pulumi.Input

	ToDnsRecordSetOutput() DnsRecordSetOutput
	ToDnsRecordSetOutputWithContext(ctx context.Context) DnsRecordSetOutput
}

func (*DnsRecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecordSet)(nil)).Elem()
}

func (i *DnsRecordSet) ToDnsRecordSetOutput() DnsRecordSetOutput {
	return i.ToDnsRecordSetOutputWithContext(context.Background())
}

func (i *DnsRecordSet) ToDnsRecordSetOutputWithContext(ctx context.Context) DnsRecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordSetOutput)
}

// DnsRecordSetArrayInput is an input type that accepts DnsRecordSetArray and DnsRecordSetArrayOutput values.
// You can construct a concrete instance of `DnsRecordSetArrayInput` via:
//
//          DnsRecordSetArray{ DnsRecordSetArgs{...} }
type DnsRecordSetArrayInput interface {
	pulumi.Input

	ToDnsRecordSetArrayOutput() DnsRecordSetArrayOutput
	ToDnsRecordSetArrayOutputWithContext(context.Context) DnsRecordSetArrayOutput
}

type DnsRecordSetArray []DnsRecordSetInput

func (DnsRecordSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecordSet)(nil)).Elem()
}

func (i DnsRecordSetArray) ToDnsRecordSetArrayOutput() DnsRecordSetArrayOutput {
	return i.ToDnsRecordSetArrayOutputWithContext(context.Background())
}

func (i DnsRecordSetArray) ToDnsRecordSetArrayOutputWithContext(ctx context.Context) DnsRecordSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordSetArrayOutput)
}

// DnsRecordSetMapInput is an input type that accepts DnsRecordSetMap and DnsRecordSetMapOutput values.
// You can construct a concrete instance of `DnsRecordSetMapInput` via:
//
//          DnsRecordSetMap{ "key": DnsRecordSetArgs{...} }
type DnsRecordSetMapInput interface {
	pulumi.Input

	ToDnsRecordSetMapOutput() DnsRecordSetMapOutput
	ToDnsRecordSetMapOutputWithContext(context.Context) DnsRecordSetMapOutput
}

type DnsRecordSetMap map[string]DnsRecordSetInput

func (DnsRecordSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecordSet)(nil)).Elem()
}

func (i DnsRecordSetMap) ToDnsRecordSetMapOutput() DnsRecordSetMapOutput {
	return i.ToDnsRecordSetMapOutputWithContext(context.Background())
}

func (i DnsRecordSetMap) ToDnsRecordSetMapOutputWithContext(ctx context.Context) DnsRecordSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordSetMapOutput)
}

type DnsRecordSetOutput struct{ *pulumi.OutputState }

func (DnsRecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecordSet)(nil)).Elem()
}

func (o DnsRecordSetOutput) ToDnsRecordSetOutput() DnsRecordSetOutput {
	return o
}

func (o DnsRecordSetOutput) ToDnsRecordSetOutputWithContext(ctx context.Context) DnsRecordSetOutput {
	return o
}

type DnsRecordSetArrayOutput struct{ *pulumi.OutputState }

func (DnsRecordSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecordSet)(nil)).Elem()
}

func (o DnsRecordSetArrayOutput) ToDnsRecordSetArrayOutput() DnsRecordSetArrayOutput {
	return o
}

func (o DnsRecordSetArrayOutput) ToDnsRecordSetArrayOutputWithContext(ctx context.Context) DnsRecordSetArrayOutput {
	return o
}

func (o DnsRecordSetArrayOutput) Index(i pulumi.IntInput) DnsRecordSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsRecordSet {
		return vs[0].([]*DnsRecordSet)[vs[1].(int)]
	}).(DnsRecordSetOutput)
}

type DnsRecordSetMapOutput struct{ *pulumi.OutputState }

func (DnsRecordSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecordSet)(nil)).Elem()
}

func (o DnsRecordSetMapOutput) ToDnsRecordSetMapOutput() DnsRecordSetMapOutput {
	return o
}

func (o DnsRecordSetMapOutput) ToDnsRecordSetMapOutputWithContext(ctx context.Context) DnsRecordSetMapOutput {
	return o
}

func (o DnsRecordSetMapOutput) MapIndex(k pulumi.StringInput) DnsRecordSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsRecordSet {
		return vs[0].(map[string]*DnsRecordSet)[vs[1].(string)]
	}).(DnsRecordSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordSetInput)(nil)).Elem(), &DnsRecordSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordSetArrayInput)(nil)).Elem(), DnsRecordSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordSetMapInput)(nil)).Elem(), DnsRecordSetMap{})
	pulumi.RegisterOutputType(DnsRecordSetOutput{})
	pulumi.RegisterOutputType(DnsRecordSetArrayOutput{})
	pulumi.RegisterOutputType(DnsRecordSetMapOutput{})
}
