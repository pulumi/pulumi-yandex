// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of [Yandex.Cloud IAM service account authorized keys](https://cloud.yandex.com/docs/iam/concepts/authorization/key).
// Generated pair of keys is used to create a [JSON Web Token](https://tools.ietf.org/html/rfc7519) which is necessary for requesting an [IAM Token](https://cloud.yandex.com/docs/iam/concepts/authorization/iam-token) for a [service account](https://cloud.yandex.com/docs/iam/concepts/users/service-accounts).
//
// ## Example Usage
//
// This snippet creates an authorized keys pair.
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
// 		_, err := yandex.NewIamServiceAccountKey(ctx, "sa-auth-key", &yandex.IamServiceAccountKeyArgs{
// 			Description:      pulumi.String("key for service account"),
// 			KeyAlgorithm:     pulumi.String("RSA_4096"),
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
type IamServiceAccountKey struct {
	pulumi.CustomResourceState

	// Creation timestamp of the static access key.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the key pair.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The encrypted private key, base64 encoded. This is only populated when `pgpKey` is supplied.
	EncryptedPrivateKey pulumi.StringOutput `pulumi:"encryptedPrivateKey"`
	// The output format of the keys. `PEM_FILE` is the default format.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The algorithm used to generate the key. `RSA_2048` is the default algorithm.
	// Valid values are listed in the [API reference](https://cloud.yandex.com/docs/iam/api-ref/Key).
	KeyAlgorithm pulumi.StringPtrOutput `pulumi:"keyAlgorithm"`
	// The fingerprint of the PGP key used to encrypt the private key. This is only populated when `pgpKey` is supplied.
	KeyFingerprint pulumi.StringOutput `pulumi:"keyFingerprint"`
	// An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey pulumi.StringPtrOutput `pulumi:"pgpKey"`
	// The private key. This is only populated when no `pgpKey` is provided.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The public key.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// ID of the service account to create a pair for.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
}

