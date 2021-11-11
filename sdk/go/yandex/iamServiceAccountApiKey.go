// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of a [Yandex.Cloud IAM service account API key](https://cloud.yandex.com/docs/iam/concepts/authorization/api-key).
// The API key is a private key used for simplified authorization in the Yandex.Cloud API. API keys are only used for [service accounts](https://cloud.yandex.com/docs/iam/concepts/users/service-accounts).
//
// API keys do not expire. This means that this authentication method is simpler, but less secure. Use it if you can't automatically request an [IAM token](https://cloud.yandex.com/docs/iam/concepts/authorization/iam-token).
//
// ## Example Usage
//
// This snippet creates an API key.
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
// 		_, err := yandex.NewIamServiceAccountApiKey(ctx, "sa_api_key", &yandex.IamServiceAccountApiKeyArgs{
// 			Description:      pulumi.String("api key for authorization"),
// 			PgpKey:           pulumi.String("keybase:keybaseusername"),
// 			ServiceAccountId: pulumi.String("some_sa_id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IamServiceAccountApiKey struct {
	pulumi.CustomResourceState

	// Creation timestamp of the static access key.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the key.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The encrypted secret key, base64 encoded. This is only populated when `pgpKey` is supplied.
	EncryptedSecretKey pulumi.StringOutput `pulumi:"encryptedSecretKey"`
	// The fingerprint of the PGP key used to encrypt the secret key. This is only populated when `pgpKey` is supplied.
	KeyFingerprint pulumi.StringOutput `pulumi:"keyFingerprint"`
	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey pulumi.StringPtrOutput `pulumi:"pgpKey"`
	// The secret key. This is only populated when no `pgpKey` is provided.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// ID of the service account to an API key for.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
}

// NewIamServiceAccountApiKey registers a new resource with the given unique name, arguments, and options.
func NewIamServiceAccountApiKey(ctx *pulumi.Context,
	name string, args *IamServiceAccountApiKeyArgs, opts ...pulumi.ResourceOption) (*IamServiceAccountApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	var resource IamServiceAccountApiKey
	err := ctx.RegisterResource("yandex:index/iamServiceAccountApiKey:IamServiceAccountApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamServiceAccountApiKey gets an existing IamServiceAccountApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamServiceAccountApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamServiceAccountApiKeyState, opts ...pulumi.ResourceOption) (*IamServiceAccountApiKey, error) {
	var resource IamServiceAccountApiKey
	err := ctx.ReadResource("yandex:index/iamServiceAccountApiKey:IamServiceAccountApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamServiceAccountApiKey resources.
type iamServiceAccountApiKeyState struct {
	// Creation timestamp of the static access key.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the key.
	Description *string `pulumi:"description"`
	// The encrypted secret key, base64 encoded. This is only populated when `pgpKey` is supplied.
	EncryptedSecretKey *string `pulumi:"encryptedSecretKey"`
	// The fingerprint of the PGP key used to encrypt the secret key. This is only populated when `pgpKey` is supplied.
	KeyFingerprint *string `pulumi:"keyFingerprint"`
	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey *string `pulumi:"pgpKey"`
	// The secret key. This is only populated when no `pgpKey` is provided.
	SecretKey *string `pulumi:"secretKey"`
	// ID of the service account to an API key for.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

type IamServiceAccountApiKeyState struct {
	// Creation timestamp of the static access key.
	CreatedAt pulumi.StringPtrInput
	// The description of the key.
	Description pulumi.StringPtrInput
	// The encrypted secret key, base64 encoded. This is only populated when `pgpKey` is supplied.
	EncryptedSecretKey pulumi.StringPtrInput
	// The fingerprint of the PGP key used to encrypt the secret key. This is only populated when `pgpKey` is supplied.
	KeyFingerprint pulumi.StringPtrInput
	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey pulumi.StringPtrInput
	// The secret key. This is only populated when no `pgpKey` is provided.
	SecretKey pulumi.StringPtrInput
	// ID of the service account to an API key for.
	ServiceAccountId pulumi.StringPtrInput
}

func (IamServiceAccountApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamServiceAccountApiKeyState)(nil)).Elem()
}

type iamServiceAccountApiKeyArgs struct {
	// The description of the key.
	Description *string `pulumi:"description"`
	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey *string `pulumi:"pgpKey"`
	// ID of the service account to an API key for.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a IamServiceAccountApiKey resource.
type IamServiceAccountApiKeyArgs struct {
	// The description of the key.
	Description pulumi.StringPtrInput
	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey pulumi.StringPtrInput
	// ID of the service account to an API key for.
	ServiceAccountId pulumi.StringInput
}

func (IamServiceAccountApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamServiceAccountApiKeyArgs)(nil)).Elem()
}

type IamServiceAccountApiKeyInput interface {
	pulumi.Input

	ToIamServiceAccountApiKeyOutput() IamServiceAccountApiKeyOutput
	ToIamServiceAccountApiKeyOutputWithContext(ctx context.Context) IamServiceAccountApiKeyOutput
}

func (*IamServiceAccountApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((*IamServiceAccountApiKey)(nil))
}

func (i *IamServiceAccountApiKey) ToIamServiceAccountApiKeyOutput() IamServiceAccountApiKeyOutput {
	return i.ToIamServiceAccountApiKeyOutputWithContext(context.Background())
}

func (i *IamServiceAccountApiKey) ToIamServiceAccountApiKeyOutputWithContext(ctx context.Context) IamServiceAccountApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountApiKeyOutput)
}

