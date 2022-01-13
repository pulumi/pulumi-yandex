// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of Yandex Cloud Serverless Containers
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
// 		_, err := yandex.NewServerlessContainer(ctx, "test-container", &yandex.ServerlessContainerArgs{
// 			CoreFraction:     pulumi.Int(100),
// 			Cores:            pulumi.Int(1),
// 			Description:      pulumi.String("any description"),
// 			ExecutionTimeout: pulumi.String("15s"),
// 			Image: &ServerlessContainerImageArgs{
// 				Url: pulumi.String("cr.yandex/yc/test-image:v1"),
// 			},
// 			Memory:           pulumi.Int(256),
// 			ServiceAccountId: pulumi.String("are1service2account3id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
// 		_, err := yandex.NewServerlessContainer(ctx, "test-container-with-digest", &yandex.ServerlessContainerArgs{
// 			Image: &ServerlessContainerImageArgs{
// 				Digest: pulumi.String("sha256:e1d772fa8795adac847a2420c87d0d2e3d38fb02f168cab8c0b5fe2fb95c47f4"),
// 				Url:    pulumi.String("cr.yandex/yc/test-image:v1"),
// 			},
// 			Memory: pulumi.Int(128),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ServerlessContainer struct {
	pulumi.CustomResourceState

	// Concurrency of Yandex Cloud Serverless Container
	Concurrency pulumi.IntPtrOutput `pulumi:"concurrency"`
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container
	CoreFraction pulumi.IntOutput    `pulumi:"coreFraction"`
	Cores        pulumi.IntPtrOutput `pulumi:"cores"`
	// Creation timestamp of the Yandex Cloud Serverless Container
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the Yandex Cloud Serverless Container
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
	ExecutionTimeout pulumi.StringOutput `pulumi:"executionTimeout"`
	// Folder ID for the Yandex Cloud Serverless Container
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Revision deployment image for Yandex Cloud Serverless Container
	// * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
	// * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
	// * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
	//   If presented, should be equal to digest that will be resolved at server side by URL.
	//   Container will be updated on digest change even if `image.0.url` stays the same.
	//   If field not specified then its value will be computed.
	// * `image.0.command` - List of commands for Yandex Cloud Serverless Container
	// * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
	// * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
	Image ServerlessContainerImageOutput `pulumi:"image"`
	// A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
	Memory pulumi.IntOutput `pulumi:"memory"`
	// Yandex Cloud Serverless Container name
	Name pulumi.StringOutput `pulumi:"name"`
	// Last revision ID of the Yandex Cloud Serverless Container
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// Service account ID for Yandex Cloud Serverless Container
	ServiceAccountId pulumi.StringPtrOutput `pulumi:"serviceAccountId"`
	// Invoke URL for the Yandex Cloud Serverless Container
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServerlessContainer registers a new resource with the given unique name, arguments, and options.
func NewServerlessContainer(ctx *pulumi.Context,
	name string, args *ServerlessContainerArgs, opts ...pulumi.ResourceOption) (*ServerlessContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	var resource ServerlessContainer
	err := ctx.RegisterResource("yandex:index/serverlessContainer:ServerlessContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerlessContainer gets an existing ServerlessContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerlessContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerlessContainerState, opts ...pulumi.ResourceOption) (*ServerlessContainer, error) {
	var resource ServerlessContainer
	err := ctx.ReadResource("yandex:index/serverlessContainer:ServerlessContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerlessContainer resources.
type serverlessContainerState struct {
	// Concurrency of Yandex Cloud Serverless Container
	Concurrency *int `pulumi:"concurrency"`
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container
	CoreFraction *int `pulumi:"coreFraction"`
	Cores        *int `pulumi:"cores"`
	// Creation timestamp of the Yandex Cloud Serverless Container
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the Yandex Cloud Serverless Container
	Description *string `pulumi:"description"`
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
	ExecutionTimeout *string `pulumi:"executionTimeout"`
	// Folder ID for the Yandex Cloud Serverless Container
	FolderId *string `pulumi:"folderId"`
	// Revision deployment image for Yandex Cloud Serverless Container
	// * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
	// * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
	// * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
	//   If presented, should be equal to digest that will be resolved at server side by URL.
	//   Container will be updated on digest change even if `image.0.url` stays the same.
	//   If field not specified then its value will be computed.
	// * `image.0.command` - List of commands for Yandex Cloud Serverless Container
	// * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
	// * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
	Image *ServerlessContainerImage `pulumi:"image"`
	// A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
	Labels map[string]string `pulumi:"labels"`
	// Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
	Memory *int `pulumi:"memory"`
	// Yandex Cloud Serverless Container name
	Name *string `pulumi:"name"`
	// Last revision ID of the Yandex Cloud Serverless Container
	RevisionId *string `pulumi:"revisionId"`
	// Service account ID for Yandex Cloud Serverless Container
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// Invoke URL for the Yandex Cloud Serverless Container
	Url *string `pulumi:"url"`
}

type ServerlessContainerState struct {
	// Concurrency of Yandex Cloud Serverless Container
	Concurrency pulumi.IntPtrInput
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container
	CoreFraction pulumi.IntPtrInput
	Cores        pulumi.IntPtrInput
	// Creation timestamp of the Yandex Cloud Serverless Container
	CreatedAt pulumi.StringPtrInput
	// Description of the Yandex Cloud Serverless Container
	Description pulumi.StringPtrInput
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
	ExecutionTimeout pulumi.StringPtrInput
	// Folder ID for the Yandex Cloud Serverless Container
	FolderId pulumi.StringPtrInput
	// Revision deployment image for Yandex Cloud Serverless Container
	// * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
	// * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
	// * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
	//   If presented, should be equal to digest that will be resolved at server side by URL.
	//   Container will be updated on digest change even if `image.0.url` stays the same.
	//   If field not specified then its value will be computed.
	// * `image.0.command` - List of commands for Yandex Cloud Serverless Container
	// * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
	// * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
	Image ServerlessContainerImagePtrInput
	// A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
	Labels pulumi.StringMapInput
	// Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
	Memory pulumi.IntPtrInput
	// Yandex Cloud Serverless Container name
	Name pulumi.StringPtrInput
	// Last revision ID of the Yandex Cloud Serverless Container
	RevisionId pulumi.StringPtrInput
	// Service account ID for Yandex Cloud Serverless Container
	ServiceAccountId pulumi.StringPtrInput
	// Invoke URL for the Yandex Cloud Serverless Container
	Url pulumi.StringPtrInput
}

func (ServerlessContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessContainerState)(nil)).Elem()
}

type serverlessContainerArgs struct {
	// Concurrency of Yandex Cloud Serverless Container
	Concurrency *int `pulumi:"concurrency"`
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container
	CoreFraction *int `pulumi:"coreFraction"`
	Cores        *int `pulumi:"cores"`
	// Description of the Yandex Cloud Serverless Container
	Description *string `pulumi:"description"`
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
	ExecutionTimeout *string `pulumi:"executionTimeout"`
	// Folder ID for the Yandex Cloud Serverless Container
	FolderId *string `pulumi:"folderId"`
	// Revision deployment image for Yandex Cloud Serverless Container
	// * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
	// * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
	// * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
	//   If presented, should be equal to digest that will be resolved at server side by URL.
	//   Container will be updated on digest change even if `image.0.url` stays the same.
	//   If field not specified then its value will be computed.
	// * `image.0.command` - List of commands for Yandex Cloud Serverless Container
	// * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
	// * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
	Image ServerlessContainerImage `pulumi:"image"`
	// A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
	Labels map[string]string `pulumi:"labels"`
	// Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
	Memory int `pulumi:"memory"`
	// Yandex Cloud Serverless Container name
	Name *string `pulumi:"name"`
	// Service account ID for Yandex Cloud Serverless Container
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a ServerlessContainer resource.
type ServerlessContainerArgs struct {
	// Concurrency of Yandex Cloud Serverless Container
	Concurrency pulumi.IntPtrInput
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container
	CoreFraction pulumi.IntPtrInput
	Cores        pulumi.IntPtrInput
	// Description of the Yandex Cloud Serverless Container
	Description pulumi.StringPtrInput
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container
	ExecutionTimeout pulumi.StringPtrInput
	// Folder ID for the Yandex Cloud Serverless Container
	FolderId pulumi.StringPtrInput
	// Revision deployment image for Yandex Cloud Serverless Container
	// * `image.0.url` (Required) - URL of image that will be deployed as Yandex Cloud Serverless Container
	// * `image.0.work_dir` - Working directory for Yandex Cloud Serverless Container
	// * `image.0.digest` - Digest of image that will be deployed as Yandex Cloud Serverless Container.
	//   If presented, should be equal to digest that will be resolved at server side by URL.
	//   Container will be updated on digest change even if `image.0.url` stays the same.
	//   If field not specified then its value will be computed.
	// * `image.0.command` - List of commands for Yandex Cloud Serverless Container
	// * `image.0.args` - List of arguments for Yandex Cloud Serverless Container
	// * `image.0.environment` -  A set of key/value environment variable pairs for Yandex Cloud Serverless Container
	Image ServerlessContainerImageInput
	// A set of key/value label pairs to assign to the Yandex Cloud Serverless Container
	Labels pulumi.StringMapInput
	// Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Serverless Container
	Memory pulumi.IntInput
	// Yandex Cloud Serverless Container name
	Name pulumi.StringPtrInput
	// Service account ID for Yandex Cloud Serverless Container
	ServiceAccountId pulumi.StringPtrInput
}

func (ServerlessContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessContainerArgs)(nil)).Elem()
}