// NewIamServiceAccountKey registers a new resource with the given unique name, arguments, and options.
func NewIamServiceAccountKey(ctx *pulumi.Context,
	name string, args *IamServiceAccountKeyArgs, opts ...pulumi.ResourceOption) (*IamServiceAccountKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	var resource IamServiceAccountKey
	err := ctx.RegisterResource("yandex:index/iamServiceAccountKey:IamServiceAccountKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamServiceAccountKey gets an existing IamServiceAccountKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamServiceAccountKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamServiceAccountKeyState, opts ...pulumi.ResourceOption) (*IamServiceAccountKey, error) {
	var resource IamServiceAccountKey
	err := ctx.ReadResource("yandex:index/iamServiceAccountKey:IamServiceAccountKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamServiceAccountKey resources.
type iamServiceAccountKeyState struct {
	// Creation timestamp of the static access key.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the key pair.
	Description *string `pulumi:"description"`
	// The encrypted private key, base64 encoded. This is only populated when `pgpKey` is supplied.
	EncryptedPrivateKey *string `pulumi:"encryptedPrivateKey"`
	// The output format of the keys. `PEM_FILE` is the default format.
	Format *string `pulumi:"format"`
	// The algorithm used to generate the key. `RSA_2048` is the default algorithm.
	// Valid values are listed in the [API reference](https://cloud.yandex.com/docs/iam/api-ref/Key).
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// The fingerprint of the PGP key used to encrypt the private key. This is only populated when `pgpKey` is supplied.
	KeyFingerprint *string `pulumi:"keyFingerprint"`
	// An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey *string `pulumi:"pgpKey"`
	// The private key. This is only populated when no `pgpKey` is provided.
	PrivateKey *string `pulumi:"privateKey"`
	// The public key.
	PublicKey *string `pulumi:"publicKey"`
	// ID of the service account to create a pair for.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

type IamServiceAccountKeyState struct {
	// Creation timestamp of the static access key.
	CreatedAt pulumi.StringPtrInput
	// The description of the key pair.
	Description pulumi.StringPtrInput
	// The encrypted private key, base64 encoded. This is only populated when `pgpKey` is supplied.
	EncryptedPrivateKey pulumi.StringPtrInput
	// The output format of the keys. `PEM_FILE` is the default format.
	Format pulumi.StringPtrInput
	// The algorithm used to generate the key. `RSA_2048` is the default algorithm.
	// Valid values are listed in the [API reference](https://cloud.yandex.com/docs/iam/api-ref/Key).
	KeyAlgorithm pulumi.StringPtrInput
	// The fingerprint of the PGP key used to encrypt the private key. This is only populated when `pgpKey` is supplied.
	KeyFingerprint pulumi.StringPtrInput
	// An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey pulumi.StringPtrInput
	// The private key. This is only populated when no `pgpKey` is provided.
	PrivateKey pulumi.StringPtrInput
	// The public key.
	PublicKey pulumi.StringPtrInput
	// ID of the service account to create a pair for.
	ServiceAccountId pulumi.StringPtrInput
}

func (IamServiceAccountKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamServiceAccountKeyState)(nil)).Elem()
}

type iamServiceAccountKeyArgs struct {
	// The description of the key pair.
	Description *string `pulumi:"description"`
	// The output format of the keys. `PEM_FILE` is the default format.
	Format *string `pulumi:"format"`
	// The algorithm used to generate the key. `RSA_2048` is the default algorithm.
	// Valid values are listed in the [API reference](https://cloud.yandex.com/docs/iam/api-ref/Key).
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey *string `pulumi:"pgpKey"`
	// ID of the service account to create a pair for.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a IamServiceAccountKey resource.
type IamServiceAccountKeyArgs struct {
	// The description of the key pair.
	Description pulumi.StringPtrInput
	// The output format of the keys. `PEM_FILE` is the default format.
	Format pulumi.StringPtrInput
	// The algorithm used to generate the key. `RSA_2048` is the default algorithm.
	// Valid values are listed in the [API reference](https://cloud.yandex.com/docs/iam/api-ref/Key).
	KeyAlgorithm pulumi.StringPtrInput
	// An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form `keybase:keybaseusername`.
	PgpKey pulumi.StringPtrInput
	// ID of the service account to create a pair for.
	ServiceAccountId pulumi.StringInput
}

func (IamServiceAccountKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamServiceAccountKeyArgs)(nil)).Elem()
}

type IamServiceAccountKeyInput interface {
	pulumi.Input

	ToIamServiceAccountKeyOutput() IamServiceAccountKeyOutput
	ToIamServiceAccountKeyOutputWithContext(ctx context.Context) IamServiceAccountKeyOutput
}

func (*IamServiceAccountKey) ElementType() reflect.Type {
	return reflect.TypeOf((**IamServiceAccountKey)(nil)).Elem()
}

func (i *IamServiceAccountKey) ToIamServiceAccountKeyOutput() IamServiceAccountKeyOutput {
	return i.ToIamServiceAccountKeyOutputWithContext(context.Background())
}

func (i *IamServiceAccountKey) ToIamServiceAccountKeyOutputWithContext(ctx context.Context) IamServiceAccountKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountKeyOutput)
}

// IamServiceAccountKeyArrayInput is an input type that accepts IamServiceAccountKeyArray and IamServiceAccountKeyArrayOutput values.
// You can construct a concrete instance of `IamServiceAccountKeyArrayInput` via:
//
//          IamServiceAccountKeyArray{ IamServiceAccountKeyArgs{...} }
type IamServiceAccountKeyArrayInput interface {
	pulumi.Input

	ToIamServiceAccountKeyArrayOutput() IamServiceAccountKeyArrayOutput
	ToIamServiceAccountKeyArrayOutputWithContext(context.Context) IamServiceAccountKeyArrayOutput
}

type IamServiceAccountKeyArray []IamServiceAccountKeyInput

func (IamServiceAccountKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamServiceAccountKey)(nil)).Elem()
}

func (i IamServiceAccountKeyArray) ToIamServiceAccountKeyArrayOutput() IamServiceAccountKeyArrayOutput {
	return i.ToIamServiceAccountKeyArrayOutputWithContext(context.Background())
}