func (i *IamServiceAccountApiKey) ToIamServiceAccountApiKeyPtrOutput() IamServiceAccountApiKeyPtrOutput {
	return i.ToIamServiceAccountApiKeyPtrOutputWithContext(context.Background())
}

func (i *IamServiceAccountApiKey) ToIamServiceAccountApiKeyPtrOutputWithContext(ctx context.Context) IamServiceAccountApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountApiKeyPtrOutput)
}

type IamServiceAccountApiKeyPtrInput interface {
	pulumi.Input

	ToIamServiceAccountApiKeyPtrOutput() IamServiceAccountApiKeyPtrOutput
	ToIamServiceAccountApiKeyPtrOutputWithContext(ctx context.Context) IamServiceAccountApiKeyPtrOutput
}

type iamServiceAccountApiKeyPtrType IamServiceAccountApiKeyArgs

func (*iamServiceAccountApiKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IamServiceAccountApiKey)(nil))
}

func (i *iamServiceAccountApiKeyPtrType) ToIamServiceAccountApiKeyPtrOutput() IamServiceAccountApiKeyPtrOutput {
	return i.ToIamServiceAccountApiKeyPtrOutputWithContext(context.Background())
}

func (i *iamServiceAccountApiKeyPtrType) ToIamServiceAccountApiKeyPtrOutputWithContext(ctx context.Context) IamServiceAccountApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountApiKeyPtrOutput)
}

// IamServiceAccountApiKeyArrayInput is an input type that accepts IamServiceAccountApiKeyArray and IamServiceAccountApiKeyArrayOutput values.
// You can construct a concrete instance of `IamServiceAccountApiKeyArrayInput` via:
//
//          IamServiceAccountApiKeyArray{ IamServiceAccountApiKeyArgs{...} }
type IamServiceAccountApiKeyArrayInput interface {
	pulumi.Input

	ToIamServiceAccountApiKeyArrayOutput() IamServiceAccountApiKeyArrayOutput
	ToIamServiceAccountApiKeyArrayOutputWithContext(context.Context) IamServiceAccountApiKeyArrayOutput
}

type IamServiceAccountApiKeyArray []IamServiceAccountApiKeyInput

func (IamServiceAccountApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamServiceAccountApiKey)(nil)).Elem()
}

func (i IamServiceAccountApiKeyArray) ToIamServiceAccountApiKeyArrayOutput() IamServiceAccountApiKeyArrayOutput {
	return i.ToIamServiceAccountApiKeyArrayOutputWithContext(context.Background())
}

func (i IamServiceAccountApiKeyArray) ToIamServiceAccountApiKeyArrayOutputWithContext(ctx context.Context) IamServiceAccountApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountApiKeyArrayOutput)
}

// IamServiceAccountApiKeyMapInput is an input type that accepts IamServiceAccountApiKeyMap and IamServiceAccountApiKeyMapOutput values.
// You can construct a concrete instance of `IamServiceAccountApiKeyMapInput` via:
//
//          IamServiceAccountApiKeyMap{ "key": IamServiceAccountApiKeyArgs{...} }
type IamServiceAccountApiKeyMapInput interface {
	pulumi.Input

	ToIamServiceAccountApiKeyMapOutput() IamServiceAccountApiKeyMapOutput
	ToIamServiceAccountApiKeyMapOutputWithContext(context.Context) IamServiceAccountApiKeyMapOutput
}

type IamServiceAccountApiKeyMap map[string]IamServiceAccountApiKeyInput

func (IamServiceAccountApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamServiceAccountApiKey)(nil)).Elem()
}

func (i IamServiceAccountApiKeyMap) ToIamServiceAccountApiKeyMapOutput() IamServiceAccountApiKeyMapOutput {
	return i.ToIamServiceAccountApiKeyMapOutputWithContext(context.Background())
}

