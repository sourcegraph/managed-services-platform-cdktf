package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesRequireEmail struct {
	// The email of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#email ZeroTrustAccessApplication#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