type ServerlessContainerInput interface {
	pulumi.Input

	ToServerlessContainerOutput() ServerlessContainerOutput
	ToServerlessContainerOutputWithContext(ctx context.Context) ServerlessContainerOutput
}

func (*ServerlessContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessContainer)(nil)).Elem()
}

func (i *ServerlessContainer) ToServerlessContainerOutput() ServerlessContainerOutput {
	return i.ToServerlessContainerOutputWithContext(context.Background())
}

func (i *ServerlessContainer) ToServerlessContainerOutputWithContext(ctx context.Context) ServerlessContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessContainerOutput)
}

// ServerlessContainerArrayInput is an input type that accepts ServerlessContainerArray and ServerlessContainerArrayOutput values.
// You can construct a concrete instance of `ServerlessContainerArrayInput` via:
//
//          ServerlessContainerArray{ ServerlessContainerArgs{...} }
type ServerlessContainerArrayInput interface {
	pulumi.Input

	ToServerlessContainerArrayOutput() ServerlessContainerArrayOutput
	ToServerlessContainerArrayOutputWithContext(context.Context) ServerlessContainerArrayOutput
}

type ServerlessContainerArray []ServerlessContainerInput

func (ServerlessContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerlessContainer)(nil)).Elem()
}

func (i ServerlessContainerArray) ToServerlessContainerArrayOutput() ServerlessContainerArrayOutput {
	return i.ToServerlessContainerArrayOutputWithContext(context.Background())
}

