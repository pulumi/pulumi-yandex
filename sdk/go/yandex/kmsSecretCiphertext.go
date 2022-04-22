// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Encrypts given plaintext with the specified Yandex KMS key and provides access to the ciphertext.
//
// > **Note:** Using this resource will allow you to conceal secret data within your
// resource definitions, but it does not take care of protecting that data in the
// logging output, plan output, or state output.  Please take care to secure your secret
// data outside of resource definitions.
//
// For more information, see [the official documentation](https://cloud.yandex.com/docs/kms/concepts/).
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
// 		example, err := yandex.NewKmsSymmetricKey(ctx, "example", &yandex.KmsSymmetricKeyArgs{
// 			Description: pulumi.String("description for key"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewKmsSecretCiphertext(ctx, "password", &yandex.KmsSecretCiphertextArgs{
// 			AadContext: pulumi.String("additional authenticated data"),
// 			KeyId:      example.ID(),
// 			Plaintext:  pulumi.String("strong password"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type KmsSecretCiphertext struct {
	pulumi.CustomResourceState

	// Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the `SymmetricDecryptRequest`
	AadContext pulumi.StringPtrOutput `pulumi:"aadContext"`
	// Resulting ciphertext, encoded with "standard" base64 alphabet as defined in RFC 4648 section 4
	Ciphertext pulumi.StringOutput `pulumi:"ciphertext"`
	// ID of the symmetric KMS key to use for encryption.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// Plaintext to be encrypted.
	Plaintext pulumi.StringOutput `pulumi:"plaintext"`
}

// NewKmsSecretCiphertext registers a new resource with the given unique name, arguments, and options.
func NewKmsSecretCiphertext(ctx *pulumi.Context,
	name string, args *KmsSecretCiphertextArgs, opts ...pulumi.ResourceOption) (*KmsSecretCiphertext, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	if args.Plaintext == nil {
		return nil, errors.New("invalid value for required argument 'Plaintext'")
	}
	var resource KmsSecretCiphertext
	err := ctx.RegisterResource("yandex:index/kmsSecretCiphertext:KmsSecretCiphertext", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsSecretCiphertext gets an existing KmsSecretCiphertext resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsSecretCiphertext(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsSecretCiphertextState, opts ...pulumi.ResourceOption) (*KmsSecretCiphertext, error) {
	var resource KmsSecretCiphertext
	err := ctx.ReadResource("yandex:index/kmsSecretCiphertext:KmsSecretCiphertext", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsSecretCiphertext resources.
type kmsSecretCiphertextState struct {
	// Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the `SymmetricDecryptRequest`
	AadContext *string `pulumi:"aadContext"`
	// Resulting ciphertext, encoded with "standard" base64 alphabet as defined in RFC 4648 section 4
	Ciphertext *string `pulumi:"ciphertext"`
	// ID of the symmetric KMS key to use for encryption.
	KeyId *string `pulumi:"keyId"`
	// Plaintext to be encrypted.
	Plaintext *string `pulumi:"plaintext"`
}

type KmsSecretCiphertextState struct {
	// Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the `SymmetricDecryptRequest`
	AadContext pulumi.StringPtrInput
	// Resulting ciphertext, encoded with "standard" base64 alphabet as defined in RFC 4648 section 4
	Ciphertext pulumi.StringPtrInput
	// ID of the symmetric KMS key to use for encryption.
	KeyId pulumi.StringPtrInput
	// Plaintext to be encrypted.
	Plaintext pulumi.StringPtrInput
}

func (KmsSecretCiphertextState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsSecretCiphertextState)(nil)).Elem()
}

type kmsSecretCiphertextArgs struct {
	// Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the `SymmetricDecryptRequest`
	AadContext *string `pulumi:"aadContext"`
	// ID of the symmetric KMS key to use for encryption.
	KeyId string `pulumi:"keyId"`
	// Plaintext to be encrypted.
	Plaintext string `pulumi:"plaintext"`
}

// The set of arguments for constructing a KmsSecretCiphertext resource.
type KmsSecretCiphertextArgs struct {
	// Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the `SymmetricDecryptRequest`
	AadContext pulumi.StringPtrInput
	// ID of the symmetric KMS key to use for encryption.
	KeyId pulumi.StringInput
	// Plaintext to be encrypted.
	Plaintext pulumi.StringInput
}

func (KmsSecretCiphertextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsSecretCiphertextArgs)(nil)).Elem()
}

type KmsSecretCiphertextInput interface {
	pulumi.Input

	ToKmsSecretCiphertextOutput() KmsSecretCiphertextOutput
	ToKmsSecretCiphertextOutputWithContext(ctx context.Context) KmsSecretCiphertextOutput
}

func (*KmsSecretCiphertext) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsSecretCiphertext)(nil)).Elem()
}

func (i *KmsSecretCiphertext) ToKmsSecretCiphertextOutput() KmsSecretCiphertextOutput {
	return i.ToKmsSecretCiphertextOutputWithContext(context.Background())
}

func (i *KmsSecretCiphertext) ToKmsSecretCiphertextOutputWithContext(ctx context.Context) KmsSecretCiphertextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSecretCiphertextOutput)
}

