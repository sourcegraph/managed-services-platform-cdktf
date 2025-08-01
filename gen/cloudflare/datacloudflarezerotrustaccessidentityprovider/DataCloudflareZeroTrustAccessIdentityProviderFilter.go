package datacloudflarezerotrustaccessidentityprovider


type DataCloudflareZeroTrustAccessIdentityProviderFilter struct {
	// Indicates to Access to only retrieve identity providers that have the System for Cross-Domain Identity Management (SCIM) enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zero_trust_access_identity_provider#scim_enabled DataCloudflareZeroTrustAccessIdentityProvider#scim_enabled}
	ScimEnabled *string `field:"optional" json:"scimEnabled" yaml:"scimEnabled"`
}

