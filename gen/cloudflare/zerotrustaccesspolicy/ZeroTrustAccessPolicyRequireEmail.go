package zerotrustaccesspolicy


type ZeroTrustAccessPolicyRequireEmail struct {
	// The email of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#email ZeroTrustAccessPolicy#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

