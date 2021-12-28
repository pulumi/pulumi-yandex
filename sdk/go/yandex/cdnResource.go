// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of [Yandex.Cloud CDN Resource](https://cloud.yandex.ru/docs/cdn/concepts/resource).
//
// > **_NOTE:_**  CDN provider must be activated prior usage of CDN resources, either via UI console or via yc cli command: ```yc cdn provider activate --folder-id <folder-id> --type gcore```
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
// 		_, err := yandex.NewCdnResource(ctx, "myResource", &yandex.CdnResourceArgs{
// 			Cname:          pulumi.String("cdn1.yandex-example.ru"),
// 			Active:         pulumi.Bool(false),
// 			OriginProtocol: pulumi.String("https"),
// 			SecondaryHostnames: pulumi.StringArray{
// 				pulumi.String("cdn-example-1.yandex.ru"),
// 				pulumi.String("cdn-example-2.yandex.ru"),
// 			},
// 			OriginGroupId: pulumi.Any(yandex_cdn_origin_group.Foo_cdn_group_by_id.Id),
// 			Options: &CdnResourceOptionsArgs{
// 				EdgeCacheSettings: pulumi.Int(345600),
// 				IgnoreCookie:      pulumi.Bool(true),
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
// A origin group can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import yandex:index/cdnResource:CdnResource default origin_group_id
// ```
type CdnResource struct {
	pulumi.CustomResourceState

	// Flag to create Resource either in active or disabled state. True - the content from CDN is available to clients.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// CDN endpoint CNAME, must be unique among resources.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// Creation timestamp of the IoT Core Device
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	FolderId  pulumi.StringOutput `pulumi:"folderId"`
	// CDN Resource settings and options to tune CDN edge behavior.
	Options         CdnResourceOptionsOutput `pulumi:"options"`
	OriginGroupId   pulumi.IntPtrOutput      `pulumi:"originGroupId"`
	OriginGroupName pulumi.StringPtrOutput   `pulumi:"originGroupName"`
	OriginProtocol  pulumi.StringPtrOutput   `pulumi:"originProtocol"`
	// list of secondary hostname strings.
	SecondaryHostnames pulumi.StringArrayOutput           `pulumi:"secondaryHostnames"`
	SslCertificate     CdnResourceSslCertificatePtrOutput `pulumi:"sslCertificate"`
	UpdatedAt          pulumi.StringOutput                `pulumi:"updatedAt"`
}

