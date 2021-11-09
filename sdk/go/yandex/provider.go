// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the yandex package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// ID of Yandex.Cloud tenant.
	CloudId pulumi.StringPtrOutput `pulumi:"cloudId"`
	// The API endpoint for Yandex.Cloud SDK client.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// The default folder ID where resources will be placed.
	FolderId pulumi.StringPtrOutput `pulumi:"folderId"`
	// Either the path to or the contents of a Service Account key file in JSON format.
	ServiceAccountKeyFile pulumi.StringPtrOutput `pulumi:"serviceAccountKeyFile"`
	// Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly
	// specified.
	StorageAccessKey pulumi.StringPtrOutput `pulumi:"storageAccessKey"`
	// Yandex.Cloud storage service endpoint. Default is storage.yandexcloud.net
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
	// Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly
	// specified.
	StorageSecretKey pulumi.StringPtrOutput `pulumi:"storageSecretKey"`
	// The access token for API operations.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// Yandex.Cloud Message Queue service access key. Used when a message queue resource doesn't have an access key explicitly
	// specified.
	YmqAccessKey pulumi.StringPtrOutput `pulumi:"ymqAccessKey"`
	// Yandex.Cloud Message Queue service endpoint. Default is message-queue.api.cloud.yandex.net
	YmqEndpoint pulumi.StringPtrOutput `pulumi:"ymqEndpoint"`
	// Yandex.Cloud Message Queue service secret key. Used when a message queue resource doesn't have a secret key explicitly
	// specified.
	YmqSecretKey pulumi.StringPtrOutput `pulumi:"ymqSecretKey"`
	// The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:yandex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// ID of Yandex.Cloud tenant.
	CloudId *string `pulumi:"cloudId"`
	// The API endpoint for Yandex.Cloud SDK client.
	Endpoint *string `pulumi:"endpoint"`
	// The default folder ID where resources will be placed.
	FolderId *string `pulumi:"folderId"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
	Insecure *bool `pulumi:"insecure"`
	// The maximum number of times an API request is being executed. If the API request still fails, an error is thrown.
	MaxRetries *int `pulumi:"maxRetries"`
	// Disable use of TLS. Default value is `false`.
	Plaintext *bool `pulumi:"plaintext"`
	// Either the path to or the contents of a Service Account key file in JSON format.
	ServiceAccountKeyFile *string `pulumi:"serviceAccountKeyFile"`
	// Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly
	// specified.
	StorageAccessKey *string `pulumi:"storageAccessKey"`
	// Yandex.Cloud storage service endpoint. Default is storage.yandexcloud.net
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly
	// specified.
	StorageSecretKey *string `pulumi:"storageSecretKey"`
	// The access token for API operations.
	Token *string `pulumi:"token"`
	// Yandex.Cloud Message Queue service access key. Used when a message queue resource doesn't have an access key explicitly
	// specified.
	YmqAccessKey *string `pulumi:"ymqAccessKey"`
	// Yandex.Cloud Message Queue service endpoint. Default is message-queue.api.cloud.yandex.net
	YmqEndpoint *string `pulumi:"ymqEndpoint"`
	// Yandex.Cloud Message Queue service secret key. Used when a message queue resource doesn't have a secret key explicitly
	// specified.
	YmqSecretKey *string `pulumi:"ymqSecretKey"`
	// The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// ID of Yandex.Cloud tenant.
	CloudId pulumi.StringPtrInput
	// The API endpoint for Yandex.Cloud SDK client.
	Endpoint pulumi.StringPtrInput
	// The default folder ID where resources will be placed.
	FolderId pulumi.StringPtrInput
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
	Insecure pulumi.BoolPtrInput
	// The maximum number of times an API request is being executed. If the API request still fails, an error is thrown.
	MaxRetries pulumi.IntPtrInput
	// Disable use of TLS. Default value is `false`.
	Plaintext pulumi.BoolPtrInput
	// Either the path to or the contents of a Service Account key file in JSON format.
	ServiceAccountKeyFile pulumi.StringPtrInput
	// Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly
	// specified.
	StorageAccessKey pulumi.StringPtrInput
	// Yandex.Cloud storage service endpoint. Default is storage.yandexcloud.net
	StorageEndpoint pulumi.StringPtrInput
	// Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly
	// specified.
	StorageSecretKey pulumi.StringPtrInput
	// The access token for API operations.
	Token pulumi.StringPtrInput
	// Yandex.Cloud Message Queue service access key. Used when a message queue resource doesn't have an access key explicitly
	// specified.
	YmqAccessKey pulumi.StringPtrInput
	// Yandex.Cloud Message Queue service endpoint. Default is message-queue.api.cloud.yandex.net
	YmqEndpoint pulumi.StringPtrInput
	// Yandex.Cloud Message Queue service secret key. Used when a message queue resource doesn't have a secret key explicitly
	// specified.
	YmqSecretKey pulumi.StringPtrInput
	// The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
	Zone pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *Provider) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderPtrInput interface {
	pulumi.Input

	ToProviderPtrOutput() ProviderPtrOutput
	ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput
}

type providerPtrType ProviderArgs

func (*providerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (i *providerPtrType) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *providerPtrType) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o.ToProviderPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Provider) *Provider {
		return &v
	}).(ProviderPtrOutput)
}

type ProviderPtrOutput struct{ *pulumi.OutputState }

func (ProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (o ProviderPtrOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) Elem() ProviderOutput {
	return o.ApplyT(func(v *Provider) Provider {
		if v != nil {
			return *v
		}
		var ret Provider
		return ret
	}).(ProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderPtrInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderPtrOutput{})
}
