package zerotrustaccessgroup


type ZeroTrustAccessGroupIncludeServiceToken struct {
	// The ID of a Service Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#token_id ZeroTrustAccessGroup#token_id}
	TokenId *string `field:"required" json:"tokenId" yaml:"tokenId"`
}

