package zerotrustaccessgroup


type ZeroTrustAccessGroupExcludeGsuite struct {
	// The email of the Google Workspace group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#email ZeroTrustAccessGroup#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The ID of your Google Workspace identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#identity_provider_id ZeroTrustAccessGroup#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
}

