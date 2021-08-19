// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of [Yandex Cloud API Gateway](https://cloud.yandex.com/docs/api-gateway/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := yandex.NewApiGateway(ctx, "test_api_gateway", &yandex.ApiGatewayArgs{
// 			Description: pulumi.String("any description"),
// 			Labels: pulumi.StringMap{
// 				"label":       pulumi.String("label"),
// 				"empty-label": pulumi.String(""),
// 			},
// 			Spec: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "openapi: \"3.0.0\"\n", "info:\n", "  version: 1.0.0\n", "  title: Test API\n", "paths:\n", "  /hello:\n", "    get:\n", "      summary: Say hello\n", "      operationId: hello\n", "      parameters:\n", "        - name: user\n", "          in: query\n", "          description: User name to appear in greetings\n", "          required: false\n", "          schema:\n", "            type: string\n", "            default: 'world'\n", "      responses:\n", "        '200':\n", "          description: Greeting\n", "          content:\n", "            'text/plain':\n", "              schema:\n", "                type: \"string\"\n", "      x-yc-apigateway-integration:\n", "        type: dummy\n", "        http_code: 200\n", "        http_headers:\n", "          'Content-Type': \"text/plain\"\n", "        content:\n", "          'text/plain': \"Hello again, {user}!\\n\"\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ApiGateway struct {
	pulumi.CustomResourceState

	// Creation timestamp of the Yandex Cloud API Gateway.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the Yandex Cloud API Gateway.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Default domain for the Yandex API Gateway. Generated at creation time.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
	Labels     pulumi.StringMapOutput `pulumi:"labels"`
	LogGroupId pulumi.StringOutput    `pulumi:"logGroupId"`
	// Yandex Cloud API Gateway name used to define API Gateway.
	Name pulumi.StringOutput `pulumi:"name"`
	// OpenAPI specification for Yandex API Gateway.
	Spec pulumi.StringOutput `pulumi:"spec"`
	// Status of the Yandex API Gateway.
	Status pulumi.StringOutput `pulumi:"status"`
	// Set of user domains attached to Yandex API Gateway.
	UserDomains pulumi.StringArrayOutput `pulumi:"userDomains"`
}

// NewApiGateway registers a new resource with the given unique name, arguments, and options.
func NewApiGateway(ctx *pulumi.Context,
	name string, args *ApiGatewayArgs, opts ...pulumi.ResourceOption) (*ApiGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	var resource ApiGateway
	err := ctx.RegisterResource("yandex:index/apiGateway:ApiGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiGateway gets an existing ApiGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiGatewayState, opts ...pulumi.ResourceOption) (*ApiGateway, error) {
	var resource ApiGateway
	err := ctx.ReadResource("yandex:index/apiGateway:ApiGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiGateway resources.
type apiGatewayState struct {
	// Creation timestamp of the Yandex Cloud API Gateway.
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the Yandex Cloud API Gateway.
	Description *string `pulumi:"description"`
	// Default domain for the Yandex API Gateway. Generated at creation time.
	Domain *string `pulumi:"domain"`
	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
	Labels     map[string]string `pulumi:"labels"`
	LogGroupId *string           `pulumi:"logGroupId"`
	// Yandex Cloud API Gateway name used to define API Gateway.
	Name *string `pulumi:"name"`
	// OpenAPI specification for Yandex API Gateway.
	Spec *string `pulumi:"spec"`
	// Status of the Yandex API Gateway.
	Status *string `pulumi:"status"`
	// Set of user domains attached to Yandex API Gateway.
	UserDomains []string `pulumi:"userDomains"`
}

type ApiGatewayState struct {
	// Creation timestamp of the Yandex Cloud API Gateway.
	CreatedAt pulumi.StringPtrInput
	// Description of the Yandex Cloud API Gateway.
	Description pulumi.StringPtrInput
	// Default domain for the Yandex API Gateway. Generated at creation time.
	Domain pulumi.StringPtrInput
	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
	Labels     pulumi.StringMapInput
	LogGroupId pulumi.StringPtrInput
	// Yandex Cloud API Gateway name used to define API Gateway.
	Name pulumi.StringPtrInput
	// OpenAPI specification for Yandex API Gateway.
	Spec pulumi.StringPtrInput
	// Status of the Yandex API Gateway.
	Status pulumi.StringPtrInput
	// Set of user domains attached to Yandex API Gateway.
	UserDomains pulumi.StringArrayInput
}

func (ApiGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayState)(nil)).Elem()
}

type apiGatewayArgs struct {
	// Description of the Yandex Cloud API Gateway.
	Description *string `pulumi:"description"`
	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
	Labels map[string]string `pulumi:"labels"`
	// Yandex Cloud API Gateway name used to define API Gateway.
	Name *string `pulumi:"name"`
	// OpenAPI specification for Yandex API Gateway.
	Spec string `pulumi:"spec"`
}

// The set of arguments for constructing a ApiGateway resource.
type ApiGatewayArgs struct {
	// Description of the Yandex Cloud API Gateway.
	Description pulumi.StringPtrInput
	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
	Labels pulumi.StringMapInput
	// Yandex Cloud API Gateway name used to define API Gateway.
	Name pulumi.StringPtrInput
	// OpenAPI specification for Yandex API Gateway.
	Spec pulumi.StringInput
}

func (ApiGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayArgs)(nil)).Elem()
}

type ApiGatewayInput interface {
	pulumi.Input

	ToApiGatewayOutput() ApiGatewayOutput
	ToApiGatewayOutputWithContext(ctx context.Context) ApiGatewayOutput
}

func (*ApiGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiGateway)(nil))
}

func (i *ApiGateway) ToApiGatewayOutput() ApiGatewayOutput {
	return i.ToApiGatewayOutputWithContext(context.Background())
}

