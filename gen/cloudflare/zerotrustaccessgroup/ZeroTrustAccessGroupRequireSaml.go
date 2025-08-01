package zerotrustaccessgroup


type ZeroTrustAccessGroupRequireSaml struct {
	// The name of the SAML attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#attribute_name ZeroTrustAccessGroup#attribute_name}
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The SAML attribute value to look for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#attribute_value ZeroTrustAccessGroup#attribute_value}
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
	// The ID of your SAML identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#identity_provider_id ZeroTrustAccessGroup#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
}