func (i ServerlessContainerArray) ToServerlessContainerArrayOutputWithContext(ctx context.Context) ServerlessContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessContainerArrayOutput)
}

// ServerlessContainerMapInput is an input type that accepts ServerlessContainerMap and ServerlessContainerMapOutput values.
// You can construct a concrete instance of `ServerlessContainerMapInput` via:
//
//          ServerlessContainerMap{ "key": ServerlessContainerArgs{...} }
type ServerlessContainerMapInput interface {
	pulumi.Input

	ToServerlessContainerMapOutput() ServerlessContainerMapOutput
	ToServerlessContainerMapOutputWithContext(context.Context) ServerlessContainerMapOutput
}

type ServerlessContainerMap map[string]ServerlessContainerInput

func (ServerlessContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerlessContainer)(nil)).Elem()
}

func (i ServerlessContainerMap) ToServerlessContainerMapOutput() ServerlessContainerMapOutput {
	return i.ToServerlessContainerMapOutputWithContext(context.Background())
}

func (i ServerlessContainerMap) ToServerlessContainerMapOutputWithContext(ctx context.Context) ServerlessContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessContainerMapOutput)
}

type ServerlessContainerOutput struct{ *pulumi.OutputState }

func (ServerlessContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessContainer)(nil)).Elem()
}

func (o ServerlessContainerOutput) ToServerlessContainerOutput() ServerlessContainerOutput {
	return o
}

func (o ServerlessContainerOutput) ToServerlessContainerOutputWithContext(ctx context.Context) ServerlessContainerOutput {
	return o
}

type ServerlessContainerArrayOutput struct{ *pulumi.OutputState }

func (ServerlessContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerlessContainer)(nil)).Elem()
}

func (o ServerlessContainerArrayOutput) ToServerlessContainerArrayOutput() ServerlessContainerArrayOutput {
	return o
}

func (o ServerlessContainerArrayOutput) ToServerlessContainerArrayOutputWithContext(ctx context.Context) ServerlessContainerArrayOutput {
	return o
}

func (o ServerlessContainerArrayOutput) Index(i pulumi.IntInput) ServerlessContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerlessContainer {
		return vs[0].([]*ServerlessContainer)[vs[1].(int)]
	}).(ServerlessContainerOutput)
}

type ServerlessContainerMapOutput struct{ *pulumi.OutputState }

func (ServerlessContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerlessContainer)(nil)).Elem()
}

func (o ServerlessContainerMapOutput) ToServerlessContainerMapOutput() ServerlessContainerMapOutput {
	return o
}

func (o ServerlessContainerMapOutput) ToServerlessContainerMapOutputWithContext(ctx context.Context) ServerlessContainerMapOutput {
	return o
}

func (o ServerlessContainerMapOutput) MapIndex(k pulumi.StringInput) ServerlessContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerlessContainer {
		return vs[0].(map[string]*ServerlessContainer)[vs[1].(string)]
	}).(ServerlessContainerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessContainerInput)(nil)).Elem(), &ServerlessContainer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessContainerArrayInput)(nil)).Elem(), ServerlessContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessContainerMapInput)(nil)).Elem(), ServerlessContainerMap{})
	pulumi.RegisterOutputType(ServerlessContainerOutput{})
	pulumi.RegisterOutputType(ServerlessContainerArrayOutput{})
	pulumi.RegisterOutputType(ServerlessContainerMapOutput{})
}