// NewCdnResource registers a new resource with the given unique name, arguments, and options.
func NewCdnResource(ctx *pulumi.Context,
	name string, args *CdnResourceArgs, opts ...pulumi.ResourceOption) (*CdnResource, error) {
	if args == nil {
		args = &CdnResourceArgs{}
	}

	var resource CdnResource
	err := ctx.RegisterResource("yandex:index/cdnResource:CdnResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCdnResource gets an existing CdnResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCdnResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CdnResourceState, opts ...pulumi.ResourceOption) (*CdnResource, error) {
	var resource CdnResource
	err := ctx.ReadResource("yandex:index/cdnResource:CdnResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CdnResource resources.
type cdnResourceState struct {
	// Flag to create Resource either in active or disabled state. True - the content from CDN is available to clients.
	Active *bool `pulumi:"active"`
	// CDN endpoint CNAME, must be unique among resources.
	Cname *string `pulumi:"cname"`
	// Creation timestamp of the IoT Core Device
	CreatedAt *string `pulumi:"createdAt"`
	FolderId  *string `pulumi:"folderId"`
	// CDN Resource settings and options to tune CDN edge behavior.
	Options         *CdnResourceOptions `pulumi:"options"`
	OriginGroupId   *int                `pulumi:"originGroupId"`
	OriginGroupName *string             `pulumi:"originGroupName"`
	OriginProtocol  *string             `pulumi:"originProtocol"`
	// list of secondary hostname strings.
	SecondaryHostnames []string                   `pulumi:"secondaryHostnames"`
	SslCertificate     *CdnResourceSslCertificate `pulumi:"sslCertificate"`
	UpdatedAt          *string                    `pulumi:"updatedAt"`
}

type CdnResourceState struct {
	// Flag to create Resource either in active or disabled state. True - the content from CDN is available to clients.
	Active pulumi.BoolPtrInput
	// CDN endpoint CNAME, must be unique among resources.
	Cname pulumi.StringPtrInput
	// Creation timestamp of the IoT Core Device
	CreatedAt pulumi.StringPtrInput
	FolderId  pulumi.StringPtrInput
	// CDN Resource settings and options to tune CDN edge behavior.
	Options         CdnResourceOptionsPtrInput
	OriginGroupId   pulumi.IntPtrInput
	OriginGroupName pulumi.StringPtrInput
	OriginProtocol  pulumi.StringPtrInput
	// list of secondary hostname strings.
	SecondaryHostnames pulumi.StringArrayInput
	SslCertificate     CdnResourceSslCertificatePtrInput
	UpdatedAt          pulumi.StringPtrInput
}

func (CdnResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cdnResourceState)(nil)).Elem()
}

type cdnResourceArgs struct {
	// Flag to create Resource either in active or disabled state. True - the content from CDN is available to clients.
	Active *bool `pulumi:"active"`
	// CDN endpoint CNAME, must be unique among resources.
	Cname *string `pulumi:"cname"`
	// CDN Resource settings and options to tune CDN edge behavior.
	Options         *CdnResourceOptions `pulumi:"options"`
	OriginGroupId   *int                `pulumi:"originGroupId"`
	OriginGroupName *string             `pulumi:"originGroupName"`
	OriginProtocol  *string             `pulumi:"originProtocol"`
	// list of secondary hostname strings.
	SecondaryHostnames []string                   `pulumi:"secondaryHostnames"`
	SslCertificate     *CdnResourceSslCertificate `pulumi:"sslCertificate"`
	UpdatedAt          *string                    `pulumi:"updatedAt"`
}

// The set of arguments for constructing a CdnResource resource.
type CdnResourceArgs struct {
	// Flag to create Resource either in active or disabled state. True - the content from CDN is available to clients.
	Active pulumi.BoolPtrInput
	// CDN endpoint CNAME, must be unique among resources.
	Cname pulumi.StringPtrInput
	// CDN Resource settings and options to tune CDN edge behavior.
	Options         CdnResourceOptionsPtrInput
	OriginGroupId   pulumi.IntPtrInput
	OriginGroupName pulumi.StringPtrInput
	OriginProtocol  pulumi.StringPtrInput
	// list of secondary hostname strings.
	SecondaryHostnames pulumi.StringArrayInput
	SslCertificate     CdnResourceSslCertificatePtrInput
	UpdatedAt          pulumi.StringPtrInput
}

func (CdnResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cdnResourceArgs)(nil)).Elem()
}

type CdnResourceInput interface {
	pulumi.Input

	ToCdnResourceOutput() CdnResourceOutput
	ToCdnResourceOutputWithContext(ctx context.Context) CdnResourceOutput
}

func (*CdnResource) ElementType() reflect.Type {
	return reflect.TypeOf((*CdnResource)(nil))
}

func (i *CdnResource) ToCdnResourceOutput() CdnResourceOutput {
	return i.ToCdnResourceOutputWithContext(context.Background())
}

func (i *CdnResource) ToCdnResourceOutputWithContext(ctx context.Context) CdnResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnResourceOutput)
}

func (i *CdnResource) ToCdnResourcePtrOutput() CdnResourcePtrOutput {
	return i.ToCdnResourcePtrOutputWithContext(context.Background())
}

func (i *CdnResource) ToCdnResourcePtrOutputWithContext(ctx context.Context) CdnResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnResourcePtrOutput)
}

type CdnResourcePtrInput interface {
	pulumi.Input

	ToCdnResourcePtrOutput() CdnResourcePtrOutput
	ToCdnResourcePtrOutputWithContext(ctx context.Context) CdnResourcePtrOutput
}

type cdnResourcePtrType CdnResourceArgs

func (*cdnResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CdnResource)(nil))
}

func (i *cdnResourcePtrType) ToCdnResourcePtrOutput() CdnResourcePtrOutput {
	return i.ToCdnResourcePtrOutputWithContext(context.Background())
}

func (i *cdnResourcePtrType) ToCdnResourcePtrOutputWithContext(ctx context.Context) CdnResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnResourcePtrOutput)
}

// CdnResourceArrayInput is an input type that accepts CdnResourceArray and CdnResourceArrayOutput values.
// You can construct a concrete instance of `CdnResourceArrayInput` via:
//
//          CdnResourceArray{ CdnResourceArgs{...} }
type CdnResourceArrayInput interface {
	pulumi.Input

	ToCdnResourceArrayOutput() CdnResourceArrayOutput
	ToCdnResourceArrayOutputWithContext(context.Context) CdnResourceArrayOutput
}

type CdnResourceArray []CdnResourceInput

func (CdnResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CdnResource)(nil)).Elem()
}

