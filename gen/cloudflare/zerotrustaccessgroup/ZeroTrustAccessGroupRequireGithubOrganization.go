package zerotrustaccessgroup


type ZeroTrustAccessGroupRequireGithubOrganization struct {
	// The ID of your Github identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#identity_provider_id ZeroTrustAccessGroup#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
	// The name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#name ZeroTrustAccessGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#team ZeroTrustAccessGroup#team}
	Team *string `field:"optional" json:"team" yaml:"team"`
}

