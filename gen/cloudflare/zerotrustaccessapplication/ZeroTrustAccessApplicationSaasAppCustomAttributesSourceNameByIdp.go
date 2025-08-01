package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdp struct {
	// The UID of the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#idp_id ZeroTrustAccessApplication#idp_id}
	IdpId *string `field:"optional" json:"idpId" yaml:"idpId"`
	// The name of the IdP provided attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#source_name ZeroTrustAccessApplication#source_name}
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
}

