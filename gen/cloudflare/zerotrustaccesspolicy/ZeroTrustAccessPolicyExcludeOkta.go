package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeOkta struct {
	// The ID of your Okta identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#identity_provider_id ZeroTrustAccessPolicy#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
	// The name of the Okta group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#name ZeroTrustAccessPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

