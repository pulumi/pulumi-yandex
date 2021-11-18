module github.com/pulumi/pulumi-yandex/provider

go 1.16

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.10.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
	github.com/yandex-cloud/terraform-provider-yandex v0.66.0
)
