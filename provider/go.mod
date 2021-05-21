module github.com/springload/pulumi-tf-provider-ec/provider

go 1.16

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0

require (
	github.com/elastic/terraform-provider-ec v0.1.1 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)
