package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeSaml struct {
	// The name of the SAML attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#attribute_name ZeroTrustAccessPolicy#attribute_name}
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The SAML attribute value to look for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#attribute_value ZeroTrustAccessPolicy#attribute_value}
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
	// The ID of your SAML identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#identity_provider_id ZeroTrustAccessPolicy#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
}