func (i IamServiceAccountApiKeyMap) ToIamServiceAccountApiKeyMapOutputWithContext(ctx context.Context) IamServiceAccountApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountApiKeyMapOutput)
}

type IamServiceAccountApiKeyOutput struct{ *pulumi.OutputState }

func (IamServiceAccountApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamServiceAccountApiKey)(nil))
}

func (o IamServiceAccountApiKeyOutput) ToIamServiceAccountApiKeyOutput() IamServiceAccountApiKeyOutput {
	return o
}

func (o IamServiceAccountApiKeyOutput) ToIamServiceAccountApiKeyOutputWithContext(ctx context.Context) IamServiceAccountApiKeyOutput {
	return o
}

func (o IamServiceAccountApiKeyOutput) ToIamServiceAccountApiKeyPtrOutput() IamServiceAccountApiKeyPtrOutput {
	return o.ToIamServiceAccountApiKeyPtrOutputWithContext(context.Background())
}

func (o IamServiceAccountApiKeyOutput) ToIamServiceAccountApiKeyPtrOutputWithContext(ctx context.Context) IamServiceAccountApiKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IamServiceAccountApiKey) *IamServiceAccountApiKey {
		return &v
	}).(IamServiceAccountApiKeyPtrOutput)
}

type IamServiceAccountApiKeyPtrOutput struct{ *pulumi.OutputState }

func (IamServiceAccountApiKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamServiceAccountApiKey)(nil))
}

func (o IamServiceAccountApiKeyPtrOutput) ToIamServiceAccountApiKeyPtrOutput() IamServiceAccountApiKeyPtrOutput {
	return o
}

func (o IamServiceAccountApiKeyPtrOutput) ToIamServiceAccountApiKeyPtrOutputWithContext(ctx context.Context) IamServiceAccountApiKeyPtrOutput {
	return o
}

func (o IamServiceAccountApiKeyPtrOutput) Elem() IamServiceAccountApiKeyOutput {
	return o.ApplyT(func(v *IamServiceAccountApiKey) IamServiceAccountApiKey {
		if v != nil {
			return *v
		}
		var ret IamServiceAccountApiKey
		return ret
	}).(IamServiceAccountApiKeyOutput)
}

type IamServiceAccountApiKeyArrayOutput struct{ *pulumi.OutputState }

func (IamServiceAccountApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamServiceAccountApiKey)(nil))
}

func (o IamServiceAccountApiKeyArrayOutput) ToIamServiceAccountApiKeyArrayOutput() IamServiceAccountApiKeyArrayOutput {
	return o
}

func (o IamServiceAccountApiKeyArrayOutput) ToIamServiceAccountApiKeyArrayOutputWithContext(ctx context.Context) IamServiceAccountApiKeyArrayOutput {
	return o
}

func (o IamServiceAccountApiKeyArrayOutput) Index(i pulumi.IntInput) IamServiceAccountApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IamServiceAccountApiKey {
		return vs[0].([]IamServiceAccountApiKey)[vs[1].(int)]
	}).(IamServiceAccountApiKeyOutput)
}

type IamServiceAccountApiKeyMapOutput struct{ *pulumi.OutputState }

func (IamServiceAccountApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamServiceAccountApiKey)(nil))
}

func (o IamServiceAccountApiKeyMapOutput) ToIamServiceAccountApiKeyMapOutput() IamServiceAccountApiKeyMapOutput {
	return o
}

func (o IamServiceAccountApiKeyMapOutput) ToIamServiceAccountApiKeyMapOutputWithContext(ctx context.Context) IamServiceAccountApiKeyMapOutput {
	return o
}

func (o IamServiceAccountApiKeyMapOutput) MapIndex(k pulumi.StringInput) IamServiceAccountApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IamServiceAccountApiKey {
		return vs[0].(map[string]IamServiceAccountApiKey)[vs[1].(string)]
	}).(IamServiceAccountApiKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamServiceAccountApiKeyInput)(nil)).Elem(), &IamServiceAccountApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamServiceAccountApiKeyPtrInput)(nil)).Elem(), &IamServiceAccountApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamServiceAccountApiKeyArrayInput)(nil)).Elem(), IamServiceAccountApiKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamServiceAccountApiKeyMapInput)(nil)).Elem(), IamServiceAccountApiKeyMap{})
	pulumi.RegisterOutputType(IamServiceAccountApiKeyOutput{})
	pulumi.RegisterOutputType(IamServiceAccountApiKeyPtrOutput{})
	pulumi.RegisterOutputType(IamServiceAccountApiKeyArrayOutput{})
	pulumi.RegisterOutputType(IamServiceAccountApiKeyMapOutput{})
}
