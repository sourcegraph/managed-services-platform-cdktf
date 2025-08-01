package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppCustomAttributes struct {
	// The SAML FriendlyName of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#friendly_name ZeroTrustAccessApplication#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// The name of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A globally unique name for an identity or service provider. Available values: "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified", "urn:oasis:names:tc:SAML:2.0:attrname-format:basic", "urn:oasis:names:tc:SAML:2.0:attrname-format:uri".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#name_format ZeroTrustAccessApplication#name_format}
	NameFormat *string `field:"optional" json:"nameFormat" yaml:"nameFormat"`
	// If the attribute is required when building a SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#required ZeroTrustAccessApplication#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#source ZeroTrustAccessApplication#source}.
	Source *ZeroTrustAccessApplicationSaasAppCustomAttributesSource `field:"optional" json:"source" yaml:"source"`
}

