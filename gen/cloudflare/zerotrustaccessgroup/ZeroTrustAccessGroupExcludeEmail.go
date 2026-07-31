package zerotrustaccessgroup


type ZeroTrustAccessGroupExcludeEmail struct {
	// The email of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_group#email ZeroTrustAccessGroup#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

