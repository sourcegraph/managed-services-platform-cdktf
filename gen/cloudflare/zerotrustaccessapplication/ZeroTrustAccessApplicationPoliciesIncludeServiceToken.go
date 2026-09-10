package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesIncludeServiceToken struct {
	// The ID of a Service Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#token_id ZeroTrustAccessApplication#token_id}
	TokenId *string `field:"required" json:"tokenId" yaml:"tokenId"`
}

