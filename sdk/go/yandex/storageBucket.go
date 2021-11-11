// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of [Yandex.Cloud Storage Bucket](https://cloud.yandex.com/docs/storage/concepts/bucket).
//
// > **Note:** Your need to provide [static access key](https://cloud.yandex.com/docs/iam/concepts/authorization/access-key) (Access and Secret) to create storage client to work with Storage Service. To create them you need Service Account and proper permissions.
//
// ## Example Usage
// ### Simple Private Bucket
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
// 		folderId := "<folder-id>"
// 		sa, err := yandex.NewIamServiceAccount(ctx, "sa", &yandex.IamServiceAccountArgs{
// 			FolderId: pulumi.String(folderId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewResourcemanagerFolderIamMember(ctx, "sa_editor", &yandex.ResourcemanagerFolderIamMemberArgs{
// 			FolderId: pulumi.String(folderId),
// 			Role:     pulumi.String("storage.editor"),
// 			Member: sa.ID().ApplyT(func(id string) (string, error) {
// 				return fmt.Sprintf("%v%v", "serviceAccount:", id), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewIamServiceAccountStaticAccessKey(ctx, "sa_static_key", &yandex.IamServiceAccountStaticAccessKeyArgs{
// 			ServiceAccountId: sa.ID(),
// 			Description:      pulumi.String("static access key for object storage"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewStorageBucket(ctx, "test", &yandex.StorageBucketArgs{
// 			AccessKey: sa_static_key.AccessKey,
// 			SecretKey: sa_static_key.SecretKey,
// 			Bucket:    pulumi.String("tf-test-bucket"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Static Website Hosting
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
// 		_, err := yandex.NewStorageBucket(ctx, "test", &yandex.StorageBucketArgs{
// 			Acl:    pulumi.String("public-read"),
// 			Bucket: pulumi.String("storage-website-test.hashicorp.com"),
// 			Website: &StorageBucketWebsiteArgs{
// 				ErrorDocument: pulumi.String("error.html"),
// 				IndexDocument: pulumi.String("index.html"),
// 				RoutingRules:  pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v", "[{\n", "    \"Condition\": {\n", "        \"KeyPrefixEquals\": \"docs/\"\n", "    },\n", "    \"Redirect\": {\n", "        \"ReplaceKeyPrefixWith\": \"documents/\"\n", "    }\n", "}]\n", "\n")),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using ACL policy grants
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
// 		_, err := yandex.NewStorageBucket(ctx, "test", &yandex.StorageBucketArgs{
// 			Bucket: pulumi.String("mybucket"),
// 			Grants: StorageBucketGrantArray{
// 				&StorageBucketGrantArgs{
// 					Id: pulumi.String("myuser"),
// 					Permissions: pulumi.StringArray{
// 						pulumi.String("FULL_CONTROL"),
// 					},
// 					Type: pulumi.String("CanonicalUser"),
// 				},
// 				&StorageBucketGrantArgs{
// 					Permissions: pulumi.StringArray{
// 						pulumi.String("READ"),
// 						pulumi.String("WRITE"),
// 					},
// 					Type: pulumi.String("Group"),
// 					Uri:  pulumi.String("http://acs.amazonaws.com/groups/global/AllUsers"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using CORS
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
// 		_, err := yandex.NewStorageBucket(ctx, "storageBucket", &yandex.StorageBucketArgs{
// 			Acl:    pulumi.String("public-read"),
// 			Bucket: pulumi.String("s3-website-test.hashicorp.com"),
// 			CorsRules: StorageBucketCorsRuleArray{
// 				&StorageBucketCorsRuleArgs{
// 					AllowedHeaders: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 					AllowedMethods: pulumi.StringArray{
// 						pulumi.String("PUT"),
// 						pulumi.String("POST"),
// 					},
// 					AllowedOrigins: pulumi.StringArray{
// 						pulumi.String("https://s3-website-test.hashicorp.com"),
// 					},
// 					ExposeHeaders: pulumi.StringArray{
// 						pulumi.String("ETag"),
// 					},
// 					MaxAgeSeconds: pulumi.Int(3000),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using versioning
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
// 		_, err := yandex.NewStorageBucket(ctx, "storageBucket", &yandex.StorageBucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("my-tf-test-bucket"),
// 			Versioning: &StorageBucketVersioningArgs{
// 				Enabled: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Enable Logging
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
// 		logBucket, err := yandex.NewStorageBucket(ctx, "logBucket", &yandex.StorageBucketArgs{
// 			Bucket: pulumi.String("my-tf-log-bucket"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewStorageBucket(ctx, "storageBucket", &yandex.StorageBucketArgs{
// 			Bucket: pulumi.String("my-tf-test-bucket"),
// 			Acl:    pulumi.String("private"),
// 			Loggings: StorageBucketLoggingArray{
// 				&StorageBucketLoggingArgs{
// 					TargetBucket: logBucket.ID(),
// 					TargetPrefix: pulumi.String("log/"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using object lifecycle
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
// 		_, err := yandex.NewStorageBucket(ctx, "bucket", &yandex.StorageBucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("my-bucket"),
// 			LifecycleRules: StorageBucketLifecycleRuleArray{
// 				&StorageBucketLifecycleRuleArgs{
// 					Enabled: pulumi.Bool(true),
// 					Expiration: &StorageBucketLifecycleRuleExpirationArgs{
// 						Days: pulumi.Int(90),
// 					},
// 					Id:     pulumi.String("log"),
// 					Prefix: pulumi.String("log/"),
// 					Transitions: StorageBucketLifecycleRuleTransitionArray{
// 						&StorageBucketLifecycleRuleTransitionArgs{
// 							Days:         pulumi.Int(30),
// 							StorageClass: pulumi.String("COLD"),
// 						},
// 					},
// 				},
// 				&StorageBucketLifecycleRuleArgs{
// 					Enabled: pulumi.Bool(true),
// 					Expiration: &StorageBucketLifecycleRuleExpirationArgs{
// 						Date: pulumi.String("2020-12-21"),
// 					},
// 					Id:     pulumi.String("tmp"),
// 					Prefix: pulumi.String("tmp/"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewStorageBucket(ctx, "versioningBucket", &yandex.StorageBucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("my-versioning-bucket"),
// 			LifecycleRules: StorageBucketLifecycleRuleArray{
// 				&StorageBucketLifecycleRuleArgs{
// 					Enabled: pulumi.Bool(true),
// 					NoncurrentVersionExpiration: &StorageBucketLifecycleRuleNoncurrentVersionExpirationArgs{
// 						Days: pulumi.Int(90),
// 					},
// 					NoncurrentVersionTransitions: StorageBucketLifecycleRuleNoncurrentVersionTransitionArray{
// 						&StorageBucketLifecycleRuleNoncurrentVersionTransitionArgs{
// 							Days:         pulumi.Int(30),
// 							StorageClass: pulumi.String("COLD"),
// 						},
// 					},
// 					Prefix: pulumi.String("config/"),
// 				},
// 			},
// 			Versioning: &StorageBucketVersioningArgs{
// 				Enabled: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using SSE
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
// 		_, err := yandex.NewKmsSymmetricKey(ctx, "key_a", &yandex.KmsSymmetricKeyArgs{
// 			Description:      pulumi.String("description for key"),
// 			DefaultAlgorithm: pulumi.String("AES_128"),
// 			RotationPeriod:   pulumi.String("8760h"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewStorageBucket(ctx, "test", &yandex.StorageBucketArgs{
// 			Bucket: pulumi.String("mybucket"),
// 			ServerSideEncryptionConfiguration: &StorageBucketServerSideEncryptionConfigurationArgs{
// 				Rule: &StorageBucketServerSideEncryptionConfigurationRuleArgs{
// 					ApplyServerSideEncryptionByDefault: &StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
// 						KmsMasterKeyId: key_a.ID(),
// 						SseAlgorithm:   pulumi.String("aws:kms"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Bucket Policy
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
// 		_, err := yandex.NewStorageBucket(ctx, "storageBucket", &yandex.StorageBucketArgs{
// 			Bucket: pulumi.String("my-policy-bucket"),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"s3:*\",\n", "      \"Resource\": [\n", "        \"arn:aws:s3:::my-policy-bucket/*\",\n", "        \"arn:aws:s3:::my-policy-bucket\"\n", "      ]\n", "    },\n", "    {\n", "      \"Effect\": \"Deny\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"s3:PutObject\",\n", "      \"Resource\": [\n", "        \"arn:aws:s3:::my-policy-bucket/*\",\n", "        \"arn:aws:s3:::my-policy-bucket\"\n", "      ]\n", "    }\n", "  ]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### All settings example
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
// 		logBucket, err := yandex.NewStorageBucket(ctx, "logBucket", &yandex.StorageBucketArgs{
// 			Bucket: pulumi.String("my-tf-log-bucket"),
// 			LifecycleRules: StorageBucketLifecycleRuleArray{
// 				&StorageBucketLifecycleRuleArgs{
// 					Id:      pulumi.String("cleanupoldlogs"),
// 					Enabled: pulumi.Bool(true),
// 					Expiration: &StorageBucketLifecycleRuleExpirationArgs{
// 						Days: pulumi.Int(365),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewKmsSymmetricKey(ctx, "key_a", &yandex.KmsSymmetricKeyArgs{
// 			Description:      pulumi.String("description for key"),
// 			DefaultAlgorithm: pulumi.String("AES_128"),
// 			RotationPeriod:   pulumi.String("8760h"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = yandex.NewStorageBucket(ctx, "allSettings", &yandex.StorageBucketArgs{
// 			Bucket: pulumi.String("example-tf-settings-bucket"),
// 			Website: &StorageBucketWebsiteArgs{
// 				IndexDocument: pulumi.String("index.html"),
// 				ErrorDocument: pulumi.String("error.html"),
// 			},
// 			LifecycleRules: StorageBucketLifecycleRuleArray{
// 				&StorageBucketLifecycleRuleArgs{
// 					Id:      pulumi.String("test"),
// 					Enabled: pulumi.Bool(true),
// 					Prefix:  pulumi.String("prefix/"),
// 					Expiration: &StorageBucketLifecycleRuleExpirationArgs{
// 						Days: pulumi.Int(30),
// 					},
// 				},
// 				&StorageBucketLifecycleRuleArgs{
// 					Id:      pulumi.String("log"),
// 					Enabled: pulumi.Bool(true),
// 					Prefix:  pulumi.String("log/"),
// 					Transitions: StorageBucketLifecycleRuleTransitionArray{
// 						&StorageBucketLifecycleRuleTransitionArgs{
// 							Days:         pulumi.Int(30),
// 							StorageClass: pulumi.String("COLD"),
// 						},
// 					},
// 					Expiration: &StorageBucketLifecycleRuleExpirationArgs{
// 						Days: pulumi.Int(90),
// 					},
// 				},
// 				&StorageBucketLifecycleRuleArgs{
// 					Id:      pulumi.String("everything180"),
// 					Prefix:  pulumi.String(""),
// 					Enabled: pulumi.Bool(true),
// 					Expiration: &StorageBucketLifecycleRuleExpirationArgs{
// 						Days: pulumi.Int(180),
// 					},
// 				},
// 				&StorageBucketLifecycleRuleArgs{
// 					Id:      pulumi.String("cleanupoldversions"),
// 					Prefix:  pulumi.String("config/"),
// 					Enabled: pulumi.Bool(true),
// 					NoncurrentVersionTransitions: StorageBucketLifecycleRuleNoncurrentVersionTransitionArray{
// 						&StorageBucketLifecycleRuleNoncurrentVersionTransitionArgs{
// 							Days:         pulumi.Int(30),
// 							StorageClass: pulumi.String("COLD"),
// 						},
// 					},
// 					NoncurrentVersionExpiration: &StorageBucketLifecycleRuleNoncurrentVersionExpirationArgs{
// 						Days: pulumi.Int(90),
// 					},
// 				},
// 				&StorageBucketLifecycleRuleArgs{
// 					Id:                                 pulumi.String("abortmultiparts"),
// 					Prefix:                             pulumi.String(""),
// 					Enabled:                            pulumi.Bool(true),
// 					AbortIncompleteMultipartUploadDays: pulumi.Int(7),
// 				},
// 			},
// 			CorsRules: StorageBucketCorsRuleArray{
// 				&StorageBucketCorsRuleArgs{
// 					AllowedHeaders: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 					AllowedMethods: pulumi.StringArray{
// 						pulumi.String("GET"),
// 						pulumi.String("PUT"),
// 					},
// 					AllowedOrigins: pulumi.StringArray{
// 						pulumi.String("https://storage-cloud.example.com"),
// 					},
// 					ExposeHeaders: pulumi.StringArray{
// 						pulumi.String("ETag"),
// 					},
// 					MaxAgeSeconds: pulumi.Int(3000),
// 				},
// 			},
// 			Versioning: &StorageBucketVersioningArgs{
// 				Enabled: pulumi.Bool(true),
// 			},
// 			ServerSideEncryptionConfiguration: &StorageBucketServerSideEncryptionConfigurationArgs{
// 				Rule: &StorageBucketServerSideEncryptionConfigurationRuleArgs{
// 					ApplyServerSideEncryptionByDefault: &StorageBucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
// 						KmsMasterKeyId: key_a.ID(),
// 						SseAlgorithm:   pulumi.String("aws:kms"),
// 					},
// 				},
// 			},
// 			Loggings: StorageBucketLoggingArray{
// 				&StorageBucketLoggingArgs{
// 					TargetBucket: logBucket.ID(),
// 					TargetPrefix: pulumi.String("tf-logs/"),
// 				},
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
// Storage bucket can be imported using the `bucket`, e.g.
//
// ```sh
//  $ pulumi import yandex:index/storageBucket:StorageBucket bucket bucket-name
// ```
//
//  `false` in state. If you've set it to `true` in config, run `terraform apply` to update the value set in state. If you delete this resource before updating the value, objects in the bucket will not be destroyed.
type StorageBucket struct {
	pulumi.CustomResourceState

	// The access key to use when applying changes. If omitted, `storageAccessKey` specified in provider config is used.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	Acl    pulumi.StringPtrOutput `pulumi:"acl"`
	Bucket pulumi.StringOutput    `pulumi:"bucket"`
	// The bucket domain name.
	BucketDomainName pulumi.StringOutput `pulumi:"bucketDomainName"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrOutput `pulumi:"bucketPrefix"`
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/cors/) (documented below).
	CorsRules StorageBucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants StorageBucketGrantArrayOutput `pulumi:"grants"`
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules StorageBucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings StorageBucketLoggingArrayOutput `pulumi:"loggings"`
	Policy   pulumi.StringPtrOutput          `pulumi:"policy"`
	// The secret key to use when applying changes. If omitted, `storageSecretKey` specified in provider config is used.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
	// A configuration of server-side encryption for the bucket (documented below)
	ServerSideEncryptionConfiguration StorageBucketServerSideEncryptionConfigurationPtrOutput `pulumi:"serverSideEncryptionConfiguration"`
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	Versioning StorageBucketVersioningOutput `pulumi:"versioning"`
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website StorageBucketWebsitePtrOutput `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain pulumi.StringOutput `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringOutput `pulumi:"websiteEndpoint"`
}

// NewStorageBucket registers a new resource with the given unique name, arguments, and options.
func NewStorageBucket(ctx *pulumi.Context,
	name string, args *StorageBucketArgs, opts ...pulumi.ResourceOption) (*StorageBucket, error) {
	if args == nil {
		args = &StorageBucketArgs{}
	}

	var resource StorageBucket
	err := ctx.RegisterResource("yandex:index/storageBucket:StorageBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBucket gets an existing StorageBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBucketState, opts ...pulumi.ResourceOption) (*StorageBucket, error) {
	var resource StorageBucket
	err := ctx.ReadResource("yandex:index/storageBucket:StorageBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBucket resources.
type storageBucketState struct {
	// The access key to use when applying changes. If omitted, `storageAccessKey` specified in provider config is used.
	AccessKey *string `pulumi:"accessKey"`
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	Acl    *string `pulumi:"acl"`
	Bucket *string `pulumi:"bucket"`
	// The bucket domain name.
	BucketDomainName *string `pulumi:"bucketDomainName"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/cors/) (documented below).
	CorsRules []StorageBucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants []StorageBucketGrant `pulumi:"grants"`
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules []StorageBucketLifecycleRule `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings []StorageBucketLogging `pulumi:"loggings"`
	Policy   *string                `pulumi:"policy"`
	// The secret key to use when applying changes. If omitted, `storageSecretKey` specified in provider config is used.
	SecretKey *string `pulumi:"secretKey"`
	// A configuration of server-side encryption for the bucket (documented below)
	ServerSideEncryptionConfiguration *StorageBucketServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	Versioning *StorageBucketVersioning `pulumi:"versioning"`
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website *StorageBucketWebsite `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

