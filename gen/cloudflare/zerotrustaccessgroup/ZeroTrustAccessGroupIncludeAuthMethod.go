package zerotrustaccessgroup


type ZeroTrustAccessGroupIncludeAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#auth_method ZeroTrustAccessGroup#auth_method}
	AuthMethod *string `field:"required" json:"authMethod" yaml:"authMethod"`
}

