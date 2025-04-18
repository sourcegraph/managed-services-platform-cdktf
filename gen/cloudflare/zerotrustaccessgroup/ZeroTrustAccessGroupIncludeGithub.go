package zerotrustaccessgroup


type ZeroTrustAccessGroupIncludeGithub struct {
	// The ID of your Github identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_group#identity_provider_id ZeroTrustAccessGroup#identity_provider_id}
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
	// The name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_group#name ZeroTrustAccessGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The teams that should be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_group#teams ZeroTrustAccessGroup#teams}
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
}