func (i IamServiceAccountKeyArray) ToIamServiceAccountKeyArrayOutputWithContext(ctx context.Context) IamServiceAccountKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountKeyArrayOutput)
}

// IamServiceAccountKeyMapInput is an input type that accepts IamServiceAccountKeyMap and IamServiceAccountKeyMapOutput values.
// You can construct a concrete instance of `IamServiceAccountKeyMapInput` via:
//
//          IamServiceAccountKeyMap{ "key": IamServiceAccountKeyArgs{...} }
type IamServiceAccountKeyMapInput interface {
	pulumi.Input

	ToIamServiceAccountKeyMapOutput() IamServiceAccountKeyMapOutput
	ToIamServiceAccountKeyMapOutputWithContext(context.Context) IamServiceAccountKeyMapOutput
}

type IamServiceAccountKeyMap map[string]IamServiceAccountKeyInput

func (IamServiceAccountKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamServiceAccountKey)(nil)).Elem()
}

func (i IamServiceAccountKeyMap) ToIamServiceAccountKeyMapOutput() IamServiceAccountKeyMapOutput {
	return i.ToIamServiceAccountKeyMapOutputWithContext(context.Background())
}

func (i IamServiceAccountKeyMap) ToIamServiceAccountKeyMapOutputWithContext(ctx context.Context) IamServiceAccountKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountKeyMapOutput)
}

type IamServiceAccountKeyOutput struct{ *pulumi.OutputState }

func (IamServiceAccountKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamServiceAccountKey)(nil)).Elem()
}

func (o IamServiceAccountKeyOutput) ToIamServiceAccountKeyOutput() IamServiceAccountKeyOutput {
	return o
}

func (o IamServiceAccountKeyOutput) ToIamServiceAccountKeyOutputWithContext(ctx context.Context) IamServiceAccountKeyOutput {
	return o
}

type IamServiceAccountKeyArrayOutput struct{ *pulumi.OutputState }

func (IamServiceAccountKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamServiceAccountKey)(nil)).Elem()
}

func (o IamServiceAccountKeyArrayOutput) ToIamServiceAccountKeyArrayOutput() IamServiceAccountKeyArrayOutput {
	return o
}

func (o IamServiceAccountKeyArrayOutput) ToIamServiceAccountKeyArrayOutputWithContext(ctx context.Context) IamServiceAccountKeyArrayOutput {
	return o
}

func (o IamServiceAccountKeyArrayOutput) Index(i pulumi.IntInput) IamServiceAccountKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamServiceAccountKey {
		return vs[0].([]*IamServiceAccountKey)[vs[1].(int)]
	}).(IamServiceAccountKeyOutput)
}

type IamServiceAccountKeyMapOutput struct{ *pulumi.OutputState }

func (IamServiceAccountKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamServiceAccountKey)(nil)).Elem()
}

func (o IamServiceAccountKeyMapOutput) ToIamServiceAccountKeyMapOutput() IamServiceAccountKeyMapOutput {
	return o
}

func (o IamServiceAccountKeyMapOutput) ToIamServiceAccountKeyMapOutputWithContext(ctx context.Context) IamServiceAccountKeyMapOutput {
	return o
}

func (o IamServiceAccountKeyMapOutput) MapIndex(k pulumi.StringInput) IamServiceAccountKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamServiceAccountKey {
		return vs[0].(map[string]*IamServiceAccountKey)[vs[1].(string)]
	}).(IamServiceAccountKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamServiceAccountKeyInput)(nil)).Elem(), &IamServiceAccountKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamServiceAccountKeyArrayInput)(nil)).Elem(), IamServiceAccountKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamServiceAccountKeyMapInput)(nil)).Elem(), IamServiceAccountKeyMap{})
	pulumi.RegisterOutputType(IamServiceAccountKeyOutput{})
	pulumi.RegisterOutputType(IamServiceAccountKeyArrayOutput{})
	pulumi.RegisterOutputType(IamServiceAccountKeyMapOutput{})
}
