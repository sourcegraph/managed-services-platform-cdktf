package zerotrustaccesspolicy


type ZeroTrustAccessPolicyIncludeOkta struct {
	// The ID of your Okta identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#identity_provider_id ZeroTrustAccessPolicy#identity_provider_id}
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
	// The name of the Okta Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#name ZeroTrustAccessPolicy#name}
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}