func (i CdnResourceArray) ToCdnResourceArrayOutput() CdnResourceArrayOutput {
	return i.ToCdnResourceArrayOutputWithContext(context.Background())
}

func (i CdnResourceArray) ToCdnResourceArrayOutputWithContext(ctx context.Context) CdnResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnResourceArrayOutput)
}

// CdnResourceMapInput is an input type that accepts CdnResourceMap and CdnResourceMapOutput values.
// You can construct a concrete instance of `CdnResourceMapInput` via:
//
//          CdnResourceMap{ "key": CdnResourceArgs{...} }
type CdnResourceMapInput interface {
	pulumi.Input

	ToCdnResourceMapOutput() CdnResourceMapOutput
	ToCdnResourceMapOutputWithContext(context.Context) CdnResourceMapOutput
}

type CdnResourceMap map[string]CdnResourceInput

func (CdnResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CdnResource)(nil)).Elem()
}

func (i CdnResourceMap) ToCdnResourceMapOutput() CdnResourceMapOutput {
	return i.ToCdnResourceMapOutputWithContext(context.Background())
}

func (i CdnResourceMap) ToCdnResourceMapOutputWithContext(ctx context.Context) CdnResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnResourceMapOutput)
}

type CdnResourceOutput struct{ *pulumi.OutputState }

func (CdnResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CdnResource)(nil))
}

func (o CdnResourceOutput) ToCdnResourceOutput() CdnResourceOutput {
	return o
}

func (o CdnResourceOutput) ToCdnResourceOutputWithContext(ctx context.Context) CdnResourceOutput {
	return o
}

func (o CdnResourceOutput) ToCdnResourcePtrOutput() CdnResourcePtrOutput {
	return o.ToCdnResourcePtrOutputWithContext(context.Background())
}

func (o CdnResourceOutput) ToCdnResourcePtrOutputWithContext(ctx context.Context) CdnResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CdnResource) *CdnResource {
		return &v
	}).(CdnResourcePtrOutput)
}

type CdnResourcePtrOutput struct{ *pulumi.OutputState }

func (CdnResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CdnResource)(nil))
}

func (o CdnResourcePtrOutput) ToCdnResourcePtrOutput() CdnResourcePtrOutput {
	return o
}

func (o CdnResourcePtrOutput) ToCdnResourcePtrOutputWithContext(ctx context.Context) CdnResourcePtrOutput {
	return o
}

func (o CdnResourcePtrOutput) Elem() CdnResourceOutput {
	return o.ApplyT(func(v *CdnResource) CdnResource {
		if v != nil {
			return *v
		}
		var ret CdnResource
		return ret
	}).(CdnResourceOutput)
}

type CdnResourceArrayOutput struct{ *pulumi.OutputState }

func (CdnResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CdnResource)(nil))
}

func (o CdnResourceArrayOutput) ToCdnResourceArrayOutput() CdnResourceArrayOutput {
	return o
}

func (o CdnResourceArrayOutput) ToCdnResourceArrayOutputWithContext(ctx context.Context) CdnResourceArrayOutput {
	return o
}

func (o CdnResourceArrayOutput) Index(i pulumi.IntInput) CdnResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CdnResource {
		return vs[0].([]CdnResource)[vs[1].(int)]
	}).(CdnResourceOutput)
}

type CdnResourceMapOutput struct{ *pulumi.OutputState }

func (CdnResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CdnResource)(nil))
}

func (o CdnResourceMapOutput) ToCdnResourceMapOutput() CdnResourceMapOutput {
	return o
}

func (o CdnResourceMapOutput) ToCdnResourceMapOutputWithContext(ctx context.Context) CdnResourceMapOutput {
	return o
}

func (o CdnResourceMapOutput) MapIndex(k pulumi.StringInput) CdnResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CdnResource {
		return vs[0].(map[string]CdnResource)[vs[1].(string)]
	}).(CdnResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CdnResourceInput)(nil)).Elem(), &CdnResource{})
	pulumi.RegisterInputType(reflect.TypeOf((*CdnResourcePtrInput)(nil)).Elem(), &CdnResource{})
	pulumi.RegisterInputType(reflect.TypeOf((*CdnResourceArrayInput)(nil)).Elem(), CdnResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CdnResourceMapInput)(nil)).Elem(), CdnResourceMap{})
	pulumi.RegisterOutputType(CdnResourceOutput{})
	pulumi.RegisterOutputType(CdnResourcePtrOutput{})
	pulumi.RegisterOutputType(CdnResourceArrayOutput{})
	pulumi.RegisterOutputType(CdnResourceMapOutput{})
}
