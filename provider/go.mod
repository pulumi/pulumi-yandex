module github.com/pulumi/pulumi-yandex/provider

go 1.16

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	cloud.google.com/go/kms v1.1.0 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.16.0
	github.com/pulumi/pulumi/sdk/v3 v3.22.0
	github.com/yandex-cloud/terraform-provider-yandex v0.70.0
)