// KmsSecretCiphertextArrayInput is an input type that accepts KmsSecretCiphertextArray and KmsSecretCiphertextArrayOutput values.
// You can construct a concrete instance of `KmsSecretCiphertextArrayInput` via:
//
//          KmsSecretCiphertextArray{ KmsSecretCiphertextArgs{...} }
type KmsSecretCiphertextArrayInput interface {
	pulumi.Input

	ToKmsSecretCiphertextArrayOutput() KmsSecretCiphertextArrayOutput
	ToKmsSecretCiphertextArrayOutputWithContext(context.Context) KmsSecretCiphertextArrayOutput
}

type KmsSecretCiphertextArray []KmsSecretCiphertextInput

func (KmsSecretCiphertextArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsSecretCiphertext)(nil)).Elem()
}

func (i KmsSecretCiphertextArray) ToKmsSecretCiphertextArrayOutput() KmsSecretCiphertextArrayOutput {
	return i.ToKmsSecretCiphertextArrayOutputWithContext(context.Background())
}

func (i KmsSecretCiphertextArray) ToKmsSecretCiphertextArrayOutputWithContext(ctx context.Context) KmsSecretCiphertextArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSecretCiphertextArrayOutput)
}

// KmsSecretCiphertextMapInput is an input type that accepts KmsSecretCiphertextMap and KmsSecretCiphertextMapOutput values.
// You can construct a concrete instance of `KmsSecretCiphertextMapInput` via:
//
//          KmsSecretCiphertextMap{ "key": KmsSecretCiphertextArgs{...} }
type KmsSecretCiphertextMapInput interface {
	pulumi.Input

	ToKmsSecretCiphertextMapOutput() KmsSecretCiphertextMapOutput
	ToKmsSecretCiphertextMapOutputWithContext(context.Context) KmsSecretCiphertextMapOutput
}

type KmsSecretCiphertextMap map[string]KmsSecretCiphertextInput

func (KmsSecretCiphertextMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsSecretCiphertext)(nil)).Elem()
}

func (i KmsSecretCiphertextMap) ToKmsSecretCiphertextMapOutput() KmsSecretCiphertextMapOutput {
	return i.ToKmsSecretCiphertextMapOutputWithContext(context.Background())
}

func (i KmsSecretCiphertextMap) ToKmsSecretCiphertextMapOutputWithContext(ctx context.Context) KmsSecretCiphertextMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSecretCiphertextMapOutput)
}

type KmsSecretCiphertextOutput struct{ *pulumi.OutputState }

func (KmsSecretCiphertextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsSecretCiphertext)(nil)).Elem()
}

func (o KmsSecretCiphertextOutput) ToKmsSecretCiphertextOutput() KmsSecretCiphertextOutput {
	return o
}

func (o KmsSecretCiphertextOutput) ToKmsSecretCiphertextOutputWithContext(ctx context.Context) KmsSecretCiphertextOutput {
	return o
}

type KmsSecretCiphertextArrayOutput struct{ *pulumi.OutputState }

func (KmsSecretCiphertextArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsSecretCiphertext)(nil)).Elem()
}

func (o KmsSecretCiphertextArrayOutput) ToKmsSecretCiphertextArrayOutput() KmsSecretCiphertextArrayOutput {
	return o
}

func (o KmsSecretCiphertextArrayOutput) ToKmsSecretCiphertextArrayOutputWithContext(ctx context.Context) KmsSecretCiphertextArrayOutput {
	return o
}

func (o KmsSecretCiphertextArrayOutput) Index(i pulumi.IntInput) KmsSecretCiphertextOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KmsSecretCiphertext {
		return vs[0].([]*KmsSecretCiphertext)[vs[1].(int)]
	}).(KmsSecretCiphertextOutput)
}

type KmsSecretCiphertextMapOutput struct{ *pulumi.OutputState }

func (KmsSecretCiphertextMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsSecretCiphertext)(nil)).Elem()
}

func (o KmsSecretCiphertextMapOutput) ToKmsSecretCiphertextMapOutput() KmsSecretCiphertextMapOutput {
	return o
}

func (o KmsSecretCiphertextMapOutput) ToKmsSecretCiphertextMapOutputWithContext(ctx context.Context) KmsSecretCiphertextMapOutput {
	return o
}

func (o KmsSecretCiphertextMapOutput) MapIndex(k pulumi.StringInput) KmsSecretCiphertextOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KmsSecretCiphertext {
		return vs[0].(map[string]*KmsSecretCiphertext)[vs[1].(string)]
	}).(KmsSecretCiphertextOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSecretCiphertextInput)(nil)).Elem(), &KmsSecretCiphertext{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSecretCiphertextArrayInput)(nil)).Elem(), KmsSecretCiphertextArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSecretCiphertextMapInput)(nil)).Elem(), KmsSecretCiphertextMap{})
	pulumi.RegisterOutputType(KmsSecretCiphertextOutput{})
	pulumi.RegisterOutputType(KmsSecretCiphertextArrayOutput{})
	pulumi.RegisterOutputType(KmsSecretCiphertextMapOutput{})
}