type StorageBucketState struct {
	// The access key to use when applying changes. If omitted, `storageAccessKey` specified in provider config is used.
	AccessKey pulumi.StringPtrInput
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	Acl    pulumi.StringPtrInput
	Bucket pulumi.StringPtrInput
	// The bucket domain name.
	BucketDomainName pulumi.StringPtrInput
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/cors/) (documented below).
	CorsRules StorageBucketCorsRuleArrayInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants StorageBucketGrantArrayInput
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules StorageBucketLifecycleRuleArrayInput
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings StorageBucketLoggingArrayInput
	Policy   pulumi.StringPtrInput
	// The secret key to use when applying changes. If omitted, `storageSecretKey` specified in provider config is used.
	SecretKey pulumi.StringPtrInput
	// A configuration of server-side encryption for the bucket (documented below)
	ServerSideEncryptionConfiguration StorageBucketServerSideEncryptionConfigurationPtrInput
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	Versioning StorageBucketVersioningPtrInput
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website StorageBucketWebsitePtrInput
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (StorageBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketState)(nil)).Elem()
}

type storageBucketArgs struct {
	// The access key to use when applying changes. If omitted, `storageAccessKey` specified in provider config is used.
	AccessKey *string `pulumi:"accessKey"`
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	Acl    *string `pulumi:"acl"`
	Bucket *string `pulumi:"bucket"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/cors/) (documented below).
	CorsRules []StorageBucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants []StorageBucketGrant `pulumi:"grants"`
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules []StorageBucketLifecycleRule `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings []StorageBucketLogging `pulumi:"loggings"`
	Policy   *string                `pulumi:"policy"`
	// The secret key to use when applying changes. If omitted, `storageSecretKey` specified in provider config is used.
	SecretKey *string `pulumi:"secretKey"`
	// A configuration of server-side encryption for the bucket (documented below)
	ServerSideEncryptionConfiguration *StorageBucketServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	Versioning *StorageBucketVersioning `pulumi:"versioning"`
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website *StorageBucketWebsite `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

// The set of arguments for constructing a StorageBucket resource.
type StorageBucketArgs struct {
	// The access key to use when applying changes. If omitted, `storageAccessKey` specified in provider config is used.
	AccessKey pulumi.StringPtrInput
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	Acl    pulumi.StringPtrInput
	Bucket pulumi.StringPtrInput
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/cors/) (documented below).
	CorsRules StorageBucketCorsRuleArrayInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants StorageBucketGrantArrayInput
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules StorageBucketLifecycleRuleArrayInput
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings StorageBucketLoggingArrayInput
	Policy   pulumi.StringPtrInput
	// The secret key to use when applying changes. If omitted, `storageSecretKey` specified in provider config is used.
	SecretKey pulumi.StringPtrInput
	// A configuration of server-side encryption for the bucket (documented below)
	ServerSideEncryptionConfiguration StorageBucketServerSideEncryptionConfigurationPtrInput
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	Versioning StorageBucketVersioningPtrInput
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website StorageBucketWebsitePtrInput
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (StorageBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketArgs)(nil)).Elem()
}

type StorageBucketInput interface {
	pulumi.Input

	ToStorageBucketOutput() StorageBucketOutput
	ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput
}

func (*StorageBucket) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBucket)(nil))
}

func (i *StorageBucket) ToStorageBucketOutput() StorageBucketOutput {
	return i.ToStorageBucketOutputWithContext(context.Background())
}

func (i *StorageBucket) ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketOutput)
}

func (i *StorageBucket) ToStorageBucketPtrOutput() StorageBucketPtrOutput {
	return i.ToStorageBucketPtrOutputWithContext(context.Background())
}

func (i *StorageBucket) ToStorageBucketPtrOutputWithContext(ctx context.Context) StorageBucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketPtrOutput)
}

type StorageBucketPtrInput interface {
	pulumi.Input

	ToStorageBucketPtrOutput() StorageBucketPtrOutput
	ToStorageBucketPtrOutputWithContext(ctx context.Context) StorageBucketPtrOutput
}

type storageBucketPtrType StorageBucketArgs

func (*storageBucketPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucket)(nil))
}

func (i *storageBucketPtrType) ToStorageBucketPtrOutput() StorageBucketPtrOutput {
	return i.ToStorageBucketPtrOutputWithContext(context.Background())
}

func (i *storageBucketPtrType) ToStorageBucketPtrOutputWithContext(ctx context.Context) StorageBucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketPtrOutput)
}

// StorageBucketArrayInput is an input type that accepts StorageBucketArray and StorageBucketArrayOutput values.
// You can construct a concrete instance of `StorageBucketArrayInput` via:
//
//          StorageBucketArray{ StorageBucketArgs{...} }
type StorageBucketArrayInput interface {
	pulumi.Input

	ToStorageBucketArrayOutput() StorageBucketArrayOutput
	ToStorageBucketArrayOutputWithContext(context.Context) StorageBucketArrayOutput
}

type StorageBucketArray []StorageBucketInput

func (StorageBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageBucket)(nil)).Elem()
}

func (i StorageBucketArray) ToStorageBucketArrayOutput() StorageBucketArrayOutput {
	return i.ToStorageBucketArrayOutputWithContext(context.Background())
}

func (i StorageBucketArray) ToStorageBucketArrayOutputWithContext(ctx context.Context) StorageBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketArrayOutput)
}

// StorageBucketMapInput is an input type that accepts StorageBucketMap and StorageBucketMapOutput values.
// You can construct a concrete instance of `StorageBucketMapInput` via:
//
//          StorageBucketMap{ "key": StorageBucketArgs{...} }
type StorageBucketMapInput interface {
	pulumi.Input

	ToStorageBucketMapOutput() StorageBucketMapOutput
	ToStorageBucketMapOutputWithContext(context.Context) StorageBucketMapOutput
}

type StorageBucketMap map[string]StorageBucketInput

func (StorageBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageBucket)(nil)).Elem()
}

func (i StorageBucketMap) ToStorageBucketMapOutput() StorageBucketMapOutput {
	return i.ToStorageBucketMapOutputWithContext(context.Background())
}

func (i StorageBucketMap) ToStorageBucketMapOutputWithContext(ctx context.Context) StorageBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketMapOutput)
}

type StorageBucketOutput struct{ *pulumi.OutputState }

func (StorageBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBucket)(nil))
}

func (o StorageBucketOutput) ToStorageBucketOutput() StorageBucketOutput {
	return o
}

func (o StorageBucketOutput) ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput {
	return o
}

func (o StorageBucketOutput) ToStorageBucketPtrOutput() StorageBucketPtrOutput {
	return o.ToStorageBucketPtrOutputWithContext(context.Background())
}

func (o StorageBucketOutput) ToStorageBucketPtrOutputWithContext(ctx context.Context) StorageBucketPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageBucket) *StorageBucket {
		return &v
	}).(StorageBucketPtrOutput)
}

type StorageBucketPtrOutput struct{ *pulumi.OutputState }

func (StorageBucketPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucket)(nil))
}

func (o StorageBucketPtrOutput) ToStorageBucketPtrOutput() StorageBucketPtrOutput {
	return o
}

func (o StorageBucketPtrOutput) ToStorageBucketPtrOutputWithContext(ctx context.Context) StorageBucketPtrOutput {
	return o
}

func (o StorageBucketPtrOutput) Elem() StorageBucketOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucket {
		if v != nil {
			return *v
		}
		var ret StorageBucket
		return ret
	}).(StorageBucketOutput)
}

type StorageBucketArrayOutput struct{ *pulumi.OutputState }

func (StorageBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageBucket)(nil))
}

func (o StorageBucketArrayOutput) ToStorageBucketArrayOutput() StorageBucketArrayOutput {
	return o
}

func (o StorageBucketArrayOutput) ToStorageBucketArrayOutputWithContext(ctx context.Context) StorageBucketArrayOutput {
	return o
}

func (o StorageBucketArrayOutput) Index(i pulumi.IntInput) StorageBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageBucket {
		return vs[0].([]StorageBucket)[vs[1].(int)]
	}).(StorageBucketOutput)
}

type StorageBucketMapOutput struct{ *pulumi.OutputState }

func (StorageBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageBucket)(nil))
}

func (o StorageBucketMapOutput) ToStorageBucketMapOutput() StorageBucketMapOutput {
	return o
}

func (o StorageBucketMapOutput) ToStorageBucketMapOutputWithContext(ctx context.Context) StorageBucketMapOutput {
	return o
}

func (o StorageBucketMapOutput) MapIndex(k pulumi.StringInput) StorageBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StorageBucket {
		return vs[0].(map[string]StorageBucket)[vs[1].(string)]
	}).(StorageBucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketInput)(nil)).Elem(), &StorageBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketPtrInput)(nil)).Elem(), &StorageBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketArrayInput)(nil)).Elem(), StorageBucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketMapInput)(nil)).Elem(), StorageBucketMap{})
	pulumi.RegisterOutputType(StorageBucketOutput{})
	pulumi.RegisterOutputType(StorageBucketPtrOutput{})
	pulumi.RegisterOutputType(StorageBucketArrayOutput{})
	pulumi.RegisterOutputType(StorageBucketMapOutput{})
}
