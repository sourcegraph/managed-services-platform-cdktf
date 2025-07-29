package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppCustomClaimsSource struct {
	// The name of the IdP claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A mapping from IdP ID to claim name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#name_by_idp ZeroTrustAccessApplication#name_by_idp}
	NameByIdp *map[string]*string `field:"optional" json:"nameByIdp" yaml:"nameByIdp"`
}