func (i *ApiGateway) ToApiGatewayOutputWithContext(ctx context.Context) ApiGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayOutput)
}

func (i *ApiGateway) ToApiGatewayPtrOutput() ApiGatewayPtrOutput {
	return i.ToApiGatewayPtrOutputWithContext(context.Background())
}

func (i *ApiGateway) ToApiGatewayPtrOutputWithContext(ctx context.Context) ApiGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayPtrOutput)
}

type ApiGatewayPtrInput interface {
	pulumi.Input

	ToApiGatewayPtrOutput() ApiGatewayPtrOutput
	ToApiGatewayPtrOutputWithContext(ctx context.Context) ApiGatewayPtrOutput
}

type apiGatewayPtrType ApiGatewayArgs

func (*apiGatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGateway)(nil))
}

func (i *apiGatewayPtrType) ToApiGatewayPtrOutput() ApiGatewayPtrOutput {
	return i.ToApiGatewayPtrOutputWithContext(context.Background())
}

func (i *apiGatewayPtrType) ToApiGatewayPtrOutputWithContext(ctx context.Context) ApiGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayPtrOutput)
}

// ApiGatewayArrayInput is an input type that accepts ApiGatewayArray and ApiGatewayArrayOutput values.
// You can construct a concrete instance of `ApiGatewayArrayInput` via:
//
//          ApiGatewayArray{ ApiGatewayArgs{...} }
type ApiGatewayArrayInput interface {
	pulumi.Input

	ToApiGatewayArrayOutput() ApiGatewayArrayOutput
	ToApiGatewayArrayOutputWithContext(context.Context) ApiGatewayArrayOutput
}

type ApiGatewayArray []ApiGatewayInput

func (ApiGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApiGateway)(nil))
}

func (i ApiGatewayArray) ToApiGatewayArrayOutput() ApiGatewayArrayOutput {
	return i.ToApiGatewayArrayOutputWithContext(context.Background())
}

func (i ApiGatewayArray) ToApiGatewayArrayOutputWithContext(ctx context.Context) ApiGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayArrayOutput)
}

// ApiGatewayMapInput is an input type that accepts ApiGatewayMap and ApiGatewayMapOutput values.
// You can construct a concrete instance of `ApiGatewayMapInput` via:
//
//          ApiGatewayMap{ "key": ApiGatewayArgs{...} }
type ApiGatewayMapInput interface {
	pulumi.Input

	ToApiGatewayMapOutput() ApiGatewayMapOutput
	ToApiGatewayMapOutputWithContext(context.Context) ApiGatewayMapOutput
}

type ApiGatewayMap map[string]ApiGatewayInput

func (ApiGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApiGateway)(nil))
}

func (i ApiGatewayMap) ToApiGatewayMapOutput() ApiGatewayMapOutput {
	return i.ToApiGatewayMapOutputWithContext(context.Background())
}

func (i ApiGatewayMap) ToApiGatewayMapOutputWithContext(ctx context.Context) ApiGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayMapOutput)
}

type ApiGatewayOutput struct {
	*pulumi.OutputState
}

func (ApiGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiGateway)(nil))
}

func (o ApiGatewayOutput) ToApiGatewayOutput() ApiGatewayOutput {
	return o
}

func (o ApiGatewayOutput) ToApiGatewayOutputWithContext(ctx context.Context) ApiGatewayOutput {
	return o
}

func (o ApiGatewayOutput) ToApiGatewayPtrOutput() ApiGatewayPtrOutput {
	return o.ToApiGatewayPtrOutputWithContext(context.Background())
}

func (o ApiGatewayOutput) ToApiGatewayPtrOutputWithContext(ctx context.Context) ApiGatewayPtrOutput {
	return o.ApplyT(func(v ApiGateway) *ApiGateway {
		return &v
	}).(ApiGatewayPtrOutput)
}

type ApiGatewayPtrOutput struct {
	*pulumi.OutputState
}

func (ApiGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGateway)(nil))
}

func (o ApiGatewayPtrOutput) ToApiGatewayPtrOutput() ApiGatewayPtrOutput {
	return o
}

func (o ApiGatewayPtrOutput) ToApiGatewayPtrOutputWithContext(ctx context.Context) ApiGatewayPtrOutput {
	return o
}

type ApiGatewayArrayOutput struct{ *pulumi.OutputState }

func (ApiGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiGateway)(nil))
}

func (o ApiGatewayArrayOutput) ToApiGatewayArrayOutput() ApiGatewayArrayOutput {
	return o
}

func (o ApiGatewayArrayOutput) ToApiGatewayArrayOutputWithContext(ctx context.Context) ApiGatewayArrayOutput {
	return o
}

func (o ApiGatewayArrayOutput) Index(i pulumi.IntInput) ApiGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiGateway {
		return vs[0].([]ApiGateway)[vs[1].(int)]
	}).(ApiGatewayOutput)
}

type ApiGatewayMapOutput struct{ *pulumi.OutputState }

func (ApiGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiGateway)(nil))
}

func (o ApiGatewayMapOutput) ToApiGatewayMapOutput() ApiGatewayMapOutput {
	return o
}

func (o ApiGatewayMapOutput) ToApiGatewayMapOutputWithContext(ctx context.Context) ApiGatewayMapOutput {
	return o
}

func (o ApiGatewayMapOutput) MapIndex(k pulumi.StringInput) ApiGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiGateway {
		return vs[0].(map[string]ApiGateway)[vs[1].(string)]
	}).(ApiGatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiGatewayOutput{})
	pulumi.RegisterOutputType(ApiGatewayPtrOutput{})
	pulumi.RegisterOutputType(ApiGatewayArrayOutput{})
	pulumi.RegisterOutputType(ApiGatewayMapOutput{})
}
