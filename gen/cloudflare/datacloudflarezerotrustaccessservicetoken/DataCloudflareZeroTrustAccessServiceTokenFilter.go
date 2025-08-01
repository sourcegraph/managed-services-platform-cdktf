package datacloudflarezerotrustaccessservicetoken


type DataCloudflareZeroTrustAccessServiceTokenFilter struct {
	// The name of the service token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zero_trust_access_service_token#name DataCloudflareZeroTrustAccessServiceToken#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Search for service tokens by other listed query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zero_trust_access_service_token#search DataCloudflareZeroTrustAccessServiceToken#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

