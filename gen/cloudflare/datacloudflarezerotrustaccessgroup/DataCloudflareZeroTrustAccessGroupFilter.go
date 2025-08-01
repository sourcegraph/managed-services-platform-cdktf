package datacloudflarezerotrustaccessgroup


type DataCloudflareZeroTrustAccessGroupFilter struct {
	// The name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zero_trust_access_group#name DataCloudflareZeroTrustAccessGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Search for groups by other listed query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zero_trust_access_group#search DataCloudflareZeroTrustAccessGroup#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

