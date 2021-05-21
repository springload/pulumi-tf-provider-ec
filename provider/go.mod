module github.com/springload/pulumi-tf-provider-ec/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
)

require (
	github.com/elastic/terraform-provider-ec v0.1.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.1.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)
